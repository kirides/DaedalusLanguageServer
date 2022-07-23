package langserver

import (
	"context"
	"fmt"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/javadoc"
	lsp "go.lsp.dev/protocol"
)

func (h *LspHandler) getSignatureInfo(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.SignatureHelp, error) {
	fnCtx, err := getFunctionCallContext(h, params.TextDocument.URI, params.Position)
	if err != nil {
		return lsp.SignatureHelp{}, err
	}

	fn := fnCtx.Function

	var fnParams []lsp.ParameterInformation
	for _, p := range fn.Parameters {
		doc := javadoc.FindParam(fn.Documentation(), p.Name())
		var mdDoc interface{}
		if doc != "" {
			mdDoc = &lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: fmt.Sprintf("**%s** - *%s*", p.Name(), javadoc.RemoveTokens(doc)),
			}
		}
		fnParams = append(fnParams, lsp.ParameterInformation{
			Label:         p.String(),
			Documentation: mdDoc,
		})
	}

	return lsp.SignatureHelp{
		Signatures: []lsp.SignatureInformation{
			{
				Documentation: &lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: javadoc.MarkdownSimple(fn),
				},
				Label:      fn.String(),
				Parameters: fnParams,
			},
		},
		ActiveParameter: uint32(fnCtx.ParamIdx),
		ActiveSignature: 0,
	}, nil
}

func (h *LspHandler) handleTextDocumentSignatureHelp(req dls.RpcContext, data lsp.TextDocumentPositionParams) error {
	result, err := h.getSignatureInfo(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), result, nil)
}
