package DaedalusLanguageServer

import (
	"context"
	"errors"
	"sync"

	"go.lsp.dev/jsonrpc2"
)

// Replier delivers the response of a request. Only the first reply counts.
type Replier func(ctx context.Context, result interface{}, err error) error

var errAlreadyReplied = errors.New("already replied")

// ReplyRecorder records the first reply of a handler, so it can be returned
// from a direct-return [jsonrpc2.Handler].
type ReplyRecorder struct {
	mu      sync.Mutex
	replied bool
	result  interface{}
	err     error
}

// Reply is a [Replier] that records the response.
func (r *ReplyRecorder) Reply(ctx context.Context, result interface{}, err error) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.replied {
		return errAlreadyReplied
	}
	r.replied = true
	r.result, r.err = result, err
	return nil
}

// Result returns the recorded response and whether there was one.
func (r *ReplyRecorder) Result() (result interface{}, err error, replied bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.result, r.err, r.replied
}

type RpcContext interface {
	Context() context.Context
	Reply(ctx context.Context, result interface{}, err error) error
	ReplyEither(ctx context.Context, result interface{}, err error) error
	Request() *jsonrpc2.Request
}

type rpcContext struct {
	ctx   context.Context
	reply Replier
	req   *jsonrpc2.Request
}

func (d rpcContext) Context() context.Context { return d.ctx }
func (d rpcContext) Reply(ctx context.Context, result interface{}, err error) error {
	return d.reply(ctx, result, err)
}
func (d rpcContext) Request() *jsonrpc2.Request { return d.req }

func (d rpcContext) ReplyEither(ctx context.Context, result interface{}, err error) error {
	if err != nil {
		return d.Reply(ctx, nil, err)
	}
	return d.Reply(ctx, result, nil)
}
