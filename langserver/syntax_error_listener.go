package langserver

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SyntaxErrorListener ...
type SyntaxErrorListener struct {
	antlr.DefaultErrorListener
	SyntaxErrors []SyntaxError
}

// SyntaxError ...
func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.SyntaxErrors = append(l.SyntaxErrors, NewSyntaxError(line, column, NewSyntaxErrorCode("D0000", msg, SeverityWarning)))
}
