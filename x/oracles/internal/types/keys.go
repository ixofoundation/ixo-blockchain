package types

import "github.com/ixofoundation/ixo-cosmos/x/ixo"

const (
	ModuleName   = "oracles"
	StoreKey     = ModuleName
	RouterKey    = StoreKey
	QuerierRoute = RouterKey
)

var (
	OracleKey = []byte{0x00}
)

func GetOraclePrefixKey(did ixo.Did) []byte {
	return append(OracleKey, []byte(did)...)
}
