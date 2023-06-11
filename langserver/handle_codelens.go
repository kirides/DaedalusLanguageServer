package langserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/goccy/go-json"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

type codeLensProvider interface {
	Provide(ctx context.Context, source *ParseResult) []lsp.CodeLens
}

type funcCodeLensProvider func(context.Context, *ParseResult) []lsp.CodeLens

func (f funcCodeLensProvider) Provide(ctx context.Context, source *ParseResult) []lsp.CodeLens {
	return f(ctx, source)
}

type codeLensHandler struct {
	lsp            *LspWorkspace
	symbolProvider SymbolProvider
	providers      []codeLensProvider
}

func (p *codeLensHandler) Provide(ctx context.Context, source *ParseResult) []lsp.CodeLens {
	var result []lsp.CodeLens
	for _, v := range p.providers {
		result = append(result, v.Provide(ctx, source)...)
	}
	return result
}

func (p *codeLensHandler) provideImplementations(ctx context.Context, source *ParseResult) []lsp.CodeLens {
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
			Command: &lsp.Command{
				Title:     txt,
				Command:   "editor.action.showReferences",
				Arguments: []any{getSymbolLocation(v).URI, location.Range.Start, locations},
			},
		}

		result = append(result, lens)
	}
	return result
}

type resolveType string

const (
	resolveReferences      resolveType = "references"
	resolveImplementations resolveType = "implementations"
)

type codeLensResolveMetadata struct {
	Command         string
	SourceURI       lsp.DocumentURI
	Type            resolveType
	SourceLocation  lsp.Location
	ReferenceParams lsp.ReferenceParams
}

func (p *codeLensHandler) provideReferences(ctx context.Context, source *ParseResult) []lsp.CodeLens {

	allSymbolsCount := 0
	source.WalkGlobalSymbols(func(s symbol.Symbol) error {
		allSymbolsCount++
		return nil
	}, SymbolAll)

	if allSymbolsCount == 0 {
		return nil
	}

	allSymbols := make([]symbol.Symbol, 0, allSymbolsCount)
	source.WalkGlobalSymbols(func(s symbol.Symbol) error {
		allSymbols = append(allSymbols, s)
		return nil
	}, SymbolAll)

	result := make([]lsp.CodeLens, 0, len(allSymbols))

	for _, v := range allSymbols {
		location := getSymbolLocation(v)

		meta := codeLensResolveMetadata{
			SourceURI:      location.URI,
			SourceLocation: location,
			Type:           resolveReferences,
			Command:        "editor.action.showReferences",
			ReferenceParams: lsp.ReferenceParams{
				TextDocumentPositionParams: lsp.TextDocumentPositionParams{
					TextDocument: lsp.TextDocumentIdentifier{
						URI: lsp.DocumentURI(uri.File(v.Source())),
					},
					Position: lsp.Position{
						Line:      uint32(v.Definition().Start.Line - 1),
						Character: uint32(v.Definition().Start.Column + 1),
					},
				}},
		}
		metaJson, err := json.Marshal(meta)
		if err != nil {
			continue
		}
		lens := lsp.CodeLens{
			Command: nil,
			Range:   location.Range,
			Data:    metaJson,
		}

		result = append(result, lens)
	}

	return result
}

func newCodeLensHandler(h *LspWorkspace, symbolProvider SymbolProvider) *codeLensHandler {
	handler := &codeLensHandler{
		lsp:            h,
		symbolProvider: symbolProvider,
	}
	handler.providers = []codeLensProvider{
		funcCodeLensProvider(handler.provideImplementations),
		funcCodeLensProvider(handler.provideReferences),
	}

	return handler
}

func (h *LspHandler) handleCodeLensResolve(req dls.RpcContext, params lsp.CodeLens) error {
	var meta codeLensResolveMetadata
	if err := json.Unmarshal(params.Data, &meta); err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}
	if meta.Type != resolveReferences {
		return req.Reply(req.Context(), nil, nil)
	}

	ws := h.GetWorkspace(meta.SourceURI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	refsCh := ws.getAllReferences(req.Context(), meta.ReferenceParams)
	var locations []lsp.Location

	for v := range refsCh {
		if locations == nil {
			locations = make([]lsp.Location, 0, 100)
		}
		locations = append(locations, v)
	}
	var txt string
	if len(locations) > 1 {
		txt = fmt.Sprintf("%d references", len(locations))
	} else if len(locations) == 1 {
		txt = "1 references"
	} else {
		txt = "no references"
	}

	params.Command = &lsp.Command{
		Title: txt,
	}
	if len(locations) != 0 {
		params.Command.Command = meta.Command
		params.Command.Arguments = []any{meta.SourceURI, meta.SourceLocation.Range.Start, locations}
	}
	return req.Reply(req.Context(), params, nil)
}

func (h *LspHandler) handleTextDocumentCodeLens(req dls.RpcContext, params lsp.CodeLensParams) error {
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	source := uriToFilename(params.TextDocument.URI)
	if source == "" {

		return req.Reply(req.Context(), nil, nil)
	}

	r, err := ws.parsedDocuments.GetCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}
	locations := newCodeLensHandler(ws, ws.parsedDocuments).Provide(req.Context(), r)
	return req.Reply(req.Context(), locations, nil)
}
