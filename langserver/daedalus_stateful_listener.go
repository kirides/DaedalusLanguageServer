package langserver

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// DaedalusStatefulListener ...
type DaedalusStatefulListener struct {
	parser.BaseDaedalusListener
	knownSymbols SymbolProvider
	Globals      struct {
		Instances  map[string]symbol.ProtoTypeOrInstance
		Variables  map[string]symbol.Symbol
		Constants  map[string]symbol.Symbol
		Functions  map[string]symbol.Function
		Classes    map[string]symbol.Class
		Prototypes map[string]symbol.ProtoTypeOrInstance
	}
	Namespaces     map[string]symbol.Namespace
	summaryBuilder *bytes.Buffer
	source         string

	curentNamespace *symbol.Namespace

	instances  []symbol.ProtoTypeOrInstance
	variables  []symbol.Symbol
	constants  []symbol.Symbol
	functions  []symbol.Function
	classes    []symbol.Class
	prototypes []symbol.ProtoTypeOrInstance
}

type SymbolProvider interface {
	WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error
	LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool)

	GetGlobalSymbols(types SymbolType) []symbol.Symbol
}

// NewDaedalusStatefulListener ...
func NewDaedalusStatefulListener(source string, knownSymbols SymbolProvider) *DaedalusStatefulListener {
	return &DaedalusStatefulListener{
		source:     source,
		Namespaces: make(map[string]symbol.Namespace),
		Globals: struct {
			Instances  map[string]symbol.ProtoTypeOrInstance
			Variables  map[string]symbol.Symbol
			Constants  map[string]symbol.Symbol
			Functions  map[string]symbol.Function
			Classes    map[string]symbol.Class
			Prototypes map[string]symbol.ProtoTypeOrInstance
		}{
			Variables:  make(map[string]symbol.Symbol, 1),
			Constants:  make(map[string]symbol.Symbol, 1),
			Functions:  make(map[string]symbol.Function, 1),
			Classes:    make(map[string]symbol.Class, 1),
			Prototypes: make(map[string]symbol.ProtoTypeOrInstance, 1),
			Instances:  make(map[string]symbol.ProtoTypeOrInstance, 1),
		},
		summaryBuilder: &bytes.Buffer{},
		knownSymbols:   knownSymbols,
	}
}

func (l *DaedalusStatefulListener) symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) symbol.Definition {
	return symbol.NewDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

type hiddenTokens interface {
	GetHiddenTokensToLeft(tokenIndex int, channel int) []antlr.Token
}

func (l *DaedalusStatefulListener) getHiddenTokens(ctx antlr.ParserRuleContext) (hiddenTokens, bool) {
	parserGetter, ok := ctx.(interface{ GetParser() antlr.Parser })
	if !ok {
		return nil, false
	}
	tokenStream := parserGetter.GetParser().GetTokenStream()
	hidden, ok := tokenStream.(hiddenTokens)
	return hidden, ok
}

func (l *DaedalusStatefulListener) symbolSummaryForContext(ctx antlr.ParserRuleContext) string {
	tokens, ok := l.getHiddenTokens(ctx)
	if !ok {
		return ""
	}
	hidden := tokens.GetHiddenTokensToLeft(ctx.GetStart().GetTokenIndex(), parser.COMMENTS)

	l.summaryBuilder.Reset()
	for _, h := range hidden {
		txt := h.GetText()
		if !strings.HasPrefix(txt, "///") {
			continue
		}
		if l.summaryBuilder.Len() > 0 {
			l.summaryBuilder.WriteRune('\n')
		}
		sumReplaced := strings.TrimLeft(txt, "/ \t")
		sumReplaced = strings.ReplaceAll(sumReplaced, "/// ", "  \n")
		l.summaryBuilder.WriteString(sumReplaced)
	}

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
					l.symbolDefinitionForRuleContext(val.NameNode()),
				))
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ival := range innerVal.AllVarValueDecl() {
				val := ival.(*parser.VarValueDeclContext)
				result = append(result,
					symbol.NewVariable(val.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						summary, // documentation
						l.symbolDefinitionForRuleContext(val.NameNode()),
					))
			}
		} else if innerVal, ok := ch.(*parser.VarArrayDeclContext); ok {
			result = append(result,
				symbol.NewArrayVariable(innerVal.NameNode().GetText(),
					v.TypeReference().GetText(),
					innerVal.ArraySize().GetText(),
					l.source,
					summary, // documentation
					l.symbolDefinitionForRuleContext(innerVal.NameNode()),
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
				l.symbolDefinitionForRuleContext(cv.NameNode()),
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
				l.symbolDefinitionForRuleContext(innerVal.NameNode()),
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
			l.constants = append(l.constants, l.constsFromContext(c)...)
		} else if v, ok := ch.(*parser.VarDeclContext); ok {
			l.variables = append(l.variables, l.variablesFromContext(v)...)
		} else if c, ok := ch.(*parser.InstanceDeclContext); ok {
			walkSymbols(c, func(name parser.INameNodeContext) error {
				psym := symbol.NewPrototypeOrInstance(
					name.GetText(),
					c.ParentReference().GetText(),
					l.source,
					"",
					l.symbolDefinitionForRuleContext(name),
					l.symbolDefinitionForRuleContext(c.ParentReference()),
					true)
				l.instances = append(l.instances, psym)
				return nil
			})
		}
	}
}

