package langserver

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestFindFile(t *testing.T) {
	fname := "FINDfile.go"
	f, err := findFile(fname)
	if err != nil {
		t.Fatalf("Err should be nil. %v\n", err)
	}

	expectedfname := "findfile.go"
	if runtime.GOOS == "windows" {
		expectedfname = fname
	}
	p, _ := filepath.Abs(expectedfname)
	if f != p {
		t.Fatalf("filepath does not match. Expected %q, got %q,\n", p, f)
	}
}
