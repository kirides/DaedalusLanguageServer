package langserver

import (
	"fmt"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"golang.org/x/exp/slices"

	"github.com/antlr4-go/antlr/v4"
)

// pendingCall is a function call whose argument count could not be checked yet,
// because the functions of all other scripts of a workspace are not known while parsing.
type pendingCall struct {
	name      string
	args      int32
	line, col int32
}

// DaedalusValidatingListener ...
type DaedalusValidatingListener struct {
	parser.BaseDaedalusListener
	knownSymbols SymbolProvider
	source       string

	// deferCalls collects function calls in pendingCalls instead of validating them directly.
	deferCalls   bool
	pendingCalls []pendingCall
}

// NewDeferringValidatingListener creates a listener that validates everything that does not
// depend on other scripts, and collects the function calls to be checked later.
func NewDeferringValidatingListener(source string) *DaedalusValidatingListener {
	return &DaedalusValidatingListener{source: source, deferCalls: true}
}

// NewDaedalusValidatingListener ...
func NewDaedalusValidatingListener(source string, knownSymbols SymbolProvider) *DaedalusValidatingListener {
	return &DaedalusValidatingListener{
		source:       source,
		knownSymbols: knownSymbols,
	}
}

type listenerSyntaxError struct {
	*antlr.BaseRecognitionException
	Code SyntaxErrorCode
}

func (l *DaedalusValidatingListener) report(parser antlr.Parser, ctx antlr.RuleContext, symbol antlr.Token, code SyntaxErrorCode) {
	ex := antlr.NewBaseRecognitionException(code.Description, parser, parser.GetInputStream(), ctx)
	parser.NotifyErrorListeners(ex.GetMessage(), symbol, &listenerSyntaxError{BaseRecognitionException: ex, Code: code})
}

// EnterStatementBlock ...
func (l *DaedalusValidatingListener) EnterStatementBlock(ctx *parser.StatementBlockContext) {
	for i, stmt := range ctx.GetChildren() {
		if ifStmt, ok := stmt.(*parser.IfBlockStatementContext); ok {
			if ctx.GetChildCount()-1 == i {
				l.report(ctx.GetParser(), ctx, ifStmt.GetStop(), D0003MissingSemicolonMightCauseIssues)
				continue
			}
			child := ctx.GetChild(i + 1)
			if termNode, ok := child.(*antlr.TerminalNodeImpl); !(ok && termNode.GetText() == ";") {
				l.report(ctx.GetParser(), ctx, ifStmt.GetStop(), D0003MissingSemicolonMightCauseIssues)
			}
		}
	}
}

func (l *DaedalusValidatingListener) checkIdentifierStart(p antlr.Parser, ctx antlr.RuleContext, name antlr.Token) {
	if name == nil || name.GetTokenType() != parser.DaedalusParserIdentifier {
		return
	}
	id := name.GetText()
	if len(id) > 0 && (id[0] >= 0x30 && id[0] <= 0x39) {
		l.report(p, ctx, name, D0001NoIdentifierWithStartingDigits)
	}
}

// EnterNameNode ...
func (l *DaedalusValidatingListener) EnterNameNode(ctx *parser.NameNodeContext) {
	if id := ctx.Identifier(); id != nil {
		l.checkIdentifierStart(ctx.GetParser(), ctx, id.GetSymbol())
	}
}

// EnterReferenceExpression checks the (inlined) names of references.
func (l *DaedalusValidatingListener) EnterReferenceExpression(ctx *parser.ReferenceExpressionContext) {
	l.checkIdentifierStart(ctx.GetParser(), ctx, ctx.GetName())
	l.checkIdentifierStart(ctx.GetParser(), ctx, ctx.GetMember())
}

func (l *DaedalusValidatingListener) lookupSymbol(name string, symbol SymbolType) symbol.Symbol {
	foundSymbol, _ := l.knownSymbols.LookupGlobalSymbol(name, symbol)
	return foundSymbol
}

// EnterFuncCallExpression ...
func (l *DaedalusValidatingListener) EnterFuncCallExpression(ctx *parser.FuncCallExpressionContext) {
	name := ctx.GetName()
	l.checkIdentifierStart(ctx.GetParser(), ctx, name)

	if l.deferCalls {
		l.pendingCalls = append(l.pendingCalls, pendingCall{
			name: name.GetText(), args: int32(len(ctx.AllExpression())), line: int32(name.GetLine()), col: int32(name.GetColumn()),
		})
		return
	}

	sym := l.lookupSymbol(strings.ToUpper(name.GetText()), SymbolFunction)
	if sym == nil {
		return
	}

	numArgs := len(ctx.AllExpression())
	if numArgs < len(sym.(symbol.Function).Parameters) {
		l.report(ctx.GetParser(), ctx, name, D0004NotEnoughArgumentsSpecified)
	} else if numArgs > len(sym.(symbol.Function).Parameters) {
		l.report(ctx.GetParser(), ctx, name, D0005TooManyArgumentsSpecified)
	}
}

// zParserExtender

var metaFields = map[string][]string{
	strings.ToUpper("Parser"):      {"GAME", "SFX", "PFX", "VFX", "CAMERA", "MENU", "MUSIC"},
	strings.ToUpper("MergeMode"):   {"0", "1", "TRUE", "FALSE"},
	strings.ToUpper("Engine"):      {"G1", "G1A", "G2", "G2A"},
	strings.ToUpper("NativeWhile"): {"0", "1", "TRUE", "FALSE"},
	strings.ToUpper("Namespace"):   nil,
	strings.ToUpper("Using"):       nil,
	strings.ToUpper("Mod"):         nil,
	strings.ToUpper("After"):       nil,
}

// EnterZParserExtenderMeta implements parser.DaedalusListener
func (l *DaedalusValidatingListener) EnterZParserExtenderMeta(ctx *parser.ZParserExtenderMetaContext) {
	nameNode := ctx.NameNode()
	if nameNode == nil || ctx.MetaValue() == nil {
		return
	}
	key := strings.ToUpper(nameNode.GetText())
	if values, ok := metaFields[key]; ok {
		// nil means there are no checks
		if values == nil {
			return
		}
		switch key {
		// these are comma separated
		case "ENGINE", "NAMESPACE", "USING", "AFTER":
			items := strings.Split(ctx.MetaValue().GetText(), ",")
			for _, v := range items {
				if !slices.Contains(values, strings.ToUpper(strings.TrimSpace(v))) {
					l.report(ctx.GetParser(),
						ctx.MetaValue(),
						ctx.MetaValue().GetStart(),
						NewSyntaxErrorCode("D0000", fmt.Sprintf("Unrecognized values. Known values are: %q", strings.Join(values, ", ")), SeverityError))
					return
				}
			}
		default:
			if !slices.Contains(values, strings.ToUpper(ctx.MetaValue().GetText())) {
				l.report(ctx.GetParser(),
					ctx.MetaValue(),
					ctx.MetaValue().GetStart(),
					NewSyntaxErrorCode("D0000", fmt.Sprintf("Unrecognized value. Known values are: %q", strings.Join(values, ", ")), SeverityError))
			}
		}
	} else {
		l.report(ctx.GetParser(), nameNode, nameNode.GetStart(), D0006UnrecognizedMetaField)
	}
}
