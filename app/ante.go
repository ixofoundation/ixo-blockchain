package app

import (
	corestoretypes "cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	entityante "github.com/ixofoundation/ixo-blockchain/v3/x/entity/ante"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/v3/x/entity/keeper"
	iidante "github.com/ixofoundation/ixo-blockchain/v3/x/iid/ante"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v3/x/iid/keeper"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	authante.HandlerOptions

	IidKeeper         iidkeeper.Keeper
	EntityKeeper      entitykeeper.Keeper
	WasmConfig        wasmtypes.WasmConfig
	IBCKeeper         *ibckeeper.Keeper
	TxCounterStoreKey corestoretypes.KVStoreService
}

// IxoAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func IxoAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.TxCounterStoreKey == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "tx counter key is required for ante builder")
	}

	var sigGasConsumer = options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		// outermost AnteDecorator. SetUpContext must be called first
		authante.NewSetUpContextDecorator(),
		// wasm ante handlers after setup context to enforce limits early
		wasmkeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit),
		wasmkeeper.NewCountTXDecorator(options.TxCounterStoreKey),
		// standard SDK AnteDecorators
		// TODO add circuit breaker everywhere
		// circuitante.NewCircuitBreakerDecorator(options.CircuitKeeper),
		authante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewSetPubKeyDecorator(options.AccountKeeper),
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		// custom ixo handlers
		iidante.NewIidResolutionDecorator(options.IidKeeper),
		entityante.NewBlockNftContractTransferForEntityDecorator(options.EntityKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

// NewIxoAnteHandler returns a new sdk.AnteHandler or panics if fail to create.
func NewIxoAnteHandler(options HandlerOptions) sdk.AnteHandler {
	ixoAnteHandler, err := IxoAnteHandler(options)
	if err != nil {
		panic(err)
	}
	return ixoAnteHandler
}
