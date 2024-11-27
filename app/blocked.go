package app

import (
	"strings"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	liquidstaketypes "github.com/ixofoundation/ixo-blockchain/v4/x/liquidstake/types"
)

var (
	// Module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName:       true,
		liquidstaketypes.ModuleName: true,
	}
)

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *IxoApp) BlockedAddresses() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	// We can add any addresses we want to block below
	addresses := []string{}
	for _, addr := range addresses {
		blockedAddrs[addr] = true
		blockedAddrs[strings.ToLower(addr)] = true
	}

	return blockedAddrs
}
