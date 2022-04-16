package langserver

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

type EmptyHandler struct{}

var errNotImplemented = errors.New("not implemented")
var ErrUnhandled = errors.New("unhandled")

func (h *EmptyHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
	return errNotImplemented
}

type baseLspHandler struct {
	EmptyHandler
	conn   jsonrpc2.Conn
	logger Logger
}

func (h *baseLspHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
	return fmt.Errorf("%w: %s", ErrUnhandled, req.Method())
}

func replyEither(ctx context.Context, rpc RpcContext, result interface{}, err error) error {
	if err != nil {
		return rpc.Reply(ctx, nil, err)
	}
	return rpc.Reply(ctx, result, nil)
}

func (h *baseLspHandler) replyEither(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request, result interface{}, err error) error {
	if err != nil {
		return reply(ctx, nil, err)
	}
	return reply(ctx, result, nil)
}

func (h *baseLspHandler) LogDebug(format string, params ...interface{}) {
	h.logger.Debugf(format, params...)
}
func (h *baseLspHandler) LogInfo(format string, params ...interface{}) {
	h.logger.Infof(format, params...)
}
func (h *baseLspHandler) LogWarn(format string, params ...interface{}) {
	h.logger.Warnf(format, params...)
}
func (h *baseLspHandler) LogError(format string, params ...interface{}) {
	h.logger.Errorf(format, params...)
}

// workaround for unsupported file paths (invalid file://-prefix )
func fixURI(s string) (string, bool) {
	if !strings.HasPrefix(s, "file:///") {
		// VS Code sends URLs with only two slashes, which are invalid. golang/go#39789.
		if strings.HasPrefix(s, "file://") {
			return "file:///" + s[len("file://"):], true
		}
		return "", false
	}
	return s, true
}

func (h *baseLspHandler) uriToFilename(v uri.URI) string {
	s := string(v)
	if strings.HasPrefix(s, "git:/") {
		return ""
	}
	fixed, ok := fixURI(s)

	if !ok {
		unescaped, err := url.PathUnescape(s)
		if err != nil {
			h.LogWarn("Unsupported URI (not a filepath): %q\n", s)

		} else {
			h.LogWarn("Unsupported URI (not a filepath): %q\n", unescaped)
		}
		return ""
	}
	v = uri.URI(fixed)

	return v.Filename()
}
