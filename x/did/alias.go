package did

import (
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
)

type (
	Keeper       = keeper.Keeper
	Did           = exported.Did
	DidCredential = types.DidCredential
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	IsValidDid    = types.IsValidDid

	// variable aliases
	ErrInvalidDid        = types.ErrInvalidDid
)