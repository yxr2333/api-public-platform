@echo off
setlocal

REM Change the directory to the project root, assuming the script is in the /scripts directory
cd /d "%~dp0.."

REM Build the Go project
go build -o build/main cmd/main.go

endlocal