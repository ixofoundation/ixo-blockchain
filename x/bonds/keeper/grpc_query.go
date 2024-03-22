package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/bonds keeper providing gRPC method
// handlers.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) Bonds(c context.Context, _ *types.QueryBondsRequest) (*types.QueryBondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList []string
	iterator := q.Keeper.GetBondIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		q.Keeper.cdc.MustUnmarshal(iterator.Value(), &bond)
		bondsList = append(bondsList, bond.BondDid)
	}

	return &types.QueryBondsResponse{Bonds: bondsList}, nil
}

func (q Querier) BondsDetailed(c context.Context, _ *types.QueryBondsDetailedRequest) (*types.QueryBondsDetailedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var bondsList []*types.BondDetails
	iterator := q.Keeper.GetBondIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		q.Keeper.cdc.MustUnmarshal(iterator.Value(), &bond)

		reserveBalances := q.Keeper.GetReserveBalances(ctx, bond.BondDid)
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

func (q Querier) Bond(c context.Context, req *types.QueryBondRequest) (*types.QueryBondResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	return &types.QueryBondResponse{Bond: &bond}, nil
}

func (q Querier) Batch(c context.Context, req *types.QueryBatchRequest) (*types.QueryBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	if !q.Keeper.BatchExists(ctx, bondDid) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "batch for '%s' does not exist", bondDid)
	}

	batch := q.Keeper.MustGetBatch(ctx, bondDid)

	return &types.QueryBatchResponse{Batch: &batch}, nil
}

func (q Querier) LastBatch(c context.Context, req *types.QueryLastBatchRequest) (*types.QueryLastBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	if !q.Keeper.LastBatchExists(ctx, bondDid) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "last batch for '%s' does not exist", bondDid)
	}

	batch := q.Keeper.MustGetLastBatch(ctx, bondDid)

	return &types.QueryLastBatchResponse{LastBatch: &batch}, nil
}

func (q Querier) CurrentPrice(c context.Context, req *types.QueryCurrentPriceRequest) (*types.QueryCurrentPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := q.Keeper.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

	return &types.QueryCurrentPriceResponse{CurrentPrice: reservePrices}, nil
}

func (q Querier) CurrentReserve(c context.Context, req *types.QueryCurrentReserveRequest) (*types.QueryCurrentReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := zeroReserveTokensIfEmpty(bond.CurrentReserve, bond)

	return &types.QueryCurrentReserveResponse{CurrentReserve: reserveBalances}, nil
}

func (q Querier) AvailableReserve(c context.Context, req *types.QueryAvailableReserveRequest) (*types.QueryAvailableReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	availableReserve := zeroReserveTokensIfEmpty(bond.AvailableReserve, bond)

	return &types.QueryAvailableReserveResponse{AvailableReserve: availableReserve}, nil
}

func (q Querier) CustomPrice(c context.Context, req *types.QueryCustomPriceRequest) (*types.QueryCustomPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := q.Keeper.GetBond(ctx, bondDid)
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

	return &types.QueryCustomPriceResponse{Price: reservePrices}, nil
}

func (q Querier) BuyPrice(c context.Context, req *types.QueryBuyPriceRequest) (*types.QueryBuyPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bondCoin, err := client.ParseTwoPartCoin(bondAmount, bond.Token)
	if err != nil {
		return nil, err

	}

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := q.Keeper.GetSupplyAdjustedForBuy(ctx, bondDid)
	if bond.MaxSupply.IsLT(adjustedSupply.Add(bondCoin)) {
		return nil, types.ErrCannotMintMoreThanMaxSupply
	}

	reserveBalances := q.Keeper.GetReserveBalances(ctx, bondDid)
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

func (q Querier) SellReturn(c context.Context, req *types.QuerySellReturnRequest) (*types.QuerySellReturnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid
	bondAmount := req.BondAmount

	bond, found := q.Keeper.GetBond(ctx, bondDid)
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
	adjustedSupply := q.Keeper.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(bondCoin) {
		return nil, types.ErrCannotBurnMoreThanSupply
	}

	reserveBalances := q.Keeper.GetReserveBalances(ctx, bondDid)
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
		TotalReturns:   zeroReserveTokensIfEmpty(reserveReturnsRounded.Sub(totalFees), bond),
		TotalFees:      zeroReserveTokensIfEmpty(totalFees, bond),
	}, nil
}

func (q Querier) SwapReturn(c context.Context, req *types.QuerySwapReturnRequest) (*types.QuerySwapReturnResponse, error) {
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

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	reserveBalances := q.Keeper.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(fromCoin, toToken, reserveBalances)
	if err != nil {
		return nil, err
	}

	if reserveReturns.Empty() {
		reserveReturns = sdk.Coins{sdk.Coin{Denom: toToken, Amount: sdk.ZeroInt()}}
	}

	return &types.QuerySwapReturnResponse{
		TotalFees:    sdk.Coins{txFee},
		TotalReturns: reserveReturns,
	}, nil
}

func (q Querier) AlphaMaximums(c context.Context, req *types.QueryAlphaMaximumsRequest) (*types.QueryAlphaMaximumsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bondDid := req.BondDid

	bond, found := q.Keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	if bond.FunctionType != types.AugmentedFunction {
		return nil, sdkerrors.Wrapf(types.ErrFunctionNotAvailableForFunctionType, bond.FunctionType)
	} else if !bond.AlphaBond {
		return nil, sdkerrors.Wrap(types.ErrFunctionNotAvailableForFunctionType, "bond is not an alpha bond")
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

	return &types.QueryAlphaMaximumsResponse{
		MaxSystemAlphaIncrease: maxSystemAlphaIncrease,
		MaxSystemAlpha:         maxSystemAlpha,
	}, nil
}

func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}

// HELPERS
// ==================================

func zeroReserveTokensIfEmpty(reserveCoins sdk.Coins, bond types.Bond) sdk.Coins {
	if reserveCoins.IsZero() {
		zeroes, _ := bond.GetNewReserveDecCoins(sdk.OneDec()).TruncateDecimal()
		for i := range zeroes {
			zeroes[i].Amount = sdk.ZeroInt()
		}
		reserveCoins = zeroes
	}
	return reserveCoins
}

func zeroReserveTokensIfEmptyDec(reserveCoins sdk.DecCoins, bond types.Bond) sdk.DecCoins {
	if reserveCoins.IsZero() {
		zeroes := bond.GetNewReserveDecCoins(sdk.OneDec())
		for i := range zeroes {
			zeroes[i].Amount = sdk.ZeroDec()
		}
		reserveCoins = zeroes
	}
	return reserveCoins
}
