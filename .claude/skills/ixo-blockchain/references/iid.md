# Iid module — `x/iid`

**Proto package:** `ixo.iid.v1beta1` · **TS typeUrl prefix:** `/ixo.iid.v1beta1.` · **CLI:** `ixod tx iid …` / `ixod query iid …`

## Purpose
The `iid` module is a W3C-DID-Core-compliant registry for Interchain Identifiers (IIDs) — decentralized identifiers and their DID documents — stored on-chain. It is the **foundational identity module**: a signer must own an `iid`/DID document before most other modules (entity, claims, token, bonds) will let that account act, because those modules reference iid DIDs as owners/controllers. It supports create / read / update / deactivate of DID documents and management of their verification methods, relationships, services, controllers, contexts, and linked resources/claims/entities/rights. It is a CRU(no-D) registry — documents are deactivated, never deleted.

## Concepts & state
- **DID / IID** — a string identifier. The canonical user-created form is `did:ixo:<bech32-account>` (prefix `IxoDidPrefix = "did:ixo:"`). A wasm-contract form `did:ixo:wasm:<bech32-contract>` is also allowed. `did:x:` (`DidChainPrefix`) is a legacy chain-DID prefix produced by `NewChainDID(chainName, didID)` → `did:x:<chain>:<id>` (used by the aries helper, not by `MsgCreateIidDocument`).
- **IidDocument** (stored, key `0x01 | id`): `context []Context`, `id string`, `controller []string`, `verificationMethod []VerificationMethod`, `service []Service`, relationship arrays `authentication/assertionMethod/keyAgreement/capabilityInvocation/capabilityDelegation []string`, `linkedResource []LinkedResource`, `linkedClaim []LinkedClaim`, `accordedRight []AccordedRight`, `linkedEntity []LinkedEntity`, `alsoKnownAs string`, `metadata IidMetadata`. Note the wire `IidDocument` field order differs (linkedResource=11, linkedClaim=12, accordedRight=13, linkedEntity=14).
- **VerificationMethod** — `id`, `type`, `controller`, plus a `oneof verificationMaterial`: `blockchainAccountID` | `publicKeyHex` | `publicKeyMultibase` | `publicKeyBase58` (all `string`). The cryptographic key bound to the DID.
- **Verification relationships** (string constants, exact values): `authentication`, `assertionMethod`, `keyAgreement`, `capabilityInvocation`, `capabilityDelegation`. These arrays list verification-method IDs authorized for that capability. The `authentication` relationship is what authorizes updates to the document.
- **Verification** (input wrapper, defined in tx.proto) — couples a `VerificationMethod` with the `relationships` it is granted plus optional `context`. Used when creating a doc or via `MsgAddVerification`.
- **Service** — `id`, `type`, `serviceEndpoint`. Ways to communicate with the DID subject (e.g. `DIDCommMessaging`).
- **Context** — `key`, `val`. JSON-LD `@context` entries (stored under jsontag `@context`).
- **Controller** — a DID string authorized to modify the document (separate from a verification-method signer). NOTE: in the iid keeper the signer is always an address, so the controller branch of the auth check effectively never matches; in practice authorization comes from the `authentication` relationship.
- **AccordedRight** — `type`, `id`, `mechanism`, `message`, `service`.
- **LinkedResource** — `type`, `id`, `description`, `mediaType`, `serviceEndpoint`, `proof`, `encrypted`, `right`.
- **LinkedClaim** — `type`, `id`, `description`, `serviceEndpoint`, `proof`, `encrypted`, `right`.
- **LinkedEntity** — `type`, `id`, `relationship`, `service`.
- **IidMetadata** — `versionId`, `created`, `updated` (timestamps), `deactivated bool`. Maintained by the keeper; not user-set.
- **Reserved DID namespaces** — `ReservedDidPrefixes` (currently `did:ixo:entity:`) are minted by other modules via `IidKeeper.SetDidDocument` and cannot be created through `MsgCreateIidDocument`.

## Messages
Source order from `tx.proto`. Every Msg has `signer` and uses `(cosmos.msg.v1.signer) = "signer"`. Shared nested types (`Verification`, `VerificationMethod`, `Service`, `Context`, `AccordedRight`, `LinkedResource`, `LinkedClaim`, `LinkedEntity`) are documented in Concepts above.

