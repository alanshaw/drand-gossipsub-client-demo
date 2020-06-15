package util

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
)

// NewPubsub constructs a basic libp2p pubsub module for use with the drand client.
func NewPubsub(listenAddr string, relayAddr string) *pubsub.PubSub {
	h, err := libp2p.New(context.Background(), libp2p.ListenAddrStrings(listenAddr))
	if err != nil {
		panic(err)
	}

	relayMa, err := multiaddr.NewMultiaddr(relayAddr)
	if err != nil {
		panic(err)
	}

	relayAi, err := peer.AddrInfoFromP2pAddr(relayMa)
	if err != nil {
		panic(err)
	}

	dps := []peer.AddrInfo{*relayAi}
	p, err := pubsub.NewGossipSub(context.Background(), h, pubsub.WithDirectPeers(dps))
	if err != nil {
		panic(err)
	}

	return p
}
