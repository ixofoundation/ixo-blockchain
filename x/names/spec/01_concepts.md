# Concepts

## Names and namespaces

A **name** is a human-readable handle (e.g. `alice`) that resolves to an `owner_did` (e.g. `did:ixo:abc…`). Names always live inside a **namespace** — a governance-managed bucket that defines:

- the validation rules for member names (length bounds, optional regex, charset),
- the set of registrar accounts allowed to operate on member names on behalf of users,
- whether users may self-register, and
- whether registrars can override owner consent for transfers and updates.

Uniqueness is **scoped to the namespace**: a single chain can have `yoid:alice` and `twitter:alice` simultaneously, owned by different DIDs. The full address of a name is the `(namespace, normalized_name)` tuple.

Namespaces are created and updated only through chain governance — the gov module address is the sole valid signer for `MsgCreateNamespace` and `MsgUpdateNamespace`. This makes a new namespace a deliberate, auditable event recorded as a passed proposal.

## Normalisation and uniqueness

Every name has two forms:

- **`display_name`** — the original case-preserving form supplied by the registrant (e.g. `AliceCAPS`).
- **`normalized_name`** — the canonical lookup form: trimmed of whitespace and lower-cased ASCII (e.g. `alicecaps`).

Uniqueness is enforced on `normalized_name` only. Two display variants that normalise to the same string (`alice`, `Alice`, ` ALICE `) cannot coexist in the same namespace. Resolution queries normalise the supplied name server-side before lookup, so callers may pass either form.

The v1 normalisation rule is intentionally minimal — trim + ASCII lowercase — combined with a charset rejection of any non-ASCII input (`[a-z0-9_-]+`). This keeps the rule auditable, deterministic, and free of Unicode-table dependencies. Confusables defence (homoglyph attacks across scripts) is deferred until a future Unicode/NFKC pass; for now non-ASCII registrations are simply refused.

Each namespace can layer additional restrictions on top of the chain-wide charset:
- `min_length` / `max_length` bounds on the normalised name.
- An optional Go-regexp `regex` that the normalised name must additionally match.

## Self-register vs registrar-only namespaces

Namespaces fall into two operational modes determined by the `allow_self_register` flag:

- **Self-register** namespaces (e.g. a public handle namespace like `yoid`) accept `MsgRegisterName` directly from any user — provided the signer controls the `owner_did` they're registering. Records created this way start with `verified = false` and `source = "self"`.
- **Registrar-only** namespaces (e.g. an attested namespace like `twitter` or a KYC-gated namespace) reject `MsgRegisterName` with `ErrSelfRegisterNotAllowed`. Names in these namespaces only enter through `MsgRegisterNameByRegistrar`, sent by an account in the namespace's `registrar_accounts` list.

`MsgUpdateNamespace` can flip a namespace between modes — gov can lock down a self-register namespace (existing names stay; future self-registrations are refused) or open a registrar-only namespace.

## Registrars and on-behalf registration

A **registrar** is a normal account address listed in a namespace's `registrar_accounts`. Registrars are designed for three use cases:

1. **Off-chain identity attestation** — a registrar (typically an oracle account holding off-chain proof) registers `twitter:alice` for a user's DID after verifying their Twitter OAuth flow off-chain. The on-chain record carries `verified=true`, `verified_by=<registrar address>`, `evidence_hash=<hash of the attestation>`, and `source` describing the attestation source.
2. **Custodial / USSD / no-gas users** — when a user has no IXO for fees or transacts via USSD-mediated custody, the registrar submits the registration on their behalf. The user's DID is still bound as `owner_did`; the registrar's signer is recorded as the actor.
3. **Moderation** — when a name violates ToS or law, registrars can change its status (suspend / revoke / tombstone) without the owner's consent.

A namespace can enable registrar-driven transfers (taking an active name from one DID and giving it to another) by setting `allow_registrar_override = true`. When `false`, only the current owner can transfer.

Registrars are *not* authorised to change `owner_did` directly via `MsgUpdateNameByRegistrar` — that message updates verification metadata only. Ownership change always flows through `MsgTransferName` which has its own auth rules.

## Verified records and off-chain attestation

The chain does not perform OAuth or KYC itself. Verification is a *claim* made by a registrar after off-chain proof, recorded on-chain as the four fields:

- `verified: bool` — has this name been attested to?
- `verified_by: string` — which registrar (or DID) made the attestation?
- `evidence_hash: string` — content hash of the off-chain proof (e.g. the JWT or signed VC).
- `source: string` — free-form tag describing where the proof came from (`twitter-oauth`, `workos`, `ussd`, etc.).

