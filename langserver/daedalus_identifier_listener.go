package langserver

import (
	"fmt"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Identifier struct {
	nameValue  string
	definition symbol.Definition
}

func (s Identifier) Name() string                  { return s.nameValue }
func (s Identifier) Definition() symbol.Definition { return s.definition }

// Function ...
type FunctionWithIdentifiers struct {
	symbol.Function
	Identifiers []Identifier
}

// GetType ...
func (s FunctionWithIdentifiers) GetType() string {
	return s.ReturnType
}

// String ...
func (s FunctionWithIdentifiers) String() string {
	return s.Function.String()
}

// DaedalusIdentifierListener ...
type DaedalusIdentifierListener struct {
	parser.BaseDaedalusListener

	Functions         map[string]FunctionWithIdentifiers
	GlobalIdentifiers []Identifier
	source            string

	currentIdentifiers []Identifier
}

// NewDaedalusIdentifierListener ...
func NewDaedalusIdentifierListener(source string) *DaedalusIdentifierListener {
	return &DaedalusIdentifierListener{
		source:             source,
		Functions:          map[string]FunctionWithIdentifiers{},
		GlobalIdentifiers:  []Identifier{},
		currentIdentifiers: []Identifier{},
	}
}

func (l *DaedalusIdentifierListener) variablesFromContext(v *parser.VarDeclContext) []symbol.Symbol {
	result := []symbol.Symbol{}

	for _, ch := range v.GetChildren() {
		if val, ok := ch.(*parser.VarValueDeclContext); ok {
			result = append(result,
				symbol.NewVariable(val.NameNode().GetText(),
					v.TypeReference().GetText(),
					l.source,
					"", // documentation
					symbolDefinitionForRuleContext(val.NameNode()),
				))
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ival := range innerVal.AllVarValueDecl() {
				val := ival.(*parser.VarValueDeclContext)
				result = append(result,
					symbol.NewVariable(val.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						"", // documentation
						symbolDefinitionForRuleContext(val.NameNode()),
					))
			}
		} else if innerVal, ok := ch.(*parser.VarArrayDeclContext); ok {
			result = append(result,
				symbol.NewArrayVariable(innerVal.NameNode().GetText(),
					v.TypeReference().GetText(),
					innerVal.ArraySize().GetText(),
					l.source,
					"", // documentation
					symbolDefinitionForRuleContext(innerVal.NameNode()),
				))
		}
	}

	return result
}

func (l *DaedalusIdentifierListener) maxNOfConstValues(n int, c *parser.ConstArrayAssignmentContext) string {
	result := &strings.Builder{}
	result.WriteString("{ ")
	counter := 0

	walkSymbols(c, func(v parser.IExpressionBlockContext) error {
		if counter > n {
			return fmt.Errorf("ok")
		}
		if counter > 0 {
			result.WriteString(", ")
		}
		result.WriteString(v.(*parser.ExpressionBlockContext).GetText())
		counter++
		return nil
	})

	if counter == 0 {
		result.WriteString("}")
	} else {
		result.WriteString(" ... }")
	}
	return result.String()
}

func (l *DaedalusIdentifierListener) constsFromContext(c *parser.ConstDefContext) []symbol.Symbol {
	result := []symbol.Symbol{}
	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		result = append(result,
			symbol.NewConstant(cv.NameNode().GetText(),
				c.TypeReference().GetText(),
				l.source,
				"", // documentation
				symbolDefinitionForRuleContext(cv.NameNode()),
				cv.ConstValueAssignment().(*parser.ConstValueAssignmentContext).ExpressionBlock().GetText(),
			))
		return nil
	})

	walkSymbols(c, func(innerVal *parser.ConstArrayDefContext) error {
		result = append(result,
			symbol.NewConstantArray(innerVal.NameNode().GetText(),
				c.TypeReference().GetText(),
				innerVal.ArraySize().GetText(),
				l.source,
				"", // documentation
				symbolDefinitionForRuleContext(innerVal.NameNode()),
				l.maxNOfConstValues(3, innerVal.ConstArrayAssignment().(*parser.ConstArrayAssignmentContext)),
			))
		return nil
	})

	return result
}

func (l *DaedalusIdentifierListener) findVarsConstsInStatements(root antlr.Tree) []symbol.Symbol {
	result := []symbol.Symbol{}

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

func (l *DaedalusIdentifierListener) onIdentifier(ctx antlr.ParserRuleContext) {
	txt := ctx.GetText()

	// only plain tokens for now
	if strings.ContainsAny(txt, ".[]") || // no arrays or properties
		// strings.EqualFold(txt, "null") || // people actually use "null" as variable names
		strings.EqualFold(txt, "true") ||
		strings.EqualFold(txt, "false") ||

		strings.EqualFold(txt, "hero") ||
		strings.EqualFold(txt, "self") ||
		strings.EqualFold(txt, "other") ||
		strings.EqualFold(txt, "victim") ||

		strings.EqualFold(txt, "end") ||
		strings.EqualFold(txt, "repeat") ||
		strings.EqualFold(txt, "while") ||

		strings.EqualFold(txt, "item") {
		return
	}

	l.currentIdentifiers = append(l.currentIdentifiers, Identifier{
		nameValue:  txt,
		definition: symbolDefinitionForRuleContext(ctx),
	})
}

func (l *DaedalusIdentifierListener) EnterReferenceValue(ctx *parser.ReferenceValueContext) {
	l.onIdentifier(ctx)
}

func (l *DaedalusIdentifierListener) EnterReferenceAtom(ctx *parser.ReferenceAtomContext) {
	l.onIdentifier(ctx)
}

func (l *DaedalusIdentifierListener) EnterNullLiteralValue(ctx *parser.NullLiteralValueContext) {
	l.onIdentifier(ctx)
}

func (l *DaedalusIdentifierListener) copyIdentifiers(o []Identifier) []Identifier {
	if len(l.currentIdentifiers) > 0 {
		o = append(o, l.currentIdentifiers...)
		l.currentIdentifiers = l.currentIdentifiers[:0]
	}
	return o
}

func (l *DaedalusIdentifierListener) ExitDaedalusFile(ctx *parser.DaedalusFileContext) {
	l.GlobalIdentifiers = l.copyIdentifiers(l.GlobalIdentifiers)
}

func (l *DaedalusIdentifierListener) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	statements := ctx.StatementBlock().(*parser.StatementBlockContext)

	params := []symbol.Variable{}
	locals := l.findVarsConstsInStatements(statements)

	walkSymbols(ctx.ParameterList(), func(pdef *parser.ParameterDeclContext) error {
		params = append(params,
			symbol.NewVariable(pdef.NameNode().GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(pdef.NameNode())))
		return nil
	})

	fnc := symbol.NewFunction(ctx.NameNode().GetText(),
		l.source,
		"",
		symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	identifiers := []Identifier{}
	identifiers = l.copyIdentifiers(identifiers)

	l.Functions[strings.ToUpper(fnc.Name())] = FunctionWithIdentifiers{
		Function:    fnc,
		Identifiers: identifiers,
	}
}

// EnterFunctionDef ...
func (l *DaedalusIdentifierListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	l.GlobalIdentifiers = l.copyIdentifiers(l.GlobalIdentifiers)
}
