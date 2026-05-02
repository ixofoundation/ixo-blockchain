package v7

import (
	store "cosmossdk.io/store/types"

	"github.com/ixofoundation/ixo-blockchain/v6/app/upgrades"
)

// UpgradeName is the on-chain identifier for the v7 multi-pool liquidstake
// upgrade. The chain governance proposal must use exactly this string.
const (
	UpgradeName = "v7"

	// LegacyPoolID is the pool_id assigned to the pre-v7 single-pool state
	// during migration. Holders of the existing "uzero" LST denom continue
	// to interact with this pool seamlessly after the upgrade.
	LegacyPoolID = "zero"
)

// Upgrade is registered in app.go alongside earlier versions. No new store
// keys are added — v7 reshapes the existing liquidstake KV store.
var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
