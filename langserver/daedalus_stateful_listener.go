package langserver

import (
	"bytes"
	"langsrv/langserver/parser"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// DaedalusStatefulListener ...
type DaedalusStatefulListener struct {
	parser.BaseDaedalusListener

	source          string
	GlobalVariables map[string]VariableSymbol
	GlobalConstants map[string]ConstantSymbol
	Functions       map[string]FunctionSymbol
	Classes         map[string]ClassSymbol
	Prototypes      map[string]ProtoTypeOrInstanceSymbol
	Instances       map[string]ProtoTypeOrInstanceSymbol

	summaryBuilder *bytes.Buffer
	knownSymbols   symbolWalker
}

type symbolWalker interface {
	WalkGlobalSymbols(walkFn func(Symbol) error, types SymbolType) error
	LookupGlobalSymbol(name string, types SymbolType) (Symbol, bool)
}

// NewDaedalusStatefulListener ...
func NewDaedalusStatefulListener(source string, knownSymbols symbolWalker) *DaedalusStatefulListener {
	return &DaedalusStatefulListener{
		source:          source,
		GlobalVariables: map[string]VariableSymbol{},
		GlobalConstants: map[string]ConstantSymbol{},
		Functions:       map[string]FunctionSymbol{},
		Classes:         map[string]ClassSymbol{},
		Prototypes:      map[string]ProtoTypeOrInstanceSymbol{},
		Instances:       map[string]ProtoTypeOrInstanceSymbol{},
		summaryBuilder:  &bytes.Buffer{},
		knownSymbols:    knownSymbols,
	}
}

var constDefContextType = reflect.TypeOf((*parser.IConstDefContext)(nil)).Elem()
var varDeclContextType = reflect.TypeOf((*parser.IVarDeclContext)(nil)).Elem()

var classDefContextType = reflect.TypeOf((*parser.IClassDefContext)(nil)).Elem()
var prototypeDefContextType = reflect.TypeOf((*parser.IPrototypeDefContext)(nil)).Elem()
var instanceDefContextType = reflect.TypeOf((*parser.IInstanceDefContext)(nil)).Elem()

func symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) SymbolDefinition {
	return NewSymbolDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

func (l *DaedalusStatefulListener) symbolSummaryForContext(summaries []parser.ISymbolSummaryContext) string {
	if summaries == nil || len(summaries) == 0 {
		return ""
	}

	l.summaryBuilder.Reset()

	for i, isummary := range summaries {
		sum := isummary.(*parser.SymbolSummaryContext)
		if i > 0 {
			l.summaryBuilder.WriteRune('\n')
		}
		sumReplaced := strings.TrimLeft(sum.GetText(), "/ \t")
		sumReplaced = strings.ReplaceAll(sumReplaced, "/// ", "  \n")
		l.summaryBuilder.WriteString(sumReplaced)
	}

	return l.summaryBuilder.String()
}

func (l *DaedalusStatefulListener) variablesFromContext(v *parser.VarDeclContext) []Symbol {
	result := []Symbol{}
	summary := ""
	if p, ok := v.GetParent().(*parser.InlineDefContext); ok {
		summary = l.symbolSummaryForContext(p.AllSymbolSummary())
	}
	for _, ival := range v.AllVarValueDecl() {
		val := ival.(*parser.VarValueDeclContext)
		result = append(result,
			NewVariableSymbol(val.NameNode().GetText(),
				v.TypeReference().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(val.NameNode()),
			))
	}

	for _, iInnerVar := range v.AllVarDecl() {
		innerVal := iInnerVar.(*parser.VarDeclContext)
		for _, ival := range innerVal.AllVarValueDecl() {
			val := ival.(*parser.VarValueDeclContext)
			result = append(result,
				NewVariableSymbol(val.NameNode().GetText(),
					v.TypeReference().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(val.NameNode()),
				))
		}
	}
	for _, iInnerVar := range v.AllVarArrayDecl() {
		innerVal := iInnerVar.(*parser.VarArrayDeclContext)
		if innerVal != nil {
			result = append(result,
				NewVariableSymbol(innerVal.NameNode().GetText(),
					v.TypeReference().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(innerVal.NameNode()),
				))
		}
	}

	return result
}

func (l *DaedalusStatefulListener) maxNOfConstValues(n int, c *parser.ConstArrayAssignmentContext) string {
	result := "{ "
	counter := 0
	for i, v := range c.AllExpressionBlock() {
		if i >= n {
			break
		}
		counter++
		if i > 0 {
			result += ", " + v.GetText()
		} else {
			result += v.GetText()
		}
	}
	if counter == 0 {
		return result + "}"
	}
	return result + " ... }"
}

func (l *DaedalusStatefulListener) constsFromContext(c *parser.ConstDefContext) []Symbol {
	result := []Symbol{}
	summary := ""
	if p, ok := c.GetParent().(*parser.InlineDefContext); ok {
		summary = l.symbolSummaryForContext(p.AllSymbolSummary())
	}
	for _, icv := range c.AllConstValueDef() {
		cv := icv.(*parser.ConstValueDefContext)

		result = append(result,
			NewConstantSymbol(cv.NameNode().GetText(),
				c.TypeReference().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(cv.NameNode()),
				cv.ConstValueAssignment().(*parser.ConstValueAssignmentContext).ExpressionBlock().GetText(),
			))
	}

	for _, iInnerVar := range c.AllConstArrayDef() {
		innerVal := iInnerVar.(*parser.ConstArrayDefContext)
		result = append(result,
			NewConstantSymbol(innerVal.NameNode().GetText(),
				c.TypeReference().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(innerVal.NameNode()),
				l.maxNOfConstValues(3, innerVal.ConstArrayAssignment().(*parser.ConstArrayAssignmentContext)),
			))
	}

	return result
}

// EnterInlineDef ...
func (l *DaedalusStatefulListener) EnterInlineDef(ctx *parser.InlineDefContext) {

	consts := ctx.GetTypedRuleContexts(constDefContextType) // Does not really work somehow ...
	for _, ic := range consts {
		c, ok := ic.(*parser.ConstDefContext)
		if !ok {
			continue
		}
		for _, s := range l.constsFromContext(c) {
			l.GlobalConstants[strings.ToUpper(s.Name())] = s.(ConstantSymbol)
		}
	}
	vars := ctx.GetTypedRuleContexts(varDeclContextType) // Does not really work somehow ...
	for _, iv := range vars {
		v, ok := iv.(*parser.VarDeclContext)
		if !ok {
			continue
		}
		for _, s := range l.variablesFromContext(v) {
			l.GlobalVariables[strings.ToUpper(s.Name())] = s.(VariableSymbol)
		}
	}
}

// EnterBlockDef ...
func (l *DaedalusStatefulListener) EnterBlockDef(ctx *parser.BlockDefContext) {
	// Classes

	classes := ctx.GetTypedRuleContexts(classDefContextType)
	for _, ic := range classes {
		c, ok := ic.(*parser.ClassDefContext)
		if !ok {
			continue
		}
		cFields := []VariableSymbol{}
		for _, iv := range c.AllVarDecl() {
			v := iv.(*parser.VarDeclContext)
			for _, ivv := range v.AllVarValueDecl() {
				vv := ivv.(*parser.VarValueDeclContext)
				cFields = append(cFields,
					NewVariableSymbol(vv.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						"",
						symbolDefinitionForRuleContext(vv.NameNode())),
				)
			}
		}
		csym := NewClassSymbol(c.NameNode().GetText(),
			l.source,
			l.symbolSummaryForContext(c.AllSymbolSummary()),
			symbolDefinitionForRuleContext(c.NameNode()),
			SymbolDefinition{
				Start: DefinitionIndex{
					Line:   c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetLine(),
					Column: c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetColumn(),
				},
				End: DefinitionIndex{
					Line:   c.GetChild(c.GetChildCount() - 1).(antlr.TerminalNode).GetSymbol().GetLine(),
					Column: c.GetChild(c.GetChildCount() - 1).(antlr.TerminalNode).GetSymbol().GetColumn(),
				},
			},
			cFields)
		l.Classes[strings.ToUpper(csym.Name())] = csym
	}

	prototypes := ctx.GetTypedRuleContexts(prototypeDefContextType)
	for _, ic := range prototypes {
		c, ok := ic.(*parser.PrototypeDefContext)
		if !ok {
			continue
		}
		psym := NewPrototypeOrInstanceSymbol(
			c.NameNode().GetText(),
			c.ParentReference().(*parser.ParentReferenceContext).Identifier().GetText(),
			l.source,
			"",
			symbolDefinitionForRuleContext(c.NameNode()),
			symbolDefinitionForRuleContext(c.StatementBlock()),
			false)
		l.Prototypes[strings.ToUpper(psym.Name())] = psym
	}

	instances := ctx.GetTypedRuleContexts(instanceDefContextType)
	for _, ic := range instances {
		c, ok := ic.(*parser.InstanceDefContext)
		if !ok {
			continue
		}
		psym := NewPrototypeOrInstanceSymbol(
			c.NameNode().GetText(),
			c.ParentReference().(*parser.ParentReferenceContext).Identifier().GetText(),
			l.source,
			"",
			symbolDefinitionForRuleContext(c.NameNode()),
			symbolDefinitionForRuleContext(c.StatementBlock()),
			true)
		l.Prototypes[strings.ToUpper(psym.Name())] = psym
	}
}

func (l *DaedalusStatefulListener) findVarsConstsInStatements(root antlr.Tree) []Symbol {
	result := []Symbol{}

	for _, s := range root.GetChildren() {
		if varDecl, ok := s.(*parser.VarDeclContext); ok {
			result = append(result, l.variablesFromContext(varDecl)...)
		} else if constDef, ok := s.(*parser.ConstDefContext); ok {
			result = append(result, l.constsFromContext(constDef)...)
		} else if s.GetChildCount() > 0 {
			result = append(result, l.findVarsConstsInStatements(s)...)
		}
	}

	return result
}

// EnterFunctionDef ...
func (l *DaedalusStatefulListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	statements := ctx.StatementBlock().(*parser.StatementBlockContext)

	params := []VariableSymbol{}
	locals := l.findVarsConstsInStatements(statements)

	for _, s := range ctx.ParameterList().(*parser.ParameterListContext).AllParameterDecl() {
		pdef := s.(*parser.ParameterDeclContext)
		params = append(params,
			NewVariableSymbol(pdef.NameNode().GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(pdef.NameNode())))
	}
	fnc := NewFunctionSymbol(ctx.NameNode().GetText(),
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext).AllSymbolSummary()),
		symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	l.Functions[strings.ToUpper(fnc.Name())] = fnc
}
