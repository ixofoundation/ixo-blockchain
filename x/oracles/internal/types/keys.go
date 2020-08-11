package types

import "github.com/ixofoundation/ixo-blockchain/x/did"

const (
	ModuleName   = "oracles"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	OracleKey = []byte{0x00}
)

func GetOraclePrefixKey(did did.Did) []byte {
	return append(OracleKey, []byte(did)...)
}
