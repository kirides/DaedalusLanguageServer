package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func (h *LspHandler) handleTextDocumentTypeDefinition(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	ws := h.GetWorkspace(data.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	ssymbol, err := ws.lookUpSymbol(uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		sym, err := ws.getDotCompletedSymbol(&data)
		if err != nil {
			return req.Reply(req.Context(), nil, nil)
		}
		ssymbol = FoundSymbol{Symbol: sym, Location: FoundGlobal}
	}
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}

	switch s := ssymbol.Symbol.(type) {
	case symbol.ProtoTypeOrInstance, symbol.Class, symbol.Function:
		return req.Reply(req.Context(), getSymbolLocation(s), nil)
	}

	if getTyper, ok := ssymbol.Symbol.(interface{ GetType() string }); ok {
		typ := getTyper.GetType()
		sym, ok := ws.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(typ), SymbolClass|SymbolInstance|SymbolPrototype)
		if ok {
			return req.Reply(req.Context(), getSymbolLocation(sym), nil)
		}
	}

	return req.Reply(req.Context(), nil, nil)
}
