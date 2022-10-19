[![Build](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml/badge.svg)](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml)

## Usage

see: https://github.com/kirides/vscode-daedalus

## Custom externals

The server looks for a `_externals\` directory located in the workspace root.  
When there is a `_externals\externals.src` it will try to parse it and all referenced files right after parsing the built-in externals and before parsing user scripts.  
- If there is no `_externals\externals.src` we look for a `_externals\externals.d` and try to parse that.

This externals should be provided from Union plugins such as zParserExtender.


## neovim configuration

Minimum configuration needed to run the language server


_somewhere in init.lua or any lua script_
```lua
vim.api.nvim_create_autocmd('LspAttach', {
  callback = function(args)
    local client = vim.lsp.get_client_by_id(args.data.client_id)
    if client.server_capabilities.hoverProvider then
      vim.keymap.set('n', 'K', vim.lsp.buf.hover, { buffer = args.buf })
    end
	if client.server_capabilities.signatureHelpProvider then
	  vim.keymap.set('i', '<c-s>', vim.lsp.buf.signature_help, { buffer = args.buf })
	end
  end,
})


local client_id = vim.lsp.start({
	name = 'DLS',
	cmd = {'C:/languageserver/DaedalusLanguageServer.exe'},
})

vim.api.nvim_create_autocmd('FileType', {
  pattern = { 'd' },
  callback = function(args)
    vim.lsp.buf_attach_client(0, client_id)
  end,
})


vim.cmd([[highlight! link LspSignatureActiveParameter Search]])
```