package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	//"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	BankKeeper    bank.Keeper
	//SupplyKeeper  supply.Keeper
	accountKeeper auth.AccountKeeper
	StakingKeeper staking.Keeper
	DidKeeper     did.Keeper

	storeKey   sdk.StoreKey
	paramSpace params.Subspace

	cdc *codec.BinaryMarshaler
}

func NewKeeper(bankKeeper bank.Keeper, /*supplyKeeper supply.Keeper,*/
	accountKeeper auth.AccountKeeper, stakingKeeper staking.Keeper,
	didKeeper did.Keeper, storeKey sdk.StoreKey, paramSpace params.Subspace,
	cdc *codec.BinaryMarshaler) Keeper {

	// ensure batches module account is set
	if addr := accountKeeper.GetModuleAddress(types.BatchesIntermediaryAccount); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BatchesIntermediaryAccount))
	}

	return Keeper{
		BankKeeper:    bankKeeper,
		//SupplyKeeper:  supplyKeeper,
		accountKeeper: accountKeeper,
		StakingKeeper: stakingKeeper,
		DidKeeper:     didKeeper,
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
