package project

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	//"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
)

// SetUpContextDecorator sets the GasMeter in the Context and wraps the next AnteHandler with a defer clause
// to recover from any downstream OutOfGas panics in the AnteHandler chain to return an error with information
// on gas provided and gas used.
// CONTRACT: Must be first decorator in the chain
// CONTRACT: Tx must implement GasTx interface
type SetUpContextDecorator struct{}

func NewSetUpContextDecorator() SetUpContextDecorator {
	return SetUpContextDecorator{}
}

func (sud SetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must implement GasTx
	gasTx, ok := tx.(ante.GasTx)
	if !ok {
		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
		// during runTx.
		newCtx = ante.SetGasMeter(simulate, ctx, 0)
		return newCtx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be GasTx")
	}

	// Addding of DID uses an infinite gas meter
	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	// Decorator will catch an OutOfGasPanic caused in the next antehandler
	// AnteHandlers must have their own defer/recover in order for the BaseApp
	// to know how much gas was used! This is because the GasMeter is created in
	// the AnteHandler, but if it panics the context won't be set properly in
	// runTx's recover call.
	defer func() {
		if r := recover(); r != nil {
			switch rType := r.(type) {
			case sdk.ErrorOutOfGas:
				log := fmt.Sprintf(
					"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
					rType.Descriptor, gasTx.GetGas(), newCtx.GasMeter().GasConsumed())

				err = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, log)
			default:
				panic(r)
			}
		}
	}()

	return next(newCtx, tx, simulate)
}

type SetPubKeyDecorator struct {
	ak  keeper.AccountKeeper
	pkg ixo.PubKeyGetter
}

func NewSetPubKeyDecorator(ak keeper.AccountKeeper, pkg ixo.PubKeyGetter) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(authsigning.SigVerifiableTx)//(ante.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	// message must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
	}

	// Get project pubKey
	pubKey, err := spkd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// Fetch signer (project itself). Account expected to not exist
	signerAddr := sdk.AccAddress(pubKey.Address())
	_, err = ante.GetSignerAcc(ctx, spkd.ak, signerAddr)
	if err == nil {
		return ctx, fmt.Errorf("expected project account to not exist")
	}

	// Create signer's account
	signerAcc := spkd.ak.NewAccountWithAddress(ctx, signerAddr)
	spkd.ak.SetAccount(ctx, signerAcc)

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
	//supplyKeeper authtypes.SupplyKeeper
	bk           bankkeeper.Keeper
	didKeeper    did.Keeper
	pkg          ixo.PubKeyGetter
}

func NewDeductFeeDecorator(ak keeper.AccountKeeper, bk bankkeeper.Keeper,
	didKeeper did.Keeper, pkg ixo.PubKeyGetter) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:           ak,
		bk:           bk,
		didKeeper:    didKeeper,
		pkg:          pkg,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", authtypes.FeeCollectorName))
	}

	// all messages must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
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

	// confirm that fee is the exact amount expected
	stdTx, ok := tx.(legacytx.StdTx) //(auth.StdTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTotalFee)))
	if !stdTx.Fee.Amount.IsEqual(expectedTotalFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")
	}

	// Calculate transaction fee and project funding
	transactionFee := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTransactionFee)))
	projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

	// deduct the fees
	if !stdTx.Fee.Amount.IsZero() {
		// fetch fee payer account
		feePayerDidDoc, err := dfd.didKeeper.GetDidDoc(ctx, msg.SenderDid)
		if err != nil {
			return ctx, err
		}
		feePayerAcc, err := ante.GetSignerAcc(ctx, dfd.ak, feePayerDidDoc.Address())
		if err != nil {
			return ctx, err
		}

		err = ante.DeductFees(dfd.bk, ctx, feePayerAcc, transactionFee)
		if err != nil {
			return ctx, err
		}

		projectAddr := sdk.AccAddress(pubKey.Address())
		err = deductProjectFundingFees(dfd.bk, ctx, feePayerAcc, projectAddr, projectFunding)
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
	pkg 			ixo.PubKeyGetter
}

