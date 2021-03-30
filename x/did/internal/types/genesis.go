package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

func NewGenesisState(dd []exported.DidDoc) *GenesisState {
	didDocs := make([]*types.Any, len(dd))

	for i, diddoc := range dd {
		msg, ok := diddoc.(proto.Message)
		if !ok {
			panic(fmt.Errorf("cannot proto marshal %T", diddoc))
		}
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			didDocs[i] = any
		}
	}

	return &GenesisState{
		Diddocs: didDocs,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

// ValidateGenesis performs no inspection
func ValidateGenesis(data GenesisState) error {
	return nil
}