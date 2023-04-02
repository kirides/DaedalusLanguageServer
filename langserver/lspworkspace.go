package langserver

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

type LspWorkspace struct {
	path            string
	uri             lsp.DocumentURI
	logger          dls.Logger
	bufferManager   *BufferManager
	parsedDocuments *parseResultsManager
	config          LspConfig
	onceParseAll    sync.Once
	conn            jsonrpc2.Conn

	parsedKnownSrcFiles  concurrentSet[string]
	workspaceInitialized bool

	workspaceCtx       context.Context
	cancelWorkspaceCtx context.CancelFunc
}

// TODO: We need to figure out a better way to handle workspaces.
// We might have to separate the parsers for Gothic/Music/SFX/VFX/...
func (ws *LspWorkspace) tryInitializeWorkspace(ctx context.Context, params *lsp.DidOpenTextDocumentParams, config LspConfig) {
	if ws.workspaceInitialized {
		return
	}
	wd := uriToFilename(params.TextDocument.URI)
	if wd == "" {
		ws.logger.Errorf("Error locating current file")
		return
	}

	exe, _ := os.Executable()
	if f, err := findPath(filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src")); err == nil {
		_, err = ws.parsedDocuments.ParseSource(ws.workspaceCtx, f)
		if err != nil {
			ws.logger.Errorf("Error parsing %q: %v", f, err)
			return
		}
	}

	if _, err := findPath(filepath.Join(ws.path, ".dls", "externals", "builtins.src")); err != nil {
		req := DlsMessageRequest{
			ShowMessageRequestParams: lsp.ShowMessageRequestParams{
				Type:    lsp.Info,
				Message: "Create neccessery workspace files",
				Actions: []lsp.MessageActionItem{
					{Title: "Cancel"},
					{Title: "Gothic 1"},
					{Title: "Gothic 2"},
				},
			},
			WorkspaceURI: ws.uri,
		}
		var result lsp.MessageActionItem
		id, err := ws.conn.Call(context.Background(), lsp.MethodWindowShowMessageRequest, req, &result)
		if err != nil {
			ws.logger.Errorf("Error requesting message %q: %v", id, err)
		}
		ws.logger.Debugf("Result: %#v", result)
		switch result.Title {
		case "Gothic 2":
			ws.commandSetupWorkspace(ws, "G2A")
		case "Gothic 1":
			ws.commandSetupWorkspace(ws, "G1")
		}
	}

	if f, err := findPath(filepath.Join(ws.path, ".dls", "externals", "builtins.src")); err == nil {
		_, err = ws.parsedDocuments.ParseSource(ws.workspaceCtx, f)
		if err != nil {
			ws.logger.Errorf("Error parsing %q: %v", f, err)
			return
		}
	}

	// Try to locate a workspace file
	foundProjectFile := false
	for _, v := range config.ProjectFiles {
		if foundProjectFile {
			break
		}

		full := v
		if filepath.IsAbs(full) || strings.ContainsAny(full, "\\/") {
			if _, err := os.Stat(full); err == nil {
				ws.logger.Errorf("Error user-define project file %s: %v", full, err)
				foundProjectFile = true
			}
		} else {
			_, err := findPathAnywhereUpToRoot(wd, v)
			if err == nil {
				foundProjectFile = true
			} else {
				for _, knownDir := range []string{"SYSTEM", "CONTENT"} {
					if _, err := findPathAnywhereUpToRoot(wd, filepath.Join(knownDir, v)); err == nil {
						foundProjectFile = true
						break
					}
				}
			}
		}
	}
	if !foundProjectFile {
		ws.logger.Debugf("No project file found, deferring workspace initialization")
		return
	}
	ws.workspaceInitialized = true

	ws.onceParseAll.Do(func() {
		var resultsX []*ParseResult
		if externalsSrc, err := findPathAnywhereUpToRoot(wd, filepath.Join("_externals", "externals.src")); err == nil {
			if !ws.parsedKnownSrcFiles.Contains(externalsSrc) {
				ws.parsedKnownSrcFiles.Store(externalsSrc)
				customBuiltins, err := ws.parsedDocuments.ParseSource(ws.workspaceCtx, externalsSrc)
				if err != nil {
					ws.logger.Errorf("Error parsing %q: %v", externalsSrc, err)
				} else {
					resultsX = append(resultsX, customBuiltins...)
				}
			}
		} else if externalsDaedalus, err := findPathAnywhereUpToRoot(wd, filepath.Join("_externals", "externals.d")); err == nil {
			if !ws.parsedKnownSrcFiles.Contains(externalsDaedalus) {
				ws.parsedKnownSrcFiles.Store(externalsDaedalus)
				parsed, err := ws.parsedDocuments.ParseFile(externalsDaedalus)
				if err != nil {
					ws.logger.Errorf("Error parsing %q: %v", externalsDaedalus, err)
				} else {
					resultsX = append(resultsX, parsed)
				}
			}
		}

		for _, v := range config.ProjectFiles {
			var err error
			full := v
			if filepath.IsAbs(full) || strings.ContainsAny(full, "\\/") {

				if _, err := os.Stat(full); err != nil {
					ws.logger.Errorf("Error user-define project file %s: %v", full, err)
					continue
				}
			} else {
				full, err = findPathAnywhereUpToRoot(wd, v)
				if err != nil {
					for _, knownDir := range []string{"SYSTEM", "CONTENT"} {
						full, err = findPathAnywhereUpToRoot(wd, filepath.Join(knownDir, v))
						if err == nil {
							ws.logger.Infof("Found %s at %s", v, full)
							break
						}
					}
				}
			}
			if err != nil {
				ws.logger.Debugf("Did not parse %q: %v", v, err)
				continue
			}
			if ws.parsedKnownSrcFiles.Contains(full) {
				continue
			}
			ws.parsedKnownSrcFiles.Store(full)
			results, err := ws.parsedDocuments.ParseSource(ws.workspaceCtx, full)
			if err != nil {
				ws.logger.Errorf("Error parsing %s: %v", full, err)
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
				ws.logger.Infof("> %s", p.Source)
				ws.conn.Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
					URI:         lsp.DocumentURI(uri.File(p.Source)),
					Diagnostics: diagnostics,
				})
			}
		}
		ws.logger.Infof("Initial diagnostics completed")
	})
}

