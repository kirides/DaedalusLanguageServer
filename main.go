package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"langsrv/langserver"

	"net/http"
	// _ "net/http/pprof"

	"go.lsp.dev/jsonrpc2"
)

func main() {
	pprofPort := flag.Int("pprof", -1, "enables pprof on the specified port")
	flag.Parse()
	if *pprofPort > 0 {
		go func() {
			http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *pprofPort), nil)
		}()
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	conn := connectLanguageServer(&rwc{os.Stdin, os.Stdout})
	lspHandler := langserver.NewLspHandler(conn)

	handlers := []jsonrpc2.Handler{
		lspHandler.TextDocumentSyncHandler,
		lspHandler.Handle,
	}
	conn.Go(ctx, func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		var err error
		for _, h := range handlers {
			err = h(ctx, reply, req)
			if err == nil {
				return nil
			}
		}
		if err != nil {
			if !errors.Is(err, langserver.ErrUnhandled) {
				fmt.Fprintf(os.Stderr, "[ERR] %s: %v\n", req.Method(), err)
			}
		}
		return nil
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
