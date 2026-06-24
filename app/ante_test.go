package app

import (
	"testing"

	"cosmossdk.io/log"
	wasm "github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	sims "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/lib/ixo"
)

// TestIxoAnteHandler_Wiring confirms the IxoAnteHandler constructor accepts
// the full HandlerOptions populated from a live IxoApp without erroring.
// This is the only "full chain" test we need: per-decorator behaviour is
// covered in x/iid/ante, x/entity/ante, x/smart-account/ante, and the
// decorator ordering is structural — if any of the keepers were nil or the
// ante helper signatures drifted, this test would fail.
func TestIxoAnteHandler_Wiring(t *testing.T) {
	a := NewIxoApp(
		log.NewNopLogger(), dbm.NewMemDB(), nil, true, map[int64]bool{},
		DefaultNodeHome, sims.EmptyAppOptions{}, EmptyWasmOpts,
		baseapp.SetChainID(TestChainID),
	)

	wasmConfig, err := wasm.ReadWasmConfig(sims.EmptyAppOptions{})
	require.NoError(t, err)

	handler, err := IxoAnteHandler(HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   a.AccountKeeper,
			BankKeeper:      a.BankKeeper,
			FeegrantKeeper:  a.FeeGrantKeeper,
			SignModeHandler: a.txConfig.SignModeHandler(),
			SigGasConsumer:  ixo.IxoSigVerificationGasConsumer,
		},
		IidKeeper:          a.IidKeeper,
		EntityKeeper:       a.EntityKeeper,
		WasmConfig:         wasmConfig,
		IBCKeeper:          a.IBCKeeper,
		TxCounterStoreKey:  runtime.NewKVStoreService(a.GetKey(wasmtypes.StoreKey)),
		appCodec:           a.appCodec,
		smartAccountKeeper: a.SmartAccountKeeper,
	})
	require.NoError(t, err)
	require.NotNil(t, handler, "constructed ante handler must be non-nil")
}

// TestIxoAnteHandler_MissingDeps asserts the constructor's nil-keeper guards
// fire — these prevent a misconfigured chain from booting with a partial
// ante chain that would silently skip signature checks etc.
func TestIxoAnteHandler_MissingDeps(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(opts *HandlerOptions)
		errSub string
	}{
		{
			name:   "nil AccountKeeper",
			mutate: func(o *HandlerOptions) { o.AccountKeeper = nil },
			errSub: "account keeper is required",
		},
		{
			name:   "nil BankKeeper",
			mutate: func(o *HandlerOptions) { o.BankKeeper = nil },
			errSub: "bank keeper is required",
		},
		{
			name:   "nil SignModeHandler",
			mutate: func(o *HandlerOptions) { o.SignModeHandler = nil },
			errSub: "sign mode handler is required",
		},
		{
			name:   "nil TxCounterStoreKey",
			mutate: func(o *HandlerOptions) { o.TxCounterStoreKey = nil },
			errSub: "tx counter key is required",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := NewIxoApp(
				log.NewNopLogger(), dbm.NewMemDB(), nil, true, map[int64]bool{},
				DefaultNodeHome, sims.EmptyAppOptions{}, EmptyWasmOpts,
				baseapp.SetChainID(TestChainID),
			)
			wasmConfig, err := wasm.ReadWasmConfig(sims.EmptyAppOptions{})
			require.NoError(t, err)

			opts := HandlerOptions{
				HandlerOptions: ante.HandlerOptions{
					AccountKeeper:   a.AccountKeeper,
					BankKeeper:      a.BankKeeper,
					FeegrantKeeper:  a.FeeGrantKeeper,
					SignModeHandler: a.txConfig.SignModeHandler(),
					SigGasConsumer:  ixo.IxoSigVerificationGasConsumer,
				},
				IidKeeper:          a.IidKeeper,
				EntityKeeper:       a.EntityKeeper,
				WasmConfig:         wasmConfig,
				IBCKeeper:          a.IBCKeeper,
				TxCounterStoreKey:  runtime.NewKVStoreService(a.GetKey(wasmtypes.StoreKey)),
				appCodec:           a.appCodec,
				smartAccountKeeper: a.SmartAccountKeeper,
			}
			tc.mutate(&opts)

			_, err = IxoAnteHandler(opts)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.errSub)
		})
	}
}
