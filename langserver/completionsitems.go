package langserver

import (
	"fmt"

	lsp "go.lsp.dev/protocol"
)

func completionItemFromSymbol(docs *parseResultsManager, s Symbol) (lsp.CompletionItem, error) {
	kind, err := completionItemKindForSymbol(s)
	if err != nil {
		return lsp.CompletionItem{}, err
	}
	return lsp.CompletionItem{
		Kind:   kind,
		Label:  s.Name(),
		Detail: docs.getSymbolCode(s),
		Documentation: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: simpleJavadocMD(s),
		},
	}, nil
}

func completionItemKindForSymbol(s Symbol) (lsp.CompletionItemKind, error) {
	switch s.(type) {
	case ArrayVariableSymbol, VariableSymbol:
		return lsp.CompletionItemKindVariable, nil
	case ConstantSymbol, ConstantArraySymbol:
		return lsp.CompletionItemKindConstant, nil
	case FunctionSymbol:
		return lsp.CompletionItemKindFunction, nil
	case ClassSymbol:
		return lsp.CompletionItemKindClass, nil
	case ProtoTypeOrInstanceSymbol:
		return lsp.CompletionItemKindClass, nil
	}
	return lsp.CompletionItemKind(-1), fmt.Errorf("Symbol not found")
}

func fieldsToCompletionItems(docs *parseResultsManager, fields []Symbol) []lsp.CompletionItem {
	result := make([]lsp.CompletionItem, 0, len(fields))
	for _, v := range fields {
		ci, err := completionItemFromSymbol(docs, v)
		if err != nil {
			continue
		}
		result = append(result, ci)
	}
	return result
}
