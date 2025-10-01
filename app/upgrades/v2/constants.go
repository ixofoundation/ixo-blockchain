package v2

import (
	"github.com/ixofoundation/ixo-blockchain/v6/app/upgrades"

	store "cosmossdk.io/store/types"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
)

// UpgradeName defines the on-chain upgrade name for the Ixo v2 upgrade.
const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			icqtypes.StoreKey,
			packetforwardtypes.StoreKey,
		},
		Deleted: []string{},
	},
}
