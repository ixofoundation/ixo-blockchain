package types

import (
	"bytes"
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type SetPubKeyDecorator struct {
	ak  ante.AccountKeeper
	pkg PubKeyGetter
}

func NewSetPubKeyDecorator(ak ante.AccountKeeper, pkg PubKeyGetter) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	// all messages must be of type IxoMsg
	msg, ok := tx.GetMsgs()[0].(IxoMsg)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.IxoMsg")
	}

	// Get pubKey
	pubKey, err := spkd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer
	signerAddr := sdk.AccAddress(pubKey.Address())

	pubkeys := []cryptotypes.PubKey{pubKey}
	signers := []sdk.AccAddress{signerAddr}

	for i, pk := range pubkeys {
		// PublicKey was omitted from slice since it has already been set in context
		if pk == nil {
			if !simulate {
				continue
			}
			pk = &simEd25519Pubkey
		}
		// Only make check if simulate=false
		if !simulate && !bytes.Equal(pk.Address(), signers[i]) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey,
				"pubKey does not match signer address %s with signer index: %d", signers[i], i)
		}

		acc, err := ante.GetSignerAcc(ctx, spkd.ak, signers[i])
		if err != nil {
			return ctx, err
		}
		// account already has pubkey set,no need to reset
		if acc.GetPubKey() != nil {
			continue
		}
		err = acc.SetPubKey(pk)
		if err != nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
		}
		spkd.ak.SetAccount(ctx, acc)
	}

	return next(ctx, tx, simulate)
}

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	ak  ante.AccountKeeper
	bk  types.BankKeeper
	pkg PubKeyGetter
}

func NewDeductFeeDecorator(ak ante.AccountKeeper, bk types.BankKeeper, pkg PubKeyGetter) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:  ak,
		bk:  bk,
		pkg: pkg,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.ak.GetModuleAddress(types.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.FeeCollectorName))
	}

	// all messages must be of type IxoMsg
	msg, ok := tx.GetMsgs()[0].(IxoMsg)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.IxoMsg")
	}

	// Get pubKey
	pubKey, err := dfd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer, who's going to pay the fees
	feePayer := sdk.AccAddress(pubKey.Address())
	feePayerAcc := dfd.ak.GetAccount(ctx, feePayer)

	if feePayerAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
	}

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, feePayerAcc, feeTx.GetFee())
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

// Consume parameter-defined amount of gas for each signature according to the passed-in SignatureVerificationGasConsumer function
// before calling the next AnteHandler
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigGasConsumeDecorator struct {
	ak             keeper.AccountKeeper
	sigGasConsumer ante.SignatureVerificationGasConsumer
	pkg            PubKeyGetter
}

func NewSigGasConsumeDecorator(ak keeper.AccountKeeper, sigGasConsumer ante.SignatureVerificationGasConsumer, pkg PubKeyGetter) SigGasConsumeDecorator {
	return SigGasConsumeDecorator{
		ak:             ak,
		sigGasConsumer: sigGasConsumer,
		pkg:            pkg,
	}
}

func (sgcd SigGasConsumeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	params := sgcd.ak.GetParams(ctx)
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	// all messages must be of type IxoMsg
	msg, ok := tx.GetMsgs()[0].(IxoMsg)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.IxoMsg")
	}

	// Get pubKey
	pubKey, err := sgcd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer
	signerAddrs := []sdk.AccAddress{sdk.AccAddress(pubKey.Address())}

	for i, sig := range sigs {
		signerAcc, err := ante.GetSignerAcc(ctx, sgcd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}
		pubKey := signerAcc.GetPubKey()

		// In simulate mode the transaction comes with no signatures, thus if the
		// account's pubkey is nil, both signature verification and gasKVStore.Set()
		// shall consume the largest amount, i.e. it takes more gas to verify
		// secp256k1 keys than ed25519 ones.
		if simulate && pubKey == nil {
			pubKey = &simEd25519Pubkey
		}

		// make a SignatureV2 with PubKey filled in from above
		sig = signing.SignatureV2{
			PubKey:   pubKey,
			Data:     sig.Data,
			Sequence: sig.Sequence,
		}

		err = sgcd.sigGasConsumer(ctx.GasMeter(), sig, params)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator decorator will not get executed on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              keeper.AccountKeeper
	signModeHandler authsigning.SignModeHandler
	pkg             PubKeyGetter
}

func NewSigVerificationDecorator(ak keeper.AccountKeeper, signModeHandler authsigning.SignModeHandler, pkg PubKeyGetter) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,
		pkg:             pkg,
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	// all messages must be of type IxoMsg
	msg, ok := tx.GetMsgs()[0].(IxoMsg)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.IxoMsg")
	}

	// Get pubKey
	pubKey, err := svd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer
	signerAddrs := []sdk.AccAddress{sdk.AccAddress(pubKey.Address())}

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc, err := ante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}

		// retrieve signBytes of tx

		// retrieve pubkey
		pubKey := acc.GetPubKey()
		if !simulate && pubKey == nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		// Check account sequence number.
		// When using Amino StdSignatures, we actually don't have the Sequence in
		// the SignatureV2 struct (it's only in the SignDoc). In this case, we
		// cannot check sequence directly, and must do it via signature
		// verification (in the VerifySignature call below).
		onlyAminoSigners := ante.OnlyLegacyAminoSigners(sig.Data)
		if !onlyAminoSigners {
			if sig.Sequence != acc.GetSequence() {
				return ctx, sdkerrors.Wrapf(
					sdkerrors.ErrWrongSequence,
					"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
				)
			}
		}

		// retrieve signer data
		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}
		signerData := authsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNum,
			Sequence:      acc.GetSequence(),
		}

		if !simulate {
			err := authsigning.VerifySignature(pubKey, signerData, sig.Data, svd.signModeHandler, tx)
			if err != nil {
				var errMsg string
				if onlyAminoSigners {
					// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
					// and therefore communicate sequence number as a potential cause of error.
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d), sequence (%d) and chain-id (%s)", accNum, acc.GetSequence(), chainID)
				} else {
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID)
				}
				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)

			}
		}
	}

	return next(ctx, tx, simulate)
}

// IncrementSequenceDecorator handles incrementing sequences of all signers.
// Use the IncrementSequenceDecorator decorator to prevent replay attacks. Note,
// there is no need to execute IncrementSequenceDecorator on RecheckTX since
// CheckTx would already bump the sequence number.
//
// NOTE: Since CheckTx and DeliverTx state are managed separately, subsequent and
// sequential txs orginating from the same account cannot be handled correctly in
// a reliable way unless sequence numbers are managed and tracked manually by a
// client. It is recommended to instead use multiple messages in a tx.
type IncrementSequenceDecorator struct {
	ak  keeper.AccountKeeper
	pkg PubKeyGetter
}

func NewIncrementSequenceDecorator(ak keeper.AccountKeeper, pkg PubKeyGetter) IncrementSequenceDecorator {
	return IncrementSequenceDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (isd IncrementSequenceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// all messages must be of type IxoMsg
	msg, ok := tx.GetMsgs()[0].(IxoMsg)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.IxoMsg")
	}

	// Get pubKey
	pubKey, err := isd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer
	signerAddrs := []sdk.AccAddress{sdk.AccAddress(pubKey.Address())}

	// increment sequence of all signers
	for _, addr := range signerAddrs {
		acc := isd.ak.GetAccount(ctx, addr)
		if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
			panic(err)
		}

		isd.ak.SetAccount(ctx, acc)
	}

	return next(ctx, tx, simulate)
}
