@ECHO OFF
PUSHD %~dp0
REM Change this to antlr command line:
SET ANTLR=java -jar "antlr-4.13.0-complete.jar"
%ANTLR% -o ../daedalus/parser -Dlanguage=Go -no-visitor Daedalus.g4
REM %ANTLR% -visitor -o ../langserver/parser -Dlanguage=Go Daedalus.g4
popd
PAUSE
