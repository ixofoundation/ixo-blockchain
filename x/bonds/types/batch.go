package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
)

func (b Batch) MoreBuysThanSells() bool { return b.TotalSellAmount.IsLT(b.TotalBuyAmount) }
func (b Batch) MoreSellsThanBuys() bool { return b.TotalBuyAmount.IsLT(b.TotalSellAmount) }
func (b Batch) EqualBuysAndSells() bool { return b.TotalBuyAmount.IsEqual(b.TotalSellAmount) }
func (b Batch) HasNextAlpha() bool      { return !b.NextPublicAlpha.IsNegative() }
func (b Batch) Empty() bool             { return len(b.Buys) == 0 && len(b.Sells) == 0 && len(b.Swaps) == 0 }

func NewBatch(bondDid did.Did, token string, blocks sdk.Uint) Batch {
	return Batch{
		BondDid:         bondDid,
		BlocksRemaining: blocks,
		NextPublicAlpha: sdk.OneDec().Neg(),
		TotalBuyAmount:  sdk.NewInt64Coin(token, 0),
		TotalSellAmount: sdk.NewInt64Coin(token, 0),
	}
}

func NewBaseOrder(accountDid did.Did, amount sdk.Coin) BaseOrder {
	return BaseOrder{
		AccountDid:   accountDid,
		Amount:       amount,
		Cancelled:    false,
		CancelReason: "",
	}
}

func (bo BaseOrder) IsCancelled() bool {
	return bo.Cancelled == true
}

func NewBuyOrder(buyerDid did.Did, amount sdk.Coin, maxPrices sdk.Coins) BuyOrder {
	return BuyOrder{
		BaseOrder: NewBaseOrder(buyerDid, amount),
		MaxPrices: maxPrices,
	}
}

func NewSellOrder(sellerDid did.Did, amount sdk.Coin) SellOrder {
	return SellOrder{
		BaseOrder: NewBaseOrder(sellerDid, amount),
	}
}

func NewSwapOrder(swapperDid did.Did, from sdk.Coin, toToken string) SwapOrder {
	return SwapOrder{
		BaseOrder: NewBaseOrder(swapperDid, from),
		ToToken:   toToken,
	}
}
