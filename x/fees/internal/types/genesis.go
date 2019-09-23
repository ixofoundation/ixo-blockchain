package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func DefaultGenesis() GenesisState {
	ixoFactor := sdk.OneDec() // 1
	
	initiationFeeAmount := sdk.NewDec(500).Mul(ixo.IxoDecimals) // 500
	initiationNodeFeePercentage := sdk.ZeroDec()                // 0
	
	claimFeeAmount := sdk.NewDec(6).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals)      // 0.6
	evaluationFeeAmount := sdk.NewDec(4).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals) // 0.4
	
	serviceAgentRegistrationFeeAmount := sdk.ZeroDec().Mul(ixo.IxoDecimals)    // 0
	evaluationAgentRegistrationFeeAmount := sdk.ZeroDec().Mul(ixo.IxoDecimals) // 0
	
	nodeFeePercentage := sdk.NewDec(5).Quo(sdk.NewDec(10)) // 0.5
	
	evaluationPayFeePercentage := sdk.NewDec(1).Quo(sdk.NewDec(10))     // 0.1  TODO : Can change this value
	evaluationPayNodeFeePercentage := sdk.NewDec(2).Quo(sdk.NewDec(10)) // 0.2  TODO : Can change this value
	
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
