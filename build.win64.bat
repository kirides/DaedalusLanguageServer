@ECHO OFF
set "GOARCH=amd64"
set "UPX_PARAMS=-9 --lzma --strip-relocs=0"

echo Building executable for %GOARCH%
REM See https://en.wikipedia.org/wiki/X86-64#Microarchitecture_levels
set "GOAMD64=v2"
go build -o DaedalusLanguageServer.x64.exe -ldflags "-s -w"

EXIT /B

DEL DaedalusLanguageServer_upx.exe
echo Compressing binary using UXP
upx.exe -o"DaedalusLanguageServer_upx.x64.exe" %UPX_PARAMS% DaedalusLanguageServer.x64.exe
