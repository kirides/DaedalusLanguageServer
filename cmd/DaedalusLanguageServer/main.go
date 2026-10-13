package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"log/slog"

	"github.com/goccy/go-json"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/langserver"

	"go.lsp.dev/jsonrpc2"
)

func BV(s []debug.BuildSetting, key string) string {
	for _, v := range s {
		if v.Key == key {
			if key == "vcs.revision" && len(v.Value) > 7 {
				return v.Value[:7]
			}
			return v.Value
		}
	}
	return ""
}

func logBuildInfo(log *slog.Logger) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	BV := func(key string) string {
		for _, v := range bi.Settings {
			if v.Key == key {
				if key == "vcs.revision" && len(v.Value) > 7 {
					return v.Value[:7]
				}
				return v.Value
			}
		}
		return ""
	}

	log.Info("Running", "revision", BV("vcs.revision"), "go_version", bi.GoVersion, "built_at", BV("vcs.time"))
}

func main() {
	pprofPort := flag.Int("pprof", -1, "enables pprof on the specified port")
	logLevel := &slog.LevelVar{}
	flag.TextVar(logLevel, "loglevel", logLevel, "debug/info/warn/error")
	flag.Parse()

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})
	logger := slog.New(handler)

	logBuildInfo(logger)

	pprofServer := &pprofServer{}
	defer pprofServer.Stop()

	if *pprofPort > 0 {
		pprofServer.ChangeAddr(fmt.Sprintf("127.0.0.1:%d", *pprofPort))
	}
	ctx, stop := notifyContext()
	defer stop()

	stdin, stdout := stdio()
	conn := connectLanguageServer(&rwc{stdin, stdout})
	lspHandler := langserver.NewLspHandler(conn, newDlsSlog(logger.Handler()))

	lspHandler.OnConfigChanged(func(config langserver.LspConfig) {
		level := strings.ToLower(config.LogLevel)
		if len(level) < 1 {
			return
		}
		switch level[0] {
		case 'd':
			logLevel.Set(slog.LevelDebug)
		case 'w':
			logLevel.Set(slog.LevelWarn)
		case 'e':
			logLevel.Set(slog.LevelError)
		case 'i':
			logLevel.Set(slog.LevelInfo)
		default:
			logLevel.Set(slog.LevelInfo)
		}
	})

	lspHandler.OnConfigChanged(func(config langserver.LspConfig) {
		if *pprofPort <= 0 {
			// only when not set by args
			logger.Info("Updating pprof address", "addr", config.PprofAddr)
			pprofServer.ChangeAddr(config.PprofAddr)
		}
	})

	handleRequest := func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		// Release the read loop, so requests (and cancellations) are processed concurrently.
		jsonrpc2.Async(ctx)

		var rec dls.ReplyRecorder
		err := lspHandler.Handle(ctx, rec.Reply, req)
		result, replyErr, replied := rec.Result()

		if err != nil {
			if errors.Is(err, langserver.ErrUnhandled) {
				logger.Debug(err.Error(), "method", req.Method(), "request", string(req.Params()))
			} else {
				logger.Error("Error", "method", req.Method(), "err", err)
			}
		}
		if !req.IsCall() {
			// an error returned for a notification would tear down the connection
			return nil, nil
		}
		if replied {
			return result, replyErr
		}
		if errors.Is(err, langserver.ErrUnhandled) {
			return nil, fmt.Errorf("%q: %w", req.Method(), jsonrpc2.ErrMethodNotFound)
		}
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	cancelHandler, cancelRequest := jsonrpc2.CancelHandler(handleRequest)

	conn.Go(ctx, func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		if req.Method() == "$/cancelRequest" {
			var payload struct {
				ID json.RawMessage `json:"id"`
			}
			if err := json.Unmarshal(req.Params(), &payload); err != nil || len(payload.ID) == 0 {
				logger.Warn("invalid request, missing \"id\"")
				return nil, nil
			}
			var id jsonrpc2.ID
			if s, err := strconv.Unquote(string(payload.ID)); err == nil {
				id = jsonrpc2.NewStringID(s)
			} else if n, err := strconv.ParseInt(string(payload.ID), 10, 64); err == nil {
				id = jsonrpc2.NewNumberID(n)
			} else {
				logger.Warn("invalid request, malformed \"id\"")
				return nil, nil
			}
			logger.Debug("cancelling request", "request_id", id)
			cancelRequest(id)
			return nil, nil
		}
		return cancelHandler(ctx, req)
	})
	<-conn.Done()
}

type rwc struct {
	r io.ReadCloser
	w io.WriteCloser
}

func (rwc *rwc) Read(b []byte) (int, error)  { return rwc.r.Read(b) }
func (rwc *rwc) Write(b []byte) (int, error) { return rwc.w.Write(b) }
func (rwc *rwc) Close() error {
	rwc.r.Close()
	return rwc.w.Close()
}

func connectLanguageServer(rwc io.ReadWriteCloser) jsonrpc2.Conn {
	bufStream := jsonrpc2.NewStream(rwc)
	rootConn := jsonrpc2.NewConn(bufStream, jsonrpc2.WithCodec(dls.Codec{}))
	return rootConn

}
