#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-chat-mdns exp_chat_mdns/main.go exp_chat_mdns/flags.go exp_chat_mdns/mdns.go

echo "build completed! you can run it using: output/libp2p-exp-chat-mdns"