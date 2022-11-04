package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	tokenkeeper "github.com/ixofoundation/ixo-blockchain/x/token/keeper"
	tokentypes "github.com/ixofoundation/ixo-blockchain/x/token/types"
)

type BlockNftContractTransferForTokenDecorator struct {
	tokenKeeper tokenkeeper.Keeper
}

func NewBlockNftContractTransferForTokenDecorator(tokenKeeper tokenkeeper.Keeper) BlockNftContractTransferForTokenDecorator {
	return BlockNftContractTransferForTokenDecorator{
		tokenKeeper: tokenKeeper,
	}
}

func (sud BlockNftContractTransferForTokenDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {

	for _, msg := range tx.GetMsgs() {
		wasmMsg, ok := msg.(*wasmtypes.MsgExecuteContract)
		if !ok {
			continue
		}

		var params tokentypes.Params
		sud.tokenKeeper.ParamSpace.GetParamSetIfExists(ctx, &params)

		if wasmMsg.Contract == params.NftContractAddress {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cant execute contract set as the nft contract address")
		}
	}

	return next(newCtx, tx, simulate)
}
