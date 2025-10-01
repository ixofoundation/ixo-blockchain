package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v6/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/v6/x/bonds/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Bonds(c context.Context, _ *types.QueryBondsRequest) (*types.QueryBondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList []string
	iterator := k.GetBondIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		k.cdc.MustUnmarshal(iterator.Value(), &bond)
		bondsList = append(bondsList, bond.BondDid)
	}

	return &types.QueryBondsResponse{Bonds: bondsList}, nil
}

func (k Keeper) BondsDetailed(c context.Context, _ *types.QueryBondsDetailedRequest) (*types.QueryBondsDetailedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList []*types.BondDetails
	iterator := k.GetBondIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		k.cdc.MustUnmarshal(iterator.Value(), &bond)

		reserveBalances := k.GetReserveBalances(ctx, bond.BondDid)
		reservePrices, _ := bond.GetCurrentPricesPT(reserveBalances)
		reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

		bondsList = append(bondsList, &types.BondDetails{
			BondDid:   bond.BondDid,
			SpotPrice: reservePrices,
			Supply:    bond.CurrentSupply,
			Reserve:   reserveBalances,
		})
	}

	return &types.QueryBondsDetailedResponse{BondsDetailed: bondsList}, nil
}

func (k Keeper) Bond(c context.Context, req *types.QueryBondRequest) (*types.QueryBondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	return &types.QueryBondResponse{Bond: &bond}, nil
}

func (k Keeper) Batch(c context.Context, req *types.QueryBatchRequest) (*types.QueryBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	if !k.BatchExists(ctx, bondDid) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "batch for '%s' does not exist", bondDid)
	}

	batch := k.MustGetBatch(ctx, bondDid)

	return &types.QueryBatchResponse{Batch: &batch}, nil
}

func (k Keeper) LastBatch(c context.Context, req *types.QueryLastBatchRequest) (*types.QueryLastBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	if !k.LastBatchExists(ctx, bondDid) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "last batch for '%s' does not exist", bondDid)
	}

	batch := k.MustGetLastBatch(ctx, bondDid)

	return &types.QueryLastBatchResponse{LastBatch: &batch}, nil
}

func (k Keeper) CurrentPrice(c context.Context, req *types.QueryCurrentPriceRequest) (*types.QueryCurrentPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

	return &types.QueryCurrentPriceResponse{CurrentPrice: reservePrices}, nil
}

func (k Keeper) CurrentReserve(c context.Context, req *types.QueryCurrentReserveRequest) (*types.QueryCurrentReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := zeroReserveTokensIfEmpty(bond.CurrentReserve, bond)

	return &types.QueryCurrentReserveResponse{CurrentReserve: reserveBalances}, nil
}

func (k Keeper) AvailableReserve(c context.Context, req *types.QueryAvailableReserveRequest) (*types.QueryAvailableReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	availableReserve := zeroReserveTokensIfEmpty(bond.AvailableReserve, bond)

	return &types.QueryAvailableReserveResponse{AvailableReserve: availableReserve}, nil
}

func (k Keeper) CustomPrice(c context.Context, req *types.QueryCustomPriceRequest) (*types.QueryCustomPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bondCoin, err := client.ParseTwoPartCoin(bondAmount, bond.Token)
	if err != nil {
		return nil, err
	}

	reservePrices, err := bond.GetPricesAtSupply(bondCoin.Amount)
	if err != nil {
		return nil, err
	}
	reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

	return &types.QueryCustomPriceResponse{Price: reservePrices}, nil
}

