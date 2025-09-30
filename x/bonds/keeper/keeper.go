package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/bonds/types"
)

type Keeper struct {
	BankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	StakingKeeper types.StakingKeeper
	iidKeeper     types.IidKeeper

	storeKey   storetypes.StoreKey
	paramSpace paramstypes.Subspace

	cdc codec.BinaryCodec
}

func NewKeeper(
	cdc codec.BinaryCodec,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper,
	iidKeeper types.IidKeeper,
	storeKey storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
) Keeper {
	// ensure batches module account is set
	if addr := accountKeeper.GetModuleAddress(types.BatchesIntermediaryAccount); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BatchesIntermediaryAccount))
	}

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		BankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		StakingKeeper: stakingKeeper,
		iidKeeper:     iidKeeper,
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		cdc:           cdc,
	}
}

// GetParams returns the total set of bonds parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of bonds parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
