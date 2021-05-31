package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	//"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	stakingKeeper types.StakingKeeper
	didKeeper     types.DidKeeper

	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace

	cdc codec.BinaryMarshaler
}

func NewKeeper(bankKeeper types.BankKeeper, accountKeeper types.AccountKeeper, stakingKeeper types.StakingKeeper,
	didKeeper types.DidKeeper, storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, cdc codec.BinaryMarshaler) Keeper {

	// ensure batches module account is set
	if addr := accountKeeper.GetModuleAddress(types.BatchesIntermediaryAccount); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BatchesIntermediaryAccount))
	}

	return Keeper{
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		stakingKeeper: stakingKeeper,
		didKeeper:     didKeeper,
		storeKey:      storeKey,
		paramSpace:    paramSpace.WithKeyTable(types.ParamKeyTable()),
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
