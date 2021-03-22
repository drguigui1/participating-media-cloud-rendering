#!/bin/sh

if [ -f "go.mod" ]; then
    rm go.mod
fi

if [ -f "go.sum" ]; then
    rm go.sum
fi

cd src/ && go mod init volumetric-cloud

if [ $1 = "release" ]; then
    go build -ldflags "-s -w"
else
    go build
fi

if [ -f "go.sum" ]; then
    mv go.sum ../
fi

if [ -f "go.mod" ]; then
    mv go.mod ../
fi

mv volumetric-cloud ../
