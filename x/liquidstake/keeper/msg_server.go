package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the liquidstake MsgServer
// interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// authorisedByGovOrPoolAdmin returns nil if the supplied authority equals
// either the chain governance authority or the pool's current admin. Used
// by the per-pool admin operations.
func (s msgServer) authorisedByGovOrPoolAdmin(authority string, p types.Pool) error {
	if authority == s.Keeper.Authority() || authority == p.WhitelistAdminAddress {
		return nil
	}
	return errors.Wrapf(
		sdkerrors.ErrorInvalidSigner,
		"invalid authority; expected governance %s or pool admin %s, got %s",
		s.Keeper.Authority(), p.WhitelistAdminAddress, authority,
	)
}

// authorisedByGov returns nil if the supplied authority equals the chain
// governance authority. Used for module-wide and pool-creation operations.
func (s msgServer) authorisedByGov(authority string) error {
	if authority == s.Keeper.Authority() {
		return nil
	}
	return errors.Wrapf(
		sdkerrors.ErrorInvalidSigner,
		"invalid authority; expected governance %s, got %s",
		s.Keeper.Authority(), authority,
	)
}

// ----------------------------------------------------------------------------
// LiquidStake
// ----------------------------------------------------------------------------

func (s msgServer) LiquidStake(goCtx context.Context, msg *types.MsgLiquidStake) (*types.MsgLiquidStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}

	mintedAmount, err := s.Keeper.LiquidStake(ctx, pool, msg.GetDelegator(), msg.Amount)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.LiquidStakeEvent{
		PoolId:             pool.PoolId,
		Delegator:          msg.DelegatorAddress,
		LiquidAmount:       msg.Amount,
		StkIxoMintedAmount: sdk.NewCoin(pool.LiquidBondDenom, mintedAmount),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit liquid stake event")
	}
	return &types.MsgLiquidStakeResponse{}, nil
}

// ----------------------------------------------------------------------------
// LiquidUnstake
// ----------------------------------------------------------------------------

func (s msgServer) LiquidUnstake(goCtx context.Context, msg *types.MsgLiquidUnstake) (*types.MsgLiquidUnstakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}

	completionTime, unbondingAmount, _, immediatelyReturned, err := s.Keeper.LiquidUnstake(ctx, pool, msg.GetDelegator(), msg.Amount)
	if err != nil {
		return nil, err
	}

	bondDenom, err := s.Keeper.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return nil, err
	}
	if err := ctx.EventManager().EmitTypedEvent(&types.LiquidUnstakeEvent{
		PoolId:          pool.PoolId,
		Delegator:       msg.DelegatorAddress,
		UnstakeAmount:   msg.Amount,
		UnbondingAmount: sdk.NewCoin(bondDenom, unbondingAmount),
		UnbondedAmount:  sdk.NewCoin(bondDenom, immediatelyReturned),
		CompletionTime:  completionTime,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit liquid unstake event")
	}
	return &types.MsgLiquidUnstakeResponse{CompletionTime: completionTime}, nil
}

// ----------------------------------------------------------------------------
// CreatePool (governance only)
// ----------------------------------------------------------------------------

func (s msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := s.authorisedByGov(msg.Authority); err != nil {
		return nil, err
	}

	pool, err := s.Keeper.registerPool(ctx, msg.PoolId, msg.LiquidBondDenom, msg.InitialAdminAddress, msg.InitialFeeAccountAddress)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.PoolCreatedEvent{
		PoolId:    pool.PoolId,
		Pool:      &pool,
		Authority: msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit pool created event")
	}
	return &types.MsgCreatePoolResponse{ProxyAccountAddress: pool.ProxyAccountAddress}, nil
}

// ----------------------------------------------------------------------------
// UpdateModuleParams (governance only)
// ----------------------------------------------------------------------------

