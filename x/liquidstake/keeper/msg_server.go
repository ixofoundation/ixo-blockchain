package keeper

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/liquidstake/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the liquidstake MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// --------------------------
// LIQUID STAKE
// --------------------------
func (s msgServer) LiquidStake(goCtx context.Context, msg *types.MsgLiquidStake) (*types.MsgLiquidStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Perform the liquid stake
	stkIXOMintAmount, err := s.Keeper.LiquidStake(ctx, types.LiquidStakeProxyAcc, msg.GetDelegator(), msg.Amount)
	if err != nil {
		return nil, err
	}

	liquidBondDenom := s.LiquidBondDenom(ctx)

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidStakeEvent{
			Delegator:          msg.DelegatorAddress,
			LiquidAmount:       msg.Amount.String(),
			StkIxoMintedAmount: sdk.Coin{Denom: liquidBondDenom, Amount: stkIXOMintAmount}.String(),
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid unstake event")
	}

	return &types.MsgLiquidStakeResponse{}, nil
}

// --------------------------
// LIQUID UNSTAKE
// --------------------------
func (s msgServer) LiquidUnstake(goCtx context.Context, msg *types.MsgLiquidUnstake) (*types.MsgLiquidUnstakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Perform the liquid unstake
	completionTime, unbondingAmount, _, unbondedAmount, err := s.Keeper.LiquidUnstake(ctx, types.LiquidStakeProxyAcc, msg.GetDelegator(), msg.Amount)
	if err != nil {
		return nil, err
	}

	bondDenom, err := s.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return nil, err
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidUnstakeEvent{
			Delegator:       msg.DelegatorAddress,
			UnstakeAmount:   msg.Amount.String(),
			UnbondingAmount: sdk.Coin{Denom: bondDenom, Amount: unbondingAmount}.String(),
			UnbondedAmount:  sdk.Coin{Denom: bondDenom, Amount: unbondedAmount}.String(),
			CompletionTime:  completionTime.Format(time.RFC3339),
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid unstake event")
	}

	return &types.MsgLiquidUnstakeResponse{
		CompletionTime: completionTime,
	}, nil
}

// --------------------------
// UPDATE PARAMS
// --------------------------
func (s msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.GetParams(ctx)

	// Validate who can update the params
	if msg.Authority != s.authority && msg.Authority != params.WhitelistAdminAddress {
		return nil, errors.Wrapf(sdkerrors.ErrorInvalidSigner, "invalid authority; expected %s, got %s", s.authority, msg.Authority)
	}

	// List of all updateable param
	params.UnstakeFeeRate = msg.Params.UnstakeFeeRate
	params.MinLiquidStakeAmount = msg.Params.MinLiquidStakeAmount
	params.FeeAccountAddress = msg.Params.FeeAccountAddress
	params.AutocompoundFeeRate = msg.Params.AutocompoundFeeRate
	params.WhitelistAdminAddress = msg.Params.WhitelistAdminAddress

	// Persist the params
	err := s.SetParams(ctx, params)
	if err != nil {
		return nil, err
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidStakeParamsUpdatedEvent{
			Authority: msg.Authority,
			Params:    &params,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid stake params updated event")
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

// --------------------------
// UPDATE WHITELISTED VALIDATORS
// --------------------------
func (s msgServer) UpdateWhitelistedValidators(goCtx context.Context, msg *types.MsgUpdateWhitelistedValidators) (*types.MsgUpdateWhitelistedValidatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.GetParams(ctx)

	// Validate who can update the whitelisted validators
	if msg.Authority != s.authority && msg.Authority != params.WhitelistAdminAddress {
		return nil, errors.Wrapf(sdkerrors.ErrorInvalidSigner, "invalid authority; expected %s, got %s", params.WhitelistAdminAddress, msg.Authority)
	}

	// Validate the whitelisted validators list is valid and sum of weights adds up
	totalWeight := math.NewInt(0)
	for _, val := range msg.WhitelistedValidators {
		totalWeight = totalWeight.Add(val.TargetWeight)

		valAddr := val.GetValidatorAddress()
		fullVal, err := s.stakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			return nil, errors.Wrapf(
				types.ErrWhitelistedValidatorsList,
				"validator not found: %s", valAddr,
			)
		}

		if fullVal.Status != stakingtypes.Bonded {
			return nil, errors.Wrapf(
				types.ErrWhitelistedValidatorsList,
				"validator status %s: expected %s; got %s", valAddr, stakingtypes.Bonded.String(), fullVal.Status.String(),
			)
		}
	}

	if !totalWeight.Equal(types.TotalValidatorWeight) {
		return nil, errors.Wrapf(
			types.ErrWhitelistedValidatorsList,
			"weights don't add up; expected %s, got %s", types.TotalValidatorWeight.String(), totalWeight.String(),
		)
	}

	// Update the params and persist
	params.WhitelistedValidators = msg.WhitelistedValidators
	err := s.SetParams(ctx, params)
	if err != nil {
		return nil, err
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidStakeParamsUpdatedEvent{
			Authority: msg.Authority,
			Params:    &params,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid stake params updated event")
	}

	return &types.MsgUpdateWhitelistedValidatorsResponse{}, nil
}

// --------------------------
// UPDATE WEIGHTED REWARDS RECEIVERS
// --------------------------
func (s msgServer) UpdateWeightedRewardsReceivers(goCtx context.Context, msg *types.MsgUpdateWeightedRewardsReceivers) (*types.MsgUpdateWeightedRewardsReceiversResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.GetParams(ctx)

	// Validate only whitelist admin address can update this list
	if msg.Authority != params.WhitelistAdminAddress {
		return nil, errors.Wrapf(sdkerrors.ErrorInvalidSigner, "invalid authority; expected %s, got %s", params.WhitelistAdminAddress, msg.Authority)
	}

	// Update the params and persist
	params.WeightedRewardsReceivers = msg.WeightedRewardsReceivers
	err := s.SetParams(ctx, params)
	if err != nil {
		return nil, err
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidStakeParamsUpdatedEvent{
			Authority: msg.Authority,
			Params:    &params,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid stake params updated event")
	}

	return &types.MsgUpdateWeightedRewardsReceiversResponse{}, nil
}

// --------------------------
// SET MODULE PAUSED
// --------------------------
func (s msgServer) SetModulePaused(goCtx context.Context, msg *types.MsgSetModulePaused) (*types.MsgSetModulePausedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.GetParams(ctx)

	// Validate who can set the module paused
	if msg.Authority != s.authority && msg.Authority != params.WhitelistAdminAddress {
		return nil, errors.Wrapf(sdkerrors.ErrorInvalidSigner, "invalid authority; expected %s, got %s", params.WhitelistAdminAddress, msg.Authority)
	}

	// Update the params and persist
	params.ModulePaused = msg.IsPaused
	err := s.SetParams(ctx, params)
	if err != nil {
		return nil, err
	}

	// Emit the events
	if err := ctx.EventManager().EmitTypedEvent(
		&types.LiquidStakeParamsUpdatedEvent{
			Authority: msg.Authority,
			Params:    &params,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit liquid stake params updated event")
	}

	return &types.MsgSetModulePausedResponse{}, nil
}
