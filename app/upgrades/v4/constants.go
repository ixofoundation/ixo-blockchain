package v4

import (
	"cosmossdk.io/math"
	store "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	cosmosminttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	"github.com/ixofoundation/ixo-blockchain/v4/app/upgrades"
	epochstypes "github.com/ixofoundation/ixo-blockchain/v4/x/epochs/types"
	liquidstaketypes "github.com/ixofoundation/ixo-blockchain/v4/x/liquidstake/types"
	minttypes "github.com/ixofoundation/ixo-blockchain/v4/x/mint/types"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v4/x/smart-account/types"
)

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
	// Set to the Impact Hub DAO Security Group
	CircuitBreakerController = "ixo1rrf8ydfh23y59vdv2aslvtfsua0vp84yyrg6gwflqht5dpexhxns45j7cd"
)

var (
	// LIQUID STAKE PARAMS
	// -------------------------------------------------
	// total weights must be 10000
	WhitelistedValidators = []liquidstaketypes.WhitelistedValidator{
		// Internet of Impacts Launchpad Validator to start with
		{
			ValidatorAddress: "ixovaloper1kr8v9qt46ysltgmzcrtgyw5ckke83up673u2lu",
			TargetWeight:     math.NewIntFromUint64(10000),
		},
	}
	// LSMUnstakeFeeRate is the Unstake Fee Rate. (note the fee amount stays as staked amount in proxy account, it is not
	// instaked and transferred to the FeeAccountAddress)
	LSMUnstakeFeeRate = math.LegacyZeroDec()
	// LSMAutocompoundFeeRate is the fee rate for auto redelegating the stake rewards.
	LSMAutocompoundFeeRate = math.LegacyZeroDec()
	// LSMWhitelistAdminAddress is the address of the whitelist admin, who is allowed to add/remove whitelisted validators,
	// pause/unpause the liquid stake module, and set the weighted rewards receivers.
	// Set to the ZERO DAO Members Group
	LSMWhitelistAdminAddress = "ixo1dwaypqeva5j5p8a0dux9n35lp0aspepnfy9zj45fc9s3wmpvcz4sk6duea"
	// LSMWeightedRewardsReceivers is the list of weighted rewards receivers who will recieve the staking rewards based
	// on their weights.
	LSMWeightedRewardsReceivers = []liquidstaketypes.WeightedAddress{}
	// LSMFeeAccountAddress is the address of the fee account, which will receive the autocompound fees.
	// Set to the ZERO DAO "LS Fees" entity account
	LSMFeeAccountAddress = "ixo1qzngrk3hnpytp9apt7442u3gyr0fyv9y4wvy6u"

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
