package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/did"
)

// Parameter store keys
var (
	KeyIxoDid                       = []byte("IxoDid")
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
)

// project parameters
type Params struct {
	IxoDid                       did.Did `json:"ixo_did" yaml:"ixo_did"`
	ProjectMinimumInitialFunding sdk.Int `json:"project_minimum_initial_funding" yaml:"project_minimum_initial_funding"`
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Int, ixoDid did.Did) Params {
	return Params{
		IxoDid:                       ixoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
	}

}

// default project module parameters
func DefaultParams() Params {
	return Params{
		IxoDid:                       did.Did(""),  // blank
		ProjectMinimumInitialFunding: sdk.OneInt(), // 1
	}
}

// validate params
func ValidateParams(params Params) error {
	if len(params.IxoDid) == 0 {
		return fmt.Errorf("ixo did cannot be empty")
	}
	if params.ProjectMinimumInitialFunding.LT(sdk.ZeroInt()) {
		return fmt.Errorf("project parameter ProjectMinimumInitialFunding should be positive, is %s ", params.ProjectMinimumInitialFunding.String())
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Ixo Did: %s
  Project Minimum Initial Funding: %s

`, p.ProjectMinimumInitialFunding, p.IxoDid)
}

func validateIxoDid(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("ixo factor must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("ixo factor must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("ixo factor too large: %s", v)
	}

	return nil
}

func validateProjectMinimumInitialFunding(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid project minimum initial funding type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("validate project minimum initial funding must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("validate project minimum initial funding must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("validate project minimum initial funding too large: %s", v)
	}

	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid, validateIxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding, validateProjectMinimumInitialFunding},
	}
}
