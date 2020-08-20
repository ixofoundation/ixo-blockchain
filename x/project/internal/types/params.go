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
)

// project parameters
type Params struct {
	IxoDid                       did.Did   `json:"ixo_did" yaml:"ixo_did"`
	ProjectMinimumInitialFunding sdk.Coins `json:"project_minimum_initial_funding" yaml:"project_minimum_initial_funding"`
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Coins, ixoDid did.Did) Params {
	return Params{
		IxoDid:                       ixoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
	}

}

// default project module parameters
func DefaultParams() Params {
	defaultInvalidBlankDid := did.Did("")
	defaultMinInitFunding := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.OneInt()))

	return Params{
		IxoDid:                       defaultInvalidBlankDid, // invalid blank
		ProjectMinimumInitialFunding: defaultMinInitFunding,  // 1uixo
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
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Ixo Did: %s
  Project Minimum Initial Funding: %s

`, p.ProjectMinimumInitialFunding, p.IxoDid)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding},
	}
}
