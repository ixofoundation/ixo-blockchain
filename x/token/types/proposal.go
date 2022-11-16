package types

import (
	fmt "fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	types "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var (
	ProposalTypeInitializeNftContract                  = "InitializeTokenContract"
	_                                 govtypes.Content = &InitializeTokenContract{}
)

func NftModuleAddress() string { return fmt.Sprintf("%s-minter", types.ModuleName) }

func NewInitializeNftContract(nftContractCodeId uint64, nftminteraddress string) InitializeTokenContract {
	return InitializeTokenContract{
		NftContractCodeId: nftContractCodeId,
		NftMinterAddress:  nftminteraddress,
	}
}

func (p *InitializeTokenContract) GetDescription() string {
	return "update token paramaters"
}

func (p *InitializeTokenContract) GetTitle() string {
	return "update token paramaters"
}

func (sup *InitializeTokenContract) ProposalRoute() string { return RouterKey }
func (sup *InitializeTokenContract) ProposalType() string  { return ProposalTypeInitializeNftContract }

func (sup *InitializeTokenContract) ValidateBasic() error { return nil }

func init() {
	govtypes.RegisterProposalType(ProposalTypeInitializeNftContract)
	govtypes.RegisterProposalTypeCodec(&InitializeTokenContract{}, "token.ixo.token.v1beta1.InitializeTokenContract")
}
