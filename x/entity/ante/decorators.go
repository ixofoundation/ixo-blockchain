package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/x/entity/types"
)

type BlockNftContractTransferForEntityDecorator struct {
	entityKeeper entitykeeper.Keeper
}

func NewBlockNftContractTransferForEntityDecorator(entityKeeper entitykeeper.Keeper) BlockNftContractTransferForEntityDecorator {
	return BlockNftContractTransferForEntityDecorator{
		entityKeeper: entityKeeper,
	}
}

func (sud BlockNftContractTransferForEntityDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {

	for _, msg := range tx.GetMsgs() {
		wasmMsg, ok := msg.(*wasmtypes.MsgExecuteContract)
		if !ok {
			continue
		}

		var params entitytypes.Params
		sud.entityKeeper.ParamSpace.GetParamSetIfExists(ctx, &params)

		if wasmMsg.Contract == params.NftContractAddress {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cant execute contract set as the nft contract address")
		}
	}

	return next(newCtx, tx, simulate)
}
