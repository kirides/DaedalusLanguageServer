package langserver

import (
	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "go.lsp.dev/protocol"
)

func getDocumentSymbol(s symbol.Symbol) lsp.DocumentSymbol {
	rn := getSymbolLocation(s).Range
	return lsp.DocumentSymbol{
		Name:           s.Name(),
		Kind:           getSymbolKind(s),
		Range:          rn,
		SelectionRange: rn,
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
