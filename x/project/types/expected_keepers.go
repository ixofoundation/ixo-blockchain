package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	did "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	payments "github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI
	SetAccount(ctx sdk.Context, acc auth.AccountI)

	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI
}

type DidKeeper interface {
	GetDidDoc(ctx sdk.Context, did did.Did) (did.DidDoc, error)
}

type PaymentsKeeper interface {
	GetPaymentTemplate(ctx sdk.Context, templateId string) (payments.PaymentTemplate, error)
	SetPaymentContract(ctx sdk.Context, contract payments.PaymentContract)

	MustGetPaymentTemplate(ctx sdk.Context, templateId string) payments.PaymentTemplate
	MustGetPaymentContract(ctx sdk.Context, contractId string) payments.PaymentContract

	PaymentContractExists(ctx sdk.Context, contractId string) bool

	EffectPayment(ctx sdk.Context, contractId string) (effected bool, err error)
}

type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}
