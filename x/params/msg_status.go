package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ActivatedTypes []string `json:"activated-types"`
}

func ActivatedParamKey(ty string) string {
	return "Activated/" + ty
}

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, ty := range data.ActivatedTypes {
		k.Set(ctx, ActivatedParamKey(ty), true)
	}
	
	return []abci.ValidatorUpdate{}
}
