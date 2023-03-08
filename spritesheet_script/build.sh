#!/bin/sh
GOOS=darwin GOARCH=amd64 go build -o plugins/bin/macos/spritesheet spritesheet._go && \
GOOS=linux GOARCH=amd64 go build -o plugins/bin/linux/spritesheet spritesheet._go && \
GOOS=windows GOARCH=amd64 go build -o plugins/bin/win32/spritesheet.exe spritesheet._go