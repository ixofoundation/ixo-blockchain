package payments

import (
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
)

const (
	ModuleName = types.ModuleName

	PayRemainderPool = types.PayRemainderPool

	PaymentIdPrefix         = types.PaymentIdPrefix
	PaymentTemplateIdPrefix = types.PaymentTemplateIdPrefix
	PaymentContractIdPrefix = types.PaymentContractIdPrefix
	SubscriptionIdPrefix    = types.SubscriptionIdPrefix

	DefaultParamspace = types.DefaultParamspace
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey

	FeeClaimTransaction      = types.FeeClaimTransaction
	FeeEvaluationTransaction = types.FeeEvaluationTransaction
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	FeeType = types.FeeType

	PaymentTemplate   = types.PaymentTemplate
	PaymentContract   = types.PaymentContract
	Distribution      = types.Distribution
	DistributionShare = types.DistributionShare

	Discount  = types.Discount
	Discounts = types.Discounts

	Subscription = types.Subscription
	Period       = types.Period
	BlockPeriod  = types.BlockPeriod
	TimePeriod   = types.TimePeriod

	MsgSetPaymentContractAuthorisation = types.MsgSetPaymentContractAuthorisation
	MsgCreatePaymentTemplate           = types.MsgCreatePaymentTemplate
	MsgCreatePaymentContract           = types.MsgCreatePaymentContract
	MsgCreateSubscription              = types.MsgCreateSubscription
	MsgGrantDiscount                   = types.MsgGrantDiscount
	MsgRevokeDiscount                  = types.MsgRevokeDiscount
	MsgEffectPayment                   = types.MsgEffectPayment
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec
	DefaultParams = types.DefaultParams

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	NewPaymentTemplate           = types.NewPaymentTemplate
	NewPaymentContract           = types.NewPaymentContract
	NewPaymentContractNoDiscount = types.NewPaymentContractNoDiscount
	NewDistribution              = types.NewDistribution
	NewDistributionShare         = types.NewDistributionShare

	NewDiscount  = types.NewDiscount
	NewDiscounts = types.NewDiscounts

	NewSubscription = types.NewSubscription
	NewBlockPeriod  = types.NewBlockPeriod
	NewTimePeriod   = types.NewTimePeriod

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
