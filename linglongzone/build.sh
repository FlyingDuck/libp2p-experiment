#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/linglongzone main.go

echo "build completed! you can run it using: output/linglongzone"