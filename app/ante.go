package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ibcante "github.com/cosmos/ibc-go/v4/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	entityante "github.com/ixofoundation/ixo-blockchain/x/entity/ante"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	iidante "github.com/ixofoundation/ixo-blockchain/x/iid/ante"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	AccountKeeper     authante.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	FeegrantKeeper    authante.FeegrantKeeper
	IidKeeper         iidkeeper.Keeper
	EntityKeeper      entitykeeper.Keeper
	wasmConfig        *wasmtypes.WasmConfig
	IBCKeeper         *ibckeeper.Keeper
	txCounterStoreKey sdk.StoreKey
	SignModeHandler   authsigning.SignModeHandler
	SigGasConsumer    func(meter sdk.GasMeter, sig txsigning.SignatureV2, params authtypes.Params) error
}

// IxoAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
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
		// outermost AnteDecorator. SetUpContext must be called first
		authante.NewSetUpContextDecorator(),
		// wasm ante handlers after setup context to enforce limits early
		wasmkeeper.NewLimitSimulationGasDecorator(options.wasmConfig.SimulationGasLimit),
		wasmkeeper.NewCountTXDecorator(options.txCounterStoreKey),
		authante.NewRejectExtensionOptionsDecorator(),
		authante.NewMempoolFeeDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewSetPubKeyDecorator(options.AccountKeeper),
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
		// custom ixo handlers
		iidante.NewIidResolutionDecorator(options.IidKeeper),
		entityante.NewBlockNftContractTransferForEntityDecorator(options.EntityKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
