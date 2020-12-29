@ECHO OFF
PUSHD %~dp0
REM Change this to antlr command line:
SET ANTLR=java -jar "antlr-4.9-complete.jar"
%ANTLR% -o ../langserver/parser -Dlanguage=Go Daedalus.g4
REM %ANTLR% -visitor -o ../langserver/parser -Dlanguage=Go Daedalus.g4
popd
PAUSE
