package v4

import (
	"cosmossdk.io/math"
	store "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	cosmosminttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	"github.com/ixofoundation/ixo-blockchain/v3/app/upgrades"
	epochstypes "github.com/ixofoundation/ixo-blockchain/v3/x/epochs/types"
	liquidstaketypes "github.com/ixofoundation/ixo-blockchain/v3/x/liquidstake/types"
	minttypes "github.com/ixofoundation/ixo-blockchain/v3/x/mint/types"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v3/x/smart-account/types"
)

// TODO: need details below!!

const (
	// UpgradeName defines the on-chain upgrade name for the Ixo v4 upgrade.
	UpgradeName = "Dominia"

	// CONSENSUS PARAMS
	// -------------------------------------------------
	// BlockMaxBytes is the max bytes for a block, 30mb (current 22020096)
	BlockMaxBytes = int64(30000000)
	// BlockMaxGas is the max gas allowed in a block (current 200000000)
	BlockMaxGas = int64(400000000)

	// GOV PARAMS
	// -------------------------------------------------
	// Normal proposal deposit is 10k ixo, make expedited proposal deposit 3x
	ExpeditedProposalDeposit = 30000000000
	MinInitialDepositRatio   = "0.100000000000000000"

	// SMART ACCOUNT PARAMS
	// -------------------------------------------------
	// MaximumUnauthenticatedGas for smart account transactions to verify the fee payer
	MaximumUnauthenticatedGas = uint64(250_000)
	// IsSmartAccountActive is used for the smart account circuit breaker, smartaccounts are activated for v4
	IsSmartAccountActive = true
	// CircuitBreakerController is a DAODAO address, used only to deactivate the smart account module
	// https://daodao.zone/dao/osmo1wn58hxkv0869ua7qmz3gvek3sz773l89a778fjqvenl6anwuhgnq6ks7kl/home
	CircuitBreakerController = "ixo1kqmtxkggcqa9u34lnr6shy0euvclgatw4f9zz5"
)

var (
	// LIQUID STAKE PARAMS
	// -------------------------------------------------
	// total weights must be 10000
	WhitelistedValidators = []liquidstaketypes.WhitelistedValidator{
		{
			ValidatorAddress: "ixovaloper1n8yrmeatsk74dw0zs95ess9sgzptd6thzncf20",
			TargetWeight:     math.NewIntFromUint64(5000),
		},
		{
			ValidatorAddress: "ixovaloper1n8yrmeatsk74dw0zs95ess9sgzptd6thzncf20",
			TargetWeight:     math.NewIntFromUint64(5000),
		},
	}
	// LSMUnstakeFeeRate is the Unstake Fee Rate. (note the fee amount stays as staked amount in proxy account, it is not
	// instaked and transferred to the FeeAccountAddress)
	LSMUnstakeFeeRate = math.LegacyNewDecWithPrec(3333, 4) // "0.333300000000000000"
	// LSMAutocompoundFeeRate is the fee rate for auto redelegating the stake rewards.
	LSMAutocompoundFeeRate = math.LegacyNewDecWithPrec(3333, 4) // "0.333300000000000000"
	// LSMWhitelistAdminAddress is the address of the whitelist admin, who is allowed to add/remove whitelisted validators,
	// pause/unpause the liquid stake module, and set the weighted rewards receivers.
	LSMWhitelistAdminAddress = "ixo1kqmtxkggcqa9u34lnr6shy0euvclgatw4f9zz5"
	// LSMWeightedRewardsReceivers is the list of weighted rewards receivers who will recieve the staking rewards based
	// on their weights.
	LSMWeightedRewardsReceivers = []liquidstaketypes.WeightedAddress{}
	// LSMFeeAccountAddress is the address of the fee account, which will receive the autocompound fees.
	LSMFeeAccountAddress = "ixo1kqmtxkggcqa9u34lnr6shy0euvclgatw4f9zz5"

	// STAKING PARAMS
	// -------------------------------------------------
	// The ValidatorBondFactor dictates the cap on the liquid shares
	// for a validator - determined as a multiple to their validator bond
	// (e.g. ValidatorBondShares = 1000, BondFactor = 250 -> LiquidSharesCap: 250,000)
	// ValidatorBondFactor of -1 indicates that it's disabled
	ValidatorBondFactor = math.LegacyNewDecFromInt(math.NewInt(-1))
	// GlobalLiquidStakingCap represents a cap on the portion of stake that
	// comes from liquid staking providers for a specific validator
	ValidatorLiquidStakingCap = math.LegacyOneDec() // 100%
	// GlobalLiquidStakingCap represents the percentage cap on
	// the portion of a chain's total stake can be liquid
	GlobalLiquidStakingCap = math.LegacyOneDec() // 100%
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
			liquidstaketypes.StoreKey,
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
