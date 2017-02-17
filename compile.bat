@echo off&title simo cmd&prompt ^>
set GOPATH=%cd%
set /p biname=Enter the index file name:
if biname=="" biname=main
go build -o bin/%biname%.exe src/%biname%.go
echo +++++++++++++++++++++++++++++++++++++++++++++++++++++++++
echo compile success!
echo ---------------------------------------------------------
echo Usage^:
echo %biname%.exe [args]
echo +++++++++++++++++++++++++++++++++++++++++++++++++++++++++
pause
cd bin
cmd.exe