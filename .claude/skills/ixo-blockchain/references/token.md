# Token module — `x/token`

**Proto package:** `ixo.token.v1beta1` · **TS typeUrl prefix:** `/ixo.token.v1beta1.` · **CLI:** `ixod tx token …` (no `ixod query token …` — `GetQueryCmd` returns `nil`, query only via gRPC)

## Purpose
The token module manages ixo's "impact tokens" / carbon-style credits as EIP-1155-style multi-tokens. Each token class (a `Token`, e.g. "CARBON") is backed by an on-chain CosmWasm `ixo1155` contract (cw1155, not cw721) that is instantiated automatically when the class is created; the contract holds balances and is what actually mints/burns. The module tracks supply, mint properties, and retirement/cancellation records on chain. A token class is bound to a protocol entity DID (`class`), and mints reference a `collection` DID — so tokens are tied to entities/iid identifiers.

## Concepts & state
- **Token** (`token.proto` `Token`): the token class/collection overview. Keyed by `minter + contract_address`. Holds `class` (entity DID), `name` (unique namespace), `cap`, `supply`, `paused`, `stopped`, and the `retired`/`cancelled`/`transferred` ledgers.
- **TokenProperties** (`token.proto` `TokenProperties`): one per minted token id, keyed by `id`. Persists `index`, `name`, `collection`, and `tokenData` (linked resources surfaced in metadata).
- **TokenData** (`token.proto` `TokenData`): a linked resource — `uri` (e.g. credential `.ipfs`), `encrypted`, `proof`, `type` (should be `application/json`), `id` (entity DID, may be empty).
- **MintBatch** (`tx.proto`): mint input — `name`, `index`, `amount`, `collection`, `token_data`. Token `id` = hex md5 of `name+index` (no separator).
- **TokenBatch** (`tx.proto`): transfer/retire/cancel unit — `id` + `amount` only (no name/index).
- **minter**: address authorized to create the class and mint; also the only signer for pause/stop. Set once at `CreateToken` (no separate "setup minter" message).
- **contract_address**: the cw1155 contract instantiated at `CreateToken` (code id = `Params.ixo1155_contract_code`); referenced by mint/pause/stop and resolved via the token for transfer/retire/cancel.
- **cap / supply**: `cap` 0 = unlimited; `supply` increases on mint, decreases on cancel, **unchanged on retire**.
- **collection**: a DID (e.g. Supamoto Malawi) the minted token is grouped under; stored on `TokenProperties`.

## Messages
Source order from `tx.proto` service `Msg`. Note: `MsgPauseToken` and `MsgStopToken` are **not** registered in `RegisterInterfaces`/amino (`codec.go`) — only Create/Mint/Transfer/Retire/TransferCredit/Cancel are. There is no `MsgSetupMinter` and no `MsgUpdateToken`.

### MsgCreateToken
- **Purpose:** Create a token class and auto-instantiate its `ixo1155` cw1155 contract.
- **Signer / auth:** `minter` signs. CLI overwrites `minter` with the from-address. Requires the `class` DID document to already exist (iid). No authz.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| minter | string | yes | Minter address (signer) |
| class | string (`DIDFragment` casttype) | yes | Token protocol entity DID; must be a valid DID and the DID doc must exist |
| name | string | yes | Token name, unique namespace; must not be empty/duplicate |
| description | string | no | Arbitrary description |
| image | string | yes | Image URL; must be a valid RFC3986 URI |
| token_type | string | no | Token type, e.g. `ixo1155` |
| cap | string (`math.Uint`, non-nullable) | yes | Max mintable for this name; `0` = unlimited |

