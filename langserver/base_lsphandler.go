package langserver

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

type logLevel int

var (
	// LogLevelDebug ...
	LogLevelDebug logLevel = 0
	// LogLevelInfo ...
	LogLevelInfo logLevel = 1
	// LogLevelWarn ...
	LogLevelWarn logLevel = 2
	// LogLevelErr ...
	LogLevelErr logLevel = 3
)

type baseLspHandler struct {
	jsonrpc2.EmptyHandler
	LogLevel logLevel
}

func (h *baseLspHandler) replyEither(ctx context.Context, r *jsonrpc2.Request, result interface{}, err error) {
	if err != nil {
		r.Reply(ctx, nil, err)
	} else {
		r.Reply(ctx, result, nil)
	}
}

func (h *baseLspHandler) LogDebug(format string, params ...interface{}) {
	if h.LogLevel <= LogLevelDebug {
		fmt.Fprintf(os.Stderr, "DEBUG: "+format, params...)
	}
}
func (h *baseLspHandler) LogInfo(format string, params ...interface{}) {
	if h.LogLevel <= LogLevelInfo {
		fmt.Fprintf(os.Stderr, "INFO: "+format, params...)
	}
}
func (h *baseLspHandler) LogWarn(format string, params ...interface{}) {
	if h.LogLevel <= LogLevelWarn {
		fmt.Fprintf(os.Stderr, "WARN: "+format, params...)
	}
}
func (h *baseLspHandler) LogError(format string, params ...interface{}) {
	if h.LogLevel <= LogLevelErr {
		fmt.Fprintf(os.Stderr, "ERROR: "+format, params...)
	}
}

// workaround for unsupported file paths (git + invalid file://-prefix )
func fixURI(s string) (string, bool) {
	if strings.HasPrefix(s, "git:/") {
		return "file:///" + s[5:], true
	}
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