func walkSymbols[T antlr.ParserRuleContext](rule antlr.RuleNode, walkFn func(symbol T) error) {
	if rule == nil {
		return
	}
	for _, v := range rule.GetChildren() {
		if s, ok := v.(T); ok {
			if err := walkFn(s); err != nil {
				return
			}
		}
	}
}

func (l *DaedalusStatefulListener) getAssignedFields(ctx parser.IStatementBlockContext) []symbol.Constant {
	cFields := []symbol.Constant{}

	walkSymbols(ctx, func(v *parser.StatementContext) error {
		if v.Assignment() != nil {
			v := v.Assignment().(*parser.AssignmentContext)
			ref := v.Reference().(*parser.ReferenceContext)
			txt := ref.GetText()

			expr := v.ExpressionBlock().GetText()

			cFields = append(cFields,
				symbol.NewConstant(txt,
					"_",
					l.source,
					"",
					l.symbolDefinitionForRuleContext(ref),
					expr),
			)
		}
		return nil
	})
	return cFields
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
							l.symbolDefinitionForRuleContext(vv.NameNode())),
					)
				} else if vv, ok := ch.(*parser.VarArrayDeclContext); ok {
					cFields = append(cFields,
						symbol.NewArrayVariable(vv.NameNode().GetText(),
							v.TypeReference().GetText(),
							vv.ArraySize().GetText(),
							l.source,
							"",
							l.symbolDefinitionForRuleContext(vv.NameNode())),
					)
				}
			}
			return nil
		})

		csym := symbol.NewClass(c.NameNode().GetText(),
			l.source,
			l.symbolSummaryForContext(c),
			l.symbolDefinitionForRuleContext(c.NameNode()),
			symbol.NewDefinition(
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetColumn(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetColumn(),
			),
			cFields)
		l.classes = append(l.classes, csym)
		return nil
	})

	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.PrototypeDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(c.NameNode()),
				l.symbolDefinitionForRuleContext(c.StatementBlock()),
				false)
			psym.Fields = l.getAssignedFields(c.StatementBlock())
			l.prototypes = append(l.prototypes, psym)
		} else if c, ok := ch.(*parser.InstanceDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(c.NameNode()),
				l.symbolDefinitionForRuleContext(c.StatementBlock()),
				true)
			psym.Fields = l.getAssignedFields(c.StatementBlock())
			l.instances = append(l.instances, psym)
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
	if ctx.NameNode() == nil {
		return
	}

	statements := ctx.StatementBlock().(*parser.StatementBlockContext)

	params := []symbol.Variable{}
	locals := l.findVarsConstsInStatements(statements)

	walkSymbols(ctx.ParameterList(), func(pdef *parser.ParameterDeclContext) error {
		if pdef.NameNode() == nil {
			return nil
		}
		params = append(params,
			symbol.NewVariable(pdef.NameNode().GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(pdef.NameNode())))
		return nil
	})

	fnc := symbol.NewFunction(ctx.NameNode().GetText(),
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		l.symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		l.symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	l.functions = append(l.functions, fnc)
}

/* zParserExtender */

// EnterNamespaceDef implements parser.DaedalusListener
func (l *DaedalusStatefulListener) EnterNamespaceDef(ctx *parser.NamespaceDefContext) {
	l.copySymbolsToCurrentScope()

	if ctx.NameNode() == nil {
		return
	}

	parent := l.curentNamespace

	bodyDef := symbol.NewDefinition(
		ctx.NameNode().GetStop().GetLine(), ctx.GetStop().GetColumn(),
		ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())

	ns := symbol.NewNamespace(ctx.NameNode().GetText(),
		parent,
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		l.symbolDefinitionForRuleContext(ctx.NameNode()),
		bodyDef,
	)
	l.Namespaces[strings.ToUpper(ns.FullName())] = ns
	l.curentNamespace = &ns
}

func (l *DaedalusStatefulListener) copySymbolsToCurrentScope() {
	if l.curentNamespace != nil {
		for _, v := range l.constants {
			l.curentNamespace.Constants[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.variables {
			l.curentNamespace.Variables[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.instances {
			l.curentNamespace.Instances[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.prototypes {
			l.curentNamespace.Prototypes[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.classes {
			l.curentNamespace.Classes[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.functions {
			l.curentNamespace.Functions[strings.ToUpper(v.Name())] = v
		}
	} else {
		// global
		for _, v := range l.constants {
			l.Globals.Constants[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.variables {
			l.Globals.Variables[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.instances {
			l.Globals.Instances[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.prototypes {
			l.Globals.Prototypes[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.classes {
			l.Globals.Classes[strings.ToUpper(v.Name())] = v
		}
		for _, v := range l.functions {
			l.Globals.Functions[strings.ToUpper(v.Name())] = v
		}
	}

	l.constants = l.constants[:0]
	l.variables = l.variables[:0]
	l.instances = l.instances[:0]
	l.prototypes = l.prototypes[:0]
	l.classes = l.classes[:0]
	l.functions = l.functions[:0]
}

// ExitNamespaceDef implements parser.DaedalusListener
func (l *DaedalusStatefulListener) ExitNamespaceDef(ctx *parser.NamespaceDefContext) {
	l.copySymbolsToCurrentScope()
	l.curentNamespace = l.curentNamespace.Parent
}

func (l *DaedalusStatefulListener) EnterZParserExtenderMetaBlock(ctx *parser.ZParserExtenderMetaBlockContext) {
}

// ExitNamespaceDef implements parser.DaedalusListener
func (l *DaedalusStatefulListener) ExitDaedalusFile(ctx *parser.DaedalusFileContext) {
	l.copySymbolsToCurrentScope()
}
