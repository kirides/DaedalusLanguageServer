[![Build](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml/badge.svg)](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml)

## Usage

see: https://github.com/kirides/vscode-daedalus

## Custom externals

The server looks for a `_externals\` directory located in the workspace root.  
When there is a `_externals\externals.src` it will try to parse it and all referenced files right after parsing the built-in externals and before parsing user scripts.  
- If there is no `_externals\externals.src` we look for a `_externals\externals.d` and try to parse that.

This externals should be provided from Union plugins such as zParserExtender.

## Aknowledgements

### Gothic Classic - Nintendo Switch

https://github.com/kirides/DaedalusLanguageServer/assets/13602143/91162090-0b7f-4bbd-a461-4e6f9b1f6078


## neovim configuration

Minimum configuration needed to run the language server


_somewhere in init.lua or any lua script_
```lua
vim.api.nvim_create_autocmd("FileType", {
    pattern = "d",
    callback = function(ev)
        local root_dir = vim.fs.dirname(
            vim.fs.find({ 'Gothic.src', 'Camera.src', 'Menu.src', 'Music.src', 'ParticleFX.src', 'SFX.src', 'VisualFX.src' }, { upward = true })[1]
        )

        local client = vim.lsp.start({
          name = 'DLS',
          cmd = {'C:/.../DaedalusLanguageServer.exe'},
          root_dir = root_dir,
          -- configure language server specific options
          settings = {
            daedalusLanguageServer = {
              loglevel = 'debug',
              inlayHints = { constants = true },
              numParserThreads = 16,
              fileEncoding = 'Windows-1252',
              srcFileEncoding ='Windows-1252',
            },
          },
        })

        vim.lsp.buf_attach_client(0, client)
    end
})


vim.cmd([[highlight! link LspSignatureActiveParameter Search]])
```
