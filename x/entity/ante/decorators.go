package ante

import (
	errorsmod "cosmossdk.io/errors"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/v7/x/entity/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v7/x/entity/types"
)

// maxMsgExecNestingDepth caps how deep we recurse into nested authz.MsgExec
// wrappers. Matches the bonds/iid ante guards.
const maxMsgExecNestingDepth = 6

type BlockNftContractTransferForEntityDecorator struct {
	entityKeeper entitykeeper.Keeper
}

func NewBlockNftContractTransferForEntityDecorator(entityKeeper entitykeeper.Keeper) BlockNftContractTransferForEntityDecorator {
	return BlockNftContractTransferForEntityDecorator{
		entityKeeper: entityKeeper,
	}
}

// AnteHandle blocks any direct MsgExecuteContract against the entity NFT
// contract — including ones nested inside authz.MsgExec — so that NFT transfers
// must go through the entity module (MsgTransferEntity), keeping NFT ownership
// and the entity DID document in sync.
//
// NOTE: this ante decorator cannot see contract-to-contract sub-messages or
// ICA-host-dispatched messages (neither runs the ante). The route-proof fix is
// a contract-level restriction (entity module as sole cw721 operator); this
// decorator closes the top-level and authz.MsgExec routes.
func (dec BlockNftContractTransferForEntityDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	var params entitytypes.Params

	if err := dec.assertNoNftContractExec(ctx, tx.GetMsgs(), &params, 0); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

// assertNoNftContractExec rejects any MsgExecuteContract targeting the entity
// NFT contract, recursing through authz.MsgExec wrappers up to
// maxMsgExecNestingDepth.
func (dec BlockNftContractTransferForEntityDecorator) assertNoNftContractExec(ctx sdk.Context, msgs []sdk.Msg, params *entitytypes.Params, depth int) error {
	if depth > maxMsgExecNestingDepth {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "authz.MsgExec nesting too deep")
	}
	for _, msg := range msgs {
		if exec, ok := msg.(*authz.MsgExec); ok {
			inner, err := exec.GetMessages()
			if err != nil {
				return errorsmod.Wrap(err, "failed to unpack authz.MsgExec messages")
			}
			if err := dec.assertNoNftContractExec(ctx, inner, params, depth+1); err != nil {
				return err
			}
			continue
		}

		wasmMsg, ok := msg.(*wasmtypes.MsgExecuteContract)
		if !ok {
			continue
		}

		// get the params if not already set
		if params.NftContractAddress == "" {
			dec.entityKeeper.ParamSpace.GetParamSetIfExists(ctx, params)
		}

		if wasmMsg.Contract == params.NftContractAddress {
			return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot execute contract set as the entity nft contract address")
		}
	}
	return nil
}
