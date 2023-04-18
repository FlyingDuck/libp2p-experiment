#!/usr/bin/env bash

echo "GOPATH=$GOPATH"

# 脚本
go build -o output/libp2p-exp-chat-pubsub exp_chat_pubsub/main.go exp_chat_pubsub/chatroom.go exp_chat_pubsub/ui.go

echo "build completed! you can run it using: output/libp2p-exp-chat-pubsub"