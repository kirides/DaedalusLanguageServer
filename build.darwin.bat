@ECHO OFF

set "GOOS=darwin"

set "GOARCH=arm64"
echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer_darwin.arm64 -ldflags "-extldflags ^"-static^" -s -w" ./cmd/DaedalusLanguageServer/

set "GOARCH=amd64"
echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer_darwin.x64 -ldflags "-extldflags ^"-static^" -s -w" ./cmd/DaedalusLanguageServer/
