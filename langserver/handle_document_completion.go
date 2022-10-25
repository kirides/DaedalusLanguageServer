package langserver

import (
	"fmt"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

type textDocumentCompletion struct {
	symbols SymbolProvider
}

func (tdc textDocumentCompletion) fieldsToCompletionItems(fields []symbol.Symbol) []lsp.CompletionItem {
	result := make([]lsp.CompletionItem, 0, len(fields))
	for _, v := range fields {
		ci, err := completionItemFromSymbol(tdc.symbols, v)
		if err != nil {
			continue
		}
		result = append(result, ci)
	}
	return result
}

func (tdc textDocumentCompletion) getTypeFieldsAsCompletionItems(docs SymbolProvider, symbolName string, locals map[string]symbol.Symbol) ([]lsp.CompletionItem, error) {
	symName := strings.ToUpper(symbolName)
	sym, ok := locals[symName]
	if !ok {
		sym, ok = docs.LookupGlobalSymbol(symName, SymbolVariable|SymbolClass|SymbolInstance|SymbolPrototype)
	}

	if !ok {
		return []lsp.CompletionItem{}, nil
	}
	if clsSym, ok := sym.(symbol.Class); ok {
		return tdc.fieldsToCompletionItems(clsSym.Fields), nil
	}
	if protoSym, ok := sym.(symbol.ProtoTypeOrInstance); ok {
		return tdc.getTypeFieldsAsCompletionItems(docs, protoSym.Parent, locals)
	}
	if varSym, ok := sym.(symbol.Variable); ok {
		return tdc.getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	}
	// no way for e.g. C_NPC arrays
	// if varSym, ok := sym.(ArrayVariableSymbol); ok {
	// 	return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	// }
	return []lsp.CompletionItem{}, nil
}

func (h *LspHandler) getTextDocumentCompletion(params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	// TODO: split up further
	tdc := textDocumentCompletion{symbols: h.parsedDocuments}

	result := make([]lsp.CompletionItem, 0, 200)
	parsedDoc, err := h.parsedDocuments.Get(uriToFilename(params.TextDocument.URI))

	partialIdentifier := ""

	if err == nil {
		doc := h.bufferManager.GetBuffer(uriToFilename(params.TextDocument.URI))
		di := symbol.DefinitionIndex{Line: int(params.Position.Line), Column: int(params.Position.Character)}

		scci, found, err := getSignatureCompletions(params, h)
		if err != nil {
			h.logger.Debugf("signature completion error %v: ", err)
		}
		if found {
			return scci, nil
		}

		// dot-completion
		proto, _, err := doc.GetParentSymbolReference(params.Position)
		if err == nil && proto != "" {
			locals := parsedDoc.ScopedVariables(di)
			return tdc.getTypeFieldsAsCompletionItems(h.parsedDocuments, proto, locals)
		}
		partialIdentifier, _ = doc.GetIdentifier(params.Position)

		// locally scoped variables ordered at the top
		localSortIdx := 0
		for _, fn := range parsedDoc.Functions {
			if fn.BodyDefinition.InBBox(di) {
				for _, p := range fn.Parameters {
					if partialIdentifier != "" && !stringContainsAllAnywhere(p.Name(), partialIdentifier) {
						continue
					}
					ci, err := completionItemFromSymbol(h.parsedDocuments, p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (parameter)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				for _, p := range fn.LocalVariables {
					if partialIdentifier != "" && !stringContainsAllAnywhere(p.Name(), partialIdentifier) {
						continue
					}
					ci, err := completionItemFromSymbol(h.parsedDocuments, p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (local)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				break
			}
		}
		for _, fn := range parsedDoc.Instances {
			if fn.BodyDefinition.InBBox(di) {
				return tdc.getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
		for _, fn := range parsedDoc.Prototypes {
			if fn.BodyDefinition.InBBox(di) {
				return tdc.getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
	}

	symbols := make([]symbol.Symbol, 0, 200)
	h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if partialIdentifier != "" && !stringContainsAllAnywhere(s.Name(), partialIdentifier) {
			return nil
		}
		symbols = append(symbols, s)
		return nil
	}, SymbolAll)

	for _, v := range symbols {
		ci, err := completionItemFromSymbol(h.parsedDocuments, v)
		if err != nil {
			continue
		}
		result = append(result, ci)
	}

	return result, nil
}

func (h *LspHandler) handleTextDocumentCompletion(req dls.RpcContext, data lsp.CompletionParams) error {
	items, err := h.getTextDocumentCompletion(&data)
	req.ReplyEither(req.Context(), items, err)
	return nil
}
