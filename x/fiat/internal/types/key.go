package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName   = "fiat"
	StoreKey     = ModuleName
	RouterKey    = StoreKey
	QuerierRoute = RouterKey
)

var (
	FiatAccountKey = []byte{0x04}
	FiatPegHashKey = []byte{0x05}
)

func FiatAccountStoreKey(address sdk.AccAddress) []byte {
	return append(FiatAccountKey, address...)
}
