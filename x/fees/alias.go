package fees

import (
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

const (
	ModuleName = types.ModuleName

	FeeRemainderPool = types.FeeRemainderPool

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

	Fee                = types.Fee
	FeeContent         = types.FeeContent
	FeeContract        = types.FeeContract
	FeeContractContent = types.FeeContractContent
	Distribution       = types.Distribution
	DistributionShare  = types.DistributionShare

	Discount       = types.Discount
	Discounts      = types.Discounts
	DiscountHolder = types.DiscountHolder

	Subscription        = types.Subscription
	SubscriptionContent = types.SubscriptionContent
	BlockSubscription   = types.BlockSubscriptionContent
	TimeSubscription    = types.TimeSubscriptionContent
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

	NewFeeContent         = types.NewFeeContent
	NewFeeContractContent = types.NewFeeContractContent
	NewDistribution       = types.NewDistribution
	NewDistributionShare  = types.NewDistributionShare

	NewDiscount       = types.NewDiscount
	NewDiscounts      = types.NewDiscounts
	NewDiscountHolder = types.NewDiscountHolder

	NewSubscription             = types.NewSubscription
	NewBlockSubscriptionContent = types.NewBlockSubscriptionContent
	NewTimeSubscriptionContent  = types.NewTimeSubscriptionContent

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
