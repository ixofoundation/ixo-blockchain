package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type IidResolutionDecorator struct {
	iidKeeper iidkeeper.Keeper
}

func NewIidResolutionDecorator(iidKeeper iidkeeper.Keeper) IidResolutionDecorator {
	return IidResolutionDecorator{iidKeeper: iidKeeper}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	iidTx, ok := tx.(IidTx)
	if !ok {
		// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
		return next(ctx, tx, simulate)
	}

	if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	}

	return next(ctx, tx, simulate)
}

// type InjectIidAddressDecorator struct {
// 	iidKeeper iidkeeper.Keeper
// }

// func NewInjectIidAddressDecorator(iidKeeper iidkeeper.Keeper) InjectIidAddressDecorator {
// 	return InjectIidAddressDecorator{iidKeeper: iidKeeper}
// }

// func (dec InjectIidAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	iidTx, ok := tx.(IidTx)
// 	if !ok {
// 		// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
// 		return next(ctx, tx, simulate)
// 	}

// 	if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
// 	}

// 	return next(ctx, tx, simulate)
// }

// type IidSetPubKeyDecorator struct {
// 	iidKeeper           iidkeeper.Keeper
// 	defaultPubKeyGetter authante.SetPubKeyDecorator
// }

// func NewIidSetPubKeyDecorator(iidKeeper iidkeeper.Keeper, defaultPubKeyGetter authante.SetPubKeyDecorator) IidSetPubKeyDecorator {
// 	return IidSetPubKeyDecorator{iidKeeper: iidKeeper, defaultPubKeyGetter: defaultPubKeyGetter}
// }

// func (dec IidSetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	iidTx, ok := tx.(IidTx)
// 	if !ok {
// 		dec.defaultPubKeyGetter.AnteHandle(ctx, tx, simulate, next)
// 	}

// 	sigTx, ok := tx.(authsigning.SigVerifiableTx)
// 	if !ok {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
// 	}

// 	pubkeys, err := sigTx.GetPubKeys()
// 	if err != nil {
// 		return ctx, err
// 	}
// 	signers := sigTx.GetSigners()

// 	for i, pk := range pubkeys {
// 		// PublicKey was omitted from slice since it has already been set in context
// 		if pk == nil {
// 			if !simulate {
// 				continue
// 			}
// 			pk = simSecp256k1Pubkey
// 		}
// 		// Only make check if simulate=false
// 		if !simulate && !bytes.Equal(pk.Address(), signers[i]) {
// 			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey,
// 				"pubKey does not match signer address %s with signer index: %d", signers[i], i)
// 		}

// 		acc, err := GetSignerAcc(ctx, spkd.ak, signers[i])
// 		if err != nil {
// 			return ctx, err
// 		}
// 		// account already has pubkey set,no need to reset
// 		if acc.GetPubKey() != nil {
// 			continue
// 		}
// 		err = acc.SetPubKey(pk)
// 		if err != nil {
// 			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
// 		}
// 		spkd.ak.SetAccount(ctx, acc)
// 	}

// 	// Also emit the following events, so that txs can be indexed by these
// 	// indices:
// 	// - signature (via `tx.signature='<sig_as_base64>'`),
// 	// - concat(address,"/",sequence) (via `tx.acc_seq='cosmos1abc...def/42'`).
// 	sigs, err := sigTx.GetSignaturesV2()
// 	if err != nil {
// 		return ctx, err
// 	}

// 	var events sdk.Events
// 	for i, sig := range sigs {
// 		events = append(events, sdk.NewEvent(sdk.EventTypeTx,
// 			sdk.NewAttribute(sdk.AttributeKeyAccountSequence, fmt.Sprintf("%s/%d", signers[i], sig.Sequence)),
// 		))

// 		sigBzs, err := signatureDataToBz(sig.Data)
// 		if err != nil {
// 			return ctx, err
// 		}
// 		for _, sigBz := range sigBzs {
// 			events = append(events, sdk.NewEvent(sdk.EventTypeTx,
// 				sdk.NewAttribute(sdk.AttributeKeySignature, base64.StdEncoding.EncodeToString(sigBz)),
// 			))
// 		}
// 	}

// 	ctx.EventManager().EmitEvents(events)

// 	return next(ctx, tx, simulate)
// }

type IidCapabilityVerificationDectorator struct {
	// iidKeeper iidkeeper.Keeper
}

func NewIidCapabilityVerificationDectorator() IidCapabilityVerificationDectorator {
	return IidCapabilityVerificationDectorator{}
}

func (dec IidCapabilityVerificationDectorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_, ok := tx.(IidTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
	}

	// if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
	// 	return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	// }

	return next(ctx, tx, simulate)
}
