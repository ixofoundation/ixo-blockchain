// Invariant smoke tests. RegisterInvariants is the standard cosmos-sdk
// hook for crisis-style invariants — broken ones halt the chain. We
// can't trigger crisis-halt in a unit test, but we CAN pull every
// registered invariant out of the InvariantRouter and run it against a
// freshly-bootstrapped IxoApp to confirm the genesis state is
// invariant-clean.
//
// Run with:
//
//	go test -run TestIxoSim_AllInvariantsCleanAtGenesis ./tests/simulator/...
package simulator

import (
	"strings"
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/app"
)

// TestIxoSim_AllInvariantsCleanAtGenesis runs every registered
// invariant against a fresh IxoApp and asserts none reports a broken
// state. A custom-module invariant that's wrong about its own genesis
// expectations breaks here.
func TestIxoSim_AllInvariantsCleanAtGenesis(t *testing.T) {
	// app.Setup boots an IxoApp with the GenesisStateWithValSet helper —
	// a single bonded validator + funded genesis accounts — so InitChain
	// succeeds and the invariants have something real to check.
	a := app.Setup(false)
	ctx := a.NewContextLegacy(false, cmtproto.Header{})

	invs := a.CrisisKeeper.Invariants()
	require.NotEmpty(t, invs, "IxoApp must register at least one invariant")

	var broken []string
	for _, inv := range invs {
		msg, isBroken := inv(ctx)
		if isBroken {
			broken = append(broken, msg)
		}
	}

	require.Empty(t, broken,
		"these invariants broke against fresh genesis — debug each before merging:\n%s",
		strings.Join(broken, "\n---\n"))
}
