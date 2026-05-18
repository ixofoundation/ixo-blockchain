// Module-ordering tests. The OrderBeginBlockers / OrderEndBlockers /
// OrderInitGenesis lists in app/modules.go control hook execution order
// at each block. A custom module silently dropped from one of these
// lists keeps compiling and running — but its hooks never fire,
// silently breaking inflation, slashing, or epoch advancement.
package simulator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v6/app"
)

// modulesNeedingBeginBlocker lists custom ixo modules that publish a
// BeginBlocker hook. If a module here is missing from
// OrderBeginBlockers, its hook never fires.
var modulesNeedingBeginBlocker = []string{
	"epochs", // epochs MUST run first; mint's hook depends on it
	"mint",   // custom inflation triggered by epoch hook
}

// modulesNeedingEndBlocker lists custom ixo modules that publish an
// EndBlocker. liquidstake autocompounds and rebalances at end of block;
// claims processes timeouts.
var modulesNeedingEndBlocker = []string{
	"liquidstake",
	"claims",
	"bonds", // batch settlement runs in EndBlocker
}

// modulesNeedingInitGenesis lists every custom module that has genesis
// state. Order matters when modules cross-reference each other (entity
// reads iid; claims reads entity).
var modulesNeedingInitGenesis = []string{
	"iid",
	"bonds",
	"entity",
	"token",
	"claims",
	"smartaccount",
	"epochs",
	"liquidstake",
	"names",
	"mint",
}

// TestIxoSim_BeginBlockersIncludeAllExpected asserts every module in
// modulesNeedingBeginBlocker appears in OrderBeginBlockers.
func TestIxoSim_BeginBlockersIncludeAllExpected(t *testing.T) {
	order := app.OrderBeginBlockers()
	present := make(map[string]bool, len(order))
	for _, name := range order {
		present[name] = true
	}

	for _, name := range modulesNeedingBeginBlocker {
		require.True(t, present[name],
			"module %q must appear in OrderBeginBlockers; current: %v", name, order)
	}

	// epochs MUST come before mint in the begin-blocker order.
	epochsIdx, mintIdx := -1, -1
	for i, name := range order {
		switch name {
		case "epochs":
			epochsIdx = i
		case "mint":
			mintIdx = i
		}
	}
	require.GreaterOrEqual(t, epochsIdx, 0, "epochs must be in OrderBeginBlockers")
	require.GreaterOrEqual(t, mintIdx, 0, "mint must be in OrderBeginBlockers")
	require.Less(t, epochsIdx, mintIdx,
		"epochs BeginBlocker must run before mint (mint's inflation depends on the epoch tick)")
}

// TestIxoSim_EndBlockersIncludeAllExpected asserts every module in
// modulesNeedingEndBlocker appears in OrderEndBlockers.
func TestIxoSim_EndBlockersIncludeAllExpected(t *testing.T) {
	order := app.OrderEndBlockers()
	present := make(map[string]bool, len(order))
	for _, name := range order {
		present[name] = true
	}

	for _, name := range modulesNeedingEndBlocker {
		require.True(t, present[name],
			"module %q must appear in OrderEndBlockers; current: %v", name, order)
	}
}

// TestIxoSim_InitGenesisOrderIncludesAllCustomModules confirms every
// custom ixo module is present in OrderInitGenesis. A module dropped
// here means its InitGenesis never runs and any default state silently
// stays unwritten.
func TestIxoSim_InitGenesisOrderIncludesAllCustomModules(t *testing.T) {
	order := app.OrderInitGenesis()
	present := make(map[string]bool, len(order))
	for _, name := range order {
		present[name] = true
	}

	for _, name := range modulesNeedingInitGenesis {
		require.True(t, present[name],
			"custom module %q must appear in OrderInitGenesis; current: %v", name, order)
	}
}
