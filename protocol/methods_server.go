package protocol

// list of server methods.
const (
	// MethodCancelRequest method name of "$/cancelRequest".
	MethodCancelRequest = "$/cancelRequest"

	// MethodInitialize method name of "initialize".
	MethodInitialize = "initialize"

	// MethodInitialized method name of "initialized".
	MethodInitialized = "initialized"

	// MethodShutdown method name of "shutdown".
	MethodShutdown = "shutdown"

	// MethodExit method name of "exit".
	MethodExit = "exit"

	// MethodWorkDoneProgressCancel method name of "window/workDoneProgress/cancel".
	MethodWorkDoneProgressCancel = "window/workDoneProgress/cancel"

	// MethodLogTrace method name of "$/logTrace".
	MethodLogTrace = "$/logTrace"

	// MethodSetTrace method name of "$/setTrace".
	MethodSetTrace = "$/setTrace"

	// MethodTextDocumentCodeAction method name of "textDocument/codeAction".
	MethodTextDocumentCodeAction = "textDocument/codeAction"

	// MethodTextDocumentCodeLens method name of "textDocument/codeLens".
	MethodTextDocumentCodeLens = "textDocument/codeLens"

	// MethodCodeLensResolve method name of "codeLens/resolve".
	MethodCodeLensResolve = "codeLens/resolve"

	// MethodTextDocumentColorPresentation method name of "textDocument/colorPresentation".
	MethodTextDocumentColorPresentation = "textDocument/colorPresentation"

	// MethodTextDocumentCompletion method name of "textDocument/completion".
	MethodTextDocumentCompletion = "textDocument/completion"

	// MethodCompletionItemResolve method name of "completionItem/resolve".
	MethodCompletionItemResolve = "completionItem/resolve"

	// MethodTextDocumentDeclaration method name of "textDocument/declaration".
	MethodTextDocumentDeclaration = "textDocument/declaration"

	// MethodTextDocumentDefinition method name of "textDocument/definition".
	MethodTextDocumentDefinition = "textDocument/definition"

	// MethodTextDocumentDidChange method name of "textDocument/didChange".
	MethodTextDocumentDidChange = "textDocument/didChange"

	// MethodWorkspaceDidChangeConfiguration method name of "workspace/didChangeConfiguration".
	MethodWorkspaceDidChangeConfiguration = "workspace/didChangeConfiguration"

	// MethodWorkspaceDidChangeWatchedFiles method name of "workspace/didChangeWatchedFiles".
	MethodWorkspaceDidChangeWatchedFiles = "workspace/didChangeWatchedFiles"

	// MethodWorkspaceDidChangeWorkspaceFolders method name of "workspace/didChangeWorkspaceFolders".
	MethodWorkspaceDidChangeWorkspaceFolders = "workspace/didChangeWorkspaceFolders"

	// MethodTextDocumentDidClose method name of "textDocument/didClose".
	MethodTextDocumentDidClose = "textDocument/didClose"

	// MethodTextDocumentDidOpen method name of "textDocument/didOpen".
	MethodTextDocumentDidOpen = "textDocument/didOpen"

	// MethodTextDocumentDidSave method name of "textDocument/didSave".
	MethodTextDocumentDidSave = "textDocument/didSave"

	// MethodTextDocumentDocumentColor method name of"textDocument/documentColor".
	MethodTextDocumentDocumentColor = "textDocument/documentColor"

	// MethodTextDocumentDocumentHighlight method name of "textDocument/documentHighlight".
	MethodTextDocumentDocumentHighlight = "textDocument/documentHighlight"

	// MethodTextDocumentDocumentLink method name of "textDocument/documentLink".
	MethodTextDocumentDocumentLink = "textDocument/documentLink"

	// MethodDocumentLinkResolve method name of "documentLink/resolve".
	MethodDocumentLinkResolve = "documentLink/resolve"

	// MethodTextDocumentDocumentSymbol method name of "textDocument/documentSymbol".
	MethodTextDocumentDocumentSymbol = "textDocument/documentSymbol"

	// MethodWorkspaceExecuteCommand method name of "workspace/executeCommand".
	MethodWorkspaceExecuteCommand = "workspace/executeCommand"

	// MethodTextDocumentFoldingRange method name of "textDocument/foldingRange".
	MethodTextDocumentFoldingRange = "textDocument/foldingRange"

	// MethodTextDocumentFormatting method name of "textDocument/formatting".
	MethodTextDocumentFormatting = "textDocument/formatting"

	// MethodTextDocumentHover method name of "textDocument/hover".
	MethodTextDocumentHover = "textDocument/hover"

	// MethodTextDocumentImplementation method name of "textDocument/implementation".
	MethodTextDocumentImplementation = "textDocument/implementation"

	// MethodTextDocumentOnTypeFormatting method name of "textDocument/onTypeFormatting".
	MethodTextDocumentOnTypeFormatting = "textDocument/onTypeFormatting"

	// MethodTextDocumentPrepareRename method name of "textDocument/prepareRename".
	MethodTextDocumentPrepareRename = "textDocument/prepareRename"

	// MethodTextDocumentRangeFormatting method name of "textDocument/rangeFormatting".
	MethodTextDocumentRangeFormatting = "textDocument/rangeFormatting"

	// MethodTextDocumentReferences method name of "textDocument/references".
	MethodTextDocumentReferences = "textDocument/references"

	// MethodTextDocumentRename method name of "textDocument/rename".
	MethodTextDocumentRename = "textDocument/rename"

	// MethodTextDocumentSignatureHelp method name of "textDocument/signatureHelp".
	MethodTextDocumentSignatureHelp = "textDocument/signatureHelp"

	// MethodWorkspaceSymbol method name of "workspace/symbol".
	MethodWorkspaceSymbol = "workspace/symbol"

	// MethodTextDocumentTypeDefinition method name of "textDocument/typeDefinition".
	MethodTextDocumentTypeDefinition = "textDocument/typeDefinition"

	// MethodTextDocumentWillSave method name of "textDocument/willSave".
	MethodTextDocumentWillSave = "textDocument/willSave"

	// MethodTextDocumentWillSaveWaitUntil method name of "textDocument/willSaveWaitUntil".
	MethodTextDocumentWillSaveWaitUntil = "textDocument/willSaveWaitUntil"

	// MethodShowDocument method name of "window/showDocument".
	MethodShowDocument = "window/showDocument"

	// MethodWillCreateFiles method name of "workspace/willCreateFiles".
	MethodWillCreateFiles = "workspace/willCreateFiles"

	// MethodDidCreateFiles method name of "workspace/didCreateFiles".
	MethodDidCreateFiles = "workspace/didCreateFiles"

	// MethodWillRenameFiles method name of "workspace/willRenameFiles".
	MethodWillRenameFiles = "workspace/willRenameFiles"

	// MethodDidRenameFiles method name of "workspace/didRenameFiles".
	MethodDidRenameFiles = "workspace/didRenameFiles"

	// MethodWillDeleteFiles method name of "workspace/willDeleteFiles".
	MethodWillDeleteFiles = "workspace/willDeleteFiles"

	// MethodDidDeleteFiles method name of "workspace/didDeleteFiles".
	MethodDidDeleteFiles = "workspace/didDeleteFiles"

	// MethodCodeLensRefresh method name of "workspace/codeLens/refresh".
	MethodCodeLensRefresh = "workspace/codeLens/refresh"

	// MethodTextDocumentPrepareCallHierarchy method name of "textDocument/prepareCallHierarchy".
	MethodTextDocumentPrepareCallHierarchy = "textDocument/prepareCallHierarchy"

	// MethodCallHierarchyIncomingCalls method name of "callHierarchy/incomingCalls".
	MethodCallHierarchyIncomingCalls = "callHierarchy/incomingCalls"

	// MethodCallHierarchyOutgoingCalls method name of "callHierarchy/outgoingCalls".
	MethodCallHierarchyOutgoingCalls = "callHierarchy/outgoingCalls"

	// MethodSemanticTokensFull method name of "textDocument/semanticTokens/full".
	MethodSemanticTokensFull = "textDocument/semanticTokens/full"

	// MethodSemanticTokensFullDelta method name of "textDocument/semanticTokens/full/delta".
	MethodSemanticTokensFullDelta = "textDocument/semanticTokens/full/delta"

	// MethodSemanticTokensRange method name of "textDocument/semanticTokens/range".
	MethodSemanticTokensRange = "textDocument/semanticTokens/range"

	// MethodSemanticTokensRefresh method name of "workspace/semanticTokens/refresh".
	MethodSemanticTokensRefresh = "workspace/semanticTokens/refresh"

	// MethodLinkedEditingRange method name of "textDocument/linkedEditingRange".
	MethodLinkedEditingRange = "textDocument/linkedEditingRange"

	// MethodMoniker method name of "textDocument/moniker".
	MethodMoniker = "textDocument/moniker"
)
