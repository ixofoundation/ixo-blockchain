package ante

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"

	libixo "github.com/ixofoundation/ixo-blockchain/lib/ixo"

	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidutil "github.com/ixofoundation/ixo-blockchain/x/iid/util"
	projectkeeper "github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	// issuerante "github.com/allinbits/cosmos-cash/v3/x/issuer/ante"
	// issuerkeeper "github.com/allinbits/cosmos-cash/v3/x/issuer/keeper"
	// vcskeeper "github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/keeper"
)

func checkAllMsgs(tx sdk.Tx) (*types.MsgCreateProject, error) {
	msg, ok := tx.GetMsgs()[0].(*types.MsgCreateProject)
	if !ok {
		return nil, errors.New("msg must be MsgCreateProject")
	}

	if len(tx.GetMsgs()) > 1 {
		return nil, errors.New("transactions with a MsgCreateProject can only contain 1 MsgCreateProject")
	}

	return msg, nil
}

func pubKeyGetter(keeper projectkeeper.Keeper, iidKeeper iidkeeper.Keeper) libixo.PubKeyGetter {
	return func(ctx sdk.Context, msg libixo.IxoMsg) (pubKey cryptotypes.PubKey, err error) {

		// MsgCreateProject: pubkey from msg since project does not exist yet
		// MsgWithdrawFunds: signer is user DID, so get pubkey from did module
		// Other: signer is project DID, so get pubkey from project module

		var pubKeyEd25519 ed25519.PubKey
		switch msg := msg.(type) {
		case *types.MsgCreateProject:
			pubKeyEd25519.Key = base58.Decode(msg.GetPubKey())
		default:
			return nil, errors.New("pubkey not found")
		}
		return &pubKeyEd25519, nil
	}
}

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
	_, err = checkAllMsgs(tx)
	if err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, err.Error())
	}

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
	projectKeeper projectkeeper.Keeper
	ak            AccountKeeper
}

func NewSetPubKeyDecorator(projectKeeper projectkeeper.Keeper, ak AccountKeeper) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		projectKeeper: projectKeeper,
		ak:            ak,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	// message must be of type MsgCreateProject
	msg, err := checkAllMsgs(tx)
	if err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, err.Error())
	}

	// Get project pubKey
	pubKey, err := pubKeyGetter(spkd.projectKeeper, spkd.projectKeeper.IidKeeper)(ctx, msg)
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
	projectKeeper projectkeeper.Keeper
	ak            authante.AccountKeeper
	bk            bankkeeper.Keeper
	iidKeeper     iidkeeper.Keeper
}

func NewDeductFeeDecorator(projectKeeper projectkeeper.Keeper, ak authante.AccountKeeper, bk bankkeeper.Keeper,
	iidKeeper iidkeeper.Keeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		projectKeeper: projectKeeper,
		ak:            ak,
		bk:            bk,
		iidKeeper:     iidKeeper,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", authtypes.FeeCollectorName))
	}

	// all messages must be of type MsgCreateProject
	msg, err := checkAllMsgs(tx)
	if err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, err.Error())
	}

	// Get pubKey
	pubKey, err := pubKeyGetter(dfd.projectKeeper, dfd.projectKeeper.IidKeeper)(ctx, msg)
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
	expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
		libixo.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTotalFee)))
	if !feeTx.GetFee().IsEqual(expectedTotalFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")
	}

	// Calculate transaction fee and project funding
	transactionFee := sdk.NewCoins(sdk.NewCoin(
		libixo.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTransactionFee)))
	projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		// fetch fee payer account
		feePayerDidDoc, exists := dfd.iidKeeper.GetDidDocument(ctx, []byte(msg.SenderDid))
		if !exists {
			return ctx, errors.New("sender did does not exist")
		}

		account, err := iidutil.GetAccountForVerificationMethod(ctx, dfd.ak, feePayerDidDoc, feePayerDidDoc.Id)
		if err != nil {
			return ctx, err
		}

		feePayerAcc, err := ante.GetSignerAcc(ctx, dfd.ak, account.GetAddress())
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

// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              AccountKeeper
	signModeHandler authsigning.SignModeHandler
	projectKeeper   projectkeeper.Keeper
}

