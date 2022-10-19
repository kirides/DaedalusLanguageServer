package langserver

import (
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

type SemanticParseResult struct {
	ParseResult
	GlobalIdentifiers []Identifier
}

func (r *SemanticParseResult) CountSymbols() int64 {
	numSymbols := r.ParseResult.CountSymbols() + int64(len(r.GlobalIdentifiers))

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
func (p *SemanticParseResult) WalkSymbols(walkFn func(token) error) error {
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

// WalkScoped walks all higher-class symbols that are contained in `di`
func (p *SemanticParseResult) WalkScoped(bbox symbol.Definition, walkFn func(sym token) error) error {
	var err error
	scopedWalk := func(t token) error {
		if bbox.InBBox(t.Definition().Start) || bbox.InBBox(t.Definition().End) {
			return walkFn(t)
		}
		return nil
	}
	err = walkSymbolsMapHelper(p.Classes, scopedWalk, err)
	err = walkSymbolsMapHelper(p.GlobalConstants, scopedWalk, err)
	err = walkSymbolsMapHelper(p.Functions, scopedWalk, err)
	err = walkSymbolsMapHelper(p.Instances, scopedWalk, err)
	err = walkSymbolsMapHelper(p.Prototypes, scopedWalk, err)
	err = walkSymbolsMapHelper(p.GlobalVariables, scopedWalk, err)
	err = walkSymbolsSliceHelper(p.GlobalIdentifiers, scopedWalk, err)
	return err
}

type ScopedVariable struct {
	Location FoundLocation
}

func (parsedDoc *SemanticParseResult) FindScopedVariableDeclaration(di symbol.DefinitionIndex, name string) (FoundSymbol, bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if strings.EqualFold(p.Name(), name) {
					return FoundSymbol{p, FoundParameter}, true
				}
			}
			for _, p := range fn.LocalVariables {
				if strings.EqualFold(p.Name(), name) {
					return FoundSymbol{p, FoundLocal}, true
				}
			}
			break
		}
	}
	for _, fn := range parsedDoc.Classes {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Fields {
				if strings.EqualFold(p.Name(), name) {
					return FoundSymbol{p, FoundField}, true
				}
			}
			break
		}
	}
	for _, fn := range parsedDoc.Prototypes {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Fields {
				if strings.EqualFold(p.Name(), name) {
					return FoundSymbol{p, FoundField}, true
				}
			}
			break
		}
	}
	for _, fn := range parsedDoc.Instances {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Fields {
				if strings.EqualFold(p.Name(), name) {
					return FoundSymbol{p, FoundField}, true
				}
			}
			break
		}
	}
	return FoundSymbol{}, false
}

func (parsedDoc *SemanticParseResult) WalkScopedVariables(di symbol.DefinitionIndex, walkFn func(sym symbol.Symbol, isParam bool) bool) {
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
