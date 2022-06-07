package langserver

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.lsp.dev/jsonrpc2"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// LspHandler ...
type LspHandler struct {
	TextDocumentSync   *textDocumentSync
	bufferManager      *BufferManager
	parsedDocuments    *parseResultsManager
	initialDiagnostics map[string][]lsp.Diagnostic
	handlers           RpcMux
	baseLspHandler
	initialized bool
}

var (
	// ErrWalkAbort should be returned if a walk function should abort early
	ErrWalkAbort = fmt.Errorf("OK")
)

// NewLspHandler ...
func NewLspHandler(conn jsonrpc2.Conn, logger Logger) *LspHandler {
	bufferManager := NewBufferManager()
	parsedDocuments := newParseResultsManager(logger)
	return &LspHandler{
		baseLspHandler: baseLspHandler{
			logger: logger,
			conn:   conn,
		},
		handlers: RpcMux{
			pathToType: map[string]Handler{},
		},
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
	}
}

func getTypeFieldsAsCompletionItems(docs *parseResultsManager, symbolName string, locals map[string]Symbol) ([]lsp.CompletionItem, error) {
	symName := strings.ToUpper(symbolName)
	sym, ok := locals[symName]
	if !ok {
		sym, ok = docs.LookupGlobalSymbol(symName, SymbolVariable|SymbolClass|SymbolInstance|SymbolPrototype)
	}

	if !ok {
		return []lsp.CompletionItem{}, nil
	}
	if clsSym, ok := sym.(ClassSymbol); ok {
		return fieldsToCompletionItems(clsSym.Fields), nil
	}
	if protoSym, ok := sym.(ProtoTypeOrInstanceSymbol); ok {
		return getTypeFieldsAsCompletionItems(docs, protoSym.Parent, locals)
	}
	if varSym, ok := sym.(VariableSymbol); ok {
		return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	}
	return []lsp.CompletionItem{}, nil
}

func (h *LspHandler) getTextDocumentCompletion(params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(params.TextDocument.URI))
	if err == nil {
		doc := h.bufferManager.GetBuffer(h.uriToFilename(params.TextDocument.URI))
		di := DefinitionIndex{Line: int(params.Position.Line), Column: int(params.Position.Character)}

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
					ci, err := completionItemFromSymbol(p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (parameter)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				for _, p := range fn.LocalVariables {
					ci, err := completionItemFromSymbol(p)
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
	}

	h.parsedDocuments.WalkGlobalSymbols(func(s Symbol) error {
		ci, err := completionItemFromSymbol(s)
		if err != nil {
			return nil
		}
		result = append(result, ci)
		return nil
	}, SymbolAll)

	return result, nil
}

func (h *LspHandler) lookUpScopedSymbol(documentURI, identifier string, position lsp.Position) Symbol {
	p, err := h.parsedDocuments.Get(documentURI)
	if err != nil {
		return nil
	}

	di := DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}

	var sym Symbol
	p.WalkScopedVariables(di, func(s Symbol) bool {
		if strings.EqualFold(s.Name(), identifier) {
			sym = s
			return false
		}
		return true
	})
	return sym
}

func (h *LspHandler) lookUpSymbol(documentURI string, position lsp.Position) (Symbol, error) {
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
	doc := h.bufferManager.GetBuffer(h.uriToFilename(params.TextDocument.URI))

	methodCallLine := doc.GetMethodCall(params.Position)
	// The expected method call turned out to be a `func void something( ... )` -> a function definition
	if rxFunctionDef.MatchString(methodCallLine) {
		return lsp.SignatureHelp{}, nil
	}

	methodCallLine = rxStringValues.ReplaceAllLiteralString(methodCallLine, "")
	oldLen := -1
	for len(methodCallLine) != oldLen {
		oldLen = len(methodCallLine)
		methodCallLine = rxFuncCall.ReplaceAllLiteralString(methodCallLine, "")
	}

	// If for some reason the parenthesis of the methodcall went missing
	idxParen := strings.LastIndexByte(methodCallLine, '(')
	if idxParen < 0 {
		return lsp.SignatureHelp{}, fmt.Errorf("the parenthesis of the methodcall went missing")
	}

	word := ""
	for i := idxParen - 1; i > 0; i-- {
		if !isIdentifier(methodCallLine[i]) {
			start := i + 1
			if start+idxParen > len(methodCallLine) {
				return lsp.SignatureHelp{}, fmt.Errorf("idx out of bounds. Bad format :/")
			}
			word = methodCallLine[start : start+idxParen]
		}
	}
	if word == "" {
		word = methodCallLine[:idxParen]
	}
	word = strings.ToUpper(strings.TrimSpace(word))

	funcSymbol, found := h.parsedDocuments.LookupGlobalSymbol(word, SymbolFunction)
	if !found {
		return lsp.SignatureHelp{}, fmt.Errorf("no function symbol found")
	}
	sigCtx := methodCallLine[idxParen+1:]
	fn := funcSymbol.(FunctionSymbol)

	var fnParams []lsp.ParameterInformation
	for _, p := range fn.Parameters {
		doc := findJavadocParam(fn.Documentation(), p.Name())
		var mdDoc interface{}
		if doc != "" {
			mdDoc = &lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: fmt.Sprintf("**%s** - *%s*", p.Name(), doc),
			}
		}
		fnParams = append(fnParams, lsp.ParameterInformation{
			Label:         p.String(),
			Documentation: mdDoc,
		})
	}

	paramIdx := uint32(strings.Count(sigCtx, ","))
	return lsp.SignatureHelp{
		Signatures: []lsp.SignatureInformation{
			{
				Documentation: &lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: simpleJavadocMD(fn),
				},
				Label:      fn.String(),
				Parameters: fnParams,
			},
		},
		ActiveParameter: paramIdx,
		ActiveSignature: 0,
	}, nil
}

