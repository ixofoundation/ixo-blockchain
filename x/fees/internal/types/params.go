package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// Parameter store keys
var (
	KeyIxoFactor                            = []byte("ixoFactor")
	KeyInitiationFeeAmount                  = []byte("InitiationFeeAmount")
	KeyInitiationNodeFeePercentage          = []byte("InitiationNodeFeePercentage")
	KeyClaimFeeAmount                       = []byte("ClaimFeeAmount")
	KeyEvaluationFeeAmount                  = []byte("EvaluationFeeAmount")
	KeyServiceAgentRegistrationFeeAmount    = []byte("ServiceAgentRegistrationFeeAmount")
	KeyEvaluationAgentRegistrationFeeAmount = []byte("EvaluationAgentRegistrationFeeAmount")
	KeyNodeFeePercentage                    = []byte("NodeFeePercentage")
	KeyEvaluationPayFeePercentage           = []byte("EvaluationPayFeePercentage")
	KeyEvaluationPayNodeFeePercentage       = []byte("EvaluationPayNodeFeePercentage")
)

// fees parameters
type Params struct {
	IxoFactor sdk.Dec `json:"ixo_factor" yaml:"ixo_factor"`

	InitiationFeeAmount         sdk.Dec `json:"initiation_fee_amount" yaml:"initiation_fee_amount"`
	InitiationNodeFeePercentage sdk.Dec `json:"initiation_node_fee_percentage"`

	ClaimFeeAmount      sdk.Dec `json:"claim_fee_amount" yaml:"claim_fee_amount"`
	EvaluationFeeAmount sdk.Dec `json:"evaluation_fee_amount" yaml:"evaluation_fee_amount"`

	ServiceAgentRegistrationFeeAmount    sdk.Dec `json:"service_agent_registration_fee_amount" yaml:"service_agent_registration_fee_amount"`
	EvaluationAgentRegistrationFeeAmount sdk.Dec `json:"evaluation_agent_registration_fee_amount" yaml:"evaluation_agent_registration_fee_amount"`

	NodeFeePercentage sdk.Dec `json:"node_fee_percentage" yaml:"node_fee_percentage"`

	EvaluationPayFeePercentage     sdk.Dec `json:"evaluation_pay_fee_percentage" yaml:"evaluation_pay_fee_percentage"`
	EvaluationPayNodeFeePercentage sdk.Dec `json:"evaluation_pay_node_fee_percentage" yaml:"evaluation_pay_node_fee_percentage"`
}

// ParamTable for fees module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(ixoFactor, initiationFeeAmount, initiationNodeFeePercentage,
	claimFeeAmount, evaluationFeeAmount, serviceAgentRegistrationFeeAmount,
	evaluationAgentRegistrationFeeAmount, nodeFeePercentage,
	evaluationPayFeePercentage, evaluationPayNodeFeePercentage sdk.Dec) Params {

	return Params{
		IxoFactor:                            ixoFactor,
		InitiationFeeAmount:                  initiationFeeAmount,
		InitiationNodeFeePercentage:          initiationNodeFeePercentage,
		ClaimFeeAmount:                       claimFeeAmount,
		EvaluationFeeAmount:                  evaluationFeeAmount,
		ServiceAgentRegistrationFeeAmount:    serviceAgentRegistrationFeeAmount,
		EvaluationAgentRegistrationFeeAmount: evaluationAgentRegistrationFeeAmount,
		NodeFeePercentage:                    nodeFeePercentage,
		EvaluationPayFeePercentage:           evaluationPayFeePercentage,
		EvaluationPayNodeFeePercentage:       evaluationPayNodeFeePercentage,
	}

}

