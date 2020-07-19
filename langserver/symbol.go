package langserver

import "strings"

// Symbol ...
type Symbol interface {
	Name() string
	Source() string
	Documentation() string
	Definition() SymbolDefinition
	String() string
}

type symbolBase struct {
	stringerCache       string
	NameValue           string
	SymbolSource        string
	SymbolDocumentation string
	SymbolDefinition    SymbolDefinition
}

var _ Symbol = (*FunctionSymbol)(nil)
var _ Symbol = (*VariableSymbol)(nil)
var _ Symbol = (*ConstantSymbol)(nil)
var _ Symbol = (*ConstantSymbol)(nil)

// FunctionSymbol ...
type FunctionSymbol struct {
	symbolBase
	ReturnType     string
	Parameters     []VariableSymbol
	LocalVariables []Symbol
	BodyDefinition SymbolDefinition
}

// NewFunctionSymbol ...
func NewFunctionSymbol(name, source, doc string, def SymbolDefinition, retType string, bodyDef SymbolDefinition, params []VariableSymbol, locals []Symbol) FunctionSymbol {
	return FunctionSymbol{
		symbolBase: symbolBase{
			NameValue:           name,
			SymbolSource:        source,
			SymbolDocumentation: doc,
			SymbolDefinition:    def,
		},
		ReturnType:     retType,
		BodyDefinition: bodyDef,
		Parameters:     params,
		LocalVariables: locals,
	}
}

// NewVariableSymbol ...
func NewVariableSymbol(name, varType, source, documentation string, definiton SymbolDefinition) VariableSymbol {
	return VariableSymbol{
		Type: varType,
		symbolBase: symbolBase{
			NameValue:           name,
			SymbolSource:        source,
			SymbolDocumentation: documentation,
			SymbolDefinition:    definiton,
		},
	}
}

// NewConstantSymbol ...
func NewConstantSymbol(name, varType, source, documentation string, definiton SymbolDefinition, value string) ConstantSymbol {
	return ConstantSymbol{
		VariableSymbol: VariableSymbol{
			Type: varType,
			symbolBase: symbolBase{
				NameValue:           name,
				SymbolSource:        source,
				SymbolDocumentation: documentation,
				SymbolDefinition:    definiton,
			},
		},
		Value: value,
	}
}

// NewClassSymbol ...
func NewClassSymbol(name, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, fields []VariableSymbol) ClassSymbol {
	return ClassSymbol{
		symbolBase: symbolBase{
			NameValue:           name,
			SymbolSource:        source,
			SymbolDocumentation: documentation,
			SymbolDefinition:    definiton,
		},
		BodyDefinition: bodyDef,
		Fields:         fields,
	}
}

// NewPrototypeOrInstanceSymbol ...
func NewPrototypeOrInstanceSymbol(name, parent, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, isInstance bool) ProtoTypeOrInstanceSymbol {
	return ProtoTypeOrInstanceSymbol{
		symbolBase: symbolBase{
			NameValue:           name,
			SymbolSource:        source,
			SymbolDocumentation: documentation,
			SymbolDefinition:    definiton,
		},
		IsInstance:     isInstance,
		BodyDefinition: bodyDef,
	}
}

// TODO: Refactor for generics, `[]T where T: Stringer`
func joinVars(symbols []VariableSymbol) string {
	syms := make([]string, 0, len(symbols))
	for _, s := range symbols {
		syms = append(syms, s.String())
	}
	return strings.Join(syms, ", ")
}

// Name ...
func (s symbolBase) Name() string {
	return s.NameValue
}

// Source ...
func (s symbolBase) Source() string {
	return s.SymbolSource
}

// SetSource ...
func (s symbolBase) SetSource(src string) {
	s.SymbolSource = src
}

// Documentation ...
func (s symbolBase) Documentation() string {
	return s.SymbolDocumentation
}

// Definition ...
func (s symbolBase) Definition() SymbolDefinition {
	return s.SymbolDefinition
}

// String ...
func (s FunctionSymbol) String() string {
	return "func " + s.ReturnType + " " + s.Name() + "(" + joinVars(s.Parameters) + ")"
}

// VariableSymbol ...
type VariableSymbol struct {
	symbolBase
	Type string
}

// String ...
func (s VariableSymbol) String() string {
	return "var " + s.Type + " " + s.Name()
}

// ConstantSymbol ...
type ConstantSymbol struct {
	VariableSymbol
	Value string
}

// String ...
func (s ConstantSymbol) String() string {
	return "const " + s.Type + " " + s.Name() + " = " + s.Value
}

// ClassSymbol ...
type ClassSymbol struct {
	symbolBase
	Fields         []VariableSymbol
	BodyDefinition SymbolDefinition
}

// String ...
func (s ClassSymbol) String() string {
	return "class " + s.Name()
}

// ProtoTypeOrInstanceSymbol ...
type ProtoTypeOrInstanceSymbol struct {
	symbolBase
	IsInstance     bool
	Parent         string
	Fields         []VariableSymbol
	BodyDefinition SymbolDefinition
}

// String ...
func (s ProtoTypeOrInstanceSymbol) String() string {
	if s.IsInstance {
		return "instance " + s.Name() + "(" + s.Parent + ")"
	}
	return "prototype " + s.Name() + "(" + s.Parent + ")"
}

// SymbolDefinition ...
type SymbolDefinition struct {
	Start DefinitionIndex
	End   DefinitionIndex
}

// NewSymbolDefinition ...
func NewSymbolDefinition(startLine, startCol, endLine, endCol int) SymbolDefinition {
	return SymbolDefinition{
		Start: DefinitionIndex{
			Line:   startLine,
			Column: startCol,
		},
		End: DefinitionIndex{
			Line:   endLine,
			Column: endCol,
		},
	}
}

// InBBox ...
func (sd SymbolDefinition) InBBox(di DefinitionIndex) bool {
	if di.Line < sd.Start.Line {
		return false
	}
	if di.Line > sd.End.Line {
		return false
	}
	if sd.Start.Line <= di.Line && di.Line < sd.End.Line {
		return true
	}
	if di.Line == sd.End.Line && di.Column <= sd.End.Column {
		return true
	}
	if sd.Start.Line <= di.Line &&
		sd.Start.Column <= di.Column &&
		di.Line <= sd.End.Line &&
		di.Column <= sd.End.Column {
		return true
	}

	return false
}

// DefinitionIndex ...
type DefinitionIndex struct {
	Line   int
	Column int
}
