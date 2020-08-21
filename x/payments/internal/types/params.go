package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// ParamTable for payments module.
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

// validate params
func ValidateParams(params Params) error {
	if params.IxoFactor.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter IxoFactor should be positive, is %s ", params.IxoFactor.String())
	}
	if params.InitiationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter InitiationFeeAmount should be positive, is %s ", params.InitiationFeeAmount.String())
	}
	if params.InitiationNodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter InitiationNodeFeePercentage should be positive, is %s ", params.InitiationNodeFeePercentage.String())
	}
	if params.ClaimFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter ClaimFeeAmount should be positive, is %s ", params.ClaimFeeAmount.String())
	}
	if params.EvaluationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationFeeAmount should be positive, is %s ", params.EvaluationFeeAmount.String())
	}
	if params.ServiceAgentRegistrationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter ServiceAgentRegistrationFeeAmount should be positive, is %s ", params.ServiceAgentRegistrationFeeAmount.String())
	}
	if params.EvaluationAgentRegistrationFeeAmount.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationAgentRegistrationFeeAmount should be positive, is %s ", params.EvaluationAgentRegistrationFeeAmount.String())
	}
	if params.NodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter NodeFeePercentage should be positive, is %s ", params.NodeFeePercentage.String())
	}
	if params.EvaluationPayFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationPayFeePercentage should be positive, is %s ", params.EvaluationPayFeePercentage.String())
	}
	if params.EvaluationPayNodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationPayNodeFeePercentage should be positive, is %s ", params.EvaluationPayNodeFeePercentage.String())
	}
	// TODO: validate according to param upper limits
	return nil
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
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateFeeAmount(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateFeePercentage(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateClaimFeeAmount(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateEvaluationFeeAmount(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateRegistrationFeeAmount(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateAgentRegistrationFeeAmount(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateNodeFeePercentage(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateEvaluationPayFeePercentage(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

func validateEvaluationPayNodeFeePercentage(i interface{}) error {
	ok := i
	if ok != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdkerrors.Wrap(ErrInternal, "unknown type")
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoFactor, &p.IxoFactor, validateFactor},
		{KeyInitiationFeeAmount, &p.InitiationFeeAmount, validateFeeAmount},
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
