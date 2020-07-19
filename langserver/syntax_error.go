package langserver

import (
	lsp "go.lsp.dev/protocol"
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
	Line      int
	Column    int
	ErrorCode SyntaxErrorCode
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
	D0003MissingSemicolonMightCauseIssues = NewSyntaxErrorCode("D0003", "Missing ';' might cause issues.", SeverityWarning)
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
				Line:      float64(se.Line - 1),
				Character: float64(se.Column),
			},
			End: lsp.Position{
				Line:      float64(se.Line - 1),
				Character: float64(se.Column),
			},
		},
	}
}
