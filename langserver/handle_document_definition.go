package langserver

import (
	"fmt"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

func getSymbolLocation(symbol symbol.Symbol) lsp.Location {
	return lsp.Location{
		URI: lsp.DocumentURI(uri.File(symbol.Source())),
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

func (h *LspWorkspace) getDotCompletedSymbol(params *lsp.TextDocumentPositionParams) (symbol.Symbol, error) {
	parsedDoc, err := h.parsedDocuments.Get(uriToFilename(params.TextDocument.URI))

	if err != nil {
		return nil, fmt.Errorf("not found")

	}
	doc := h.bufferManager.GetBuffer(uriToFilename(params.TextDocument.URI))
	di := symbol.DefinitionIndex{Line: int(params.Position.Line), Column: int(params.Position.Character)}

	// dot-completion
	proto, _, err := doc.GetParentSymbolReference(params.Position)
	if err != nil || proto == "" {
		identifier, err := doc.GetIdentifier(params.Position)
		if err != nil {
			return nil, fmt.Errorf("could not find identifier below cursor")
		}

		sym, ok := parsedDoc.LookupScopedVariable(di, identifier)
		if !ok {
			sym, ok = h.parsedDocuments.LookupGlobalSymbol(identifier, SymbolAll)
		}
		if !ok {
			return nil, fmt.Errorf("could not find local or global variable")
		}

		return sym, nil
	}

	fieldName, _ := doc.GetIdentifier(params.Position)

	parentIdentifier := strings.ToUpper(strings.TrimSpace(proto))
	sym, ok := parsedDoc.LookupScopedVariable(di, parentIdentifier)
	if !ok {
		sym, ok = h.parsedDocuments.LookupGlobalSymbol(parentIdentifier, SymbolAll)
	}
	if !ok {
		return nil, fmt.Errorf("could not find parent for member access (PARENT.Member)")
	}

	switch s := sym.(type) {
	case symbol.Variable:
		sym, ok = h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(s.GetType()), SymbolClass|SymbolInstance|SymbolPrototype)
		for ok {
			switch s := sym.(type) {
			case symbol.ProtoTypeOrInstance:
				for _, f := range s.Fields {
					if strings.EqualFold(f.NameValue, fieldName) {
						return f, nil
					}
				}
				sym, ok = h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(s.Parent), SymbolClass|SymbolInstance|SymbolPrototype)
			case symbol.Class:
				for _, f := range s.Fields {
					if strings.EqualFold(f.Name(), fieldName) {
						return f, nil
					}
				}
				// still not found the actual definition, just return the class
				return s, nil
			}
		}
	}

	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return sym, nil
}

func (h *LspHandler) handleTextDocumentDefinition(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	ws := h.GetWorkspace(data.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	symbol, err := ws.lookUpSymbol(uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		sym, err := ws.getDotCompletedSymbol(&data)
		if err != nil {
			return req.Reply(req.Context(), nil, nil)
		}
		symbol = FoundSymbol{Symbol: sym, Location: FoundGlobal}
	}

	return req.Reply(req.Context(), getSymbolLocation(symbol.Symbol), nil)
}
