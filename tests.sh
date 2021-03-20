#!/bin/sh

cd src/
go mod init volumetrical-cloud
go test ./...
rm go.mod
