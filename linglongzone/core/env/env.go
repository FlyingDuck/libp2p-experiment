package env

import (
	"github.com/FlyingDuck/libp2p-experiment/linglongzone/core"
	"github.com/FlyingDuck/libp2p-experiment/linglongzone/core/api"
	coreiface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
)

func GetAPI() (coreiface.CoreAPI, error) {
	node := &core.LingLongNode{} // todo construct

	fetchBlocks := true
	newAPI, err := api.NewCoreAPI(node, options.Api.FetchBlocks(fetchBlocks))
	if err != nil {
		return nil, err
	}
	return newAPI, nil
}
