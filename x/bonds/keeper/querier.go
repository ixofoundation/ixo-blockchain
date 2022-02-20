package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryBonds            = "bonds"
	QueryBondsDetailed    = "bonds_detailed"
	QueryBond             = "bond"
	QueryBatch            = "batch"
	QueryLastBatch        = "last_batch"
	QueryCurrentPrice     = "current_price"
	QueryCurrentReserve   = "current_reserve"
	QueryAvailableReserve = "available_reserve"
	QueryCustomPrice      = "custom_price"
	QueryBuyPrice         = "buy_price"
	QuerySellReturn       = "sell_return"
	QuerySwapReturn       = "swap_return"
	QueryAlphaMaximums    = "alpha_maximums"
	QueryParams           = "params"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryBonds:
			return queryBonds(ctx, keeper, legacyQuerierCdc)
		case QueryBondsDetailed:
			return queryBondsDetailed(ctx, keeper, legacyQuerierCdc)
		case QueryBond:
			return queryBond(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryBatch:
			return queryBatch(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryLastBatch:
			return queryLastBatch(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryCurrentPrice:
			return queryCurrentPrice(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryCurrentReserve:
			return queryCurrentReserve(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryAvailableReserve:
			return queryAvailableReserve(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryCustomPrice:
			return queryCustomPrice(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryBuyPrice:
			return queryBuyPrice(ctx, path[1:], keeper, legacyQuerierCdc)
		case QuerySellReturn:
			return querySellReturn(ctx, path[1:], keeper, legacyQuerierCdc)
		case QuerySwapReturn:
			return querySwapReturn(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryAlphaMaximums:
			return queryAlphaMaximums(ctx, path[1:], keeper, legacyQuerierCdc)
		case QueryParams:
			return queryParams(ctx, keeper, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown bonds query endpoint")
		}
	}
}

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

func queryBonds(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	var bondsList []string
	iterator := keeper.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		keeper.cdc.MustUnmarshal(iterator.Value(), &bond)
		bondsList = append(bondsList, bond.BondDid)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, bondsList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBondsDetailed(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	var bondsList []*types.BondDetails
	iterator := keeper.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		keeper.cdc.MustUnmarshal(iterator.Value(), &bond)

		reserveBalances := keeper.GetReserveBalances(ctx, bond.BondDid)
		reservePrices, _ := bond.GetCurrentPricesPT(reserveBalances)
		reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

		bondsList = append(bondsList, &types.BondDetails{
			BondDid:   bond.BondDid,
			SpotPrice: reservePrices,
			Supply:    bond.CurrentSupply,
			Reserve:   reserveBalances,
		})
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, bondsList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBond(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, bond)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBatch(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	if !keeper.BatchExists(ctx, bondDid) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "batch for '%s' does not exist", bondDid)
	}

	batch := keeper.MustGetBatch(ctx, bondDid)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, batch)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryLastBatch(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	if !keeper.LastBatchExists(ctx, bondDid) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "last batch for '%s' does not exist", bondDid)
	}

	batch := keeper.MustGetLastBatch(ctx, bondDid)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, batch)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCurrentPrice(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePrices = zeroReserveTokensIfEmptyDec(reservePrices, bond)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, reservePrices)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCurrentReserve(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	reserveBalances := zeroReserveTokensIfEmpty(bond.CurrentReserve, bond)
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, reserveBalances)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryAvailableReserve(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	availableReserve := zeroReserveTokensIfEmpty(bond.AvailableReserve, bond)
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, availableReserve)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCustomPrice(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
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

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, reservePrices)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBuyPrice(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "bond '%s' does not exist", bondDid)
	}

	bondCoin, err := client.ParseTwoPartCoin(bondAmount, bond.Token)
	if err != nil {
		return nil, err

	}

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := keeper.GetSupplyAdjustedForBuy(ctx, bondDid)
	if bond.MaxSupply.IsLT(adjustedSupply.Add(bondCoin)) {
		return nil, types.ErrCannotMintMoreThanMaxSupply
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetPricesToMint(bondCoin.Amount, reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePricesRounded := types.RoundReservePrices(reservePrices)
	txFee := bond.GetTxFees(reservePrices)

	var result = types.QueryBuyPriceResponse{
		AdjustedSupply: adjustedSupply,
		Prices:         zeroReserveTokensIfEmpty(reservePricesRounded, bond),
		TxFees:         zeroReserveTokensIfEmpty(txFee, bond),
		TotalPrices:    zeroReserveTokensIfEmpty(reservePricesRounded.Add(txFee...), bond),
		TotalFees:      zeroReserveTokensIfEmpty(txFee, bond),
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func querySellReturn(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
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
	adjustedSupply := keeper.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(bondCoin) {
		return nil, types.ErrCannotBurnMoreThanSupply
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reserveReturns, err := bond.GetReturnsForBurn(bondCoin.Amount, reserveBalances)
	if err != nil {
		return nil, err
	}
	reserveReturnsRounded := types.RoundReserveReturns(reserveReturns)

	txFees := bond.GetTxFees(reserveReturns)
	exitFees := bond.GetExitFees(reserveReturns)
	totalFees := types.AdjustFees(txFees.Add(exitFees...), reserveReturnsRounded)

	var result = types.QuerySellReturnResponse{
		AdjustedSupply: adjustedSupply,
		Returns:        zeroReserveTokensIfEmpty(reserveReturnsRounded, bond),
		TxFees:         zeroReserveTokensIfEmpty(txFees, bond),
		ExitFees:       zeroReserveTokensIfEmpty(exitFees, bond),
		TotalReturns:   zeroReserveTokensIfEmpty(reserveReturnsRounded.Sub(totalFees), bond),
		TotalFees:      zeroReserveTokensIfEmpty(totalFees, bond),
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func querySwapReturn(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]
	fromToken := path[1]
	fromAmount := path[2]
	toToken := path[3]

	fromCoin, err := client.ParseTwoPartCoin(fromAmount, fromToken)
	if err != nil {
		return nil, err
	}

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBondDoesNotExist, bondDid)
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(fromCoin, toToken, reserveBalances)
	if err != nil {
		return nil, err
	}

	if reserveReturns.Empty() {
		reserveReturns = sdk.Coins{sdk.Coin{Denom: toToken, Amount: sdk.ZeroInt()}}
	}

	var result = types.QuerySwapReturnResponse{
		TotalFees:    sdk.Coins{txFee},
		TotalReturns: reserveReturns,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryAlphaMaximums(ctx sdk.Context, path []string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) (res []byte, err error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
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

	var result = types.QueryAlphaMaximumsResponse{
		MaxSystemAlphaIncrease: maxSystemAlphaIncrease,
		MaxSystemAlpha:         maxSystemAlpha,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal data %s", err)
	}

	return bz, nil
}
