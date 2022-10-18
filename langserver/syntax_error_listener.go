package langserver

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SyntaxErrorListener ...
type SyntaxErrorListener struct {
	antlr.DefaultErrorListener
	SyntaxErrors []SyntaxError
}

// SyntaxError ...
func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if se, ok := e.(*listenerSyntaxError); ok {
		l.SyntaxErrors = append(l.SyntaxErrors, NewSyntaxError(line, column, se.Code))
	} else {
		l.SyntaxErrors = append(l.SyntaxErrors, NewSyntaxError(line, column, NewSyntaxErrorCode("D0000", msg, SeverityWarning)))
	}
}

// NoOpErrorListener ...
type NoOpErrorListener struct {
	antlr.DefaultErrorListener
}

// SyntaxError ...
func (l *NoOpErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
}
