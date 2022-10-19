package langserver

import (
	"fmt"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"github.com/kirides/DaedalusLanguageServer/javadoc"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
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
		return lsp.VariableCompletion, nil
	case symbol.Constant, symbol.ConstantArray:
		return lsp.ConstantCompletion, nil
	case symbol.Function:
		return lsp.FunctionCompletion, nil
	case symbol.Class:
		return lsp.ClassCompletion, nil
	case symbol.ProtoTypeOrInstance:
		return lsp.ClassCompletion, nil
	}
	return lsp.CompletionItemKind(-1), fmt.Errorf("symbol not found")
}
