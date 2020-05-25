package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

type Keeper struct {
	cdc        *codec.Codec
	storeKey   sdk.StoreKey
	paramSpace params.Subspace
	bankKeeper bank.Keeper
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, paramSpace params.Subspace,
	bankKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace.WithKeyTable(types.ParamKeyTable()),
		bankKeeper: bankKeeper,
	}
}

// GetParams returns the total set of fees parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of fees parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
