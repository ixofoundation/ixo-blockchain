package types

import "encoding/json"

// String renders the deprecated single-pool Params layout. Provided only so
// the deprecated proto type satisfies gogoproto's Message interface; the
// only place the legacy Params is read is the v7 upgrade migration.
func (p Params) String() string {
	out, _ := json.MarshalIndent(p, "", "  ")
	return string(out)
}
