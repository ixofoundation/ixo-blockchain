package fees

import sdk "github.com/cosmos/cosmos-sdk/types"

// IxoFactor The factor used to adjust fees across the board if the value of the IXO token varies against Fiat 1 = 100%
const KeyIxoFactor = "ixoFactor"

// Fee amounts in ixo tokens

// InitiationFeeAmount The Fee for initiating a project
const KeyInitiationFeeAmount = "InitiationFeeAmount"

// InitiationNodeFeePercentage The percentage of the initiation fee that goes to the node when unsuccessful the remainder goes to ixo
const KeyInitiationNodeFeePercentage = "InitiationNodeFeePercentage"

// ClaimFeeAmount The fee for processing a claim
const KeyClaimFeeAmount = "ClaimFeeAmount"

// EvaluationFeeAmount The fee for processing an evaluation
const KeyEvaluationFeeAmount = "EvaluationFeeAmount"

// ServiceAgentRegistrationFeeAmount The fee for processing a service agent registration
const KeyServiceAgentRegistrationFeeAmount = "ServiceAgentRegistrationFeeAmount"

// EvaluationAgentRegistrationFeeAmount The fee for processing an evaluation agent registration
const KeyEvaluationAgentRegistrationFeeAmount = "EvaluationAgentRegistrationFeeAmount"

// NodeFeePercentage The percentage of the fee that the initiating node will receive the remainder goes to ixo
const KeyNodeFeePercentage = "NodeFeePercentage"

// EvaluationPayFeePercentage The generated off the Evaluation pay as a percentage of the payment
const KeyEvaluationPayFeePercentage = "EvaluationPayFeePercentage"

// EvaluationPayNodeFeePercentage The percentage of the evaluation payment fee that the initiating node will receive the remainder goes to ixo
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

// GenesisAccount doesn't need pubkey or sequence
type GenesisState struct {
	IxoFactor sdk.Rat `json:"ixoFactor"`

	InitiationFeeAmount         sdk.Rat `json:"initiationFeeAmount"`
	InitiationNodeFeePercentage sdk.Rat `json:"initiationNodeFeePercentage"`

	ClaimFeeAmount      sdk.Rat `json:"claimFeeAmount"`
	EvaluationFeeAmount sdk.Rat `json:"evaluationFeeAmount"`

	ServiceAgentRegistrationFeeAmount    sdk.Rat `json:"serviceAgentRegistrationFeeAmount"`
	EvaluationAgentRegistrationFeeAmount sdk.Rat `json:"evaluationAgentRegistrationFeeAmount"`

	NodeFeePercentage sdk.Rat `json:"nodeFeePercentage"`

	EvaluationPayFeePercentage     sdk.Rat `json:"evaluationPayFeePercentage"`
	EvaluationPayNodeFeePercentage sdk.Rat `json:"evaluationPayNodeFeePercentage"`
}
