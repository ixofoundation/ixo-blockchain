package fees

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets the fees onto the chain
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) error {
	keeper.SetRat(ctx, KeyIxoFactor, data.IxoFactor)
	keeper.SetRat(ctx, KeyInitiationFeeAmount, data.InitiationFeeAmount)
	keeper.SetRat(ctx, KeyInitiationNodeFeePercentage, data.InitiationNodeFeePercentage)

	keeper.SetRat(ctx, KeyClaimFeeAmount, data.ClaimFeeAmount)
	keeper.SetRat(ctx, KeyEvaluationFeeAmount, data.EvaluationFeeAmount)

	keeper.SetRat(ctx, KeyServiceAgentRegistrationFeeAmount, data.ServiceAgentRegistrationFeeAmount)
	keeper.SetRat(ctx, KeyEvaluationAgentRegistrationFeeAmount, data.EvaluationAgentRegistrationFeeAmount)

	keeper.SetRat(ctx, KeyNodeFeePercentage, data.NodeFeePercentage)

	keeper.SetRat(ctx, KeyEvaluationPayFeePercentage, data.EvaluationPayFeePercentage)
	keeper.SetRat(ctx, KeyEvaluationPayNodeFeePercentage, data.EvaluationPayNodeFeePercentage)

	return nil
}

// WriteGenesis returns a GenesisState for a given context and keeper. The
// GenesisState will contain the pool, params, validators, and bonds found in
// the keeper.
func WriteGenesis(ctx sdk.Context, keeper Keeper) GenesisState {

	return GenesisState{
		IxoFactor: keeper.GetRat(ctx, KeyIxoFactor),

		InitiationFeeAmount:         keeper.GetRat(ctx, KeyInitiationFeeAmount),
		InitiationNodeFeePercentage: keeper.GetRat(ctx, KeyInitiationNodeFeePercentage),

		ClaimFeeAmount:      keeper.GetRat(ctx, KeyClaimFeeAmount),
		EvaluationFeeAmount: keeper.GetRat(ctx, KeyEvaluationFeeAmount),

		ServiceAgentRegistrationFeeAmount:    keeper.GetRat(ctx, KeyServiceAgentRegistrationFeeAmount),
		EvaluationAgentRegistrationFeeAmount: keeper.GetRat(ctx, KeyEvaluationAgentRegistrationFeeAmount),

		NodeFeePercentage: keeper.GetRat(ctx, KeyNodeFeePercentage),

		EvaluationPayFeePercentage:     keeper.GetRat(ctx, KeyEvaluationPayFeePercentage),
		EvaluationPayNodeFeePercentage: keeper.GetRat(ctx, KeyEvaluationPayNodeFeePercentage),
	}
}

// WriteGenesis returns a GenesisState for a given context and keeper. The
// GenesisState will contain the pool, params, validators, and bonds found in
// the keeper.
func DefaultGenesis() GenesisState {
	ixoFactor := sdk.OneRat() // 1

	initiationFeeAmount := sdk.NewRat(500, 1)    //500
	initiationNodeFeePercentage := sdk.ZeroRat() // 0

	claimFeeAmount := sdk.NewRat(6, 10)      // 0.6
	evaluationFeeAmount := sdk.NewRat(4, 10) //0.4

	serviceAgentRegistrationFeeAmount := sdk.ZeroRat()    // 0
	evaluationAgentRegistrationFeeAmount := sdk.ZeroRat() // 0

	nodeFeePercentage := sdk.NewRat(5, 10) // 0.5

	evaluationPayFeePercentage := sdk.ZeroRat()     // 0
	evaluationPayNodeFeePercentage := sdk.ZeroRat() // 0

	return GenesisState{
		IxoFactor: ixoFactor,

		InitiationFeeAmount:         initiationFeeAmount,
		InitiationNodeFeePercentage: initiationNodeFeePercentage,

		ClaimFeeAmount:      claimFeeAmount,
		EvaluationFeeAmount: evaluationFeeAmount,

		ServiceAgentRegistrationFeeAmount:    serviceAgentRegistrationFeeAmount,
		EvaluationAgentRegistrationFeeAmount: evaluationAgentRegistrationFeeAmount,

		NodeFeePercentage: nodeFeePercentage,

		EvaluationPayFeePercentage:     evaluationPayFeePercentage,
		EvaluationPayNodeFeePercentage: evaluationPayNodeFeePercentage,
	}
}

func floatToRat(val float64) big.Rat {
	r := new(big.Rat)
	r.SetFloat64(val)
	return *r
}
