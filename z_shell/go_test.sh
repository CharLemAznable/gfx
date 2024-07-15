#!/bin/sh

# gfx
go test -v ./... -race -test.bench=.* -coverprofile=coverage.txt -covermode=atomic

# gfx/ext/agollox
cd ./ext/agollox
go test -v ./... -race -test.bench=.* -coverprofile=coverage.txt -covermode=atomic
cd ./../..
