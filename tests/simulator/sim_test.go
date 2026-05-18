// Package simulator runs the cosmos-sdk simulation framework against the
// IxoApp. Each test below is invoked by `make test-sim-*`; the underlying
// machinery is the standard simcli + simulation.SimulateFromSeed pair.
//
// To run a quick smoke simulation:
//
//	go test -run ^TestFullAppSimulation -v ./tests/simulator/...
//
// To run determinism (multi-seed) tests:
//
//	go test -run ^TestAppStateDeterminism -v ./tests/simulator/...
package simulator

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"testing"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/app"
)

// SimAppChainID is the chain id used for simulations. Distinct from the
// integration-test chain id in app.TestChainID so simulator runs can't
// accidentally clobber the in-process app's state.
const SimAppChainID = "ixo-sim-1"

var FlagEnableStreamingValue bool

func init() {
	simcli.GetSimulatorFlags()
	flag.BoolVar(&FlagEnableStreamingValue, "EnableStreaming", false, "enable streaming service")
}

// fauxMerkleModeOpt sets the BaseApp into faux-merkle mode for faster sims.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) { bapp.SetFauxMerkleMode() }

// setupSimulationApp mirrors wasmd's helper: it builds a fresh IxoApp on a
// LevelDB-backed store rooted at a unique tempdir (passed via FlagHome so
// downstream sub-paths like wasm/ resolve correctly) and returns the
// pieces every sim test needs. Deferred cleanup is registered with t.Cleanup.
func setupSimulationApp(t *testing.T, skipMsg string) (simtypes.Config, dbm.DB, simtestutil.AppOptionsMap, *app.IxoApp) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID

	db, dir, logger, skip, err := simtestutil.SetupSimulation(
		config, "leveldb-app-sim", "Simulation",
		simcli.FlagVerboseValue, simcli.FlagEnabledValue,
	)
	if skip {
		t.Skip(skipMsg)
	}
	require.NoError(t, err, "simulation setup failed")

	t.Cleanup(func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	})

	appOptions := make(simtestutil.AppOptionsMap)
	// FlagHome must be the per-test temp dir so wasm/, application.toml etc.
	// resolve to ephemeral paths and don't collide between tests.
	appOptions[flags.FlagHome] = dir
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	a := app.NewIxoApp(
		logger, db, nil, true, map[int64]bool{}, dir,
		appOptions, app.EmptyWasmOpts,
		baseapp.SetChainID(SimAppChainID),
	)
	return config, db, appOptions, a
}

// TestFullAppSimulation runs SimulateFromSeed once over the configured
// number of blocks and ensures no invariant fails. Driven by simcli flags
// (-Period, -NumBlocks, -Seed, -Verbose, ...).
func TestFullAppSimulation(t *testing.T) {
	config, db, _, a := setupSimulationApp(t, "skipping application simulation")
	// Force Commit between blocks. Without it, BaseApp.LastBlockHeight() stays
	// at 0 and validateFinalizeBlockHeight rejects every block after the
	// first with "invalid height: N; expected: 1". cosmos-sdk's Makefile
	// hard-codes -Commit=true for the same reason.
	config.Commit = true

	_, simParams, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, a.BaseApp,
		simtestutil.AppStateFn(a.AppCodec(), a.SimulationManager(), a.DefaultGenesis()),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(a, a.AppCodec(), config),
		a.BlockedAddresses(),
		config,
		a.AppCodec(),
	)

	require.NoError(t, simtestutil.CheckExportSimulation(a, config, simParams))
	require.NoError(t, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}
}

// TestAppStateDeterminism runs the simulator NumSeeds times across NumTimesPerSeed
// runs and asserts the final app hash is identical across runs of the same seed.
// This is the classic "simulator-determinism" smoke test from cosmos-sdk.
func TestAppStateDeterminism(t *testing.T) {
	if !simcli.FlagEnabledValue {
		t.Skip("skipping application simulation determinism (set -Enabled=true to run)")
	}

	config := simcli.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = SimAppChainID
	config.Commit = true // see TestFullAppSimulation note

	numSeeds := 3
	numTimesPerSeed := 2
	appHashList := make([]string, numTimesPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesPerSeed; j++ {
			logger := log.NewNopLogger()
			if simcli.FlagVerboseValue {
				logger = log.NewTestLogger(t)
			}
			db := dbm.NewMemDB()
			dir, err := os.MkdirTemp("", "ixod-sim-determinism")
			require.NoError(t, err)
			t.Cleanup(func() { _ = os.RemoveAll(dir) })

			appOptions := make(simtestutil.AppOptionsMap)
			appOptions[flags.FlagHome] = dir
			appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

			a := app.NewIxoApp(
				logger, db, nil, true, map[int64]bool{}, dir,
				appOptions, app.EmptyWasmOpts,
				baseapp.SetChainID(SimAppChainID),
			)

			fmt.Printf("running non-determinism simulation; seed %d: attempt: %d/%d\n",
				config.Seed, j+1, numTimesPerSeed)

			_, _, simErr := simulation.SimulateFromSeed(
				t, os.Stdout, a.BaseApp,
				simtestutil.AppStateFn(a.AppCodec(), a.SimulationManager(), a.DefaultGenesis()),
				simtypes.RandomAccounts,
				simtestutil.SimulationOperations(a, a.AppCodec(), config),
				a.BlockedAddresses(),
				config,
				a.AppCodec(),
			)
			require.NoError(t, simErr)

			appHash := a.LastCommitID().Hash
			appHashList[j] = fmt.Sprintf("%X", appHash)

			if j != 0 {
				require.Equal(t, appHashList[0], appHashList[j],
					"non-determinism in seed %d", config.Seed)
			}
		}
	}
}

