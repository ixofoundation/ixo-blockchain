package types

const (
	// ModuleName is the module name.
	ModuleName = "mint"

	// StoreKey is the default store key for mint.
	StoreKey = ModuleName + "-store"
)

var (
	// MinterKey is the key to use for the keeper store at which
	// the Minter and its EpochProvisions are stored.
	MinterKey = []byte{0x00}

	// LastReductionEpochKey is the key to use for the keeper store
	// for storing the last epoch at which reduction occurred.
	LastReductionEpochKey = []byte{0x03}

	// Parameter store keys.
	KeyMintDenom                            = []byte("MintDenom")
	KeyGenesisEpochProvisions               = []byte("GenesisEpochProvisions")
	KeyEpochIdentifier                      = []byte("EpochIdentifier")
	KeyReductionPeriodInEpochs              = []byte("ReductionPeriodInEpochs")
	KeyReductionFactor                      = []byte("ReductionFactor")
	KeyPoolAllocationRatio                  = []byte("PoolAllocationRatio")
	KeyImpactRewardsReceivers               = []byte("ImpactRewardsReceivers")
	KeyMintingRewardsDistributionStartEpoch = []byte("MintingRewardsDistributionStartEpoch")
)
