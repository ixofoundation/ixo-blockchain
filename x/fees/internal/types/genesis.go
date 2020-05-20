package types

type GenesisState struct {
	Params       Params        `json:"params" yaml:"params"`
	Fees         []Fee         `json:"fees" yaml:"fees"`
	FeeContracts []FeeContract `json:"fee_contracts" yaml:"fee_contracts"`
}

func NewGenesisState(params Params, fees []Fee, feeContracts []FeeContract) GenesisState {
	return GenesisState{
		Params:       params,
		Fees:         fees,
		FeeContracts: feeContracts,
	}
}

func ValidateGenesis(data GenesisState) error {
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

	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:       DefaultParams(),
		Fees:         nil,
		FeeContracts: nil,
	}
}
