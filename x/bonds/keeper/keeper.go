package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	BankKeeper    bankkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper
	StakingKeeper stakingkeeper.Keeper
	didKeeper     types.DidKeeper

	storeKey   storetypes.StoreKey
	paramSpace paramstypes.Subspace

	cdc codec.BinaryCodec
}

func NewKeeper(bankKeeper bankkeeper.Keeper, accountKeeper authkeeper.AccountKeeper, stakingKeeper stakingkeeper.Keeper,
	didKeeper types.DidKeeper, storeKey storetypes.StoreKey, paramSpace paramstypes.Subspace, cdc codec.BinaryCodec) Keeper {

	// ensure batches module account is set
	if addr := accountKeeper.GetModuleAddress(types.BatchesIntermediaryAccount); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BatchesIntermediaryAccount))
	}

	return Keeper{
		BankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		StakingKeeper: stakingKeeper,
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
