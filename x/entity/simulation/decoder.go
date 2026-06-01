// Package simulation provides x/entity hooks for the cosmos-sdk simulation
// framework.
package simulation

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

// NewDecodeStore returns a decoder for x/entity KV pairs.
func NewDecodeStore(_ codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		return fmt.Sprintf("key=%X\nA=%s\nB=%s",
			kvA.Key,
			hex.EncodeToString(kvA.Value),
			hex.EncodeToString(kvB.Value),
		)
	}
}
