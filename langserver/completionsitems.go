package langserver

import (
	"fmt"

	lsp "go.lsp.dev/protocol"
)

func completionItemFromSymbol(s Symbol) (lsp.CompletionItem, error) {
	kind, err := completionItemKindForSymbol(s)
	if err != nil {
		return lsp.CompletionItem{}, err
	}
	return lsp.CompletionItem{
		Kind:   kind,
		Label:  s.Name(),
		Detail: s.String(),
		Documentation: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: simpleJavadocMD(s),
		},
	}, nil
}

func completionItemKindForSymbol(s Symbol) (lsp.CompletionItemKind, error) {
	switch s.(type) {
	case VariableSymbol:
		return lsp.CompletionItemKindVariable, nil
	case ArrayVariableSymbol:
		return lsp.CompletionItemKindVariable, nil
	case ConstantSymbol:
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

func fieldsToCompletionItems(fields []Symbol) []lsp.CompletionItem {
	result := make([]lsp.CompletionItem, 0, len(fields))
	for _, v := range fields {
		ci, err := completionItemFromSymbol(v)
		if err != nil {
			continue
		}
		result = append(result, ci)
	}
	return result
}
