#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-chat-rendezvous exp_chat_rendezvous/chat.go exp_chat_rendezvous/flags.go

echo "build completed! you can run it using: output/libp2p-exp-chat-rendezvous"