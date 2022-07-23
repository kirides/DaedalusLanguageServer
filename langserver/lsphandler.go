package langserver

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"go.lsp.dev/jsonrpc2"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

var defaultProjectFiles = []string{"Gothic.src", "Camera.src", "Menu.src", "Music.src", "ParticleFX.src", "SFX.src", "VisualFX.src"}

type LspConfig struct {
	FileEncoding     string   `json:"fileEncoding"`
	SrcFileEncoding  string   `json:"srcFileEncoding"`
	LogLevel         string   `json:"loglevel"`
	PprofAddr        string   `json:"pprofAddr"`
	NumParserThreads int      `json:"numParserThreads"`
	ProjectFiles     []string `json:"projectFiles"`
}

// LspHandler ...
type LspHandler struct {
	TextDocumentSync *textDocumentSync
	bufferManager    *BufferManager
	parsedDocuments  *parseResultsManager
	handlers         *dls.RpcMux
	config           LspConfig
	onceParseAll     sync.Once

	onConfigChangedHandlers []func(config LspConfig)
	parsedKnownSrcFiles     concurrentSet[string]
	baseLspHandler
	initialized bool
}

var (
	// ErrWalkAbort should be returned if a walk function should abort early
	ErrWalkAbort = fmt.Errorf("OK")
)

// NewLspHandler ...
func NewLspHandler(conn jsonrpc2.Conn, logger dls.Logger) *LspHandler {
	bufferManager := NewBufferManager()
	parsedDocuments := newParseResultsManager(logger)
	return &LspHandler{
		baseLspHandler: baseLspHandler{
			logger: logger,
			conn:   conn,
		},
		handlers:        dls.NewMux(),
		initialized:     false,
		bufferManager:   bufferManager,
		parsedDocuments: parsedDocuments,
		TextDocumentSync: &textDocumentSync{
			baseLspHandler: baseLspHandler{
				logger: logger,
				conn:   conn,
			},
			bufferManager:   bufferManager,
			parsedDocuments: parsedDocuments,
		},
		onConfigChangedHandlers: []func(config LspConfig){},
		config: LspConfig{
			FileEncoding:    "1252",
			SrcFileEncoding: "1252",
			LogLevel:        "info",
			ProjectFiles:    defaultProjectFiles,
		},
	}
}

func (h *LspHandler) OnConfigChanged(handler func(config LspConfig)) {
	h.onConfigChangedHandlers = append(h.onConfigChangedHandlers, handler)
}

func (h *LspHandler) lookUpScopedSymbol(documentURI, identifier string, position lsp.Position) symbol.Symbol {
	p, err := h.parsedDocuments.Get(documentURI)
	if err != nil {
		return nil
	}

	di := symbol.DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}

	var sym symbol.Symbol
	p.WalkScopedVariables(di, func(s symbol.Symbol) bool {
		if strings.EqualFold(s.Name(), identifier) {
			sym = s
			return false
		}
		return true
	})
	return sym
}

func (h *LspHandler) lookUpSymbol(documentURI string, position lsp.Position) (symbol.Symbol, error) {
	doc := h.bufferManager.GetBuffer(documentURI)
	if doc == "" {
		return nil, fmt.Errorf("document %q not found", documentURI)
	}
	identifier := doc.GetWordRangeAtPosition(position)

	if v := h.lookUpScopedSymbol(documentURI, identifier, position); v != nil {
		return v, nil
	}

	symbol, found := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(identifier), SymbolAll)

	if !found {
		return nil, fmt.Errorf("identifier %q not found", identifier)
	}

	return symbol, nil
}

func getSymbolKind(s symbol.Symbol) lsp.SymbolKind {
	switch s.(type) {
	case symbol.ArrayVariable, symbol.ConstantArray:
		return lsp.SymbolKindArray
	case symbol.Class, symbol.ProtoTypeOrInstance:
		return lsp.SymbolKindClass
	case symbol.Function:
		return lsp.SymbolKindFunction
	case symbol.Constant:
		return lsp.SymbolKindConstant
	case symbol.Variable:
		return lsp.SymbolKindVariable
	}
	return lsp.SymbolKindNull
}

