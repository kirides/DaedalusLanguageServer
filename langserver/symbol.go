package langserver

import (
	"io"
	"strconv"
	"strings"
)

// Symbol ...
type Symbol interface {
	Name() string
	Source() string
	Documentation() string
	Definition() SymbolDefinition
	String() string
}

type symbolBase struct {
	NameValue           string
	SymbolSource        string
	SymbolDocumentation string
	SymbolDefinition    SymbolDefinition
}

var _ Symbol = (*FunctionSymbol)(nil)
var _ Symbol = (*ArrayVariableSymbol)(nil)
var _ Symbol = (*VariableSymbol)(nil)
var _ Symbol = (*ConstantSymbol)(nil)
var _ Symbol = (*FunctionSymbol)(nil)
var _ Symbol = (*ProtoTypeOrInstanceSymbol)(nil)
var _ Symbol = (*ClassSymbol)(nil)
var _ Symbol = (*ConstantArraySymbol)(nil)

// FunctionSymbol ...
type FunctionSymbol struct {
	ReturnType     string
	Parameters     []VariableSymbol
	LocalVariables []Symbol
	symbolBase
	BodyDefinition SymbolDefinition
}

// GetType ...
func (s FunctionSymbol) GetType() string {
	return s.ReturnType
}

func newSymbolBase(name, source, doc string, def SymbolDefinition) symbolBase {
	return symbolBase{
		NameValue:           name,
		SymbolSource:        source,
		SymbolDocumentation: doc,
		SymbolDefinition:    def,
	}
}

// NewFunctionSymbol ...
func NewFunctionSymbol(name, source, doc string, def SymbolDefinition, retType string, bodyDef SymbolDefinition, params []VariableSymbol, locals []Symbol) FunctionSymbol {
	return FunctionSymbol{
		symbolBase:     newSymbolBase(name, source, doc, def),
		ReturnType:     retType,
		BodyDefinition: bodyDef,
		Parameters:     params,
		LocalVariables: locals,
	}
}

// NewVariableSymbol ...
func NewVariableSymbol(name, varType, source, documentation string, definiton SymbolDefinition) VariableSymbol {
	return VariableSymbol{
		Type:       varType,
		symbolBase: newSymbolBase(name, source, documentation, definiton),
	}
}

// NewVariableSymbol ...
func NewArrayVariableSymbol(name, varType, sizeText, source, documentation string, definiton SymbolDefinition) ArrayVariableSymbol {
	return ArrayVariableSymbol{
		Type:          varType,
		ArraySizeText: sizeText,
		symbolBase:    newSymbolBase(name, source, documentation, definiton),
	}
}

// NewConstantSymbol ...
func NewConstantSymbol(name, varType, source, documentation string, definiton SymbolDefinition, value string) ConstantSymbol {
	return ConstantSymbol{
		VariableSymbol: NewVariableSymbol(name, varType, source, documentation, definiton),
		Value:          value,
	}
}

// NewConstantArraySymbol ...
func NewConstantArraySymbol(name, varType, arraySizeText, source, documentation string, definiton SymbolDefinition, value string) ConstantArraySymbol {
	return ConstantArraySymbol{
		VariableSymbol: NewVariableSymbol(name, varType, source, documentation, definiton),
		Value:          value,
		ArraySizeText:  arraySizeText,
	}
}

// NewClassSymbol ...
func NewClassSymbol(name, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, fields []Symbol) ClassSymbol {
	return ClassSymbol{
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		BodyDefinition: bodyDef,
		Fields:         fields,
	}
}

// NewPrototypeOrInstanceSymbol ...
func NewPrototypeOrInstanceSymbol(name, parent, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, isInstance bool) ProtoTypeOrInstanceSymbol {
	return ProtoTypeOrInstanceSymbol{
		Parent:         parent,
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		IsInstance:     isInstance,
		BodyDefinition: bodyDef,
	}
}

func writeParameterVariables(w io.StringWriter, symbols []VariableSymbol) {
	for _, s := range symbols {
		w.WriteString(s.String())
	}
}

// Name ...
func (s symbolBase) Name() string {
	return s.NameValue
}

// Source ...
func (s symbolBase) Source() string {
	return s.SymbolSource
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
	sb := strings.Builder{}
	sb.WriteString("func ")
	sb.WriteString(s.ReturnType)
	sb.WriteString(" ")
	sb.WriteString(s.Name())
	sb.WriteString("(")
	writeParameterVariables(&sb, s.Parameters)
	sb.WriteString(")")

	return sb.String()
}

// VariableSymbol ...
type VariableSymbol struct {
	Type string
	symbolBase
}

// String ...
func (s VariableSymbol) String() string {
	return "var " + s.Type + " " + s.Name()
}

// GetType ...
func (s VariableSymbol) GetType() string {
	return s.Type
}

// VariableSymbol ...
type ArrayVariableSymbol struct {
	Type          string
	ArraySizeText string
	symbolBase
}

// String ...
func (s ArrayVariableSymbol) String() string {
	return "var " + s.Type + " " + s.Name() + "[" + s.ArraySizeText + "]"
}

func (s ArrayVariableSymbol) Format(w io.StringWriter, resolvedSize int) {
	w.WriteString("var ")
	w.WriteString(s.Type)
	w.WriteString(" ")
	w.WriteString(s.Name())
	w.WriteString("[")
	w.WriteString(s.ArraySizeText)
	if resolvedSize != -1 {
		w.WriteString(" /* ")
		w.WriteString(strconv.Itoa(resolvedSize))
		w.WriteString(" */")
	}
	w.WriteString("]")
}

// GetType ...
func (s ArrayVariableSymbol) GetType() string {
	return s.Type
}

// ConstantSymbol ...
type ConstantSymbol struct {
	Value string
	VariableSymbol
}

// String ...
func (s ConstantSymbol) String() string {
	return "const " + s.Type + " " + s.Name() + " = " + s.Value
}

// GetType ...
func (s ConstantSymbol) GetType() string {
	return s.Type
}

type ConstantArraySymbol struct {
	Value         string
	ArraySizeText string
	VariableSymbol
}

// String ...
func (s ConstantArraySymbol) String() string {
	sb := strings.Builder{}
	s.Format(&sb, -1)
	return sb.String()
}

func (s ConstantArraySymbol) Format(w io.StringWriter, resolvedSize int) {
	w.WriteString("const ")
	w.WriteString(s.Type)
	w.WriteString(" ")
	w.WriteString(s.Name())
	w.WriteString("[")
	w.WriteString(s.ArraySizeText)
	if resolvedSize != -1 {
		w.WriteString(" /* ")
		w.WriteString(strconv.Itoa(resolvedSize))
		w.WriteString(" */")
	}
	w.WriteString("] = ")
	w.WriteString(s.Value)
}

// GetType ...
func (s ConstantArraySymbol) GetType() string {
	return s.Type
}

// ClassSymbol ...
type ClassSymbol struct {
	Fields []Symbol
	symbolBase
	BodyDefinition SymbolDefinition
}

// String ...
func (s ClassSymbol) String() string {
	return "class " + s.Name()
}

// ProtoTypeOrInstanceSymbol ...
type ProtoTypeOrInstanceSymbol struct {
	Parent string
	Fields []VariableSymbol
	symbolBase
	BodyDefinition SymbolDefinition
	IsInstance     bool
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
