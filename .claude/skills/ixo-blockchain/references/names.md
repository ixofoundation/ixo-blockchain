# Names module — `x/names`

**Proto package:** `ixo.names.v1beta1` · **TS typeUrl prefix:** `/ixo.names.v1beta1.` · **CLI:** no tx commands registered (construct via SDK); queries ARE registered via autocli (`ixod query names ...`)

## Purpose
The names module is an on-chain name service that binds human-readable handles to DIDs. A `name` (e.g. `alice`) resolves to an `owner_did` and always lives inside a governance-managed `namespace` (e.g. `yoid`, `twitter`) that scopes uniqueness and defines validation rules, registrar accounts, and self-register/override policy. Names are never hard-deleted; a `status` lifecycle (active/suspended/revoked/tombstoned) governs visibility. It depends one-way on the `iid` module to verify DID controllers and confirm target DIDs exist.

## Concepts & state
- **Namespace** (`0x01`, key `name`): a governance-created bucket. Holds validation rules (`min_length`, `max_length`, `regex`), the `registrar_accounts` list, and the `allow_self_register` / `allow_registrar_override` / `allow_expiry` flags. Created/updated only by the gov authority. Name is immutable; `MsgUpdateNamespace` replaces every other field wholesale and does not retro-invalidate existing records.
- **NameRecord** (`0x02`, key `(namespace, normalized_name)`): a registered name bound to `owner_did`. Carries `display_name` (case-preserved), `verified` + `verified_by` + `evidence_hash` + `source` (off-chain attestation metadata), `status`, reserved `valid_until` (always `0` in v1), and `created_at`/`updated_at`. Uniqueness is on `normalized_name` (trim + ASCII lowercase).
- **Owner index** (`0x03`, key `(owner_did, namespace, normalized_name)`, empty value): reverse lookup powering `NamesByOwner`. Maintained on register and transfer; not touched by status changes.
- **NameStatus enum:** `NAME_STATUS_UNSPECIFIED=0` (never persisted), `NAME_STATUS_ACTIVE=1`, `NAME_STATUS_SUSPENDED=2`, `NAME_STATUS_REVOKED=3`, `NAME_STATUS_TOMBSTONED=4` (terminal). Only `ACTIVE` resolves via `ResolveName`.
- No module params store (no module-wide params in v1).

## Messages
Source order from `tx.proto`. Nested `Namespace` / `NameStatus` types resolved from `names.proto`.

### MsgCreateNamespace
- **Purpose:** Register a new namespace.
- **Signer / auth:** `authority` signs; must equal the gov module address (`ErrInvalidAuthority` otherwise). Authority-only (governance).
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| authority | string | yes | Bech32 gov module address allowed to create namespaces. |
| namespace | Namespace | yes | Full configuration of the new namespace (nested message, see below). |

