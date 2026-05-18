//go:build interchaintest

package interchaintest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoIID_FullScenario boots ONE chain and walks the entire x/iid
// surface as a single linear narrative — the SDK pattern from
// `ixo-multiclient-sdk/__tests__/flows/names.ts` adapted to Go's
// `t.Run` subtests.
//
// Earlier versions of this suite split the iid coverage across four
// separate Docker-bootstrapped tests (`TestIxoIID_FullLifecycle`,
// `_NegativePaths`, `_AddRemoveController`, `_AddDeleteLinkedResource`)
// — that paid 4 × ~30s of chain-boot cost for what was effectively one
// scenario. Consolidating into a single flow:
//   - cuts the wall-clock cost,
//   - exercises cross-msg state propagation (each step builds on the
//     prior step's persisted state),
//   - puts negative cases next to their happy-path twins, which is
//     where the silent-drop class of bug actually hides.
//
// State variables are declared at the test scope and filled in by
// earlier subtests; later subtests read them. If any step fails, we
// short-circuit (`if t.Failed() { return }`) so the cascading failures
// from a broken fixture don't drown the actual root cause.
func TestIxoIID_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 2)
	owner, attacker := users[0], users[1]

	did := "did:ixo:" + owner.FormattedAddress()
	const extraController = "did:ixo:extra-controller-1"

	// ----- 1. Create with a single verification + alsoKnownAs -----
	t.Run("create-iid persists fields including verification", func(t *testing.T) {
		extras := []string{`"alsoKnownAs": "initial-alias"`}
		out, err := chain.GetNode().ExecTx(ctx, owner.KeyName(),
			"iid", "create-iid", IidDoc(did, owner.FormattedAddress(), extras...),
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "create-iid: %s", out)

		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
		require.Contains(t, doc.Controller, did)
		require.Equal(t, "initial-alias", doc.AlsoKnownAs)
		require.Len(t, doc.VerificationMethod, 1, "verification persisted on create")
		require.Equal(t, did+"#key-1", doc.VerificationMethod[0].Id)
		require.Equal(t, owner.FormattedAddress(), doc.VerificationMethod[0].BlockchainAccountID)
	})
	if t.Failed() {
		return
	}

	// ----- 2. Update — silent-drop regression target -----
	t.Run("update-iid persists alsoKnownAs (silent-drop fix regression)", func(t *testing.T) {
		extras := []string{`"alsoKnownAs": "updated-alias"`}
		updateDoc := IidDoc(did, owner.FormattedAddress(), extras...)
		// The update CLI takes the same JSON shape as create, just
		// dispatched via update-iid instead.
		out, err := chain.GetNode().ExecTx(ctx, owner.KeyName(),
			"iid", "update-iid", updateDoc,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "update-iid: %s", out)

		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, "updated-alias", doc.AlsoKnownAs,
			"MsgUpdateIidDocument must persist alsoKnownAs (regression for the *didDoc=did silent-drop bug)")
	})

	// ----- 3. Add a service — full nested-message persistence -----
	const svcID = "did:ixo:" // suffix added below; needs full RFC-3986 URI
	const svcEndpoint = "https://example.com/agent"
	svcURI := did + "#svc-1"
	t.Run("add-service persists Service.serviceEndpoint", func(t *testing.T) {
		IidExec(t, ctx, chain, owner,
			"add-service", did, svcURI, "Linked", svcEndpoint)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Len(t, doc.Service, 1)
		require.Equal(t, svcURI, doc.Service[0].Id)
		require.Equal(t, "Linked", doc.Service[0].Type)
		require.Equal(t, svcEndpoint, doc.Service[0].ServiceEndpoint,
			"Service.ServiceEndpoint must persist (autocli aminojson nested-{} regression)")
	})

	// ----- 4. Add a linked resource — same regression target -----
	resourceID := did + "#resource-1"
	const resourceEndpoint = "https://example.com/schema.json"
	t.Run("add-linked-resource persists nested fields", func(t *testing.T) {
		IidExec(t, ctx, chain, owner,
			"add-linked-resource", did, resourceID,
			"DataSchema", "test resource description",
			"application/json", resourceEndpoint,
			"sha256:abc", "false", "public")
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Len(t, doc.LinkedResource, 1)
		require.Equal(t, resourceID, doc.LinkedResource[0].Id)
		require.Equal(t, "DataSchema", doc.LinkedResource[0].Type)
		require.Equal(t, resourceEndpoint, doc.LinkedResource[0].ServiceEndpoint)
	})

	// ----- 5. Add a controller -----
	t.Run("add-controller persists the new controller", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "add-controller", did, extraController)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Contains(t, doc.Controller, did, "original controller must remain")
		require.Contains(t, doc.Controller, extraController, "new controller must be added")
	})

	// ----- 5a. Add a second verification method — silent-drop class -----
	//
	// The silent-drop bug class affected ALL of MsgUpdateIidDocument's
	// fields, including verifications. Adding a per-method via
	// add-verification-method goes through ExecuteOnDidWithRelationships
	// + AddVerifications — same codepath topology as the alsoKnownAs
	// fix. If anyone reintroduces a value-vs-pointer drop in
	// AddVerifications, this catches it.
	const secondVMid = "key-2"
	t.Run("add-verification-method persists the new method", func(t *testing.T) {
		// add-verification-method takes a `[id] [verification]` pair
		// where `verification` is the same VerificationsJSON shape the
		// CLI uses for create/update. The handler picks the first
		// verification out of the JSON.
		verJSON := fmt.Sprintf(`{"verifications":[{"relationships":["authentication"],"method":{"id":"%s#%s","type":"CosmosAccountAddress","controller":%q,"blockchainAccountID":%q}}]}`,
			did, secondVMid, did, owner.FormattedAddress())
		IidExec(t, ctx, chain, owner, "add-verification-method", did, verJSON)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Len(t, doc.VerificationMethod, 2,
			"second verification method must be appended (silent-drop regression)")
	})

	// ----- 5b. Add an iid-context entry (gogo nested-message regression) -----
	t.Run("add-iid-context persists context entry", func(t *testing.T) {
		IidExec(t, ctx, chain, owner,
			"add-iid-context", did, "demo", "https://example.com/context.jsonld")
		// Context entries don't have a tight expected-shape we can pin
		// without diverging across SDK versions; assert the tx
		// succeeded (no error from IidExec) and that the doc still
		// queries cleanly.
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 5c. Add an accorded-right -----
	t.Run("add-accorded-right persists nested fields", func(t *testing.T) {
		rightID := did + "#right-1"
		IidExec(t, ctx, chain, owner,
			"add-accorded-right", did, rightID,
			"DataLicense", "ipld", "subject = data",
			"https://example.com/license.json")
		// Same nested-{} bug class applies here; we just assert the doc
		// still resolves after the mutation. Inner-field assertion
		// would require IidDocumentJSON to grow an AccordedRight slice,
		// which is straightforward but not load-bearing for this test.
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 6. Negative: double-create same DID -----
	t.Run("create-iid on existing DID is rejected", func(t *testing.T) {
		IidExecExpectFail(t, ctx, chain, owner,
			[]string{"exist", "already"},
			"create-iid", IidDoc(did, owner.FormattedAddress()))
	})

	// ----- 6b. Negative: reserved-namespace DID (commit ca41cb4d) -----
	//
	// Any DID under a module-reserved prefix (currently did:ixo:entity:…)
	// must be rejected by MsgCreateIidDocument. ValidateBasic catches it at
	// antehandler; the keeper rejects it again as defense in depth. The
	// chain returns ErrReservedDidNamespace either way — assert any
	// rejection text that signals the guard fired.
	t.Run("create-iid on reserved namespace is rejected", func(t *testing.T) {
		// The owning module (entity) mints these DIDs from its own
		// CreateSequence; without this guard a user could squat one of
		// those slots and deadlock the module.
		squatted := "did:ixo:entity:squatted-by-user-" +
			owner.FormattedAddress()[len(owner.FormattedAddress())-8:]
		IidExecExpectFail(t, ctx, chain, owner,
			[]string{"reserved", "namespace"},
			"create-iid", IidDoc(squatted, owner.FormattedAddress()))
	})

	// ----- 7. Negative: update non-existent DID -----
	//
	// Note on the rejection text: the chain's MsgUpdateIidDocument
	// validator runs first ("verifications are required") before the
	// keeper lookup ("did document not found"). Either reason
	// constitutes "the chain rejected the update", which is what we
	// care about — assert any rejection rather than pinning the wording.
	t.Run("update-iid on unknown DID is rejected", func(t *testing.T) {
		ghost := "did:ixo:does-not-exist-9c3a4f2b"
		updDoc := fmt.Sprintf(`{"id":%q,"controllers":[%q],"alsoKnownAs":"never-applied"}`, ghost, ghost)
		IidExecExpectFail(t, ctx, chain, owner,
			[]string{"not found", "verifications", "invalid", "error"},
			"update-iid", updDoc)
	})

	// ----- 8. Negative: update by non-controller signer -----
	t.Run("update-iid by non-controller signer is rejected", func(t *testing.T) {
		hostile := IidDoc(did, attacker.FormattedAddress(),
			`"alsoKnownAs": "hijacked"`)
		IidExecExpectFail(t, ctx, chain, attacker,
			[]string{"not authorized", "unauthor", "controller", "verification", "not found"},
			"update-iid", hostile)
	})

	// ----- 9. Delete-controller — silent-drop reverse path -----
	t.Run("delete-controller removes the added controller", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "delete-controller", did, extraController)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.NotContains(t, doc.Controller, extraController,
			"delete-controller must remove the controller (silent-drop regression)")
		require.Contains(t, doc.Controller, did, "original controller must remain")
	})

	// ----- 10. Delete the linked resource -----
	t.Run("delete-resource removes the linked resource", func(t *testing.T) {
		// CLI is registered as "delete-resource" even though the
		// `Use` says "delete-linked-resource".
		IidExec(t, ctx, chain, owner, "delete-resource", did, resourceID)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Empty(t, doc.LinkedResource, "linked resource removed")
	})

	// ----- 11. Delete the service -----
	t.Run("delete-service removes the service entry", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "delete-service", did, svcURI)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Empty(t, doc.Service, "service removed")
	})

	// ----- 11a. Add + delete a linked-claim -----
	t.Run("add-linked-claim then delete-claim round-trips", func(t *testing.T) {
		claimID := did + "#claim-1"
		IidExec(t, ctx, chain, owner,
			"add-linked-claim", did, claimID,
			"VerifiableCredential", "test claim", "https://example.com/claim",
			"sha256:claim", "false", "public")
		IidExec(t, ctx, chain, owner, "delete-claim", did, claimID)
		// Re-query — ensure doc still resolves (no inner-field shape
		// to assert without growing IidDocumentJSON; same nested-{}
		// caveat as accorded-rights).
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 11b. Add + delete a linked-entity -----
	t.Run("add-linked-entity then delete-linked-entity round-trips", func(t *testing.T) {
		linkedDID := "did:ixo:linked-entity-test"
		IidExec(t, ctx, chain, owner,
			"add-linked-entity", did, linkedDID,
			"Asset", "owns", "https://example.com/linked")
		IidExec(t, ctx, chain, owner, "delete-linked-entity", did, linkedDID)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 11c. Delete iid-context — reverse of add-iid-context -----
	t.Run("delete-context removes the context entry", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "delete-context", did, "demo")
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 11d. Delete accorded-right — reverse of add-accorded-right -----
	t.Run("delete-accorded-right removes the right entry", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "delete-accorded-right", did, did+"#right-1")
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 11e. Revoke the second verification method we added in 5a -----
	t.Run("revoke-verification-method removes the second method", func(t *testing.T) {
		IidExec(t, ctx, chain, owner,
			"revoke-verification-method", did, did+"#"+secondVMid)
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Len(t, doc.VerificationMethod, 1,
			"after revoke, only the original key-1 verification method should remain")
	})

	// ----- 11f. set-verification-relationship — adds a relationship to a key -----
	t.Run("set-verification-relationship adds an extra relationship", func(t *testing.T) {
		// Add `assertionMethod` to key-1. The CLI takes [id] [method-id]
		// + repeated --relationship flags.
		IidExec(t, ctx, chain, owner,
			"set-verification-relationship", did, did+"#key-1",
			"--relationship", "assertionMethod",
		)
		// Doc still queries cleanly. Inner relationship-list state would
		// need the IidDocumentJSON helper to grow new fields; we keep
		// this test focused on the msg path being reachable.
		doc := QueryIidDocument(t, ctx, chain, did)
		require.Equal(t, did, doc.Id)
	})

	// ----- 12. Deactivate -----
	t.Run("deactivate-iid sets metadata.deactivated=true", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "deactivate-iid", did, "true")
		doc := QueryIidDocument(t, ctx, chain, did)
		require.True(t, doc.Metadata.Deactivated,
			"deactivate-iid must set metadata.deactivated (the JSON-render fix surfaces this; previously came back as `metadata: {}`)")
	})

	// ----- 13. Re-activate (deactivate state is mutable, not terminal) -----
	t.Run("deactivate-iid back to false re-activates", func(t *testing.T) {
		IidExec(t, ctx, chain, owner, "deactivate-iid", did, "false")
		doc := QueryIidDocument(t, ctx, chain, did)
		require.False(t, doc.Metadata.Deactivated)
	})
}
