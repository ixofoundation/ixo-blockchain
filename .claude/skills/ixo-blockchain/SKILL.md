---
name: ixo-blockchain
description: >-
  Authoritative reference for building transactions and queries on the ixo
  Blockchain (a Cosmos SDK L1; binary `ixod`, module
  github.com/ixofoundation/ixo-blockchain/v7). Use this whenever the user wants
  to interact with ixo on-chain features — creating or updating DIDs / IID
  documents, entities / digital twins, claim collections, submitting or
  evaluating verifiable claims, minting / transferring / retiring / cancelling
  token credits, bonding curves (bonds), liquid staking, names/namespaces,
  smart-account authenticators, epochs or mint params — via either the `ixod`
  CLI or the `@ixo/impactxclient-sdk` TypeScript SDK. Trigger this skill for any
  mention of ixo, ixod, impactxclient, IID/`did:ixo:`, impact tokens/credits,
  carbon credit retirement, claim evaluation, or ixo entity domains, even when
  the user doesn't name the exact module — message names, proto field names,
  types, CLI commands, and SDK typeUrls must be copied exactly, so consult the
  reference rather than guessing.
---

# ixo Blockchain SDK skill

The ixo Blockchain is a Cosmos SDK v0.50 Layer-1 for coordinating, financing, and verifying
real-world impact. It models the physical world as **entities** (digital twins) identified by
**DIDs**, lets agents file and evaluate **verifiable claims** about those entities, settles
**payments** and **impact credits** (tokens) on verification, and supports **bonding-curve**
fundraising and **liquid staking**.

This skill exists because agents copy message names, field names, types, CLI commands, and SDK
`typeUrl`s **verbatim** into transactions — a wrong field name or denom silently fails or
moves the wrong funds. Treat the reference files as the source of truth; don't improvise field
names from memory.

## How to use this skill

1. **Always start with [references/transactions.md](references/transactions.md)** — it defines
   the conventions every module assumes: networks/chain-ids, the `uixo` denom, the DID/address
   model, `ixod` CLI flags (keys, gas, fees, broadcast), and the universal
   `@ixo/impactxclient-sdk` flow (signer → `{ typeUrl, value: …fromPartial(…) }` →
   `signAndBroadcast`). 
