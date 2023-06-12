#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-chat-rendezvous-v2 exp_chat_rendezvous_v2/main.go

echo "build completed! you can run it using: output/libp2p-exp-chat-rendezvous-v2"