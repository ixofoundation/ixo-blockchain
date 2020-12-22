package types

import "github.com/ixofoundation/ixo-blockchain/x/did/exported"

type GenesisState struct {
	DidDocs []exported.DidDoc `json:"did_docs" yaml:"did_docs"`
}

// TODO Implement for proto.Message interface

func (g GenesisState) Reset() {
	panic("implement me")
}

func (g GenesisState) String() string {
	panic("implement me")
}

func (g GenesisState) ProtoMessage() {
	panic("implement me")
}

func NewGenesisState(didDocs []exported.DidDoc) GenesisState {
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
