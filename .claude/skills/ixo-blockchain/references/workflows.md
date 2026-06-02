# End-to-end workflows (recipes)

Goal-oriented sequences that wire the modules together in the right order. Each recipe shows
the happy path with the key fields; open the linked module reference for the full field
tables, every optional field, and the error/gotcha list. Conventions (denom `uixo`, gas/fees,
DID rules, SDK `{ typeUrl, value }` shape) are in [transactions.md](transactions.md).

CLI snippets assume you've exported common flags:
```bash
TX="--from mykey --chain-id pandora-8 --gas auto --gas-adjustment 1.3 --gas-prices 0.015uixo -y -o json"
```

---

## 0. Prerequisites: key + funds

```bash
ixod keys add mykey                       # or --recover to import a mnemonic
ADDR=$(ixod keys show mykey -a)           # ixo1…
ixod query bank balances "$ADDR" -o json  # confirm you hold uixo for fees
```
Your account DID is simply `did:ixo:$ADDR`. You do not "register" it separately — see step 1.

---

## 1. Create a DID document (identity foundation)

Every higher-level object (entity, claim agent, etc.) is anchored to a DID. The DID's account
suffix **must equal your signer address**.

```bash
DID="did:ixo:$ADDR"
ixod tx iid create-iid "$(cat <<JSON
{ "id": "$DID", "controllers": ["$DID"] }
JSON
)" $TX
```
The CLI sets `signer` from `--from` and regenerates `verifications` from the JSON. Full field
list (verification methods, services, linked resources/claims/entities, accorded rights,
contexts) and all 20 messages: [iid.md](iid.md).

TS SDK:
```ts
const msg = {
  typeUrl: "/ixo.iid.v1beta1.MsgCreateIidDocument",
  value: ixo.iid.v1beta1.MsgCreateIidDocument.fromPartial({
    id: `did:ixo:${address}`,
    signer: address,
    controllers: [`did:ixo:${address}`],
    verifications: [/* VerificationMethod relationships */],
  }),
};
await signingClient.signAndBroadcast(address, [msg], fee);
```

Query it back (no entity/claims/token query CLI, but iid has one):
```bash
ixod query iid iid "$DID" -o json          # single DID;  `ixod query iid iids` lists all
```

---

## 2. Create an entity domain (digital twin)

`MsgCreateEntity` is **atomic**: it creates an IID document, mints a **CW721 NFT** (ownership),
and stores the entity record. **The entity DID is derived and returned by the chain** —
`did:ixo:entity:<md5(nftContract/sequence)>` — you do not choose it.

Prereq: the chain's entity params (`nftContractAddress`, `nftContractMinter`) must be set
(done once via governance: `update-entity-params` legacy proposal — see
[entity.md](entity.md)). On a configured network you just submit:

```bash
ixod tx entity create "$(cat <<JSON
{ "entity_type": "protocol", "entity_status": 0, "owner_did": "$DID" }
JSON
)" $TX
```
The CLI overwrites `owner_address` with `--from` and regenerates `verification` from the JSON.
The response/events carry the new entity DID — capture it for later steps. Full field table
(context, verifications, services, accorded rights, linked resources/entities/claims,
credentials, dates, relayer node) and the other 6 messages (update, verify, transfer, entity
accounts + authz): [entity.md](entity.md).

```ts
const msg = {
  typeUrl: "/ixo.entity.v1beta1.MsgCreateEntity",
  value: ixo.entity.v1beta1.MsgCreateEntity.fromPartial({
    entityType: "protocol",
    ownerDid: `did:ixo:${address}`,
    ownerAddress: address,
  }),
};
```

**Entity accounts + delegation:** to let an agent transact on the entity's behalf, create a
named entity account and grant it `authz`:
```bash
ixod tx entity create-entity-account "$ENTITY_DID" "agent1" $TX
# then MsgGrantEntityAccountAuthz (see entity.md) to grant a specific Msg type to a grantee
```

Entity has **no query CLI** — query via gRPC/REST/SDK (`queryClient.ixo.entity.v1beta1.entity({ id })`).

---

## 3. Issue & manage impact credits (tokens)

Impact credits are **CW1155** tokens minted through the `ixo1155` contract. Flow:
create a token class → mint batches → transfer → retire (or cancel). Full field tables and the
retire-vs-cancel semantics: [token.md](token.md).

