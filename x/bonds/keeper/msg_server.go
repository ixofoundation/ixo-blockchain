package keeper

import (
	"context"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v5/x/bonds/types"
	"golang.org/x/exp/slices"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bonds MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func augmentedFunctionBuilder(msg *types.MsgCreateBond) error {
	paramsMap := msg.FunctionParameters.AsMap()
	d0 := paramsMap["d0"]
	p0 := paramsMap["p0"]
	theta := paramsMap["theta"]
	kappa := paramsMap["kappa"]

	R0 := d0.Mul(math.LegacyOneDec().Sub(theta))
	S0 := d0.Quo(p0)
	V0, err := types.Invariant(R0, S0, kappa)
	if err != nil {
		return err
	}

	msg.FunctionParameters = msg.FunctionParameters.AddParams(
		types.FunctionParams{
			types.NewFunctionParam("R0", R0),
			types.NewFunctionParam("S0", S0),
			types.NewFunctionParam("V0", V0),
		})

	if msg.AlphaBond {
		publicAlpha := types.StartingPublicAlpha
		systemAlpha := types.SystemAlpha(publicAlpha, math.OneInt(),
			math.OneInt(), R0.TruncateInt(), msg.OutcomePayment)

		I0 := types.InvariantI(msg.OutcomePayment, systemAlpha, math.ZeroInt())

		msg.FunctionParameters = msg.FunctionParameters.AddParams(
			types.FunctionParams{
				types.NewFunctionParam("I0", I0),
				types.NewFunctionParam("publicAlpha", publicAlpha),
				types.NewFunctionParam("systemAlpha", systemAlpha),
			})
	}
	return nil
}

// This is where you would add the default initial function parameters
func augmentedFunction2Builder(msg *types.MsgCreateBond) {
	// if msg.AlphaBond {
	publicAlpha := types.StartingPublicAlpha
	msg.FunctionParameters = msg.FunctionParameters.AddParams(
		types.FunctionParams{
			types.NewFunctionParam("INITIAL_PUBLIC_ALPHA", publicAlpha),
		})
}

func (k msgServer) CreateBond(goCtx context.Context, msg *types.MsgCreateBond) (*types.MsgCreateBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	feeAddr, err := sdk.AccAddressFromBech32(msg.FeeAddress)
	if err != nil {
		return nil, err
	}

	reserveWithdrawalAddress, err := sdk.AccAddressFromBech32(msg.ReserveWithdrawalAddress)
	if err != nil {
		return nil, err
	}

	if k.BankKeeper.BlockedAddr(feeAddr) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", msg.FeeAddress)
	}

	// Check that bond and bond DID do not already exist
	if k.BondExists(ctx, msg.BondDid) {
		return nil, errorsmod.Wrap(types.ErrBondAlreadyExists, msg.BondDid)
	} else if k.BondDidExists(ctx, msg.Token) {
		return nil, errorsmod.Wrap(types.ErrBondTokenIsTaken, msg.Token)
	}

	stakingParams, err := k.StakingKeeper.GetParams(ctx)
	if err != nil {
		return nil, err
	}
	if msg.Token == stakingParams.BondDenom {
		return nil, types.ErrBondTokenCannotBeStakingToken
	}

	// Check that bond token not reserved
	if k.ReservedBondToken(ctx, msg.Token) {
		return nil, types.ErrReservedBondToken
	}

	// Set state to open by default (overridden below if augmented function)
	state := types.OpenState

	// If augmented, add R0, S0, V0 as parameters for quick access
	// Also, override AllowSells and set to False if S0 > 0

	switch msg.FunctionType {
	case types.AugmentedFunction:
		err := augmentedFunctionBuilder(msg)
		if err != nil {
			panic(err)
		}
		// The starting state for augmented bonding curves is the Hatch state.
		// Note that we can never start with OpenState since S0>0 (S0=d0/p0 and d0>0).
		state = types.HatchState
	case types.BondingFunction:
		augmentedFunction2Builder(msg)
		state = types.HatchState
	}

	bond := types.NewBond(msg.Token, msg.Name, msg.Description, msg.CreatorDid,
		msg.ControllerDid, msg.OracleDid, msg.FunctionType, msg.FunctionParameters,
		msg.ReserveTokens, msg.TxFeePercentage, msg.ExitFeePercentage,
		feeAddr, reserveWithdrawalAddress, msg.MaxSupply, msg.OrderQuantityLimits,
		msg.SanityRate, msg.SanityMarginPercentage, msg.AllowSells,
		msg.AllowReserveWithdrawals, msg.AlphaBond, msg.BatchBlocks,
		msg.OutcomePayment, state, msg.BondDid)

	k.SetBond(ctx, bond.BondDid, bond)
	k.SetBondDid(ctx, bond.Token, bond.BondDid)
	k.SetBatch(ctx, bond.BondDid, types.NewBatch(bond.BondDid, bond.Token, msg.BatchBlocks))

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s [%s] with reserve(s) [%s] created by %s", msg.Token,
		msg.FunctionType, strings.Join(bond.ReserveTokens, ","), msg.CreatorDid))

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondCreatedEvent{
			Bond: &bond,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateBondResponse{}, nil
}

