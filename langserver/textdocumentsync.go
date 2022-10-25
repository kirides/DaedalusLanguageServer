package langserver

import (
	"context"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

type textDocumentSync struct {
	bufferManager   *BufferManager
	parsedDocuments *parseResultsManager
	baseLspHandler
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

func (h *textDocumentSync) updateBuffer(ctx context.Context, documentURI, content string) {
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

	h.LogDebug("Updated buffer for %q with %d chars", documentURI, len(chars))
}

func (h *textDocumentSync) handleTextDocumentDidClose(req dls.RpcContext, data lsp.DidCloseTextDocumentParams) error {
	documentUri := uriToFilename(data.TextDocument.URI)
	if documentUri != "" {
		h.bufferManager.DeleteBuffer(documentUri)
	}
	return nil
}
func (h *textDocumentSync) handleTextDocumentDidOpen(req dls.RpcContext, data lsp.DidOpenTextDocumentParams) error {
	documentUri := uriToFilename(data.TextDocument.URI)
	if documentUri != "" {
		h.updateBuffer(req.Context(), documentUri, data.TextDocument.Text)
	}
	return nil
}
func (h *textDocumentSync) handleTextDocumentDidChange(req dls.RpcContext, data lsp.DidChangeTextDocumentParams) error {
	documentUri := uriToFilename(data.TextDocument.URI)
	if documentUri != "" && len(data.ContentChanges) > 0 {
		h.updateBuffer(req.Context(), documentUri, data.ContentChanges[0].Text)
	}
	return nil
}

func (h *textDocumentSync) handleTextDocumentDidSave(req dls.RpcContext, data lsp.DidSaveTextDocumentParams) error {
	documentUri := uriToFilename(data.TextDocument.URI)
	if documentUri != "" {
		text := ""
		if data.Text != nil {
			text = *data.Text
		}
		h.updateBuffer(req.Context(), documentUri, text)
	}
	return nil
}
