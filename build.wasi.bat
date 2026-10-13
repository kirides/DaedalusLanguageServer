@ECHO OFF
REM Builds a single, platform independent WebAssembly (WASI preview 1) module.
REM It runs in any WASI runtime, see README.md for how to start it.
set "GOOS=wasip1"
set "GOARCH=wasm"

echo Building WebAssembly module
go build -o DaedalusLanguageServer.wasm -ldflags "-s -w" ./cmd/DaedalusLanguageServer/