// default fees module parameters
func DefaultParams() Params {
	return Params{
		IxoFactor:                            sdk.OneDec(),                                           // 1
		InitiationFeeAmount:                  sdk.NewDec(500).Mul(ixo.IxoDecimals),                   // 500
		InitiationNodeFeePercentage:          sdk.ZeroDec(),                                          // 0
		ClaimFeeAmount:                       sdk.NewDec(6).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals), // 0.6
		EvaluationFeeAmount:                  sdk.NewDec(4).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals), // 0.4
		ServiceAgentRegistrationFeeAmount:    sdk.ZeroDec().Mul(ixo.IxoDecimals),                     // 0
		EvaluationAgentRegistrationFeeAmount: sdk.ZeroDec().Mul(ixo.IxoDecimals),                     // 0
		NodeFeePercentage:                    sdk.NewDec(5).Quo(sdk.NewDec(10)),                      // 0.5
		EvaluationPayFeePercentage:           sdk.NewDec(1).Quo(sdk.NewDec(10)),                      // 0.1  TODO : Can change this value
		EvaluationPayNodeFeePercentage:       sdk.NewDec(2).Quo(sdk.NewDec(10)),                      // 0.2  TODO : Can change this value
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.IxoFactor.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter IxoFactor should be positive, is %s ", params.IxoFactor.String())
	}
	if params.InitiationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter InitiationFeeAmount should be positive, is %s ", params.InitiationFeeAmount.String())
	}
	if params.InitiationNodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter InitiationNodeFeePercentage should be positive, is %s ", params.InitiationNodeFeePercentage.String())
	}
	if params.ClaimFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter ClaimFeeAmount should be positive, is %s ", params.ClaimFeeAmount.String())
	}
	if params.EvaluationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter EvaluationFeeAmount should be positive, is %s ", params.EvaluationFeeAmount.String())
	}
	if params.ServiceAgentRegistrationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter ServiceAgentRegistrationFeeAmount should be positive, is %s ", params.ServiceAgentRegistrationFeeAmount.String())
	}
	if params.EvaluationAgentRegistrationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter EvaluationAgentRegistrationFeeAmount should be positive, is %s ", params.EvaluationAgentRegistrationFeeAmount.String())
	}
	if params.NodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter NodeFeePercentage should be positive, is %s ", params.NodeFeePercentage.String())
	}
	if params.EvaluationPayFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter EvaluationPayFeePercentage should be positive, is %s ", params.EvaluationPayFeePercentage.String())
	}
	if params.EvaluationPayNodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("fees parameter EvaluationPayNodeFeePercentage should be positive, is %s ", params.EvaluationPayNodeFeePercentage.String())
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Fees Params:
  Ixo Factor:                               %s
  Initiation Fee Amount:                    %s
  Initiation Node Fee Percentage:           %s
  Claim Fee Amount:                         %s
  Evaluation Fee Amount:                    %s
  Service Agent Registration Fee Amount:    %s
  Evaluation Agent Registration Fee Amount: %s
  Node Fee Percentage:                      %s
  Evaluation Pay Fee Percentage:            %s
  Evaluation Pay Node Fee Percentage:       %s

`,
		p.IxoFactor, p.InitiationFeeAmount, p.InitiationNodeFeePercentage,
		p.ClaimFeeAmount, p.EvaluationFeeAmount, p.ServiceAgentRegistrationFeeAmount,
		p.EvaluationAgentRegistrationFeeAmount, p.NodeFeePercentage,
		p.EvaluationPayFeePercentage, p.EvaluationPayNodeFeePercentage,
	)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoFactor, &p.IxoFactor},
		{KeyInitiationFeeAmount, &p.InitiationFeeAmount},
		{KeyInitiationNodeFeePercentage, &p.InitiationNodeFeePercentage},
		{KeyClaimFeeAmount, &p.ClaimFeeAmount},
		{KeyEvaluationFeeAmount, &p.EvaluationFeeAmount},
		{KeyServiceAgentRegistrationFeeAmount, &p.ServiceAgentRegistrationFeeAmount},
		{KeyEvaluationAgentRegistrationFeeAmount, &p.EvaluationAgentRegistrationFeeAmount},
		{KeyNodeFeePercentage, &p.NodeFeePercentage},
		{KeyEvaluationPayFeePercentage, &p.EvaluationPayFeePercentage},
		{KeyEvaluationPayNodeFeePercentage, &p.EvaluationPayNodeFeePercentage},
	}
}
