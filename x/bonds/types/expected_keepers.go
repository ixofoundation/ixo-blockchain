package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	did "github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	IterateAllBalances(ctx sdk.Context, cb func(sdk.AccAddress, sdk.Coin) bool)
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []bank.Balance
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	BlockedAddr(addr sdk.AccAddress) bool
}

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type StakingKeeper interface {
	GetParams(ctx sdk.Context) staking.Params
}

type DidKeeper interface {
	MustGetDidDoc(ctx sdk.Context, did did.Did) did.DidDoc
	GetDidDoc(ctx sdk.Context, did did.Did) (did.DidDoc, error)
}
