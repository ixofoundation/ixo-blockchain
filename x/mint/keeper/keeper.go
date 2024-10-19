package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/v3/ixomath"
	"github.com/ixofoundation/ixo-blockchain/v3/ixoutils"
	"github.com/ixofoundation/ixo-blockchain/v3/x/mint/types"
)

// Keeper of the mint store.
type Keeper struct {
	storeKey            storetypes.StoreKey
	paramSpace          paramtypes.Subspace
	accountKeeper       types.AccountKeeper
	bankKeeper          types.BankKeeper
	communityPoolKeeper types.CommunityPoolKeeper
	hooks               types.MintHooks
	feeCollectorName    string
}

type invalidRatioError struct {
	ActualRatio ixomath.Dec
}

func (e invalidRatioError) Error() string {
	return fmt.Sprintf("mint allocation ratio (%s) is greater than 1", e.ActualRatio)
}

const emptyWeightedAddressReceiver = ""

// NewKeeper creates a new mint Keeper instance.
func NewKeeper(
	key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	ak types.AccountKeeper, bk types.BankKeeper, ck types.CommunityPoolKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:            key,
		paramSpace:          paramSpace,
		accountKeeper:       ak,
		bankKeeper:          bk,
		communityPoolKeeper: ck,
		feeCollectorName:    feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// Set the mint hooks.
func (k *Keeper) SetHooks(h types.MintHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set mint hooks twice")
	}

	k.hooks = h

	return k
}

// GetMinter gets the minter.
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	ixoutils.MustGet(ctx.KVStore(k.storeKey), types.MinterKey, &minter)
	return
}

// SetMinter sets the minter.
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	ixoutils.MustSet(store, types.MinterKey, &minter)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// DistributeMintedCoin implements distribution of a minted coin from mint to external modules.
func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {
	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	// allocate staking incentives into fee collector account to be moved to on next begin blocker by staking module account.
	stakingIncentivesAmount, err := k.distributeToModule(ctx, k.feeCollectorName, mintedCoin, proportions.Staking)
	if err != nil {
		return err
	}

	// allocate impact rewards to respective accounts from mint module account.
	impactRewardAmount, err := k.distributeImpactRewards(ctx, mintedCoin, proportions.ImpactRewards, params.WeightedImpactRewardsReceivers)
	if err != nil {
		return err
	}

	// subtract from original provision to ensure no coins left over after the allocations
	communityPoolAmount := mintedCoin.Amount.Sub(stakingIncentivesAmount).Sub(impactRewardAmount)
	err = k.communityPoolKeeper.FundCommunityPool(ctx, sdk.NewCoins(sdk.NewCoin(params.MintDenom, communityPoolAmount)), k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	// call an hook after the minting and distribution of new coins
	k.hooks.AfterDistributeMintedCoin(ctx)

	return err
}

// getLastReductionEpochNum returns last reduction epoch number.
func (k Keeper) getLastReductionEpochNum(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.LastReductionEpochKey)
	if b == nil {
		return 0
	}

	return int64(sdk.BigEndianToUint64(b))
}

// setLastReductionEpochNum set last reduction epoch number.
func (k Keeper) setLastReductionEpochNum(ctx sdk.Context, epochNum int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastReductionEpochKey, sdk.Uint64ToBigEndian(uint64(epochNum)))
}

// mintCoins implements an alias call to the underlying bank keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) mintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// distributeToModule distributes mintedCoin multiplied by proportion to the recepientModule
func (k Keeper) distributeToModule(ctx sdk.Context, recipientModule string, mintedCoin sdk.Coin, proportion ixomath.Dec) (ixomath.Int, error) {
	distributionCoin, err := getProportions(mintedCoin, proportion)
	if err != nil {
		return ixomath.Int{}, err
	}
	ctx.Logger().Info("distributeToModule", "module", types.ModuleName, "recipientModule", recipientModule, "distributionCoin", distributionCoin, "height", ctx.BlockHeight())
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, recipientModule, sdk.NewCoins(distributionCoin)); err != nil {
		return ixomath.Int{}, err
	}
	return distributionCoin.Amount, nil
}

// distributeImpactRewards distributes impact rewards from mint module account
// to the respective account receivers by weight (impactRewardsReceivers).
// If no impact reward receivers given, funds the community pool instead.
// Returns the total amount distributed from the mint module account.
// With respect to input parameters, errors occur when:
// - impactRewardsReceivers is greater than 1.
// - invalid address in impact rewards receivers.
// - the balance of mint module account is less than totalMintedCoin * impactRewardsProportion.
func (k Keeper) distributeImpactRewards(ctx sdk.Context, totalMintedCoin sdk.Coin, impactRewardsProportion ixomath.Dec, impactRewardsReceivers []types.WeightedAddress) (ixomath.Int, error) {
	impactRewardCoin, err := getProportions(totalMintedCoin, impactRewardsProportion)
	if err != nil {
		return ixomath.Int{}, err
	}

	// counter for total distributed amount, used instead of impactRewardCoin.Amount
	// to avoid rounding discrepancies.
	totalDistributedAmount := ixomath.ZeroInt()

	// If no impact rewards receivers provided, fund the community pool.
	if len(impactRewardsReceivers) == 0 {
		err = k.communityPoolKeeper.FundCommunityPool(ctx, sdk.NewCoins(impactRewardCoin), k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return ixomath.Int{}, err
		}
		totalDistributedAmount = impactRewardCoin.Amount
	} else {
		// allocate impact rewards to addresses by weight
		for _, w := range impactRewardsReceivers {
			impactPortionCoin, err := getProportions(impactRewardCoin, w.Weight)
			if err != nil {
				return ixomath.Int{}, err
			}
			impactRewardPortionCoins := sdk.NewCoins(impactPortionCoin)

			// fund community pool when rewards address is empty.
			if w.Address == emptyWeightedAddressReceiver {
				err := k.communityPoolKeeper.FundCommunityPool(ctx, impactRewardPortionCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
				if err != nil {
					return ixomath.Int{}, err
				}
			} else {
				// distribute impact rewards to the address
				impactRewardsAddr, err := sdk.AccAddressFromBech32(w.Address)
				if err != nil {
					return ixomath.Int{}, err
				}
				err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, impactRewardsAddr, impactRewardPortionCoins)
				if err != nil {
					return ixomath.Int{}, err
				}
			}
			totalDistributedAmount = totalDistributedAmount.Add(impactPortionCoin.Amount)
		}
	}

	return totalDistributedAmount, nil
}

// getProportions gets the balance of the `MintedDenom` from minted coins and returns coins according to the
// allocation ratio. Returns error if ratio is greater than 1.
func getProportions(mintedCoin sdk.Coin, ratio ixomath.Dec) (sdk.Coin, error) {
	if ratio.GT(ixomath.OneDec()) {
		return sdk.Coin{}, invalidRatioError{ratio}
	}
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToLegacyDec().Mul(ratio).TruncateInt()), nil
}
