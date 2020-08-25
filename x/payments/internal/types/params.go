package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
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

// payments parameters
type Params struct {
	IxoFactor                            sdk.Dec `json:"ixo_factor" yaml:"ixo_factor"`
	InitiationFeeAmount                  sdk.Dec `json:"initiation_fee_amount" yaml:"initiation_fee_amount"`                   // NOT USED
	InitiationNodeFeePercentage          sdk.Dec `json:"initiation_node_fee_percentage" yaml:"initiation_node_fee_percentage"` // NOT USED
	ClaimFeeAmount                       sdk.Dec `json:"claim_fee_amount" yaml:"claim_fee_amount"`
	EvaluationFeeAmount                  sdk.Dec `json:"evaluation_fee_amount" yaml:"evaluation_fee_amount"`
	ServiceAgentRegistrationFeeAmount    sdk.Dec `json:"service_agent_registration_fee_amount" yaml:"service_agent_registration_fee_amount"`       // NOT USED
	EvaluationAgentRegistrationFeeAmount sdk.Dec `json:"evaluation_agent_registration_fee_amount" yaml:"evaluation_agent_registration_fee_amount"` // NOT USED
	NodeFeePercentage                    sdk.Dec `json:"node_fee_percentage" yaml:"node_fee_percentage"`
	EvaluationPayFeePercentage           sdk.Dec `json:"evaluation_pay_fee_percentage" yaml:"evaluation_pay_fee_percentage"`
	EvaluationPayNodeFeePercentage       sdk.Dec `json:"evaluation_pay_node_fee_percentage" yaml:"evaluation_pay_node_fee_percentage"`
}

// ParamTable for payments module
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

// default payments module parameters
func DefaultParams() Params {
	return Params{
		IxoFactor:                            sdk.OneDec(),                                           // 1
		InitiationFeeAmount:                  sdk.NewDec(500).Mul(ixo.IxoDecimals),                   // 500 * 1e8 = 50000000000
		InitiationNodeFeePercentage:          sdk.ZeroDec(),                                          // 0
		ClaimFeeAmount:                       sdk.NewDec(6).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals), // 0.6 * 1e8 = 60000000
		EvaluationFeeAmount:                  sdk.NewDec(4).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals), // 0.4 * 1e8 = 40000000
		ServiceAgentRegistrationFeeAmount:    sdk.ZeroDec().Mul(ixo.IxoDecimals),                     // 0 * 1e8 = 0
		EvaluationAgentRegistrationFeeAmount: sdk.ZeroDec().Mul(ixo.IxoDecimals),                     // 0 * 1e8 = 0
		NodeFeePercentage:                    sdk.NewDec(5).Quo(sdk.NewDec(10)),                      // 0.5
		EvaluationPayFeePercentage:           sdk.NewDec(1).Quo(sdk.NewDec(10)),                      // 0.1
		EvaluationPayNodeFeePercentage:       sdk.NewDec(2).Quo(sdk.NewDec(10)),                      // 0.2
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Payments Params:
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

func validateFactor(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("ixo factor must be positive: %s", v)
	}
	return nil
}

func validateInitiationFeeAmount(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("initiation feeAmount must be positive: %s", v)
	}

	return nil
}

func validateFeePercentage(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("fee percentage must be positive: %s", v)
	}

	return nil
}

func validateClaimFeeAmount(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("claim fee amount must be positive: %s", v)
	}
	return nil
}

func validateEvaluationFeeAmount(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("evaluation fee amount must be positive: %s", v)
	}
	return nil
}

func validateRegistrationFeeAmount(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("registration fee amount must be positive: %s", v)
	}
	return nil
}

func validateAgentRegistrationFeeAmount(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("agent registration fee amount must be positive: %s", v)
	}
	return nil
}

func validateNodeFeePercentage(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("node fee percentage must be positive: %s", v)
	}
	return nil
}

func validateEvaluationPayFeePercentage(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("evaluation pay fee percentage must be positive: %s", v)
	}
	return nil
}

func validateEvaluationPayNodeFeePercentage(i interface{}) error {
	v := i.(sdk.Dec)
	if v.IsNegative() {
		return fmt.Errorf("evaluation pay node fee percentage must be positive: %s", v)
	}
	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoFactor, &p.IxoFactor, validateFactor},
		{KeyInitiationFeeAmount, &p.InitiationFeeAmount, validateInitiationFeeAmount},
		{KeyInitiationNodeFeePercentage, &p.InitiationNodeFeePercentage, validateFeePercentage},
		{KeyClaimFeeAmount, &p.ClaimFeeAmount, validateClaimFeeAmount},
		{KeyEvaluationFeeAmount, &p.EvaluationFeeAmount, validateEvaluationFeeAmount},
		{KeyServiceAgentRegistrationFeeAmount, &p.ServiceAgentRegistrationFeeAmount, validateRegistrationFeeAmount},
		{KeyEvaluationAgentRegistrationFeeAmount, &p.EvaluationAgentRegistrationFeeAmount, validateAgentRegistrationFeeAmount},
		{KeyNodeFeePercentage, &p.NodeFeePercentage, validateNodeFeePercentage},
		{KeyEvaluationPayFeePercentage, &p.EvaluationPayFeePercentage, validateEvaluationPayFeePercentage},
		{KeyEvaluationPayNodeFeePercentage, &p.EvaluationPayNodeFeePercentage, validateEvaluationPayNodeFeePercentage},
	}
}
