#!/bin/bash
go get
go mod tidy
CompileDaemon -log-prefix=false -build="go build -race -v -o /usr/build/$GO_BINARY" -command="/usr/build/$GO_BINARY"