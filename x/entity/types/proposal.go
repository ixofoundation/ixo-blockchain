package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	ProposalTypeInitializeNftContract                  = "InitializeNftContract"
	_                                 govtypes.Content = &InitializeNftContract{}
)

func NewInitializeNftContract(nftContractCodeId uint64, nftminteraddress string) InitializeNftContract {
	return InitializeNftContract{
		NftContractCodeId: nftContractCodeId,
		NftMinterAddress:  nftminteraddress,
	}
}

func (p *InitializeNftContract) GetDescription() string {
	return "update entity paramaters"
}

func (p *InitializeNftContract) GetTitle() string {
	return "update entity paramaters"
}

func (sup *InitializeNftContract) ProposalRoute() string { return RouterKey }
func (sup *InitializeNftContract) ProposalType() string  { return ProposalTypeInitializeNftContract }

func (sup *InitializeNftContract) ValidateBasic() error { return nil }

func init() {
	govtypes.RegisterProposalType(ProposalTypeInitializeNftContract)
	govtypes.RegisterProposalTypeCodec(&InitializeNftContract{}, "entity.InitializeNftContract")
}
