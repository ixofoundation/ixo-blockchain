package types

import (
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeSetTokenContractCodes = "SetTokenContractCodes"
)

var (
	_ govtypesv1.Content = &SetTokenContractCodes{}
)

func NewSetTokenContract(ixo1155Code uint64) SetTokenContractCodes {
	return SetTokenContractCodes{
		Ixo1155ContractCode: ixo1155Code,
	}
}

func (p *SetTokenContractCodes) GetDescription() string {
	return "Update token contract codes"
}

func (p *SetTokenContractCodes) GetTitle() string {
	return "Set token contract codes"
}

func (p *SetTokenContractCodes) ProposalRoute() string { return RouterKey }
func (p *SetTokenContractCodes) ProposalType() string  { return ProposalTypeSetTokenContractCodes }

// TODO add validation
func (p *SetTokenContractCodes) ValidateBasic() error { return nil }

func init() {
	govtypesv1.RegisterProposalType(ProposalTypeSetTokenContractCodes)
}
