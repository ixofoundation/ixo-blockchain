package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ixoante "github.com/ixofoundation/ixo-blockchain/lib/ante"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	projectante "github.com/ixofoundation/ixo-blockchain/x/project/ante"
	projectkeeper "github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	projecttypes "github.com/ixofoundation/ixo-blockchain/x/project/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
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
	AccountKeeper     projectante.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	FeegrantKeeper    authante.FeegrantKeeper
	IidKeeper         iidkeeper.Keeper
	ProjectKeeper     projectkeeper.Keeper
	wasmConfig        wasmtypes.WasmConfig
	txCounterStoreKey sdk.StoreKey
	SignModeHandler   authsigning.SignModeHandler
	SigGasConsumer    func(meter sdk.GasMeter, sig txsigning.SignatureV2, params authtypes.Params) error
}

func checkForCreateProjectMessages(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if _, ok := msg.(*projecttypes.MsgCreateProject); ok {
			return true
		}
	}
	return false
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

	// var sigGasConsumer = options.SigGasConsumer
	// if sigGasConsumer == nil {
	// 	sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	// }

	handler := func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		var anteDecorators []sdk.AnteDecorator

		if checkForCreateProjectMessages(tx) {
			// Handles the custom AnteHandlers necessary for maintaining backwords capability with creating projects via cellnode.
			// NOTE: PLEASE REMOVE THIS AT SOME POINT IN THE FUTURE DUE TO POSSIBLE SECURITY ISSUES AND POOR DESIGN.
			// NOTE: REFER TO THIS FILE FOR MORE INFORMATION: app/ante.md
			anteDecorators = []sdk.AnteDecorator{
				projectante.NewSetUpContextDecorator(),
				ixoante.NewCheckTxForIncompatibleMsgsDecorator(), // outermost AnteDecorator. SetUpContext must be called first
				//ante.NewMempoolFeeDecorator(),
				authante.NewValidateBasicDecorator(),
				authante.NewValidateMemoDecorator(options.AccountKeeper),
				authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
				projectante.NewSetPubKeyDecorator(options.ProjectKeeper, options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
				ante.NewValidateSigCountDecorator(options.AccountKeeper),
				projectante.NewDeductFeeDecorator(options.ProjectKeeper, options.AccountKeeper, options.BankKeeper, options.IidKeeper),
				//ixo.NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
				projectante.NewSigVerificationDecorator(options.AccountKeeper, options.ProjectKeeper, options.SignModeHandler),
				authante.NewIncrementSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator
			}
		} else {
			anteDecorators = []sdk.AnteDecorator{
				authante.NewSetUpContextDecorator(),
				ixoante.NewCheckTxForIncompatibleMsgsDecorator(),
				wasmkeeper.NewLimitSimulationGasDecorator(options.wasmConfig.SimulationGasLimit), // after setup context to enforce limits early
				wasmkeeper.NewCountTXDecorator(options.txCounterStoreKey),
				authante.NewRejectExtensionOptionsDecorator(),
				authante.NewMempoolFeeDecorator(),
				authante.NewValidateBasicDecorator(),
				authante.NewTxTimeoutHeightDecorator(),
				authante.NewValidateMemoDecorator(options.AccountKeeper),
				authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
				authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
				authante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
				authante.NewValidateSigCountDecorator(options.AccountKeeper),
				authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
				authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
				authante.NewIncrementSequenceDecorator(options.AccountKeeper),
				// iidante.NewIidResolutionDecorator(options.IidKeeper),
			}
		}

		return sdk.ChainAnteDecorators(anteDecorators...)(ctx, tx, simulate)
	}

	return handler, nil
}
