package project

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519.PubKeyEd25519
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, err error) {

		// Get signer PubKey
		var pubKeyEd25519 ed25519.PubKeyEd25519
		switch msg := msg.(type) {
		case MsgCreateProject:
			copy(pubKeyEd25519[:], base58.Decode(msg.GetPubKey()))
		case MsgWithdrawFunds:
			signerDid := msg.GetSignerDid()
			signerDoc, _ := didKeeper.GetDidDoc(ctx, signerDid)
			if signerDoc == nil {
				return pubKey, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "signer did not found")
			}
			copy(pubKeyEd25519[:], base58.Decode(signerDoc.GetPubKey()))
		default:
			// For the remaining messages, the project is the signer
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.GetSignerDid())
			if err != nil {
				return pubKey, sdkerrors.Wrap(did.ErrInvalidDid, "project did not found")
			}
			copy(pubKeyEd25519[:], base58.Decode(projectDoc.PubKey))
		}
		return pubKeyEd25519, nil
	}
}

// Identical to Cosmos DeductFees function, but tokens sent to project account
func deductProjectFundingFees(bankKeeper bank.Keeper, ctx sdk.Context,
	acc exported.Account, projectAddr sdk.AccAddress, fees sdk.Coins) error {
	blockTime := ctx.BlockHeader().Time
	coins := acc.GetCoins()

	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount %s", fees)
	}

	// verify the account has enough funds to pay for fees
	_, hasNeg := coins.SafeSub(fees)
	if hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees %s < %s", coins, fees)
	}

	// Validate the account has enough "spendable" coins as this will cover cases
	// such as vesting accounts.
	spendableCoins := acc.SpendableCoins(blockTime)
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees; %s < %s", spendableCoins, fees)
	}

	err := bankKeeper.SendCoins(ctx, acc.GetAddress(), projectAddr, fees)
	if err != nil {
		return err
	}

	return nil
}

func getProjectCreationSignBytes(chainID string, tx auth.StdTx, acc exported.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewProjectCreationAnteHandler(ak auth.AccountKeeper, supplyKeeper supply.Keeper,
	bk bank.Keeper, didKeeper did.Keeper, pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		//ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, supplyKeeper, bk, didKeeper, pubKeyGetter),
		//ixo.NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, pubKeyGetter),
		ixo.NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}
