package langserver

import (
	"fmt"
	"sync"

	lsp "go.lsp.dev/protocol"
)

var (
	// ErrDocumentNotFound ...
	ErrDocumentNotFound = fmt.Errorf("document not found")
	// ErrSymbolNotFound ...
	ErrSymbolNotFound = fmt.Errorf("symbol not found")
)

// DocumentsManager ...
type DocumentsManager struct {
	parsedResults sync.Map
}

// GetParsed ...
func (d *DocumentsManager) GetParsed(documentURI string) (*ParseResult, error) {
	if v, ok := d.parsedResults.Load(documentURI); ok {
		return v.(*ParseResult), nil
	}
	return nil, ErrDocumentNotFound
}

// UpdateParseResult ...
func (d *DocumentsManager) UpdateParseResult(documentURI string, result *ParseResult) {
	d.parsedResults.Store(documentURI, result)
}

// GetCompletionSymbols ...
func (d *DocumentsManager) GetCompletionSymbols(documentURI string, position *lsp.Position) []lsp.CompletionItem {
	return []lsp.CompletionItem{}
}

// LookupSymbol ...
func (d *DocumentsManager) LookupSymbol(identifier string) (Symbol, error) {
	return nil, ErrSymbolNotFound
}