func (k msgServer) EditBond(goCtx context.Context, msg *types.MsgEditBond) (*types.MsgEditBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	if bond.CreatorDid != msg.EditorDid {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the creator of the bond")
	}

	if msg.Name != types.DoNotModifyField {
		bond.Name = msg.Name
	}
	if msg.Description != types.DoNotModifyField {
		bond.Description = msg.Description
	}

	if msg.OrderQuantityLimits != types.DoNotModifyField {
		orderQuantityLimits, err := sdk.ParseCoinsNormalized(msg.OrderQuantityLimits)
		if err != nil {
			return nil, err
		}
		bond.OrderQuantityLimits = orderQuantityLimits
	}

	if msg.SanityRate != types.DoNotModifyField {
		var sanityRate, sanityMarginPercentage math.LegacyDec
		if msg.SanityRate == "" {
			sanityRate = math.LegacyZeroDec()
			sanityMarginPercentage = math.LegacyZeroDec()
		} else {
			parsedSanityRate, err := math.LegacyNewDecFromStr(msg.SanityRate)
			if err != nil {
				return nil, errorsmod.Wrap(types.ErrArgumentMissingOrNonFloat, "sanity rate")
			} else if parsedSanityRate.IsNegative() {
				return nil, errorsmod.Wrap(types.ErrArgumentCannotBeNegative, "sanity rate")
			}
			parsedSanityMarginPercentage, err := math.LegacyNewDecFromStr(msg.SanityMarginPercentage)
			if err != nil {
				return nil, errorsmod.Wrap(types.ErrArgumentMissingOrNonFloat, "sanity margin percentage")
			} else if parsedSanityMarginPercentage.IsNegative() {
				return nil, errorsmod.Wrap(types.ErrArgumentCannotBeNegative, "sanity margin percentage")
			}
			sanityRate = parsedSanityRate
			sanityMarginPercentage = parsedSanityMarginPercentage
		}
		bond.SanityRate = sanityRate
		bond.SanityMarginPercentage = sanityMarginPercentage
	}

	k.SetBond(ctx, bond.BondDid, bond)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s edited by %s",
		msg.BondDid, msg.EditorDid))

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondUpdatedEvent{
			Bond: &bond,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgEditBondResponse{}, nil
}

