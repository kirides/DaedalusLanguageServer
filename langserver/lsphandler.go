package langserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/google/uuid"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/jsonrpc2"
)

var defaultProjectFiles = []string{"Gothic.src", "Camera.src", "Menu.src", "Music.src", "ParticleFX.src", "SFX.src", "VisualFX.src"}

type InlayHintConfig struct {
	Constants bool `json:"constants"`
}
type LspConfig struct {
	FileEncoding     string          `json:"fileEncoding"`
	SrcFileEncoding  string          `json:"srcFileEncoding"`
	LogLevel         string          `json:"loglevel"`
	PprofAddr        string          `json:"pprofAddr"`
	NumParserThreads int             `json:"numParserThreads"`
	ProjectFiles     []string        `json:"projectFiles"`
	InlayHints       InlayHintConfig `json:"inlayHints"`
}

// LspHandler ...
type LspHandler struct {
	TextDocumentSync *textDocumentSync
	handlers         *dls.RpcMux
	config           LspConfig

	onConfigChangedHandlers []func(config LspConfig)
	baseLspHandler
	initialized     context.Context
	markInitialized func()

	workspaces map[string]*LspWorkspace
}

var (
	// ErrWalkAbort should be returned if a walk function should abort early
	ErrWalkAbort = fmt.Errorf("OK")
)

// NewLspHandler ...
func NewLspHandler(conn jsonrpc2.Conn, logger dls.Logger) *LspHandler {
	initialized, markInitialized := context.WithCancel(context.Background())
	h := &LspHandler{
		baseLspHandler: baseLspHandler{
			logger: logger,
			conn:   conn,
		},
		workspaces:              make(map[string]*LspWorkspace),
		handlers:                dls.NewMux(),
		initialized:             initialized,
		markInitialized:         markInitialized,
		onConfigChangedHandlers: []func(config LspConfig){},
		config: LspConfig{
			FileEncoding:    "1252",
			SrcFileEncoding: "1252",
			LogLevel:        "info",
			ProjectFiles:    defaultProjectFiles,
			InlayHints: InlayHintConfig{
				Constants: true,
			},
		},
	}

	h.TextDocumentSync = &textDocumentSync{
		baseLspHandler: baseLspHandler{
			logger: logger,
			conn:   conn,
		},
		GetWorkspace: h.GetWorkspace,
	}
	return h
}

func (h *LspHandler) OnConfigChanged(handler func(config LspConfig)) {
	h.onConfigChangedHandlers = append(h.onConfigChangedHandlers, handler)
}

func (h *LspHandler) handleWorkspaceExecuteCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	return req.Reply(req.Context(), nil, nil)
}

func (h *LspHandler) onInitialized() {
	h.handlers.Register(lsp.MethodTextDocumentCompletion, dls.MakeHandler(h.handleTextDocumentCompletion))
	h.handlers.Register(lsp.MethodTextDocumentDefinition, dls.MakeHandler(h.handleTextDocumentDefinition))
	h.handlers.Register(lsp.MethodTextDocumentTypeDefinition, dls.MakeHandler(h.handleTextDocumentTypeDefinition))
	h.handlers.Register(lsp.MethodTextDocumentHover, dls.MakeHandler(h.handleTextDocumentHover))
	h.handlers.Register(lsp.MethodTextDocumentSignatureHelp, dls.MakeHandler(h.handleTextDocumentSignatureHelp))
	h.handlers.Register(lsp.MethodTextDocumentImplementation, dls.MakeHandler(h.handleTextDocumentImplementation))
	h.handlers.Register(lsp.MethodSemanticTokensFull, dls.MakeHandler(h.handleSemanticTokensFull))
	h.handlers.Register(lsp.MethodSemanticTokensRange, dls.MakeHandler(h.handleSemanticTokensRange))
	h.handlers.Register(lsp.MethodTextDocumentInlayHint, dls.MakeHandler(h.handleInlayHints))

	h.handlers.Register(lsp.MethodTextDocumentDidOpen, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidOpen))
	h.handlers.Register(lsp.MethodTextDocumentDidChange, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidChange))
	h.handlers.Register(lsp.MethodTextDocumentDidSave, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidSave))
	h.handlers.Register(lsp.MethodTextDocumentDidClose, dls.MakeHandler(h.TextDocumentSync.handleTextDocumentDidClose))

	h.handlers.Register(lsp.MethodTextDocumentDocumentSymbol, dls.MakeHandler(h.handleDocumentSymbol))
	h.handlers.Register(lsp.MethodWorkspaceSymbol, dls.MakeHandler(h.handleWorkspaceSymbol))

	h.handlers.Register(lsp.MethodTextDocumentCodeLens, dls.MakeHandler(h.handleTextDocumentCodeLens))
	h.handlers.Register(lsp.MethodCodeLensResolve, dls.MakeHandler(h.handleCodeLensResolve))
	h.handlers.Register(lsp.MethodTextDocumentReferences, dls.MakeHandler(h.handleTextDocumentReferences))
	h.handlers.Register(lsp.MethodWorkspaceExecuteCommand, dls.MakeHandler(h.handleWorkspaceExecuteCommand))
}

