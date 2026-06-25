# Smart-account module — `x/smart-account`

**Proto package:** `ixo.smartaccount.v1beta1` · **TS typeUrl prefix:** `/ixo.smartaccount.v1beta1.` · **CLI:** no `client/cli` directory and `GetTxCmd()`/`GetQueryCmd()` return `nil`, **but** the module ships an `autocli.go` (`AutoCLIOptions`) that auto-generates `ixod tx smartaccount …` / `ixod query smartaccount …` commands. Prefer constructing messages via the SDK; the autocli `data` arg takes raw bytes and is awkward for structured authenticator configs.

> Directory is `x/smart-account` (hyphen) but the Go package inside is `authenticator`, the proto package is `ixo.smartaccount.v1beta1`, and the on-chain module name / store key is `smartaccount`.

## Purpose
The smart-account module is an account-abstraction / authenticator framework (a port of the Osmosis smart-account design). It replaces the default Cosmos SDK signature check with a pluggable set of per-account **authenticators**, enabling session/hot keys scoped to specific messages, WebAuthn passkeys, spend-limit and cosigner contracts, partitioned multisig, and arbitrary CosmWasm authentication logic. It runs as an ante/post handler behind a circuit-breaker param and is fully opt-in per transaction.

## Concepts & state
- **Authenticator** — a Go struct implementing the `Authenticator` interface (`Type`, `StaticGas`, `Initialize`, `Authenticate`, `Track`, `ConfirmExecution`, `OnAuthenticatorAdded`, `OnAuthenticatorRemoved`). Defined in `x/smart-account/authenticator/iface.go`.
- **AccountAuthenticator** (`models.proto`) — the stored per-account instance: `id` (uint64, globally-unique incrementing), `type` (string, e.g. `SignatureVerification`), `config` (bytes — the `data` passed in `MsgAddAuthenticator`, replayed into `Initialize`). Stored keyed by `account` + `id`.
- **Authenticator types** (registered in `app/keepers/keepers.go` via `InitializeAuthenticators`): `SignatureVerification`, `MessageFilter`, `AllOf`, `AnyOf`, `PartitionedAnyOf`, `PartitionedAllOf`, `CosmwasmAuthenticatorV1`, `AuthnVerification`. See [Authenticator types](#authenticator-types).
- **AuthenticatorManager** (`authenticator/manager.go`) — in-memory registry mapping type string → base authenticator code; `MsgAddAuthenticator` rejects any `authenticator_type` not registered here. Not chain state; configured at app wiring.
- **Active state / circuit breaker** — `Params.is_smart_account_active`. When `false`, the module is bypassed and classic SDK auth is used. Toggled by `MsgSetActiveState` (see auth rules there) or governance param change.
- **Authenticator ids are global** — ids increment across all accounts (user1's first authenticator = id 1, user2's = id 2, user1's second = id 3). `FirstAuthenticatorId = 1`.
- **Composite ids** — a composite authenticator with id `a` passes id `a.i` to its `i`-th sub-authenticator (recursively `a.i.j`).
- **Params** (`params.proto`): `maximum_unauthenticated_gas` (uint64 — gas cap for authenticating before the fee payer is verified, spam protection), `is_smart_account_active` (bool), `circuit_breaker_controllers` (repeated string — addresses allowed to set active=false without governance).
- **TxExtension** (`tx.proto`) — `selected_authenticators` (repeated uint64): the chosen authenticator id **per message**, carried in the tx's `non_critical_extension_options`. Absent → classic SDK auth.

## Messages
The proto `service Msg` RPC names are `AddAuthenticator` / `RemoveAuthenticator` / `SetActiveState`; the request message types are `MsgAddAuthenticator` / `MsgRemoveAuthenticator` / `MsgSetActiveState`. `ValidateBasic` on all three only validates that `sender` is a valid bech32 address.

### MsgAddAuthenticator
- **Purpose:** Register a new authenticator instance on the sender's account.
- **Signer / auth:** `sender` signs (proto `cosmos.msg.v1.signer = "sender"`). The account adds the authenticator to **its own** account.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `sender` | string | **yes** | Bech32 (`ixo…`) account that owns the new authenticator. |
| `authenticator_type` | string | **yes** | Must exactly match a type registered in the `AuthenticatorManager` (see list above). Unregistered → error `authenticator type %s is not registered`. |
| `data` | bytes | **yes** (per type) | Type-specific config bytes. Validated by that type's `OnAuthenticatorAdded` and stored as `AccountAuthenticator.config`. Encoding differs per type — see [Authenticator types](#authenticator-types). |

- **CLI:** `ixod tx smartaccount add-authenticator [authenticator-type] [data] --from …` (autocli; `data` is raw bytes — impractical for JSON/proto configs, use the SDK).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.smartaccount.v1beta1.MsgAddAuthenticator',
    value: ixo.smartaccount.v1beta1.MsgAddAuthenticator.fromPartial({
      sender: 'ixo1...',
      authenticatorType: 'SignatureVerification',
      data: pubKeyBytes, // Uint8Array; meaning depends on authenticatorType
    }),
  };
  ```
- **Gotchas:** The module must be active (`is_smart_account_active = true`) or the msg errors `smartaccount module is not active` — this is independent of whether a tx *uses* authenticators. `data` semantics are entirely type-dependent (raw 33-byte secp256k1 pubkey for `SignatureVerification`, proto-encoded `AuthnPubKey` for `AuthnVerification`, JSON for `MessageFilter`/`CosmwasmAuthenticatorV1`, JSON array of sub-authenticators for `AllOf`/`AnyOf`/`Partitioned*`). On success emits `AuthenticatorAddedEvent` and returns `success = true`.

### MsgRemoveAuthenticator
- **Purpose:** Remove an authenticator (by id) from the sender's account.
- **Signer / auth:** `sender` signs. Only the owning account can remove its authenticators (store key is scoped to `sender`).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `sender` | string | **yes** | Bech32 account that owns the authenticator. |
| `id` | uint64 | **yes** | The global `AccountAuthenticator.id` to remove. |

- **CLI:** `ixod tx smartaccount remove-authenticator [id] --from …` (autocli).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.smartaccount.v1beta1.MsgRemoveAuthenticator',
    value: ixo.smartaccount.v1beta1.MsgRemoveAuthenticator.fromPartial({
      sender: 'ixo1...',
      id: 1n, // uint64
    }),
  };
  ```
- **Gotchas:** Module must be active or it errors `smartaccount module is not active`. Errors if no authenticator with that `id` exists for the account. The authenticator's `OnAuthenticatorRemoved` runs first and **can block removal** (used sparingly). Emits `AuthenticatorRemovedEvent`, returns `success = true`.

### MsgSetActiveState
- **Purpose:** Flip the global circuit breaker (`Params.is_smart_account_active`). Primarily for emergency disable.
- **Signer / auth:** `sender` signs, but msg_server enforces extra authorization:
  - `active = true` → signer **must** equal the keeper's `CircuitBreakerGovernor` (the gov module address) — i.e. only re-enabling via governance.
  - `active = false` → signer must be one of `Params.circuit_breaker_controllers` (disable without governance).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `sender` | string | **yes** | Bech32 signer; must satisfy the auth rule above for the requested `active` value. |
| `active` | bool | **yes** | New value for `is_smart_account_active`. |

- **CLI:** `ixod tx smartaccount set-active-state [active] --from …` (autocli).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.smartaccount.v1beta1.MsgSetActiveState',
    value: ixo.smartaccount.v1beta1.MsgSetActiveState.fromPartial({
      sender: 'ixo1...',
      active: false,
    }),
  };
  ```
- **Gotchas:** Unauthorized signer → `signer is not the circuit breaker governor` (for `true`) or `signer is not a circuit breaker controller` (for `false`). Emits `AuthenticatorSetActiveStateEvent`. Response is empty.

## Authenticator types
The value of `MsgAddAuthenticator.data` (and thus stored `config`) for each registered type. Type strings are returned by each implementation's `Type()` method (verbatim below). All composite configs are JSON arrays of `{ "type": <string>, "config": <bytes> }` objects with **at least 2** sub-authenticators.

- **`SignatureVerification`** (`signature_authenticator.go`, default) — `data` = raw secp256k1 public-key bytes; **must be exactly 33 bytes** (`secp256k1.PubKeySize`), else `OnAuthenticatorAdded` rejects it. Verifies the tx signature against this pubkey.
- **`MessageFilter`** (`message_filter.go`) — `data` = a JSON pattern object. The incoming message (marshalled to JSON) must be a **superset** of the pattern (pattern fields must match; unspecified fields ignored). Numbers **must be encoded as strings**, not JSON floats (`OnAuthenticatorAdded` rejects floats). E.g. `{"@type":"/cosmos.bank.v1beta1.MsgSend","amount":[{"denom":"uixo"}]}`.
- **`AllOf`** (`all_of.go`) — `data` = JSON array of sub-authenticator init data (`[{ "type", "config" }, …]`, ≥2). Authenticates iff **all** sub-authenticators pass. Single signature passed to every sub-authenticator.
- **`AnyOf`** (`any_of.go`) — same JSON array shape. Authenticates iff **any** sub-authenticator passes. Single shared signature.
- **`PartitionedAnyOf`** (`any_of.go`, `NewPartitionedAnyOf`) — as `AnyOf`, but the tx signature must be a JSON array `[sig0, sig1, …]` (one per sub-authenticator) split via `splitSignatures`.
- **`PartitionedAllOf`** (`all_of.go`, `NewPartitionedAllOf`) — as `AllOf` with partitioned (JSON-array) signatures; this is how a real multisig is expressed.
- **`CosmwasmAuthenticatorV1`** (`cosmwasm.go`) — `data` = JSON `{ "contract": "<bech32_contract_addr>", "params": <byte_array> }`. `contract` is required and must be a valid bech32 address (contract must already be instantiated); `params` is optional but, if present, **must be valid JSON bytes** (user-specific config passed to the contract). The contract is invoked via `sudo` with `AuthenticatorSudoMsg` variants (`OnAuthenticatorAdded`, `OnAuthenticatorRemoved`, `Authenticate`, `Track`, `ConfirmExecution`).
- **`AuthnVerification`** (`authn_authenticator.go`) — WebAuthn/passkey. `data` = a proto-encoded `ixo.smartaccount.crypto.AuthnPubKey` (TS: `ixo.smartaccount.crypto.AuthnPubKey`) with fields `key_id` (string, credential id — required non-empty), `cose_algorithm` (int32 — only `-7` ES256 or `-257` RS256 supported), `key` (bytes — public key in DER/PKIX; length-checked and parsed via `x509.ParsePKIXPublicKey`). Signature is verified against the SHA-256 hash of the sign bytes.

> The `crypto.proto` `AuthnPubKey` has exactly three fields (`key_id`, `cose_algorithm`, `key`). The `relayingPartyId`/`relaying_party_id` field shown in `README_authn_authenticator.md`'s JS example does **not** exist in the proto — do not set it.

## Queries
gRPC only. (autocli also exposes these as `ixod query smartaccount …` subcommands; there is no hand-written `client/cli`.)

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| Module params | `ixo.smartaccount.v1beta1.Query/Params` | `ixod query smartaccount params` (autocli) | — | `Params` |
| Single authenticator | `ixo.smartaccount.v1beta1.Query/GetAuthenticator` | `ixod query smartaccount get-authenticator [account] [id]` (autocli) | `account` (string), `authenticator_id` (uint64) | `AccountAuthenticator` |
| All for account | `ixo.smartaccount.v1beta1.Query/GetAuthenticators` | `ixod query smartaccount get-authenticators [account]` (autocli) | `account` (string) | repeated `AccountAuthenticator` |

## Events
Typed events (`event.proto`), emitted from `keeper/msg_server.go`:
- **`AuthenticatorAddedEvent`** — on add. Fields: `sender`, `authenticator_type`, `authenticator_id` (string).
- **`AuthenticatorRemovedEvent`** — on remove. Fields: `sender`, `authenticator_id` (string).
- **`AuthenticatorSetActiveStateEvent`** — on circuit-breaker change. Fields: `sender`, `is_smart_account_active` (bool).

## Module gotchas
- **This changes how transactions are authenticated.** It is wired as an ante handler (plus a post handler for `ConfirmExecution`) behind the `is_smart_account_active` circuit breaker. With the breaker off, classic SDK auth runs and authenticators are ignored.
- **Opt-in per tx via `selected_authenticators`.** To authenticate a tx with a custom authenticator, add a `/ixo.smartaccount.v1beta1.TxExtension` to the tx body's `non_critical_extension_options`, with one `selected_authenticators` entry (the authenticator `id`) **per message**. If absent or not one-per-message, the module falls back to classic SDK auth. Most txs use no extension and the default `SignatureVerification` flow.
- **Tx restrictions when using authenticators:** each message may have only one signer, and the fee payer must be the first signer of the first message.
- **Gas before fee-payer auth is capped** by `maximum_unauthenticated_gas` (anti-spam), relevant for expensive (e.g. CosmWasm) authenticators.
- **Hook call guarantees are weak** inside composites: `Track` runs on all sub-authenticators, but `ConfirmExecution` is stateless w.r.t. which sub-authenticator authenticated, so an `AnyOf` may call `Authenticate` on one branch and `ConfirmExecution` on another — authenticator authors must not assume both run on the same instance.
- **`Track` state is not reverted** even if message execution later fails; `Authenticate`/`ConfirmExecution` state changes are discarded/reverted on failure.
- This is advanced functionality; for a normal address signing a normal tx, none of the above applies.
