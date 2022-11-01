package langserver

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"unicode"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

type LspWorkspace struct {
	logger          dls.Logger
	bufferManager   *BufferManager
	parsedDocuments *parseResultsManager
	config          LspConfig
	onceParseAll    sync.Once

	parsedKnownSrcFiles  concurrentSet[string]
	workspaceInitialized bool

	workspaceCtx       context.Context
	cancelWorkspaceCtx context.CancelFunc
}

func (h *LspWorkspace) lookUpScope(documentURI string, position lsp.Position) (symbol.Symbol, bool) {
	p, err := h.parsedDocuments.Get(documentURI)
	if err != nil {
		return nil, false
	}

	di := symbol.DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}

	return p.FindContainingScope(di)
}

func (h *LspWorkspace) getEffectiveValue(sym symbol.ProtoTypeOrInstance, field string) (symbol.Constant, bool) {
	var own symbol.Constant
	found := false
	for _, v := range sym.Fields {
		if strings.EqualFold(v.Name(), field) {
			own = v
			found = true
			break
		}
	}

	parent, ok := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(sym.Parent), SymbolPrototype|SymbolClass)
	if ok {
		switch p := parent.(type) {
		case symbol.Class:
			val := ""
			for _, v := range p.Fields {
				if strings.EqualFold(v.Name(), field) {
					getType, ok := v.(interface{ GetType() string })
					if ok {
						typ := getType.GetType()
						switch {
						case strings.EqualFold(typ, "int"):
							val = "0"
							typ = "int"
						case strings.EqualFold(typ, "float"):
							val = "0.0"
							typ = "float"
						case strings.EqualFold(typ, "string"):
							val = ""
							typ = "string"
						}
						if found {
							// we have the value, just use it for finding the type
							return symbol.NewConstant(v.Name(), typ, own.Source(), own.Documentation(), own.Definition(), own.Value), true
						}
						return symbol.NewConstant(v.Name(), typ, v.Source(), v.Documentation(), v.Definition(), val), true
					}
					break
				}
			}
		case symbol.ProtoTypeOrInstance:
			v, ok := h.getEffectiveValue(p, field)
			if !ok {
				return v, ok
			}

			if !found {
				return v, true
			}
			own.Type = v.Type // fix type information, should be brought up from class->proto->inst
		}
	}

	return own, found
}

type FoundLocation int

func (l FoundLocation) String() string {
	switch l {
	case FoundGlobal:
		return "global"
	case FoundParameter:
		return "parameter"
	case FoundLocal:
		return "local"
	case FoundField:
		return "field"
	}
	return "-"
}

const (
	// Global symbol
	FoundGlobal FoundLocation = iota
	// Function parameter symbol
	FoundParameter
	// Local symbol
	FoundLocal
	// Instance/prototype/class field symbol
	FoundField
)

type FoundSymbol struct {
	Symbol   symbol.Symbol
	Location FoundLocation
}

func (h *LspWorkspace) lookUpSymbol(documentURI string, position lsp.Position) (FoundSymbol, error) {
	var notFound FoundSymbol
	doc := h.bufferManager.GetBuffer(documentURI)
	if doc == "" {
		return notFound, fmt.Errorf("document %q not found", documentURI)
	}
	identifier, _ := doc.GetWordRangeAtPosition(position)

	if v, ok := h.lookUpScope(documentURI, position); ok {
		var ok bool
		switch s := v.(type) {
		case symbol.Function:
			for _, v := range s.Parameters {
				if strings.EqualFold(v.Name(), identifier) {
					return FoundSymbol{
						Symbol:   v,
						Location: FoundParameter,
					}, nil
				}
			}
			for _, v := range s.LocalVariables {
				if strings.EqualFold(v.Name(), identifier) {
					return FoundSymbol{
						Symbol:   v,
						Location: FoundLocal,
					}, nil
				}
			}
		case symbol.ProtoTypeOrInstance:
			v, ok = h.getEffectiveValue(s, identifier)
			if ok {
				return FoundSymbol{
					Symbol:   v,
					Location: FoundField,
				}, nil
			}
		}
	}

	symbol, found := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(identifier), SymbolAll)

	if !found {
		return notFound, fmt.Errorf("identifier %q not found", identifier)
	}
	return FoundSymbol{
		Symbol:   symbol,
		Location: FoundGlobal,
	}, nil
}

func getSymbolKind(s symbol.Symbol) lsp.SymbolKind {
	switch s.(type) {
	case symbol.ArrayVariable, symbol.ConstantArray:
		return lsp.Array
	case symbol.Class, symbol.ProtoTypeOrInstance:
		return lsp.Class
	case symbol.Function:
		return lsp.Function
	case symbol.Constant:
		return lsp.Constant
	case symbol.Variable:
		return lsp.Variable
	}
	return lsp.Null
}

func stringContainsAllAnywhere(value, set string) bool {
	found := true

	for _, v := range set {
		if !(strings.ContainsRune(value, unicode.ToLower(v)) ||
			strings.ContainsRune(value, unicode.ToUpper(v))) {
			found = false
			break
		}
	}

	return found
}
