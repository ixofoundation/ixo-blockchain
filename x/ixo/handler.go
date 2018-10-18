package ixo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func NewHandler(k Keeper) sdk.Handler {
// 	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
// 		switch msg := msg.(type) {
// 		case AddEthWalletMsg:
// 			return handleAddEthWalletMsg(ctx, k, msg)
// 		default:
// 			return sdk.ErrUnknownRequest("No match for message type.").Result()
// 		}
// 	}
// }

func handleAddEthWalletMsg(ctx sdk.Context, k Keeper, msg AddEthWalletMsg) sdk.Result {
	k.SetETHWallet(ctx, msg.Data.Id, msg.Data.WalletAddress)

	return sdk.Result{
		Code: sdk.ABCICodeOK,
	}
}
