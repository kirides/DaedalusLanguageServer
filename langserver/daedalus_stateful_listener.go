package langserver

import (
	"bytes"
	"fmt"
	"langsrv/langserver/parser"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// DaedalusStatefulListener ...
type DaedalusStatefulListener struct {
	parser.BaseDaedalusListener
	knownSymbols    symbolWalker
	Instances       map[string]ProtoTypeOrInstanceSymbol
	GlobalVariables map[string]Symbol
	Functions       map[string]FunctionSymbol
	Classes         map[string]ClassSymbol
	Prototypes      map[string]ProtoTypeOrInstanceSymbol
	summaryBuilder  *bytes.Buffer
	GlobalConstants map[string]Symbol
	source          string
}

type symbolWalker interface {
	WalkGlobalSymbols(walkFn func(Symbol) error, types SymbolType) error
	LookupGlobalSymbol(name string, types SymbolType) (Symbol, bool)
}

// NewDaedalusStatefulListener ...
func NewDaedalusStatefulListener(source string, knownSymbols symbolWalker) *DaedalusStatefulListener {
	return &DaedalusStatefulListener{
		source:          source,
		GlobalVariables: map[string]Symbol{},
		GlobalConstants: map[string]Symbol{},
		Functions:       map[string]FunctionSymbol{},
		Classes:         map[string]ClassSymbol{},
		Prototypes:      map[string]ProtoTypeOrInstanceSymbol{},
		Instances:       map[string]ProtoTypeOrInstanceSymbol{},
		summaryBuilder:  &bytes.Buffer{},
		knownSymbols:    knownSymbols,
	}
}

func symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) SymbolDefinition {
	return NewSymbolDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

func (l *DaedalusStatefulListener) symbolSummaryForContext(ctx antlr.ParserRuleContext) string {
	l.summaryBuilder.Reset()

	walkSymbols(ctx, func(sum *parser.SymbolSummaryContext) error {
		if l.summaryBuilder.Len() > 0 {
			l.summaryBuilder.WriteRune('\n')
		}
		sumReplaced := strings.TrimLeft(sum.GetText(), "/ \t")
		sumReplaced = strings.ReplaceAll(sumReplaced, "/// ", "  \n")
		l.summaryBuilder.WriteString(sumReplaced)

		return nil
	})

	return l.summaryBuilder.String()
}

func (l *DaedalusStatefulListener) variablesFromContext(v *parser.VarDeclContext) []Symbol {
	result := []Symbol{}
	summary := ""
	if p, ok := v.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
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
		result = append(result,
			NewArrayVariableSymbol(innerVal.NameNode().GetText(),
				v.TypeReference().GetText(),
				innerVal.ArraySize().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(innerVal.NameNode()),
			))
	}

	return result
}

func (l *DaedalusStatefulListener) maxNOfConstValues(n int, c *parser.ConstArrayAssignmentContext) string {
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

func (l *DaedalusStatefulListener) constsFromContext(c *parser.ConstDefContext) []Symbol {
	result := []Symbol{}
	summary := ""
	if p, ok := c.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		result = append(result,
			NewConstantSymbol(cv.NameNode().GetText(),
				c.TypeReference().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(cv.NameNode()),
				cv.ConstValueAssignment().(*parser.ConstValueAssignmentContext).ExpressionBlock().GetText(),
			))
		return nil
	})

	walkSymbols(c, func(innerVal *parser.ConstArrayDefContext) error {
		result = append(result,
			NewConstantArraySymbol(innerVal.NameNode().GetText(),
				c.TypeReference().GetText(),
				innerVal.ArraySize().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(innerVal.NameNode()),
				l.maxNOfConstValues(3, innerVal.ConstArrayAssignment().(*parser.ConstArrayAssignmentContext)),
			))
		return nil
	})

	return result
}

// EnterInlineDef ...
func (l *DaedalusStatefulListener) EnterInlineDef(ctx *parser.InlineDefContext) {

	walkSymbols(ctx, func(c *parser.ConstDefContext) error {
		for _, s := range l.constsFromContext(c) {
			l.GlobalConstants[strings.ToUpper(s.Name())] = s
		}
		return nil
	})

	walkSymbols(ctx, func(v *parser.VarDeclContext) error {
		for _, s := range l.variablesFromContext(v) {
			l.GlobalVariables[strings.ToUpper(s.Name())] = s
		}
		return nil
	})

	walkSymbols(ctx, func(c *parser.InstanceDeclContext) error {
		walkSymbols(c, func(name parser.INameNodeContext) error {
			psym := NewPrototypeOrInstanceSymbol(
				name.GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(name),
				symbolDefinitionForRuleContext(c.ParentReference()),
				true)
			l.Instances[strings.ToUpper(psym.Name())] = psym
			return nil
		})
		return nil
	})
}

func walkSymbols[T antlr.ParserRuleContext](rule antlr.RuleNode, walkFn func(symbol T) error) {
	for _, v := range rule.GetChildren() {
		if s, ok := v.(T); ok {
			if err := walkFn(s); err != nil {
				return
			}
		}
	}
}

// EnterBlockDef ...
func (l *DaedalusStatefulListener) EnterBlockDef(ctx *parser.BlockDefContext) {
	// Classes
	walkSymbols(ctx, func(c *parser.ClassDefContext) error {
		cFields := []Symbol{}
		walkSymbols(c, func(v *parser.VarDeclContext) error {
			walkSymbols(v, func(vv *parser.VarValueDeclContext) error {
				cFields = append(cFields,
					NewVariableSymbol(vv.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						"",
						symbolDefinitionForRuleContext(vv.NameNode())),
				)
				return nil
			})

			walkSymbols(v, func(vv *parser.VarArrayDeclContext) error {
				cFields = append(cFields,
					NewArrayVariableSymbol(vv.NameNode().GetText(),
						v.TypeReference().GetText(),
						vv.ArraySize().GetText(),
						l.source,
						"",
						symbolDefinitionForRuleContext(vv.NameNode())),
				)
				return nil
			})
			return nil
		})

		csym := NewClassSymbol(c.NameNode().GetText(),
			l.source,
			l.symbolSummaryForContext(c),
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
		return nil
	})

	walkSymbols(ctx, func(c *parser.PrototypeDefContext) error {
		psym := NewPrototypeOrInstanceSymbol(
			c.NameNode().GetText(),
			c.ParentReference().GetText(),
			l.source,
			"",
			symbolDefinitionForRuleContext(c.NameNode()),
			symbolDefinitionForRuleContext(c.StatementBlock()),
			false)
		l.Prototypes[strings.ToUpper(psym.Name())] = psym
		return nil
	})

	walkSymbols(ctx, func(c *parser.InstanceDefContext) error {
		psym := NewPrototypeOrInstanceSymbol(
			c.NameNode().GetText(),
			c.ParentReference().GetText(),
			l.source,
			"",
			symbolDefinitionForRuleContext(c.NameNode()),
			symbolDefinitionForRuleContext(c.StatementBlock()),
			true)
		l.Instances[strings.ToUpper(psym.Name())] = psym
		return nil
	})
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

	walkSymbols(ctx.ParameterList(), func(pdef *parser.ParameterDeclContext) error {
		params = append(params,
			NewVariableSymbol(pdef.NameNode().GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(pdef.NameNode())))
		return nil
	})

	fnc := NewFunctionSymbol(ctx.NameNode().GetText(),
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	l.Functions[strings.ToUpper(fnc.Name())] = fnc
}
