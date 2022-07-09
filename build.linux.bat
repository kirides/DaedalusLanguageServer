@ECHO OFF
set "GOARCH=amd64"
set "GOOS=linux"


set "GOARCH=amd64"
echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer.x64 -ldflags "-extldflags ^"-static^" -s -w"

set "GOARCH=386"
echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer.x86 -ldflags "-extldflags ^"-static^" -s -w"
