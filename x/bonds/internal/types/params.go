package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Parameter store keys
var (
	KeyReservedBondTokens = []byte("ReservedBondTokens")
)

// bonds parameters
type Params struct {
	ReservedBondTokens []string `json:"reserved_bond_tokens" yaml:"reserved_bond_tokens"`
}

// ParamTable for bonds module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(reservedBondTokens []string) Params {
	return Params{
		ReservedBondTokens: reservedBondTokens,
	}

}

// default bonds module parameters
func DefaultParams() Params {
	return Params{
		ReservedBondTokens: []string{}, // no reserved bond tokens
	}
}

// validate params
func ValidateParams(params Params) error {
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Bonds Params:
  Reserved Bond Tokens: %s

`,
		p.ReservedBondTokens)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyReservedBondTokens, &p.ReservedBondTokens},
	}
}
