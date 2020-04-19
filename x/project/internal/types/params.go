package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Parameter store keys
var (
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
)

// project parameters
type Params struct {
	ProjectMinimumInitialFunding sdk.Int `json:"project_minimum_initial_funding" yaml:"project_minimum_initial_funding"`
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Int) Params {
	return Params{
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
	}

}

// default project module parameters
func DefaultParams() Params {
	return Params{
		ProjectMinimumInitialFunding: sdk.OneInt(), // 1
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.ProjectMinimumInitialFunding.LT(sdk.ZeroInt()) {
		return fmt.Errorf("project parameter ProjectMinimumInitialFunding should be positive, is %s ", params.ProjectMinimumInitialFunding.String())
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Project Minimum Initial Funding: %s

`, p.ProjectMinimumInitialFunding)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding},
	}
}
