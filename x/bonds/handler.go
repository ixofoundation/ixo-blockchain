package bonds

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateBond:
			return handleMsgCreateBond(ctx, keeper, msg)
		case types.MsgEditBond:
			return handleMsgEditBond(ctx, keeper, msg)
		case types.MsgSetNextAlpha:
			return handleMsgSetNextAlpha(ctx, keeper, msg)
		case types.MsgUpdateBondState:
			return handleMsgUpdateBondState(ctx, keeper, msg)
		case types.MsgBuy:
			return handleMsgBuy(ctx, keeper, msg)
		case types.MsgSell:
			return handleMsgSell(ctx, keeper, msg)
		case types.MsgSwap:
			return handleMsgSwap(ctx, keeper, msg)
		case types.MsgMakeOutcomePayment:
			return handleMsgMakeOutcomePayment(ctx, keeper, msg)
		case types.MsgWithdrawShare:
			return handleMsgWithdrawShare(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized bonds Msg type: %v", msg.Type())
		}
	}
}

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bond := keeper.MustGetBondByKey(ctx, iterator.Key())
		batch := keeper.MustGetBatch(ctx, bond.BondDid)

		// Subtract one block
		batch.BlocksRemaining = batch.BlocksRemaining.SubUint64(1)
		keeper.SetBatch(ctx, bond.BondDid, batch)

		// If blocks remaining > 0 do not perform orders
		if !batch.BlocksRemaining.IsZero() {
			continue
		}

		// Store current reserve to check if this has changed later on
		reserveBeforeOrderProcessing := bond.CurrentReserve

		// Perform orders
		keeper.PerformOrders(ctx, bond.BondDid)

		// Get bond again just in case current supply was updated
		// Get batch again just in case orders were cancelled
		bond = keeper.MustGetBond(ctx, bond.BondDid)
		batch = keeper.MustGetBatch(ctx, bond.BondDid)

		// For augmented, if hatch phase and newSupply >= S0, go to open phase
		if bond.FunctionType == types.AugmentedFunction &&
			bond.State == types.HatchState {
			args := bond.FunctionParameters.AsMap()
			if bond.CurrentSupply.Amount.ToDec().GTE(args["S0"]) {
				keeper.SetBondState(ctx, bond.BondDid, types.OpenState)
			}
		}

		// Update alpha value if in open state and next alpha is not null
		if bond.State == types.OpenState && batch.HasNextAlpha() {
			keeper.UpdateAlpha(ctx, bond.BondDid)
		}

		// Save current batch as last batch and reset current batch
		keeper.SetLastBatch(ctx, bond.BondDid, batch)
		keeper.SetBatch(ctx, bond.BondDid, types.NewBatch(bond.BondDid, bond.Token, bond.BatchBlocks))

		// If reserve has not changed, no need to recalculate I0; rest of function can be skipped
		if bond.CurrentReserve.IsEqual(reserveBeforeOrderProcessing) {
			continue
		}

		// Recalculate and re-set I0 if alpha bond
		if bond.AlphaBond {
			paramsMap := bond.FunctionParameters.AsMap()
			newI0 := types.InvariantI(bond.OutcomePayment, paramsMap["alpha"],
				bond.CurrentReserve[0].Amount)
			bond.FunctionParameters.ReplaceParam("I0", newI0)
		}

		// Save bond
		keeper.SetBond(ctx, bond.BondDid, bond)
	}
	return []abci.ValidatorUpdate{}
}

