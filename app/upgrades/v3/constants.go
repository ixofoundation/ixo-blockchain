package v3

import (
	store "cosmossdk.io/store/types"

	"github.com/ixofoundation/ixo-blockchain/v5/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Ixo v3 upgrade.
const UpgradeName = "v3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
