#!/bin/sh

# gfx
echo "get package: gfx"
go get -t ./...

# gfx/ext/agollox
echo "get package: gfx/ext/agollox"
cd ./ext/agollox
go get -t ./...
cd ./../..