func handleMsgCreateBond(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgCreateBond) (*sdk.Result, error) {
	if keeper.BankKeeper.BlacklistedAddr(msg.FeeAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", msg.FeeAddress)
	}

	// Check that bond and bond DID do not already exist
	if keeper.BondExists(ctx, msg.BondDid) {
		return nil, sdkerrors.Wrap(types.ErrBondAlreadyExists, msg.BondDid)
	} else if keeper.BondDidExists(ctx, msg.Token) {
		return nil, sdkerrors.Wrap(types.ErrBondTokenIsTaken, msg.Token)
	} else if msg.Token == keeper.StakingKeeper.GetParams(ctx).BondDenom {
		return nil, types.ErrBondTokenCannotBeStakingToken
	}

	// Check that bond token not reserved
	if keeper.ReservedBondToken(ctx, msg.Token) {
		return nil, types.ErrReservedBondToken
	}

	// Set state to open by default (overridden below if augmented function)
	state := types.OpenState

	// If augmented, add R0, S0, V0 as parameters for quick access
	// Also, override AllowSells and set to False if S0 > 0
	if msg.FunctionType == types.AugmentedFunction {
		paramsMap := msg.FunctionParameters.AsMap()
		d0, _ := paramsMap["d0"]
		p0, _ := paramsMap["p0"]
		theta, _ := paramsMap["theta"]
		kappa, _ := paramsMap["kappa"]

		R0 := d0.Mul(sdk.OneDec().Sub(theta))
		S0 := d0.Quo(p0)
		V0, err := types.Invariant(R0, S0, kappa)
		if err != nil {
			return nil, err
		}

		msg.FunctionParameters = msg.FunctionParameters.AddParams(
			types.FunctionParams{
				types.NewFunctionParam("R0", R0),
				types.NewFunctionParam("S0", S0),
				types.NewFunctionParam("V0", V0),
			})

		if msg.AlphaBond {
			alpha := types.Alpha(sdk.OneInt(), sdk.OneInt(),
				R0.TruncateInt(), msg.OutcomePayment)

			I0 := types.InvariantI(msg.OutcomePayment, alpha, sdk.ZeroInt())

			msg.FunctionParameters = msg.FunctionParameters.AddParams(
				types.FunctionParams{
					types.NewFunctionParam("I0", I0),
					types.NewFunctionParam("alpha", alpha),
				})
		}

		// The starting state for augmented bonding curves is the Hatch state.
		// Note that we can never start with OpenState since S0>0 (S0=d0/p0 and d0>0).
		state = types.HatchState
	}

	bond := types.NewBond(msg.Token, msg.Name, msg.Description, msg.CreatorDid,
		msg.ControllerDid, msg.FunctionType, msg.FunctionParameters,
		msg.ReserveTokens, msg.TxFeePercentage, msg.ExitFeePercentage,
		msg.FeeAddress, msg.MaxSupply, msg.OrderQuantityLimits, msg.SanityRate,
		msg.SanityMarginPercentage, msg.AllowSells, msg.AlphaBond, msg.BatchBlocks,
		msg.OutcomePayment, state, msg.BondDid)

	keeper.SetBond(ctx, bond.BondDid, bond)
	keeper.SetBondDid(ctx, bond.Token, bond.BondDid)
	keeper.SetBatch(ctx, bond.BondDid, types.NewBatch(bond.BondDid, bond.Token, msg.BatchBlocks))

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s [%s] with reserve(s) [%s] created by %s", msg.Token,
		msg.FunctionType, strings.Join(bond.ReserveTokens, ","), msg.CreatorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateBond,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types.AttributeKeyFunctionType, msg.FunctionType),
			sdk.NewAttribute(types.AttributeKeyFunctionParameters, msg.FunctionParameters.String()),
			sdk.NewAttribute(types.AttributeKeyCreatorDid, msg.CreatorDid),
			sdk.NewAttribute(types.AttributeKeyControllerDid, msg.ControllerDid),
			sdk.NewAttribute(types.AttributeKeyReserveTokens, types.StringsToString(msg.ReserveTokens)),
			sdk.NewAttribute(types.AttributeKeyTxFeePercentage, msg.TxFeePercentage.String()),
			sdk.NewAttribute(types.AttributeKeyExitFeePercentage, msg.ExitFeePercentage.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAddress, msg.FeeAddress.String()),
			sdk.NewAttribute(types.AttributeKeyMaxSupply, msg.MaxSupply.String()),
			sdk.NewAttribute(types.AttributeKeyOrderQuantityLimits, msg.OrderQuantityLimits.String()),
			sdk.NewAttribute(types.AttributeKeySanityRate, msg.SanityRate.String()),
			sdk.NewAttribute(types.AttributeKeySanityMarginPercentage, msg.SanityMarginPercentage.String()),
			sdk.NewAttribute(types.AttributeKeyAllowSells, strconv.FormatBool(msg.AllowSells)),
			sdk.NewAttribute(types.AttributeKeyAlphaBond, strconv.FormatBool(msg.AlphaBond)),
			sdk.NewAttribute(types.AttributeKeyBatchBlocks, msg.BatchBlocks.String()),
			sdk.NewAttribute(types.AttributeKeyOutcomePayment, msg.OutcomePayment.String()),
			sdk.NewAttribute(types.AttributeKeyState, string(state)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.CreatorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgEditBond(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgEditBond) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	if bond.CreatorDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the creator of the bond")
	}

	if msg.Name != types.DoNotModifyField {
		bond.Name = msg.Name
	}
	if msg.Description != types.DoNotModifyField {
		bond.Description = msg.Description
	}

	if msg.OrderQuantityLimits != types.DoNotModifyField {
		orderQuantityLimits, err := sdk.ParseCoins(msg.OrderQuantityLimits)
		if err != nil {
			return nil, err
		}
		bond.OrderQuantityLimits = orderQuantityLimits
	}

	if msg.SanityRate != types.DoNotModifyField {
		var sanityRate, sanityMarginPercentage sdk.Dec
		if msg.SanityRate == "" {
			sanityRate = sdk.ZeroDec()
			sanityMarginPercentage = sdk.ZeroDec()
		} else {
			parsedSanityRate, err := sdk.NewDecFromStr(msg.SanityRate)
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "sanity rate")
			} else if parsedSanityRate.IsNegative() {
				return nil, sdkerrors.Wrap(types.ErrArgumentCannotBeNegative, "sanity rate")
			}
			parsedSanityMarginPercentage, err := sdk.NewDecFromStr(msg.SanityMarginPercentage)
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "sanity margin percentage")
			} else if parsedSanityMarginPercentage.IsNegative() {
				return nil, sdkerrors.Wrap(types.ErrArgumentCannotBeNegative, "sanity margin percentage")
			}
			sanityRate = parsedSanityRate
			sanityMarginPercentage = parsedSanityMarginPercentage
		}
		bond.SanityRate = sanityRate
		bond.SanityMarginPercentage = sanityMarginPercentage
	}

	keeper.SetBond(ctx, bond.BondDid, bond)

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s edited by %s",
		msg.BondDid, msg.EditorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditBond,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types.AttributeKeyOrderQuantityLimits, msg.OrderQuantityLimits),
			sdk.NewAttribute(types.AttributeKeySanityRate, msg.SanityRate),
			sdk.NewAttribute(types.AttributeKeySanityMarginPercentage, msg.SanityMarginPercentage),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSetNextAlpha(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgSetNextAlpha) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	newAlpha := msg.Alpha

	if bond.FunctionType != types.AugmentedFunction {
		return nil, sdkerrors.Wrap(types.ErrFunctionNotAvailableForFunctionType,
			"bond is not an augmented bonding curve")
	} else if !bond.AlphaBond {
		return nil, sdkerrors.Wrap(types.ErrFunctionNotAvailableForFunctionType,
			"bond is not an alpha bond")
	} else if bond.State != types.OpenState {
		return nil, types.ErrInvalidStateForAction
	}

	if bond.ControllerDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the controller of the bond")
	}

	// Get supply, reserve, outcome payment. Note that we get the adjusted
	// supply in order to take into consideration the influence of the buys and
	// sells in the current batch. We then get the reserve based on this supply.
	S := keeper.GetSupplyAdjustedForAlphaEdit(ctx, bond.BondDid).Amount
	R, err := bond.ReserveAtSupply(S)
	if err != nil {
		return nil, err
	}
	C := bond.OutcomePayment

	// Get current parameters
	paramsMap := bond.FunctionParameters.AsMap()

	// Check 1 (newAlpha != alpha)
	if newAlpha.Equal(paramsMap["alpha"]) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAlpha,
			"cannot change alpha to the current value of alpha")
	}
	// Check 2 (I > C * alpha)
	if paramsMap["I0"].LTE(newAlpha.MulInt(C)) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAlpha,
			"cannot change alpha to that value due to violated restriction [1]")
	}
	// Check 3 (R / C > newAlpha - alpha)
	if R.QuoInt(C).LTE(newAlpha.Sub(paramsMap["alpha"])) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAlpha,
			"cannot change alpha to that value due to violated restriction [2]")
	}

	// Recalculate kappa and V0 using new alpha
	newKappa := types.Kappa(paramsMap["I0"], C, newAlpha)
	_, err = types.Invariant(R, S.ToDec(), newKappa)
	if err != nil {
		return nil, err
	}

	// Get batch to set new alpha
	batch := keeper.MustGetBatch(ctx, bond.BondDid)
	batch.NextAlpha = newAlpha
	keeper.SetBatch(ctx, bond.BondDid, batch)

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s next alpha set by %s",
		msg.BondDid, msg.EditorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetNextAlpha,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types.AttributeKeyAlpha, newAlpha.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateBondState(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgUpdateBondState) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}
	batch := keeper.MustGetBatch(ctx, msg.BondDid)

	if bond.FunctionType != types.AugmentedFunction {
		return nil, types.ErrFunctionNotAvailableForFunctionType
	} else if !msg.State.IsValidProgressionFrom(bond.State) {
		return nil, types.ErrInvalidStateProgression
	} // Also, next state must be SETTLE or FAILED -- checked by ValidateBasic

	if bond.ControllerDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the controller of the bond")
	}

	// If state is settle or failed, move all outcome payment to reserve, so
	// that it is available for share withdrawal (MsgWithdrawShare)
	if msg.State == types.SettleState || msg.State == types.FailedState {
		if !batch.Empty() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				"cannot update bond state to SETTLE/FAILED while there are orders in the batch")
		}
		keeper.MoveOutcomePaymentToReserve(ctx, bond.BondDid)
	}

	// Update bond state
	keeper.SetBondState(ctx, bond.BondDid, msg.State)

	ctx.EventManager().EmitEvents(sdk.Events{
		// No need to emit event/log for state change, as SetBondState does this
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBuy(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgBuy) (*sdk.Result, error) {
	buyerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.BuyerDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types.ErrBondTokenDoesNotMatchBond
	}

	// Check current state is HATCH/OPEN, max prices, order quantity limits
	if bond.State != types.OpenState && bond.State != types.HatchState {
		return nil, types.ErrInvalidStateForAction
	} else if !bond.ReserveDenomsEqualTo(msg.MaxPrices) {
		return nil, sdkerrors.Wrap(types.ErrReserveDenomsMismatch, msg.MaxPrices.String())
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// For the swapper, the first buy is the initialisation of the reserves
	// The max prices are used as the actual prices and one token is minted
	// The amount of token serves to define the price of adding more liquidity
	if bond.CurrentSupply.IsZero() && bond.FunctionType == types.SwapperFunction {
		return performFirstSwapperFunctionBuy(ctx, keeper, msg)
	}

	// Take max that buyer is willing to pay (enforces maxPrice <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, buyerAddr,
		types.BatchesIntermediaryAccount, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewBuyOrder(msg.BuyerDid, msg.Amount, msg.MaxPrices)

	// Get buy price and check if can add buy order to batch
	buyPrices, sellPrices, err := keeper.GetUpdatedBatchPricesAfterBuy(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add buy order to batch
	keeper.AddBuyOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	// Cancel unfulfillable orders
	keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBuy,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyMaxPrices, msg.MaxPrices.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BuyerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func performFirstSwapperFunctionBuy(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgBuy) (*sdk.Result, error) {
	buyerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.BuyerDid).Address()

	// TODO: investigate effect that a high amount has on future buyers' ability to buy.

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
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
	err := keeper.DepositReserve(ctx, bond.BondDid, buyerAddr, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Mint bond tokens
	err = keeper.SupplyKeeper.MintCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Send bond tokens to buyer
	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types.BondsMintBurnAccount, buyerAddr, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Update supply
	keeper.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Add(msg.Amount))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeInitSwapper,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyChargedPrices, msg.MaxPrices.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BuyerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSell(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgSell) (*sdk.Result, error) {
	sellerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SellerDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check sells allowed, current state is OPEN, and order limits not exceeded
	if !bond.AllowSells {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotAllowSelling, msg.BondDid)
	} else if bond.State != types.OpenState {
		return nil, types.ErrInvalidStateForAction
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types.ErrBondTokenDoesNotMatchBond
	}

	// Send coins to be burned from seller (enforces sellAmount <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, sellerAddr,
		types.BondsMintBurnAccount, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Burn bond tokens to be sold
	err = keeper.SupplyKeeper.BurnCoins(ctx, types.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewSellOrder(msg.SellerDid, msg.Amount)

	// Get sell price and check if can add sell order to batch
	buyPrices, sellPrices, err := keeper.GetUpdatedBatchPricesAfterSell(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add sell order to batch
	keeper.AddSellOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSell,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SellerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSwap(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgSwap) (*sdk.Result, error) {
	swapperAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SwapperDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that function type is swapper_function and state is OPEN
	if bond.FunctionType != types.SwapperFunction {
		return nil, types.ErrFunctionNotAvailableForFunctionType
	} else if bond.State != types.OpenState {
		return nil, types.ErrInvalidStateForAction
	}

	// Check that from and to use reserve token names
	fromAndTo := sdk.NewCoins(msg.From, sdk.NewCoin(msg.ToToken, sdk.OneInt()))
	fromAndToDenoms := msg.From.Denom + "," + msg.ToToken
	if !bond.ReserveDenomsEqualTo(fromAndTo) {
		return nil, sdkerrors.Wrap(types.ErrReserveDenomsMismatch, fromAndToDenoms)
	}

	// Check if order quantity limit exceeded
	if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.From}) {
		return nil, types.ErrOrderQuantityLimitExceeded
	}

	// Take coins to be swapped from swapper (enforces swapAmount <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, swapperAddr,
		types.BatchesIntermediaryAccount, sdk.Coins{msg.From})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types.NewSwapOrder(msg.SwapperDid, msg.From, msg.ToToken)

	// Add swap order to batch
	keeper.AddSwapOrder(ctx, bond.BondDid, order)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSwap,
			sdk.NewAttribute(types.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.From.Amount.String()),
			sdk.NewAttribute(types.AttributeKeySwapFromToken, msg.From.Denom),
			sdk.NewAttribute(types.AttributeKeySwapToToken, msg.ToToken),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SwapperDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgMakeOutcomePayment(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgMakeOutcomePayment) (*sdk.Result, error) {
	senderAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SenderDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that state is OPEN and that outcome payment is not nil
	if bond.State != types.OpenState {
		return nil, types.ErrInvalidStateForAction
	}

	// Send outcome payment to outcome payment reserve
	outcomePayment := bond.GetNewReserveCoins(msg.Amount)
	err := keeper.DepositOutcomePayment(ctx, bond.BondDid, senderAddr, outcomePayment)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMakeOutcomePayment,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, outcomePayment.String()),
			sdk.NewAttribute(types.AttributeKeyAddress, senderAddr.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SenderDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgWithdrawShare(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgWithdrawShare) (*sdk.Result, error) {
	recipientAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.RecipientDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that state is SETTLE or FAILED
	if bond.State != types.SettleState && bond.State != types.FailedState {
		return nil, types.ErrInvalidStateForAction
	}

	// Get number of bond tokens owned by the recipient
	bondTokensOwnedAmount := keeper.BankKeeper.GetCoins(ctx, recipientAddr).AmountOf(bond.Token)
	if bondTokensOwnedAmount.IsZero() {
		return nil, types.ErrNoBondTokensOwned
	}
	bondTokensOwned := sdk.NewCoin(bond.Token, bondTokensOwnedAmount)

	// Send coins to be burned from recipient
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(
		ctx, recipientAddr, types.BondsMintBurnAccount, sdk.NewCoins(bondTokensOwned))
	if err != nil {
		return nil, err
	}

	// Burn bond tokens
	err = keeper.SupplyKeeper.BurnCoins(ctx, types.BondsMintBurnAccount,
		sdk.NewCoins(sdk.NewCoin(bond.Token, bondTokensOwnedAmount)))
	if err != nil {
		return nil, err
	}

	// Calculate amount owned
	remainingReserve := keeper.GetReserveBalances(ctx, bond.BondDid)
	bondTokensShare := bondTokensOwnedAmount.ToDec().QuoInt(bond.CurrentSupply.Amount)
	reserveOwedDec := sdk.NewDecCoinsFromCoins(remainingReserve...).MulDec(bondTokensShare)
	reserveOwed, _ := reserveOwedDec.TruncateDecimal()

	// Send coins owed to recipient
	err = keeper.WithdrawReserve(ctx, bond.BondDid, recipientAddr, reserveOwed)
	if err != nil {
		return nil, err
	}

	// Update supply
	keeper.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Sub(bondTokensOwned))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawShare,
			sdk.NewAttribute(types.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types.AttributeKeyAddress, recipientAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, reserveOwed.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.RecipientDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
