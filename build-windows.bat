@echo off
go generate
go build -ldflags "-H windowsgui" -o print.exe
