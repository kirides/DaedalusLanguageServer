package langserver

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode"
)

// findPath resolves file to an existing path, ignoring the case of its
// components on platforms with case-sensitive file systems (Gothic's scripts
// are written for Windows and routinely disagree with the case on disk).
//
// Only the part of the path that does not exist as spelled is looked up in
// directory listings, so no directory above the deepest existing ancestor has
// to be readable. This matters in sandboxes (WASI), which expose nothing but
// the directories that were explicitly mounted.
func findPath(file string) (string, error) {
	abs, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	if _, err := os.Lstat(abs); err == nil {
		return abs, nil
	}
	if runtime.GOOS == "windows" {
		return "", os.ErrNotExist
	}

	cur := abs
	var missing []string
	for {
		parent := filepath.Dir(cur)
		if parent == cur {
			return "", os.ErrNotExist
		}
		missing = append(missing, filepath.Base(cur))
		cur = parent
		if _, err := os.Lstat(cur); err == nil {
			break
		}
	}
	if caseInsensitiveDir(cur) {
		// the file system already compared the names without regard to case
		return "", os.ErrNotExist
	}
	for i := len(missing) - 1; i >= 0; i-- {
		entries, err := os.ReadDir(cur)
		if err != nil {
			return "", os.ErrNotExist
		}
		found := false
		for _, e := range entries {
			if strings.EqualFold(e.Name(), missing[i]) {
				cur = filepath.Join(cur, e.Name())
				found = true
				break
			}
		}
		if !found {
			return "", os.ErrNotExist
		}
	}
	return cur, nil
}

// caseInsensitiveDir reports whether dir, an existing directory, lives on a
// file system that ignores the case of names (macOS, or Windows underneath
// a sandbox). Listing directories to find a different spelling is pointless
// there, and can be very slow when the sandbox has to describe every entry.
func caseInsensitiveDir(dir string) bool {
	// the probe needs a name with letters in it, an ancestor will do
	for {
		name := filepath.Base(dir)
		other := strings.ToUpper(name)
		if other == name {
			other = strings.ToLower(name)
		}
		if other != name {
			a, err := os.Lstat(dir)
			if err != nil {
				return false
			}
			b, err := os.Lstat(filepath.Join(filepath.Dir(dir), other))
			return err == nil && os.SameFile(a, b)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return false
		}
		dir = parent
	}
}

func findPathAnywhereUpToRoot(dir, segment string) (string, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	p, err := findPath(filepath.Join(dir, segment))
	if err == nil {
		return p, nil
	}

	nextDir := filepath.Dir(absDir)
	if nextDir == absDir {
		return "", fmt.Errorf("could not find %q anywhere", segment)
	}
	return findPathAnywhereUpToRoot(nextDir, segment)
}

func insensitiveRxPattern(path string) string {
	p := strings.Builder{}
	for _, r := range path {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&p, "[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else if r == '*' {
			p.WriteRune('.')
			p.WriteRune(r)
		} else {
			p.WriteString(regexp.QuoteMeta(string(r)))
		}
	}
	return p.String()
}

func ResolvePathsCaseInsensitive(directory, fileName string) (files []string, err error) {
	containingDir, err := findPath(directory)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(containingDir)
	if err != nil {
		return nil, err
	}
	pattern := insensitiveRxPattern(fileName)
	rxName, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	for _, file := range entries {
		if rxName.MatchString(file.Name()) {
			files = append(files, filepath.Join(containingDir, file.Name()))
		}
	}

	return files, nil
}
