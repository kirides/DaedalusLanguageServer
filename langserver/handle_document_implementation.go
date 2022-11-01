package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func (h *LspHandler) handleTextDocumentImplementation(req dls.RpcContext, params lsp.ImplementationParams) error {
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	found, err := ws.lookUpSymbol(uriToFilename(params.TextDocument.URI), params.Position)
	if err != nil {
		req.Reply(req.Context(), nil, nil)
		h.LogDebug("Failed to lookup symbol. %v", err)
		return nil
	}

	if _, ok := found.Symbol.(symbol.Class); !ok {
		protoOrSymbol, ok := found.Symbol.(symbol.ProtoTypeOrInstance)
		if !ok || protoOrSymbol.IsInstance {
			req.Reply(req.Context(), nil, nil)
			return nil
		}
	}

	result := []lsp.Location{}
	ws.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		switch v := s.(type) {
		case symbol.ProtoTypeOrInstance:
			if strings.EqualFold(found.Symbol.Name(), v.Parent) {
				result = append(result, getSymbolLocation(v))
			}
		}
		return nil
	}, SymbolInstance|SymbolPrototype)

	req.Reply(req.Context(), result, nil)
	return nil
}
