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
	FeeType      = types.FeeType
)

var (
	// function aliases
	NewKeeper      = keeper.NewKeeper
	NewQuerier     = keeper.NewQuerier
	ParamKeyTable  = types.ParamKeyTable
	NewParams      = types.NewParams
	DefaultParams  = types.DefaultParams
	ValidateParams = types.ValidateParams

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc                               = types.ModuleCdc
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
)
