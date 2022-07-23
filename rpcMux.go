package DaedalusLanguageServer

import (
	"context"
	"encoding/json"
	"fmt"

	"go.lsp.dev/jsonrpc2"
)

type Handler func(RpcContext) error

func NewMux() *RpcMux {
	return &RpcMux{
		pathToType: map[string]Handler{},
	}
}

type RpcMux struct {
	pathToType map[string]Handler
}

func (d *RpcMux) Handle(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) (bool, error) {
	if handler, ok := d.pathToType[req.Method()]; ok {
		return true, handler(rpcContext{ctx, reply, req})
	}
	return false, fmt.Errorf("no handler")
}

func (d *RpcMux) Register(p string, fn Handler) {
	d.pathToType[p] = fn
}

func MakeHandler[T any](fn func(req RpcContext, data T) error) Handler {
	return func(req RpcContext) error {
		var val T

		if err := json.Unmarshal(req.Request().Params(), &val); err != nil {
			return err
		}

		return fn(req, val)
	}
}