- **CLI:** `ixod tx token create [create_token_doc] [flags]` — `create_token_doc` is raw JSON of `MsgCreateToken` (the `minter` field is ignored/overwritten).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgCreateToken',
    value: ixo.token.v1beta1.MsgCreateToken.fromPartial({
      minter: minterAddress,
      class: classDid,
      name: 'CARBON',
      image: 'https://...',
      cap: '0',
    }),
  };
  ```
- **Gotchas:** Fails with `class did document not found` if `class` DID doc absent; `token name is already taken` (`ErrTokenNameDuplicate`) on duplicate name. `ValidateBasic` requires non-empty `name`, valid DID, valid RFC3986 `image`.

### MsgMintToken
- **Purpose:** Mint tokens on the class's cw1155 contract; creates a `TokenProperties` per id.
- **Signer / auth:** `minter` signs (CLI overwrites with from-address). Authz: gated by `MintAuthorization` (grantee mints under granted constraints). Token must not be `paused`/`stopped`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| minter | string | yes | Minter address (signer) |
| contract_address | string | yes | The class's cw1155 contract address |
| owner | string | yes | Recipient address the tokens are minted to |
| mint_batch | repeated MintBatch | yes | Batches to mint (non-empty) |

  **MintBatch:** `name` string (must equal Token.name); `index` string (hexstring id); `amount` string `math.Uint` (>0); `collection` string (collection DID); `token_data` repeated TokenData.

- **CLI:** `ixod tx token mint [mint_token_doc] [flags]` — raw JSON of `MsgMintToken` (`minter` overwritten).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgMintToken',
    value: ixo.token.v1beta1.MsgMintToken.fromPartial({
      minter: minterAddress,
      contractAddress: contractAddress,
      owner: ownerAddress,
      mintBatch: [{ name: 'CARBON', index: '01', amount: '100', collection: collectionDid, tokenData: [] }],
    }),
  };
  ```
- **Gotchas:** Class must exist first (`CreateToken`). `ValidateBasic` rejects empty `mint_batch`, empty batch `name`/`index`, or zero `amount`. Server rejects `token name is not same as class token name` (`ErrTokenNameIncorrect`) and over-cap mints (`ErrTokenAmountIncorrect`) when `supply+amount > cap` (cap≠0). Returns `ErrTokenPaused`/`ErrTokenStopped`. id = `md5(name+index)` hex; minting the same `name+index` again adds to that id's supply.

### MsgTransferToken
- **Purpose:** Transfer minted tokens between owners on the cw1155 contract.
- **Signer / auth:** `owner` signs (CLI overwrites with from-address). No authz. Contract resolved from `tokens[0].id`'s Token; **all tokens must be in the same contract**.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| owner | string | yes | Current owner (signer) |
| recipient | string | yes | Receiving address |
| tokens | repeated TokenBatch | yes | `id`+`amount` batches (non-empty); same contract |