func prettyJSON(val interface{}) string {
	v, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(v)
}

func copyAndCastToStringSlice[T ~string](items []T) []string {
	result := make([]string, len(items))
	for i, v := range items {
		result[i] = string(v)
	}
	return result
}

func (h *LspHandler) GetWorkspace(docURI lsp.DocumentURI) *LspWorkspace {
	p := uriToFilename(docURI)

	for dir, ws := range h.workspaces {
		if strings.HasPrefix(p, dir) {
			return ws
		}
	}
	return nil
}

func (h *LspHandler) initializeWorkspaces(ctx context.Context, workspaceURIs ...string) {
	for _, v := range workspaceURIs {
		p := uriToFilename(lsp.DocumentURI(v))
		if _, ok := h.workspaces[p]; ok {
			// already set-up
			continue
		}
		h.logger.Infof("Setting up workspace %q", p)
		ws := &LspWorkspace{
			bufferManager:   NewBufferManager(),
			parsedDocuments: newParseResultsManager(h.logger),
			config:          h.config,
			logger:          h.logger,
			conn:            h.conn,
			path:            p,
			uri:             lsp.DocumentURI(v),
		}
		ws.workspaceCtx, ws.cancelWorkspaceCtx = context.WithCancel(context.Background())
		h.workspaces[p] = ws
	}
}

func (h *LspHandler) removeWorkspaces(ctx context.Context, workspaceURIs ...string) {
	for _, v := range workspaceURIs {
		p := uriToFilename(lsp.DocumentURI(v))
		if ws, ok := h.workspaces[p]; ok {
			ws.cancelWorkspaceCtx()
			delete(h.workspaces, p)
		}
	}
}

