package types

func NewGenesisState(iids []IidDocument, meta []IidMetadata) *GenesisState {
	return &GenesisState{
		IidDocs: iids,
	}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		IidDocs: nil,
	}
}