This keeps OAuth providers and KYC vendors fully off-chain — consensus only needs to understand "registrar X attests Y" — while preserving an auditable trail. A registrar wishing to retract an attestation can call `MsgUpdateNameByRegistrar` with `verified=false`.

## Status lifecycle (no hard-delete)

Records are never deleted. Every `NameRecord` has a `status` field (a `NameStatus` enum) that drives its visibility:

- `NAME_STATUS_ACTIVE` — the resolvable state. `ResolveName` returns it; `NamesByNamespace` and `NamesByOwner` include it.
- `NAME_STATUS_SUSPENDED` — temporarily hidden from `ResolveName` (returns NotFound). Can be restored to `ACTIVE` by a registrar.
- `NAME_STATUS_REVOKED` — terminally taken down for the current owner; conventionally used after an owner-side breach. Distinguishable from `SUSPENDED` for audit but treated identically by `ResolveName`. Can be restored by a registrar (the chain doesn't enforce "terminal" semantically — that's for moderation policy).
- `NAME_STATUS_TOMBSTONED` — registrar-level take-down. **Terminal**: any further status transition is rejected with `ErrInvalidStatusTransition`. Used for permanent burns where the `(namespace, normalized_name)` slot must never be reused.

`GetName` returns records regardless of status (audit / moderation surface), while `ResolveName` is `ACTIVE`-only (the application surface).

This avoids the failure modes of silent deletion — name slots that "free up" again let a malicious party re-register a previously banned identity, and disappearing rows lose audit trail. Status transitions are recorded as `NameStatusChangedEvent`s with `reason` for downstream auditing.

## Transfers and registrar override

`MsgTransferName` reassigns a record's `owner_did`. Two authorisation paths:

- **Owner-driven** — the signer controls the current `owner_did` (via an authentication verification method on the IID document, or by being listed as a controller of it). This works in any namespace.
- **Registrar-override** — the signer is a registrar of the namespace AND the namespace has `allow_registrar_override = true`. Used for moderation-driven reassignment (e.g. taking over a squatted handle for a verified entity).

The new `owner_did` must reference an existing IID document; transfers to non-existent DIDs are rejected with `ErrInvalidDID`. Self-transfers (new owner equals old owner) are rejected with `ErrInvalidRequest`.

Transfers always update the reverse-lookup index — see [Reverse lookup by owner DID](#reverse-lookup-by-owner-did).

## Reverse lookup by owner DID

In addition to the forward `(namespace, normalized_name) → NameRecord` mapping, the module maintains a secondary index `(owner_did, namespace, normalized_name) → []` that supports the `NamesByOwner(owner_did)` query. The reverse index is kept in lockstep with the primary record by:

- writing both keys on `MsgRegisterName` / `MsgRegisterNameByRegistrar`,
- removing the old `owner_did` entry and writing the new one on `MsgTransferName`.

Status changes do not touch the reverse index; a suspended or tombstoned record is still listed when querying its owner (use the record's `status` field to filter on the client side).

## Relationship to the iid module

The names module is a **separate** module from `iid`. The two interact one-way:

- The names module imports the iid keeper to verify that a tx signer controls a given DID (`verifyDidController`) and to confirm a target DID exists (`GetDidDocument`).
- The iid module is unaware of the names module — `IidDocument.alsoKnownAs` is unchanged and unused by names. Apps can still set `alsoKnownAs` as a free-text hint; the canonical handle registry is the names module.

This split keeps the iid module focused on W3C DID-core compliance, and lets the names module own its own params, gov messages, and indexes without bloating the iid surface.

## Reserved fields for future expansion

Two `NameRecord` fields are reserved for forward compatibility but **not used in v1**:

- `valid_until: int64` — Unix-second expiry timestamp. Always `0` (no expiry) in v1; ignored by `ResolveName`. The hook for a future renewal flow.
- Per-namespace `allow_expiry: bool` — gates whether non-zero `valid_until` may be set. Always `false` by default. Genesis validation rejects records with `valid_until != 0` in namespaces where `allow_expiry = false`.

Adding paid expiry / renewal is a forward-compatible chain change: a future `MsgRenewName` can be added without touching existing record schemas.

Other deliberately deferred features (callable as forward-compatible additions in a later upgrade): NFKC Unicode normalisation, confusables-table rejection, IBC outpost / ICS-721 portability, on-chain JWT/VC verification.
