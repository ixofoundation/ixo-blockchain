package types

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/crypto"
)

type SetPubKeyDecorator struct {
	ak  keeper.AccountKeeper
	pkg PubKeyGetter
}

func NewSetPubKeyDecorator(ak keeper.AccountKeeper, pkg PubKeyGetter) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(ante.SigVerifiableTx)
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

	pubkeys := []crypto.PubKey{pubKey}
	signers := []sdk.AccAddress{signerAddr}

	for i, pk := range pubkeys {
		// PublicKey was omitted from slice since it has already been set in context
		if pk == nil {
			if !simulate {
				continue
			}
			pk = simEd25519Pubkey
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
	ak           keeper.AccountKeeper
	supplyKeeper types.SupplyKeeper
	pkg          PubKeyGetter
}

func NewDeductFeeDecorator(ak keeper.AccountKeeper, sk types.SupplyKeeper, pkg PubKeyGetter) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:           ak,
		supplyKeeper: sk,
		pkg:          pkg,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(ante.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.supplyKeeper.GetModuleAddress(types.FeeCollectorName); addr == nil {
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
		err = ante.DeductFees(dfd.supplyKeeper, ctx, feePayerAcc, feeTx.GetFee())
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
	sigTx, ok := tx.(ante.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	params := sgcd.ak.GetParams(ctx)
	sigs := sigTx.GetSignatures()

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

		if simulate && pubKey == nil {
			// In simulate mode the transaction comes with no signatures, thus if the
			// account's pubkey is nil, both signature verification and gasKVStore.Set()
			// shall consume the largest amount, i.e. it takes more gas to verify
			// secp256k1 keys than ed25519 ones.
			if pubKey == nil {
				pubKey = simEd25519Pubkey
			}
		}
		err = sgcd.sigGasConsumer(ctx.GasMeter(), sig, pubKey, params)
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
	ak  keeper.AccountKeeper
	pkg PubKeyGetter
}

func NewSigVerificationDecorator(ak keeper.AccountKeeper, pkg PubKeyGetter) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}
	sigTx, ok := tx.(ante.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs := sigTx.GetSignatures()

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
	signerAccs := make([]authexported.Account, len(signerAddrs))

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		signerAccs[i], err = ante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}

		// retrieve signBytes of tx
		signBytes := sigTx.GetSignBytes(ctx, signerAccs[i])

		// retrieve pubkey
		pubKey := signerAccs[i].GetPubKey()
		if !simulate && pubKey == nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		// verify signature
		if !simulate && !pubKey.VerifyBytes(signBytes, sig) {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "signature verification failed; verify correct account sequence and chain-id")
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
	_, ok := tx.(ante.SigVerifiableTx)
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
