package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
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

	// tx.GetMsgs()[0].

	return next(newCtx, tx, simulate)
}
