package DaedalusLanguageServer

import (
	"github.com/goccy/go-json"
	"go.lsp.dev/jsonrpc2"
)

// Codec is a [jsonrpc2.Codec] backed by goccy/go-json, which the protocol
// types are written against (jsonrpc2's default codec uses json/v2 semantics).
type Codec struct{}

var _ jsonrpc2.Codec = Codec{}

func (Codec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case jsonrpc2.RawMessage:
		if m == nil {
			return []byte("null"), nil
		}
		return m, nil
	case json.RawMessage:
		if m == nil {
			return []byte("null"), nil
		}
		return m, nil
	}
	return json.MarshalWithOption(v, json.DisableHTMLEscape())
}

func (Codec) Unmarshal(data []byte, v interface{}) error {
	switch p := v.(type) {
	case *jsonrpc2.RawMessage:
		*p = append(jsonrpc2.RawMessage(nil), data...)
		return nil
	case *json.RawMessage:
		*p = append(json.RawMessage(nil), data...)
		return nil
	}
	return json.Unmarshal(data, v)
}
