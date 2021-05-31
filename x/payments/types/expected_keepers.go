package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	did "github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

type BankKeeper interface {
	InputOutputCoins(ctx sdk.Context, inputs []bank.Input, outputs []bank.Output) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	BlockedAddr(addr sdk.AccAddress) bool
}

type DidKeeper interface {
	GetDidDoc(ctx sdk.Context, did did.Did) (did.DidDoc, error)
}
