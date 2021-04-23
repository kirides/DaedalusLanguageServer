package langserver

import (
	"langsrv/langserver/parser"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// DaedalusValidatingListener ...
type DaedalusValidatingListener struct {
	parser.BaseDaedalusListener
	knownSymbols symbolWalker
	source       string
}

// NewDaedalusValidatingListener ...
func NewDaedalusValidatingListener(source string, knownSymbols symbolWalker) *DaedalusValidatingListener {
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

// EnterAnyIdentifier ...
func (l *DaedalusValidatingListener) EnterAnyIdentifier(ctx *parser.AnyIdentifierContext) {
	if ctx.Identifier() != nil {
		id := ctx.Identifier().GetText()

		if len(id) > 0 && (id[0] >= 0x30 && id[0] <= 0x39) {
			l.report(ctx.GetParser(), ctx, ctx.Identifier().GetSymbol(), D0001NoIdentifierWithStartingDigits)
		}
	}
}

func (l *DaedalusValidatingListener) lookupSymbol(name string, symbol SymbolType) Symbol {
	foundSymbol, _ := l.knownSymbols.LookupGlobalSymbol(name, symbol)
	return foundSymbol
}

// EnterFuncCall ...
func (l *DaedalusValidatingListener) EnterFuncCall(ctx *parser.FuncCallContext) {
	funcName := ctx.NameNode().GetText()
	funcName = strings.ToUpper(funcName)
	sym := l.lookupSymbol(strings.ToUpper(funcName), SymbolFunction)

	if sym == nil {
		return
	}

	if len(ctx.AllFuncArgExpression()) < len(sym.(FunctionSymbol).Parameters) {
		l.report(ctx.GetParser(), ctx, ctx.NameNode().GetStop(), D0004NotEnoughArgumentsSpecified)
	}
}
