package langserver

import (
	"fmt"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

type codeLensProvider interface {
	Provide(source *ParseResult) []lsp.CodeLens
}

type funcCodeLensProvider func(*ParseResult) []lsp.CodeLens

func (f funcCodeLensProvider) Provide(source *ParseResult) []lsp.CodeLens {
	return f(source)
}

type codeLensHandler struct {
	symbolProvider SymbolProvider
	providers      []codeLensProvider
}

func (p *codeLensHandler) Provide(source *ParseResult) []lsp.CodeLens {
	var result []lsp.CodeLens
	for _, v := range p.providers {
		result = append(result, v.Provide(source)...)
	}
	return result
}

func (p *codeLensHandler) provideImplementations(source *ParseResult) []lsp.CodeLens {
	countClassesAndPrototypes := 0

	source.WalkGlobalSymbols(func(s symbol.Symbol) error {
		countClassesAndPrototypes++
		return nil
	}, SymbolClass|SymbolPrototype)

	if countClassesAndPrototypes == 0 {
		return nil
	}

	// Collect all classes and prototypes from the current document
	currentClassesAndPrototypes := make(map[string]symbol.Symbol, countClassesAndPrototypes)
	source.WalkGlobalSymbols(func(s symbol.Symbol) error {
		currentClassesAndPrototypes[strings.ToUpper(s.Name())] = s
		return nil
	}, SymbolClass|SymbolPrototype)

	// grab all instances and prototypes
	allInstancesAndPrototypesByParentCount := 0
	p.symbolProvider.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if s.Source() == "" {
			return nil
		}

		allInstancesAndPrototypesByParentCount++
		return nil
	}, SymbolInstance|SymbolPrototype)

	allInstancesAndPrototypesByParent := make(map[string][]lsp.Location, allInstancesAndPrototypesByParentCount)
	p.symbolProvider.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if s.Source() == "" {
			return nil
		}

		key := strings.ToUpper(s.(symbol.ProtoTypeOrInstance).Parent)
		if _, ok := currentClassesAndPrototypes[key]; !ok {
			return nil
		}
		items, ok := allInstancesAndPrototypesByParent[key]
		if !ok {
			items = make([]lsp.Location, 0, 10)
		}
		allInstancesAndPrototypesByParent[key] = append(items, getSymbolLocation(s))
		return nil
	}, SymbolInstance|SymbolPrototype)

	result := make([]lsp.CodeLens, 0, len(currentClassesAndPrototypes))
	for k, v := range currentClassesAndPrototypes {
		locations, ok := allInstancesAndPrototypesByParent[k]
		if !ok {
			continue
		}

		var txt string
		if len(locations) > 1 {
			txt = fmt.Sprintf("%d implementations", len(locations))
		} else {
			txt = "1 implementation"
		}
		location := getSymbolLocation(v)
		lens := lsp.CodeLens{
			Range: location.Range,
			Command: lsp.Command{
				Title:     txt,
				Command:   "editor.action.showReferences",
				Arguments: []any{getSymbolLocation(v).URI, location.Range.Start, locations},
			},
		}

		result = append(result, lens)
	}
	return result
}

func newCodeLensHandler(symbolProvider SymbolProvider) *codeLensHandler {
	h := &codeLensHandler{
		symbolProvider: symbolProvider,
	}
	h.providers = []codeLensProvider{
		funcCodeLensProvider(h.provideImplementations),
	}

	return h
}

func (h *LspHandler) handleTextDocumentCodeLens(req dls.RpcContext, params lsp.CodeLensParams) error {
	source := uriToFilename(params.TextDocument.URI)
	if source == "" {

		return req.Reply(req.Context(), nil, nil)
	}

	r, err := h.parsedDocuments.GetCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}
	result := newCodeLensHandler(h.parsedDocuments).Provide(r)
	return req.Reply(req.Context(), result, nil)
}
