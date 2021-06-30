package types

func NewGenesisState(bonds []Bond, batches []Batch, params Params) *GenesisState {
	return &GenesisState{
		Bonds:   bonds,
		Batches: batches,
		Params:  params,
	}
}

func ValidateGenesis(data *GenesisState) error {
	err := ValidateParams(data.Params)
	if err != nil {
		return err
	}

	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Bonds:   nil,
		Batches: nil,
		Params:  DefaultParams(),
	}
}
