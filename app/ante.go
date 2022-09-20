package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	iidante "github.com/ixofoundation/ixo-blockchain/x/iid/ante"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"

	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	// issuerante "github.com/allinbits/cosmos-cash/v3/x/issuer/ante"
	// issuerkeeper "github.com/allinbits/cosmos-cash/v3/x/issuer/keeper"
	// vcskeeper "github.com/allinbits/cosmos-cash/v3/x/verifiable-credential/keeper"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
// func IxoAnteHandler(
// 	ak authante.AccountKeeper,
// 	bankKeeper authtypes.BankKeeper,
// 	feeGrantKeeper authante.FeegrantKeeper,
// 	ik issuerkeeper.Keeper,
// 	dk didkeeper.Keeper,
// 	vcsk vcskeeper.Keeper,
// 	sigGasConsumer authante.SignatureVerificationGasConsumer,
// 	signModeHandler signing.SignModeHandler,
// ) sdk.AnteHandler {
// 	return sdk.ChainAnteDecorators(
// 		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
// 		authante.NewRejectExtensionOptionsDecorator(),
// 		authante.NewMempoolFeeDecorator(),
// 		authante.NewValidateBasicDecorator(),
// 		authante.TxTimeoutHeightDecorator{},
// 		authante.NewValidateMemoDecorator(ak),
// 		authante.NewConsumeGasForTxSizeDecorator(ak),
// 		authante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
// 		authante.NewValidateSigCountDecorator(ak),
// 		authante.NewDeductFeeDecorator(ak, bankKeeper, feeGrantKeeper),
// 		authante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
// 		authante.NewSigVerificationDecorator(ak, signModeHandler),
// 		authante.NewIncrementSequenceDecorator(ak),
// 		issuerante.NewCheckUserCredentialsDecorator(ak, ik, dk, vcsk),
// 	)
// }

type HandlerOptions struct {
	AccountKeeper   authante.AccountKeeper
	BankKeeper      bankkeeper.Keeper
	FeegrantKeeper  authante.FeegrantKeeper
	IidKeeper       iidkeeper.Keeper
	SignModeHandler authsigning.SignModeHandler
	SigGasConsumer  func(meter sdk.GasMeter, sig txsigning.SignatureV2, params authtypes.Params) error
}

func IxoAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	var sigGasConsumer = options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		authante.NewRejectExtensionOptionsDecorator(),
		authante.NewMempoolFeeDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		authante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
		iidante.NewIidResolutionDecorator(options.IidKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