func (s msgServer) UpdateModuleParams(goCtx context.Context, msg *types.MsgUpdateModuleParams) (*types.MsgUpdateModuleParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := s.authorisedByGov(msg.Authority); err != nil {
		return nil, err
	}
	if err := s.Keeper.SetModuleParams(ctx, msg.ModuleParams); err != nil {
		return nil, err
	}
	if err := ctx.EventManager().EmitTypedEvent(&types.ModuleParamsUpdatedEvent{
		ModuleParams: &msg.ModuleParams,
		Authority:    msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit module params updated event")
	}
	return &types.MsgUpdateModuleParamsResponse{}, nil
}

// ----------------------------------------------------------------------------
// UpdatePool (governance or pool admin)
// ----------------------------------------------------------------------------

func (s msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}
	if err := s.authorisedByGovOrPoolAdmin(msg.Authority, pool); err != nil {
		return nil, err
	}

	pool.UnstakeFeeRate = msg.UnstakeFeeRate
	pool.AutocompoundFeeRate = msg.AutocompoundFeeRate
	pool.FeeAccountAddress = msg.FeeAccountAddress
	pool.WhitelistAdminAddress = msg.WhitelistAdminAddress

	if err := pool.Validate(); err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	s.Keeper.SetPool(ctx, pool)

	if err := ctx.EventManager().EmitTypedEvent(&types.PoolUpdatedEvent{
		PoolId:    pool.PoolId,
		Pool:      &pool,
		Authority: msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit pool updated event")
	}
	return &types.MsgUpdatePoolResponse{}, nil
}

// ----------------------------------------------------------------------------
// UpdateWhitelistedValidators (governance or pool admin)
// ----------------------------------------------------------------------------

func (s msgServer) UpdateWhitelistedValidators(goCtx context.Context, msg *types.MsgUpdateWhitelistedValidators) (*types.MsgUpdateWhitelistedValidatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}
	if err := s.authorisedByGovOrPoolAdmin(msg.Authority, pool); err != nil {
		return nil, err
	}

	// Validate every entry parses + has a positive non-nil weight + uniqueness.
	if err := types.ValidateWhitelistedValidators(msg.WhitelistedValidators); err != nil {
		return nil, errors.Wrap(types.ErrWhitelistedValidatorsList, err.Error())
	}

	// Per-validator: must exist on chain in Bonded status.
	totalWeight := math.NewInt(0)
	for _, val := range msg.WhitelistedValidators {
		totalWeight = totalWeight.Add(val.TargetWeight)
		valAddr := val.GetValidatorAddress()
		fullVal, err := s.Keeper.stakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			return nil, errors.Wrapf(types.ErrWhitelistedValidatorsList, "validator not found: %s", valAddr)
		}
		if fullVal.Status != stakingtypes.Bonded {
			return nil, errors.Wrapf(
				types.ErrWhitelistedValidatorsList,
				"validator status %s: expected %s; got %s",
				valAddr, stakingtypes.Bonded.String(), fullVal.Status.String(),
			)
		}
	}
	if !totalWeight.Equal(types.TotalValidatorWeight) {
		return nil, errors.Wrapf(
			types.ErrWhitelistedValidatorsList,
			"weights don't add up; expected %s, got %s",
			types.TotalValidatorWeight.String(), totalWeight.String(),
		)
	}

	pool.WhitelistedValidators = msg.WhitelistedValidators
	s.Keeper.SetPool(ctx, pool)

	if err := ctx.EventManager().EmitTypedEvent(&types.PoolUpdatedEvent{
		PoolId:    pool.PoolId,
		Pool:      &pool,
		Authority: msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit pool updated event")
	}
	return &types.MsgUpdateWhitelistedValidatorsResponse{}, nil
}

// ----------------------------------------------------------------------------
// UpdateWeightedRewardsReceivers (pool admin only)
// ----------------------------------------------------------------------------

func (s msgServer) UpdateWeightedRewardsReceivers(goCtx context.Context, msg *types.MsgUpdateWeightedRewardsReceivers) (*types.MsgUpdateWeightedRewardsReceiversResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}
	if msg.Authority != pool.WhitelistAdminAddress {
		return nil, errors.Wrapf(
			sdkerrors.ErrorInvalidSigner,
			"invalid authority; expected pool admin %s, got %s",
			pool.WhitelistAdminAddress, msg.Authority,
		)
	}

	if err := types.ValidateWeightedRewardsReceivers(msg.WeightedRewardsReceivers); err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	pool.WeightedRewardsReceivers = msg.WeightedRewardsReceivers
	s.Keeper.SetPool(ctx, pool)

	if err := ctx.EventManager().EmitTypedEvent(&types.PoolUpdatedEvent{
		PoolId:    pool.PoolId,
		Pool:      &pool,
		Authority: msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit pool updated event")
	}
	return &types.MsgUpdateWeightedRewardsReceiversResponse{}, nil
}

// ----------------------------------------------------------------------------
// SetPoolPaused (governance or pool admin)
// ----------------------------------------------------------------------------

func (s msgServer) SetPoolPaused(goCtx context.Context, msg *types.MsgSetPoolPaused) (*types.MsgSetPoolPausedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := s.Keeper.MustGetPool(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}
	if err := s.authorisedByGovOrPoolAdmin(msg.Authority, pool); err != nil {
		return nil, err
	}

	pool.Paused = msg.IsPaused
	s.Keeper.SetPool(ctx, pool)

	if err := ctx.EventManager().EmitTypedEvent(&types.PoolUpdatedEvent{
		PoolId:    pool.PoolId,
		Pool:      &pool,
		Authority: msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit pool updated event")
	}
	return &types.MsgSetPoolPausedResponse{}, nil
}

// ----------------------------------------------------------------------------
// SetModulePaused (governance only — global kill switch)
// ----------------------------------------------------------------------------

func (s msgServer) SetModulePaused(goCtx context.Context, msg *types.MsgSetModulePaused) (*types.MsgSetModulePausedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := s.authorisedByGov(msg.Authority); err != nil {
		return nil, err
	}

	params := s.Keeper.GetModuleParams(ctx)
	params.ModulePaused = msg.IsPaused
	if err := s.Keeper.SetModuleParams(ctx, params); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.ModuleParamsUpdatedEvent{
		ModuleParams: &params,
		Authority:    msg.Authority,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to emit module params updated event")
	}
	return &types.MsgSetModulePausedResponse{}, nil
}

// ----------------------------------------------------------------------------
// Burn (uixo only, module-level)
// ----------------------------------------------------------------------------

func (s msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Amount.Denom != "uixo" {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "burning amount must be in uixo")
	}
	coins := sdk.NewCoins(msg.Amount)

	burnerAddr, err := sdk.AccAddressFromBech32(msg.GetBurner())
	if err != nil {
		return nil, err
	}
	if err := s.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, burnerAddr, types.ModuleName, coins); err != nil {
		return nil, err
	}
	if err := s.Keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, err
	}
	return &types.MsgBurnResponse{}, nil
}
