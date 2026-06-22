package v8

import (
	store "cosmossdk.io/store/types"

	"github.com/ixofoundation/ixo-blockchain/v7/app/upgrades"
)

// UpgradeName is the on-chain identifier for the v8 emergency security upgrade
// that disables the x/bonds module and closes the authz.MsgExec recursion gap
// in the IID ante handler, following the 2026-06-20 incident.
const UpgradeName = "v8"

// Upgrade is registered in app.go alongside earlier versions. v8 introduces no
// new KV stores and deletes none: the bonds module is disabled behaviourally in
// the binary (disabled msg server + no-op EndBlocker + ante guard), and its
// existing store is retained so genesis/state/queries continue to load.
var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
