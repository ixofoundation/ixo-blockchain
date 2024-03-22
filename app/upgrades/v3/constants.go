package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/ixofoundation/ixo-blockchain/v2/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Ixo v3 upgrade.
const UpgradeName = "v3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
