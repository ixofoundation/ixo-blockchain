package v6

import (
	store "cosmossdk.io/store/types"
	"github.com/ixofoundation/ixo-blockchain/v5/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name for the Ixo v6 upgrade.
	UpgradeName = "v6"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
