package types

type GenesisState struct {
	PaymentTemplates []PaymentTemplate `json:"payment_templates" yaml:"payment_templates"`
	PaymentContracts []PaymentContract `json:"payment_contracts" yaml:"payment_contracts"`
	Subscriptions    []Subscription    `json:"subscriptions" yaml:"subscriptions"`
}

// TODO Implement for proto.Message interface

func (g GenesisState) Reset() {
	panic("implement me")
}

func (g GenesisState) String() string {
	panic("implement me")
}

func (g GenesisState) ProtoMessage() {
	panic("implement me")
}

func NewGenesisState(templates []PaymentTemplate, contracts []PaymentContract,
	subscriptions []Subscription) GenesisState {
	return GenesisState{
		PaymentTemplates: templates,
		PaymentContracts: contracts,
		Subscriptions:    subscriptions,
	}
}

func ValidateGenesis(data GenesisState) error {
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

func DefaultGenesisState() GenesisState {
	return GenesisState{
		PaymentTemplates: nil,
		PaymentContracts: nil,
		Subscriptions:    nil,
	}
}
