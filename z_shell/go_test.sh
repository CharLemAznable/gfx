#!/bin/sh

# gfx
echo ">>>>>>>> test package: gfx"
go test -v ./... -race -test.bench=.* -coverprofile=coverage.txt -covermode=atomic

# gfx/ext/agollox
echo ">>>>>>>> test package: gfx/ext/agollox"
cd ./ext/agollox
go test -v ./... -race -test.bench=.* -coverprofile=coverage.txt -covermode=atomic
cd ./../..

# gfx/ext/gviewx/apollo
echo ">>>>>>>> test package: gfx/ext/gviewx/apollo"
cd ./ext/gviewx/apollo
go test -v ./... -race -test.bench=.* -coverprofile=coverage.txt -covermode=atomic
cd ./../../..
