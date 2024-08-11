package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyNftContractAddress = []byte("NftContractAddress")
	KeyNftContractMinter  = []byte("NftContractMinter")
	KeyCreateSequence     = []byte("CreateSequence")
)

var _ paramstypes.ParamSet = (*Params)(nil)

func validateNftContractAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}

	if len(addr) == 0 {
		return fmt.Errorf("nft contract addresses can not be empty")
	}

	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	return nil
}

func validateCreateSequence(i interface{}) error {
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

func NewParams(nftContractAddress string, nftContractMinter string, createSequence uint64) Params {
	return Params{
		NftContractAddress: nftContractAddress,
		NftContractMinter:  nftContractAddress,
		CreateSequence:     createSequence,
	}
}

func DefaultParams() Params {
	return Params{
		NftContractAddress: "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
		NftContractMinter:  "ixo1g7xtrvc8ejkenee8a3gryvx6d4n9uu6gpsx63z",
		CreateSequence:     0,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.ParamSetPair{Key: KeyNftContractAddress, Value: &p.NftContractAddress, ValidatorFn: validateNftContractAddress},
		paramstypes.ParamSetPair{Key: KeyNftContractMinter, Value: &p.NftContractMinter, ValidatorFn: validateNftContractAddress},
		paramstypes.ParamSetPair{Key: KeyCreateSequence, Value: &p.CreateSequence, ValidatorFn: validateCreateSequence},
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateNftContractAddress(p.NftContractAddress)
	if err != nil {
		return err
	}
	err = validateNftContractAddress(p.NftContractMinter)
	if err != nil {
		return err
	}
	err = validateCreateSequence(p.CreateSequence)
	if err != nil {
		return err
	}
	return nil
}
