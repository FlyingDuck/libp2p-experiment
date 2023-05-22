package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/FlyingDuck/libp2p-experiment/linglongchat/chatroom"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

func main() {
	nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	roomFlag := flag.String("room", "awesome-chat-room", "name of chat room to join")
	flag.Parse()

	ctx := context.Background()

	// create a new libp2p Host that listens on a random TCP port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	// create a new PubSub service using the GossipSub router
	pbsb, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{ctx: ctx, h: h})
	err = s.Start()
	if err != nil {
		panic(err)
	}

	// - use the nickname from the cli flag, or a default if blank
	// - join the room from the cli flag, or the flag default
	nickname := *nickFlag
	if len(nickname) == 0 {
		nickname = fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(h.ID()))
	}
	roomname := *roomFlag

	room, err := chatroom.JoinRoom(ctx, pbsb, h.ID(), nickname, roomname)
	if err != nil {
		panic(err)
	}

	// draw the UI
	ui := chatroom.NewChatUI(room)
	if err = ui.Run(); err != nil {
		printErr("error running text UI: %s", err)
	}

}

const DiscoveryServiceTag = "pubsub-chat-example"

type discoveryNotifee struct {
	ctx context.Context
	h   host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (notifee *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID)
	err := notifee.h.Connect(notifee.ctx, pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID, err)
	}
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.String()
	return pretty[len(pretty)-8:]
}

func printErr(m string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, m, args...)
}
