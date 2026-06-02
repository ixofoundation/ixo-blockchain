# Entity module — `x/entity`

**Proto package:** `ixo.entity.v1beta1` · **TS typeUrl prefix:** `/ixo.entity.v1beta1.` · **CLI:** `ixod tx entity …` / `ixod query entity …`

## Purpose
An entity is a sovereign "digital twin domain" in the Spatial Web. Creating one performs a trifecta atomically: an iid DID document (id `did:ixo:entity:<hash>`) is generated via the `iid` module, a CW721 NFT is minted to represent ownership, and an `Entity` record is stored in the entity KV store with supplementary metadata. The three share the same deterministically-generated on-chain id. The entity therefore *wraps* an iid DID document — its controllers, services, verification methods, linked resources/claims/entities and accorded rights all live on the underlying `IidDocument`, while NFT ownership governs transfer and account control.

## Concepts & state
- **Entity** (`entity.proto`): KV record `0x01 | entityId(DID) -> Entity`. Fields: `id`, `type`, `start_date`, `end_date`, `status` (int32), `relayer_node`, `credentials` (repeated string), `entity_verified` (bool), `metadata` (EntityMetadata), `accounts` (repeated EntityAccount).
- **EntityMetadata**: `version_id`, `created`, `updated` (timestamps). Bumped on every create/update.
- **EntityAccount**: `name`, `address` — a Cosmos module account derived deterministically from the entity DID + name. A default `admin` account (`EntityAdminAccountName = "admin"`) is created at entity creation.
- **iid DID document**: the entity *owns* an `IidDocument` with the same id; it holds controllers, verificationMethod, service, linkedResource, linkedClaim, accordedRight, linkedEntity, alsoKnownAs. Entity create/transfer mutate this document.
- **NFT backing**: a CW721 token (token_id = entity DID) minted on the contract at `Params.nftContractAddress` by `Params.nftContractMinter`. NFT owner == entity owner; transfer moves the NFT.
- **Controllers / owner**: `MsgUpdateEntity`/`MsgTransferEntity` authorize via the iid document's controller list + `Authentication` relationship; account ops authorize via NFT ownership (`CheckIfOwner`).
- **Relayer node**: the operator DID (`relayer_node`) through which the entity was created; only it may flip `entity_verified`.
- **Params** (`entity.proto`): `nftContractAddress`, `nftContractMinter`, `createSequence` (uint64, drives id generation). Set via gov proposal (see below).

## Messages

### MsgCreateEntity
- **Purpose:** Create an entity: generate its iid DID document, mint the CW721 NFT to `owner_address`, create the `admin` module account, and store the `Entity`.
- **Signer / auth:** `owner_address` signs. `relayer_node` DID **must already exist** as an iid document (checked in msg_server). `owner_did` and `relayer_node` must be valid DIDs. The entity's own DID does NOT pre-exist (it is generated and must be free). `Params.nftContractAddress` must be set or the tx errors.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `entity_type` | string | yes (recommended) | Entity type, e.g. `protocol`, `dao`, `asset/device`. |
| `entity_status` | int32 | no | Implementer-defined status. |
| `controller` | repeated string | no | Controller DIDs; each must be a valid DID. Added to iid controllers along with the entity DID and owner DID. |
| `context` | repeated `ixo.iid.v1beta1.Context` | no | JSON-LD contexts. Each: `key` (string), `val` (string). |
| `verification` | repeated `ixo.iid.v1beta1.Verification` | **yes** | Verification methods/relationships; must be non-empty. Each: `relationships` (repeated string), `method` (VerificationMethod), `context` (repeated string). |
| `service` | repeated `ixo.iid.v1beta1.Service` | no | Each: `id`, `type`, `serviceEndpoint` (all string). |
| `accorded_right` | repeated `ixo.iid.v1beta1.AccordedRight` | no | Each: `type`, `id`, `mechanism`, `message`, `service` (all string). |
| `linked_resource` | repeated `ixo.iid.v1beta1.LinkedResource` | no | Each: `type`, `id`, `description`, `mediaType`, `serviceEndpoint`, `proof`, `encrypted`, `right` (all string). |
| `linked_entity` | repeated `ixo.iid.v1beta1.LinkedEntity` | no | Each: `type`, `id`, `relationship`, `service` (all string). |
| `start_date` | google.protobuf.Timestamp | no | stdtime. |
| `end_date` | google.protobuf.Timestamp | no | stdtime. |
| `relayer_node` | string | **yes** | Operator DID; must already exist on-chain. |
| `credentials` | repeated string | no | CID/hash of public verifiable credentials. |
| `owner_did` | string (`iidtypes.DIDFragment`) | **yes** | Owner DID; added to iid controllers; used as event signer. |
| `owner_address` | string | **yes** | Cosmos address signing the tx; NFT mint recipient. |
| `data` | bytes (`encoding/json.RawMessage`) | no | Extension data; passed verbatim as NFT `extension`. |
| `alsoKnownAs` | string | no | iid `alsoKnownAs`. |
| `linked_claim` | repeated `ixo.iid.v1beta1.LinkedClaim` | no | Each: `type`, `id`, `description`, `serviceEndpoint`, `proof`, `encrypted`, `right` (all string). |

  `VerificationMethod` (inside `Verification.method`): `id`, `type`, `controller` (string), and a `verificationMaterial` **oneof** — exactly one of `blockchainAccountID` / `publicKeyHex` / `publicKeyMultibase` / `publicKeyBase58` (all string).