func (k Keeper) BuyPrice(c context.Context, req *types.QueryBuyPriceRequest) (*types.QueryBuyPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bondCoin, err := client.ParseTwoPartCoin(bondAmount, bond.Token)
	if err != nil {
		return nil, err
	}

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := k.GetSupplyAdjustedForBuy(ctx, bondDid)
	if bond.MaxSupply.IsLT(adjustedSupply.Add(bondCoin)) {
		return nil, types.ErrCannotMintMoreThanMaxSupply
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetPricesToMint(bondCoin.Amount, reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePricesRounded := types.RoundReservePrices(reservePrices)
	txFee := bond.GetTxFees(reservePrices)

	return &types.QueryBuyPriceResponse{
		AdjustedSupply: adjustedSupply,
		Prices:         zeroReserveTokensIfEmpty(reservePricesRounded, bond),
		TxFees:         zeroReserveTokensIfEmpty(txFee, bond),
		TotalPrices:    zeroReserveTokensIfEmpty(reservePricesRounded.Add(txFee...), bond),
		TotalFees:      zeroReserveTokensIfEmpty(txFee, bond),
	}, nil
}

func (k Keeper) SellReturn(c context.Context, req *types.QuerySellReturnRequest) (*types.QuerySellReturnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bondCoin, err := client.ParseTwoPartCoin(bondAmount, bond.Token)
	if err != nil {
		return nil, err
	}

	if !bond.AllowSells {
		return nil, types.ErrBondDoesNotAllowSelling
	}

	// Cannot burn more tokens than what exists
	adjustedSupply := k.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(bondCoin) {
		return nil, types.ErrCannotBurnMoreThanSupply
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reserveReturns, err := bond.GetReturnsForBurn(bondCoin.Amount, reserveBalances)
	if err != nil {
		return nil, err
	}
	reserveReturnsRounded := types.RoundReserveReturns(reserveReturns)

	txFees := bond.GetTxFees(reserveReturns)
	exitFees := bond.GetExitFees(reserveReturns)
	totalFees := types.AdjustFees(txFees.Add(exitFees...), reserveReturnsRounded)

	return &types.QuerySellReturnResponse{
		AdjustedSupply: adjustedSupply,
		Returns:        zeroReserveTokensIfEmpty(reserveReturnsRounded, bond),
		TxFees:         zeroReserveTokensIfEmpty(txFees, bond),
		ExitFees:       zeroReserveTokensIfEmpty(exitFees, bond),
		TotalReturns:   zeroReserveTokensIfEmpty(reserveReturnsRounded.Sub(totalFees...), bond),
		TotalFees:      zeroReserveTokensIfEmpty(totalFees, bond),
	}, nil
}

func (k Keeper) SwapReturn(c context.Context, req *types.QuerySwapReturnRequest) (*types.QuerySwapReturnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	toToken := req.ToToken
	fromCoin, err := sdk.ParseCoinNormalized(req.FromTokenWithAmount)
	if err != nil {
		return nil, err
	}

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(fromCoin, toToken, reserveBalances)
	if err != nil {
		return nil, err
	}

	if reserveReturns.Empty() {
		reserveReturns = sdk.Coins{sdk.Coin{Denom: toToken, Amount: math.ZeroInt()}}
	}

	return &types.QuerySwapReturnResponse{
		TotalFees:    sdk.Coins{txFee},
		TotalReturns: reserveReturns,
	}, nil
}

func (k Keeper) AlphaMaximums(c context.Context, req *types.QueryAlphaMaximumsRequest) (*types.QueryAlphaMaximumsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	if bond.FunctionType != types.AugmentedFunction {
		return nil, errorsmod.Wrapf(types.ErrFunctionNotAvailableForFunctionType, bond.FunctionType)
	} else if !bond.AlphaBond {
		return nil, errorsmod.Wrap(types.ErrFunctionNotAvailableForFunctionType, "bond is not an alpha bond")
	}

	var maxSystemAlphaIncrease, maxSystemAlpha math.LegacyDec
	if len(bond.CurrentReserve) == 0 {
		maxSystemAlphaIncrease = math.LegacyZeroDec()
		maxSystemAlpha = math.LegacyZeroDec()
	} else {
		R := bond.CurrentReserve[0].Amount // common reserve balance
		C := bond.OutcomePayment
		maxSystemAlphaIncrease = math.LegacyNewDecFromInt(R).QuoInt(C)

		paramsMap := bond.FunctionParameters.AsMap()
		I := paramsMap["I0"]
		maxSystemAlpha = I.QuoInt(C)
	}

	return &types.QueryAlphaMaximumsResponse{
		MaxSystemAlphaIncrease: maxSystemAlphaIncrease,
		MaxSystemAlpha:         maxSystemAlpha,
	}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}

// HELPERS
// ==================================

func zeroReserveTokensIfEmpty(reserveCoins sdk.Coins, bond types.Bond) sdk.Coins {
	if reserveCoins.IsZero() {
		zeroes, _ := bond.GetNewReserveDecCoins(math.LegacyOneDec()).TruncateDecimal()
		for i := range zeroes {
			zeroes[i].Amount = math.ZeroInt()
		}
		reserveCoins = zeroes
	}
	return reserveCoins
}

func zeroReserveTokensIfEmptyDec(reserveCoins sdk.DecCoins, bond types.Bond) sdk.DecCoins {
	if reserveCoins.IsZero() {
		zeroes := bond.GetNewReserveDecCoins(math.LegacyOneDec())
		for i := range zeroes {
			zeroes[i].Amount = math.LegacyZeroDec()
		}
		reserveCoins = zeroes
	}
	return reserveCoins
}
