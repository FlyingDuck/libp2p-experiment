package api

import (
	"context"
	"sync"

	"github.com/FlyingDuck/libp2p-experiment/linglongzone/core"
	coreiface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/coreiface/path"
	"github.com/ipfs/boxo/files"
)

type UnixfsAPI CoreAPI

var nilNode *core.LingLongNode
var once sync.Once

func (api *UnixfsAPI) Add(ctx context.Context, files files.Node, opts ...options.UnixfsAddOption) (path.Resolved, error) {

}

func (api *UnixfsAPI) Get(ctx context.Context, p path.Path) (files.Node, error) {

}

// Ls returns the contents of an IPFS or IPNS object(s) at path p, with the format:
// `<link base58 hash> <link size in bytes> <link name>`
func (api *UnixfsAPI) Ls(ctx context.Context, p path.Path, opts ...options.UnixfsLsOption) (<-chan coreiface.DirEntry, error) {

}
