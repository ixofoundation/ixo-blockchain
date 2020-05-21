package types

type GenesisState struct {
	Params                 Params         `json:"params" yaml:"params"`
	Fees                   []Fee          `json:"fees" yaml:"fees"`
	FeeContracts           []FeeContract  `json:"fee_contracts" yaml:"fee_contracts"`
	Subscriptions          []Subscription `json:"subscriptions" yaml:"subscriptions"`
	StartingFeeId          uint64         `json:"starting_fee_id" yaml:"starting_fee_id"`
	StartingFeeContractId  uint64         `json:"starting_fee_contract_id" yaml:"starting_fee_contract_id"`
	StartingSubscriptionId uint64         `json:"starting_subscription_id" yaml:"starting_subscription_id"`
}

func NewGenesisState(params Params, fees []Fee, feeContracts []FeeContract,
	subscriptions []Subscription, startingFeeID, startingFeeContractID,
	startingSubscriptionID uint64) GenesisState {
	return GenesisState{
		Params:                 params,
		Fees:                   fees,
		FeeContracts:           feeContracts,
		Subscriptions:          subscriptions,
		StartingFeeId:          startingFeeID,
		StartingFeeContractId:  startingFeeContractID,
		StartingSubscriptionId: startingSubscriptionID,
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
		Params:                 DefaultParams(),
		Fees:                   nil,
		FeeContracts:           nil,
		Subscriptions:          nil,
		StartingFeeId:          1,
		StartingFeeContractId:  1,
		StartingSubscriptionId: 1,
	}
}
