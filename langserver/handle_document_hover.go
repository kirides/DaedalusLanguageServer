package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"github.com/kirides/DaedalusLanguageServer/javadoc"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func (h *LspHandler) handleTextDocumentHover(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	ws := h.GetWorkspace(data.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	found, err := ws.lookUpSymbol(uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	h.LogDebug("Found Symbol for Hover: %s", found.Symbol.String())

	codeSnippet := SymbolToReadableCode(ws.parsedDocuments, found.Symbol)
	codeSnippetFinal := codeSnippet

	if found.Location == FoundField {
		cnst, ok := found.Symbol.(symbol.Constant)
		ogVal := cnst.Value
		for ok {
			var cnst2 symbol.Symbol
			cnst2, ok = ws.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(cnst.Value), SymbolConstant)
			if ok {
				cnst = cnst2.(symbol.Constant)
			}
		}
		if ogVal != cnst.Value {
			codeSnippetFinal = codeSnippet + " // " + cnst.Value
		}
	}

	return req.Reply(req.Context(), lsp.Hover{
		Range: lsp.Range{
			Start: data.Position,
			End:   data.Position,
		},
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: strings.TrimSpace(javadoc.MarkdownSimple(found.Symbol) + "\n\n```daedalus\n" + codeSnippetFinal + "\n```"),
		},
	}, nil)
}
