#!/bin/sh
GOOS=darwin GOARCH=amd64 go build -o plugins/bin/macos/spritesheet spritesheet.go && \
GOOS=linux GOARCH=amd64 go build -o plugins/bin/linux/spritesheet spritesheet.go && \
GOOS=windows GOARCH=amd64 go build -o plugins/bin/win32/spritesheet.exe spritesheet.go