package apptesting

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateRandomAccounts returns numAccts deterministic-but-unique addresses
// derived from freshly generated ed25519 pubkeys. Useful when a test needs
// "some address that hasn't seen any state yet".
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		addrs[i] = sdk.AccAddress(pk.Address())
	}
	return addrs
}

// RandomAccountAddress returns a single random ixo account address.
func RandomAccountAddress() sdk.AccAddress {
	return CreateRandomAccounts(1)[0]
}
