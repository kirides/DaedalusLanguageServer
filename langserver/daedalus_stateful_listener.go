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
	summary := ""
	if p, ok := v.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	nVars := 0
	for _, ch := range v.GetChildren() {
		if _, ok := ch.(*parser.VarValueDeclContext); ok {
			nVars++
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ctx := range innerVal.GetChildren() {
				if _, ok := ctx.(parser.IVarValueDeclContext); ok {
					nVars++
				}
			}
		} else if _, ok := ch.(*parser.VarArrayDeclContext); ok {
			nVars++
		}
	}

	result := make([]symbol.Symbol, 0, nVars)

	for _, ch := range v.GetChildren() {
		if val, ok := ch.(*parser.VarValueDeclContext); ok {
			nameNode := val.NameNode()
			result = append(result,
				symbol.NewVariable(nameNode.GetText(),
					v.TypeReference().GetText(),
					l.source,
					summary, // documentation
					l.symbolDefinitionForRuleContext(nameNode),
				))
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ival := range innerVal.AllVarValueDecl() {
				val := ival.(*parser.VarValueDeclContext)
				nameNode := val.NameNode()
				result = append(result,
					symbol.NewVariable(nameNode.GetText(),
						v.TypeReference().GetText(),
						l.source,
						summary, // documentation
						l.symbolDefinitionForRuleContext(nameNode),
					))
			}
		} else if innerVal, ok := ch.(*parser.VarArrayDeclContext); ok {
			nameNode := innerVal.NameNode()
			result = append(result,
				symbol.NewArrayVariable(nameNode.GetText(),
					v.TypeReference().GetText(),
					innerVal.ArraySize().GetText(),
					l.source,
					summary, // documentation
					l.symbolDefinitionForRuleContext(nameNode),
				))
		}
	}
	return result
}

func (l *DaedalusStatefulListener) maxNOfConstValues(n int, c parser.IConstArrayAssignmentContext) string {
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
		result.WriteString(v.GetText())
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

func (l *DaedalusStatefulListener) arrayElementsFromContext(c parser.IConstArrayAssignmentContext) []symbol.ArrayElement {
	result := make([]symbol.ArrayElement, 0, 66)
	for _, v := range c.AllExpressionBlock() {
		result = append(result, symbol.NewArrayElement(
			v.GetText(),
			symbol.NewDefinition(v.GetStart().GetLine(),
				v.GetStart().GetColumn(),
				v.GetStop().GetLine(),
				v.GetStop().GetColumn(),
			)))
	}
	return result
}

func (l *DaedalusStatefulListener) constsFromContext(c *parser.ConstDefContext) []symbol.Symbol {
	summary := ""
	if p, ok := c.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	nDefs := 0
	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		nDefs++
		return nil
	})

	walkSymbols(c, func(innerVal *parser.ConstArrayDefContext) error {
		nDefs++
		return nil
	})

	result := make([]symbol.Symbol, 0, nDefs)
	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		nameNode := cv.NameNode()
		result = append(result,
			symbol.NewConstant(nameNode.GetText(),
				c.TypeReference().GetText(),
				l.source,
				summary, // documentation
				l.symbolDefinitionForRuleContext(nameNode),
				cv.ConstValueAssignment().ExpressionBlock().GetText(),
			))
		return nil
	})

	walkSymbols(c, func(innerVal *parser.ConstArrayDefContext) error {
		nameNode := innerVal.NameNode()
		constAssignment := innerVal.ConstArrayAssignment()
		result = append(result,
			symbol.NewConstantArray(nameNode.GetText(),
				c.TypeReference().GetText(),
				innerVal.ArraySize().GetText(),
				l.source,
				summary, // documentation
				l.symbolDefinitionForRuleContext(nameNode),
				l.maxNOfConstValues(3, constAssignment),
				l.arrayElementsFromContext(constAssignment),
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
	nFields := 0
	walkSymbols(ctx, func(v *parser.StatementContext) error {
		assignment := v.Assignment()
		if assignment != nil {
			nFields++
		}
		return nil
	})

	cFields := make([]symbol.Constant, 0, nFields)
	walkSymbols(ctx, func(v *parser.StatementContext) error {
		assignment := v.Assignment()
		if assignment != nil {
			v := assignment
			ref := v.Reference()
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
					nameNode := vv.NameNode()
					cFields = append(cFields,
						symbol.NewVariable(nameNode.GetText(),
							v.TypeReference().GetText(),
							l.source,
							"",
							l.symbolDefinitionForRuleContext(nameNode)),
					)
				} else if vv, ok := ch.(*parser.VarArrayDeclContext); ok {
					nameNode := vv.NameNode()
					cFields = append(cFields,
						symbol.NewArrayVariable(nameNode.GetText(),
							v.TypeReference().GetText(),
							vv.ArraySize().GetText(),
							l.source,
							"",
							l.symbolDefinitionForRuleContext(nameNode)),
					)
				}
			}
			return nil
		})
		classNameNode := c.NameNode()
		startNode := c.GetChild(2).(antlr.TerminalNode).GetSymbol()
		endNode := c.GetChild(c.GetChildCount() - 1).(antlr.TerminalNode).GetSymbol()
		csym := symbol.NewClass(classNameNode.GetText(),
			l.source,
			l.symbolSummaryForContext(c),
			l.symbolDefinitionForRuleContext(classNameNode),
			symbol.NewDefinition(
				startNode.GetLine(),
				startNode.GetColumn(),
				endNode.GetLine(),
				endNode.GetColumn(),
			),
			cFields)
		l.classes = append(l.classes, csym)
		return nil
	})

	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.PrototypeDefContext); ok {
			cNameNode := c.NameNode()
			psym := symbol.NewPrototypeOrInstance(
				cNameNode.GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(cNameNode),
				l.symbolDefinitionForRuleContext(c.StatementBlock()),
				false)
			psym.Fields = l.getAssignedFields(c.StatementBlock())
			l.prototypes = append(l.prototypes, psym)
		} else if c, ok := ch.(*parser.InstanceDefContext); ok {
			cNameNode := c.NameNode()
			psym := symbol.NewPrototypeOrInstance(
				cNameNode.GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(cNameNode),
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
	ctxNameNode := ctx.NameNode()
	if ctxNameNode == nil {
		return
	}

	statements := ctx.StatementBlock()

	params := []symbol.Variable{}
	locals := l.findVarsConstsInStatements(statements)

	walkSymbols(ctx.ParameterList(), func(pdef *parser.ParameterDeclContext) error {
		nameNode := pdef.NameNode()
		if nameNode == nil {
			return nil
		}
		params = append(params,
			symbol.NewVariable(nameNode.GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				l.symbolDefinitionForRuleContext(nameNode)))
		return nil
	})

	fnc := symbol.NewFunction(ctxNameNode.GetText(),
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		l.symbolDefinitionForRuleContext(ctxNameNode),
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
	ctxNameNode := ctx.NameNode()
	if ctxNameNode == nil {
		return
	}

	parent := l.curentNamespace

	bodyDef := symbol.NewDefinition(
		ctxNameNode.GetStop().GetLine(), ctx.GetStop().GetColumn(),
		ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())

	ns := symbol.NewNamespace(ctxNameNode.GetText(),
		parent,
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		l.symbolDefinitionForRuleContext(ctxNameNode),
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
