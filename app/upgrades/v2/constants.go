package v2

import (
	"github.com/ixofoundation/ixo-blockchain/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"
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
