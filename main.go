package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/alanshaw/drand-gossipsub-client-demo/util"
	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
	gclient "github.com/drand/drand/lp2p/client"
)

const (
	listenAddr   = "/ip4/0.0.0.0/tcp/4453"
	relayP2PAddr = "/ip4/192.168.1.124/tcp/44544/p2p/12D3KooWAe637xuWdRCYkuaZZce13P1F9zJX5gzGUPWZJpsUGUSH"
)

func main() {
	c := newBasicClient()
	// c := newClientWithInfoFromHTTPEndpoint()
	// c := newClientWithInfoFromHTTPEndpointInsecurely()

	for res := range c.Watch(context.Background()) {
		fmt.Printf("round=%v randomness=%v\n", res.Round(), res.Randomness())
	}
}

func newBasicClient() client.Client {
	// Create libp2p pubsub
	ps := util.NewPubsub(listenAddr, relayP2PAddr)

	// Extract chain info from group TOML
	info := util.ChainInfoFromGroupTOML("/Users/alan/.drand0/groups/drand_group.toml")

	c, err := client.New(gclient.WithPubsub(ps), client.WithChainInfo(info))
	if err != nil {
		panic(err)
	}

	return c
}

func newClientWithInfoFromHTTPEndpoint() client.Client {
	// Create libp2p pubsub
	ps := util.NewPubsub(listenAddr, relayP2PAddr)

	// Chain hash is used to verify endpoints
	hash, err := hex.DecodeString("c599c267a0dd386606f7d6132da8327d57e1004760897c9dd4fb8495c29942b2")
	if err != nil {
		panic(err)
	}

	c, err := client.New(
		gclient.WithPubsub(ps),
		client.WithChainHash(hash),
		client.From(http.ForURLs([]string{"http://127.0.0.1:3002"}, hash)...),
	)
	if err != nil {
		panic(err)
	}

	return c
}

func newClientWithInfoFromHTTPEndpointInsecurely() client.Client {
	// Create libp2p pubsub
	ps := util.NewPubsub(listenAddr, relayP2PAddr)

	c, err := client.New(
		gclient.WithPubsub(ps),
		client.From(http.ForURLs([]string{"http://127.0.0.1:3002"}, nil)...),
		client.Insecurely(),
	)
	if err != nil {
		panic(err)
	}

	return c
}
