package langserver

import (
	"strconv"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func (h *LspHandler) handleInlayHints(req dls.RpcContext, params lsp.InlayHintParams) error {
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	var hints []lsp.InlayHint

	if !h.config.InlayHints.Constants {
		// the only setting currently
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	source := uriToFilename(params.TextDocument.URI)
	if source == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	buf, err := ws.bufferManager.GetBufferCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	parsed, err := ws.parsedDocuments.ParseSemanticsContentRange(req.Context(), source, string(buf), lsp.Range{})
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

	getSymbolFromToken := func(v token) string {
		name := v.Name()
		idxOpen := strings.Index(name, "[")
		if idxOpen == -1 {
			return name
		}
		actualName := name[:idxOpen]
		return actualName
	}

	onSymbolFound := func(v token, found FoundSymbol) {
		kind := lsp.InlayHintKind(0)
		if found.Location == FoundParameter {
			kind = lsp.Parameter
		}

		value := found.Location.String()
		if cnst, ok := found.Symbol.(symbol.Constant); ok {
			for ok {
				var cnst2 symbol.Symbol
				cnst2, ok = ws.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(cnst.Value), SymbolConstant)
				if ok {
					cnst = cnst2.(symbol.Constant)
				}
			}
			value = ":" + cnst.Value

		} else if cnst, ok := found.Symbol.(symbol.ConstantArray); ok {
			name := v.Name()
			idxOpen := strings.Index(name, "[")
			idxClose := strings.Index(name, "]")
			if idxOpen > 0 && idxClose > idxOpen {
				index := name[idxOpen+1 : idxClose]
				if i, err := strconv.Atoi(index); err == nil && i >= 0 && i < len(cnst.Elements) {
					value = ":" + cnst.Elements[i].GetValue()
				} else {
					// probably a symbol, so lets find it!
					resolvedIndex := index
					ok := true
					for ok {
						var cnst2 symbol.Symbol
						cnst2, ok = ws.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(resolvedIndex), SymbolConstant)
						if ok {
							resolvedIndex = cnst2.(symbol.Constant).Value
						}
					}
					if i, err := strconv.Atoi(resolvedIndex); err == nil && i >= 0 && i < len(cnst.Elements) {
						value = ":" + cnst.Elements[i].GetValue()
					}
				}
			}
			if value == "" {
				value = ":" + cnst.Value // atleast show preview for constant
			}
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
			found, ok := parsed.FindScopedVariableDeclaration(v.Definition().Start, getSymbolFromToken(v))
			if !ok {
				var symb symbol.Symbol
				symb, ok = ws.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(getSymbolFromToken(v)), SymbolAll)
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
