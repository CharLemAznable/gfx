#!/bin/sh

# gfx
echo ">>>>>>>> build package: gfx"
go build -v ./...

# gfx/ext/agollox
echo ">>>>>>>> build package: gfx/ext/agollox"
cd ./ext/agollox
go build -v ./...
cd ./../..

# gfx/ext/gviewx/apollo
echo ">>>>>>>> build package: gfx/ext/gviewx/apollo"
cd ./ext/gviewx/apollo
go build -v ./...
cd ./../../..