```bash
# (a) Create the token class. `class` is the protocol/token entity DID; `name` is the unique namespace.
ixod tx token create "$(cat <<JSON
{ "class": "$ENTITY_DID", "name": "CARBON", "description": "Verified carbon credit",
  "image": "https://…", "token_type": "ixo1155", "cap": "0" }
JSON
)" $TX
# cap "0" = unlimited. This deploys/links the cw1155 contract; capture contract_address from events.

# (b) Mint a batch to an owner.
ixod tx token mint "$(cat <<JSON
{ "contract_address": "$CONTRACT", "owner": "$ADDR",
  "mint_batch": [ { "name": "CARBON", "index": "1", "amount": "1000",
                    "collection": "$COLLECTION_DID", "token_data": [] } ] }
JSON
)" $TX

# (c) Transfer credits (all tokens must be in the same contract).
ixod tx token transfer "$(cat <<JSON
{ "recipient": "$RECIPIENT", "tokens": [ { "id": "<tokenId>", "amount": "100" } ] }
JSON
)" $TX

# (d) RETIRE credits — permanent burn, keeps Supply, REQUIRES a jurisdiction. This is the
#     "offset / consume the credit" operation (e.g. carbon retirement).
ixod tx token retire-token "$(cat <<JSON
{ "tokens": [ { "id": "<tokenId>", "amount": "100" } ],
  "jurisdiction": "US-CA 94101", "reason": "offset 2026 Q2 emissions" }
JSON
)" $TX
```
**Retire vs Cancel:** both burn permanently on the cw1155 contract.
*Retire* (`retire-token`) records `TokensRetired`, **requires `jurisdiction`**, leaves
`Token.supply` unchanged — the credit was legitimately consumed/offset.
*Cancel* (`cancel-token`) records `TokensCancelled`, takes **no jurisdiction**, and does
`supply.Sub(amount)` — for error correction. (Note: in the current code `cancel-token` emits a
`TokenRetiredEvent`; rely on state, not the event name.)

`minter`/`owner` fields in the JSON are overwritten from `--from`. Token has **no query CLI**.

```ts
const msg = {
  typeUrl: "/ixo.token.v1beta1.MsgRetireToken",
  value: ixo.token.v1beta1.MsgRetireToken.fromPartial({
    owner: address,
    tokens: [{ id: tokenId, amount: "100" }],
    jurisdiction: "US-CA 94101",
    reason: "offset 2026 Q2 emissions",
  }),
};
```

---

## 4. Run a claim collection (submit + evaluate with payment)

This is the core impact-verification loop. A **collection** ties an entity DID + a protocol
DID and defines the payments that flow on submission/evaluation/approval/rejection. Service
agents submit claims; oracle agents evaluate them — **almost always through cosmos `authz`
`MsgExec` signed by the collection admin**, not the agent's own key. Full field tables (18
messages, 4 authz types, intents, disputes, performance deposits): [claims.md](claims.md).

```bash
# (a) Create the collection. signer must be the entity's NFT owner; the persisted admin is the
#     entity's admin account. `payments` defines the four payout slots (use uixo amounts).
ixod tx claims create-collection "$(cat <<JSON
{ "entity": "$ENTITY_DID", "signer": "$ADDR", "protocol": "$PROTOCOL_DID",
  "state": 0, "intents": 0,
  "payments": {
    "submission": { "account": "$ENTITY_ACCOUNT", "amount": [{"denom":"uixo","amount":"1000"}] },
    "evaluation": { "account": "$ENTITY_ACCOUNT", "amount": [{"denom":"uixo","amount":"1000"}] },
    "approval":   { "account": "$ENTITY_ACCOUNT", "amount": [{"denom":"uixo","amount":"5000"}] },
    "rejection":  { "account": "$ENTITY_ACCOUNT", "amount": [] }
  } }
JSON
)" $TX
# capture the new collection_id (from Params.collection_sequence) from events.
```

```bash
# (b) Authorize agents. Grant the service agent a SubmitClaimAuthorization and the evaluator an
#     EvaluateClaimAuthorization (cosmos authz). Done via the standard authz grant carrying the
#     claims authorization type — see "Authz authorizations" in claims.md for the exact grant
#     shape and constraints (collection_id, agent quotas, max amounts).

# (c) Submit a claim — admin_address must equal collection.Admin; normally wrapped in MsgExec so
#     the admin signs while the service agent is the grantee.
ixod tx claims submit-claim "$(cat <<JSON
{ "collection_id": "$COLLECTION_ID", "claim_id": "<cidHash>",
  "agent_did": "$AGENT_DID", "agent_address": "$AGENT_ADDR", "admin_address": "$ADMIN_ADDR" }
JSON
)" $TX

# (d) Evaluate the claim — status 1 = APPROVED fires the approval payment to the submitter.
ixod tx claims evaluate-claim "$(cat <<JSON
{ "collection_id": "$COLLECTION_ID", "claim_id": "<cidHash>", "oracle": "$ORACLE_DID",
  "agent_did": "$EVAL_DID", "agent_address": "$EVAL_ADDR", "admin_address": "$ADMIN_ADDR",
  "status": 1, "verification_proof": "<cid>" }
JSON
)" $TX
```
EvaluationStatus: `PENDING=0, APPROVED=1, REJECTED=2, INVALIDATED=4, FLAGGED=5` (DISPUTED is
deprecated for new txs). If a payment has `timeout_ns > 0`, payout is deferred: the chain
grants a `WithdrawPaymentAuthorization` and the payee later runs
`ixod tx claims withdraw-payment [doc]`.

