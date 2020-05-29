package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"strings"
)

type Keeper struct {
	cdc                *codec.Codec
	storeKey           sdk.StoreKey
	paramSpace         params.Subspace
	bankKeeper         bank.Keeper
	reservedIdPrefixes []string
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, paramSpace params.Subspace,
	bankKeeper bank.Keeper, reservedIdPrefixes []string) Keeper {
	return Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		paramSpace:         paramSpace.WithKeyTable(types.ParamKeyTable()),
		bankKeeper:         bankKeeper,
		reservedIdPrefixes: reservedIdPrefixes,
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

// IdReserved checks whether the id (belonging to a Fee, FeeContract, or
// Subscription) has a prefix that has been reserved (i.e. should not be used).
func (k Keeper) IdReserved(id string) bool {
	for _, prefix := range k.reservedIdPrefixes {
		if strings.HasPrefix(id, types.FeeIdPrefix+prefix) {
			return true
		}
	}
	return false
}
