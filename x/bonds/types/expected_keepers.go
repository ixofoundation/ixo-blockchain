package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (iidtypes.IidDocument, bool)
}

type StakingKeeper interface {
	GetParams(ctx context.Context) (params stakingtypes.Params, err error)
}

type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	BlockedAddr(addr sdk.AccAddress) bool
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	IterateAllBalances(ctx context.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool))
}
