package types

func NewGenesisState(iids []IidDocument, meta []IidMetadata) *GenesisState {
	return &GenesisState{
		IidDocs: iids,
		IidMeta: meta,
	}
}

//func ValidateGenesis(data *GenesisState) error {
//	err := ValidateParams(data.Params)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		IidDocs: nil,
		IidMeta: nil,
	}
}
