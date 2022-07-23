package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/javadoc"
	lsp "go.lsp.dev/protocol"
)

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
			Value: strings.TrimSpace(javadoc.MarkdownSimple(found) + "\n\n```daedalus\n" + SymbolToReadableCode(h.parsedDocuments, found) + "\n```"),
		},
	}, nil)
}
