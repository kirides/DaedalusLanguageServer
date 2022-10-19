package symbol

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
	Definition() Definition
	String() string
}

type symbolBase struct {
	NameValue           string
	SymbolSource        string
	SymbolDocumentation string
	SymbolDefinition    Definition
}

var _ Symbol = (*Function)(nil)
var _ Symbol = (*ArrayVariable)(nil)
var _ Symbol = (*Variable)(nil)
var _ Symbol = (*Constant)(nil)
var _ Symbol = (*ProtoTypeOrInstance)(nil)
var _ Symbol = (*Class)(nil)
var _ Symbol = (*ConstantArray)(nil)

func newSymbolBase(name, source, doc string, def Definition) symbolBase {
	return symbolBase{
		NameValue:           name,
		SymbolSource:        source,
		SymbolDocumentation: doc,
		SymbolDefinition:    def,
	}
}

// NewFunction ...
func NewFunction(name, source, doc string, def Definition, retType string, bodyDef Definition, params []Variable, locals []Symbol) Function {
	return Function{
		symbolBase:     newSymbolBase(name, source, doc, def),
		ReturnType:     retType,
		BodyDefinition: bodyDef,
		Parameters:     params,
		LocalVariables: locals,
	}
}

// NewVariable ...
func NewVariable(name, varType, source, documentation string, definiton Definition) Variable {
	return Variable{
		Type:       varType,
		symbolBase: newSymbolBase(name, source, documentation, definiton),
	}
}

// NewVariableSymbol ...
func NewArrayVariable(name, varType, sizeText, source, documentation string, definiton Definition) ArrayVariable {
	return ArrayVariable{
		Type:          varType,
		ArraySizeText: sizeText,
		symbolBase:    newSymbolBase(name, source, documentation, definiton),
	}
}

// NewConstant ...
func NewConstant(name, varType, source, documentation string, definiton Definition, value string) Constant {
	return Constant{
		Variable: NewVariable(name, varType, source, documentation, definiton),
		Value:    value,
	}
}

// NewConstantArray ...
func NewConstantArray(name, varType, arraySizeText, source, documentation string, definiton Definition, value string) ConstantArray {
	return ConstantArray{
		Variable:      NewVariable(name, varType, source, documentation, definiton),
		Value:         value,
		ArraySizeText: arraySizeText,
	}
}

// NewClass ...
func NewClass(name, source, documentation string, definiton Definition, bodyDef Definition, fields []Symbol) Class {
	return Class{
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		BodyDefinition: bodyDef,
		Fields:         fields,
	}
}

// NewPrototypeOrInstance ...
func NewPrototypeOrInstance(name, parent, source, documentation string, definiton Definition, bodyDef Definition, isInstance bool) ProtoTypeOrInstance {
	return ProtoTypeOrInstance{
		Parent:         parent,
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		IsInstance:     isInstance,
		BodyDefinition: bodyDef,
	}
}

func writeParameterVariables(w io.StringWriter, symbols []Variable) {
	for i, s := range symbols {
		if i > 0 {
			w.WriteString(", ")
		}
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
func (s symbolBase) Definition() Definition {
	return s.SymbolDefinition
}

// Function ...
type Function struct {
	ReturnType     string
	Parameters     []Variable
	LocalVariables []Symbol
	symbolBase
	BodyDefinition Definition
}

// GetType ...
func (s Function) GetType() string {
	return s.ReturnType
}

// String ...
func (s Function) String() string {
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

// Variable ...
type Variable struct {
	Type string
	symbolBase
}

// String ...
func (s Variable) String() string {
	return "var " + s.Type + " " + s.Name()
}

// GetType ...
func (s Variable) GetType() string {
	return s.Type
}

// VariableSymbol ...
type ArrayVariable struct {
	Type          string
	ArraySizeText string
	symbolBase
}

// String ...
func (s ArrayVariable) String() string {
	return "var " + s.Type + " " + s.Name() + "[" + s.ArraySizeText + "]"
}

func (s ArrayVariable) Format(w io.StringWriter, resolvedSize int) {
	w.WriteString("var ")
	w.WriteString(s.Type)
	w.WriteString(" ")
	w.WriteString(s.Name())
	w.WriteString("[")
	w.WriteString(s.ArraySizeText)
	if resolvedSize != -1 {
		w.WriteString(":")
		w.WriteString(strconv.Itoa(resolvedSize))
	}
	w.WriteString("]")
}

// GetType ...
func (s ArrayVariable) GetType() string {
	return s.Type
}

// Constant ...
type Constant struct {
	Value string
	Variable
}

// String ...
func (s Constant) String() string {
	return "const " + s.Type + " " + s.Name() + " = " + s.Value
}

// GetType ...
func (s Constant) GetType() string {
	return s.Type
}

type ConstantArray struct {
	Value         string
	ArraySizeText string
	Variable
}

// String ...
func (s ConstantArray) String() string {
	sb := strings.Builder{}
	s.Format(&sb, -1)
	return sb.String()
}

func (s ConstantArray) Format(w io.StringWriter, resolvedSize int) {
	w.WriteString("const ")
	w.WriteString(s.Type)
	w.WriteString(" ")
	w.WriteString(s.Name())
	w.WriteString("[")
	w.WriteString(s.ArraySizeText)
	if resolvedSize != -1 {
		w.WriteString(":")
		w.WriteString(strconv.Itoa(resolvedSize))
	}
	w.WriteString("] = ")
	w.WriteString(s.Value)
}

// GetType ...
func (s ConstantArray) GetType() string {
	return s.Type
}

// Class ...
type Class struct {
	Fields []Symbol
	symbolBase
	BodyDefinition Definition
}

// String ...
func (s Class) String() string {
	return "class " + s.Name()
}

// ProtoTypeOrInstance ...
type ProtoTypeOrInstance struct {
	Parent string
	Fields []Constant
	symbolBase
	BodyDefinition Definition
	IsInstance     bool
}

// String ...
func (s ProtoTypeOrInstance) String() string {
	if s.IsInstance {
		return "instance " + s.Name() + "(" + s.Parent + ")"
	}
	return "prototype " + s.Name() + "(" + s.Parent + ")"
}

// Definition ...
type Definition struct {
	Start DefinitionIndex
	End   DefinitionIndex
}

// NewDefinition ...
func NewDefinition(startLine, startCol, endLine, endCol int) Definition {
	return Definition{
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
func (sd Definition) InBBox(di DefinitionIndex) bool {
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
