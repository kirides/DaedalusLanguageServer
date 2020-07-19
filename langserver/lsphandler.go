package langserver

import (
	"context"
	"encoding/json"
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
	baseLspHandler

	initialized   bool
	bufferManager *BufferManager

	parsedDocuments         *parseResultsManager
	TextDocumentSyncHandler jsonrpc2.Handler
}

var (
	// ErrWalkAbort should be returned if a walk function should abort early
	ErrWalkAbort = fmt.Errorf("OK")
)

// NewLspHandler ...
func NewLspHandler() *LspHandler {
	bufferManager := NewBufferManager()
	parsedDocuments := newParseResultsManager()
	return &LspHandler{
		initialized:     false,
		bufferManager:   bufferManager,
		parsedDocuments: parsedDocuments,
		TextDocumentSyncHandler: &textDocumentSyncHandler{
			bufferManager:   bufferManager,
			parsedDocuments: parsedDocuments,
		},
	}
}

func (h *LspHandler) handleTextDocumentCompletion(ctx context.Context, params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)
	h.parsedDocuments.WalkGlobalSymbols(func(s Symbol) error {
		var kind lsp.CompletionItemKind
		switch s.(type) {
		case VariableSymbol:
			kind = lsp.VariableCompletion
		case ConstantSymbol:
			kind = lsp.ConstantCompletion
		case FunctionSymbol:
			kind = lsp.FunctionCompletion
		case ClassSymbol:
			kind = lsp.ClassCompletion
		case ProtoTypeOrInstanceSymbol:
			kind = lsp.ClassCompletion
		default:
			return nil
		}
		result = append(result, lsp.CompletionItem{
			Kind:   kind,
			Label:  s.Name(),
			Detail: s.String(),
			Documentation: lsp.MarkupContent{
				Kind:  lsp.PlainText,
				Value: s.Documentation(),
			},
		})
		return nil
	})

	return result, nil
}

func (h *LspHandler) lookUpSymbol(documentURI string, position lsp.Position) (Symbol, error) {
	doc := h.bufferManager.GetBuffer(documentURI)
	identifier := doc.GetWordRangeAtPosition(position)

	p, err := h.parsedDocuments.Get(documentURI)
	if err == nil {
		di := DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}
		for _, f := range p.Functions {
			if f.BodyDefinition.InBBox(di) {
				for _, param := range f.Parameters {
					if strings.EqualFold(param.Name(), identifier) {
						return param, nil
					}
				}
				for _, local := range f.LocalVariables {
					if strings.EqualFold(local.Name(), identifier) {
						return local, nil
					}
				}
			}
		}
	}

	var symbol Symbol
	h.parsedDocuments.WalkGlobalSymbols(func(s Symbol) error {
		if strings.EqualFold(s.Name(), identifier) {
			symbol = s
			return fmt.Errorf("OK")
		}
		return nil
	})

	if symbol == nil {
		return nil, fmt.Errorf("Identifier %q not found", identifier)
	}

	return symbol, nil
}

func (h *LspHandler) handleSignatureInfo(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.SignatureHelp, error) {
	doc := h.bufferManager.GetBuffer(params.TextDocument.URI.Filename())
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
				return lsp.SignatureHelp{}, fmt.Errorf("Idx out of bounds. Bad format :/")
			}
			word = methodCallLine[start : start+idxParen]
		}
	}
	if word == "" {
		word = methodCallLine[:idxParen]
	}
	word = strings.TrimSpace(word)

	var funcSymbol Symbol

	h.parsedDocuments.WalkGlobalSymbols(func(s Symbol) error {
		if _, ok := s.(FunctionSymbol); ok {
			if strings.EqualFold(s.Name(), word) {
				funcSymbol = s
				return ErrWalkAbort
			}
		}
		return nil
	})
	if funcSymbol == nil {
		return lsp.SignatureHelp{}, fmt.Errorf("no functino symbol found")
	}
	sigCtx := methodCallLine[idxParen+1:]
	fn := funcSymbol.(FunctionSymbol)

	var fnParams []lsp.ParameterInformation
	for _, p := range fn.Parameters {
		fnParams = append(fnParams, lsp.ParameterInformation{
			Label: p.String(),
		})
	}

	return lsp.SignatureHelp{
		Signatures: []lsp.SignatureInformation{
			{
				Documentation: &lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: fn.Documentation(),
				},
				Label:      fn.String(),
				Parameters: fnParams,
			},
		},
		ActiveParameter: float64(strings.Count(sigCtx, ",")),
		ActiveSignature: 0,
	}, nil
}

