package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyNftContractAddress = []byte("KeyNftContractAddress")
)

func validateNftContractAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("nft contract adderess can not be empty cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(add)
	if err != nil {
		return err
	}

	return nil
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{}
}

// default project module parameters
func DefaultParams() Params {
	return Params{}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{KeyNftContractAddress, &p.NftContractAddress, validateNftContractAddress},
	}
}
