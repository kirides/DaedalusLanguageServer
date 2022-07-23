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
	"github.com/kirides/DaedalusLanguageServer/javadoc"
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

func getTypeFieldsAsCompletionItems(docs *parseResultsManager, symbolName string, locals map[string]symbol.Symbol) ([]lsp.CompletionItem, error) {
	symName := strings.ToUpper(symbolName)
	sym, ok := locals[symName]
	if !ok {
		sym, ok = docs.LookupGlobalSymbol(symName, SymbolVariable|SymbolClass|SymbolInstance|SymbolPrototype)
	}

	if !ok {
		return []lsp.CompletionItem{}, nil
	}
	if clsSym, ok := sym.(symbol.Class); ok {
		return fieldsToCompletionItems(docs, clsSym.Fields), nil
	}
	if protoSym, ok := sym.(symbol.ProtoTypeOrInstance); ok {
		return getTypeFieldsAsCompletionItems(docs, protoSym.Parent, locals)
	}
	if varSym, ok := sym.(symbol.Variable); ok {
		return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	}
	// no way for e.g. C_NPC arrays
	// if varSym, ok := sym.(ArrayVariableSymbol); ok {
	// 	return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	// }
	return []lsp.CompletionItem{}, nil
}

func (h *LspHandler) getTextDocumentCompletion(params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(params.TextDocument.URI))
	if err == nil {
		doc := h.bufferManager.GetBuffer(h.uriToFilename(params.TextDocument.URI))
		di := symbol.DefinitionIndex{Line: int(params.Position.Line), Column: int(params.Position.Character)}

		scci, err := getSignatureCompletions(params, h)
		if err != nil {
			h.logger.Debugf("signature completion error %v: ", err)
		}
		if len(scci) > 0 {
			return scci, nil
		}

		// dot-completion
		proto, _, err := doc.GetParentSymbolReference(params.Position)
		if err == nil && proto != "" {
			locals := parsedDoc.ScopedVariables(di)
			return getTypeFieldsAsCompletionItems(h.parsedDocuments, proto, locals)
		}

		// locally scoped variables ordered at the top
		localSortIdx := 0
		for _, fn := range parsedDoc.Functions {
			if fn.BodyDefinition.InBBox(di) {
				for _, p := range fn.Parameters {
					ci, err := completionItemFromSymbol(h.parsedDocuments, p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (parameter)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				for _, p := range fn.LocalVariables {
					ci, err := completionItemFromSymbol(h.parsedDocuments, p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (local)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				break
			}
		}
		for _, fn := range parsedDoc.Instances {
			if fn.BodyDefinition.InBBox(di) {
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
		for _, fn := range parsedDoc.Prototypes {
			if fn.BodyDefinition.InBBox(di) {
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
	}

	h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		ci, err := completionItemFromSymbol(h.parsedDocuments, s)
		if err != nil {
			return nil
		}
		result = append(result, ci)
		return nil
	}, SymbolAll)

	return result, nil
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

func (h *LspHandler) handleSignatureInfo(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.SignatureHelp, error) {
	fnCtx, err := getFunctionCallContext(h, params.TextDocument.URI, params.Position)
	if err != nil {
		return lsp.SignatureHelp{}, err
	}

	fn := fnCtx.Function

	var fnParams []lsp.ParameterInformation
	for _, p := range fn.Parameters {
		doc := javadoc.FindParam(fn.Documentation(), p.Name())
		var mdDoc interface{}
		if doc != "" {
			mdDoc = &lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: fmt.Sprintf("**%s** - *%s*", p.Name(), javadoc.RemoveTokens(doc)),
			}
		}
		fnParams = append(fnParams, lsp.ParameterInformation{
			Label:         p.String(),
			Documentation: mdDoc,
		})
	}

	return lsp.SignatureHelp{
		Signatures: []lsp.SignatureInformation{
			{
				Documentation: &lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: javadoc.MarkdownSimple(fn),
				},
				Label:      fn.String(),
				Parameters: fnParams,
			},
		},
		ActiveParameter: uint32(fnCtx.ParamIdx),
		ActiveSignature: 0,
	}, nil
}

func getSymbolLocation(symbol symbol.Symbol) lsp.Location {
	return lsp.Location{
		URI: uri.File(symbol.Source()),
		Range: lsp.Range{
			Start: lsp.Position{
				Character: uint32(symbol.Definition().Start.Column),
				Line:      uint32(symbol.Definition().Start.Line - 1),
			},
			End: lsp.Position{
				Character: uint32(symbol.Definition().Start.Column + len(symbol.Name())),
				Line:      uint32(symbol.Definition().Start.Line - 1),
			},
		}}
}

func (h *LspHandler) handleGoToDefinition(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.Location, error) {
	symbol, err := h.lookUpSymbol(h.uriToFilename(params.TextDocument.URI), params.Position)
	if err != nil {
		return lsp.Location{}, err
	}

	return getSymbolLocation(symbol), nil
}

func (h *LspHandler) handleTextDocumentCompletion(req dls.RpcContext, data lsp.CompletionParams) error {
	items, err := h.getTextDocumentCompletion(&data)
	req.ReplyEither(req.Context(), items, err)
	return nil
}

func (h *LspHandler) handleTextDocumentDefinition(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.handleGoToDefinition(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), found, nil)
}

func (h *LspHandler) handleTextDocumentHover(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.lookUpSymbol(h.uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	h.LogDebug("Found Symbol for Hover: %s", found.String())

	return req.Reply(req.Context(), lsp.Hover{
		Range: &lsp.Range{
			Start: data.Position,
			End:   data.Position,
		},
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: strings.TrimSpace(javadoc.MarkdownSimple(found) + "\n\n```daedalus\n" + h.parsedDocuments.getSymbolCode(found) + "\n```"),
		},
	}, nil)
}
func (h *LspHandler) handleTextDocumentSignatureHelp(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	result, err := h.handleSignatureInfo(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), result, nil)
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

func getDocumentSymbol(s symbol.Symbol) lsp.DocumentSymbol {
	rn := getSymbolLocation(s).Range
	return lsp.DocumentSymbol{
		Name:           s.Name(),
		Kind:           getSymbolKind(s),
		Range:          rn,
		SelectionRange: rn,
	}
}

func getSymbolInformation(s symbol.Symbol) lsp.SymbolInformation {
	return lsp.SymbolInformation{
		Name:     s.Name(),
		Kind:     getSymbolKind(s),
		Location: getSymbolLocation(s),
	}
}

func collectDocumentSymbols(result []lsp.DocumentSymbol, s symbol.Symbol) []lsp.DocumentSymbol {
	mainSymb := getDocumentSymbol(s)

	if cls, ok := s.(symbol.Class); ok {
		for _, v := range cls.Fields {
			si := getDocumentSymbol(v)
			mainSymb.Children = append(mainSymb.Children, si)
		}
		mainSymb.Range = lsp.Range{
			Start: mainSymb.Range.Start,
			End: lsp.Position{
				Line:      uint32(cls.BodyDefinition.End.Line) - 1,
				Character: uint32(cls.BodyDefinition.End.Column),
			},
		}
	} else if cls, ok := s.(symbol.ProtoTypeOrInstance); ok {
		for _, v := range cls.Fields {
			si := getDocumentSymbol(v)
			mainSymb.Children = append(mainSymb.Children, si)
		}
		mainSymb.Range = lsp.Range{
			Start: mainSymb.Range.Start,
			End: lsp.Position{
				Line:      uint32(cls.BodyDefinition.End.Line) - 1,
				Character: uint32(cls.BodyDefinition.End.Column),
			},
		}
	} else if cls, ok := s.(symbol.Function); ok {
		mainSymb.Range = lsp.Range{
			Start: mainSymb.Range.Start,
			End: lsp.Position{
				Line:      uint32(cls.BodyDefinition.End.Line) - 1,
				Character: uint32(cls.BodyDefinition.End.Column),
			},
		}
	}

	result = append(result, mainSymb)
	return result
}

func collectWorkspaceSymbols(result []lsp.SymbolInformation, s symbol.Symbol) []lsp.SymbolInformation {
	mainSymb := getSymbolInformation(s)
	result = append(result, mainSymb)

	if cls, ok := s.(symbol.Class); ok {
		for _, v := range cls.Fields {
			si := getSymbolInformation(v)
			si.ContainerName = s.Name()
			result = append(result, si)
		}
	} else if cls, ok := s.(symbol.ProtoTypeOrInstance); ok {
		for _, v := range cls.Fields {
			si := getSymbolInformation(v)
			si.ContainerName = s.Name()
			result = append(result, si)
		}
	}
	return result
}

func (h *LspHandler) handleDocumentSymbol(req dls.RpcContext, params lsp.DocumentSymbolParams) error {
	r, err := h.parsedDocuments.Get(h.uriToFilename(params.TextDocument.URI))
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}
	numSymbols := r.CountSymbols()
	result := make([]lsp.DocumentSymbol, 0, numSymbols)

	err = r.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if req.Context().Err() != nil {
			h.logger.Debugf("request cancelled", "method", req.Request().Method())
			return req.Context().Err()
		}
		result = collectDocumentSymbols(result, s)
		return nil
	}, SymbolAll)
	if err != nil {
		return nil
	}

	req.Reply(req.Context(), result, nil)
	return nil
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

func (h *LspHandler) handleWorkspaceSymbol(req dls.RpcContext, params lsp.WorkspaceSymbolParams) error {
	numSymbols := h.parsedDocuments.CountSymbols()
	result := make([]lsp.SymbolInformation, 0, numSymbols)
	buffer := make([]lsp.SymbolInformation, 0, 50)

	qlower := strings.ToLower(params.Query)

	err := h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if req.Context().Err() != nil {
			h.logger.Debugf("request cancelled", "method", req.Request().Method())
			return req.Context().Err()
		}
		if qlower == "" {
			result = collectWorkspaceSymbols(result, s)
			return nil
		}

		// pre filtering
		buffer = buffer[:0]
		buffer = collectWorkspaceSymbols(buffer, s)
		for _, v := range buffer {
			if stringContainsAllAnywhere(v.Name, params.Query) {
				result = append(result, v)
			}
		}
		return nil
	}, SymbolAll)

	if err != nil {
		return nil
	}

	req.Reply(req.Context(), result, nil)
	return nil
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
