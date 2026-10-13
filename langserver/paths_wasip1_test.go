//go:build wasip1

package langserver

import (
	"testing"

	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func TestWasiPathMapping(t *testing.T) {
	t.Cleanup(func() { windowsClient.Store(false) })

	t.Run("posix", func(t *testing.T) {
		windowsClient.Store(false)
		const u = "file:///home/me/My%20Mods/a.d"
		p := uriToFilename(u)
		if p != "/home/me/My Mods/a.d" {
			t.Fatalf("unexpected path %q", p)
		}
		if got := filenameToURI(p); got != u {
			t.Fatalf("unexpected uri %q", got)
		}
		// a directory named like a drive is nothing special for a posix client
		if got := filenameToURI("/c/x.d"); got != "file:///c/x.d" {
			t.Fatalf("unexpected uri %q", got)
		}
	})

	t.Run("windows", func(t *testing.T) {
		windowsClient.Store(false)
		for in, want := range map[lsp.DocumentURI]string{
			"file:///C:/Games/G2/a.d":     "/c/Games/G2/a.d",
			"file:///c%3A/Games/G2/a.d":   "/c/Games/G2/a.d",
			"file://C:/Games/G2/a.d":      "/c/Games/G2/a.d", // VS Code, golang/go#39789
			"file:///d:":                  "/d",
			"file:///c:/with%20space/a.d": "/c/with space/a.d",
		} {
			windowsClient.Store(false)
			if got := uriToFilename(in); got != want {
				t.Errorf("uriToFilename(%q) = %q, want %q", in, got, want)
			}
			if !windowsClient.Load() {
				t.Errorf("uriToFilename(%q) did not recognize a windows client", in)
			}
		}
		if got := filenameToURI("/c/Games/G2/a.d"); got != "file:///c:/Games/G2/a.d" {
			t.Errorf("unexpected uri %q", got)
		}
		if got := filenameToURI("/c"); got != "file:///c:" {
			t.Errorf("unexpected uri %q", got)
		}
		if got := filenameToURI("/cc/x.d"); got != "file:///cc/x.d" {
			t.Errorf("unexpected uri %q", got)
		}
	})

	t.Run("settings", func(t *testing.T) {
		windowsClient.Store(false)
		if got := guestPath(`D:\Mods\Gothic.src`); got != "/d/Mods/Gothic.src" {
			t.Errorf("unexpected path %q", got)
		}
		if got := guestPath("Gothic.src"); got != "Gothic.src" {
			t.Errorf("unexpected path %q", got)
		}
		if got := guestPath("/home/me/Gothic.src"); got != "/home/me/Gothic.src" {
			t.Errorf("unexpected path %q", got)
		}
	})
}