// Deliver ...
func (h *LspHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, r jsonrpc2.Request) error {
	h.LogDebug("Requested %q", r.Method())

	// if r.Params() != nil {
	// 	var paramsMap map[string]interface{}
	// 	json.Unmarshal(r.Params(), &paramsMap)
	// 	fmt.Fprintf(os.Stderr, "Params: %+v\n", paramsMap)
	// }

	switch r.Method() {
	case lsp.MethodWorkspaceDidChangeWorkspaceFolders:
		var params lsp.DidChangeWorkspaceFoldersParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			reply(ctx, nil, err)
		}
		for _, v := range params.Event.Added {
			h.logger.Infof("Adding workspace folder: %s", v)
			h.initializeWorkspaces(ctx, v.URI)
		}
		for _, v := range params.Event.Removed {
			h.logger.Infof("Removing workspace folder: %s", v)
			h.removeWorkspaces(ctx, v.URI)
		}
		reply(ctx, nil, nil)
	case lsp.MethodWorkspaceWorkspaceFolders:
		// if r.Params() != nil {
		// 	var paramsMap map[string]interface{}
		// 	json.Unmarshal(r.Params(), &paramsMap)
		// 	fmt.Fprintf(os.Stderr, "Params: %+v\n", paramsMap)
		// }
		reply(ctx, nil, nil)
	case lsp.MethodInitialize:
		var params lsp.InitializeParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			reply(ctx, nil, err)
			return err
		}
		h.onInitialized()
		h.logger.Infof("Lsp Client: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)

		var workspaceURIs []string
		if len(params.WorkspaceFolders) > 0 {
			for _, v := range params.WorkspaceFolders {
				workspaceURIs = append(workspaceURIs, v.URI)
			}
		} else {
			workspaceURIs = append(workspaceURIs, string(params.RootURI))
		}
		h.initializeWorkspaces(ctx, workspaceURIs...)

		if err := reply(ctx, lsp.InitializeResult{
			ServerInfo: struct {
				Name    string "json:\"name\""
				Version string "json:\"version,omitempty\""
			}{
				Name:    "Daedalus Language Server",
				Version: "dev",
			},
			Capabilities: lsp.ServerCapabilities{
				ExecuteCommandProvider: lsp.ExecuteCommandOptions{
					Commands: []string{CommandSetupWorkspace},
				},
				CompletionProvider: lsp.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
				DefinitionProvider:     true,
				TypeDefinitionProvider: true,
				HoverProvider:          true,
				SignatureHelpProvider: lsp.SignatureHelpOptions{
					TriggerCharacters: []string{"(", ","},
				},
				TextDocumentSync: lsp.TextDocumentSyncOptions{
					Change:    lsp.Full,
					OpenClose: true,
					Save: lsp.SaveOptions{
						IncludeText: true,
					},
				},
				WorkspaceSymbolProvider: true,
				Workspace: lsp.Workspace6Gn{
					WorkspaceFolders: lsp.WorkspaceFolders5Gn{
						Supported:           true,
						ChangeNotifications: uuid.NewString(),
					},
				},
				DocumentSymbolProvider: true,
				ImplementationProvider: true,
				InlayHintProvider: lsp.InlayHintOptions{
					ResolveProvider: false,
				},
				SemanticTokensProvider: &lsp.SemanticTokensOptions{
					Range: true,
					Full:  true,
					Legend: lsp.SemanticTokensLegend{
						TokenTypes:     copyAndCastToStringSlice(SemanticTypes()),
						TokenModifiers: copyAndCastToStringSlice(SemanticModifiers()),
					},
				},
				CodeLensProvider:   &lsp.CodeLensOptions{ResolveProvider: true},
				ReferencesProvider: true,
			},
		}, nil); err != nil {
			return fmt.Errorf("not initialized")
		}
		return nil
	case lsp.MethodWorkspaceDidChangeConfiguration:
		var params struct {
			Settings struct {
				DaedalusLanguageServer LspConfig `json:"daedalusLanguageServer"`
			} `json:"settings"`
		}
		var paramsToMerge struct {
			Settings struct {
				DaedalusLanguageServer json.RawMessage `json:"daedalusLanguageServer"`
			} `json:"settings"`
		}

		var configRaw map[string]interface{}
		_ = json.Unmarshal(r.Params(), &configRaw)
		_ = json.Unmarshal(r.Params(), &paramsToMerge)
		h.LogDebug("%s (debug): %v", r.Method(), prettyJSON(configRaw))

		_ = json.Unmarshal(r.Params(), &params)
		json.Unmarshal(paramsToMerge.Settings.DaedalusLanguageServer, &h.config)
		h.LogInfo("%s: %v", r.Method(), prettyJSON(configRaw))

		for _, ws := range h.workspaces {
			ws.parsedDocuments.SetFileEncoding(h.config.FileEncoding)
			ws.parsedDocuments.SetSrcEncoding(h.config.SrcFileEncoding)
			ws.parsedDocuments.NumParserThreads = h.config.NumParserThreads
		}

		for _, v := range h.onConfigChangedHandlers {
			v(h.config)
		}

		if len(h.config.ProjectFiles) == 0 {
			h.config.ProjectFiles = defaultProjectFiles
		}

		reply(ctx, nil, nil)
		return nil
	case lsp.MethodInitialized:
		reply(ctx, nil, nil)
		h.LogInfo("DLS initialized")
		h.markInitialized()
		return nil
	}

	// Recover if something bad happens in the handlers...
	defer func() {
		err := recover()
		if err != nil {
			h.LogWarn("Recovered from panic at %s: %v\n", r.Method, err)
		}
	}()

	// Init should be done, just to be safe, wait for it to finish
	select {
	case <-h.initialized.Done():
	case <-ctx.Done():
		return nil
	}

	if r.Method() == lsp.MethodTextDocumentDidOpen {
		var openParams lsp.DidOpenTextDocumentParams
		json.Unmarshal(r.Params(), &openParams)

		ws := h.GetWorkspace(openParams.TextDocument.URI)
		if ws != nil && !ws.workspaceInitialized {
			ws.tryInitializeWorkspace(ctx, &openParams, h.config)
		}
	}

	handled, err := h.handlers.Handle(ctx, reply, r)
	if err != nil && handled {
		return err
	}
	if handled {
		return nil
	}
	return h.baseLspHandler.Handle(ctx, reply, r)
}
