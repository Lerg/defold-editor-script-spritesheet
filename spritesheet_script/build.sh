#!/bin/sh
GOOS=darwin GOARCH=amd64 go build -o bin/macos/spritesheet spritesheet.go && \
GOOS=linux GOARCH=amd64 go build -o bin/linux/spritesheet spritesheet.go && \
GOOS=windows GOARCH=amd64 go build -o bin/win32/spritesheet.exe spritesheet.go