#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-multipro exp_multipro/main.go exp_multipro/node.go exp_multipro/ping.go exp_multipro/echo.go

echo "build completed! you can run it using: output/libp2p-exp-multipro"