package DaedalusLanguageServer

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

type RpcContext interface {
	Context() context.Context
	Reply(ctx context.Context, result interface{}, err error) error
	ReplyEither(ctx context.Context, result interface{}, err error) error
	Request() jsonrpc2.Request
}

type rpcContext struct {
	ctx   context.Context
	reply jsonrpc2.Replier
	req   jsonrpc2.Request
}

func (d rpcContext) Context() context.Context { return d.ctx }
func (d rpcContext) Reply(ctx context.Context, result interface{}, err error) error {
	return d.reply(ctx, result, err)
}
func (d rpcContext) Request() jsonrpc2.Request { return d.req }

func (d rpcContext) ReplyEither(ctx context.Context, result interface{}, err error) error {
	if err != nil {
		return d.Reply(ctx, nil, err)
	}
	return d.Reply(ctx, result, nil)
}
