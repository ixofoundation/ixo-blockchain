package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:          DefaultParams(),
		Tokens:          []Token{},
		TokenProperties: []TokenProperties{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
