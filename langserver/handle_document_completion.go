package langserver

import (
	"fmt"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "go.lsp.dev/protocol"
)

func getTypeFieldsAsCompletionItems(docs *parseResultsManager, symbolName string, locals map[string]symbol.Symbol) ([]lsp.CompletionItem, error) {
	symName := strings.ToUpper(symbolName)
	sym, ok := locals[symName]
	if !ok {
		sym, ok = docs.LookupGlobalSymbol(symName, SymbolVariable|SymbolClass|SymbolInstance|SymbolPrototype)
	}

	if !ok {
		return []lsp.CompletionItem{}, nil
	}
	if clsSym, ok := sym.(symbol.Class); ok {
		return fieldsToCompletionItems(docs, clsSym.Fields), nil
	}
	if protoSym, ok := sym.(symbol.ProtoTypeOrInstance); ok {
		return getTypeFieldsAsCompletionItems(docs, protoSym.Parent, locals)
	}
	if varSym, ok := sym.(symbol.Variable); ok {
		return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	}
	// no way for e.g. C_NPC arrays
	// if varSym, ok := sym.(ArrayVariableSymbol); ok {
	// 	return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	// }
	return []lsp.CompletionItem{}, nil
}

func (h *LspHandler) getTextDocumentCompletion(params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(params.TextDocument.URI))
	if err == nil {
		doc := h.bufferManager.GetBuffer(h.uriToFilename(params.TextDocument.URI))
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
			return getTypeFieldsAsCompletionItems(h.parsedDocuments, proto, locals)
		}

		// locally scoped variables ordered at the top
		localSortIdx := 0
		for _, fn := range parsedDoc.Functions {
			if fn.BodyDefinition.InBBox(di) {
				for _, p := range fn.Parameters {
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
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
		for _, fn := range parsedDoc.Prototypes {
			if fn.BodyDefinition.InBBox(di) {
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]symbol.Symbol{})
			}
		}
	}

	h.parsedDocuments.WalkGlobalSymbols(func(s symbol.Symbol) error {
		ci, err := completionItemFromSymbol(h.parsedDocuments, s)
		if err != nil {
			return nil
		}
		result = append(result, ci)
		return nil
	}, SymbolAll)

	return result, nil
}

func (h *LspHandler) handleTextDocumentCompletion(req dls.RpcContext, data lsp.CompletionParams) error {
	items, err := h.getTextDocumentCompletion(&data)
	req.ReplyEither(req.Context(), items, err)
	return nil
}