func NewSigVerificationDecorator(ak AccountKeeper, projectKeeper projectkeeper.Keeper, signModeHandler authsigning.SignModeHandler) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,
		projectKeeper:   projectKeeper,
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

	// message must be of type MsgCreateProject
	msg, err := checkAllMsgs(tx)
	if err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, err.Error())
	}

	// Get signer did pubKey
	pubKey, err := pubKeyGetter(svd.projectKeeper, svd.projectKeeper.IidKeeper)(ctx, msg)
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

/////----------------------------------------------

// // SetUpContextDecorator sets the GasMeter in the Context and wraps the next AnteHandler with a defer clause
// // to recover from any downstream OutOfGas panics in the AnteHandler chain to return an error with information
// // on gas provided and gas used.
// // CONTRACT: Must be first decorator in the chain
// // CONTRACT: Tx must implement GasTx interface
// type SetUpContextDecorator struct{}

// func NewSetUpContextDecorator() SetUpContextDecorator {
// 	return SetUpContextDecorator{}
// }

// func (sud SetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	// all transactions must implement GasTx
// 	gasTx, ok := tx.(ante.GasTx)
// 	if !ok {
// 		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
// 		// during runTx.
// 		newCtx = ante.SetGasMeter(simulate, ctx, 0)
// 		return newCtx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be GasTx")
// 	}

// 	// Addding of DID uses an infinite gas meter
// 	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

// 	// Decorator will catch an OutOfGasPanic caused in the next antehandler
// 	// AnteHandlers must have their own defer/recover in order for the BaseApp
// 	// to know how much gas was used! This is because the GasMeter is created in
// 	// the AnteHandler, but if it panics the context won't be set properly in
// 	// runTx's recover call.
// 	defer func() {
// 		if r := recover(); r != nil {
// 			switch rType := r.(type) {
// 			case sdk.ErrorOutOfGas:
// 				log := fmt.Sprintf(
// 					"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
// 					rType.Descriptor, gasTx.GetGas(), newCtx.GasMeter().GasConsumed())

// 				err = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, log)
// 			default:
// 				panic(r)
// 			}
// 		}
// 	}()

// 	return next(newCtx, tx, simulate)
// }

// // type SetPubKeyDecorator struct {
// // 	ak  keeper.AccountKeeper
// // 	pkg ixotypes.PubKeyGetter
// // }

// // func NewSetPubKeyDecorator(ak keeper.AccountKeeper, pkg ixotypes.PubKeyGetter) SetPubKeyDecorator {
// // 	return SetPubKeyDecorator{
// // 		ak:  ak,
// // 		pkg: pkg,
// // 	}
// // }

// // func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
// // 	_, ok := tx.(authsigning.SigVerifiableTx)
// // 	if !ok {
// // 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
// // 	}

// // 	// message must be of type MsgCreateProject
// // 	msg, ok := tx.GetMsgs()[0].(*types.MsgCreateProject)
// // 	if !ok {
// // 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
// // 	}

// // 	// Get project pubKey
// // 	pubKey, err := spkd.pkg(ctx, msg)
// // 	if err != nil {
// // 		return ctx, err
// // 	}

// // 	// Fetch signer (project itself). Account expected to not exist
// // 	signerAddr := sdk.AccAddress(pubKey.Address())
// // 	_, err = ante.GetSignerAcc(ctx, spkd.ak, signerAddr)
// // 	if err == nil {
// // 		return ctx, fmt.Errorf("expected project account to not exist")
// // 	}

// // 	// Create signer's account
// // 	signerAcc := spkd.ak.NewAccountWithAddress(ctx, signerAddr)
// // 	spkd.ak.SetAccount(ctx, signerAcc)

// // 	pubkeys := []cryptotypes.PubKey{pubKey}
// // 	signers := []sdk.AccAddress{signerAddr}

// // 	for i, pk := range pubkeys {
// // 		// PublicKey was omitted from slice since it has already been set in context
// // 		if pk == nil {
// // 			if !simulate {
// // 				continue
// // 			}
// // 			pk = &simEd25519Pubkey
// // 		}
// // 		// Only make check if simulate=false
// // 		if !simulate && !bytes.Equal(pk.Address(), signers[i]) {
// // 			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey,
// // 				"pubKey does not match signer address %s with signer index: %d", signers[i], i)
// // 		}