`Namespace` fields (from `names.proto`):

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| name | string | yes | Namespace identifier and uniqueness scope. Lowercase ASCII recommended (`^[a-z0-9_-]+$`); immutable after creation. |
| description | string | no | Human-readable summary. |
| registrar_accounts | repeated string | no | Bech32 accounts authorised to register/moderate names. Each validated as bech32. |
| allow_self_register | bool | no | When true, users may register their own names via `MsgRegisterName`. |
| allow_registrar_override | bool | no | When true, registrars may transfer/update names regardless of owner. |
| min_length | uint32 | no | Minimum length of the normalized name. |
| max_length | uint32 | yes | Maximum length of the normalized name. Must be `> 0`. |
| regex | string | no | Additional Go-regexp the normalized name must match. Empty = no extra check. |
| allow_expiry | bool | no | Reserved; when true a record may carry non-zero `valid_until`. Always false in v1. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgCreateNamespace',
  value: ixo.names.v1beta1.MsgCreateNamespace.fromPartial({
    authority: govModuleAddress,
    namespace: ixo.names.v1beta1.Namespace.fromPartial({
      name: 'yoid',
      maxLength: 32,
      allowSelfRegister: true,
    }),
  }),
};
```
- **Gotchas:** Rejected if `namespace` is nil (`ErrInvalidRequest`) or the name already exists (`ErrNamespaceExists`). `ValidateNamespace` enforces the name charset/length, requires `max_length > 0` (and `<= MaxNameLengthCap = 256`), validates each `registrar_accounts` entry as bech32, requires at least one of `allow_self_register=true` OR a non-empty `registrar_accounts`, and compiles `regex` for validity. Normally reached via `cosmos.gov.v1.MsgSubmitProposal`.

### MsgUpdateNamespace
- **Purpose:** Replace an existing namespace's configuration in full.
- **Signer / auth:** `authority` signs; must equal the gov module address. Authority-only (governance).
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| authority | string | yes | Bech32 gov module address. |
| namespace | Namespace | yes | New full configuration; `namespace.name` selects the existing target (same Namespace fields as above). |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgUpdateNamespace',
  value: ixo.names.v1beta1.MsgUpdateNamespace.fromPartial({
    authority: govModuleAddress,
    namespace: ixo.names.v1beta1.Namespace.fromPartial({
      name: 'yoid',
      maxLength: 32,
      allowSelfRegister: false,
    }),
  }),
};
```
- **Gotchas:** Rejected if `namespace` is nil (`ErrInvalidRequest`) or the target name does not exist (`ErrNamespaceNotFound`). Same `ValidateNamespace` rules as create. `name` is the selector and is effectively immutable. Does NOT affect existing `NameRecord`s — only future register/update messages.

### MsgRegisterName
- **Purpose:** Self-register a name in a namespace that allows self-registration.
- **Signer / auth:** `signer` signs; must control `owner_did` (via an `authentication` verification method on the IID document, or by being a controller). User-facing.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| signer | string | yes | Bech32 address signing the tx; must control `owner_did`. |
| namespace | string | yes | Namespace to register under. |
| name | string | yes | Display name (case preserved); normalized server-side for uniqueness. |
| owner_did | string | yes | DID the name will resolve to. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgRegisterName',
  value: ixo.names.v1beta1.MsgRegisterName.fromPartial({
    signer: signerAddress,
    namespace: 'yoid',
    name: 'alice',
    ownerDid: 'did:ixo:abc123',
  }),
};
```
- **Gotchas:** Namespace must exist (`ErrNamespaceNotFound`) and have `allow_self_register=true` (`ErrSelfRegisterNotAllowed`). Signer must control `owner_did` (`ErrUnauthorized` via `verifyDidController`). Normalized name must satisfy charset `[a-z0-9_-]+`, namespace `min_length`/`max_length`, and `regex` (`ErrInvalidName`). Slot must be free in any status, including tombstoned (`ErrNameTaken`). Server writes `verified=false`, `source="self"`, `status=ACTIVE`. Response returns `normalized_name`.

### MsgRegisterNameByRegistrar
- **Purpose:** A registrar registers a name on behalf of any DID (custodial/USSD/no-gas users and attested names).
- **Signer / auth:** `registrar` signs; must be in the namespace's `registrar_accounts`. Registrar-only.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| registrar | string | yes | Bech32 address signing; must be in `registrar_accounts`. |
| namespace | string | yes | Target namespace. |
| name | string | yes | Display name; normalized server-side. |
| owner_did | string | yes | DID that will own the record; registrar need NOT control it. |
| verified | bool | no | Marks the record attested. Self-register sets false; registrars typically set true after off-chain proof. |
| evidence_hash | string | no | Optional hash of off-chain attestation evidence. |
| source | string | no | Free-form tag describing the verification source. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgRegisterNameByRegistrar',
  value: ixo.names.v1beta1.MsgRegisterNameByRegistrar.fromPartial({
    registrar: registrarAddress,
    namespace: 'twitter',
    name: 'alice',
    ownerDid: 'did:ixo:abc123',
    verified: true,
  }),
};
```
- **Gotchas:** `registrar` must be in `registrar_accounts` (`ErrNotRegistrar`). `owner_did` must reference an existing IID document (`ErrInvalidDID`) but the registrar is NOT required to control it. Same normalization/charset/length/regex and `ErrNameTaken` checks as `MsgRegisterName`. `evidence_hash` max `MaxNameRecordEvidenceHashLength = 256`, `source` max `MaxNameRecordSourceLength = 64` — enforced in both `ValidateBasic` and the keeper (`ValidateRecordMetadata`, for Wasm sub-msg dispatch that bypasses ante). Server sets `verified_by = registrar`, `status=ACTIVE`. Response returns `normalized_name`.