func stringContainsAllAnywhere(value, set string) bool {
	found := true

	for _, v := range set {
		if !(strings.ContainsRune(value, unicode.ToLower(v)) ||
			strings.ContainsRune(value, unicode.ToUpper(v))) {
			found = false
			break
		}
	}

	return found
}

func (h *LspHandler) onInitialized() {
	h.handlers.Register(lsp.MethodTextDocumentCompletion, dls.MakeHandler(h.handleTextDocumentCompletion))
	h.handlers.Register(lsp.MethodTextDocumentDefinition, dls.MakeHandler(h.handleTextDocumentDefinition))
	h.handlers.Register(lsp.MethodTextDocumentHover, dls.MakeHandler(h.handleTextDocumentHover))
	h.handlers.Register(lsp.MethodTextDocumentSignatureHelp, dls.MakeHandler(h.handleTextDocumentSignatureHelp))

	// textDocument/didOpen/didSave/didChange
	h.handlers.Register(lsp.MethodTextDocumentDidOpen, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidOpen))
	h.handlers.Register(lsp.MethodTextDocumentDidChange, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidChange))
	h.handlers.Register(lsp.MethodTextDocumentDidSave, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidSave))

	h.handlers.Register(lsp.MethodTextDocumentDocumentSymbol, dls.MakeHandler(h.handleDocumentSymbol))
	h.handlers.Register(lsp.MethodWorkspaceSymbol, dls.MakeHandler(h.handleWorkspaceSymbol))
}

func prettyJSON(val interface{}) string {
	v, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(v)
}

