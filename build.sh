#!/bin/sh

go build -ldflags="-s -w" -o="./build/goliath" "./src/main.go"