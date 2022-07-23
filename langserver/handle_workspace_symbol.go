package langserver

import (
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "go.lsp.dev/protocol"
)

func getSymbolInformation(s symbol.Symbol) lsp.SymbolInformation {
	return lsp.SymbolInformation{
		Name:     s.Name(),
		Kind:     getSymbolKind(s),
		Location: getSymbolLocation(s),
	}
}

func collectWorkspaceSymbols(result []lsp.SymbolInformation, s symbol.Symbol) []lsp.SymbolInformation {
	mainSymb := getSymbolInformation(s)
	result = append(result, mainSymb)

	if cls, ok := s.(symbol.Class); ok {
		for _, v := range cls.Fields {
			si := getSymbolInformation(v)
			si.ContainerName = s.Name()
			result = append(result, si)
		}
	} else if cls, ok := s.(symbol.ProtoTypeOrInstance); ok {
		for _, v := range cls.Fields {
			si := getSymbolInformation(v)
			si.ContainerName = s.Name()
			result = append(result, si)
		}
	}
	return result
}

func (h *LspHandler) handleWorkspaceSymbol(req dls.RpcContext, params lsp.WorkspaceSymbolParams) error {
	numSymbols := h.parsedDocuments.CountSymbols()
	result := make([]lsp.SymbolInformation, 0, numSymbols)
	buffer := make([]lsp.SymbolInformation, 0, 50)

	qlower := strings.ToLower(params.Query)

	err := h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if req.Context().Err() != nil {
			h.logger.Debugf("request cancelled", "method", req.Request().Method())
			return req.Context().Err()
		}
		if qlower == "" {
			result = collectWorkspaceSymbols(result, s)
			return nil
		}

		// pre filtering
		buffer = buffer[:0]
		buffer = collectWorkspaceSymbols(buffer, s)
		for _, v := range buffer {
			if stringContainsAllAnywhere(v.Name, params.Query) {
				result = append(result, v)
			}
		}
		return nil
	}, SymbolAll)

	if err != nil {
		return nil
	}

	req.Reply(req.Context(), result, nil)
	return nil
}
