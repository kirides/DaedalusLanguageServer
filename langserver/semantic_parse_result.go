package langserver

import "github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

type semanticParseResult struct {
	Instances         map[string]symbol.ProtoTypeOrInstance
	GlobalVariables   map[string]symbol.Symbol
	GlobalConstants   map[string]symbol.Symbol
	Functions         map[string]FunctionWithIdentifiers
	Classes           map[string]symbol.Class
	Prototypes        map[string]symbol.ProtoTypeOrInstance
	GlobalIdentifiers []Identifier
}

func (r *semanticParseResult) CountSymbols() int64 {
	numSymbols := int64(len(r.Classes)) +
		int64(len(r.Functions)) +
		int64(len(r.GlobalConstants)) +
		int64(len(r.GlobalVariables)) +
		int64(len(r.Instances)) +
		int64(len(r.Prototypes)) +
		int64(len(r.GlobalIdentifiers))

	for _, v := range r.Classes {
		numSymbols += int64(len(v.Fields))
	}

	for _, v := range r.Instances {
		numSymbols += int64(len(v.Fields))
	}
	for _, v := range r.Prototypes {
		numSymbols += int64(len(v.Fields))
	}

	for _, v := range r.Functions {
		numSymbols += int64(len(v.Parameters))
		numSymbols += int64(len(v.LocalVariables))
		numSymbols += int64(len(v.Identifiers))
	}

	return numSymbols
}

// WalkSymbols ...
func walkSymbolsMapHelper[V token](items map[string]V, walkFn func(token) error, previousError error) error {
	if previousError != nil {
		return previousError
	}

	for _, s := range items {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	return nil
}
func walkSymbolsSliceHelper[V token](items []V, walkFn func(token) error, previousError error) error {
	if previousError != nil {
		return previousError
	}

	for _, s := range items {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	return nil
}
func (p *semanticParseResult) WalkSymbols(walkFn func(token) error) error {
	var err error
	err = walkSymbolsMapHelper(p.Classes, walkFn, err)
	err = walkSymbolsMapHelper(p.GlobalConstants, walkFn, err)
	err = walkSymbolsMapHelper(p.Functions, walkFn, err)
	err = walkSymbolsMapHelper(p.Instances, walkFn, err)
	err = walkSymbolsMapHelper(p.Prototypes, walkFn, err)
	err = walkSymbolsMapHelper(p.GlobalVariables, walkFn, err)
	err = walkSymbolsSliceHelper(p.GlobalIdentifiers, walkFn, err)
	return err
}

func (parsedDoc *semanticParseResult) WalkScopedVariables(di symbol.DefinitionIndex, walkFn func(sym symbol.Symbol, isParam bool) bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if !walkFn(p, true) {
					return
				}
			}
			for _, p := range fn.LocalVariables {
				if !walkFn(p, false) {
					return
				}
			}
			break
		}
	}
}
