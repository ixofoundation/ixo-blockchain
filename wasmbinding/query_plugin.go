package wasmbinding

import (
	"fmt"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StargateQuerier dispatches whitelisted stargate queries
func StargateQuerier(queryRouter baseapp.GRPCQueryRouter, cdc codec.Codec) func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error) {
	return func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error) {
		protoResponseType, err := getWhitelistedQuery(request.Path)
		if err != nil {
			return nil, err
		}

		// no matter what happens after this point, we must return
		// the response type to prevent sync.Pool from leaking.
		defer returnStargateResponseToPool(request.Path, protoResponseType)

		route := queryRouter.Route(request.Path)
		if route == nil {
			return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("No route to query '%s'", request.Path)}
		}

		res, err := route(ctx, &abci.RequestQuery{
			Data: request.Data,
			Path: request.Path,
		})
		if err != nil {
			return nil, err
		}

		if res.Value == nil {
			return nil, fmt.Errorf("Res returned from abci query route is nil")
		}

		// TODO: use custom one like osmosis, check once packages sorted
		bz, err := wasmkeeper.ConvertProtoToJSONMarshal(cdc, protoResponseType, res.Value)
		if err != nil {
			return nil, err
		}

		return bz, nil
	}
}