### MsgUpdateNameByRegistrar
- **Purpose:** Update verification metadata (`verified`, `verified_by`, `evidence_hash`, `source`) of an existing record. Owner DID and status are NOT changed here.
- **Signer / auth:** `registrar` signs; must be in the namespace's `registrar_accounts`. Registrar-only.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| registrar | string | yes | Bech32 signer; must be in `registrar_accounts`. |
| namespace | string | yes | Target namespace. |
| normalized_name | string | yes | Canonical name to update. |
| verified | bool | no | New `verified` value; set false to retract a prior attestation. |
| evidence_hash | string | no | Replaces the existing evidence hash. |
| source | string | no | Replaces the existing source tag. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgUpdateNameByRegistrar',
  value: ixo.names.v1beta1.MsgUpdateNameByRegistrar.fromPartial({
    registrar: registrarAddress,
    namespace: 'twitter',
    normalizedName: 'alice',
    verified: false,
  }),
};
```
- **Gotchas:** Takes `normalized_name`, not a display name (no normalization applied here). `registrar` must be in `registrar_accounts` (`ErrNotRegistrar`); record must exist (`ErrNameNotFound`). Same `evidence_hash` (256) / `source` (64) length caps, dual-enforced. Sets `verified_by` to this registrar. Does NOT change `owner_did` (use `MsgTransferName`) or `status` (use `MsgSetNameStatus`); display name and namespace are immutable.

### MsgTransferName
- **Purpose:** Reassign a record's `owner_did` to another DID.
- **Signer / auth:** `signer` signs; must control the current `owner_did`, OR be a registrar when the namespace has `allow_registrar_override=true`. User-facing or registrar (override).
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| signer | string | yes | Bech32 address; current-owner controller, or registrar with override. |
| namespace | string | yes | Namespace containing the record. |
| normalized_name | string | yes | Canonical name to transfer. |
| new_owner_did | string | yes | DID that will own the name after transfer. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgTransferName',
  value: ixo.names.v1beta1.MsgTransferName.fromPartial({
    signer: signerAddress,
    namespace: 'yoid',
    normalizedName: 'alice',
    newOwnerDid: 'did:ixo:def456',
  }),
};
```
- **Gotchas:** Namespace and record must exist (`ErrNamespaceNotFound`, `ErrNameNotFound`). Auth requires `verifyDidController(owner_did, signer)` OR (`HasRegistrar` AND `allow_registrar_override`) — else `ErrUnauthorized`. `new_owner_did` must reference an existing IID document (`ErrInvalidDID`) and must differ from the current owner (`ErrInvalidRequest`). Updates the reverse index and dual-emits `NameUpdatedEvent` + `NameTransferredEvent`.

### MsgSetNameStatus
- **Purpose:** Change a record's lifecycle status (suspend, revoke, tombstone, or restore to active).
- **Signer / auth:** `signer` signs; must be a namespace registrar OR the gov module address (no `allow_registrar_override` requirement). Registrar or governance.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| signer | string | yes | Bech32 address; namespace registrar or gov authority. |
| namespace | string | yes | Target namespace. |
| normalized_name | string | yes | Canonical name. |
| status | NameStatus | yes | Target lifecycle status (enum). |
| reason | string | no | Free-form string surfaced in the audit event. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
```ts
const msg = {
  typeUrl: '/ixo.names.v1beta1.MsgSetNameStatus',
  value: ixo.names.v1beta1.MsgSetNameStatus.fromPartial({
    signer: registrarAddress,
    namespace: 'yoid',
    normalizedName: 'alice',
    status: ixo.names.v1beta1.NameStatus.NAME_STATUS_SUSPENDED,
    reason: 'ToS violation',
  }),
};
```
- **Gotchas:** `status` must be one of ACTIVE/SUSPENDED/REVOKED/TOMBSTONED; `NAME_STATUS_UNSPECIFIED` (zero value) is rejected with `ErrInvalidStatusTransition` (checked in both `ValidateBasic` and keeper, for Wasm bypass). `signer` must be a registrar or `k.authority` else `ErrUnauthorized`. Record must exist (`ErrNameNotFound`). Current status must not be `TOMBSTONED` (terminal) and must differ from the target (`ErrInvalidStatusTransition`). Dual-emits `NameUpdatedEvent` + `NameStatusChangedEvent`.

