package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	types "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var (
	ProposalTypeInitializeNftContract                  = "InitializeNftContract"
	_                                 govtypes.Content = &InitializeNftContract{}
)

func NftModuleAddress() string { return fmt.Sprintf("%s-minter", types.ModuleName) }

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

func (sup *InitializeNftContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(sup.NftMinterAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}
	return nil
}

func init() {
	govtypes.RegisterProposalType(ProposalTypeInitializeNftContract)
	govtypes.RegisterProposalTypeCodec(&InitializeNftContract{}, "entity.ixo.entity.v1beta1.InitializeNftContract")
}
