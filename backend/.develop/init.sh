#!/bin/sh
set -e

go install github.com/cosmtrek/air@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