// // 		acc, err := ante.GetSignerAcc(ctx, spkd.ak, signers[i])
// // 		if err != nil {
// // 			return ctx, err
// // 		}
// // 		// account already has pubkey set,no need to reset
// // 		if acc.GetPubKey() != nil {
// // 			continue
// // 		}
// // 		err = acc.SetPubKey(pk)
// // 		if err != nil {
// // 			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
// // 		}
// // 		spkd.ak.SetAccount(ctx, acc)
// // 	}

// // 	return next(ctx, tx, simulate)
// // }

// // FundProjectDecorator deducts fees from the first signer of the tx
// // If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// // Call next AnteHandler if fees successfully deducted
// // CONTRACT: Tx must implement FeeTx interface to use FundProjectDecorator
// type FundProjectDecorator struct {
// 	iidKeeper      iidkeeper.Keeper
// 	accountKeeper  authante.AccountKeeper
// 	bankKeeper     bankkeeper.Keeper
// 	feegrantKeeper authante.FeegrantKeeper
// }

// func NewFundProjectDecorator(iidKeeper iidkeeper.Keeper, accountKeeper authante.AccountKeeper, bankKeeper bankkeeper.Keeper, feegrantKeeper authante.FeegrantKeeper) FundProjectDecorator {
// 	return FundProjectDecorator{
// 		iidKeeper:      iidKeeper,
// 		accountKeeper:  accountKeeper,
// 		bankKeeper:     bankKeeper,
// 		feegrantKeeper: feegrantKeeper,
// 	}
// }

// func (dec FundProjectDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	fmt.Println("executeing fund handler =============================================")

// 	feeTx, ok := tx.(libante.IxoFeeTx)
// 	if !ok {
// 		return next(ctx, tx, simulate)
// 	}

// 	ixoFeeMsgs := feeTx.GetFeePayerMsgs()

// 	var fundProjectMsgs []*types.MsgCreateProject
// 	for _, msg := range ixoFeeMsgs {
// 		if cpmsg, ok := msg.(*types.MsgCreateProject); ok {
// 			fundProjectMsgs = append(fundProjectMsgs, cpmsg)
// 		}
// 	}

// 	// If no create project messages are found skip this decorator
// 	fundProjectCount := len(fundProjectMsgs)
// 	if fundProjectCount == 0 {
// 		return next(ctx, tx, simulate)
// 	}

// 	feePayer, err := ixoFeeMsgs[0].FeePayerFromIid(ctx, dec.accountKeeper, dec.iidKeeper)
// 	if err != nil {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "fee payer does not exist")
// 	}

// 	// confirm that fee is the exact amount expected
// 	expectedTotalFee := sdk.NewCoins(sdk.NewCoin(libixo.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTotalFee)))
// 	if !feeTx.GetFee().IsEqual(expectedTotalFee) {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")
// 	}

// 	// Calculate transaction fee and project funding
// 	transactionFee := sdk.NewCoins(sdk.NewCoin(libixo.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTransactionFee).MulRaw(int64(fundProjectCount))))
// 	projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

// 	fee := feeTx.GetFee()
// 	feeGranter := feeTx.FeeGranter()

// 	deductFeesFrom := feePayer.GetFeePayerAccount().GetAddress()

// 	// if feegranter set deduct fee from feegranter account.
// 	// this works with only when feegrant enabled.
// 	if feeGranter != nil {
// 		if dec.feegrantKeeper == nil {
// 			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee grants are not enabled")
// 		} else if !feeGranter.Equals(feePayer.GetRecipientAddress()) {
// 			err := dec.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer.GetRecipientAddress(), fee, tx.GetMsgs())

// 			if err != nil {
// 				return ctx, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
// 			}
// 		}

// 		deductFeesFrom = feeGranter
// 	}

// 	deductFeesFromAcc := dec.accountKeeper.GetAccount(ctx, deductFeesFrom)
// 	if deductFeesFromAcc == nil {
// 		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
// 	}

// 	// deduct the fees
// 	if !feeTx.GetFee().IsZero() {
// 		for _, fundProjectMsg := range fundProjectMsgs {
// 			projectAddr, err := sdk.AccAddressFromBech32(fundProjectMsg.PubKey)
// 			if err != nil {
// 				return ctx, err
// 			}

// 			err = deductProjectFundingFees(dec.bankKeeper, ctx, deductFeesFromAcc, projectAddr, projectFunding)
// 			if err != nil {
// 				return ctx, err
// 			}
// 		}
// 	}

