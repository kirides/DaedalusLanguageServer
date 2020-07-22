package langserver

import (
	"context"
	"encoding/json"

	"go.lsp.dev/jsonrpc2"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type textDocumentSyncHandler struct {
	baseLspHandler

	bufferManager   *BufferManager
	parsedDocuments *parseResultsManager
}

func lspSeverityFromSeverity(severity ErrorSeverity) lsp.DiagnosticSeverity {
	switch severity {
	case SeverityInfo:
		return lsp.SeverityInformation
	case SeverityWarning:
		return lsp.SeverityWarning
	}
	return lsp.SeverityError
}

func (h *textDocumentSyncHandler) updateBuffer(ctx context.Context, r *jsonrpc2.Request, documentURI, content string) {
	chars := content
	h.bufferManager.UpdateBuffer(documentURI, chars)
	p, _ := h.parsedDocuments.Update(documentURI, content)

	diagnostics := []lsp.Diagnostic{}
	if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
		for _, se := range p.SyntaxErrors {
			diagnostics = append(diagnostics, lsp.Diagnostic{
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
			})
		}
	}
	r.Conn().Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
		URI:         lsp.DocumentURI(uri.File(documentURI)),
		Diagnostics: diagnostics,
	})

	h.LogDebug("Updated buffer for %q with %d chars\n", documentURI, len(chars))
}

// Deliver ...
func (h *textDocumentSyncHandler) Deliver(ctx context.Context, r *jsonrpc2.Request, delivered bool) bool {
	// Recover if something bad happens in the handlers...
	defer func() {
		err := recover()
		if err != nil {
			h.LogWarn("Recovered from panic at %s: %v\n", r.Method, err)
		}
	}()
	switch r.Method {
	case lsp.MethodTextDocumentDidOpen:
		var params lsp.DidOpenTextDocumentParams
		json.Unmarshal(*r.Params, &params)
		h.updateBuffer(ctx, r, h.uriToFilename(params.TextDocument.URI), params.TextDocument.Text)
	case lsp.MethodTextDocumentDidChange:
		var params lsp.DidChangeTextDocumentParams
		json.Unmarshal(*r.Params, &params)
		h.updateBuffer(ctx, r, h.uriToFilename(params.TextDocument.URI), params.ContentChanges[0].Text)

	case lsp.MethodTextDocumentDidSave:
		var params lsp.DidSaveTextDocumentParams
		json.Unmarshal(*r.Params, &params)
		h.updateBuffer(ctx, r, h.uriToFilename(params.TextDocument.URI), params.Text)

	default:
		return h.baseLspHandler.Deliver(ctx, r, delivered)
	}
	return true
}
