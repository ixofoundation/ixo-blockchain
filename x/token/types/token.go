package types

import (
	fmt "fmt"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyCw20ContractCode    = []byte("Cw20ContractCode")
	KeyCw721ContractCode   = []byte("Cw721ContractCode")
	KeyIxo1155ContractCode = []byte("Ixo1155ContractCode")
)

func validateContractCode(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected uint64", i)
	}

	return nil
}

// ParamTable for project module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(cw20ContractCode uint64, cw721ContractCode uint64, ixo1155ContractCode uint64) Params {
	return Params{
		Cw20ContractCode:    cw20ContractCode,
		Cw721ContractCode:   cw721ContractCode,
		Ixo1155ContractCode: ixo1155ContractCode,
	}
}

func DefaultParams() Params {
	return Params{
		Cw20ContractCode:    0,
		Cw721ContractCode:   0,
		Ixo1155ContractCode: 0,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.ParamSetPair{Key: KeyCw20ContractCode, Value: &p.Cw20ContractCode, ValidatorFn: validateContractCode},
		paramstypes.ParamSetPair{Key: KeyCw721ContractCode, Value: &p.Cw721ContractCode, ValidatorFn: validateContractCode},
		paramstypes.ParamSetPair{Key: KeyIxo1155ContractCode, Value: &p.Ixo1155ContractCode, ValidatorFn: validateContractCode},
	}
}
