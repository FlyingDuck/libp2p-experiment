package api

import (
	"context"

	"github.com/FlyingDuck/libp2p-experiment/linglongzone/core"
	coreiface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/coreiface/path"
	ipld "github.com/ipfs/go-ipld-format"
)

type CoreAPI struct {
	// todo

	// ONLY for re-applying options in WithOptions, DO NOT USE ANYWHERE ELSE
	nd         *core.LingLongNode
	parentOpts options.ApiSettings
}

func NewCoreAPI(n *core.LingLongNode, opts ...options.ApiOption) (coreiface.CoreAPI, error) {
	parentOpts, err := options.ApiOptions()
	if err != nil {
		return nil, err
	}

	return (&CoreAPI{nd: n, parentOpts: *parentOpts}).WithOptions(opts...)
}

// Unixfs returns an implementation of Unixfs API
func (api *CoreAPI) Unixfs() coreiface.UnixfsAPI {
	return (*UnixfsAPI)(api)
}

// Block returns an implementation of Block API
func (api *CoreAPI) Block() coreiface.BlockAPI {

}

// Dag returns an implementation of Dag API
func (api *CoreAPI) Dag() coreiface.APIDagService {

}

// Name returns an implementation of Name API
func (api *CoreAPI) Name() coreiface.NameAPI {

}

// Key returns an implementation of Key API
func (api *CoreAPI) Key() coreiface.KeyAPI {

}

// Pin returns an implementation of Pin API
func (api *CoreAPI) Pin() coreiface.PinAPI {

}

// Object returns an implementation of Object API
func (api *CoreAPI) Object() coreiface.ObjectAPI {

}

// Dht returns an implementation of Dht API
func (api *CoreAPI) Dht() coreiface.DhtAPI {

}

// Swarm returns an implementation of Swarm API
func (api *CoreAPI) Swarm() coreiface.SwarmAPI {

}

// PubSub returns an implementation of PubSub API
func (api *CoreAPI) PubSub() coreiface.PubSubAPI {

}

// Routing returns an implementation of Routing API
func (api *CoreAPI) Routing() coreiface.RoutingAPI {

}

// ResolvePath resolves the path using Unixfs resolver
func (api *CoreAPI) ResolvePath(context.Context, path.Path) (path.Resolved, error) {

}

// ResolveNode resolves the path (if not resolved already) using Unixfs
// resolver, gets and returns the resolved Node
func (api *CoreAPI) ResolveNode(context.Context, path.Path) (ipld.Node, error) {

}

// WithOptions creates new instance of CoreAPI based on this instance with
// a set of options applied
func (api *CoreAPI) WithOptions(...options.ApiOption) (coreiface.CoreAPI, error) {

}
