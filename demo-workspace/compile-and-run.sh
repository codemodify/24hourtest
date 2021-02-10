#!/bin/env sh

export GO111MODULE=on
export GOPATH="`cd ~/go;pwd`"
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

cd ./contracts

protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	*.proto

cd ..

go mod tidy

go run ./server/main.go &
go run ./client/main.go
