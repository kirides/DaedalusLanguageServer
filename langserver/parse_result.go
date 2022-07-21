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

func (r *ParseResult) CountSymbols() int64 {
	numSymbols := int64(len(r.Classes)) +
		int64(len(r.Functions)) +
		int64(len(r.GlobalConstants)) +
		int64(len(r.GlobalVariables)) +
		int64(len(r.Instances)) +
		int64(len(r.Prototypes))

	for _, v := range r.Classes {
		numSymbols += int64(len(v.Fields))
	}

	for _, v := range r.Instances {
		numSymbols += int64(len(v.Fields))
	}
	for _, v := range r.Prototypes {
		numSymbols += int64(len(v.Fields))
	}

	// for _, v := range r.Functions {
	// 	numSymbols += int64(len(v.Parameters))
	// 	numSymbols += int64(len(v.LocalVariables))
	// }

	return numSymbols
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
