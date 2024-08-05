package types

func NewGenesisState(iids []IidDocument, meta []IidMetadata) *GenesisState {
	return &GenesisState{
		IidDocs: iids,
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
	}
}