### MsgCreateIidDocument
- **Purpose:** Create a new DID document.
- **Signer / auth:** `signer` (bech32 account address). The DID `id` must be the signer's own account DID; see Gotchas.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did (must be `did:ixo:<signer-bech32>` or `did:ixo:wasm:<contract>`) |
| controllers | repeated string | no | controller DIDs |
| context | repeated Context | no | JSON-LD contexts |
| verifications | repeated Verification | yes | verification methods + relationships (must be non-empty) |
| services | repeated Service | no | services |
| accordedRight | repeated AccordedRight | no | accorded rights |
| linkedResource | repeated LinkedResource | no | linked resources |
| linkedEntity | repeated LinkedEntity | no | linked entities |
| alsoKnownAs | string | no | also-known-as URI |
| signer | string | yes | signing account address |
| linkedClaim | repeated LinkedClaim | no | linked claims |

- **CLI:** `ixod tx iid create-iid [did-doc] [flags]` — single positional arg is raw JSON of `MsgCreateIidDocument`; the CLI overrides `signer` with the `--from` address and regenerates `verifications` from the JSON. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgCreateIidDocument',
    value: ixo.iid.v1beta1.MsgCreateIidDocument.fromPartial({
      id: 'did:ixo:<signer-bech32>',
      verifications: [/* ixo.iid.v1beta1.Verification.fromPartial({...}) */],
      signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires: valid bech32 `signer`; `IsValidDID(id)`; `ValidateMsgCreateDIDForm(id, signer)`; at least one verification; each verification/service valid; each controller a valid DID. `ValidateMsgCreateDIDForm` enforces: reserved prefixes rejected (`ErrReservedDidNamespace`); must start with `did:ixo:` else `ErrDIDFormNotAllowed`; `did:ixo:wasm:<x>` requires a single valid bech32 after `wasm:`; otherwise `did:ixo:<account>` must be a single bech32 segment that **equals the signer** else `ErrDIDAccountSignerMismatch` / `ErrDIDFormNotAllowed`. The same form check re-runs in the keeper (defense in depth). Existing id → `ErrDidDocumentFound` ("a document with did %s already exists").

### MsgUpdateIidDocument
- **Purpose:** Overwrite an existing DID document. **Full replace** — every field is set; omitted fields become their Go zero value (never null), so resend values you want to keep.
- **Signer / auth:** `signer` must hold a verification method in the document's `authentication` relationship (constraint `Authentication`); else `ErrUnauthorized`.
- **Fields:** identical to `MsgCreateIidDocument` (same field names/types/numbers).
- **CLI:** `ixod tx iid update-iid [did-doc] [flags]` — positional arg is raw JSON of `MsgUpdateIidDocument`; CLI sets `signer` from `--from` and regenerates `verifications`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgUpdateIidDocument',
    value: ixo.iid.v1beta1.MsgUpdateIidDocument.fromPartial({
      id: 'did:ixo:<signer-bech32>',
      verifications: [/* ... */],
      signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires valid signer, `IsValidDID(id)`, non-empty `verifications`, valid services/controllers (does NOT run the create-form/account-match check). Document metadata is preserved across the update.

### MsgAddVerification
- **Purpose:** Add one verification method (and its relationships) to a document.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| verification | Verification | yes | the verification to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-verification-method [id] [verification] [flags]` — `[id]` is the DID, `[verification]` is raw JSON of `Verification`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddVerification',
    value: ixo.iid.v1beta1.MsgAddVerification.fromPartial({
      id: did,
      verification: ixo.iid.v1beta1.Verification.fromPartial({ /* method, relationships */ }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` → valid signer, `IsValidDID(id)`, `ValidateVerification(verification)`. Document must exist (`ErrDidDocumentNotFound`).

### MsgUpdateIidDocument (see above) — listed for completeness; tx.proto declares it right after Create.

### MsgRevokeVerification
- **Purpose:** Remove a verification method and all relationships referencing it.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| method_id | string | yes | verification method id (a DID URL) |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid revoke-verification-method [id] [method-id] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgRevokeVerification',
    value: ixo.iid.v1beta1.MsgRevokeVerification.fromPartial({
      id: did, methodId: methodId, signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires `IsValidDIDURL(method_id)` else `ErrInvalidDIDURLFormat`. (TS field is `methodId`.)

### MsgSetVerificationRelationships
- **Purpose:** Overwrite the relationship set of one verification method.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| method_id | string | yes | verification method id (DID URL) |
| relationships | repeated string | yes | relationships to set (≥1) |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid set-verification-relationship [id] [method-id] --relationship NAME [--relationship NAME ...] [flags]`. Flag `-r/--relationship` (repeatable, string slice); `--unsafe` skips auto-adding the `authentication` relationship (by default `authentication` is appended). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgSetVerificationRelationships',
    value: ixo.iid.v1beta1.MsgSetVerificationRelationships.fromPartial({
      id: did, methodId: methodId, relationships: ['authentication'], signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires `IsValidDIDURL(method_id)` and non-empty `relationships` (`ErrEmptyRelationships`). Use exact relationship strings (see Concepts). Removing `authentication` from your only key can lock you out.

### MsgAddService
- **Purpose:** Add a service to a document.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| service_data | Service | yes | the service to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-service [id] [service-id] [type] [endpoint] [flags]` (positionals map to `Service{id,type,serviceEndpoint}`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddService',
    value: ixo.iid.v1beta1.MsgAddService.fromPartial({
      id: did,
      serviceData: ixo.iid.v1beta1.Service.fromPartial({ id, type, serviceEndpoint }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` → `ValidateService(service_data)`. TS field is `serviceData`.

### MsgDeleteService
- **Purpose:** Remove a service.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| service_id | string | yes | service id (RFC3986 URI) |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-service [id] [service-id] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteService',
    value: ixo.iid.v1beta1.MsgDeleteService.fromPartial({
      id: did, serviceId: serviceId, signer: address,
    }),
  };
  ```
- **Gotchas:** `service_id` must be non-empty and a valid RFC3986 URI (`ErrInvalidRFC3986UriFormat`). Doc must have services (`ErrInvalidState` "doesn't have services associated"). TS field `serviceId`.

### MsgAddController
- **Purpose:** Add a controller DID.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did of the document |
| controller_did | string | yes | DID to add as controller |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-controller [id] [controller-did] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddController',
    value: ixo.iid.v1beta1.MsgAddController.fromPartial({
      id: did, controllerDid: controllerDid, signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` calls `IsValidIIDKeyFormat(controller_did)` which currently always returns `true` (no real check). TS field `controllerDid`.

### MsgDeleteController
- **Purpose:** Remove a controller DID.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did of the document |
| controller_did | string | yes | DID to remove from controllers |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-controller [id] [controller-did] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteController',
    value: ixo.iid.v1beta1.MsgDeleteController.fromPartial({
      id: did, controllerDid: controllerDid, signer: address,
    }),
  };
  ```
- **Gotchas:** `ValidateBasic` requires `IsValidDID(controller_did)`. TS field `controllerDid`.

### MsgAddLinkedResource
- **Purpose:** Add a linked resource.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| linkedResource | LinkedResource | yes | the resource to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-linked-resource [id] [resource-id] [type] [description] [media-type] [service-endpoint] [proof] [encrypted] [privacy] [flags]` (9 positionals; map to `LinkedResource{id,type,description,mediaType,serviceEndpoint,proof,encrypted,right}` — the `[privacy]` positional fills `right`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddLinkedResource',
    value: ixo.iid.v1beta1.MsgAddLinkedResource.fromPartial({
      id: did,
      linkedResource: ixo.iid.v1beta1.LinkedResource.fromPartial({ /* ... */ }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `linkedResource` must be non-nil (`ErrInvalidInput` "linked resource cannot be nil").

### MsgDeleteLinkedResource
- **Purpose:** Remove a linked resource.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| resource_id | string | yes | resource id |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-resource [id] [resource-id] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteLinkedResource',
    value: ixo.iid.v1beta1.MsgDeleteLinkedResource.fromPartial({
      id: did, resourceId: resourceId, signer: address,
    }),
  };
  ```
- **Gotchas:** `resource_id` non-empty. Doc must have resources (`ErrInvalidState`). TS field `resourceId`.

### MsgAddLinkedClaim
- **Purpose:** Add a linked claim.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| linkedClaim | LinkedClaim | yes | the claim to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-linked-claim [id] [claim-id] [type] [description] [service-endpoint] [proof] [encrypted] [privacy] [flags]` (8 positionals; map to `LinkedClaim{id,type,description,serviceEndpoint,proof,encrypted,right}` — `[privacy]` fills `right`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddLinkedClaim',
    value: ixo.iid.v1beta1.MsgAddLinkedClaim.fromPartial({
      id: did,
      linkedClaim: ixo.iid.v1beta1.LinkedClaim.fromPartial({ /* ... */ }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `linkedClaim` must be non-nil (`ErrInvalidInput` "linked claim cannot be nil").

### MsgDeleteLinkedClaim
- **Purpose:** Remove a linked claim.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| claim_id | string | yes | claim id |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-claim [id] [claim-id] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteLinkedClaim',
    value: ixo.iid.v1beta1.MsgDeleteLinkedClaim.fromPartial({
      id: did, claimId: claimId, signer: address,
    }),
  };
  ```
- **Gotchas:** `claim_id` non-empty. Doc must have claims (`ErrInvalidState`). TS field `claimId`.

### MsgAddLinkedEntity
- **Purpose:** Add a linked entity.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the iid |
| linkedEntity | LinkedEntity | yes | the entity to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-linked-entity [id] [entity-id] [type] [relationship] [service] [flags]` (5 positionals; map to `LinkedEntity{id,type,relationship,service}`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddLinkedEntity',
    value: ixo.iid.v1beta1.MsgAddLinkedEntity.fromPartial({
      id: did,
      linkedEntity: ixo.iid.v1beta1.LinkedEntity.fromPartial({ id, type, relationship, service }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `linkedEntity` must be non-nil (`ErrInvalidInput` "linked entity cannot be nil").

### MsgDeleteLinkedEntity
- **Purpose:** Remove a linked entity.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the iid |
| entity_id | string | yes | entity id |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-linked-entity [id] [entity-id] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteLinkedEntity',
    value: ixo.iid.v1beta1.MsgDeleteLinkedEntity.fromPartial({
      id: did, entityId: entityId, signer: address,
    }),
  };
  ```
- **Gotchas:** `entity_id` non-empty. Doc must have entities (`ErrInvalidState`). TS field `entityId`.

### MsgAddAccordedRight
- **Purpose:** Add an accorded right.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| accordedRight | AccordedRight | yes | the accorded right to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-accorded-right [id] [right-id] [type] [mechanism] [message] [service-endpoint] [flags]` (6 positionals; map to `AccordedRight{id,type,mechanism,message,service}`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddAccordedRight',
    value: ixo.iid.v1beta1.MsgAddAccordedRight.fromPartial({
      id: did,
      accordedRight: ixo.iid.v1beta1.AccordedRight.fromPartial({ /* ... */ }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `accordedRight` must be non-nil (`ErrInvalidInput` "accordede right cannot be nil").

### MsgDeleteAccordedRight
- **Purpose:** Remove an accorded right.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| right_id | string | yes | accorded right id |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-accorded-right [id] [resource-id] [flags]` (note: second positional is labelled `[resource-id]` in CLI but is the right id). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteAccordedRight',
    value: ixo.iid.v1beta1.MsgDeleteAccordedRight.fromPartial({
      id: did, rightId: rightId, signer: address,
    }),
  };
  ```
- **Gotchas:** `right_id` non-empty. Doc must have rights (`ErrInvalidState`). TS field `rightId`.

### MsgAddIidContext
- **Purpose:** Add a JSON-LD context entry.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| context | Context | yes | the context to add |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid add-iid-context [id] [key] [value] [flags]` (map to `Context{key,val}`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgAddIidContext',
    value: ixo.iid.v1beta1.MsgAddIidContext.fromPartial({
      id: did,
      context: ixo.iid.v1beta1.Context.fromPartial({ key, val }),
      signer: address,
    }),
  };
  ```
- **Gotchas:** `context` must be non-nil (`ErrInvalidInput` "context cannot be nil").

### MsgDeactivateIID
- **Purpose:** Set the document's `metadata.deactivated` flag.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| state | bool | yes | new deactivated state |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid deactivate-iid [id] [state] [flags]` — `[state]` parsed with `strconv.ParseBool` (e.g. `true`/`false`). Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeactivateIID',
    value: ixo.iid.v1beta1.MsgDeactivateIID.fromPartial({
      id: did, state: true, signer: address,
    }),
  };
  ```
- **Gotchas:** Calls `didDoc.Deactivate(msg.State)` on the document. The spec note that `state` is "currently ignored" is stale — the proto passes `state` through to `Deactivate`.

### MsgDeleteIidContext
- **Purpose:** Remove a JSON-LD context entry by key.
- **Signer / auth:** `signer` via `authentication` relationship.
- **Fields:**

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| id | string | yes | the did |
| contextKey | string | yes | the context key |
| signer | string | yes | signing address |

- **CLI:** `ixod tx iid delete-context [id] [key] [flags]`. Requires `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.iid.v1beta1.MsgDeleteIidContext',
    value: ixo.iid.v1beta1.MsgDeleteIidContext.fromPartial({
      id: did, contextKey: contextKey, signer: address,
    }),
  };
  ```
- **Gotchas:** `contextKey` non-empty. Doc must have contexts (`ErrInvalidState`). TS field `contextKey`.

> Not exposed as a Msg: the `link-aries-agent` CLI (`NewLinkAriesAgentCmd`) exists in source but is **commented out** of `GetTxCmd()` and is therefore not registered/usable. It composes `MsgAddVerification` + `MsgAddService` from an aries agent key.

## Queries

| Query | gRPC method | CLI | Args | Returns |
|-------|-------------|-----|------|---------|
| All DID documents | `Query/IidDocuments` (REST `GET /ixo/did/dids`) | `ixod query iid iids [flags]` | none (pagination flags) | `QueryIidDocumentsResponse { iidDocuments []IidDocument, pagination }` |
| Single DID document | `Query/IidDocument` (REST `GET /ixo/did/dids/{id}`) | `ixod query iid iid [id]` | `id` (the DID) | `QueryIidDocumentResponse { iidDocument IidDocument }` |

- `iid [id]` resolves via `ResolveDid`; empty id → InvalidArgument; not found → NotFound. CLI uses `PrintProto` (gogo jsonpb) so nested types render correctly.

## Events
Typed protobuf events (package `ixo.iid.v1beta1`), emitted by the keeper:
- **`IidDocumentCreatedEvent`** `{ iidDocument: IidDocument }` — emitted on successful `MsgCreateIidDocument` (and also when another module, e.g. entity, creates an iid doc). Emitted via `EmitTypedEvents`.
- **`IidDocumentUpdatedEvent`** `{ iidDocument: IidDocument }` — emitted on every other successful iid Msg (all of them update the document), via `EmitTypedEvent` inside `ExecuteOnDidWithRelationships`.

## Module gotchas
- **DID format is strict for creation.** `MsgCreateIidDocument` only accepts `did:ixo:<bech32-account>` where the account **equals the signer**, or `did:ixo:wasm:<bech32-contract>`. Everything else (`did:cosmos:…`, `did:x:…`, `did:ixo:foo:bar`, reserved `did:ixo:entity:…`) is rejected. This means a normal user's DID is deterministically `did:ixo:<their-address>` — you generally don't choose it freely.
- **You must create your DID before using other modules.** entity/claims/token/bonds expect the signer to already own an iid document (its DID). Create it first with `MsgCreateIidDocument` (with at least one `Verification` whose `relationships` include `authentication`, and a verification method bound to your account, typically via `blockchainAccountID`).
- **Authorization = `authentication` relationship, not "signer field equals owner".** All mutating Msgs run `ExecuteOnDidWithRelationships` with the `authentication` constraint: the signer's blockchain account must be referenced by a verification method in the doc's `authentication` array. The controller fallback exists but, because the iid handlers pass an address (not a DID) as `signer`, it effectively never matches — rely on `authentication`.
- **`MsgUpdateIidDocument` is a full overwrite.** Omitted fields reset to Go zero values; always resend everything you want to keep.
- **Reserved namespaces are minted by their owning module** via `IidKeeper.SetDidDocument` (e.g. entity → `did:ixo:entity:<id>`), never through `MsgCreateIidDocument`.
- **Other modules reference iid DIDs** as owners/controllers/subjects; the iid document is the on-chain identity anchor those references resolve against (`ResolveDid`).
- **CLI JSON-input commands** (`create-iid`, `update-iid`, `add-verification-method`) take raw JSON; `signer` is overwritten from `--from` and `verifications` is regenerated from the JSON via `GenerateVerificationsFromJson`.
- **`IsValidIIDKeyFormat` and `IsValidDIDDocument` controller checks are currently permissive** (`IsValidIIDKeyFormat` always returns `true`), so controller-DID validity is not strongly enforced in `ValidateBasic`.
