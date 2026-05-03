# State

## Storage layout

The names module owns one KV store with the following key prefixes:

| Prefix | Key                                                                  | Value                       |
| ------ | -------------------------------------------------------------------- | --------------------------- |
| `0x01` | `NamespaceKey(name)` = `0x01 \| name`                                | `ProtocolBuffer(Namespace)` |
| `0x02` | `NameRecordKey(ns, normalized)` = `0x02 \| ns \| 0x00 \| normalized` | `ProtocolBuffer(NameRecord)`|
| `0x03` | `OwnerIndexKey(did, ns, normalized)` = `0x03 \| did \| 0x00 \| ns \| 0x00 \| normalized` | `[]byte{}` (existence-only) |

The `0x00` byte is used as a delimiter between variable-length string components. This is safe because all validated inputs (namespace names, normalised names, DIDs) are ASCII without NUL bytes.

The reverse-lookup index at `0x03` stores empty values — the key encodes the relationship and is sufficient for iteration. `NamesByOwner(did)` does a `KVStorePrefixIterator(store, OwnerIndexPrefix(did))` and re-fetches each `NameRecord` from the primary `0x02` slot.

The module has no parameters store key — there are no module-wide params in v1 (see [Parameters](./05_params.md)).

# Types

### Namespace

A governance-managed bucket of names. Defines validation rules, the registrar set, and self-register / override policies. Created via `MsgCreateNamespace` and replaced wholesale via `MsgUpdateNamespace`.

```go
type Namespace struct {
    Name                   string
    Description            string
    RegistrarAccounts      []string
    AllowSelfRegister      bool
    AllowRegistrarOverride bool
    MinLength              uint32
    MaxLength              uint32
    Regex                  string
    AllowExpiry            bool
}
```

The field's descriptions is as follows:

- `name` - the namespace identifier and uniqueness scope for member names. Lowercase ASCII alphanumeric plus `-` and `_` (`^[a-z0-9_-]+$`). Max `MaxNamespaceNameLength = 64` bytes. Used in storage keys, message routing, and event emission. Immutable after `MsgCreateNamespace`.
- `description` - human-readable summary of the namespace's purpose. Free-form, max `MaxNamespaceDescriptionLength = 4096` bytes.
- `registrar_accounts` - bech32 account addresses authorised to register / update / status-change names in this namespace on behalf of users. List size unbounded; each entry validated as a bech32 address. May be empty when `allow_self_register = true`. At least one of `allow_self_register` or a non-empty registrar list must hold (`MsgCreateNamespace` rejects otherwise with `ErrInvalidNamespace`).
- `allow_self_register` - when `true`, users may register their own names directly via `MsgRegisterName`. When `false`, only registrars may register.
- `allow_registrar_override` - when `true`, registrars may transfer names without owner consent and may status-change names regardless of who owns them. When `false`, transfer is owner-only.
- `min_length` - minimum allowed length of the **normalised** name. Combined with chain-wide charset (`[a-z0-9_-]+`) and the namespace `regex`.
- `max_length` - maximum allowed length of the normalised name. Must be `> 0` and `<= MaxNameLengthCap = 256` (a namespace with `max_length=0` would accept no names; `MsgCreateNamespace` rejects it).
- `regex` - optional Go-`regexp` pattern that the normalised name must additionally match. Empty string disables the check. Max `MaxNamespaceRegexLength = 256` bytes. Validated for compilability at `MsgCreateNamespace` / `MsgUpdateNamespace` time.
- `allow_expiry` - reserved. Gates whether a `NameRecord` in this namespace may have a non-zero `valid_until`. Always `false` in v1; genesis rejects records that violate this.

### NameRecord

A registered name bound to a DID. Keyed in storage by `(namespace, normalized_name)`. Created by `MsgRegisterName` / `MsgRegisterNameByRegistrar`; mutated by `MsgUpdateNameByRegistrar`, `MsgTransferName`, `MsgSetNameStatus`. Never deleted.

```go
type NameRecord struct {
    Namespace      string
    NormalizedName string
    DisplayName    string
    OwnerDid       string
    Verified       bool
    ValidUntil     int64
    Status         NameStatus
    VerifiedBy     string
    EvidenceHash   string
    Source         string
    CreatedAt      int64
    UpdatedAt      int64
}
```

The field's descriptions is as follows:

- `namespace` - the [Namespace](#namespace) name this record belongs to. Immutable.
- `normalized_name` - the canonical lookup form (trimmed + ASCII lowercase). The `(namespace, normalized_name)` tuple is the uniqueness key. Immutable.
- `display_name` - the original case-preserving form supplied by the registrant (e.g. `AliceCAPS`). Useful for UI display.
- `owner_did` - the DID this record resolves to. Reassignable via `MsgTransferName`.
- `verified` - registrar-attested flag. `false` for self-registered records; set by registrars after off-chain proof.
- `valid_until` - Unix-second expiry timestamp. **Reserved**: always `0` (no expiry) in v1. The hook for a future paid-renewal flow; gated per-namespace by `Namespace.allow_expiry`.
- `status` - lifecycle [NameStatus](#namestatus). `ResolveName` returns the record only when this is `NAME_STATUS_ACTIVE`; `GetName` returns it regardless.
- `verified_by` - bech32 address (or DID) of the registrar that produced the attestation. Empty for self-registered records.
- `evidence_hash` - optional content hash of the off-chain attestation (e.g. a JWT or signed VC). Free-form string; the chain does not verify the hash. Max `MaxNameRecordEvidenceHashLength = 256` bytes (enforced both via `ValidateBasic` and inside the keeper handler so Wasm sub-message dispatches that bypass ante still get the check).
- `source` - free-form tag describing the verification source (`self`, `twitter-oauth`, `workos`, `ussd`, `manual`, etc.). The default for self-registration is `"self"`. Max `MaxNameRecordSourceLength = 64` bytes (same dual enforcement as `evidence_hash`).
- `created_at` - Unix-second timestamp at which the record was first written. Set from `ctx.BlockTime().Unix()`.
- `updated_at` - Unix-second timestamp of the most recent write. Refreshed on every state-mutating message.

### NameStatus

```go
const (
    NameStatusUnspecified NameStatus = 0  // never persisted
    NameStatusActive      NameStatus = 1
    NameStatusSuspended   NameStatus = 2
    NameStatusRevoked     NameStatus = 3
    NameStatusTombstoned  NameStatus = 4
)
```

- `ACTIVE` — visible to `ResolveName`.
- `SUSPENDED` — hidden from `ResolveName` (`NotFound`); restorable by registrar.
- `REVOKED` — hidden from `ResolveName`; conventionally indicates an owner-side breach. Restorable by registrar (the chain treats it identically to `SUSPENDED` mechanically — the distinction is for moderation policy / audit).
- `TOMBSTONED` — terminal. Hidden from `ResolveName`. Any further status transition is rejected with `ErrInvalidStatusTransition`. Used for permanent burns where the `(namespace, normalized_name)` tuple must never be reused.

The `UNSPECIFIED` zero value is never a valid persisted state — `MsgSetNameStatus` rejects it via `ValidateBasic`.
