package types

import (
	"github.com/ixofoundation/ixo-cosmos/types"
)

const (
	ModuleName   = "fiat"
	StoreKey     = ModuleName
	RouterKey    = StoreKey
	QuerierRoute = RouterKey
)

var (
	PegHashKey = []byte{0x01}
)

func FiatPegHashStoreKey(fiatPegHash types.PegHash) []byte {
	return append(PegHashKey, fiatPegHash.Bytes()...)
}
