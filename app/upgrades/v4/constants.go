package v4

import (
	store "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	cosmosminttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	"github.com/ixofoundation/ixo-blockchain/v3/app/upgrades"
	epochstypes "github.com/ixofoundation/ixo-blockchain/v3/x/epochs/types"
	minttypes "github.com/ixofoundation/ixo-blockchain/v3/x/mint/types"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v3/x/smart-account/types"
)

// UpgradeName defines the on-chain upgrade name for the Ixo v4 upgrade.
const (
	UpgradeName = "Dominia"

	// BlockMaxBytes is the max bytes for a block, 10mb (current 22020096)
	BlockMaxBytes = int64(10000000)
	// BlockMaxGas is the max gas allowed in a block (current 200000000)
	BlockMaxGas = int64(300000000)

	// Normal proposal deposit is 10k ixo, make expedited proposal deposit 3x
	ExpeditedProposalDeposit = 30000000000
	MinInitialDepositRatio   = "0.100000000000000000"

	// MaximumUnauthenticatedGas for smart account transactions to verify the fee payer
	MaximumUnauthenticatedGas = uint64(250_000)
	// IsSmartAccountActive is used for the smart account circuit breaker, smartaccounts are activated for v4
	IsSmartAccountActive = false
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			smartaccounttypes.StoreKey,
			minttypes.StoreKey,
			epochstypes.StoreKey,
			ibchooks.StoreKey,
			// Add circuittypes as per 0.47 to 0.50 upgrade handler
			// https://github.com/cosmos/cosmos-sdk/blob/b7d9d4c8a9b6b8b61716d2023982d29bdc9839a6/simapp/upgrades.go#L21
			circuittypes.ModuleName,
			// v47 modules
			crisistypes.ModuleName,
			consensustypes.ModuleName,
		},
		Deleted: []string{
			cosmosminttypes.StoreKey,
			"intertx", // uninstalled module
		},
	},
}
