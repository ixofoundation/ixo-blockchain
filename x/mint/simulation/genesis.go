package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/ixofoundation/ixo-blockchain/v4/ixomath"
	"github.com/ixofoundation/ixo-blockchain/v4/x/mint/types"
)

// Simulation parameter constants.
const (
	epochProvisionsKey         = "genesis_epoch_provisions"
	reductionFactorKey         = "reduction_factor"
	reductionPeriodInEpochsKey = "reduction_period_in_epochs"

	mintingRewardsDistributionStartEpochKey = "minting_rewards_distribution_start_epoch"

	epochIdentifier = "day"
	maxInt64        = int(^uint(0) >> 1)
)

var (
	distributionProportions = types.DistributionProportions{
		Staking:       ixomath.NewDecWithPrec(85, 2),
		ImpactRewards: ixomath.NewDecWithPrec(10, 2),
		CommunityPool: ixomath.NewDecWithPrec(0o5, 2),
	}
	weightedImpactRewardReceivers = []types.WeightedAddress{
		{
			Address: "ixo1n8yrmeatsk74dw0zs95ess9sgzptd6thgjgcj2",
			Weight:  ixomath.NewDecWithPrec(90, 2),
		},
		{
			Address: "ixo12am7v5xgjh72c7xujreyvtncqwue3w0v6ud3r4",
			Weight:  ixomath.NewDecWithPrec(995, 3).Quo(ixomath.NewDec(10)), // 0.0995
		},
		{
			Address: "ixo13dy867pyn8jda82vnshy7jjjv42n69k7497jrh",
			Weight:  ixomath.NewDecWithPrec(5, 1).Quo(ixomath.NewDec(1000)), // 0.0005
		},
	}
)

// RandomizedGenState generates a random GenesisState for mint.
func RandomizedGenState(simState *module.SimulationState) {
	var epochProvisions ixomath.Dec
	simState.AppParams.GetOrGenerate(
		epochProvisionsKey, &epochProvisions, simState.Rand,
		func(r *rand.Rand) { epochProvisions = genEpochProvisions(r) },
	)

	var reductionFactor ixomath.Dec
	simState.AppParams.GetOrGenerate(
		reductionFactorKey, &reductionFactor, simState.Rand,
		func(r *rand.Rand) { reductionFactor = genReductionFactor(r) },
	)

	var reductionPeriodInEpochs int64
	simState.AppParams.GetOrGenerate(
		reductionPeriodInEpochsKey, &reductionPeriodInEpochs, simState.Rand,
		func(r *rand.Rand) { reductionPeriodInEpochs = genReductionPeriodInEpochs(r) },
	)

	var mintintRewardsDistributionStartEpoch int64
	simState.AppParams.GetOrGenerate(
		mintingRewardsDistributionStartEpochKey, &mintintRewardsDistributionStartEpoch, simState.Rand,
		func(r *rand.Rand) { mintintRewardsDistributionStartEpoch = genMintintRewardsDistributionStartEpoch(r) },
	)

	reductionStartedEpoch := genReductionStartedEpoch(simState.Rand)

	mintDenom := sdk.DefaultBondDenom
	params := types.NewParams(
		mintDenom,
		epochProvisions,
		epochIdentifier,
		reductionFactor,
		reductionPeriodInEpochs,
		distributionProportions,
		weightedImpactRewardReceivers,
		mintintRewardsDistributionStartEpoch)

	minter := types.NewMinter(epochProvisions)

	mintGenesis := types.NewGenesisState(minter, params, reductionStartedEpoch)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}

func genEpochProvisions(r *rand.Rand) ixomath.Dec {
	return ixomath.NewDec(int64(r.Intn(maxInt64)))
}

func genReductionFactor(r *rand.Rand) ixomath.Dec {
	return ixomath.NewDecWithPrec(int64(r.Intn(10)), 1)
}

func genReductionPeriodInEpochs(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}

func genMintintRewardsDistributionStartEpoch(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}

func genReductionStartedEpoch(r *rand.Rand) int64 {
	return int64(r.Intn(maxInt64))
}
