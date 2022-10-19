package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func (h *LspHandler) handleInlayHints(req dls.RpcContext, params lsp.InlayHintParams) error {
	var hints []lsp.InlayHint

	source := uriToFilename(params.TextDocument.URI)
	if source == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	buf, err := h.bufferManager.GetBufferCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	parsed, err := h.parsedDocuments.ParseSemanticsContentRange(req.Context(), source, string(buf), lsp.Range{})
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	bbox := symbol.Definition{
		Start: symbol.DefinitionIndex{
			Line:   int(params.Range.Start.Line) + 1,
			Column: int(params.Range.Start.Character),
		},
		End: symbol.DefinitionIndex{
			Line:   int(params.Range.End.Line) + 1,
			Column: int(params.Range.End.Character),
		},
	}

	onSymbolFound := func(v token, found FoundSymbol) {
		kind := lsp.InlayHintKind(0)
		if found.Location == FoundParameter {
			kind = lsp.Parameter
		}

		value := found.Location.String()
		if cnst, ok := found.Symbol.(symbol.Constant); ok {
			value = ":" + cnst.Value
		} else {
			// only show value from constants for now
			_ = value
			return
		}
		hints = append(hints, lsp.InlayHint{
			Position: &lsp.Position{
				Line:      uint32(v.Definition().End.Line - 1),
				Character: uint32(v.Definition().End.Column),
			},
			PaddingLeft:  true,
			PaddingRight: false,
			Kind:         kind,
			Label: []lsp.InlayHintLabelPart{
				{Value: value},
			},
		})
	}

	for _, v := range parsed.GlobalIdentifiers {
		if bbox.InBBox(v.definition.Start) || bbox.InBBox(v.definition.End) {
			found, ok := parsed.FindScopedVariableDeclaration(v.Definition().Start, v.Name())
			if !ok {
				var symb symbol.Symbol
				symb, ok = h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(v.Name()), SymbolAll)
				found = FoundSymbol{symb, FoundGlobal}
			}
			if !ok {
				continue
			}
			onSymbolFound(v, found)
		}
	}

	req.Reply(req.Context(), hints, nil)
	return nil
}
