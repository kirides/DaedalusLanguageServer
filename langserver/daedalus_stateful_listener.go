package langserver

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// DaedalusStatefulListener ...
type DaedalusStatefulListener struct {
	parser.BaseDaedalusListener
	knownSymbols    SymbolProvider
	Instances       map[string]symbol.ProtoTypeOrInstance
	GlobalVariables map[string]symbol.Symbol
	GlobalConstants map[string]symbol.Symbol
	Functions       map[string]symbol.Function
	Classes         map[string]symbol.Class
	Prototypes      map[string]symbol.ProtoTypeOrInstance
	summaryBuilder  *bytes.Buffer
	source          string
}

type SymbolProvider interface {
	WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error
	LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool)
}

// NewDaedalusStatefulListener ...
func NewDaedalusStatefulListener(source string, knownSymbols SymbolProvider) *DaedalusStatefulListener {
	return &DaedalusStatefulListener{
		source:          source,
		GlobalVariables: map[string]symbol.Symbol{},
		GlobalConstants: map[string]symbol.Symbol{},
		Functions:       map[string]symbol.Function{},
		Classes:         map[string]symbol.Class{},
		Prototypes:      map[string]symbol.ProtoTypeOrInstance{},
		Instances:       map[string]symbol.ProtoTypeOrInstance{},
		summaryBuilder:  &bytes.Buffer{},
		knownSymbols:    knownSymbols,
	}
}

func symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) symbol.Definition {
	return symbol.NewDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
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

func (l *DaedalusStatefulListener) variablesFromContext(v *parser.VarDeclContext) []symbol.Symbol {
	result := []symbol.Symbol{}
	summary := ""
	if p, ok := v.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	for _, ch := range v.GetChildren() {
		if val, ok := ch.(*parser.VarValueDeclContext); ok {
			result = append(result,
				symbol.NewVariable(val.NameNode().GetText(),
					v.TypeReference().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(val.NameNode()),
				))
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ival := range innerVal.AllVarValueDecl() {
				val := ival.(*parser.VarValueDeclContext)
				result = append(result,
					symbol.NewVariable(val.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						summary, // documentation
						symbolDefinitionForRuleContext(val.NameNode()),
					))
			}
		} else if innerVal, ok := ch.(*parser.VarArrayDeclContext); ok {
			result = append(result,
				symbol.NewArrayVariable(innerVal.NameNode().GetText(),
					v.TypeReference().GetText(),
					innerVal.ArraySize().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(innerVal.NameNode()),
				))
		}
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

func (l *DaedalusStatefulListener) constsFromContext(c *parser.ConstDefContext) []symbol.Symbol {
	result := []symbol.Symbol{}
	summary := ""
	if p, ok := c.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		result = append(result,
			symbol.NewConstant(cv.NameNode().GetText(),
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
			symbol.NewConstantArray(innerVal.NameNode().GetText(),
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
	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.ConstDefContext); ok {
			for _, s := range l.constsFromContext(c) {
				l.GlobalConstants[strings.ToUpper(s.Name())] = s
			}
		} else if v, ok := ch.(*parser.VarDeclContext); ok {
			for _, s := range l.variablesFromContext(v) {
				l.GlobalVariables[strings.ToUpper(s.Name())] = s
			}
		} else if c, ok := ch.(*parser.InstanceDeclContext); ok {
			walkSymbols(c, func(name parser.INameNodeContext) error {
				psym := symbol.NewPrototypeOrInstance(
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
		}
	}
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
		cFields := []symbol.Symbol{}
		walkSymbols(c, func(v *parser.VarDeclContext) error {
			for _, ch := range v.GetChildren() {
				if vv, ok := ch.(*parser.VarValueDeclContext); ok {
					cFields = append(cFields,
						symbol.NewVariable(vv.NameNode().GetText(),
							v.TypeReference().GetText(),
							l.source,
							"",
							symbolDefinitionForRuleContext(vv.NameNode())),
					)
				} else if vv, ok := ch.(*parser.VarArrayDeclContext); ok {
					cFields = append(cFields,
						symbol.NewArrayVariable(vv.NameNode().GetText(),
							v.TypeReference().GetText(),
							vv.ArraySize().GetText(),
							l.source,
							"",
							symbolDefinitionForRuleContext(vv.NameNode())),
					)
				}
			}
			return nil
		})

		csym := symbol.NewClass(c.NameNode().GetText(),
			l.source,
			l.symbolSummaryForContext(c),
			symbolDefinitionForRuleContext(c.NameNode()),
			symbol.NewDefinition(
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetColumn(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetColumn(),
			),
			cFields)
		l.Classes[strings.ToUpper(csym.Name())] = csym
		return nil
	})

	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.PrototypeDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(c.NameNode()),
				symbolDefinitionForRuleContext(c.StatementBlock()),
				false)
			l.Prototypes[strings.ToUpper(psym.Name())] = psym
		} else if c, ok := ch.(*parser.InstanceDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(c.NameNode()),
				symbolDefinitionForRuleContext(c.StatementBlock()),
				true)
			l.Instances[strings.ToUpper(psym.Name())] = psym
		}
	}
}

func (l *DaedalusStatefulListener) findVarsConstsInStatements(root antlr.Tree) []symbol.Symbol {
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

// EnterFunctionDef ...
func (l *DaedalusStatefulListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
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
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	l.Functions[strings.ToUpper(fnc.Name())] = fnc
}
