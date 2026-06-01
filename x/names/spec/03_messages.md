# Messages

In this section we describe the processing of the names messages and the corresponding updates to state. All created/modified state objects are defined within the [state](./02_state.md) section.

The messages fall into three groups by who is allowed to sign them:

- **Governance** — only the chain governance module address. Reached via `cosmos.gov.v1.MsgSubmitProposal` carrying the inner Msg.
- **User** — any account; auth is on the relationship between the signer and the `owner_did` referenced by the message.
- **Registrar** — accounts listed in the target namespace's `registrar_accounts`.

## Governance operations

### MsgCreateNamespace

Registers a new [Namespace](./02_state.md#namespace).

```go
type MsgCreateNamespace struct {
    Authority string
    Namespace *Namespace
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be the chain governance module address; otherwise rejected with `ErrInvalidAuthority`.
- `namespace` - the full namespace configuration to register. Required (rejected with `ErrInvalidRequest` if nil) and must pass internal validation (`ValidateNamespace` — see [State](./02_state.md#namespace)). Each address in `registrar_accounts` must be a valid bech32 address. Either `allow_self_register = true` or a non-empty `registrar_accounts` is required.

The namespace name is `(namespace.name)`'s uniqueness key. If a namespace by that name already exists the message is rejected with `ErrNamespaceExists`. On success, emits a `NamespaceCreatedEvent` carrying the full namespace and the gov authority signer.

### MsgUpdateNamespace

Replaces an existing namespace's configuration in full. The namespace's `name` selects the target — `name` is immutable, but every other field can be replaced (description, registrar set, allow flags, length bounds, regex, allow_expiry).

```go
type MsgUpdateNamespace struct {
    Authority string
    Namespace *Namespace
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be the chain governance module address.
- `namespace` - the new full configuration. Same validation as `MsgCreateNamespace`. The `namespace.name` field selects which existing namespace to update (`ErrNamespaceNotFound` if no such namespace).

A subsequent `MsgUpdateNamespace` does **not** affect existing `NameRecord`s — names that were validly registered under the previous configuration remain valid even if the new configuration would have rejected them. The new configuration applies only to *future* registration / update messages. Emits a `NamespaceUpdatedEvent`.

## User operations

### MsgRegisterName

Self-registers a name in a namespace whose `allow_self_register` is `true`.

```go
type MsgRegisterName struct {
    Signer    string
    Namespace string
    Name      string
    OwnerDid  string
}
```

The field's descriptions is as follows:

- `signer` - bech32 address signing the tx. Must control `owner_did` — either via an `authentication` verification method on the IID document, or by being listed as one of its `controller`s. Otherwise rejected with `ErrUnauthorized`.
- `namespace` - target namespace. Must exist (`ErrNamespaceNotFound` otherwise) and must have `allow_self_register = true` (`ErrSelfRegisterNotAllowed` otherwise).
- `name` - the display form (case preserved). Server-side normalised via `NormalizeName` (trim + ASCII lowercase). The normalised form must satisfy the chain-wide charset `[a-z0-9_-]+`, the namespace's `min_length`/`max_length` bounds, and the namespace's optional `regex`. Rejected with `ErrInvalidName` otherwise.
- `owner_did` - the DID that will own the record. Must reference an existing IID document (validated implicitly by the `verifyDidController` check).

If `(namespace, normalized_name)` is already taken (in any status, including tombstoned) the message is rejected with `ErrNameTaken`.

On success the keeper:

1. Writes a `NameRecord` with `verified = false`, `valid_until = 0`, `status = ACTIVE`, `source = "self"`, `created_at = updated_at = ctx.BlockTime().Unix()`.
2. Writes the reverse-lookup index entry under `OwnerIndexKey(owner_did, namespace, normalized_name)`.
3. Emits `NameRegisteredEvent` carrying the full record and the signer address as `registered_by`.
4. Returns `MsgRegisterNameResponse{ NormalizedName: <normalized> }` so the caller can confirm the canonical form the chain stored.

### MsgTransferName

Reassigns a `NameRecord`'s `owner_did`. Permitted for the current owner; permitted for a registrar when the namespace has `allow_registrar_override = true`.

```go
type MsgTransferName struct {
    Signer         string
    Namespace      string
    NormalizedName string
    NewOwnerDid    string
}
```

The field's descriptions is as follows:

- `signer` - bech32 address. Must satisfy at least one of:
  - controls the record's current `owner_did` (via `verifyDidController`), OR
  - is in `Namespace.registrar_accounts` AND the namespace has `allow_registrar_override = true`.
  Otherwise rejected with `ErrUnauthorized`.
- `namespace` - the namespace containing the record. Must exist.
- `normalized_name` - canonical name to transfer. Must reference an existing record (`ErrNameNotFound` otherwise).
- `new_owner_did` - the DID that becomes the new owner. Must reference an existing IID document (`ErrInvalidDID` otherwise) and must differ from the current owner (`ErrInvalidRequest` if equal).

On success the keeper:

1. Removes the old reverse-index entry under `OwnerIndexKey(old_owner, namespace, normalized_name)`.
2. Updates the record's `owner_did` to the new value and refreshes `updated_at`.
3. Persists the record (which also writes the new reverse-index entry).
4. Emits **two** events (dual-emit pattern — see [Events](./04_events.md#dual-emission-pattern)):
   - `NameUpdatedEvent` carrying the post-transfer full `NameRecord` and the signer as `updated_by`.
   - `NameTransferredEvent` carrying namespace, normalised name, from-DID, to-DID, and signer as `transferred_by`.

## Registrar operations

### MsgRegisterNameByRegistrar

Registers a name on behalf of a target DID. Used for custodial / USSD / no-gas users and for registrar-attested names (Twitter, KYC, etc.). Works in any namespace (self-register or registrar-only).

```go
type MsgRegisterNameByRegistrar struct {
    Registrar    string
    Namespace    string
    Name         string
    OwnerDid     string
    Verified     bool
    EvidenceHash string
    Source       string
}
```

The field's descriptions is as follows:

- `registrar` - bech32 address signing the tx. Must be in `Namespace.registrar_accounts`; otherwise rejected with `ErrNotRegistrar`.
- `namespace` - target namespace. Must exist.
- `name` - display form; server-side normalised. Same charset / length / regex validation as `MsgRegisterName`.
- `owner_did` - the DID that will own the record. Unlike `MsgRegisterName`, the registrar is **not** required to control this DID — the registrar acts on behalf of the user. The DID must reference an existing IID document (`ErrInvalidDID` otherwise).
- `verified` - whether the registrar is attesting to off-chain proof. Persisted as `NameRecord.verified`.
- `evidence_hash` - optional hash of the off-chain proof (e.g. JWT or signed VC). Free-form string, stored verbatim — the chain does not verify the content. Max `MaxNameRecordEvidenceHashLength = 256` bytes; enforced both via `ValidateBasic` and inside the keeper handler.
- `source` - free-form attestation source tag (e.g. `twitter-oauth`, `workos`). Persisted as `NameRecord.source`. Max `MaxNameRecordSourceLength = 64` bytes; same dual enforcement as `evidence_hash`.

If `(namespace, normalized_name)` is already taken the message is rejected with `ErrNameTaken`.

On success:

1. Writes the `NameRecord` with `verified` from the message, `valid_until = 0`, `status = ACTIVE`, `verified_by = registrar address`, `evidence_hash` and `source` from the message, `created_at = updated_at = ctx.BlockTime().Unix()`.
2. Writes the reverse-index entry.
3. Emits `NameRegisteredEvent` with `registered_by = registrar address`.
4. Returns `MsgRegisterNameByRegistrarResponse{ NormalizedName: <normalized> }`.

### MsgUpdateNameByRegistrar

Updates the verification metadata of an existing record. **Owner DID is not changed by this message** — use `MsgTransferName` for that. **Status is not changed by this message** — use `MsgSetNameStatus`. **Display name and namespace are immutable.**

```go
type MsgUpdateNameByRegistrar struct {
    Registrar      string
    Namespace      string
    NormalizedName string
    Verified       bool
    EvidenceHash   string
    Source         string
}
```

The field's descriptions is as follows:

- `registrar` - bech32 signer. Must be in `Namespace.registrar_accounts` (`ErrNotRegistrar` otherwise).
- `namespace` - target namespace.
- `normalized_name` - canonical name to update (`ErrNameNotFound` if no such record).
- `verified` - new value of `NameRecord.verified`. Set to `false` to retract a previous attestation.
- `evidence_hash` - replaces the existing evidence hash. Max `MaxNameRecordEvidenceHashLength = 256` bytes (dual-enforced).
- `source` - replaces the existing source tag. Max `MaxNameRecordSourceLength = 64` bytes (dual-enforced).

On success:

1. Mutates the record's `verified`, `verified_by` (set to this registrar), `evidence_hash`, `source`, and `updated_at`.
2. Persists the record.
3. Emits `NameUpdatedEvent` with the post-update record and `updated_by = registrar address`.

### MsgSetNameStatus

Changes a record's lifecycle status. Permitted for namespace registrars and for the chain governance module address (no `allow_registrar_override` requirement — namespace registrars always have moderation authority over their namespace's records).

```go
type MsgSetNameStatus struct {
    Signer         string
    Namespace      string
    NormalizedName string
    Status         NameStatus
    Reason         string
}
```

The field's descriptions is as follows:

- `signer` - bech32 address. Must be in `Namespace.registrar_accounts` OR equal the chain governance module address. Otherwise rejected with `ErrUnauthorized`.
- `namespace` - target namespace. Must exist.
- `normalized_name` - canonical name (`ErrNameNotFound` if no record).
- `status` - the [NameStatus](./02_state.md#namestatus) to transition to. Must be `ACTIVE`, `SUSPENDED`, `REVOKED`, or `TOMBSTONED`. `UNSPECIFIED` (the proto zero value) is rejected with `ErrInvalidStatusTransition`. The check runs both in `ValidateBasic` (caught for normal txs by the antehandler) and inside the keeper handler (caught for Wasm sub-message dispatches that bypass ante).
- `reason` - free-form string surfaced in the audit event. Empty allowed but encouraged for moderation review.

Pre-conditions on the existing record:

- The record's current status must not be `TOMBSTONED` (terminal — `ErrInvalidStatusTransition`).
- The new status must differ from the current status (`ErrInvalidStatusTransition`).

On success:

1. Updates the record's `status` and `updated_at`.
2. Emits **two** events (dual-emit — see [Events](./04_events.md#dual-emission-pattern)):
   - `NameUpdatedEvent` with the post-change full `NameRecord` and signer as `updated_by`.
   - `NameStatusChangedEvent` with namespace, normalised name, old status, new status, signer as `changed_by`, and the reason string.
