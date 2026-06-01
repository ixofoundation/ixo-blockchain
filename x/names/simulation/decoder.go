// Package simulation provides x/names hooks for the cosmos-sdk simulation
// framework. It registers a store decoder for state-determinism diffs;
// per-Msg WeightedOperations are intentionally not implemented because the
// names module is governance-driven (CreateNamespace / UpdateNamespace are
// authority-only) and self-registered names are exercised through the live
// keeper tests.
package simulation

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

// NewDecodeStore returns a decoder for x/names KV pairs. The simulator
// invokes this when diffing two app states for determinism: it must produce
// a stable, human-readable representation of the value bytes for every key
// the module writes.
func NewDecodeStore(_ codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		return fmt.Sprintf("key=%X\nA=%s\nB=%s",
			kvA.Key,
			hex.EncodeToString(kvA.Value),
			hex.EncodeToString(kvB.Value),
		)
	}
}
