package keeper

import (
	"context"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreateBond(ctx context.Context, bond *types.MsgCreateBond) (*types.MsgCreateBondResponse, error) {
	panic("implement me")
}

func (m msgServer) EditBond(ctx context.Context, bond *types.MsgEditBond) (*types.MsgEditBondResponse, error) {
	panic("implement me")
}

func (m msgServer) SetNextAlpha(ctx context.Context, alpha *types.MsgSetNextAlpha) (*types.MsgSetNextAlphaResponse, error) {
	panic("implement me")
}

func (m msgServer) UpdateBondState(ctx context.Context, state *types.MsgUpdateBondState) (*types.MsgUpdateBondStateResponse, error) {
	panic("implement me")
}

func (m msgServer) Buy(ctx context.Context, buy *types.MsgBuy) (*types.MsgBuyResponse, error) {
	panic("implement me")
}

func (m msgServer) Sell(ctx context.Context, sell *types.MsgSell) (*types.MsgSellResponse, error) {
	panic("implement me")
}

func (m msgServer) Swap(ctx context.Context, swap *types.MsgSwap) (*types.MsgSwapResponse, error) {
	panic("implement me")
}

func (m msgServer) MakeOutcomePayment(ctx context.Context, payment *types.MsgMakeOutcomePayment) (*types.MsgMakeOutcomePaymentResponse, error) {
	panic("implement me")
}

func (m msgServer) WithdrawShare(ctx context.Context, share *types.MsgWithdrawShare) (*types.MsgWithdrawShareResponse, error) {
	panic("implement me")
}