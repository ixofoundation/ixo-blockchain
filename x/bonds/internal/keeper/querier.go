package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/client"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)

const (
	QueryBonds          = "bonds"
	QueryBond           = "bond"
	QueryBatch          = "batch"
	QueryLastBatch      = "last_batch"
	QueryCurrentPrice   = "current_price"
	QueryCurrentReserve = "current_reserve"
	QueryCustomPrice    = "custom_price"
	QueryBuyPrice       = "buy_price"
	QuerySellReturn     = "sell_return"
	QuerySwapReturn     = "swap_return"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryBonds:
			return queryBonds(ctx, keeper)
		case QueryBond:
			return queryBond(ctx, path[1:], keeper)
		case QueryBatch:
			return queryBatch(ctx, path[1:], keeper)
		case QueryLastBatch:
			return queryLastBatch(ctx, path[1:], keeper)
		case QueryCurrentPrice:
			return queryCurrentPrice(ctx, path[1:], keeper)
		case QueryCurrentReserve:
			return queryCurrentReserve(ctx, path[1:], keeper)
		case QueryCustomPrice:
			return queryCustomPrice(ctx, path[1:], keeper)
		case QueryBuyPrice:
			return queryBuyPrice(ctx, path[1:], keeper)
		case QuerySellReturn:
			return querySellReturn(ctx, path[1:], keeper)
		case QuerySwapReturn:
			return querySwapReturn(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown bonds query endpoint")
		}
	}
}

func queryBonds(ctx sdk.Context, keeper Keeper) (res []byte, err sdk.Error) {
	var bondsList types.QueryBonds
	iterator := keeper.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var bond types.Bond
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bond)
		bondsList = append(bondsList, bond.BondDid)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, bondsList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBond(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, bond)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBatch(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]

	if !keeper.BatchExists(ctx, bondDid) {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("batch for '%s' does not exist", bondDid))
	}

	batch := keeper.MustGetBatch(ctx, bondDid)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, batch)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryLastBatch(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]

	if !keeper.LastBatchExists(ctx, bondDid) {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("last batch for '%s' does not exist", bondDid))
	}

	batch := keeper.MustGetLastBatch(ctx, bondDid)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, batch)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCurrentPrice(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetCurrentPricesPT(reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePriceRounded := types.RoundReservePrices(reservePrices)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, reservePriceRounded)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCurrentReserve(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	reserveBalances := keeper.CoinKeeper.GetCoins(ctx, bond.ReserveAddress)
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, reserveBalances)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryCustomPrice(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	bondCoin, err2 := client.ParseCoin(bondAmount, bond.Token)
	if err2 != nil {
		return nil, sdk.ErrInternal(err2.Error())
	}

	reservePrices, err := bond.GetPricesAtSupply(bondCoin.Amount)
	if err != nil {
		return nil, err
	}
	reservePricesRounded := types.RoundReservePrices(reservePrices)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, reservePricesRounded)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryBuyPrice(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	bondCoin, err2 := client.ParseCoin(bondAmount, bond.Token)
	if err2 != nil {
		return nil, sdk.ErrInternal(err2.Error())
	}

	// Max supply cannot be less than supply (max supply >= supply)
	adjustedSupply := keeper.GetSupplyAdjustedForBuy(ctx, bondDid)
	if bond.MaxSupply.IsLT(adjustedSupply.Add(bondCoin)) {
		return nil, types.ErrCannotMintMoreThanMaxSupply(types.DefaultCodespace)
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reservePrices, err := bond.GetPricesToMint(bondCoin.Amount, reserveBalances)
	if err != nil {
		return nil, err
	}
	reservePricesRounded := types.RoundReservePrices(reservePrices)
	txFee := bond.GetTxFees(reservePrices)

	var result types.QueryBuyPrice
	result.AdjustedSupply = adjustedSupply
	result.Prices = reservePricesRounded
	result.TxFees = txFee
	result.TotalFees = result.TxFees // used in next line
	result.TotalPrices = result.Prices.Add(result.TotalFees)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, result)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func querySellReturn(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]
	bondAmount := path[1]

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("bond '%s' does not exist", bondDid))
	}

	bondCoin, err2 := client.ParseCoin(bondAmount, bond.Token)
	if err2 != nil {
		return nil, sdk.ErrInternal(err2.Error())
	}

	if strings.ToLower(bond.AllowSells) == types.FALSE {
		return nil, types.ErrBondDoesNotAllowSelling(types.DefaultCodespace)
	}

	// Cannot burn more tokens than what exists
	adjustedSupply := keeper.GetSupplyAdjustedForSell(ctx, bondDid)
	if adjustedSupply.IsLT(bondCoin) {
		return nil, types.ErrCannotBurnMoreThanSupply(types.DefaultCodespace)
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reserveReturns := bond.GetReturnsForBurn(bondCoin.Amount, reserveBalances)
	reserveReturnsRounded := types.RoundReserveReturns(reserveReturns)

	txFees := bond.GetTxFees(reserveReturns)
	exitFees := bond.GetExitFees(reserveReturns)
	totalFees := types.AdjustFees(txFees.Add(exitFees), reserveReturnsRounded)

	var result types.QuerySellReturn
	result.AdjustedSupply = adjustedSupply
	result.Returns = reserveReturnsRounded
	result.TxFees = txFees
	result.ExitFees = exitFees
	result.TotalReturns = reserveReturnsRounded.Sub(totalFees)
	result.TotalFees = totalFees

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, result)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func querySwapReturn(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	bondDid := path[0]
	fromToken := path[1]
	fromAmount := path[2]
	toToken := path[3]

	fromCoin, err2 := client.ParseCoin(fromAmount, fromToken)
	if err2 != nil {
		return nil, sdk.ErrInternal(err2.Error())
	}

	bond, found := keeper.GetBond(ctx, bondDid)
	if !found {
		return nil, types.ErrBondDoesNotExist(types.DefaultCodespace, bondDid)
	}

	reserveBalances := keeper.GetReserveBalances(ctx, bondDid)
	reserveReturns, txFee, err := bond.GetReturnsForSwap(fromCoin, toToken, reserveBalances)
	if err != nil {
		return nil, err
	}

	var result types.QuerySwapReturn
	result.TotalFees = sdk.Coins{txFee}
	result.TotalReturns = reserveReturns

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, result)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
