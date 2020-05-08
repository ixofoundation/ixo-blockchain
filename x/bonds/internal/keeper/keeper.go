package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	CoinKeeper    bank.Keeper
	SupplyKeeper  supply.Keeper
	accountKeeper auth.AccountKeeper
	StakingKeeper staking.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, supplyKeeper supply.Keeper,
	accountKeeper auth.AccountKeeper, stakingKeeper staking.Keeper,
	storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {

	// ensure batches module account is set
	if addr := supplyKeeper.GetModuleAddress(types.BatchesIntermediaryAccount); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BatchesIntermediaryAccount))
	}

	return Keeper{
		CoinKeeper:    coinKeeper,
		SupplyKeeper:  supplyKeeper,
		accountKeeper: accountKeeper,
		StakingKeeper: stakingKeeper,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
