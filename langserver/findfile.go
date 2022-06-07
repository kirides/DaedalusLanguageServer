package langserver

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

// Generates case insensitive pattern used with the filepat Glob() funciton
// source: https://wenzr.wordpress.com/2018/04/09/go-glob-case-insensitive/
func InsensitivePattern(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}
	p := strings.Builder{}
	for _, r := range path {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&p, "[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else {
			p.WriteRune(r)
		}
	}
	return p.String()
}

func findFile(file string) (string, error) {
	abs, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	pattern := InsensitivePattern(abs)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", os.ErrNotExist
	}
	return matches[0], nil
}
