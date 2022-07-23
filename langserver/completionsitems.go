package langserver

import (
	"fmt"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"github.com/kirides/DaedalusLanguageServer/javadoc"

	lsp "go.lsp.dev/protocol"
)

func completionItemFromSymbol(symbols SymbolProvider, s symbol.Symbol) (lsp.CompletionItem, error) {
	kind, err := completionItemKindForSymbol(s)
	if err != nil {
		return lsp.CompletionItem{}, err
	}
	return lsp.CompletionItem{
		Kind:   kind,
		Label:  s.Name(),
		Detail: SymbolToReadableCode(symbols, s),
		Documentation: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: javadoc.MarkdownSimple(s),
		},
	}, nil
}

func completionItemKindForSymbol(s symbol.Symbol) (lsp.CompletionItemKind, error) {
	switch s.(type) {
	case symbol.ArrayVariable, symbol.Variable:
		return lsp.CompletionItemKindVariable, nil
	case symbol.Constant, symbol.ConstantArray:
		return lsp.CompletionItemKindConstant, nil
	case symbol.Function:
		return lsp.CompletionItemKindFunction, nil
	case symbol.Class:
		return lsp.CompletionItemKindClass, nil
	case symbol.ProtoTypeOrInstance:
		return lsp.CompletionItemKindClass, nil
	}
	return lsp.CompletionItemKind(-1), fmt.Errorf("symbol not found")
}

func fieldsToCompletionItems(docs *parseResultsManager, fields []symbol.Symbol) []lsp.CompletionItem {
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
