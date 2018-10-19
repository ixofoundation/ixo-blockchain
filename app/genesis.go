package app

import (
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

//___________________________________________________________________________________
// State to Unmarshal
type GenesisState struct {
	Accounts []*types.GenesisAccount `json:"accounts"`
	FeeData  fees.GenesisState       `json:"fees"`
	Config   ixo.GenesisState        `json:"config"`
}