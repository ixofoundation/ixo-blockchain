package types

type GenesisState struct {
	BondDocs []MsgCreateBond `json:"bond_docs" yaml:"bond_docs"`
}

func NewGenesisState(bondDocs []MsgCreateBond) GenesisState {
	return GenesisState{
		BondDocs: bondDocs,
	}
}

//noinspection GoUnusedParameter
func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		BondDocs: nil,
	}
}
