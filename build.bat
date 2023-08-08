@echo off

go build -ldflags="-s -w" -o="./build/main.exe" "./src/main.go"