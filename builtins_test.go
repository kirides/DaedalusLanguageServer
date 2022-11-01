package DaedalusLanguageServer

import (
	"io/fs"
	"path"
	"testing"
)

func TestBuiltinsCanBeFound(t *testing.T) {
	fs.WalkDir(BuiltinsFS, path.Join(DaedalusBuiltinsPath, "G2A"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		content, err := fs.ReadFile(BuiltinsFS, path)
		if err != nil {
			t.Fatalf("File not found %q", path)
			return nil
		}
		_ = content
		return nil
	})
}
