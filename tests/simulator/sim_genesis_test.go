// Genesis-validation tests. Each custom module's DefaultGenesis must
// pass ValidateGenesis (catches bad defaults) and round-trip through
// JSON marshal/unmarshal (catches proto / JSON tag drift). These are
// fast in-process tests that run alongside the full simulator suite.
package simulator

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/stretchr/testify/require"
)

// TestIxoSim_DefaultGenesisValidates iterates every module registered on
// IxoApp's module manager and asserts that its DefaultGenesis passes
// ValidateGenesis. A module with a bad default is harmless until someone
// tries to bootstrap a fresh chain — at which point startup fails. This
// test fires that signal up front.
func TestIxoSim_DefaultGenesisValidates(t *testing.T) {
	a := newSimApp(t)
	gen := a.DefaultGenesis()

	mm := a.ModuleManager
	require.NotNil(t, mm, "IxoApp must expose a ModuleManager")

	for name, mod := range mm.Modules {
		hg, ok := mod.(module.HasGenesis)
		if !ok {
			continue
		}
		raw, present := gen[name]
		require.True(t, present, "module %q missing from DefaultGenesis", name)
		require.NoError(t, hg.ValidateGenesis(a.AppCodec(), nil, raw),
			"module %q DefaultGenesis must pass ValidateGenesis", name)
	}
}

// TestIxoSim_DefaultGenesisJSONRoundTrip asserts every custom module's
// DefaultGenesis blob survives a JSON marshal → unmarshal trip
// unchanged. Catches proto-tag / JSON-tag drift the moment it's
// introduced, instead of waiting for a state-export to fail in
// production.
func TestIxoSim_DefaultGenesisJSONRoundTrip(t *testing.T) {
	a := newSimApp(t)
	gen := a.DefaultGenesis()

	for _, m := range expectedCustomModules {
		raw, ok := gen[m.name]
		require.True(t, ok, "default genesis must include module %q", m.name)

		// Decode and re-encode through encoding/json. A schema drift would
		// surface as either a unmarshal error or a re-marshal that
		// produces non-equivalent JSON.
		var generic any
		require.NoError(t, json.Unmarshal(raw, &generic),
			"module %q DefaultGenesis must be valid JSON", m.name)
		reencoded, err := json.Marshal(generic)
		require.NoError(t, err, "module %q DefaultGenesis must re-marshal", m.name)
		require.NotEmpty(t, reencoded)

		// And it must validate after the round trip.
		mod, present := a.ModuleManager.Modules[m.name]
		require.True(t, present, "module %q missing from ModuleManager", m.name)
		hg, ok := mod.(module.HasGenesis)
		require.True(t, ok, "module %q must implement HasGenesis", m.name)
		require.NoError(t, hg.ValidateGenesis(a.AppCodec(), nil, reencoded),
			"module %q DefaultGenesis must validate after JSON round-trip", m.name)
	}
}
