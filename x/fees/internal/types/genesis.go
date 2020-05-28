package types

type GenesisState struct {
	Params          Params           `json:"params" yaml:"params"`
	Fees            []Fee            `json:"fees" yaml:"fees"`
	FeeContracts    []FeeContract    `json:"fee_contracts" yaml:"fee_contracts"`
	Subscriptions   []Subscription   `json:"subscriptions" yaml:"subscriptions"`
	DiscountHolders []DiscountHolder `json:"discount_holders" yaml:"discount_holders"`
}

func NewGenesisState(params Params, fees []Fee, feeContracts []FeeContract,
	subscriptions []Subscription, discountHolders []DiscountHolder) GenesisState {
	return GenesisState{
		Params:          params,
		Fees:            fees,
		FeeContracts:    feeContracts,
		Subscriptions:   subscriptions,
		DiscountHolders: discountHolders,
	}
}

func ValidateGenesis(data GenesisState) error {
	// Validate params
	err := ValidateParams(data.Params)
	if err != nil {
		return err
	}

	// Validate fees
	for _, f := range data.Fees {
		if err := f.Validate(); err != nil {
			return err
		}
	}

	// Validate fee contracts
	for _, f := range data.FeeContracts {
		if err := f.Validate(); err != nil {
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
		Params:        DefaultParams(),
		Fees:          nil,
		FeeContracts:  nil,
		Subscriptions: nil,
	}
}
