package types

import (
	fmt "fmt"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyIxo1155ContractCode = []byte("Ixo1155ContractCode")
)

func validateContractCode(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected uint64", i)
	}

	return nil
}

// ParamTable for module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(ixo1155ContractCode uint64) Params {
	return Params{
		Ixo1155ContractCode: ixo1155ContractCode,
	}
}

func DefaultParams() Params {
	return Params{
		Ixo1155ContractCode: 0,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.ParamSetPair{Key: KeyIxo1155ContractCode, Value: &p.Ixo1155ContractCode, ValidatorFn: validateContractCode},
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateContractCode(p.Ixo1155ContractCode)
	if err != nil {
		return err
	}
	return nil
}
