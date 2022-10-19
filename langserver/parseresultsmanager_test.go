package langserver

import (
	"context"
	"path/filepath"
	"testing"
)

func TestParseSourceResolvesInsensitivee(t *testing.T) {

	mgr := newParseResultsManager(nopLogger{})
	p, _ := filepath.Abs(filepath.Join("testdata", "Gothic.src"))
	_, err := mgr.ParseSource(context.TODO(), p)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
