package types

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const (
	ModuleName   = "did"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var DidKey = []byte{0x01}

func GetDidPrefixKey(did ixo.Did) []byte {
	return append(DidKey, []byte(did)...)
}
