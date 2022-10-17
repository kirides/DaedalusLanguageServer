package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "go.lsp.dev/protocol"
)

func (h *LspHandler) handleTextDocumentImplementation(req dls.RpcContext, params lsp.ImplementationParams) error {
	sym, err := h.lookUpSymbol(uriToFilename(params.TextDocument.URI), params.Position)
	if err != nil {
		req.Reply(req.Context(), nil, nil)
		h.LogDebug("Failed to lookup symbol. %v", err)
		return nil
	}

	if _, ok := sym.(symbol.Class); !ok {
		protoOrSymbol, ok := sym.(symbol.ProtoTypeOrInstance)
		if !ok || protoOrSymbol.IsInstance {
			req.Reply(req.Context(), nil, nil)
			return nil
		}
	}

	result := []lsp.Location{}
	h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		switch v := s.(type) {
		case symbol.ProtoTypeOrInstance:
			if strings.EqualFold(sym.Name(), v.Parent) {
				result = append(result, getSymbolLocation(v))
			}
		}
		return nil
	}, SymbolInstance|SymbolPrototype)

	req.Reply(req.Context(), result, nil)
	return nil
}
