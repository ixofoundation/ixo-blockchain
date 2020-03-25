package fees

import (
	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	FeeType      = types.FeeType
)

var (
	// function aliases
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier

	// variable aliases
	ModuleCdc = types.ModuleCdc

	KeyIxoFactor                            = types.KeyIxoFactor
	KeyNodeFeePercentage                    = types.KeyNodeFeePercentage
	FeeClaimTransaction                     = types.FeeClaimTransaction
	FeeEvaluationTransaction                = types.FeeEvaluationTransaction
	KeyClaimFeeAmount                       = types.KeyClaimFeeAmount
	KeyEvaluationFeeAmount                  = types.KeyEvaluationFeeAmount
	KeyInitiationFeeAmount                  = types.KeyInitiationFeeAmount
	KeyInitiationNodeFeePercentage          = types.KeyInitiationNodeFeePercentage
	KeyServiceAgentRegistrationFeeAmount    = types.KeyServiceAgentRegistrationFeeAmount
	KeyEvaluationAgentRegistrationFeeAmount = types.KeyEvaluationAgentRegistrationFeeAmount
	KeyEvaluationPayFeePercentage           = types.KeyEvaluationPayFeePercentage
	KeyEvaluationPayNodeFeePercentage       = types.KeyEvaluationPayNodeFeePercentage

	DefaultGenesisState = types.DefaultGenesis
)
