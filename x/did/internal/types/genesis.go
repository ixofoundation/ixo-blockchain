package types

import "github.com/ixofoundation/ixo-blockchain/x/ixo"

type GenesisState struct {
	DidDocs []ixo.DidDoc `json:"did_docs" yaml:"did_docs"`
}

func NewGenesisState(didDocs []ixo.DidDoc) GenesisState {
	return GenesisState{
		DidDocs: didDocs,
	}
}

//noinspection GoUnusedParameter
func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		DidDocs: nil,
	}
}
