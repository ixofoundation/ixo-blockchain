package types

import (
	fmt "fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyNftContractAddress         = []byte("NftContractAddress")
	ProposalTypeEntityParamChange = "EntityParamChange"
)

func validateNftContractAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}

	if len(addr) == 0 {
		return fmt.Errorf("nft contract adderess can not be empty cannot be empty")
	}

	// _, err := sdk.AccAddressFromBech32(addr)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// ParamTable for project module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(nftContractAddress string) Params {
	return Params{
		NftContractAddress: nftContractAddress,
	}
}

// default project module parameters
func DefaultParams() Params {
	return Params{
		NftContractAddress: "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	fmt.Println("================= Setting Parmas:", p.NftContractAddress)
	return paramstypes.ParamSetPairs{
		{KeyNftContractAddress, &p.NftContractAddress, validateNftContractAddress},
	}
}

func (p *Params) GetDescription() string {
	return "update entity paramaters"
}

func (p *Params) GetTitle() string {
	return "update entity paramaters"
}

func (sup *Params) ProposalRoute() string { return RouterKey }
func (sup *Params) ProposalType() string  { return ProposalTypeEntityParamChange }

func (sup *Params) ValidateBasic() error { return nil }

func init() {
	govtypes.RegisterProposalType(ProposalTypeEntityParamChange)
	govtypes.RegisterProposalTypeCodec(&Params{}, "ixofoundation/EntityParamChangeProposal")
}
