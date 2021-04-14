#!/bin/sh

if [ -f "go.mod" ]; then
    rm go.mod
fi

if [ -f "go.sum" ]; then
    rm go.sum
fi

cd src/ && go mod init volumetric-cloud

dlv debug -- fullrender
