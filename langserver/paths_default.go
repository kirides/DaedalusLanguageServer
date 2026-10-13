//go:build !wasip1

package langserver

import (
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

func uriToFilename(v lsp.DocumentURI) string {
	return uriToFilenameDefault(v)
}

func filenameToURI(path string) lsp.DocumentURI {
	return lsp.DocumentURI(uri.File(path))
}

// guestPath maps a path as the user would write it to the path the server has to use.
func guestPath(p string) string { return p }
