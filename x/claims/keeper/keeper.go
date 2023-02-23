package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramstore    paramtypes.Subspace
		AccountKeeper authkeeper.AccountKeeper
		AuthzKeeper   authzkeeper.Keeper
		IidKeeper     iidkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	iidKeeper iidkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	authzKeeper authzkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		IidKeeper:     iidKeeper,
		AccountKeeper: accountKeeper,
		AuthzKeeper:   authzKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams returns the total set of project parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of project parameters.
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.paramstore.SetParamSet(ctx, params)
}
