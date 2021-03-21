#!/bin/sh

if [ "$#" -ne 1 ]; then
    exit 1
fi

feh --force-aliasing $1
