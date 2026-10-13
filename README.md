[![Build](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml/badge.svg)](https://github.com/kirides/DaedalusLanguageServer/actions/workflows/build.yml)

## Usage

see: https://github.com/kirides/vscode-daedalus

## WebAssembly (WASI)

Besides the native executables there is a single, platform independent `DaedalusLanguageServer.wasm`
that runs in any WASI preview 1 runtime (`GOOS=wasip1 GOARCH=wasm`, or `build.wasi.bat`).

```sh
# Linux / macOS: make your projects visible to the server, under their usual path
wasmtime run -W max-wasm-stack=8388608 --dir /home/me DaedalusLanguageServer.wasm

# Windows: a drive is mounted as /<letter>, so C:\Games\Gothic II is seen as /c/Games/Gothic II
wasmtime run -W max-wasm-stack=8388608 --dir C:\::/c DaedalusLanguageServer.wasm
wazero run -cachedir %TEMP%\wazero -mount C:\:/c DaedalusLanguageServer.wasm
```

- The server talks LSP over stdin/stdout like the native build, use the command line above as the server command of your editor.
- A WASI module can only access what the runtime mounts. Mount a directory that contains all your workspaces (a whole drive, or your home directory),
  as the command of an editor cannot change per project. Keep the mount at the path the editor uses (as shown above), the server translates nothing else.
  The server also needs write access to create `.dls/externals` in a workspace when asked to.
- `-W max-wasm-stack=8388608` (wasmtime) keeps the server alive on absurdly nested scripts, which otherwise exhaust the default stack of the runtime.
- The first start compiles the module, which takes a couple of seconds (wasmtime caches the result by default, wazero needs `-cachedir`).
- Parsing is single threaded, `numParserThreads` has no effect and a project takes about 5x as long to parse as with one thread of the native build.
- The pprof server is not available.

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
