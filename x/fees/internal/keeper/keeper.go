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

// idReserved checks if the id (from a Fee, FeeContract, or Subscription)
// is using a reserved prefix (i.e. should not be used). The prefixPrefix
// indicates whether we are referring to a Fee, FeeContract, Subscription
func (k Keeper) idReserved(id string, prefixPrefix string) bool {
	for _, prefix := range k.reservedIdPrefixes {
		if strings.HasPrefix(id, prefixPrefix+prefix) {
			return true
		}
	}
	return false
}

func (k Keeper) FeeIdReserved(feeId string) bool {
	return k.idReserved(feeId, types.FeeIdPrefix)
}

func (k Keeper) FeeContractIdReserved(feeId string) bool {
	return k.idReserved(feeId, types.FeeContractIdPrefix)
}

func (k Keeper) SubscriptionIdReserved(feeId string) bool {
	return k.idReserved(feeId, types.SubscriptionIdPrefix)
}