## Queries
All query CLI commands ARE registered via autocli under `ixod query names`. Names passed to `resolve` are normalized server-side; `get` takes a `normalized_name`.

| Query | gRPC method | CLI | Args | Returns |
|-------|-------------|-----|------|---------|
| Namespace | `Namespace` | `ixod query names namespace [name]` | `name` | `Namespace` |
| Namespaces | `Namespaces` | `ixod query names namespaces` | pagination | `repeated Namespace` + page |
| ResolveName | `ResolveName` | `ixod query names resolve [namespace] [name]` | `namespace`, `name` | `NameRecord` (ACTIVE only; NotFound otherwise) |
| GetName | `GetName` | `ixod query names get [namespace] [normalized_name]` | `namespace`, `normalized_name` | `NameRecord` (any status) |
| NamesByNamespace | `NamesByNamespace` | `ixod query names list-by-namespace [namespace]` | `namespace`, pagination | `repeated NameRecord` + page |
| NamesByOwner | `NamesByOwner` | `ixod query names list-by-owner [owner_did]` | `owner_did`, pagination | `repeated NameRecord` + page |

## Events
Typed events from `event.proto`:
- `NamespaceCreatedEvent` — emitted by `MsgCreateNamespace`; carries the full `namespace` and the gov `authority`.
- `NamespaceUpdatedEvent` — emitted by `MsgUpdateNamespace`; carries the full `namespace` and `authority`.
- `NameRegisteredEvent` — emitted by `MsgRegisterName` / `MsgRegisterNameByRegistrar`; carries the full `record` and `registered_by` (signer or registrar).
- `NameUpdatedEvent` — emitted by `MsgUpdateNameByRegistrar`, and (dual-emit) by `MsgTransferName` and `MsgSetNameStatus`; carries the post-write `record` and `updated_by`.
- `NameTransferredEvent` — emitted by `MsgTransferName`; carries `namespace`, `normalized_name`, `from_owner_did`, `to_owner_did`, `transferred_by`.
- `NameStatusChangedEvent` — emitted by `MsgSetNameStatus`; carries `namespace`, `normalized_name`, `old_status`, `new_status`, `changed_by`, `reason`.

## Module gotchas
- **Mixed authority model:** `MsgCreateNamespace` / `MsgUpdateNamespace` are governance-only (signer must equal the gov module address, reached via `MsgSubmitProposal`). `MsgRegisterName` and `MsgTransferName` are user-facing. `MsgRegisterNameByRegistrar` / `MsgUpdateNameByRegistrar` are registrar-only. `MsgSetNameStatus` accepts registrar OR gov authority.
- **No tx CLI:** the Tx autocli descriptor has an empty command list and `GetTxCmd` returns nil — all state-changing messages must be constructed via the SDK (or a gov proposal for namespace messages). Queries DO have CLI.
- **Names relate to DIDs, not addresses:** records bind to an `owner_did` (an IID document), not directly to a bech32 account. Authorisation to register/transfer is mediated through `verifyDidController` against the iid module; the bech32 `signer` only matters insofar as it controls the DID (or is a registrar). Target DIDs (`owner_did`, `new_owner_did`) must already exist as IID documents.
- **No hard-delete:** records persist across their lifecycle; `(namespace, normalized_name)` slots are never freed (tombstone is permanent). `ResolveName` returns ACTIVE only; `GetName` returns any status.
- **Reserved (unused in v1):** `NameRecord.valid_until` (always 0) and `Namespace.allow_expiry` (always false) are hooks for a future renewal flow.
