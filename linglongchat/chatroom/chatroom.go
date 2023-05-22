package chatroom

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
)

func JoinRoom(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, nickname string, roomName string) (*ChatRoom, error) {
	topicname := topicName(roomName)
	// join the pubsub topic
	topic, err := ps.Join(topicname)
	if err != nil {
		return nil, err
	}
	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	cr := &ChatRoom{
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		self:     selfID,
		nick:     nickname,
		roomName: roomName,
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr, nil
}

func (cr *ChatRoom) Publish(input string) error {
	chatMsg := &ChatMessage{
		Message:    input,
		SenderID:   cr.self.String(),
		SenderNick: cr.nick,
		SendTime:   time.Now().Unix(),
	}
	msg, err := jsoniter.Marshal(chatMsg)
	if err != nil {
		return err
	}
	return cr.topic.Publish(cr.ctx, msg)
}

func (cr *ChatRoom) ListPeers() []peer.ID {
	return cr.ps.ListPeers(topicName(cr.roomName))
}

func (cr *ChatRoom) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			fmt.Printf("error happend when receive message, err=%s\n", err)
			close(cr.Messages)
			return
		}

		if msg.ReceivedFrom == cr.self {
			continue
		}

		chatMsg := new(ChatMessage)
		err = jsoniter.Unmarshal(msg.Data, chatMsg)
		if err != nil {
			fmt.Printf("receive an unsupported message")
			continue
		}

		cr.Messages <- chatMsg
	}
}

const ChatRoomBufSize = 128

type ChatRoom struct {
	Messages chan *ChatMessage // Messages is a channel of messages received from other peers in the chat room

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

type ChatMessage struct {
	Message    string `json:"message"`
	SenderID   string `json:"sender_id"`
	SenderNick string `json:"sender_nick"`
	SendTime   int64  `json:"send_time"`
}

func topicName(roomName string) string {
	return "chat-room:" + roomName
}
