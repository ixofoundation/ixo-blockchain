package app

import (
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
)

//___________________________________________________________________________________
// State to Unmarshal
type GenesisState struct {
	Accounts []*types.GenesisAccount `json:"accounts"`
	FeeData  fees.GenesisState       `json:"fees"`
}