- **CLI:** `ixod tx entity create [create-entity-doc] [flags]` — `[create-entity-doc]` is the raw JSON of a `MsgCreateEntity`. The CLI overwrites `owner_address` with `--from` and regenerates `verification` from the JSON (`VerificationsJSON` → `GenerateVerificationsFromJson`). Requires `--from`, `--chain-id`, `--fees`/`--gas`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgCreateEntity',
    value: ixo.entity.v1beta1.MsgCreateEntity.fromPartial({
      entityType: 'protocol',
      ownerDid: 'did:ixo:...',
      ownerAddress: 'ixo1...',
      relayerNode: 'did:ixo:...',
      verification: [/* ixo.iid.v1beta1.Verification.fromPartial({...}) */],
    }),
  };
  ```
- **Gotchas:** ValidateBasic requires valid `owner_address`, valid `relayer_node` and `owner_did` DIDs, and a non-empty `verification` list (each validated); services and controllers validated if present. msg_server additionally: errors if `nftContractAddress` unset, if `relayer_node` DID not found, or if the generated entity DID already exists; mints the NFT via WasmKeeper.Execute (NFT mint failure rolls back the whole tx). The id is `did:ixo:entity:<md5(nftContractAddress/createSequence)>` and `createSequence` is incremented.

### MsgUpdateEntity
- **Purpose:** Update an existing entity's mutable fields. **Overwrites** `status`, `start_date`, `end_date`, `credentials` with the message values — omitted fields become Go zero values (never null), so resend existing values you want to keep.
- **Signer / auth:** `controller_address` signs. `controller_did` must be a controller on the entity's iid document with the `Authentication` relationship, and `controller_address` must be authorized to sign for it (`ExecuteOnDidWithRelationships`). The entity (DID) must exist.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID to update. |
| `entity_status` | int32 | no | New status (overwrites). |
| `start_date` | google.protobuf.Timestamp | no | stdtime (overwrites). |
| `end_date` | google.protobuf.Timestamp | no | stdtime (overwrites). |
| `credentials` | repeated string | no | New credentials list (overwrites). |
| `controller_did` | string (`iidtypes.DIDFragment`) | yes | Signer's controller DID. |
| `controller_address` | string | yes | Cosmos address signing the tx. |

- **CLI:** `ixod tx entity update [update-entity-doc] [flags]` — `[update-entity-doc]` is the raw JSON of a `MsgUpdateEntity`; CLI overwrites `controller_address` with `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgUpdateEntity',
    value: ixo.entity.v1beta1.MsgUpdateEntity.fromPartial({
      id: 'did:ixo:entity:...',
      controllerDid: 'did:ixo:...',
      controllerAddress: 'ixo1...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic requires valid `controller_address`, and valid `id` and `controller_did` DIDs. Because every listed field is overwritten, partial updates must echo prior values. Authorization fails if `controller_did` lacks the `Authentication` relationship on the entity's iid document.

### MsgUpdateEntityVerified
- **Purpose:** Set the entity's `entity_verified` flag.
- **Signer / auth:** `relayer_node_address` signs. `relayer_node_did` must exactly equal the entity's stored `relayer_node`; otherwise `ErrUpdateVerifiedFailed`. Only the relayer node may call this.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID. |
| `entity_verified` | bool | yes | New verified value. |
| `relayer_node_did` | string (`iidtypes.DIDFragment`) | yes | Relayer node DID; must match the entity's `relayer_node`. |
| `relayer_node_address` | string | yes | Cosmos address signing the tx. |

- **CLI:** `ixod tx entity update-entity-verified [id] [relayer-did] [verified] [flags]` — `[verified]` is parsed as a bool; `relayer_node_address` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgUpdateEntityVerified',
    value: ixo.entity.v1beta1.MsgUpdateEntityVerified.fromPartial({
      id: 'did:ixo:entity:...',
      entityVerified: true,
      relayerNodeDid: 'did:ixo:...',
      relayerNodeAddress: 'ixo1...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic requires valid address and valid `id` + `relayer_node_did` DIDs. msg_server rejects if `relayer_node_did` ≠ stored `relayer_node`.

### MsgTransferEntity
- **Purpose:** Transfer an entity (and its NFT) to a recipient. Rewrites the iid document's controllers to `[entityDID, recipientDid]`, removes existing verification methods matching the recipient and adds the recipient as a new `Authentication` verification method (blockchain account id), and transfers the CW721 NFT to the recipient's address.
- **Signer / auth:** `owner_address` signs. `owner_did` must be a controller with the `Authentication` relationship on the entity's iid document. `recipient_did` must resolve to an existing iid document with a blockchain-address verification method.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID to transfer. |
| `owner_did` | string (`iidtypes.DIDFragment`) | yes | Current owner's DID (signer). |
| `owner_address` | string | yes | Cosmos address signing the tx. |
| `recipient_did` | string (`iidtypes.DIDFragment`) | yes | Recipient DID; must already exist on-chain. |

- **CLI:** `ixod tx entity transfer [id] [owner-did] [recipient-did] [flags]` — `owner_address` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgTransferEntity',
    value: ixo.entity.v1beta1.MsgTransferEntity.fromPartial({
      id: 'did:ixo:entity:...',
      ownerDid: 'did:ixo:...',
      ownerAddress: 'ixo1...',
      recipientDid: 'did:ixo:...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic requires valid `owner_address` and valid `id`, `owner_did`, `recipient_did` DIDs. msg_server errors if `nftContractAddress` unset, if the recipient DID is not found, or if it lacks a blockchain-address verification method. NFT transfer is via WasmKeeper.Execute. (Note: the emitted `EntityTransferredEvent` sets `From`=recipient and `To`=owner — a known field-label quirk; ownership truly moves to the recipient.)

### MsgCreateEntityAccount
- **Purpose:** Create an additional named module account for the entity.
- **Signer / auth:** `owner_address` signs and must be the NFT owner (`CheckIfOwner`). The entity must exist; the account `name` must be unique on the entity.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID. |
| `name` | string | yes | Account name (must be non-empty and not already present). |
| `owner_address` | string | yes | Cosmos address signing the tx; must be NFT owner. |

- **CLI:** `ixod tx entity create-entity-account [id] [name] [flags]` — `owner_address` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgCreateEntityAccount',
    value: ixo.entity.v1beta1.MsgCreateEntityAccount.fromPartial({
      id: 'did:ixo:entity:...',
      name: 'savings',
      ownerAddress: 'ixo1...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic requires valid `owner_address`, valid `id` DID, non-empty `name`. msg_server: `ErrAccountDuplicate` if name exists; `unauthorized` if signer is not the NFT owner. Address is derived deterministically from entity DID + name.

### MsgGrantEntityAccountAuthz
- **Purpose:** Create an authz grant where an entity account is the *granter* and `grantee_address` is the grantee.
- **Signer / auth:** `owner_address` signs and must be the NFT owner. The named entity account must exist. No CLI — construct via SDK.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID owning the account. |
| `name` | string | yes | Entity account name to use as granter. |
| `grantee_address` | string | yes | Grantee address. |
| `grant` | `cosmos.authz.v1beta1.Grant` (non-nullable) | yes | The authz grant (`authorization` Any + `expiration`). |
| `owner_address` | string | yes | Cosmos address signing the tx; must be NFT owner. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgGrantEntityAccountAuthz',
    value: ixo.entity.v1beta1.MsgGrantEntityAccountAuthz.fromPartial({
      id: 'did:ixo:entity:...',
      name: 'admin',
      granteeAddress: 'ixo1...',
      grant: { /* cosmos.authz.v1beta1.Grant: authorization (Any), expiration */ },
      ownerAddress: 'ixo1...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic checks addresses, `id` DID, non-empty `name` (it does NOT call `Grant.ValidateBasic`). msg_server: `ErrAccountNotFound` if name missing; `unauthorized` if not NFT owner; unpacks and validates the grant, then `SaveGrant` with the account as granter.

### MsgRevokeEntityAccountAuthz
- **Purpose:** Revoke an existing authz grant where an entity account is the granter, for a given `msg_type_url`.
- **Signer / auth:** `owner_address` signs and must be the NFT owner. The named account must exist.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `id` | string | yes | Entity DID owning the account. |
| `name` | string | yes | Entity account name (granter). |
| `grantee_address` | string | yes | Grantee whose grant is revoked. |
| `msg_type_url` | string | yes | Message type URL of the grant to revoke. |
| `owner_address` | string | yes | Cosmos address signing the tx; must be NFT owner. |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.entity.v1beta1.MsgRevokeEntityAccountAuthz',
    value: ixo.entity.v1beta1.MsgRevokeEntityAccountAuthz.fromPartial({
      id: 'did:ixo:entity:...',
      name: 'admin',
      granteeAddress: 'ixo1...',
      msgTypeUrl: '/cosmos.bank.v1beta1.MsgSend',
      ownerAddress: 'ixo1...',
    }),
  };
  ```
- **Gotchas:** ValidateBasic checks addresses, `id` DID, non-empty `name` and `msg_type_url`. msg_server: `ErrAccountNotFound` if name missing; `unauthorized` if not NFT owner; `DeleteGrant` against the account as granter.

## Queries
The entity module's `GetQueryCmd` returns `nil` — there is **no `ixod query entity` CLI**. Query only via gRPC / REST.

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| Params | `Query/Params` | none | — | `QueryParamsResponse { params }` |
| Entity (with iid doc) | `Query/Entity` | none | `id` | `QueryEntityResponse { entity, iidDocument }` |
| Entity metadata | `Query/EntityMetaData` | none | `id` | `QueryEntityMetadataResponse { entity }` |
| Entity iid document | `Query/EntityIidDocument` | none | `id` | `QueryEntityIidDocumentResponse { iidDocument }` |
| Entity verified | `Query/EntityVerified` | none | `id` | `QueryEntityVerifiedResponse { entity_verified }` |
| Entity list | `Query/EntityList` | none | `pagination` | `QueryEntityListResponse { entities[], pagination }` |

REST paths: `/ixo/entity/params`, `/ixo/entity/{id}`, `/ixo/entity/{id}/metadata`, `/ixo/entity/{id}/iiddocument`, `/ixo/entity/{id}/verified`, `/ixo/entity`.

## Events
Typed events (`event.proto`), emitted via `EmitTypedEvents`:
- `EntityCreatedEvent { entity, signer }` — on `MsgCreateEntity` (alongside the iid `IidDocumentCreatedEvent`).
- `EntityUpdatedEvent { entity, signer }` — on `MsgUpdateEntity`, `MsgUpdateEntityVerified`, and `MsgCreateEntityAccount` (the entity record changed).
- `EntityVerifiedUpdatedEvent { id, signer, entity_verified }` — on `MsgUpdateEntityVerified`.
- `EntityTransferredEvent { id, from, to }` — on `MsgTransferEntity` (see field-label quirk above).
- `EntityAccountCreatedEvent { id, signer, account_name, account_address }` — on `MsgCreateEntityAccount`.
- `EntityAccountAuthzCreatedEvent { id, signer, account_name, granter, grantee, grant }` — on `MsgGrantEntityAccountAuthz`.
- `EntityAccountAuthzRevokedEvent { id, signer, account_name, granter, grantee, msg_type_url }` — on `MsgRevokeEntityAccountAuthz`.

## Module gotchas
- **Create flow:** `MsgCreateEntity` is a 3-in-1 op — it builds the iid `IidDocument` (id `did:ixo:entity:<md5(nftContractAddress/createSequence)>`), stores the `Entity`, increments `Params.createSequence`, creates the `admin` module account, and mints the CW721 NFT to `owner_address`. All in one atomic tx; NFT mint failure reverts everything. `Params.nftContractAddress` / `nftContractMinter` must be configured first.
- **Entity ≡ iid DID document:** the entity's controllers, verification, service, linkedResource, linkedClaim, accordedRight, linkedEntity, alsoKnownAs are stored on the iid document, not duplicated on `Entity`. Read them via `Query/EntityIidDocument` or the iid module. Updating those (beyond status/dates/credentials) goes through iid messages, not entity messages.
- **Authorization split:** `MsgUpdateEntity` / `MsgTransferEntity` authorize via the iid controller list + `Authentication` relationship; entity-account messages authorize via NFT ownership (`CheckIfOwner`). `entity_verified` is relayer-node-only.
- **Entity classes / "Protocol" templates:** entity `type` (e.g. `protocol`, `asset/device`, `dao`) is implementer-defined, interpreted by client apps; "Protocol" entities are conventionally used as templates/classes for other entities. The chain does not enforce type semantics.
- **Cross-module references:** claims, token and other modules commonly reference an entity DID (`did:ixo:entity:...`) as the subject/collection owner. Params (`nftContractAddress`, `nftContractMinter`) are set via gov: `ixod tx gov submit-legacy-proposal update-entity-params [nft-contract-code] [nft-minter-address] --title --description --deposit …` (the handler builds `InitializeNftContract`; it is registered as a gov legacy-proposal handler, NOT under `ixod tx entity`).
