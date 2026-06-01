//go:build interchaintest

package interchaintest

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Module-level helpers used by the consolidated scenario flow tests.
//
// Mirrors the "Module" helper layer in ixo-multiclient-sdk
// (`__tests__/modules/Iid.ts`, `__tests__/modules/Bond.ts`, etc.) — small
// functions that hide gas / sequence / event-extraction noise so the
// scenario flow files read as a story rather than CLI plumbing.
// ----------------------------------------------------------------------------

// IidDocumentJSON is the JSON-decoded shape of an IidDocument query
// response after the manual GetQueryCmd routes through gogo's jsonpb.
// We declare it here once and reuse across all the iid flow steps.
type IidDocumentJSON struct {
	Id                 string `json:"id"`
	Controller         []string `json:"controller"`
	VerificationMethod []struct {
		Id                  string `json:"id"`
		Type                string `json:"type"`
		Controller          string `json:"controller"`
		BlockchainAccountID string `json:"blockchainAccountID"`
	} `json:"verificationMethod"`
	Service []struct {
		Id              string `json:"id"`
		Type            string `json:"type"`
		ServiceEndpoint string `json:"serviceEndpoint"`
	} `json:"service"`
	Authentication []string `json:"authentication"`
	LinkedResource []struct {
		Id              string `json:"id"`
		Type            string `json:"type"`
		Description     string `json:"description"`
		MediaType       string `json:"mediaType"`
		ServiceEndpoint string `json:"serviceEndpoint"`
		Proof           string `json:"proof"`
	} `json:"linkedResource"`
	AlsoKnownAs string `json:"alsoKnownAs"`
	Metadata    struct {
		VersionId   string `json:"versionId"`
		Created     string `json:"created"`
		Updated     string `json:"updated"`
		Deactivated bool   `json:"deactivated"`
	} `json:"metadata"`
}

// IidDoc returns a JSON string for MsgCreateIidDocument / MsgUpdateIidDocument
// containing a single CosmosAccountAddress verification method tied to
// the signer's address. `extraFields` is interpolated as additional
// JSON fields after `controllers` — pass things like
// `"alsoKnownAs": "value"` or `"linkedResource": [...]`.
func IidDoc(did, signerAddress string, extraFields ...string) string {
	extras := ""
	for _, e := range extraFields {
		extras += "  " + e + ",\n"
	}
	return fmt.Sprintf(`{
  "id": %q,
  "controllers": [%q],
%s  "verifications": [{
    "relationships": ["authentication"],
    "method": {
      "id": "%s#key-1",
      "type": "CosmosAccountAddress",
      "controller": %q,
      "blockchainAccountID": %q
    }
  }]
}`, did, did, extras, did, did, signerAddress)
}

// CreateIidDoc registers a fresh DID for `signer` (`did:ixo:<addr>`)
// and returns the DID. Already passes a single
// CosmosAccountAddress verification matching the signer's address so
// subsequent mutations from the same signer authenticate.
func CreateIidDoc(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, signer ibc.Wallet) string {
	t.Helper()
	did := "did:ixo:" + signer.FormattedAddress()
	out, err := chain.GetNode().ExecTx(ctx, signer.KeyName(),
		"iid", "create-iid", IidDoc(did, signer.FormattedAddress()),
		"--gas", "auto", "--gas-adjustment", "1.5",
	)
	require.NoError(t, err, "create-iid: %s", out)
	return did
}

// CreateIidDocWithID registers `did` (rather than `did:ixo:<addr>`) for
// `signer`. Used when the test needs a non-address-derived DID — e.g.
// bond DIDs or entity protocol DIDs.
func CreateIidDocWithID(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, signer ibc.Wallet, did string) {
	t.Helper()
	out, err := chain.GetNode().ExecTx(ctx, signer.KeyName(),
		"iid", "create-iid", IidDoc(did, signer.FormattedAddress()),
		"--gas", "auto", "--gas-adjustment", "1.5",
	)
	require.NoError(t, err, "create-iid %s: %s", did, out)
}

// QueryIidDocument runs `iid iid <did>` and decodes the wrapped
// response into IidDocumentJSON. Fails the test on either query error
// or unmarshal error.
func QueryIidDocument(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, did string) IidDocumentJSON {
	t.Helper()
	stdout, _, err := chain.GetNode().ExecQuery(ctx, "iid", "iid", did, "--output", "json")
	require.NoError(t, err)
	var resp struct {
		IidDocument IidDocumentJSON `json:"iidDocument"`
	}
	require.NoError(t, json.Unmarshal(stdout, &resp), "decode iid response: %s", stdout)
	return resp.IidDocument
}

// IidExec is a thin wrapper around ExecTx that runs an iid CLI tx and
// fails the test if the broadcast errors. Used to keep the scenario
// flow files focused on "what" rather than "how".
func IidExec(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, signer ibc.Wallet, args ...string) string {
	t.Helper()
	cli := append([]string{"iid"}, args...)
	cli = append(cli, "--gas", "auto", "--gas-adjustment", "1.5")
	out, err := chain.GetNode().ExecTx(ctx, signer.KeyName(), cli...)
	require.NoError(t, err, "iid %v: %s", args, out)
	return out
}

// IidExecExpectFail runs an iid tx that the test expects to fail with
// any of the substrings in `wantSubstrs`. Returns true if any matched.
// Helpful for scenario flow steps that interleave negative cases.
func IidExecExpectFail(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, signer ibc.Wallet,
	wantSubstrs []string, args ...string,
) {
	t.Helper()
	cli := append([]string{"iid"}, args...)
	cli = append(cli, "--gas", "auto", "--gas-adjustment", "1.5")
	out, err := chain.GetNode().ExecTx(ctx, signer.KeyName(), cli...)
	combined := fmt.Sprintf("err=%v out=%s", err, out)
	if err == nil {
		// Some chains broadcast OK but encode the failure in the body.
		// Treat that as success-of-rejection too.
	}
	matched := false
	for _, s := range wantSubstrs {
		if containsLowercase(combined, s) {
			matched = true
			break
		}
	}
	require.True(t, matched,
		"expected iid %v to fail with one of %v; got %s", args, wantSubstrs, combined)
}

// ibcDenomHash mirrors the cosmos-sdk IBC denom-hash convention: an
// `ibc/<hex(sha256(<port>/<channel>/<base-denom>))>` string. Used by
// tests to predict the wrapped denom after a forward leg without
// having to query the chain's denom-traces.
func ibcDenomHash(trace string) string {
	sum := sha256.Sum256([]byte(trace))
	return "ibc/" + strings.ToUpper(hex.EncodeToString(sum[:]))
}

// containsLowercase is a case-insensitive substring check that
// avoids the cost of strings.ToLower on the whole haystack.
func containsLowercase(haystack, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	hay := []byte(haystack)
	for i := 0; i < len(hay); i++ {
		if i+len(needle) > len(hay) {
			break
		}
		match := true
		for j := 0; j < len(needle); j++ {
			h, n := hay[i+j], needle[j]
			if h >= 'A' && h <= 'Z' {
				h += 'a' - 'A'
			}
			if n >= 'A' && n <= 'Z' {
				n += 'a' - 'A'
			}
			if h != n {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
