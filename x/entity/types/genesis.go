package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Entities: []Entity{},
		Params:     DefaultParams(),
	}
}
