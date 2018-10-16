package fees

// IxoFactor The factor used to adjust fees across the board if the value of the IXO token varies against Fiat 1 = 100%
const IxoFactor = "fee/ixoFactor"

// Fee amounts in ixo tokens

// InitiationFeeAmount The Fee for initiating a project
const InitiationFeeAmount = "fee/InitiationFeeAmount"

// InitiationNodeFeePercentage The percentage of the initiation fee that goes to the node when unsuccessful the remainder goes to ixo
const InitiationNodeFeePercentage = "fee/InitiationNodeFeePercentage"

// ClaimFeeAmount The fee for processing a claim
const ClaimFeeAmount = "fee/ClaimFeeAmount"

// EvaluationFeeAmount The fee for processing an evaluation
const EvaluationFeeAmount = "fee/EvaluationFeeAmount"

// ServiceAgentRegistrationFeeAmount The fee for processing a service agent registration
const ServiceAgentRegistrationFeeAmount = "fee/ServiceAgentRegistrationFeeAmount"

// EvaluationAgentRegistrationFeeAmount The fee for processing an evaluation agent registration
const EvaluationAgentRegistrationFeeAmount = "fee/EvaluationAgentRegistrationFeeAmount"

// NodeFeePercentage The percentage of the fee that the initiating node will receive the remainder goes to ixo
const NodeFeePercentage = "fee/NodeFeePercentage"
