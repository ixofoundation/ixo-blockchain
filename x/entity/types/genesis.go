package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Entities: []Entity{},
		Params:   DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
