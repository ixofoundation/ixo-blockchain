package fees

import (
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

const (
	ModuleName        = types.ModuleName
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

	Discount  = types.Discount
	Discounts = types.Discounts

	Subscription        = types.Subscription
	SubscriptionContent = types.SubscriptionContent
	BlockSubscription   = types.BlockSubscriptionContent
	TimeSubscription    = types.TimeSubscriptionContent
)

var (
	// function aliases
	NewKeeper      = keeper.NewKeeper
	NewQuerier     = keeper.NewQuerier
	RegisterCodec  = types.RegisterCodec
	ParamKeyTable  = types.ParamKeyTable
	NewParams      = types.NewParams
	DefaultParams  = types.DefaultParams
	ValidateParams = types.ValidateParams

	NewFee                = types.NewFee
	NewFeeContent         = types.NewFeeContent
	NewFeeContract        = types.NewFeeContract
	NewFeeContractContent = types.NewFeeContractContent
	NewDistribution       = types.NewDistribution
	NewDistributionShare  = types.NewDistributionShare

	NewDiscount  = types.NewDiscount
	NewDiscounts = types.NewDiscounts

	NewSubscription      = types.NewSubscription
	NewBlockSubscription = types.NewBlockSubscriptionContent
	NewTimeSubscription  = types.NewTimeSubscriptionContent

	ErrNegativeSharePercentage       = types.ErrNegativeSharePercentage
	ErrDistributionPercentagesNot100 = types.ErrDistributionPercentagesNot100
	ErrInvalidGenesis                = types.ErrInvalidGenesis
	ErrSubscriptionHasNoNextPeriod   = types.ErrSubscriptionHasNoNextPeriod

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	KeyIxoFactor                            = types.KeyIxoFactor
	KeyNodeFeePercentage                    = types.KeyNodeFeePercentage
	KeyClaimFeeAmount                       = types.KeyClaimFeeAmount
	KeyEvaluationFeeAmount                  = types.KeyEvaluationFeeAmount
	KeyInitiationFeeAmount                  = types.KeyInitiationFeeAmount
	KeyInitiationNodeFeePercentage          = types.KeyInitiationNodeFeePercentage
	KeyServiceAgentRegistrationFeeAmount    = types.KeyServiceAgentRegistrationFeeAmount
	KeyEvaluationAgentRegistrationFeeAmount = types.KeyEvaluationAgentRegistrationFeeAmount
	KeyEvaluationPayFeePercentage           = types.KeyEvaluationPayFeePercentage
	KeyEvaluationPayNodeFeePercentage       = types.KeyEvaluationPayNodeFeePercentage

	FeeKeyPrefix         = types.FeeKeyPrefix
	FeeContractKeyPrefix = types.FeeContractKeyPrefix
	FeeIdKey             = types.FeeIdKey
	FeeContractIdKey     = types.FeeContractIdKey
)
