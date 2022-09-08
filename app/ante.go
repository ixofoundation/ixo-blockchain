package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	didkeeper "github.com/allinbits/cosmos-cash/v3/x/did/keeper"
	issuerante "github.com/allinbits/cosmos-cash/v3/x/issuer/ante"
	issuerkeeper "github.com/allinbits/cosmos-cash/v3/x/issuer/keeper"
	vcskeeper "github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/keeper"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func IxoAnteHandler(
	ak authante.AccountKeeper,
	bankKeeper authtypes.BankKeeper,
	feeGrantKeeper authante.FeegrantKeeper,
	ik issuerkeeper.Keeper,
	dk didkeeper.Keeper,
	vcsk vcskeeper.Keeper,
	sigGasConsumer authante.SignatureVerificationGasConsumer,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		authante.NewRejectExtensionOptionsDecorator(),
		authante.NewMempoolFeeDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.TxTimeoutHeightDecorator{},
		authante.NewValidateMemoDecorator(ak),
		authante.NewConsumeGasForTxSizeDecorator(ak),
		authante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(ak),
		authante.NewDeductFeeDecorator(ak, bankKeeper, feeGrantKeeper),
		authante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		authante.NewSigVerificationDecorator(ak, signModeHandler),
		authante.NewIncrementSequenceDecorator(ak),
		issuerante.NewCheckUserCredentialsDecorator(ak, ik, dk, vcsk),
	)
}
