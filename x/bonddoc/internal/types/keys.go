package types

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const (
	ModuleName   = "bonddoc"
	StoreKey     = ModuleName
	RouterKey    = StoreKey
	QuerierRoute = RouterKey
)

var (
	BondKey = []byte{0x01}
)

func GetBondPrefixKey(did ixo.Did) []byte {
	return append(BondKey, []byte(did)...)
}