func (h *LspWorkspace) lookUpScope(documentURI string, position lsp.Position) (symbol.Symbol, bool) {
	p, err := h.parsedDocuments.Get(documentURI)
	if err != nil {
		return nil, false
	}

	di := symbol.DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}

	return p.FindContainingScope(di)
}

func (h *LspWorkspace) getEffectiveValue(sym symbol.ProtoTypeOrInstance, field string) (symbol.Constant, bool) {
	var own symbol.Constant
	found := false
	for _, v := range sym.Fields {
		if strings.EqualFold(v.Name(), field) {
			own = v
			found = true
			break
		}
	}

	parent, ok := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(sym.Parent), SymbolPrototype|SymbolClass)
	if ok {
		switch p := parent.(type) {
		case symbol.Class:
			val := ""
			for _, v := range p.Fields {
				if strings.EqualFold(v.Name(), field) {
					getType, ok := v.(interface{ GetType() string })
					if ok {
						typ := getType.GetType()
						switch {
						case strings.EqualFold(typ, "int"):
							val = "0"
							typ = "int"
						case strings.EqualFold(typ, "float"):
							val = "0.0"
							typ = "float"
						case strings.EqualFold(typ, "string"):
							val = ""
							typ = "string"
						}
						if found {
							// we have the value, just use it for finding the type
							return symbol.NewConstant(v.Name(), typ, own.Source(), own.Documentation(), own.Definition(), own.Value), true
						}
						return symbol.NewConstant(v.Name(), typ, v.Source(), v.Documentation(), v.Definition(), val), true
					}
					break
				}
			}
		case symbol.ProtoTypeOrInstance:
			v, ok := h.getEffectiveValue(p, field)
			if !ok {
				return v, ok
			}

			if !found {
				return v, true
			}
			own.Type = v.Type // fix type information, should be brought up from class->proto->inst
		}
	}

	return own, found
}

type FoundLocation int

func (l FoundLocation) String() string {
	switch l {
	case FoundGlobal:
		return "global"
	case FoundParameter:
		return "parameter"
	case FoundLocal:
		return "local"
	case FoundField:
		return "field"
	}
	return "-"
}

const (
	// Global symbol
	FoundGlobal FoundLocation = iota
	// Function parameter symbol
	FoundParameter
	// Local symbol
	FoundLocal
	// Instance/prototype/class field symbol
	FoundField
)

type FoundSymbol struct {
	Symbol   symbol.Symbol
	Location FoundLocation
}

func (h *LspWorkspace) lookUpSymbol(documentURI string, position lsp.Position) (FoundSymbol, error) {
	var notFound FoundSymbol
	doc := h.bufferManager.GetBuffer(documentURI)
	if doc == "" {
		return notFound, fmt.Errorf("document %q not found", documentURI)
	}
	identifier, _ := doc.GetWordRangeAtPosition(position)

	if v, ok := h.lookUpScope(documentURI, position); ok {
		var ok bool
		switch s := v.(type) {
		case symbol.Function:
			for _, v := range s.Parameters {
				if strings.EqualFold(v.Name(), identifier) {
					return FoundSymbol{
						Symbol:   v,
						Location: FoundParameter,
					}, nil
				}
			}
			for _, v := range s.LocalVariables {
				if strings.EqualFold(v.Name(), identifier) {
					return FoundSymbol{
						Symbol:   v,
						Location: FoundLocal,
					}, nil
				}
			}
		case symbol.ProtoTypeOrInstance:
			v, ok = h.getEffectiveValue(s, identifier)
			if ok {
				return FoundSymbol{
					Symbol:   v,
					Location: FoundField,
				}, nil
			}
		}
	}

	symbol, found := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(identifier), SymbolAll)

	if !found {
		return notFound, fmt.Errorf("identifier %q not found", identifier)
	}
	return FoundSymbol{
		Symbol:   symbol,
		Location: FoundGlobal,
	}, nil
}

func getSymbolKind(s symbol.Symbol) lsp.SymbolKind {
	switch s.(type) {
	case symbol.ArrayVariable, symbol.ConstantArray:
		return lsp.Array
	case symbol.Class, symbol.ProtoTypeOrInstance:
		return lsp.Class
	case symbol.Function:
		return lsp.Function
	case symbol.Constant:
		return lsp.Constant
	case symbol.Variable:
		return lsp.Variable
	}
	return lsp.Null
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
