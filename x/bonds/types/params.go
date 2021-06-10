package types

import (
	"fmt"

	paramstypes"github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyReservedBondTokens = []byte("ReservedBondTokens")
)

// ParamTable for bonds module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
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

func validateReservedBondTokens(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{KeyReservedBondTokens, &p.ReservedBondTokens, validateReservedBondTokens},
	}
}
