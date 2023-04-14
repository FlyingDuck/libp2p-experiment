#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-proxy exp_proxy/proxy.go

echo "build completed! you can run it using: output/libp2p-exp-proxy"