// 	events := sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
// 		sdk.NewAttribute(sdk.AttributeKeyFee, feeTx.GetFee().String()),
// 	)}
// 	ctx.EventManager().EmitEvents(events)

// 	return next(ctx, tx, simulate)
// }

// // Verify all signatures for a tx and return an error if any are invalid. Note,
// // the SigVerificationDecorator decorator will not get executed on ReCheck.
// //
// // CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// // CONTRACT: Tx must implement SigVerifiableTx interface
// // type SigVerificationDecorator struct {
// // 	ak              keeper.AccountKeeper
// // 	signModeHandler authsigning.SignModeHandler
// // 	pkg             ixotypes.PubKeyGetter
// // }

// // func NewSigVerificationDecorator(ak keeper.AccountKeeper, signModeHandler authsigning.SignModeHandler, pkg ixotypes.PubKeyGetter) SigVerificationDecorator {
// // 	return SigVerificationDecorator{
// // 		ak:              ak,
// // 		signModeHandler: signModeHandler,
// // 		pkg:             pkg,
// // 	}
// // }

// // func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// // 	// no need to verify signatures on recheck tx
// // 	if ctx.IsReCheckTx() {
// // 		return next(ctx, tx, simulate)
// // 	}
// // 	sigTx, ok := tx.(authsigning.SigVerifiableTx)
// // 	if !ok {
// // 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
// // 	}

// // 	// stdSigs contains the sequence number, account number, and signatures.
// // 	// When simulating, this would just be a 0-length slice.
// // 	sigs, err := sigTx.GetSignaturesV2()
// // 	if err != nil {
// // 		return ctx, err
// // 	}

// // 	// message must be of type MsgCreateProject
// // 	msg, ok := tx.GetMsgs()[0].(*types.MsgCreateProject)
// // 	if !ok {
// // 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
// // 	}

// // 	// Get signer did pubKey
// // 	pubKey, err := svd.pkg(ctx, msg)
// // 	if err != nil {
// // 		return ctx, err
// // 	}

// // 	// Fetch signer (account underlying DID). Account expected to not exist
// // 	signerAddrs := []sdk.AccAddress{sdk.AccAddress(pubKey.Address())}

// // 	// check that signer length and signature length are the same
// // 	if len(sigs) != len(signerAddrs) {
// // 		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
// // 	}

// // 	for i, sig := range sigs {
// // 		acc, err := ante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
// // 		if err != nil {
// // 			return ctx, err
// // 		}

// // 		// check signature, return account with incremented nonce
// // 		// retrieve pubkey
// // 		pubKey := acc.GetPubKey()
// // 		if !simulate && pubKey == nil {
// // 			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
// // 		}

// // 		// Check account sequence number.
// // 		// When using Amino StdSignatures, we actually don't have the Sequence in
// // 		// the SignatureV2 struct (it's only in the SignDoc). In this case, we
// // 		// cannot check sequence directly, and must do it via signature
// // 		// verification (in the VerifySignature call below).
// // 		onlyAminoSigners := ante.OnlyLegacyAminoSigners(sig.Data)
// // 		if !onlyAminoSigners {
// // 			if sig.Sequence != acc.GetSequence() {
// // 				return ctx, sdkerrors.Wrapf(
// // 					sdkerrors.ErrWrongSequence,
// // 					"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
// // 				)
// // 			}
// // 		}

// // 		genesis := ctx.BlockHeight() == 0
// // 		chainID := ctx.ChainID()
// // 		var accNum uint64
// // 		if !genesis {
// // 			// Fixed account number used so that sign bytes do not depend on it
// // 			accNum = uint64(0)
// // 		}

// // 		signerData := authsigning.SignerData{
// // 			ChainID:       chainID,
// // 			AccountNumber: accNum,
// // 			Sequence:      acc.GetSequence(),
// // 		}

// // 		if !simulate {
// // 			err := authsigning.VerifySignature(pubKey, signerData, sig.Data, svd.signModeHandler, tx)

// // 			if err != nil {
// // 				var errMsg string
// // 				if onlyAminoSigners {
// // 					// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
// // 					// and therefore communicate sequence number as a potential cause of error.
// // 					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d), sequence (%d) and chain-id (%s)", accNum, acc.GetSequence(), chainID)
// // 				} else {
// // 					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID)
// // 				}
// // 				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)

// // 			}
// // 		}
// // 	}

// // 	return next(ctx, tx, simulate)
// // }