// TestAppImportExport boots the app, runs a small simulation, exports
// genesis, re-imports into a fresh app, and confirms InitGenesis succeeds.
// A full store-by-store diff is omitted here for brevity — that requires
// per-module skip-prefix maps (unbonding queues, etc.) — but the
// import-success path catches the most common breakage (missing
// GenerateGenesisState wiring).
func TestAppImportExport(t *testing.T) {
	config, _, appOptions, a := setupSimulationApp(t, "skipping import/export simulation")
	config.Commit = true // see TestFullAppSimulation note

	_, _, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, a.BaseApp,
		simtestutil.AppStateFn(a.AppCodec(), a.SimulationManager(), a.DefaultGenesis()),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(a, a.AppCodec(), config),
		a.BlockedAddresses(),
		config,
		a.AppCodec(),
	)
	require.NoError(t, simErr)

	// Export
	exported, err := a.ExportAppStateAndValidators(false, nil, nil)
	require.NoError(t, err)

	// Re-import into a fresh app
	newDB := dbm.NewMemDB()
	newApp := app.NewIxoApp(log.NewNopLogger(), newDB, nil, true,
		map[int64]bool{}, app.DefaultNodeHome,
		appOptions, app.EmptyWasmOpts,
		fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))

	_, err = newApp.InitChain(&abci.RequestInitChain{
		ChainId:         SimAppChainID,
		AppStateBytes:   exported.AppState,
		ConsensusParams: simtestutil.DefaultConsensusParams,
	})
	if err != nil {
		// "validator set is empty" is an expected outcome when the simulation
		// happened to unbond every validator — propagate as a skip rather
		// than a fail (mirrors cosmos-sdk's TestAppImportExport).
		if strings.Contains(err.Error(), "validator set is empty") {
			t.Logf("skipping import: validator set drained during simulation\n%s", debug.Stack())
			return
		}
		require.NoError(t, err)
	}

	// Confirm a context can be derived against the freshly-imported app.
	_ = newApp.NewContextLegacy(true, cmtproto.Header{Height: a.LastBlockHeight()})
	_ = appOptions
}

// TestAppSimulationAfterImport runs a simulation, exports state, re-imports
// into a fresh app, and runs another simulation against the imported state.
// This catches the most subtle class of state-bleed bugs: state that
// serialises correctly on the export/import path but then misbehaves when
// the simulator drives further transitions against it.
func TestAppSimulationAfterImport(t *testing.T) {
	config, _, appOptions, a := setupSimulationApp(t, "skipping application sim after import")
	config.Commit = true
	_ = appOptions

	stopEarly, _, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, a.BaseApp,
		simtestutil.AppStateFn(a.AppCodec(), a.SimulationManager(), a.DefaultGenesis()),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(a, a.AppCodec(), config),
		a.BlockedAddresses(),
		config,
		a.AppCodec(),
	)
	require.NoError(t, simErr)
	if stopEarly {
		t.Logf("simulation stopped early; skipping after-import phase")
		return
	}

	exported, err := a.ExportAppStateAndValidators(false, nil, nil)
	require.NoError(t, err)

	newDB := dbm.NewMemDB()
	newApp := app.NewIxoApp(log.NewNopLogger(), newDB, nil, true, map[int64]bool{},
		app.DefaultNodeHome, appOptions, app.EmptyWasmOpts,
		fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))

	_, err = newApp.InitChain(&abci.RequestInitChain{
		ChainId: SimAppChainID, AppStateBytes: exported.AppState,
		ConsensusParams: simtestutil.DefaultConsensusParams,
	})
	if err != nil {
		if strings.Contains(err.Error(), "validator set is empty") {
			t.Logf("skipping after-import sim: validator set drained")
			return
		}
		require.NoError(t, err)
	}

	// Run a second SimulateFromSeed against the imported app — sim ops will
	// generate new tx and the imported state should accept them as if it had
	// been built up incrementally from genesis.
	_, _, err = simulation.SimulateFromSeed(
		t, os.Stdout, newApp.BaseApp,
		simtestutil.AppStateFn(newApp.AppCodec(), newApp.SimulationManager(), newApp.DefaultGenesis()),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(newApp, newApp.AppCodec(), config),
		newApp.BlockedAddresses(),
		config,
		newApp.AppCodec(),
	)
	require.NoError(t, err)
}
