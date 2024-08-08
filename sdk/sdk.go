package sdk

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type SubstrateAPI struct {
	RPC    *rpc.RPC
	Client client.Client
}

func NewSDK(url string) (*SubstrateAPI, *types.Metadata, error) {
	cl, err := client.Connect(url)
	if err != nil {
		return nil, nil, err
	}

	newRPC, err := rpc.NewRPC(cl)
	if err != nil {
		return nil, nil, err
	}

	metadata, err := newRPC.State.GetMetadataLatest()
	if err != nil {
		return nil, nil, err
	}

	return &SubstrateAPI{
		RPC:    newRPC,
		Client: cl,
	}, metadata, nil
}

func (api *SubstrateAPI) Close() {
	api.Client.Close()
}

