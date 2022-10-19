package langserver

import (
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

func (h *LspHandler) handleTextDocumentDefinition(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	symbol, err := h.lookUpSymbol(uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		req.Reply(req.Context(), nil, nil)
		return err
	}

	return req.Reply(req.Context(), getSymbolLocation(symbol.Symbol), nil)
}
