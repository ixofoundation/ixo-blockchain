package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

type Batch struct {
	BondDid         exported.Did      `json:"bond_did" yaml:"bond_did"`
	BlocksRemaining sdk.Uint     `json:"blocks_remaining" yaml:"blocks_remaining"`
	NextPublicAlpha sdk.Dec      `json:"next_public_alpha" yaml:"next_public_alpha"`
	TotalBuyAmount  sdk.Coin     `json:"total_buy_amount" yaml:"total_buy_amount"`
	TotalSellAmount sdk.Coin     `json:"total_sell_amount" yaml:"total_sell_amount"`
	BuyPrices       sdk.DecCoins `json:"buy_prices" yaml:"buy_prices"`
	SellPrices      sdk.DecCoins `json:"sell_prices" yaml:"sell_prices"`
	Buys            []BuyOrder   `json:"buys" yaml:"buys"`
	Sells           []SellOrder  `json:"sells" yaml:"sells"`
	Swaps           []SwapOrder  `json:"swaps" yaml:"swaps"`
}

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

type BaseOrder struct {
	AccountDid   exported.Did  `json:"sender_did" yaml:"sender_did"`
	Amount       sdk.Coin `json:"amount" yaml:"amount"`
	Cancelled    bool     `json:"cancelled" yaml:"cancelled"`
	CancelReason string   `json:"cancel_reason" yaml:"cancel_reason"`
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

type BuyOrder struct {
	BaseOrder
	MaxPrices sdk.Coins `json:"max_prices" yaml:"max_prices"`
}

func NewBuyOrder(buyerDid did.Did, amount sdk.Coin, maxPrices sdk.Coins) BuyOrder {
	return BuyOrder{
		BaseOrder: NewBaseOrder(buyerDid, amount),
		MaxPrices: maxPrices,
	}
}

type SellOrder struct {
	BaseOrder
}

func NewSellOrder(sellerDid did.Did, amount sdk.Coin) SellOrder {
	return SellOrder{
		BaseOrder: NewBaseOrder(sellerDid, amount),
	}
}

type SwapOrder struct {
	BaseOrder
	ToToken string `json:"to_token" yaml:"to_token"`
}

func NewSwapOrder(swapperDid did.Did, from sdk.Coin, toToken string) SwapOrder {
	return SwapOrder{
		BaseOrder: NewBaseOrder(swapperDid, from),
		ToToken:   toToken,
	}
}
