package DaedalusLanguageServer

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

func TestRegister(t *testing.T) {
	d := &RpcMux{pathToType: make(map[string]Handler)}

	got := make(chan protocol.TextDocumentPositionParams, 1)
	d.Register("text/sync", MakeHandler(func(req RpcContext, data protocol.TextDocumentPositionParams) error {
		got <- data
		return req.Reply(req.Context(), data.TextDocument.URI, nil)
	}))

	serverPipe, clientPipe := net.Pipe()
	server := jsonrpc2.NewConn(jsonrpc2.NewStream(serverPipe), jsonrpc2.WithCodec(Codec{}))
	client := jsonrpc2.NewConn(jsonrpc2.NewStream(clientPipe), jsonrpc2.WithCodec(Codec{}))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Go(ctx, func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		var rec ReplyRecorder
		handled, err := d.Handle(ctx, rec.Reply, req)
		if !handled {
			return nil, jsonrpc2.ErrMethodNotFound
		}
		if err != nil {
			return nil, err
		}
		result, replyErr, _ := rec.Result()
		return result, replyErr
	})
	client.Go(ctx, jsonrpc2.MethodNotFoundHandler)
	defer client.Close()
	defer server.Close()

	docURI := protocol.DocumentURI(uri.File(`C:\demo.txt`))
	var result protocol.DocumentURI
	_, err := client.Call(ctx, "text/sync", protocol.TextDocumentPositionParams{TextDocument: protocol.TextDocumentIdentifier{URI: docURI}}, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result != docURI {
		t.Fatalf("got result %q, want %q", result, docURI)
	}
	if data := <-got; data.TextDocument.URI != docURI {
		t.Fatalf("handler got %q, want %q", data.TextDocument.URI, docURI)
	}

	if _, err := client.Call(ctx, "does/not/exist", nil, nil); err == nil {
		t.Fatal("expected error for unknown method")
	}
}
