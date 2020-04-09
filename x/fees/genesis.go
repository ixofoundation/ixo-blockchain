package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/fees/internal/keeper"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) []abci.ValidatorUpdate {
	keeper.SetDec(ctx, KeyIxoFactor, data.IxoFactor)
	keeper.SetDec(ctx, KeyInitiationFeeAmount, data.InitiationFeeAmount)
	keeper.SetDec(ctx, KeyInitiationNodeFeePercentage, data.InitiationNodeFeePercentage)

	keeper.SetDec(ctx, KeyClaimFeeAmount, data.ClaimFeeAmount)
	keeper.SetDec(ctx, KeyEvaluationFeeAmount, data.EvaluationFeeAmount)

	keeper.SetDec(ctx, KeyServiceAgentRegistrationFeeAmount, data.ServiceAgentRegistrationFeeAmount)
	keeper.SetDec(ctx, KeyEvaluationAgentRegistrationFeeAmount, data.EvaluationAgentRegistrationFeeAmount)

	keeper.SetDec(ctx, KeyNodeFeePercentage, data.NodeFeePercentage)

	keeper.SetDec(ctx, KeyEvaluationPayFeePercentage, data.EvaluationPayFeePercentage)
	keeper.SetDec(ctx, KeyEvaluationPayNodeFeePercentage, data.EvaluationPayNodeFeePercentage)

	return []abci.ValidatorUpdate{}
}

func WriteGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {

	return GenesisState{
		IxoFactor: keeper.GetDec(ctx, KeyIxoFactor),

		InitiationFeeAmount:         keeper.GetDec(ctx, KeyInitiationFeeAmount),
		InitiationNodeFeePercentage: keeper.GetDec(ctx, KeyInitiationNodeFeePercentage),

		ClaimFeeAmount:      keeper.GetDec(ctx, KeyClaimFeeAmount),
		EvaluationFeeAmount: keeper.GetDec(ctx, KeyEvaluationFeeAmount),

		ServiceAgentRegistrationFeeAmount:    keeper.GetDec(ctx, KeyServiceAgentRegistrationFeeAmount),
		EvaluationAgentRegistrationFeeAmount: keeper.GetDec(ctx, KeyEvaluationAgentRegistrationFeeAmount),

		NodeFeePercentage: keeper.GetDec(ctx, KeyNodeFeePercentage),

		EvaluationPayFeePercentage:     keeper.GetDec(ctx, KeyEvaluationPayFeePercentage),
		EvaluationPayNodeFeePercentage: keeper.GetDec(ctx, KeyEvaluationPayNodeFeePercentage),
	}
}
