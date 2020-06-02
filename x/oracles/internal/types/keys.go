package types

import "github.com/ixofoundation/ixo-blockchain/x/ixo"

const (
	ModuleName   = "oracles"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	OracleKey = []byte{0x00}
)

func GetOraclePrefixKey(did ixo.Did) []byte {
	return append(OracleKey, []byte(did)...)
}
