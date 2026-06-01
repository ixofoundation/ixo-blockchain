package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

type Keeper struct {
	cdc        codec.Codec
	storeKey   storetypes.StoreKey
	iidKeeper  types.IidKeeper
	authority  string
}

// NewKeeper builds a new names keeper.
//
// authority is the bech32 address allowed to send governance-only messages
// (typically the gov module address).
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	iidKeeper types.IidKeeper,
	authority string,
) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		iidKeeper: iidKeeper,
		authority: authority,
	}
}

// Authority returns the gov authority address that may send Msg{Create,Update}Namespace
// and acts as a fallback signer for SetNameStatus.
func (k Keeper) Authority() string {
	return k.authority
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