- **CLI:** `ixod tx token transfer [transfer_token_doc] [flags]` — raw JSON of `MsgTransferToken` (`owner` overwritten).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgTransferToken',
    value: ixo.token.v1beta1.MsgTransferToken.fromPartial({
      owner: ownerAddress,
      recipient: recipientAddress,
      tokens: [{ id: tokenId, amount: '50' }],
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` rejects empty `tokens`, empty `id`, or zero `amount`. Underlying cw1155 `batch_send_from` fails if `owner` lacks balance.

### MsgRetireToken
- **Purpose:** Permanently retire (burn) tokens as an impact offset/claim; records the retirement on the Token.
- **Signer / auth:** `owner` signs (CLI overwrites with from-address). No authz. Contract from `tokens[0].id`; all tokens same contract.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| owner | string | yes | Token owner (signer) |
| tokens | repeated TokenBatch | yes | `id`+`amount` batches to retire (non-empty); same contract |
| jurisdiction | string | yes | Owner jurisdiction `<country-code>[-<sub-national-code>[ <postal-code>]]`; required, must be non-empty |
| reason | string | no | Arbitrary reason for retiring |

- **CLI:** `ixod tx token retire-token [retire_token_doc] [flags]` — raw JSON of `MsgRetireToken` (`owner` overwritten).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgRetireToken',
    value: ixo.token.v1beta1.MsgRetireToken.fromPartial({
      owner: ownerAddress,
      tokens: [{ id: tokenId, amount: '10' }],
      jurisdiction: 'US-CO',
      reason: 'offset Q1 footprint',
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires non-empty `jurisdiction`, non-empty `tokens`, non-empty `id`, `amount`>0. Burns on cw1155 and appends `TokensRetired{id,amount,reason,jurisdiction,owner}` to `Token.retired`. **Does NOT reduce `Token.supply`** (retired remains counted in supply) — this is the key difference from cancel. Permanent/irreversible.

### MsgTransferCredit
- **Purpose:** Burn ("transfer") credits under a recorded authorization; logs to `Token.transferred`.
- **Signer / auth:** `owner` signs. No CLI command. Contract from `tokens[0].id`; same contract.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| owner | string | yes | Token owner (signer) |
| tokens | repeated TokenBatch | yes | `id`+`amount` batches (non-empty); same contract |
| jurisdiction | string | yes | Owner jurisdiction (same format/rules as retire); required |
| reason | string | no | Arbitrary reason |
| authorization_id | string | yes | Id of the authorization used for the credit transfer; must be non-empty |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgTransferCredit',
    value: ixo.token.v1beta1.MsgTransferCredit.fromPartial({
      owner: ownerAddress,
      tokens: [{ id: tokenId, amount: '10' }],
      jurisdiction: 'US-CO',
      reason: 'credit transfer',
      authorizationId: authId,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires non-empty `jurisdiction` and `authorization_id`. Burns like retire and appends `CreditsTransferred{...,authorization_id}` to `Token.transferred`; does **not** reduce `Token.supply`.

### MsgCancelToken
- **Purpose:** Cancel (burn) tokens, removing them from tradable supply; records on the Token.
- **Signer / auth:** `owner` signs (CLI overwrites with from-address). No authz. Contract from `tokens[0].id`; same contract.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| owner | string | yes | Token owner (signer) |
| tokens | repeated TokenBatch | yes | `id`+`amount` batches to cancel (non-empty); same contract |
| reason | string | no | Arbitrary reason for cancelling |

- **CLI:** `ixod tx token cancel-token [cancel_token_doc] [flags]` — raw JSON of `MsgCancelToken` (`owner` overwritten).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgCancelToken',
    value: ixo.token.v1beta1.MsgCancelToken.fromPartial({
      owner: ownerAddress,
      tokens: [{ id: tokenId, amount: '5' }],
      reason: 'issuance error',
    }),
  };
  ```
- **Gotchas:** No `jurisdiction` field (unlike retire). `ValidateBasic` rejects empty `tokens`, empty `id`, zero `amount`. Burns on cw1155, appends `TokensCancelled{id,amount,reason,owner}` to `Token.cancelled`, and **subtracts the amount from `Token.supply`** (the key difference vs retire). Implementation quirk: the server emits a `TokenRetiredEvent` (not a dedicated cancelled event) plus `TokenUpdatedEvent`. Permanent.

### MsgPauseToken
- **Purpose:** Temporarily pause/unpause minting for a token class.
- **Signer / auth:** `minter` signs (CLI sets from-address). Not registered in interface/amino codec. No authz.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| minter | string | yes | Minter address (signer); must match Token.minter |
| contract_address | string | yes | The class's cw1155 contract address |
| paused | bool | yes | `true` to pause minting, `false` to resume |

- **CLI:** `ixod tx token pause-token [contract_address] [paused] [flags]` — positional `contract_address` then `paused` (parsed as bool); `minter` = from-address.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgPauseToken',
    value: ixo.token.v1beta1.MsgPauseToken.fromPartial({
      minter: minterAddress,
      contractAddress: contractAddress,
      paused: true,
    }),
  };
  ```
- **Gotchas:** Token looked up by `minter + contract_address`; `ErrTokenNotFound` if the signer is not the recorded minter for that contract. Only affects `MintToken` (`ErrTokenPaused`).

### MsgStopToken
- **Purpose:** Permanently stop minting for a token class.
- **Signer / auth:** `minter` signs (CLI sets from-address). Not registered in interface/amino codec. No authz.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| minter | string | yes | Minter address (signer); must match Token.minter |
| contract_address | string | yes | The class's cw1155 contract address |

- **CLI:** `ixod tx token stop-token [contract_address] [flags]` — positional `contract_address`; `minter` = from-address.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.token.v1beta1.MsgStopToken',
    value: ixo.token.v1beta1.MsgStopToken.fromPartial({
      minter: minterAddress,
      contractAddress: contractAddress,
    }),
  };
  ```
- **Gotchas:** Always sets `stopped = true` (irreversible — there is no unstop). After stop, `MintToken` returns `ErrTokenStopped`. `ErrTokenNotFound` if signer isn't the recorded minter.

## Authz authorizations
- **MintAuthorization** (`authz.proto`): gates `MsgMintToken` only (`MsgTypeURL` = mint). Fields: `minter` (must equal the msg's `minter`), `constraints` repeated `MintConstraints`.
- **MintConstraints**: `contract_address`, `amount` (`math.Uint`), `name`, `index`, `collection`, `token_data` repeated `TokenData`. On `Accept`, each msg `MintBatch` must exactly match one unused constraint (contract, amount-equal, name, index, collection, and tokenData set); matched constraints are consumed. When all constraints are consumed the grant is deleted; otherwise it's updated. `ValidateBasic` requires ≥1 constraint, each with `amount`>0, valid `contract_address`, non-empty `name`.

## Queries
No CLI (`GetQueryCmd` returns `nil`); gRPC/REST only.

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| Params | `/ixo.token.v1beta1.Query/Params` | — | none | `Params{ixo1155_contract_code}` |
| TokenMetadata | `/ixo.token.v1beta1.Query/TokenMetadata` | — | `id` | name, description, decimals, image, index, `TokenMetadataProperties{class,collection,cap,linkedResources}` |
| TokenList | `/ixo.token.v1beta1.Query/TokenList` | — | `minter`, `pagination` | repeated `Token` (`tokenDocs`) + pagination |
| TokenDoc | `/ixo.token.v1beta1.Query/TokenDoc` | — | `minter`, `contract_address` | single `Token` (`tokenDoc`) |

REST: `/ixo/token/params`, `/ixo/token/metadata/{id}`, `/ixo/token/minter/{minter}`, `/ixo/token/minter/{minter}/{contract_address}`.

## Events
Typed events from `event.proto` (emitted via `EmitTypedEvent(s)`):
- **TokenCreatedEvent** — on `CreateToken`; carries the new `Token`.
- **TokenUpdatedEvent** — on mint/retire/transfer-credit/cancel/pause/stop when the `Token` changes; carries the updated `Token`.
- **TokenMintedEvent** — once per minted batch in `MintToken`: `contract_address`, `minter`, `owner`, `amount`, `tokenProperties`.
- **TokenTransferredEvent** — on `TransferToken`: `owner`, `recipient`, `tokens`.
- **TokenRetiredEvent** — on `RetireToken` (and also emitted by `CancelToken`, a code quirk): `owner`, `tokens`.
- **TokenCancelledEvent** — defined in proto but **not emitted** by `CancelToken` (which emits `TokenRetiredEvent` instead). `owner`, `tokens`.
- **CreditsTransferredEvent** — on `TransferCredit`: `owner`, `tokens`.
- **TokenPausedEvent** — on `PauseToken`: `minter`, `contract_address`, `paused`.
- **TokenStoppedEvent** — on `StopToken`: `minter`, `contract_address`, `stopped`.

## Module gotchas
- Tokens are backed by a CosmWasm **cw1155 (`ixo1155`)** contract — *not* cw721 — instantiated automatically by `CreateToken` using `Params.ixo1155_contract_code`. The contract address is stored on the `Token` and is the on-chain source of balances.
- Lifecycle: `CreateToken` (class + contract + minter set in one step — **no separate minter-setup message**) → `MintToken` → `TransferToken` → `RetireToken` / `CancelToken`. Minting requires the `class` entity DID document to exist; mints reference a `collection` DID.
- **Retire vs cancel (precise):** both burn on the cw1155 contract and are permanent. **Retire** records `TokensRetired` (with `jurisdiction`, required) in `Token.retired` and leaves `Token.supply` unchanged — it marks tokens as a permanent impact claim/offset while keeping them counted in total supply. **Cancel** records `TokensCancelled` (no `jurisdiction`) in `Token.cancelled` and **subtracts the amount from `Token.supply`**, removing them from tradable supply. So: retire = offset claim, supply preserved; cancel = supply reduction.
- `TransferCredit` is a third burn path (like retire but keyed to an `authorization_id`, logged to `Token.transferred`); also leaves supply unchanged.
- All of transfer/retire/transfer-credit/cancel resolve the contract from `tokens[0].id` and assume every batch is in the same contract.
- Token `id` is deterministic: `hex(md5(name+index))`; reusing a `name+index` mints into the same id.
- `MsgPauseToken`/`MsgStopToken` are not registered in `RegisterInterfaces`/amino — they are reachable via CLI but may not deserialize through paths relying on the interface registry. `MsgTransferCredit` has no CLI command.
- Params can only be changed via the legacy gov proposal `update-token-params` (content `SetTokenContractCodes{ixo1155_contract_code}`): `ixod tx gov submit-legacy-proposal update-token-params [ixo1155-code-id] --title --description --deposit`. It is *not* under `ixod tx token`.