func (h *LspHandler) handleGoToDefinition(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.Location, error) {
	symbol, err := h.lookUpSymbol(h.uriToFilename(params.TextDocument.URI), params.Position)
	if err != nil {
		return lsp.Location{}, err
	}

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
		}}, nil
}

func (h *LspHandler) handleTextDocumentCompletion(req RpcContext, data lsp.CompletionParams) error {
	items, err := h.getTextDocumentCompletion(&data)
	replyEither(req.Context(), req, items, err)
	return nil
}

func (h *LspHandler) handleTextDocumentDefinition(req RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.handleGoToDefinition(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), found, nil)
}

func (h *LspHandler) handleTextDocumentHover(req RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.lookUpSymbol(h.uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	h.LogDebug("Found Symbol for Hover: %s\n", found.String())
	return req.Reply(req.Context(), lsp.Hover{
		Range: &lsp.Range{
			Start: data.Position,
			End:   data.Position,
		},
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: strings.TrimSpace(simpleJavadocMD(found) + "\n```daedalus\n" + found.String() + "\n```"),
		},
	}, nil)
}
func (h *LspHandler) handleTextDocumentSignatureHelp(req RpcContext, data lsp.TextDocumentPositionParams) error {
	result, err := h.handleSignatureInfo(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), result, nil)
}

func (h *LspHandler) onInitialized() {
	h.handlers.Register(lsp.MethodTextDocumentCompletion, MakeHandler(h.handleTextDocumentCompletion))
	h.handlers.Register(lsp.MethodTextDocumentDefinition, MakeHandler(h.handleTextDocumentDefinition))
	h.handlers.Register(lsp.MethodTextDocumentHover, MakeHandler(h.handleTextDocumentHover))
	h.handlers.Register(lsp.MethodTextDocumentSignatureHelp, MakeHandler(h.handleTextDocumentSignatureHelp))

	// textDocument/didOpen/didSave/didChange
	h.handlers.Register(lsp.MethodTextDocumentDidOpen, MakeHandler(h.TextDocumentSync.handleTextDocumentDidOpen))
	h.handlers.Register(lsp.MethodTextDocumentDidChange, MakeHandler(h.TextDocumentSync.handleTextDocumentDidChange))
	h.handlers.Register(lsp.MethodTextDocumentDidSave, MakeHandler(h.TextDocumentSync.handleTextDocumentDidSave))
}

// Deliver ...
func (h *LspHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, r jsonrpc2.Request) error {
	h.LogDebug("Requested '%s'\n", r.Method())

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
			},
		}, nil); err != nil {
			return fmt.Errorf("not initialized")
		}
		h.initialized = true
		h.onInitialized()
		return nil
	case lsp.MethodInitialized:
		go func() {
			exe, _ := os.Executable()
			var resultsX []*ParseResult
			if f, err := findFile(filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src")); err == nil {
				resultsX, err = h.parsedDocuments.ParseSource(f)
				if err != nil {
					h.LogError("Error parsing %q: %v", f, err)
					return
				}
			}

			if externalsSrc, err := findFile(filepath.Join("_externals", "externals.src")); err == nil {
				customBuiltins, err := h.parsedDocuments.ParseSource(externalsSrc)
				if err != nil {
					h.LogError("Error parsing %q: %v", externalsSrc, err)
				} else {
					resultsX = append(resultsX, customBuiltins...)
				}
			} else if externalsDaedalus, err := findFile(filepath.Join("_externals", "externals.d")); err == nil {
				parsed, err := h.parsedDocuments.ParseFile(externalsDaedalus)
				if err != nil {
					h.LogError("Error parsing %q: %v", externalsDaedalus, err)
				} else {
					resultsX = append(resultsX, parsed)
				}
			}

			for _, v := range []string{"Gothic.src", "Camera.src", "Menu.src", "Music.src", "ParticleFX.src", "SFX.src", "VisualFX.src"} {
				if full, err := findFile(v); err == nil {
					results, err := h.parsedDocuments.ParseSource(full)
					if err != nil {
						h.LogError("Error parsing %s: %v", full, err)
						return
					}
					resultsX = append(resultsX, results...)
				} else {
					h.LogInfo("Did not parse %q: %v", v, err)
				}
			}

			var diagnostics []lsp.Diagnostic
			tmpDiags := make(map[string][]lsp.Diagnostic)

			for _, p := range resultsX {
				if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
					diagnostics = make([]lsp.Diagnostic, 0, len(p.SyntaxErrors))
					for _, se := range p.SyntaxErrors {
						diagnostics = append(diagnostics, se.Diagnostic())
					}
					tmpDiags[p.Source] = diagnostics
				}
			}
			h.initialDiagnostics = tmpDiags
		}()
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

	if h.initialDiagnostics != nil && len(h.initialDiagnostics) > 0 {
		h.LogInfo("Publishing initial diagnostics (%d).", len(h.initialDiagnostics))
		for k, v := range h.initialDiagnostics {
			h.LogInfo("> %s", k)
			h.conn.Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
				URI:         lsp.DocumentURI(uri.File(k)),
				Diagnostics: v,
			})
		}
		h.initialDiagnostics = map[string][]lsp.Diagnostic{}
	}

	handled, err := h.handlers.Handle(ctx, reply, r)
	if err != nil && handled {
		return err
	}
	return h.baseLspHandler.Handle(ctx, reply, r)
}
