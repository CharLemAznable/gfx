#!/bin/sh

# gfx
echo "build package: gfx"
go build -v ./...

# gfx/ext/agollox
echo "build package: gfx/ext/agollox"
cd ./ext/agollox
go build -v ./...
cd ./../..
