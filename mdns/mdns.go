package mdns

// MDNS
// Discover a peer in the network (using mdns), connect to it and open a chat stream. This code is heavily influenced by (and shamelessly copied from) chat-with-rendezvous example
//
// Authors:
// 	Bineesh Lazar
//
// Reference: https://github.com/libp2p/go-libp2p-examples/tree/master/chat-with-mdns

import (
	"context"
	"time"

	host "github.com/libp2p/go-libp2p-host"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

type discoveryNotifee struct {
	PeerChan chan pstore.PeerInfo
}

//interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi pstore.PeerInfo) {
	n.PeerChan <- pi
}

//Initialize the MDNS service
func InitMDNS(ctx context.Context, peerhost host.Host, rendezvous string) chan pstore.PeerInfo {
	// An hour might be a long long period in practical applications. But this is fine for us
	ser, err := discovery.NewMdnsService(ctx, peerhost, time.Hour, rendezvous)
	if err != nil {
		panic(err)
	}

	//register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan pstore.PeerInfo)

	ser.RegisterNotifee(n)
	return n.PeerChan
}
