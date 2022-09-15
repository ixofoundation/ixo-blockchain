package project

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	didtypes "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	did "github.com/ixofoundation/ixo-blockchain/x/iid"
	didkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

var (
	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519.PubKey
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	simEd25519Pubkey.Key = bz
}

func NewDefaultPubKeyGetter(keeper keeper.Keeper) ixotypes.PubKeyGetter {
	return func(ctx sdk.Context, msg ixotypes.IxoMsg) (pubKey cryptotypes.PubKey, err error) {

		projectDidDoc, err := keeper.GetProjectDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID not found")
		}

		var pubKeyRaw ed25519.PubKey
		pubKeyRaw.Key = base58.Decode(projectDidDoc.PubKey)
		return &pubKeyRaw, nil
	}
}

func NewModulePubKeyGetter(keeper keeper.Keeper, didKeeper didkeeper.Keeper) ixotypes.PubKeyGetter {
	return func(ctx sdk.Context, msg ixotypes.IxoMsg) (pubKey cryptotypes.PubKey, err error) {

		// MsgCreateProject: pubkey from msg since project does not exist yet
		// MsgWithdrawFunds: signer is user DID, so get pubkey from did module
		// Other: signer is project DID, so get pubkey from project module

		var pubKeyEd25519 ed25519.PubKey
		switch msg := msg.(type) {
		case *types.MsgCreateProject:
			pubKeyEd25519.Key = base58.Decode(msg.GetPubKey())
		case *types.MsgWithdrawFunds:
			return did.NewDefaultPubKeyGetter(didKeeper)(ctx, msg)
		default:
			return NewDefaultPubKeyGetter(keeper)(ctx, msg)
		}
		return &pubKeyEd25519, nil
	}
}

// Identical to Cosmos DeductFees function, but tokens sent to project account
func deductProjectFundingFees(bankKeeper bankkeeper.Keeper, ctx sdk.Context,
	acc authtypes.AccountI, projectAddr sdk.AccAddress, fees sdk.Coins) error {
	coins := bankKeeper.GetAllBalances(ctx, acc.GetAddress())

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
	spendableCoins := bankKeeper.SpendableCoins(ctx, acc.GetAddress())
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees; %s < %s", spendableCoins, fees)
	}

	err := bankKeeper.SendCoins(ctx, acc.GetAddress(), projectAddr, fees)
	if err != nil {
		return err
	}

	return nil
}

func getProjectCreationSignBytes(ctx sdk.Context, tx legacytx.StdTx, acc authtypes.AccountI) []byte {
	genesis := ctx.BlockHeight() == 0
	chainID := ctx.ChainID()
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return legacytx.StdSignBytes(
		chainID, accNum, acc.GetSequence(), 0, tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewProjectCreationAnteHandler(
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, didKeeper didkeeper.Keeper, signModeHandler signing.SignModeHandler,
	pubKeyGetter ixotypes.PubKeyGetter) sdk.AnteHandler {

	// Refer to inline documentation in app/app.go for introduction to why we
	// need a custom ixo AnteHandler, and especially a custom AnteHandler for
	// project creation. Below, we will discuss the differences between the
	// custom ixo AnteHandler and the project creation AnteHandler.
	//
	// It is clear below that our custom AnteHandler is not completely custom.
	// It uses various functions from the Cosmos ante module. However, it also
	// uses customised decorators, disables some decorators,
	//
	// In general:
	// - Sometimes enforces messages to be of type MsgCreateProject, especially
	//   if the decorator specifically needs to use the project creator DID.
	//
	// NewSetUpContextDecorator:
	// (Note: default ixo AnteHandler uses the Cosmos NewSetUpContextDecorator)
	// - Uses an infinite gas meter since we do not care about gas limits. This
	//   reduces the likelihood that a project creation message fails.
	//
	// NewMempoolFeeDecorator [[DISABLED]]:
	// - Disabled since we do not need to check that the provided fees meet a
	//   minimum threshold for the validator, given that we use a fixed fee.
	//
	// NewSetPubKeyDecorator:
	// - Enforces that the signer's account (the project) does not exist yet.
	// - Creates the signer's account (in the default ixo AnteHandler, this is
	//   only done if the signer does not yet exist, such as during MsgAddDid)
	//
	// NewDeductFeeDecorator:
	// - Enforces and charges a fixed MsgCreateProjectTotalFee instead of using
	//   the fee from the signed tx. This total fee is partly a transaction
	//   fee and partly funding for the project, so that it can sign future
	//   transactions (and pay gas fees) independently for a number of txs.
	// - Deducts any fees from the project creator rather than the message
	//   signer, since the message signer is actually the project.
	//
	// NewSigGasConsumeDecorator [[DISABLED]]:
	// - Similar to NewSetUpContextDecorator, we do not care about gas limits,
	//   so we do not need to consume signature-related gas.
	//
	// NewSigVerificationDecorator
	// - Project creation sign bytes are different from standard StdTx sign
	//   bytes, so one of this decorator's jobs is to construct this different
	//   sign bytes (difference discussed in next points) so that it is then
	//   able to verify the sign bytes correctly.
	// - The account number in project creation sign bytes is 0, because when
	//   the transaction is being signed, the project's account does not exist
	//   yet, so we cannot know what the account number will be. As another
	//   example, when signing a MsgAddDid, we do know the account number
	//   because we expect the account underlying the DID to have been created.
	//   Account creation typically happens by someone sending tokens to it.

	return sdk.ChainAnteDecorators(
		NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		//ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, bk, didKeeper, pubKeyGetter),
		//ixo.NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, signModeHandler, pubKeyGetter),
		ixotypes.NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}
