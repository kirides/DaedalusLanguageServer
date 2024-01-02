package langserver

import (
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

	"github.com/antlr4-go/antlr/v4"
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

	GlobalIdentifiers []Identifier
	source            string

	Bbox symbol.Definition
}

// NewDaedalusIdentifierListener ...
func NewDaedalusIdentifierListener(source string) *DaedalusIdentifierListener {
	return &DaedalusIdentifierListener{
		source:            source,
		GlobalIdentifiers: []Identifier{},
		Bbox:              symbol.Definition{},
	}
}

func (l *DaedalusIdentifierListener) symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) symbol.Definition {
	return symbol.NewDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn()+len(ctx.GetText()))
}

func (l *DaedalusIdentifierListener) filterRange(ctx antlr.ParserRuleContext) bool {
	if l.Bbox != (symbol.Definition{}) {
		if !l.Bbox.InBBox(symbol.DefinitionIndex{Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn()}) {
			return true
		}
	}
	return false
}

func (l *DaedalusIdentifierListener) onIdentifier(ctx antlr.ParserRuleContext) {
	if l.filterRange(ctx) {
		return
	}
	txt := ctx.GetText()

	if strings.Trim(txt, "0123456789") == "" {
		// just numbers
		return
	}
	// only plain tokens for now
	if strings.EqualFold(txt, "true") ||
		// strings.EqualFold(txt, "null") || // people actually use "null" as variable names
		strings.EqualFold(txt, "false") ||

		strings.EqualFold(txt, "hero") ||
		strings.EqualFold(txt, "self") ||
		strings.EqualFold(txt, "other") ||
		strings.EqualFold(txt, "victim") ||

		strings.EqualFold(txt, "end") ||
		strings.EqualFold(txt, "repeat") ||
		strings.EqualFold(txt, "continue") ||
		strings.EqualFold(txt, "break") ||
		strings.EqualFold(txt, "while") ||

		strings.EqualFold(txt, "item") {
		return
	}

	l.GlobalIdentifiers = append(l.GlobalIdentifiers, Identifier{
		nameValue:  txt,
		definition: l.symbolDefinitionForRuleContext(ctx),
	})
}

func (l *DaedalusIdentifierListener) EnterAssignment(ctx *parser.AssignmentContext) {
	refAtom := ctx.Reference()
	l.onIdentifier(refAtom)
}

func (l *DaedalusIdentifierListener) EnterArrayIndex(ctx *parser.ArrayIndexContext) {
	refAtom := ctx.ReferenceAtom()
	if refAtom == nil {
		return
	}
	nameNode := refAtom.NameNode()
	l.onIdentifier(nameNode)
}

func (l *DaedalusIdentifierListener) EnterReferenceValue(ctx *parser.ReferenceValueContext) {
	refAtom := ctx.Reference()
	l.onIdentifier(refAtom)
}

func (l *DaedalusIdentifierListener) EnterNullLiteralValue(ctx *parser.NullLiteralValueContext) {
	l.onIdentifier(ctx)
}