func NewSigVerificationDecorator(ak keeper.AccountKeeper, signModeHandler authsigning.SignModeHandler, pkg ixo.PubKeyGetter) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:  ak,
		signModeHandler: signModeHandler,
		pkg: pkg,
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
	if (err != nil) {
		return ctx, err
	}

	// message must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
	}

	// Get signer did pubKey
	pubKey, err := svd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// Fetch signer (account underlying DID). Account expected to not exist
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

		// check signature, return account with incremented nonce
		//stdTx, ok := tx.(legacytx.StdTx)//(auth.StdTx)
		//if !ok {
		//	return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
		//}

		/// retrieve pubkey
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

		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			// Fixed account number used so that sign bytes do not depend on it
			accNum = uint64(0)
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

//TODO Option 2: Uses legacy functions
//func (svd SigVerificationDecorator) AnteHandle2(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
//	// no need to verify signatures on recheck tx
//	if ctx.IsReCheckTx() {
//		return next(ctx, tx, simulate)
//	}
//	sigTx, ok := tx.(authsigning.SigVerifiableTx)
//	if !ok {
//		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
//	}
//
//	// stdSigs contains the sequence number, account number, and signatures.
//	// When simulating, this would just be a 0-length slice.
//	sigs, err := sigTx.GetSignaturesV2()
//	if (err != nil) {
//		return ctx, err
//	}
//
//	// message must be of type MsgCreateProject
//	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
//	if !ok {
//		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
//	}
//
//	// Get signer did pubKey
//	pubKey, err := svd.pkg(ctx, msg)
//	if err != nil {
//		return ctx, err
//	}
//
//	// Fetch signer (account underlying DID). Account expected to not exist
//	signerAddrs := []sdk.AccAddress{sdk.AccAddress(pubKey.Address())}
//
//	// check that signer length and signature length are the same
//	if len(sigs) != len(signerAddrs) {
//		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
//	}
//
//	for i, sig := range sigs {
//		acc, err := ante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
//		if err != nil {
//			return ctx, err
//		}
//
//		// check signature, return account with incremented nonce
//		stdTx, ok := tx.(legacytx.StdTx) //(auth.StdTx)
//		if !ok {
//			return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
//		}
//
//		/// retrieve pubkey
//		pubKey := acc.GetPubKey()
//		if !simulate && pubKey == nil {
//			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
//		}
//
//		genesis := ctx.BlockHeight() == 0
//		chainID := ctx.ChainID()
//		var accNum uint64
//		if !genesis {
//			// Fixed account number used so that sign bytes do not depend on it
//			accNum = uint64(0)
//		}
//
//		if !simulate {
//			err := VerifySignatureProjectCreation(chainID, accNum, acc.GetSequence(), stdTx, pubKey, sig.Data)
//			if err != nil {
//				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID))
//			}
//		}
//	}
//
//	return next(ctx, tx, simulate)
//}
//
//// VerifySignature verifies a transaction signature contained in SignatureData abstracting over different signing modes
//// and single vs multi-signatures.
//func VerifySignatureProjectCreation(
//	chainID string, accNum uint64, sequence uint64, tx legacytx.StdTx, pubKey crypto.PubKey, sigData signing.SignatureData,
//	) error {
//	switch data := sigData.(type) {
//	case *signing.SingleSignatureData:
//		signBytes := legacytx.StdSignBytes(chainID, accNum, sequence, 0, tx.Fee, tx.Msgs, tx.Memo)
//		if !pubKey.VerifySignature(signBytes, data.Signature) {
//			return fmt.Errorf("unable to verify single signer signature")
//		}
//		return nil
//
//	case *signing.MultiSignatureData:
//		multiPK, ok := pubKey.(multisig.PubKey)
//		if !ok {
//			return fmt.Errorf("expected %T, got %T", (multisig.PubKey)(nil), pubKey)
//		}
//		err := multiPK.VerifyMultisignature(func(mode signing.SignMode) ([]byte, error) {
//			return legacytx.StdSignBytes(chainID, accNum, sequence, 0, tx.Fee, tx.Msgs, tx.Memo), nil
//		}, data)
//		if err != nil {
//			return err
//		}
//		return nil
//	default:
//		return fmt.Errorf("unexpected SignatureData %T", sigData)
//	}
//}