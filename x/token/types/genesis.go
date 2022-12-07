package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		TokenMinters: []TokenMinter{},
		Params:       DefaultParams(),
	}
}
