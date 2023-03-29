package langserver

import (
	"strings"
	"time"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

// ParseResult ...
type ParseResult struct {
	Ast             parser.IDaedalusFileContext
	Instances       map[string]symbol.ProtoTypeOrInstance
	GlobalVariables map[string]symbol.Symbol
	GlobalConstants map[string]symbol.Symbol
	Functions       map[string]symbol.Function
	Classes         map[string]symbol.Class
	Prototypes      map[string]symbol.ProtoTypeOrInstance
	Namespaces      map[string]symbol.Namespace
	Source          string
	SyntaxErrors    []SyntaxError

	lastModifiedAt time.Time
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

func (parsedDoc *ParseResult) LookupScopedVariable(di symbol.DefinitionIndex, name string) (symbol.Symbol, bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if strings.EqualFold(p.Name(), name) {
					return p, true
				}
			}
			for _, p := range fn.LocalVariables {
				if strings.EqualFold(p.Name(), name) {
					return p, true
				}
			}
			break
		}
	}
	for _, fn := range parsedDoc.Classes {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Fields {
				if strings.EqualFold(p.Name(), name) {
					return p, true
				}
			}
			break
		}
	}
	return nil, false
}

func (parsedDoc *ParseResult) WalkScopedVariables(di symbol.DefinitionIndex, walkFn func(symbol.Symbol) bool) {
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
	for _, fn := range parsedDoc.Classes {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Fields {
				if !walkFn(p) {
					return
				}
			}
			break
		}
	}
}

func (parsedDoc *ParseResult) FindContainingScope(di symbol.DefinitionIndex) (symbol.Symbol, bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			return fn, true
		}
	}
	for _, fn := range parsedDoc.Classes {
		if fn.BodyDefinition.InBBox(di) {
			return fn, true
		}
	}
	for _, fn := range parsedDoc.Prototypes {
		if fn.BodyDefinition.InBBox(di) {
			return fn, true
		}
	}
	for _, fn := range parsedDoc.Instances {
		if fn.BodyDefinition.InBBox(di) {
			return fn, true
		}
	}
	return nil, false
}

func (parsedDoc *ParseResult) ScopedVariables(di symbol.DefinitionIndex) map[string]symbol.Symbol {
	locals := map[string]symbol.Symbol{}

	parsedDoc.WalkScopedVariables(di, func(s symbol.Symbol) bool {
		locals[strings.ToUpper(s.Name())] = s
		return true
	})

	return locals
}
