#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp0 exp0/main.go

echo "build exp0 completed. you can run it using: output/libp2p-exp0"