2. **For an end-to-end goal** ("create an entity", "retire credits", "submit and evaluate a
   claim with payment"), read [references/workflows.md](references/workflows.md) — it stitches
   the modules together in the right order with the gotchas called out.
3. **For exact message/field/CLI/SDK detail**, open the specific module reference below. Each
   lists every `Msg` type with a field table (name · type · required · description), the
   signer/auth rule, the `ixod tx`/`query` command, the TS SDK snippet, and per-message
   gotchas.

Read the file you need on demand — don't load all of them up front.

## Module map

| Module | Path | What it does | Msgs | Reference |
|---|---|---|---|---|
| **iid** | `x/iid` | W3C DID documents (`did:ixo:…`); identity, keys, verification relationships, linked resources/claims/entities. Foundation everything else builds on. | 20 | [iid.md](references/iid.md) |
| **entity** | `x/entity` | Digital twins ("entity domains") = atomic IID doc + CW721 NFT + entity record; ownership, controllers, entity accounts + authz. | 7 | [entity.md](references/entity.md) |
| **claims** | `x/claims` | Verifiable-claim lifecycle: collections, submit/evaluate/dispute claims, payments (native/CW20/CW1155), intents, performance deposits, adjudication. | 18 | [claims.md](references/claims.md) |
| **token** | `x/token` | Impact **credits** as CW1155 (`ixo1155`): create/mint/transfer, and **retire** vs **cancel** credits, pause/stop. | 8 | [token.md](references/token.md) |
| **bonds** | `x/bonds` | Token bonding curves (automated market making): create/edit bonds, buy/sell/swap, outcome payments, withdrawals. | 10 | [bonds.md](references/bonds.md) |
| **liquidstake** | `x/liquidstake` | Multi-pool liquid staking; stake `uixo` for a transferable LST, unstake, pool/validator admin. (SDK lags chain — see file.) | 10 | [liquidstake.md](references/liquidstake.md) |
| **names** | `x/names` | Human-readable namespaces/names registry (register, transfer, status). | 7 | [names.md](references/names.md) |
| **smart-account** | `x/smart-account` | Account abstraction: add/remove authenticators (signature, cosmwasm, all-of/any-of, etc.), active state. | 3 | [smart-account.md](references/smart-account.md) |
| **epochs** | `x/epochs` | On-chain periodic timers (day/hour/week) that drive mint & liquidstake hooks. Query-only, no Msgs. | 0 | [epochs.md](references/epochs.md) |
| **mint** | `x/mint` | Per-epoch inflation/deflation. No Msg service; params via gov. | 0 | [mint.md](references/mint.md) |

Counts are the number of `rpc`s in each module's `Msg` service. epochs and mint expose queries
only — their parameters change through governance, not user transactions.

## The mental model (why the modules fit together)

- **iid is the root.** A DID document (`did:ixo:<account>` or `did:ixo:wasm:<contract>`) is an
  account's on-chain identity. The DID's `<account>` suffix must equal the signer — you can
  only create a DID for yourself. Authorization to act on a DID is expressed via W3C
  *verification relationships* (`authentication`, `assertionMethod`, …).
- **An entity is a digital twin** built on iid. `MsgCreateEntity` atomically creates an IID
  document, mints a **CW721 NFT** that represents ownership, and stores an entity record. The
  entity DID is *derived* (`did:ixo:entity:<md5(nftContract/sequence)>`) and returned by the
  chain — you don't choose it. Whoever holds the NFT owns the entity. Entities can have
  **entity accounts** (module-controlled sub-accounts) that delegate via cosmos `authz`.
- **Claims hang off entities.** A **collection** ties an entity DID + a protocol DID together
  and defines the **payments** that flow on submission/evaluation/approval/rejection. Service
  agents submit claims and oracle agents evaluate them — almost always via `authz` `MsgExec`
  signed by the collection admin, not the agent directly.
- **Tokens are the credits** that payments and impact issuance move. They are **CW1155**
  tokens minted through the `ixo1155` contract. Retiring vs cancelling both burn the token
  permanently but mean different things (retire = consumed/offset, keeps supply, needs a
  jurisdiction; cancel = error correction, reduces supply). See [token.md](references/token.md).
- **bonds, liquidstake, names, smart-account** are independent capabilities layered on the
  same account/DID model.
- **epochs** ticks time; **mint** and **liquidstake** autocompound/rebalance on those ticks.

## Non-negotiable conventions (full detail in transactions.md)

- **Denom:** send amounts as integer strings in **`uixo`** (`1 ixo = 10^6 uixo`). Never floats.
- **Addresses:** account `ixo1…`, validator `ixovaloper1…`, consensus `ixovalcons1…`.
- **Chain-ids:** mainnet `ixo-5`, testnet `pandora-8`. Endpoints rotate — verify against the
  cosmos chain-registry.
- **SDK message shape (every module):**
  `{ typeUrl: '/ixo.<module>.v1beta1.Msg<Name>', value: ixo.<module>.v1beta1.Msg<Name>.fromPartial({…}) }`
  then `signingClient.signAndBroadcast(address, [msg], fee)`.
- **CLI shape:** `ixod tx <module> <action> [args…] --from <key> --chain-id <id> --gas auto
  --gas-prices 0.025uixo -y`. Several modules take a single raw-JSON document as the arg.
- **The published SDK can lag the chain proto** (confirmed for liquidstake). When a field this
  skill documents is missing from a `fromPartial` helper, the chain `tx.proto` wins.

## Verifying against source

These references were extracted directly from the repo's `proto/ixo/**/tx.proto`,
`x/<module>/client/cli/`, and keeper code. If you need to confirm or extend a detail, the
proto files (`proto/ixo/<module>/v1beta1/`) and CLI files (`x/<module>/client/cli/tx.go`) are
the ground truth; module references note where a CLI command exists vs. where you must use
gRPC/REST/SDK.
