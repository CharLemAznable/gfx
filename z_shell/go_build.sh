#!/bin/sh

# gfx
go build -v ./...

# gfx/ext/agollox
cd ./ext/agollox
go build -v ./...
cd ./../..
