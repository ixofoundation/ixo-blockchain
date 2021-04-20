package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Bonds(c context.Context, _ *types.QueryBondsRequest) (*types.QueryBondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList []string
	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bond)
		bondsList = append(bondsList, bond.BondDid)
	}

	return &types.QueryBondsResponse{Bonds: bondsList}, nil
}

func (k Keeper) BondsDetailed(c context.Context, _ *types.QueryBondsDetailedRequest) (*types.QueryBondsDetailedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList types.QueryBondsDetailed
	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bond)

		reserveBalances := k.GetReserveBalances(ctx, bond.BondDid)
		reservePrices, _ := bond.GetCurrentPricesPT(reserveBalances)
		reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

		bondsList.BondsDetailed = append(bondsList.BondsDetailed, &types.BondDetails{
			BondDid:   bond.BondDid,
			SpotPrice: reservePrices,
			Supply:    bond.CurrentSupply,
			Reserve:   reserveBalances,
		})
	}

	return &types.QueryBondsDetailedResponse{BondsDetailed: &bondsList}, nil
}

func (k Keeper) Bond(c context.Context, req *types.QueryBondRequest) (*types.QueryBondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// TODO (Stef) Change types in queries for bond dids to be of type DID and not string?
	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "batch for '%s' does not exist", bondDid)
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "last batch for '%s' does not exist", bondDid)
	}

	batch := k.MustGetLastBatch(ctx, bondDid)

	return &types.QueryLastBatchResponse{Batch: &batch}, nil
}

func (k Keeper) CurrentPrice(c context.Context, req *types.QueryCurrentPriceRequest) (*types.QueryCurrentPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

	return &types.QueryCurrentPriceResponse{BuyPrices: reservePrices}, nil
}

func (k Keeper) CurrentReserve(c context.Context, req *types.QueryCurrentReserveRequest) (*types.QueryCurrentReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := zeroReserveTokensIfEmpty(bond.CurrentReserve, bond)

	return &types.QueryCurrentReserveResponse{Coins: reserveBalances}, nil
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
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

	return &types.QueryCustomPriceResponse{DecCoins: reservePrices}, nil
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
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

	var result types.QueryBuyPrice
	result.AdjustedSupply = adjustedSupply
	result.Prices = zeroReserveTokensIfEmpty(reservePricesRounded, bond)
	result.TxFees = zeroReserveTokensIfEmpty(txFee, bond)
	result.TotalPrices = zeroReserveTokensIfEmpty(reservePricesRounded.Add(txFee...), bond)
	result.TotalFees = zeroReserveTokensIfEmpty(txFee, bond)

	return &types.QueryBuyPriceResponse{QueryBuyPrice: &result}, nil
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
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

	var result types.QuerySellReturn
	result.AdjustedSupply = adjustedSupply
	result.Returns = zeroReserveTokensIfEmpty(reserveReturnsRounded, bond)
	result.TxFees = zeroReserveTokensIfEmpty(txFees, bond)
	result.ExitFees = zeroReserveTokensIfEmpty(exitFees, bond)
	result.TotalReturns = zeroReserveTokensIfEmpty(reserveReturnsRounded.Sub(totalFees), bond)
	result.TotalFees = zeroReserveTokensIfEmpty(totalFees, bond)

	return &types.QuerySellReturnResponse{QuerySellReturn: &result}, nil
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
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	reserveBalances := k.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(fromCoin, toToken, reserveBalances)
	if err != nil {
		return nil, err
	}

	if reserveReturns.Empty() {
		reserveReturns = sdk.Coins{sdk.Coin{Denom: toToken, Amount: sdk.ZeroInt()}}
	}

	var result types.QuerySwapReturn
	result.TotalFees = sdk.Coins{txFee}
	result.TotalReturns = reserveReturns

	return &types.QuerySwapReturnResponse{QuerySwapReturn: &result}, nil
}

func (k Keeper) AlphaMaximums(c context.Context, req *types.QueryAlphaMaximumsRequest) (*types.QueryAlphaMaximumsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := k.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	if bond.FunctionType != types.AugmentedFunction {
		return nil, sdkerrors.Wrapf(types.ErrFunctionNotAvailableForFunctionType, bond.FunctionType)
	}

	var maxSystemAlphaIncrease, maxSystemAlpha sdk.Dec
	if len(bond.CurrentReserve) == 0 {
		maxSystemAlphaIncrease = sdk.ZeroDec()
		maxSystemAlpha = sdk.ZeroDec()
	} else {
		R := bond.CurrentReserve[0].Amount // common reserve balance
		C := bond.OutcomePayment
		maxSystemAlphaIncrease = sdk.NewDecFromInt(R).QuoInt(C)

		paramsMap := bond.FunctionParameters.AsMap()
		I := paramsMap["I0"]
		maxSystemAlpha = I.QuoInt(C)
	}

	var result types.QueryAlphaMaximums
	result.MaxSystemAlphaIncrease = maxSystemAlphaIncrease
	result.MaxSystemAlpha = maxSystemAlpha

	return &types.QueryAlphaMaximumsResponse{QueryAlphaMaximums: &result}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}