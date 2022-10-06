package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		EntityDocs: []EntityDoc{},
	}
}
