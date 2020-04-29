package types

type GenesisState struct {
	Oracles Oracles `json:"oracles" yaml:"oracles"`
}

func NewGenesisState(oracles Oracles) GenesisState {
	return GenesisState{
		Oracles: oracles,
	}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Oracles: nil,
	}
}
