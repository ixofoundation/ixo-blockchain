package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func NewGenesisState(templates []PaymentTemplate, contracts []PaymentContract,
	subscriptions []Subscription) *GenesisState {
	return &GenesisState{
		PaymentTemplates: templates,
		PaymentContracts: contracts,
		Subscriptions:    subscriptions,
	}
}

func ValidateGenesis(data *GenesisState) error {
	// Validate payment templates
	for _, pt := range data.PaymentTemplates {
		if err := pt.Validate(); err != nil {
			return err
		}
	}

	// Validate payment contracts
	for _, pc := range data.PaymentContracts {
		if err := pc.Validate(); err != nil {
			return err
		}
	}

	// Validate subscriptions
	for _, s := range data.Subscriptions {
		if err := s.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		PaymentTemplates: nil,
		PaymentContracts: nil,
		Subscriptions:    nil,
	}
}

var _ types.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, subscription := range data.Subscriptions {
		err := subscription.UnpackInterfaces(unpacker)
		if err != nil {
			return nil
		}
	}

	return nil
}
