package ante

import (
	errorsmod "cosmossdk.io/errors"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/v4/x/entity/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
)

type BlockNftContractTransferForEntityDecorator struct {
	entityKeeper entitykeeper.Keeper
}

func NewBlockNftContractTransferForEntityDecorator(entityKeeper entitykeeper.Keeper) BlockNftContractTransferForEntityDecorator {
	return BlockNftContractTransferForEntityDecorator{
		entityKeeper: entityKeeper,
	}
}

// if MsgExecuteContract is for nft module contract then block direct MsgExecuteContract as must be done through entity module
func (dec BlockNftContractTransferForEntityDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	var params entitytypes.Params

	// loop through messages, if MsgExecuteContract is for nft module contract then block direct MsgExecuteContract as must be done through entity module
	for _, msg := range tx.GetMsgs() {
		wasmMsg, ok := msg.(*wasmtypes.MsgExecuteContract)
		if !ok {
			continue
		}

		// get the params if not already set
		if params.NftContractAddress == "" {
			dec.entityKeeper.ParamSpace.GetParamSetIfExists(ctx, &params)
		}

		if wasmMsg.Contract == params.NftContractAddress {
			return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot execute contract set as the entity nft contract address")
		}
	}

	return next(ctx, tx, simulate)
}
