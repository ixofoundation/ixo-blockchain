package bonds

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgCreateBond:
			res, err := msgServer.CreateBond(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgEditBond:
			res, err := msgServer.EditBond(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetNextAlpha:
			res, err := msgServer.SetNextAlpha(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateBondState:
			res, err := msgServer.UpdateBondState(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBuy:
			res, err := msgServer.Buy(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSell:
			res, err := msgServer.Sell(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSwap:
			res, err := msgServer.Swap(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgMakeOutcomePayment:
			res, err := msgServer.MakeOutcomePayment(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawShare:
			res, err := msgServer.WithdrawShare(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
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
			bond.State == types.HatchState.String() {
			args := bond.FunctionParameters.AsMap()
			if bond.CurrentSupply.Amount.ToDec().GTE(args["S0"]) {
				keeper.SetBondState(ctx, bond.BondDid, types.OpenState.String())
				bond = keeper.MustGetBond(ctx, bond.BondDid) // get updated bond
			}
		}

		// Update alpha value if in open state and next alpha is not null
		if bond.State == types.OpenState.String() && batch.HasNextAlpha() {
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
			newI0 := types.InvariantI(bond.OutcomePayment, paramsMap["systemAlpha"],
				bond.CurrentReserve[0].Amount)
			bond.FunctionParameters.ReplaceParam("I0", newI0)
		}

		// Save bond
		keeper.SetBond(ctx, bond.BondDid, bond)
	}
	return []abci.ValidatorUpdate{}
}