package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:          DefaultParams(),
		Tokens:          []Token{},
		TokenProperties: []TokenProperties{},
	}
}
