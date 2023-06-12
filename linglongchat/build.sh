#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/linglongchat main.go

echo "build completed! you can run it using: output/linglongchat"