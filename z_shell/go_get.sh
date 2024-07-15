#!/bin/sh

# gfx
go get -t ./...

# gfx/ext/agollox
cd ./ext/agollox
go get -t ./...
cd ./../..
