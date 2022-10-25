package langserver

import (
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

// ErrorSeverity ...
type ErrorSeverity int

const (
	// SeverityInfo ...
	SeverityInfo ErrorSeverity = 0
	// SeverityWarning ...
	SeverityWarning ErrorSeverity = 1
	// SeverityError ...
	SeverityError ErrorSeverity = 2
)

// SyntaxError ...
type SyntaxError struct {
	ErrorCode SyntaxErrorCode
	Line      int
	Column    int
}

// SyntaxErrorCode ...
type SyntaxErrorCode struct {
	Code        string
	Description string
	Severity    ErrorSeverity
}

var (
	// D0001NoIdentifierWithStartingDigits ...
	D0001NoIdentifierWithStartingDigits = NewSyntaxErrorCode("D0001", "Do not start identifiers with digits", SeverityWarning)
	// D0002SplitMultipleVarDecl ...
	D0002SplitMultipleVarDecl = NewSyntaxErrorCode("D0002", "Split multiple 'var TYPE ..., var TYPE ...' into separate statements.", SeverityWarning)
	// D0003MissingSemicolonMightCauseIssues ...
	D0003MissingSemicolonMightCauseIssues = NewSyntaxErrorCode("D0003", "Missing ';' might cause issues.", SeverityError)
	// D0004NotEnoughArgumentsSpecified ...
	D0004NotEnoughArgumentsSpecified = NewSyntaxErrorCode("D0004", "Not enough arguments specified", SeverityError)
	// D0005TooManyArgumentsSpecified
	D0005TooManyArgumentsSpecified = NewSyntaxErrorCode("D0005", "Too many arguments specified", SeverityError)
)

// NewSyntaxError ...
func NewSyntaxError(line, col int, code SyntaxErrorCode) SyntaxError {
	return SyntaxError{
		Line:      line,
		Column:    col,
		ErrorCode: code,
	}
}

// NewSyntaxErrorCode ...
func NewSyntaxErrorCode(code, desc string, severity ErrorSeverity) SyntaxErrorCode {
	return SyntaxErrorCode{
		Code:        code,
		Description: desc,
		Severity:    severity,
	}
}

// Diagnostic returns a diagnostic based on this sytax error
func (se SyntaxError) Diagnostic() lsp.Diagnostic {
	return lsp.Diagnostic{
		Code:     se.ErrorCode.Code,
		Message:  se.ErrorCode.Description,
		Source:   "vscode-daedalus",
		Severity: lspSeverityFromSeverity(se.ErrorCode.Severity),
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      uint32(se.Line - 1),
				Character: uint32(se.Column),
			},
			End: lsp.Position{
				Line:      uint32(se.Line - 1),
				Character: uint32(se.Column),
			},
		},
	}
}
