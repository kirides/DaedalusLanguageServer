//go:build wasip1

package langserver

import (
	"net/url"
	"strings"
	"sync/atomic"
	"unicode"

	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

// Inside a WASI sandbox only the directories the runtime preopened are
// reachable, under names chosen by whoever starts the module. Paths are
// therefore the path component of the client's file URIs, with one exception:
// WASI runtimes disagree on whether a ':' may appear in a mount name, so the
// drive letter of a Windows URI becomes a directory of its own.
//
//	file:///home/me/mods     ->  /home/me/mods    (--dir /home/me)
//	file:///C:/Games/G2/x.d  ->  /c/Games/G2/x.d  (--dir C:\::/c)
//
// A client sends either Windows or POSIX style URIs, never both, so the
// mapping can be undone for URIs sent back by remembering which one it is.
var windowsClient atomic.Bool

func uriToFilename(v lsp.DocumentURI) string {
	fixed, ok := fixURI(string(v))
	if !ok {
		return ""
	}
	u, err := url.ParseRequestURI(fixed)
	if err != nil || u.Scheme != "file" {
		return ""
	}
	if p, ok := driveToDir(u.Path); ok {
		windowsClient.Store(true)
		return p
	}
	return u.Path
}

func filenameToURI(path string) lsp.DocumentURI {
	if windowsClient.Load() {
		path = dirToDrive(path)
	}
	u := url.URL{Scheme: "file", Path: path}
	return lsp.DocumentURI(u.String())
}

// guestPath maps a path as the user would write it on the host, e.g. in the
// settings, to the path it has inside the sandbox.
func guestPath(p string) string {
	if len(p) >= 3 && p[1] == ':' && (p[2] == '\\' || p[2] == '/') && isASCIILetter(p[0]) {
		windowsClient.Store(true)
		return "/" + strings.ToLower(p[:1]) + "/" + strings.ReplaceAll(p[3:], "\\", "/")
	}
	return p
}

// driveToDir turns "/C:/x" into "/c/x".
func driveToDir(p string) (string, bool) {
	if len(p) >= 3 && p[0] == '/' && p[2] == ':' && isASCIILetter(p[1]) && (len(p) == 3 || p[3] == '/') {
		return "/" + strings.ToLower(p[1:2]) + p[3:], true
	}
	return p, false
}

// dirToDrive turns "/c/x" into "/c:/x".
func dirToDrive(p string) string {
	if len(p) >= 2 && p[0] == '/' && isASCIILetter(p[1]) && (len(p) == 2 || p[2] == '/') {
		return "/" + p[1:2] + ":" + p[2:]
	}
	return p
}

func isASCIILetter(b byte) bool { return b < unicode.MaxASCII && unicode.IsLetter(rune(b)) }
