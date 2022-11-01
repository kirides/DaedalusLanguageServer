package langserver

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

const CommandSetupWorkspace = "daedalus.dls-setup-workspace"

type DirFs interface {
	fs.FS
	fs.ReadDirFS
	fs.ReadFileFS
}

func GetExternalsFSByEngine() map[string]DirFs {
	result := map[string]DirFs{}

	engines, err := dls.BuiltinsFS.ReadDir(dls.DaedalusBuiltinsPath)
	if err != nil {
		return result
	}

	for _, v := range engines {
		if !v.IsDir() {
			continue
		}

		result[v.Name()] = dls.BuiltinsFS
	}
	return result
}

type DlsMessageRequest struct {
	lsp.ShowMessageRequestParams

	WorkspaceURI lsp.DocumentURI
}

func (h *LspWorkspace) commandSetupWorkspace(ws *LspWorkspace, argStr string) error {
	engines := GetExternalsFSByEngine()
	h.logger.Debugf("Available Engines:\n%#v", engines)
	engine, ok := engines[argStr]
	if !ok {
		h.logger.Errorf("Engine not supported %q.", argStr)
		return nil
	}

	targetDir := filepath.Join(ws.path, ".dls", "externals")
	err := os.MkdirAll(targetDir, 0640)
	if err != nil && !os.IsExist(err) {
		h.logger.Errorf("Error creating directory %q. %v", targetDir, err)
		return nil
	}
	root := path.Join(dls.DaedalusBuiltinsPath, argStr)
	fs.WalkDir(engine, root, func(entryPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}

		targetFile := filepath.Join(targetDir, strings.TrimPrefix(entryPath, root))
		f, err := os.Create(targetFile)
		if err != nil {
			h.logger.Errorf("failed to create file %q. %v", targetFile, err)
			return nil
		}
		defer f.Close()
		content, err := fs.ReadFile(engine, entryPath)
		if err != nil {
			h.logger.Errorf("failed to read embedded file %q. %v", entryPath, err)
			return nil
		}

		if _, err := f.Write(content); err != nil {
			h.logger.Errorf("failed to write file %q. %v", targetFile, err)
			return nil
		}
		f.Chmod(0444)
		return nil
	})

	return nil
}
