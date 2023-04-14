#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-ping exp_ping/main.go

echo "build exp0 completed. you can run it using: output/libp2p-exp-ping"