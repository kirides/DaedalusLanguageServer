@ECHO OFF
set "GOARCH=amd64"
set "UPX_PARAMS=-9 --lzma --strip-relocs=0"

echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer.x64.exe -ldflags "-s -w"

DEL DaedalusLanguageServer_upx.exe
echo Compressing binary using UXP
upx.exe -o"DaedalusLanguageServer_upx.x64.exe" %UPX_PARAMS% DaedalusLanguageServer.x64.exe
