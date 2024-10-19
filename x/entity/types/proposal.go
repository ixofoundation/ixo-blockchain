package types

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	types "github.com/ixofoundation/ixo-blockchain/v4/x/iid/types"
)

const (
	ProposalTypeInitializeNftContract = "InitializeNftContract"
)

var (
	_ govtypesv1.Content = &InitializeNftContract{}
)

func NftModuleAddress() string { return fmt.Sprintf("%s-minter", types.ModuleName) }

func NewInitializeNftContract(nftContractCodeId uint64, nftminteraddress string) InitializeNftContract {
	return InitializeNftContract{
		NftContractCodeId: nftContractCodeId,
		NftMinterAddress:  nftminteraddress,
	}
}

func (p *InitializeNftContract) GetDescription() string {
	return "Initialize new NFT contract for entity module"
}

func (p *InitializeNftContract) GetTitle() string {
	return "Update entity parameters"
}

func (p *InitializeNftContract) ProposalRoute() string { return RouterKey }
func (p *InitializeNftContract) ProposalType() string  { return ProposalTypeInitializeNftContract }

func (p *InitializeNftContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(p.NftMinterAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}
	return nil
}

func init() {
	govtypesv1.RegisterProposalType(ProposalTypeInitializeNftContract)
}
