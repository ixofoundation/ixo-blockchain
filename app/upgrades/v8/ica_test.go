package v8

import (
	"strings"
	"testing"
)

// TestICAHostAllowMessages_ExcludesUnsafe guards the curated ICA host allow-list
// against regressions that would re-open the ante-bypass class: no allow-all, no
// ixo custom messages, no authz.MsgExec (nested bypass), no wasm execute (entity
// NFT decorator bypass).
func TestICAHostAllowMessages_ExcludesUnsafe(t *testing.T) {
	if len(icaHostAllowMessages) == 0 {
		t.Fatal("ICA host allow-list must not be empty (would disable all ICA)")
	}
	for _, m := range icaHostAllowMessages {
		switch {
		case m == "*":
			t.Errorf("ICA allow-list must not be allow-all")
		case strings.HasPrefix(m, "/ixo."):
			t.Errorf("ixo custom message %q must not be in the ICA allow-list until proven ante-safe", m)
		case m == "/cosmos.authz.v1beta1.MsgExec":
			t.Errorf("authz.MsgExec must not be in the ICA allow-list (nested ante bypass)")
		case strings.HasPrefix(m, "/cosmwasm."):
			t.Errorf("wasm message %q must not be in the ICA allow-list (entity NFT decorator bypass)", m)
		}
	}
}

// TestICAHostAllowMessages_IncludesStandardCosmos confirms the common, ante-safe
// Cosmos SDK operations remain allowed for legitimate ICA usage.
func TestICAHostAllowMessages_IncludesStandardCosmos(t *testing.T) {
	want := []string{
		"/cosmos.bank.v1beta1.MsgSend",
		"/cosmos.staking.v1beta1.MsgDelegate",
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
		"/cosmos.gov.v1.MsgVote",
		"/ibc.applications.transfer.v1.MsgTransfer",
	}
	set := make(map[string]bool, len(icaHostAllowMessages))
	for _, m := range icaHostAllowMessages {
		set[m] = true
	}
	for _, w := range want {
		if !set[w] {
			t.Errorf("expected standard cosmos message %q in the ICA allow-list", w)
		}
	}
}

func TestICAHostAllowMessages_NoDuplicates(t *testing.T) {
	seen := make(map[string]bool, len(icaHostAllowMessages))
	for _, m := range icaHostAllowMessages {
		if seen[m] {
			t.Errorf("duplicate entry %q in ICA allow-list", m)
		}
		seen[m] = true
	}
}
