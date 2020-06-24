package types

import "github.com/ixofoundation/ixo-blockchain/x/did"

const (
	ModuleName   = "bonddoc"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	BondKey = []byte{0x01}
)

func GetBondPrefixKey(did did.Did) []byte {
	return append(BondKey, []byte(did)...)
}
