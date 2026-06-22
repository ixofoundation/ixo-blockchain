package keeper

import (
	"context"

	"github.com/ixofoundation/ixo-blockchain/v7/x/bonds/types"
)

// disabledMsgServer implements types.MsgServer but rejects every message with
// ErrBondsModuleDisabled.
//
// The bonds module was disabled in the v8 emergency security upgrade following
// the 2026-06-20 reserve-drain incident, in which every bonds handler resolved
// a DID-fragment to a blockchain address and spent that address's funds without
// ever checking the address against the transaction signer.
//
// Registering this in place of NewMsgServerImpl neutralises EVERY state-changing
// bonds route with a single, clear error — top-level messages, authz.MsgExec,
// authz dispatch, CosmWasm stargate, and ICA-host — because they all funnel
// through the message-service router. The original (now unreachable) handler
// logic in msg_server.go is left intact so genesis/query/state still load and so
// the module can be re-enabled deliberately in a future release by:
//   - registering NewMsgServerImpl again in module.go RegisterServices, and
//   - restoring the EndBlocker call in module.go EndBlock,
// AFTER the signer-vs-resolved-address bug is fixed in every handler.
type disabledMsgServer struct{}

// NewDisabledMsgServerImpl returns a types.MsgServer that rejects all bonds
// messages with ErrBondsModuleDisabled.
func NewDisabledMsgServerImpl() types.MsgServer { return disabledMsgServer{} }

var _ types.MsgServer = disabledMsgServer{}

func (disabledMsgServer) CreateBond(context.Context, *types.MsgCreateBond) (*types.MsgCreateBondResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) EditBond(context.Context, *types.MsgEditBond) (*types.MsgEditBondResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) SetNextAlpha(context.Context, *types.MsgSetNextAlpha) (*types.MsgSetNextAlphaResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) UpdateBondState(context.Context, *types.MsgUpdateBondState) (*types.MsgUpdateBondStateResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) Buy(context.Context, *types.MsgBuy) (*types.MsgBuyResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) Sell(context.Context, *types.MsgSell) (*types.MsgSellResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) Swap(context.Context, *types.MsgSwap) (*types.MsgSwapResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) MakeOutcomePayment(context.Context, *types.MsgMakeOutcomePayment) (*types.MsgMakeOutcomePaymentResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) WithdrawShare(context.Context, *types.MsgWithdrawShare) (*types.MsgWithdrawShareResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}

func (disabledMsgServer) WithdrawReserve(context.Context, *types.MsgWithdrawReserve) (*types.MsgWithdrawReserveResponse, error) {
	return nil, types.ErrBondsModuleDisabled
}
