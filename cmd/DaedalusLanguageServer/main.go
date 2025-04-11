package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"

	"log/slog"

	"github.com/goccy/go-json"

	"github.com/kirides/DaedalusLanguageServer/langserver"

	_ "net/http/pprof"

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	conn := connectLanguageServer(&rwc{os.Stdin, os.Stdout})
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

	runningRequests := &sync.Map{}
	cancelRequest := func(id string) {
		v, ok := runningRequests.LoadAndDelete(id)
		if !ok {
			return
		}
		v.(context.CancelFunc)()
	}
	addRequest := func(ctx context.Context, id string) context.Context {
		ctx, cancel := context.WithCancel(ctx)
		runningRequests.Store(id, cancel)
		return ctx
	}

	conn.Go(ctx, func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		go func() {
			if req.Method() == "$/cancelRequest" {
				var idPayload struct {
					Request struct {
						ID string `json:"id"`
					} `json:"request"`
				}
				if err := json.Unmarshal(req.Params(), &idPayload); err != nil {
					logger.Warn("invalid request, missing \"id\"")
					return
				}
				logger.Debug("cancelling request", "request_id", idPayload.Request.ID)
				cancelRequest(idPayload.Request.ID)
				return
			}
			cancelId := ""
			if r, ok := req.(*jsonrpc2.Call); ok {
				id := r.ID()
				idVal, err := (&id).MarshalJSON()
				if err != nil {
					logger.Warn("invalid call, missing \"id\".", "err", err)
					return
				}
				cancelId = strings.Trim(string(idVal), "\"")

				ctx = addRequest(ctx, cancelId)
				defer cancelRequest(cancelId)
			}
			err := lspHandler.Handle(ctx, reply, req)
			if err != nil {
				if errors.Is(err, langserver.ErrUnhandled) {
					logger.Debug(err.Error(), "method", req.Method(), "request", string(req.Params()))
				} else {
					logger.Error("Error", "method", req.Method(), "err", err)
				}
			} else {
				runningRequests.Delete(cancelId)
			}
		}()

		return nil // unhandled
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
	rootConn := jsonrpc2.NewConn(bufStream)
	return rootConn

}
