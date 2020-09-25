package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

// Parameter store keys
var (
	KeyIxoDid                       = []byte("IxoDid")
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
	KeyOracleFeePercentage          = []byte("OracleFeePercentage")
	KeyNodeFeePercentage            = []byte("NodeFeePercentage")
)

// project parameters
type Params struct {
	IxoDid                       did.Did   `json:"ixo_did" yaml:"ixo_did"`
	ProjectMinimumInitialFunding sdk.Coins `json:"project_minimum_initial_funding" yaml:"project_minimum_initial_funding"`
	OracleFeePercentage          sdk.Dec   `json:"oracle_fee_percentage" yaml:"oracle_fee_percentage"`
	NodeFeePercentage            sdk.Dec   `json:"node_fee_percentage" yaml:"node_fee_percentage"`
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Coins, ixoDid did.Did,
	oracleFeePercentage, nodeFeePercentage sdk.Dec) Params {
	return Params{
		IxoDid:                       ixoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
		OracleFeePercentage:          oracleFeePercentage,
		NodeFeePercentage:            nodeFeePercentage,
	}

}

// default project module parameters
func DefaultParams() Params {
	defaultInvalidBlankDid := did.Did("")
	defaultMinInitFunding := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.OneInt()))
	tenPercentFee := sdk.NewDec(10)

	return Params{
		IxoDid:                       defaultInvalidBlankDid, // invalid blank
		ProjectMinimumInitialFunding: defaultMinInitFunding,  // 1uixo
		OracleFeePercentage:          tenPercentFee,          // 10.0 (10%)
		NodeFeePercentage:            tenPercentFee,          // 10.0 (10%)
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Ixo Did:                         %s
  Project Minimum Initial Funding: %s
  Oracle Fee Percentage:           %s
  Node Fee Percentage:             %s

`,
		p.ProjectMinimumInitialFunding, p.IxoDid,
		p.OracleFeePercentage, p.NodeFeePercentage)
}

func validateIxoDid(i interface{}) error {
	v, ok := i.(did.Did)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("ixo did cannot be empty")
	}

	return nil
}

func validateProjectMinimumInitialFunding(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsAnyNegative() {
		return fmt.Errorf("invalid project minimum initial "+
			"funding should be positive, is %s ", v.String())
	}

	return nil
}

func validateOracleFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("project parameter OracleFeePercentage should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("project parameter OracleFeePercentage should be <= 100, is %s ", v.String())
	}

	return nil
}

func validateNodeFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("project parameter NodeFeePercentage should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("project parameter NodeFeePercentage should be <= 100, is %s ", v.String())
	}

	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid, validateIxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding, validateProjectMinimumInitialFunding},
		{KeyOracleFeePercentage, &p.OracleFeePercentage, validateOracleFeePercentage},
		{KeyNodeFeePercentage, &p.NodeFeePercentage, validateNodeFeePercentage},
	}
}
