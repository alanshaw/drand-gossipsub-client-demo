package util

import (
	"github.com/BurntSushi/toml"
	"github.com/drand/drand/chain"
	"github.com/drand/drand/key"
)

// ChainInfoFromGroupTOML reads a drand group TOML file and returns the chain info.
func ChainInfoFromGroupTOML(path string) *chain.Info {
	gt := &key.GroupTOML{}
	_, err := toml.DecodeFile(path, gt)
	if err != nil {
		panic(err)
	}
	g := &key.Group{}
	err = g.FromTOML(gt)
	if err != nil {
		panic(err)
	}
	return chain.NewChainInfo(g)
}