func (k msgServer) SetNextAlpha(goCtx context.Context, msg *types.MsgSetNextAlpha) (*types.MsgSetNextAlphaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	newPublicAlpha := msg.Alpha

	supportedFunctionTypes := []string{types.AugmentedFunction, types.BondingFunction}
	switch {
	case !slices.Contains(supportedFunctionTypes, bond.FunctionType):
		return nil, errorsmod.Wrap(types.ErrFunctionNotAvailableForFunctionType, "bond is not an augmented bonding curve")
	case !bond.AlphaBond:
		return nil, errorsmod.Wrap(types.ErrFunctionNotAvailableForFunctionType, "bond is not an alpha bond")
	case bond.State != types.OpenState.String():
		return nil, types.ErrInvalidStateForAction
	case bond.OracleDid.Did() != msg.OracleDid.Did():
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "editor must be the controller of the bond")
	}

	if bond.FunctionType == types.AugmentedFunction {
		// Get supply, reserve, outcome payment. Note that we get the adjusted
		// supply in order to take into consideration the influence of the buys and
		// sells in the current batch. We then get the reserve based on this supply.
		S := k.GetSupplyAdjustedForAlphaEdit(ctx, bond.BondDid).Amount
		R, err := bond.ReserveAtSupply(S)
		if err != nil {
			return nil, err
		}
		C := bond.OutcomePayment

		// Get current parameters
		paramsMap := bond.FunctionParameters.AsMap()

		// Check (newPublicAlpha != publicAlpha)
		if newPublicAlpha.Equal(paramsMap["publicAlpha"]) {
			return nil, errorsmod.Wrap(types.ErrInvalidAlpha,
				"cannot change public alpha to the current value of public alpha")
		}

		// Calculate scaled delta public alpha, to calculate new system alpha
		prevPublicAlpha := paramsMap["publicAlpha"]
		deltaPublicAlpha := newPublicAlpha.Sub(prevPublicAlpha)
		temp, err := types.ApproxPower(
			prevPublicAlpha.Mul(math.LegacyOneDec().Sub(types.StartingPublicAlpha)),
			math.LegacyMustNewDecFromStr("2"))
		if err != nil {
			return nil, err
		}
		scaledDeltaPublicAlpha := deltaPublicAlpha.Mul(temp)

		// Calculate new system alpha
		prevSystemAlpha := paramsMap["systemAlpha"]
		var newSystemAlpha math.LegacyDec
		if deltaPublicAlpha.IsPositive() {
			// 1 - (1 - scaled_delta_public_alpha) * (1 - previous_alpha)
			temp1 := math.LegacyOneDec().Sub(scaledDeltaPublicAlpha)
			temp2 := math.LegacyOneDec().Sub(prevSystemAlpha)
			newSystemAlpha = math.LegacyOneDec().Sub(temp1.Mul(temp2))
		} else {
			// (1 - scaled_delta_public_alpha) * (previous_alpha)
			temp1 := math.LegacyOneDec().Sub(scaledDeltaPublicAlpha)
			temp2 := prevSystemAlpha
			newSystemAlpha = temp1.Mul(temp2)
		}

		// Check 1 (newSystemAlpha != prevSystemAlpha)
		if newSystemAlpha.Equal(prevSystemAlpha) {
			return nil, errorsmod.Wrap(types.ErrInvalidAlpha,
				"resultant system alpha based on public alpha is unchanged")
		}
		// Check 2 (I > C * newSystemAlpha)
		if paramsMap["I0"].LTE(newSystemAlpha.MulInt(C)) {
			return nil, errorsmod.Wrap(types.ErrInvalidAlpha,
				"cannot change alpha to that value due to violated restriction [1]")
		}
		// Check 3 (R / C > newSystemAlpha - prevSystemAlpha)
		if R.QuoInt(C).LTE(newSystemAlpha.Sub(prevSystemAlpha)) {
			return nil, errorsmod.Wrap(types.ErrInvalidAlpha,
				"cannot change alpha to that value due to violated restriction [2]")
		}

		// Recalculate kappa and V0 using new alpha
		newKappa := types.Kappa(paramsMap["I0"], C, newSystemAlpha)
		_, err = types.Invariant(R, S.ToLegacyDec(), newKappa)
		if err != nil {
			return nil, err
		}

		// Get batch to set new alpha
		batch := k.MustGetBatch(ctx, bond.BondDid)
		batch.NextPublicAlpha = newPublicAlpha
		k.SetBatch(ctx, bond.BondDid, batch)
	} else if bond.FunctionType == types.BondingFunction {
		// Get supply, reserve, outcome payment. Note that we get the adjusted
		// supply in order to take into consideration the influence of the buys and
		// sells in the current batch. We then get the reserve based on this supply.
		// S := k.GetSupplyAdjustedForAlphaEdit(ctx, bond.BondDid).Amount
		// R, err := bond.ReserveAtSupply(S)
		// if err != nil {
		// 	return nil, err
		// }
		// C := bond.OutcomePayment

		var algo types.AugmentedBondRevision1
		if err := algo.Init(bond); err != nil {
			return nil, err
		}

		algoParams := algo.ExportToMap()

		// Get batch to set new alpha
		newPublicAlpha, err := types.ConvertFloat64ToDec(algoParams["ap"])
		if err != nil {
			return nil, err
		}
		batch := k.MustGetBatch(ctx, bond.BondDid)
		batch.NextPublicAlpha = newPublicAlpha
		// batch.NextPublicAlphaDelta = math.NewDecFromIntWithPrec(math.NewIntFromUint64(5), 1)
		k.SetBatch(ctx, bond.BondDid, batch)
	}

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s next alpha set by %s",
		msg.BondDid, msg.OracleDid.String()))

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondSetNextAlphaEvent{
			BondDid:   msg.BondDid,
			NextAlpha: newPublicAlpha.String(),
			Signer:    msg.OracleDid.String(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgSetNextAlphaResponse{}, nil
}

func (k msgServer) UpdateBondState(goCtx context.Context, msg *types.MsgUpdateBondState) (*types.MsgUpdateBondStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}
	batch := k.MustGetBatch(ctx, msg.BondDid)

	if bond.FunctionType != types.AugmentedFunction {
		return nil, types.ErrFunctionNotAvailableForFunctionType
	} else if !types.BondStateFromString(msg.State).IsValidProgressionFrom(types.BondStateFromString(bond.State)) {
		return nil, types.ErrInvalidStateProgression
	} // Also, next state must be SETTLE or FAILED -- checked by ValidateBasic

	if bond.ControllerDid != msg.EditorDid {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the controller of the bond")
	}

	// If state is settle or failed, move all outcome payment to reserve, so
	// that it is available for share withdrawal (MsgWithdrawShare). Also, set
	// reserve balance to available reserve balance.
	if msg.State == types.SettleState.String() || msg.State == types.FailedState.String() {
		if !batch.Empty() {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized,
				"cannot update bond state to SETTLE/FAILED while there are orders in the batch")
		}
		k.MoveOutcomePaymentToReserve(ctx, bond.BondDid)

		bond = k.MustGetBond(ctx, bond.BondDid) // get updated bond
		k.setReserveBalances(ctx, bond.BondDid, bond.AvailableReserve)
	}

	// Update bond state
	k.SetBondState(ctx, bond.BondDid, msg.State)

	// emit the events
	// No need to emit event/log for state change, as SetBondState does this

	return &types.MsgUpdateBondStateResponse{}, nil
}

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	buyerDid, exists := k.iidKeeper.GetDidDocument(ctx, []byte(msg.BuyerDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	buyerAddr, err := buyerDid.GetVerificationMethodBlockchainAddress(msg.BuyerDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types.ErrBondTokenDoesNotMatchBond
	}

	// Check current state is HATCH/OPEN, max prices, order quantity limits
	if bond.State != types.OpenState.String() && bond.State != types.HatchState.String() {
		return nil, types.ErrInvalidStateForAction
	} else if !bond.ReserveDenomsEqualTo(msg.MaxPrices) {
		return nil, errorsmod.Wrap(types.ErrReserveDenomsMismatch, msg.MaxPrices.String())
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// For the swapper, the first buy is the initialisation of the reserves
	// The max prices are used as the actual prices and one token is minted
	// The amount of token serves to define the price of adding more liquidity
	if bond.CurrentSupply.IsZero() && bond.FunctionType == types.SwapperFunction {
		return performFirstSwapperFunctionBuy(ctx, k.Keeper, *msg)
	}

	// Take max that buyer is willing to pay (enforces maxPrice <= balance)
	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, buyerAddr,
		types.BatchesIntermediaryAccount, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewBuyOrder(msg.BuyerDid, msg.Amount, msg.MaxPrices)

	// Get buy price and check if can add buy order to batch
	buyPrices, sellPrices, err := k.GetUpdatedBatchPricesAfterBuy(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add buy order to batch
	k.AddBuyOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	// Cancel unfulfillable orders
	k.CancelUnfulfillableOrders(ctx, bond.BondDid)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondBuyOrderEvent{
			BondDid: msg.BondDid,
			Order:   &order,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgBuyResponse{}, nil
}

func performFirstSwapperFunctionBuy(ctx sdk.Context, keeper Keeper, msg types.MsgBuy) (*types.MsgBuyResponse, error) {
	buyerDid, exists := keeper.iidKeeper.GetDidDocument(ctx, []byte(msg.BuyerDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	buyerAddr, err := buyerDid.GetVerificationMethodBlockchainAddress(msg.BuyerDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}
	// TODO: investigate effect that a high amount has on future buyers' ability to buy.

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types.ErrBondTokenDoesNotMatchBond
	}

	// Check if initial liquidity violates sanity rate
	if bond.ReservesViolateSanityRate(msg.MaxPrices) {
		return nil, types.ErrValuesViolateSanityRate
	}

	// Use max prices as the amount to send to the liquidity pool (i.e. price)
	err = keeper.DepositIntoReserve(ctx, bond.BondDid, buyerAddr, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Mint bond tokens
	err = keeper.BankKeeper.MintCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Send bond tokens to buyer
	err = keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BondsMintBurnAccount, buyerAddr, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Update supply
	keeper.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Add(msg.Amount))

	return &types.MsgBuyResponse{}, nil
}

func (k msgServer) Sell(goCtx context.Context, msg *types.MsgSell) (*types.MsgSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sellerDid, exists := k.iidKeeper.GetDidDocument(ctx, []byte(msg.SellerDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	sellerAddr, err := sellerDid.GetVerificationMethodBlockchainAddress(msg.SellerDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check sells allowed, current state is OPEN, and order limits not exceeded
	if !bond.AllowSells {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotAllowSelling, msg.BondDid)
	} else if bond.State != types.OpenState.String() {
		return nil, types.ErrInvalidStateForAction
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types.ErrBondTokenDoesNotMatchBond
	}

	// Send coins to be burned from seller (enforces sellAmount <= balance)
	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, sellerAddr,
		types.BondsMintBurnAccount, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Burn bond tokens to be sold
	err = k.BankKeeper.BurnCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewSellOrder(msg.SellerDid, msg.Amount)

	// Get sell price and check if can add sell order to batch
	buyPrices, sellPrices, err := k.GetUpdatedBatchPricesAfterSell(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add sell order to batch
	k.AddSellOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondSellOrderEvent{
			BondDid: msg.BondDid,
			Order:   &order,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgSellResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	swapperDid, exists := k.iidKeeper.GetDidDocument(ctx, []byte(msg.SwapperDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	swapperAddr, err := swapperDid.GetVerificationMethodBlockchainAddress(msg.SwapperDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that function type is swapper_function and state is OPEN
	if bond.FunctionType != types.SwapperFunction {
		return nil, types.ErrFunctionNotAvailableForFunctionType
	} else if bond.State != types.OpenState.String() {
		return nil, types.ErrInvalidStateForAction
	}

	// Check that from and to use reserve token names
	fromAndTo := sdk.NewCoins(msg.From, sdk.NewCoin(msg.ToToken, math.OneInt()))
	fromAndToDenoms := msg.From.Denom + "," + msg.ToToken
	if !bond.ReserveDenomsEqualTo(fromAndTo) {
		return nil, errorsmod.Wrap(types.ErrReserveDenomsMismatch, fromAndToDenoms)
	}

	// Check if order quantity limit exceeded
	if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.From}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// Take coins to be swapped from swapper (enforces swapAmount <= balance)
	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, swapperAddr,
		types.BatchesIntermediaryAccount, sdk.Coins{msg.From})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewSwapOrder(msg.SwapperDid, msg.From, msg.ToToken)

	// Add swap order to batch
	k.AddSwapOrder(ctx, bond.BondDid, order)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondSwapOrderEvent{
			BondDid: msg.BondDid,
			Order:   &order,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgSwapResponse{}, nil
}

func (k msgServer) MakeOutcomePayment(goCtx context.Context, msg *types.MsgMakeOutcomePayment) (*types.MsgMakeOutcomePaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	senderDid, exists := k.iidKeeper.GetDidDocument(ctx, []byte(msg.SenderDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	senderAddr, err := senderDid.GetVerificationMethodBlockchainAddress(msg.SenderDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that state is OPEN and that outcome payment is not nil
	if bond.State != types.OpenState.String() {
		return nil, types.ErrInvalidStateForAction
	}

	// Send outcome payment to outcome payment reserve
	outcomePayment := bond.GetNewReserveCoins(msg.Amount)
	err = k.DepositOutcomePayment(ctx, bond.BondDid, senderAddr, outcomePayment)
	if err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondMakeOutcomePaymentEvent{
			BondDid:        msg.BondDid,
			SenderDid:      msg.SenderDid.Did(),
			SenderAddress:  senderAddr.String(),
			OutcomePayment: outcomePayment,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgMakeOutcomePaymentResponse{}, nil
}

func (k msgServer) WithdrawShare(goCtx context.Context, msg *types.MsgWithdrawShare) (*types.MsgWithdrawShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	recipientDid, exists := k.iidKeeper.GetDidDocument(ctx, []byte(msg.RecipientDid.Did()))
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	recipientAddr, err := recipientDid.GetVerificationMethodBlockchainAddress(msg.RecipientDid.String())
	if err != nil {
		return nil, errorsmod.Wrap(err, "Address not found in iid doc")
	}

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that state is SETTLE or FAILED
	if bond.State != types.SettleState.String() && bond.State != types.FailedState.String() {
		return nil, types.ErrInvalidStateForAction
	}

	// Get number of bond tokens owned by the recipient
	bondTokensOwnedAmount := k.BankKeeper.GetBalance(ctx, recipientAddr, bond.Token).Amount
	if bondTokensOwnedAmount.IsZero() {
		return nil, types.ErrNoBondTokensOwned
	}
	bondTokensOwned := sdk.NewCoin(bond.Token, bondTokensOwnedAmount)

	// Send coins to be burned from recipient
	err = k.BankKeeper.SendCoinsFromAccountToModule(
		ctx, recipientAddr, types.BondsMintBurnAccount, sdk.NewCoins(bondTokensOwned))
	if err != nil {
		return nil, err
	}

	// Burn bond tokens
	err = k.BankKeeper.BurnCoins(ctx, types.BondsMintBurnAccount,
		sdk.NewCoins(sdk.NewCoin(bond.Token, bondTokensOwnedAmount)))
	if err != nil {
		return nil, err
	}

	// Calculate amount owned
	remainingReserve := k.GetReserveBalances(ctx, bond.BondDid)
	bondTokensShare := bondTokensOwnedAmount.ToLegacyDec().QuoInt(bond.CurrentSupply.Amount)
	reserveOwedDec := sdk.NewDecCoinsFromCoins(remainingReserve...).MulDec(bondTokensShare)
	reserveOwed, _ := reserveOwedDec.TruncateDecimal()

	// Send coins owed to recipient
	err = k.WithdrawFromReserve(ctx, bond.BondDid, recipientAddr, reserveOwed)
	if err != nil {
		return nil, err
	}

	// Update supply
	k.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Sub(bondTokensOwned))

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondWithdrawShareEvent{
			BondDid:          msg.BondDid,
			RecipientDid:     msg.RecipientDid.Did(),
			RecipientAddress: recipientAddr.String(),
			WithdrawPayment:  reserveOwed,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawShareResponse{}, nil
}

func (k msgServer) WithdrawReserve(goCtx context.Context, msg *types.MsgWithdrawReserve) (*types.MsgWithdrawReserveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bond, found := k.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrBondDoesNotExist, msg.BondDid)
	}

	reserveWithdrawalAddress, err := sdk.AccAddressFromBech32(bond.ReserveWithdrawalAddress)
	if err != nil {
		return nil, err
	}

	// Confirm that function type is an alpha bond and state is OPEN
	if bond.FunctionType != types.AugmentedFunction {
		return nil, errorsmod.Wrap(types.ErrFunctionNotAvailableForFunctionType,
			"bond is not an augmented bonding curve")
	} else if !bond.AlphaBond {
		return nil, errorsmod.Wrap(types.ErrFunctionNotAvailableForFunctionType,
			"bond is not an alpha bond")
	} else if bond.State != types.OpenState.String() {
		return nil, types.ErrInvalidStateForAction
	}

	if !bond.AllowReserveWithdrawals {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized,
			"bond does not allow reserve withdrawals")
	}

	if bond.ControllerDid != msg.WithdrawerDid {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized,
			"withdrawer must be the controller of the bond")
	}

	// Check that amount is available
	if !msg.Amount.IsAllLTE(bond.AvailableReserve) {
		return nil, errorsmod.Wrapf(types.ErrInsufficientReserveForWithdraw,
			"available reserve: %s", bond.AvailableReserve.String())
	}

	// Send coins to withdrawer
	err = k.BankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.BondsReserveAccount, reserveWithdrawalAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	// Update total amount withdrawn from reserve. We do not use the WithdrawReserve
	// function here since we only want the available reserve to be updated. The
	// CurrentReserve (virtual reserve) reported by the bond will be unchanged.
	k.setAvailableReserve(ctx, bond.BondDid, bond.AvailableReserve.Sub(msg.Amount...))

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.BondWithdrawReserveEvent{
			BondDid:                  msg.BondDid,
			WithdrawerDid:            msg.WithdrawerDid.Did(),
			WithdrawerAddress:        msg.WithdrawerAddress,
			WithdrawAmount:           msg.Amount,
			ReserveWithdrawalAddress: reserveWithdrawalAddress.String(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawReserveResponse{}, nil
}
