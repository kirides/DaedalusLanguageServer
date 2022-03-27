package langserver

import (
	"context"
	"encoding/json"

	"go.lsp.dev/jsonrpc2"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type textDocumentSyncHandler struct {
	bufferManager   *BufferManager
	parsedDocuments *parseResultsManager
	baseLspHandler
}

func lspSeverityFromSeverity(severity ErrorSeverity) lsp.DiagnosticSeverity {
	switch severity {
	case SeverityInfo:
		return lsp.DiagnosticSeverityInformation
	case SeverityWarning:
		return lsp.DiagnosticSeverityWarning
	}
	return lsp.DiagnosticSeverityError
}

func (h *textDocumentSyncHandler) updateBuffer(ctx context.Context, documentURI, content string) {
	chars := content
	h.bufferManager.UpdateBuffer(documentURI, chars)
	p, _ := h.parsedDocuments.Update(documentURI, content)

	diagnostics := []lsp.Diagnostic{}
	if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
		for _, se := range p.SyntaxErrors {
			diagnostics = append(diagnostics, se.Diagnostic())
		}
	}
	h.conn.Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
		URI:         lsp.DocumentURI(uri.File(documentURI)),
		Diagnostics: diagnostics,
	})

	h.LogDebug("Updated buffer for %q with %d chars\n", documentURI, len(chars))
}

// Deliver ...
func (h *textDocumentSyncHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, r jsonrpc2.Request) error {
	// Recover if something bad happens in the handlers...
	defer func() {
		err := recover()
		if err != nil {
			h.LogWarn("Recovered from panic at %s: %v\n", r.Method(), err)
		}
	}()
	switch r.Method() {
	case lsp.MethodTextDocumentDidOpen:
		var params lsp.DidOpenTextDocumentParams
		json.Unmarshal(r.Params(), &params)
		documentUri := h.uriToFilename(params.TextDocument.URI)
		if documentUri != "" {
			h.updateBuffer(ctx, documentUri, params.TextDocument.Text)
		}
	case lsp.MethodTextDocumentDidChange:
		var params lsp.DidChangeTextDocumentParams
		json.Unmarshal(r.Params(), &params)
		documentUri := h.uriToFilename(params.TextDocument.URI)
		if documentUri != "" && len(params.ContentChanges) > 0 {
			h.updateBuffer(ctx, documentUri, params.ContentChanges[0].Text)
		}
	case lsp.MethodTextDocumentDidSave:
		var params lsp.DidSaveTextDocumentParams
		json.Unmarshal(r.Params(), &params)
		documentUri := h.uriToFilename(params.TextDocument.URI)
		if documentUri != "" {
			h.updateBuffer(ctx, documentUri, params.Text)
		}

	default:
		return h.baseLspHandler.Handle(ctx, reply, r)
	}
	return nil
}
