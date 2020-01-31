package types

type GenesisState struct {
	Bonds   []Bond  `json:"bonds" yaml:"bonds"`
	Batches []Batch `json:"batches" yaml:"batches"`
}

func NewGenesisState(bonds []Bond, batches []Batch) GenesisState {
	return GenesisState{
		Bonds:   bonds,
		Batches: batches,
	}
}

//noinspection GoUnusedParameter
func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Bonds:   nil,
		Batches: nil,
	}
}
