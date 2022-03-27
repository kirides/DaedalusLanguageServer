@ECHO OFF
set "GOARCH=386"
set "UPX_PARAMS=-9 --lzma --strip-relocs=0"

echo Building executable for %GOARCH%
go build -o DaedalusLanguageServer.exe -ldflags "-s -w"

EXIT /B

DEL DaedalusLanguageServer_upx.exe
echo Compressing binary using UXP
upx.exe -o"DaedalusLanguageServer_upx.exe" %UPX_PARAMS% DaedalusLanguageServer.exe
