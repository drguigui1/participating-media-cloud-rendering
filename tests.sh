#!/bin/sh

cd src/
go mod init volumetric-cloud
go test ./...
rm go.mod
