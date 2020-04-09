package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeType string

const (
	FeeClaimTransaction      FeeType = "ClaimTransactionnFee"
	FeeEvaluationTransaction FeeType = "FeeEvaluationTransaction"
)

const KeyIxoFactor = "ixoFactor"

const KeyInitiationFeeAmount = "InitiationFeeAmount"

const KeyInitiationNodeFeePercentage = "InitiationNodeFeePercentage"

const KeyClaimFeeAmount = "ClaimFeeAmount"

const KeyEvaluationFeeAmount = "EvaluationFeeAmount"

const KeyServiceAgentRegistrationFeeAmount = "ServiceAgentRegistrationFeeAmount"

const KeyEvaluationAgentRegistrationFeeAmount = "EvaluationAgentRegistrationFeeAmount"

const KeyNodeFeePercentage = "NodeFeePercentage"

const KeyEvaluationPayFeePercentage = "EvaluationPayFeePercentage"

const KeyEvaluationPayNodeFeePercentage = "EvaluationPayNodeFeePercentage"

var AllFees = []string{
	KeyIxoFactor,
	KeyInitiationFeeAmount,
	KeyInitiationNodeFeePercentage,
	KeyClaimFeeAmount,
	KeyEvaluationFeeAmount,
	KeyServiceAgentRegistrationFeeAmount,
	KeyEvaluationAgentRegistrationFeeAmount,
	KeyNodeFeePercentage,
	KeyEvaluationPayFeePercentage,
	KeyEvaluationPayNodeFeePercentage,
}

type GenesisState struct {
	IxoFactor sdk.Dec `json:"ixoFactor"`

	InitiationFeeAmount         sdk.Dec `json:"initiationFeeAmount"`
	InitiationNodeFeePercentage sdk.Dec `json:"initiationNodeFeePercentage"`

	ClaimFeeAmount      sdk.Dec `json:"claimFeeAmount"`
	EvaluationFeeAmount sdk.Dec `json:"evaluationFeeAmount"`

	ServiceAgentRegistrationFeeAmount    sdk.Dec `json:"serviceAgentRegistrationFeeAmount"`
	EvaluationAgentRegistrationFeeAmount sdk.Dec `json:"evaluationAgentRegistrationFeeAmount"`

	NodeFeePercentage sdk.Dec `json:"nodeFeePercentage"`

	EvaluationPayFeePercentage     sdk.Dec `json:"evaluationPayFeePercentage"`
	EvaluationPayNodeFeePercentage sdk.Dec `json:"evaluationPayNodeFeePercentage"`
}
