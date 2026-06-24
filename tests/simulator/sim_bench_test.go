package simulator

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/ixofoundation/ixo-blockchain/v8/app"
)

// BenchmarkFullAppSimulation profiles SimulateFromSeed against the IxoApp.
// Run with `go test -benchmem -run=^$ -bench=^BenchmarkFullAppSimulation
// ./tests/simulator/...` plus the standard simcli flags. Useful for tracking
// per-block latency regressions when the keeper layer changes.
func BenchmarkFullAppSimulation(b *testing.B) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID
	config.Commit = true

	db := dbm.NewMemDB()
	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "ixod-bench-sim")
	require.NoError(b, err)
	b.Cleanup(func() { _ = os.RemoveAll(dir) })

	appOptions := make(simtestutil.AppOptionsMap)
	appOptions[flags.FlagHome] = dir
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	a := app.NewIxoApp(
		logger, db, nil, true, map[int64]bool{}, dir,
		appOptions, app.EmptyWasmOpts,
		fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, simErr := simulation.SimulateFromSeed(
			b, os.Stdout, a.BaseApp,
			simtestutil.AppStateFn(a.AppCodec(), a.SimulationManager(), a.DefaultGenesis()),
			simtypes.RandomAccounts,
			simtestutil.SimulationOperations(a, a.AppCodec(), config),
			a.BlockedAddresses(),
			config,
			a.AppCodec(),
		)
		require.NoError(b, simErr)
	}
}