// Deliver ...
func (h *LspHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, r jsonrpc2.Request) error {
	h.LogDebug("Requested %q", r.Method())

	// if r.Params != nil {
	// 	var paramsMap map[string]interface{}
	// 	json.Unmarshal(*r.Params, &paramsMap)
	// 	fmt.Fprintf(os.Stderr, "Params: %+v\n", paramsMap)
	// }
	switch r.Method() {
	case lsp.MethodInitialize:
		if err := reply(ctx, lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				CompletionProvider: &lsp.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
				DefinitionProvider: true,
				HoverProvider:      true,
				SignatureHelpProvider: &lsp.SignatureHelpOptions{
					TriggerCharacters: []string{"(", ","},
				},
				TextDocumentSync: lsp.TextDocumentSyncOptions{
					Change:    lsp.TextDocumentSyncKindFull,
					OpenClose: true,
					Save: &lsp.SaveOptions{
						IncludeText: true,
					},
				},
				WorkspaceSymbolProvider: true,
				DocumentSymbolProvider:  true,
			},
		}, nil); err != nil {
			return fmt.Errorf("not initialized")
		}
		h.initialized = true
		h.onInitialized()
		return nil
	case lsp.MethodWorkspaceDidChangeConfiguration:
		var params struct {
			Settings struct {
				DaedalusLanguageServer LspConfig `json:"daedalusLanguageServer"`
			} `json:"settings"`
		}

		var configRaw map[string]interface{}
		_ = json.Unmarshal(r.Params(), &configRaw)
		h.LogDebug("%s (debug): %v", r.Method(), prettyJSON(configRaw))

		_ = json.Unmarshal(r.Params(), &params)
		h.config = params.Settings.DaedalusLanguageServer
		h.LogInfo("%s: %v", r.Method(), prettyJSON(configRaw))

		h.parsedDocuments.SetFileEncoding(h.config.FileEncoding)
		h.parsedDocuments.SetSrcEncoding(h.config.SrcFileEncoding)
		h.parsedDocuments.NumParserThreads = h.config.NumParserThreads

		for _, v := range h.onConfigChangedHandlers {
			v(h.config)
		}

		if len(h.config.ProjectFiles) == 0 {
			h.config.ProjectFiles = defaultProjectFiles
		}

		return nil
	case lsp.MethodInitialized:
		return nil
	}

	// DEFAULT / OTHERWISE

	if !h.initialized {
		reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.ServerNotInitialized, "Not initialized yet"))
		return fmt.Errorf("not initialized yet")
	}

	// Recover if something bad happens in the handlers...
	defer func() {
		err := recover()
		if err != nil {
			h.LogWarn("Recovered from panic at %s: %v\n", r.Method, err)
		}
	}()

	if r.Method() == lsp.MethodTextDocumentDidOpen {
		h.onceParseAll.Do(func() {
			exe, _ := os.Executable()
			if f, err := findPath(filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src")); err == nil {
				_, err = h.parsedDocuments.ParseSource(f)
				if err != nil {
					h.LogError("Error parsing %q: %v", f, err)
					return
				}
			}
		})
		var openParams lsp.DidOpenTextDocumentParams
		json.Unmarshal(r.Params(), &openParams)
		go func() {
			wd := h.uriToFilename(openParams.TextDocument.URI)
			if wd == "" {
				h.LogError("Error locating current file")
				return
			}

			var resultsX []*ParseResult
			if externalsSrc, err := findPathAnywhereUpToRoot(wd, filepath.Join("_externals", "externals.src")); err == nil {
				if !h.parsedKnownSrcFiles.Contains(externalsSrc) {
					h.parsedKnownSrcFiles.Store(externalsSrc)
					customBuiltins, err := h.parsedDocuments.ParseSource(externalsSrc)
					if err != nil {
						h.LogError("Error parsing %q: %v", externalsSrc, err)
					} else {
						resultsX = append(resultsX, customBuiltins...)
					}
				}
			} else if externalsDaedalus, err := findPathAnywhereUpToRoot(wd, filepath.Join("_externals", "externals.d")); err == nil {
				if !h.parsedKnownSrcFiles.Contains(externalsDaedalus) {
					h.parsedKnownSrcFiles.Store(externalsDaedalus)
					parsed, err := h.parsedDocuments.ParseFile(externalsDaedalus)
					if err != nil {
						h.LogError("Error parsing %q: %v", externalsDaedalus, err)
					} else {
						resultsX = append(resultsX, parsed)
					}
				}
			}

			for _, v := range h.config.ProjectFiles {
				var err error
				full := v
				if filepath.IsAbs(full) || strings.ContainsAny(full, "\\/") {
					// User defined hardcoded project file, either absolute or relative
					if _, err := os.Stat(full); err != nil {
						h.LogError("Error user-define project file %s: %v", full, err)
						continue
					}
				} else {
					full, err = findPathAnywhereUpToRoot(wd, v)
				}
				if err != nil {
					h.LogDebug("Did not parse %q: %v", v, err)
					continue
				}
				if h.parsedKnownSrcFiles.Contains(full) {
					continue
				}
				h.parsedKnownSrcFiles.Store(full)
				results, err := h.parsedDocuments.ParseSource(full)
				if err != nil {
					h.LogError("Error parsing %s: %v", full, err)
					continue
				}
				resultsX = append(resultsX, results...)
			}

			var diagnostics = make([]lsp.Diagnostic, 0)
			for _, p := range resultsX {
				if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
					diagnostics = diagnostics[:0]

					for _, se := range p.SyntaxErrors {
						diag := se.Diagnostic()

						// TODO: sometimes syntax errors are in here twice...
						add := true
						for _, v := range diagnostics {
							if v.Range == diag.Range {
								add = false
								break
							}
						}
						if !add {
							continue
						}
						diagnostics = append(diagnostics, diag)
					}
					h.LogInfo("> %s", p.Source)
					h.conn.Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
						URI:         lsp.DocumentURI(uri.File(p.Source)),
						Diagnostics: diagnostics,
					})
				}
			}
		}()
	}

	handled, err := h.handlers.Handle(ctx, reply, r)
	if err != nil && handled {
		return err
	}
	if handled {
		return nil
	}
	return h.baseLspHandler.Handle(ctx, reply, r)
}
