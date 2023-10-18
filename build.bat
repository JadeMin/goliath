@echo off

go build -ldflags="-s -w" -o="./build/goliath.exe" "./src/main.go"