func (h *LspHandler) handleGoToDefinition(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.Location, error) {
	symbol, err := h.lookUpSymbol(params.TextDocument.URI.Filename(), params.Position)
	if err != nil {
		return lsp.Location{}, err
	}

	return lsp.Location{
		URI: uri.File(symbol.Source()),
		Range: lsp.Range{
			Start: lsp.Position{
				Character: float64(symbol.Definition().Start.Column),
				Line:      float64(symbol.Definition().Start.Line - 1),
			},
			End: lsp.Position{
				Character: float64(symbol.Definition().Start.Column + len(symbol.Name())),
				Line:      float64(symbol.Definition().Start.Line - 1),
			},
		}}, nil
}

// Deliver ...
func (h *LspHandler) Deliver(ctx context.Context, r *jsonrpc2.Request, delivered bool) bool {
	h.Log("Requested '%s'\n", r.Method)
	if delivered {
		return false
	}

	if r.Params != nil {
		var paramsMap map[string]interface{}
		json.Unmarshal(*r.Params, &paramsMap)
		// fmt.Fprintf(os.Stderr, "Params: %+v\n", paramsMap)
	}
	switch r.Method {
	case lsp.MethodInitialize:
		if err := r.Reply(ctx, lsp.InitializeResult{
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
					Change:    float64(lsp.Full),
					OpenClose: true,
					Save: &lsp.SaveOptions{
						IncludeText: true,
					},
				},
			},
		}, nil); err != nil {
			return false
		}
		h.initialized = true
		return true
	case lsp.MethodInitialized:
		exe, _ := os.Executable()
		resultsX, err := h.parsedDocuments.ParseSource(filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src"))
		if err != nil {
			h.Log("Error parsing %q: %v", filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src"), err)
			return true
		}
		results, err := h.parsedDocuments.ParseSource("Gothic.src")
		if err != nil {
			h.Log("Error parsing Gothic.src: %v", err)
			return true
		}
		results = append(resultsX, results...)

		diagnostics := make([]lsp.Diagnostic, 0, 5)
		for _, p := range results {
			diagnostics = diagnostics[:0]
			if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
				for _, se := range p.SyntaxErrors {
					diagnostics = append(diagnostics, se.Diagnostic())
				}
			}
			r.Conn().Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
				URI:         lsp.DocumentURI(uri.File(p.Source)),
				Diagnostics: diagnostics,
			})
		}
		return true
	}

	// DEFAULT / OTHERWISE

	if !h.initialized {
		if !r.IsNotify() {
			r.Reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.ServerNotInitialized, "Not initialized yet"))
		}
		return false
	}

	// Recover if something bad happens in the handlers...

	switch r.Method {
	case lsp.MethodTextDocumentCompletion:
		var params lsp.CompletionParams
		json.Unmarshal(*r.Params, &params)
		items, err := h.handleTextDocumentCompletion(ctx, &params)
		h.replyEither(ctx, r, items, err)

		return true
	case lsp.MethodTextDocumentDefinition:
		var params lsp.TextDocumentPositionParams
		json.Unmarshal(*r.Params, &params)
		found, err := h.handleGoToDefinition(ctx, &params)
		if err != nil {
			h.replyEither(ctx, r, nil, nil)
		} else {
			h.replyEither(ctx, r, found, nil)
		}
	case lsp.MethodTextDocumentHover:
		var params lsp.TextDocumentPositionParams
		json.Unmarshal(*r.Params, &params)
		found, err := h.lookUpSymbol(params.TextDocument.URI.Filename(), params.Position)
		if err != nil {
			h.replyEither(ctx, r, nil, nil)
		} else {
			h.Log("Found Symbol for Hover: %s\n", found.String())
			h.replyEither(ctx, r, lsp.Hover{
				Range: lsp.Range{
					Start: params.Position,
					End:   params.Position,
				},
				Contents: lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(found.Documentation(), "\r", ""), "\n", "  \n") + "\n```daedalus\n" + found.String() + "\n```"),
				},
			}, nil)
		}
	case lsp.MethodTextDocumentSignatureHelp:
		var params lsp.TextDocumentPositionParams
		json.Unmarshal(*r.Params, &params)
		result, err := h.handleSignatureInfo(ctx, &params)
		if err == nil {
			r.Reply(ctx, result, nil)
		} else {
			r.Reply(ctx, nil, nil)
		}
	default:
		return h.baseLspHandler.Deliver(ctx, r, delivered)
	}
	return true
}
