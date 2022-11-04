package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		TokenDocs: []TokenDoc{},
		Params:     DefaultParams(),
	}
}
