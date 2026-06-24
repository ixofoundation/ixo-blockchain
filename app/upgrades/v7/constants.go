package v7

import (
	store "cosmossdk.io/store/types"

	"github.com/ixofoundation/ixo-blockchain/v8/app/upgrades"
	namestypes "github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

// UpgradeName is the on-chain identifier for the v7 ("Opus") multi-pool
// liquidstake upgrade. The chain governance software-upgrade proposal must
// use exactly this string.
const (
	UpgradeName = "Opus"

	// LegacyPoolID is the pool_id assigned to the pre-v7 single-pool state
	// during migration. Holders of the existing "uzero" LST denom continue
	// to interact with this pool seamlessly after the upgrade.
	LegacyPoolID = "zero"
)

// Upgrade is registered in app.go alongside earlier versions. v7 reshapes the
// existing liquidstake KV store and introduces the new x/names KV store for
// the chain-level name service (IXO-1123).
var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{namestypes.StoreKey},
		Deleted: []string{},
	},
}
