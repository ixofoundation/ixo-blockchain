package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyNftContractAddress = []byte("NftContractAddress")
	KeyNftContractMinter  = []byte("NftContractMinter")
)

func validateNftContractAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}

	if len(addr) == 0 {
		return fmt.Errorf("nft contract adderess can not be empty cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	return nil
}

// ParamTable for project module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(nftContractAddress string, nftContractMinter string) Params {
	return Params{
		NftContractAddress: nftContractAddress,
		NftContractMinter:  nftContractAddress,
	}
}

// default project module parameters
func DefaultParams() Params {
	return Params{
		NftContractAddress: "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
		NftContractMinter:  "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{KeyNftContractAddress, &p.NftContractAddress, validateNftContractAddress},
		{KeyNftContractMinter, &p.NftContractMinter, validateNftContractAddress},
	}
}
