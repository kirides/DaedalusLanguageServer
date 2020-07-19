package langserver

import (
	"context"
	"fmt"
	"os"

	"go.lsp.dev/jsonrpc2"
)

type baseLspHandler struct {
	jsonrpc2.EmptyHandler
}

func (h *baseLspHandler) replyEither(ctx context.Context, r *jsonrpc2.Request, result interface{}, err error) {
	if err != nil {
		r.Reply(ctx, nil, err)
	} else {
		r.Reply(ctx, result, nil)
	}
}

func (h *baseLspHandler) Log(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params...)
}
