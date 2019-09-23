package keeper

import (
	"encoding/json"
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/types"
)

const (
	QueryAllContracts = "queryAllContracts"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abciTypes.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAllContracts:
			return queryAllContracts(ctx, k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown did query endpoint")
		}
	}
}

func queryAllContracts(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	contracts := make(map[string]string)
	
	for _, contract := range types.AllContracts {
		address := k.GetContract(ctx, contract)
		contracts[contract] = address
	}
	
	res, errRes := json.Marshal(contracts)
	if errRes != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to marshal data %s", errRes))
	}
	
	return res, nil
}
