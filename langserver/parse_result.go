package langserver

import (
	"strings"
)

// ParseResult ...
type ParseResult struct {
	Instances       map[string]ProtoTypeOrInstanceSymbol
	GlobalVariables map[string]VariableSymbol
	GlobalConstants map[string]ConstantSymbol
	Functions       map[string]FunctionSymbol
	Classes         map[string]ClassSymbol
	Prototypes      map[string]ProtoTypeOrInstanceSymbol
	Source          string
	SyntaxErrors    []SyntaxError
}

func (parsedDoc *ParseResult) WalkScopedVariables(di DefinitionIndex, walkFn func(Symbol) bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if !walkFn(p) {
					return
				}
			}
			for _, p := range fn.LocalVariables {
				if !walkFn(p) {
					return
				}
			}
			break
		}
	}
}

func (parsedDoc *ParseResult) ScopedVariables(di DefinitionIndex) map[string]Symbol {
	locals := map[string]Symbol{}

	parsedDoc.WalkScopedVariables(di, func(s Symbol) bool {
		locals[strings.ToUpper(s.Name())] = s
		return true
	})

	return locals
}
