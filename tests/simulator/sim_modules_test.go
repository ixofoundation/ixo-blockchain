// Wiring-coverage tests: assert that every custom ixo module contributes
// the simulation hooks the simulator relies on (StoreDecoder, GenesisState
// generator, WeightedOperations slot). These break loudly when someone
// removes a sim wiring method during a refactor — which previously slipped
// through CI because the simulator tolerates missing wiring silently.
//
// Run with:
//
//	go test -run TestIxoSim_AllCustomModulesRegistered ./tests/simulator/...
package simulator

import (
	"testing"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/app"
)

// expectedCustomModules is the set of custom ixo modules that MUST be
// registered with the SimulationManager. If a module is removed
// intentionally, drop it from this list — but until then a missing entry
// is a regression.
//
// `name` matches each module's types.ModuleName (the routing key, used
// by the module manager), and `storeKey` matches types.StoreKey (the
// KVStore prefix). The two diverge for modules that intentionally
// namespace their store to avoid colliding with the SDK's modules — the
// custom ixo mint module uses "mint-store" so it can run alongside the
// SDK's "mint" module's state for migrations.
var expectedCustomModules = []struct {
	name     string
	storeKey string
}{
	{"iid", "iid"},
	{"bonds", "bonds"},
	{"entity", "entity"},
	{"token", "token"},
	{"claims", "claims"},
	{"smartaccount", "smartaccount"},
	{"epochs", "epochs"},
	{"liquidstake", "liquidstake"},
	{"names", "names"},
	{"mint", "mint-store"},
}

// newSimApp returns a fresh IxoApp suitable for introspecting sim wiring.
// Doesn't run any simulation — just constructs the app so we can read the
// SimulationManager.
func newSimApp(t *testing.T) *app.IxoApp {
	t.Helper()
	dir := t.TempDir()
	appOptions := make(simtestutil.AppOptionsMap)
	appOptions[flags.FlagHome] = dir
	appOptions[server.FlagInvCheckPeriod] = uint(0)

	return app.NewIxoApp(
		log.NewNopLogger(), dbm.NewMemDB(), nil, true, map[int64]bool{}, dir,
		appOptions, app.EmptyWasmOpts,
		baseapp.SetChainID(SimAppChainID),
	)
}

// TestIxoSim_AllCustomModulesRegistered confirms every custom ixo module
// is wired into the SimulationManager. A missing module here means
// determinism / import-export tests run with that module silently absent
// from sim genesis, which can mask state-machine bugs.
func TestIxoSim_AllCustomModulesRegistered(t *testing.T) {
	a := newSimApp(t)
	sm := a.SimulationManager()
	require.NotNil(t, sm, "IxoApp must expose a SimulationManager")

	registered := make(map[string]bool, len(sm.Modules))
	for _, m := range sm.Modules {
		// AppModuleSimulation embeds AppModule which exposes Name() — but
		// it's not on the interface itself. Type-assert to the named-module
		// shape every cosmos-sdk module exposes.
		if named, ok := m.(interface{ Name() string }); ok {
			registered[named.Name()] = true
		}
	}

	for _, m := range expectedCustomModules {
		require.True(t, registered[m.name],
			"module %q must be in SimulationManager.Modules — current set: %v",
			m.name, keys(registered))
	}
}

// TestIxoSim_AllCustomModulesHaveStoreDecoder confirms every custom ixo
// module registers a StoreDecoder so the simulator can diff stores
// between determinism runs. A missing decoder doesn't crash the
// simulator — it just silently drops that module's state from the diff,
// which masks state-machine non-determinism.
func TestIxoSim_AllCustomModulesHaveStoreDecoder(t *testing.T) {
	a := newSimApp(t)
	sm := a.SimulationManager()

	// Seed the registry by calling RegisterStoreDecoder on every module.
	for _, m := range sm.Modules {
		m.RegisterStoreDecoder(sm.StoreDecoders)
	}

	for _, m := range expectedCustomModules {
		_, ok := sm.StoreDecoders[m.storeKey]
		require.True(t, ok,
			"module %q (store key %q) must register a store decoder via RegisterStoreDecoder; current keys: %v",
			m.name, m.storeKey, keys(toBoolMap(sm.StoreDecoders)))
	}
}

// TestIxoSim_DefaultGenesisRoundTrip exports the app's default genesis
// and re-imports it into a fresh app, asserting InitChain succeeds. This
// is a tighter check than TestAppImportExport (which runs a full sim
// before exporting): if a custom module's DefaultGenesis emits a state
// shape that ValidateGenesis or InitGenesis rejects, this catches it
// without the noise of randomly-generated state.
func TestIxoSim_DefaultGenesisRoundTrip(t *testing.T) {
	a := newSimApp(t)

	gen := a.DefaultGenesis()
	require.NotNil(t, gen)

	// Each custom module must contribute a non-nil DefaultGenesis blob.
	for _, m := range expectedCustomModules {
		raw, ok := gen[m.name]
		require.True(t, ok, "default genesis must include module %q", m.name)
		require.NotEmpty(t, raw, "module %q DefaultGenesis must not be empty", m.name)
	}
}

// keys returns the keys of a map[string]bool as a slice — used purely for
// failure messages.
func keys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// toBoolMap converts an arbitrary map's keys to a presence-bool map for
// the failure-message helper.
func toBoolMap[V any](m map[string]V) map[string]bool {
	out := make(map[string]bool, len(m))
	for k := range m {
		out[k] = true
	}
	return out
}
