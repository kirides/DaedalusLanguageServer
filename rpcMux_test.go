package DaedalusLanguageServer

import (
	"context"
	"testing"

	"github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

func TestRegister(t *testing.T) {

	d := &RpcMux{pathToType: make(map[string]Handler)}

	d.Register("text/sync", MakeHandler(func(req RpcContext, d protocol.TextDocumentPositionParams) error {
		t.Logf("%#v", d)
		return nil
	}))

	call, _ := jsonrpc2.NewCall(jsonrpc2.NewNumberID(1), "text/sync", protocol.TextDocumentPositionParams{TextDocument: protocol.TextDocumentIdentifier{URI: protocol.DocumentURI(uri.File(`C:\demo.txt`))}})
	handled, err := d.Handle(context.Background(), nil, call)
	_ = handled

	if err != nil {
		txt := err.Error()

		_ = txt
	}

	t.Name()
}
