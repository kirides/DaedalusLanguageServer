package main

import (
	"context"
	"flag"
	"fmt"
	"os"

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
	bufStream := jsonrpc2.NewStream(os.Stdin, os.Stdout)

	rootConn := jsonrpc2.NewConn(bufStream)

	lspHandler := langserver.NewLspHandler()
	rootConn.AddHandler(lspHandler.TextDocumentSyncHandler)
	rootConn.AddHandler(lspHandler)
	rootConn.Run(context.Background())
}
