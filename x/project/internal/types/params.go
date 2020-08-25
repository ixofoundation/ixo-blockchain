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
	tenPercentFee := sdk.NewDecWithPrec(1, 1)

	return Params{
		IxoDid:                       defaultInvalidBlankDid, // invalid blank
		ProjectMinimumInitialFunding: defaultMinInitFunding,  // 1uixo
		OracleFeePercentage:          tenPercentFee,          // 0.1 (10%)
		NodeFeePercentage:            tenPercentFee,          // 0.1 (10%)
	}
}

// validate params
func ValidateParams(params Params) error {
	if len(params.IxoDid) == 0 {
		return fmt.Errorf("ixo did cannot be empty")
	}
	if params.ProjectMinimumInitialFunding.IsAnyNegative() {
		return fmt.Errorf("project parameter ProjectMinimumInitialFunding should "+
			"be positive, is %s ", params.ProjectMinimumInitialFunding.String())
	}
	if params.OracleFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("project parameter OracleFeePercentage should be >= 0.0, is %s ", params.OracleFeePercentage.String())
	}
	if params.NodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("project parameter NodeFeePercentage should be >= 0.0, is %s ", params.NodeFeePercentage.String())
	}
	if params.OracleFeePercentage.GT(sdk.OneDec()) {
		return fmt.Errorf("project parameter OracleFeePercentage should be <= 1.0, is %s ", params.OracleFeePercentage.String())
	}
	if params.NodeFeePercentage.GT(sdk.OneDec()) {
		return fmt.Errorf("project parameter NodeFeePercentage should be <= 1.0, is %s ", params.NodeFeePercentage.String())
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Ixo Did: 							%s
  Project Minimum Initial Funding: 	%s
  Oracle Fee Percentage:			%s
  Node Fee Percentage:				%s

`,
		p.ProjectMinimumInitialFunding, p.IxoDid,
		p.OracleFeePercentage, p.NodeFeePercentage)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding},
		{KeyOracleFeePercentage, &p.OracleFeePercentage},
		{KeyNodeFeePercentage, &p.NodeFeePercentage},
	}
}