TS SDK (wrapping in authz `MsgExec`):
```ts
const inner = {
  typeUrl: "/ixo.claims.v1beta1.MsgSubmitClaim",
  value: ixo.claims.v1beta1.MsgSubmitClaim.fromPartial({
    collectionId, claimId, agentDid, agentAddress, adminAddress,
  }),
};
const exec = {
  typeUrl: "/cosmos.authz.v1beta1.MsgExec",
  value: { grantee: agentAddress, msgs: [/* Any-encoded inner */] },
};
await signingClient.signAndBroadcast(agentAddress, [exec], fee);
```

Claims has **no query CLI** — query collections/claims/disputes via gRPC/REST/SDK.

---

## 5. Bonding-curve fundraise (bonds)

A bond is an automated market maker: a continuous-token curve over reserve tokens. Create a
bond, then buyers `buy`/`sell`/`swap`; the project can take an `outcome-payment` and holders
`withdraw-share`. Bonds are addressed by **bond DID** and most args are DIDs. Full flag list
and all 10 messages: [bonds.md](bonds.md).

```bash
# create-bond uses FLAGS (no positional args). Minimal-ish:
ixod tx bonds create-bond \
  --token mybond --name "My Bond" --description "…" \
  --function-type augmented_function \
  --function-parameters "d0:500000,p0:0.01,theta:0.4,kappa:3.0" \
  --reserve-tokens uixo \
  --tx-fee-percentage 0.5 --exit-fee-percentage 0.1 \
  --fee-address "$ADDR" --reserve-withdrawal-address "$ADDR" \
  --max-supply 1000000mybond \
  --order-quantity-limits "" --sanity-rate "0" --sanity-margin-percentage "0" \
  --batch-blocks 1 \
  --bond-did "$BOND_DID" --creator-did "$DID" --controller-did "$DID" \
  --oracle-did "$DID" $TX

# buy: [bond-token-with-amount] [max-prices] [bond-did] [buyer-did]
ixod tx bonds buy 10mybond 1000uixo "$BOND_DID" "$DID" $TX
# sell: [bond-token-with-amount] [bond-did] [seller-did]
ixod tx bonds sell 5mybond "$BOND_DID" "$DID" $TX
```
Bonds **does** have a query CLI: `ixod query bonds bond [bond-did]`,
`ixod query bonds current-price [bond-did]`, `buy-price`, `sell-return`, etc.

---

## 6. Liquid staking (stake / unstake)

Stake `uixo` into a pool to mint a transferable LST (e.g. `uzero`); unstake burns it and
unbonds over the standard period. **`liquid-stake` is restricted to the pool's
`whitelist_admin_address`**; `liquid-unstake` is open to any LST holder. Pools are created by
governance. Full detail + the SDK-lag warning: [liquidstake.md](liquidstake.md).

```bash
ixod tx liquidstake liquid-stake zero 1000000uixo $TX      # [pool-id] [amount]
ixod tx liquidstake liquid-unstake zero 1000000uzero $TX   # burn LST, begin unbonding
```
There is **no query CLI** for liquidstake. **SDK caveat:** `@ixo/impactxclient-sdk@2.4.10`
predates v7 multi-pool and omits `poolId` — if your installed SDK is stale, build the `Any`
manually with the v7 `typeUrl` and fields. See [liquidstake.md](liquidstake.md).

---

## Pointers for the rest

- **Names / namespaces** (register a human-readable name, transfer, set status):
  [names.md](names.md) — 7 messages, query CLI via autocli, no tx CLI.
- **Smart accounts** (add/remove authenticators for account abstraction):
  [smart-account.md](smart-account.md) — `ixod tx smartaccount …`, 3 messages, 8 authenticator
  types.
- **Epochs** (read the day/hour/week timers that drive mint & liquidstake):
  [epochs.md](epochs.md) — query only.
- **Mint** (inflation params, per-epoch): [mint.md](mint.md) — query only; params via gov.
