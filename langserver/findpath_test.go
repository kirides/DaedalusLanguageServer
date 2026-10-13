package langserver

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestFindFile(t *testing.T) {
	fname := "FINDpath.go"
	f, err := findPath(fname)
	if err != nil {
		t.Fatalf("Err should be nil. %v\n", err)
	}

	expectedfname := "findpath.go"
	if runtime.GOOS == "windows" {
		expectedfname = fname
	}
	p, _ := filepath.Abs(expectedfname)
	if f != p {
		t.Fatalf("filepath does not match. Expected %q, got %q,\n", p, f)
	}
}

func TestFindPathIgnoresCaseOfComponents(t *testing.T) {
	dir := t.TempDir()
	actual := filepath.Join(dir, "Story", "Sub", "Item.d")
	if err := os.MkdirAll(filepath.Dir(actual), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(actual, []byte("// item"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := findPath(filepath.Join(dir, "story", "SUB", "item.D"))
	if err != nil {
		t.Fatalf("Err should be nil. %v\n", err)
	}
	// a case-insensitive file system may hand back the path as it was spelled
	if !strings.EqualFold(got, actual) {
		t.Fatalf("filepath does not match. Expected %q, got %q,\n", actual, got)
	}
	if _, err := os.Stat(got); err != nil {
		t.Fatalf("returned path should exist: %v", err)
	}
	if _, err := os.Lstat(filepath.Join(dir, "story")); err != nil && got != actual {
		t.Fatalf("filepath has to use the case on disk. Expected %q, got %q,\n", actual, got)
	}

	for _, missing := range []string{
		filepath.Join(dir, "story", "SUB", "other.d"),
		filepath.Join(dir, "nowhere", "item.d"),
		filepath.Join(dir, "story", "SUB", "item.d", "deeper"),
	} {
		if _, err := findPath(missing); !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("findPath(%q): expected ErrNotExist, got %v", missing, err)
		}
	}
}
