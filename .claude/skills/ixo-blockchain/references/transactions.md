# Transactions, signing & querying (cross-cutting)

How to connect, sign, broadcast, and query on ixo — for both the `ixod` CLI and the
`@ixo/impactxclient-sdk` TypeScript SDK. Read this first; every module reference assumes
the conventions below.

## Chain identity & networks

`ixod` is the binary; the on-chain module path is `github.com/ixofoundation/ixo-blockchain/v7`.
Public endpoints rotate — the **authoritative source** is the cosmos chain-registry
(`github.com/cosmos/chain-registry/ixo` for mainnet, `.../testnets/ixopandora` for testnet).
Verify before relying on a hard-coded URL.

| | Mainnet | Testnet (Pandora) | Local dev |
|---|---|---|---|
| chain-id | `ixo-5` | `pandora-8` | `devnet-1` (from `make run`) |
| RPC (CometBFT) | `https://impacthub.ixo.world/rpc/` | `https://rpc.testnet.ixo.earth/` | `http://localhost:26657` |
| REST (LCD) | `https://impacthub.ixo.world/rest/` | `https://rest.testnet.ixo.earth/` | `http://localhost:1317` |
| min gas price | `0.025uixo` | `0.015uixo` | `0uixo` |
| bech32 account | `ixo1…` | `ixo1…` | `ixo1…` |

Bech32 prefixes (same on every network, from `app/params/config.go`): account `ixo`,
validator operator `ixovaloper`, consensus `ixovalcons` (+ `pub` variants).

## Denominations

- **`uixo`** is the native staking/fee denom (the "micro" unit). `1 ixo = 1_000_000 uixo`
  (10^6). Display denom is `ixo`; `mixo` (milli) also registered. **Always send amounts to
  chain/SDK in `uixo`** as integer strings — never floats, never `ixo`.
- Demo/other denoms seen in tests and bonds: `res`, `rez`, `uxgbp` — these are not special,
  just example bank denoms.
- Liquid-staking derivative denoms (LSTs) are per-pool, e.g. `uzero` — not fungible across
  pools. See [liquidstake.md](liquidstake.md).

## DID & address model (read this — it governs entity/iid/claims/token)

Most ixo modules are keyed by **DIDs**, not addresses. The forms you will encounter:

| DID form | Meaning | Who/what controls it |
|---|---|---|
| `did:ixo:<bech32-account>` | a user/account DID | the account whose address is `<bech32-account>` |
| `did:ixo:wasm:<contract>` | a CosmWasm contract DID | the contract |
| `did:ixo:entity:<md5hex>` | an **entity** (digital twin) DID | derived, NFT-owned (see below) |

Hard rules enforced on-chain (`x/iid`):
- For `MsgCreateIidDocument`, only `did:ixo:<account>` and `did:ixo:wasm:<contract>` forms are
  allowed, and for the account form **the `<account>` suffix MUST equal the signer address**
  (`ErrDIDAccountSignerMismatch`). You cannot create a DID document for an account you don't
  control. There is no separate "register key" step — the DID *is* your account.
- **Entity DIDs are derived, not chosen:** `did:ixo:entity:%x` where `%x` is
  `md5( "<nftContractAddress>/<createSequence>" )` (`x/entity/keeper/msg_server.go`). You do
  not pass the entity DID into `MsgCreateEntity`; the chain computes it and returns it.
- Legacy forms still readable on older docs: `did:x:…` and the old demo
  `did:earth:pandora-4:…`. Don't create new ones.

Authorization on a DID document flows through the W3C **verification relationships**
(`authentication`, `assertionMethod`, `capabilityInvocation`, …). To let another key act for a
DID, add it under the appropriate relationship — see [iid.md](iid.md). For *account-level*
delegation between addresses (e.g. an agent submitting claims for an entity), ixo uses cosmos
**`authz`** grants wrapped in `MsgExec` — see [claims.md](claims.md) and
[entity.md](entity.md) (`MsgGrantEntityAccountAuthz`).

## ixod CLI

### Key management
```bash
ixod keys add <name>                      # create a new key (writes to default keyring)
ixod keys add <name> --recover            # import from mnemonic
ixod keys list
ixod keys show <name> -a                  # print the ixo1… address
# Keyring backend: --keyring-backend {os|file|test|kwallet|pass}. CI/scripts use `test`
# (UNENCRYPTED, on disk) — never use `test` for funds you care about.
```

### Anatomy of a tx command
Every `ixod tx <module> <action> [args…]` accepts the standard cosmos tx flags:
```bash
ixod tx <module> <action> [positional-args…] \
  --from <key-name-or-address> \
  --chain-id <ixo-5|pandora-8|…> \
  --gas auto --gas-adjustment 1.3 \         # simulate then pad; or --gas <int>
  --fees 5000uixo \                         # OR --gas-prices 0.025uixo (with --gas auto)
  --node <rpc-url> \                        # defaults to tcp://localhost:26657
  --keyring-backend <backend> \
  --broadcast-mode sync \                   # sync (default) | async | block(removed in sdk50)
  -y                                        # skip confirmation
```
- **Fees:** either fix them (`--fees 5000uixo`) or let the node price gas
  (`--gas auto --gas-prices 0.025uixo`). On mainnet the min gas price is `0.025uixo`; bump
  `--gas-adjustment` if you hit "out of gas".
- **Inspect without sending:** add `--generate-only` to emit unsigned JSON, or
  `--dry-run` to just simulate gas.
- Many ixo messages take a **single raw-JSON document** as their positional arg (e.g.
  `ixod tx claims create-collection [create-collection-doc]`,
  `ixod tx entity create-entity [entity-doc]`). The JSON matches the proto message field
  names. Module references show the exact arg name per command.

### Querying (CLI)
```bash
ixod query <module> <query> [args…] --node <rpc-url> --output json
ixod query <module> --help                # list available queries for a module
ixod query tx <txhash>                    # look up a broadcast result
ixod query bank balances <ixo1…>
```
Note several ixo modules expose **no query CLI** (entity, claims, liquidstake) — query those
over gRPC/REST instead (see each module reference). Many modules wire queries via cosmos
**autocli**, so `ixod query <module> --help` is the ground truth for what exists.

## TypeScript SDK — `@ixo/impactxclient-sdk`

```bash
npm install @ixo/impactxclient-sdk
```

### 1. Build a signer
Any cosmjs `OfflineSigner` with the `ixo` bech32 prefix works. The canonical form:
```ts
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, { prefix: "ixo" });
const [{ address }] = await wallet.getAccounts();   // ixo1…
```
The SDK also bundles helpers under `utils` (`utils.mnemonic`, `utils.address`, `utils.did`,
`utils.conversions`) for generating mnemonics, deriving DIDs, and unit conversions. For
browser/mobile, the `@ixo/signx-sdk` (SignX) provides a QR/relayer signing flow as an
alternative to holding a local mnemonic.

### 2. Signing client (to send txs)
```ts
import { createSigningClient } from "@ixo/impactxclient-sdk";
const signingClient = await createSigningClient(RPC_URL, wallet);
// Alternative explicit form:
// import { getSigningixoClient } from "@ixo/impactxclient-sdk";
// const signingClient = await getSigningixoClient({ rpcEndpoint: RPC_URL, signer: wallet });
```

### 3. Compose a message
Every ixo message is `{ typeUrl, value }`, where `value` is built with the generated
`fromPartial` helper. This is the universal pattern across all 10 modules:
```ts
import { ixo } from "@ixo/impactxclient-sdk";

const msg = {
  typeUrl: "/ixo.iid.v1beta1.MsgCreateIidDocument",          // proto path, leading slash
  value: ixo.iid.v1beta1.MsgCreateIidDocument.fromPartial({  // namespaced builder
    id: `did:ixo:${address}`,
    signer: address,
    controllers: [`did:ixo:${address}`],
    verifications: [ /* … */ ],
  }),
};
```
The `typeUrl` and the `ixo.<module>.v1beta1.Msg<Name>` namespace always mirror the proto:
package `ixo.token.v1beta1` → typeUrl `/ixo.token.v1beta1.MsgCreateToken` → builder
`ixo.token.v1beta1.MsgCreateToken.fromPartial`. Each module reference lists the exact
`typeUrl` and required fields.

### 4. Fee + broadcast
```ts
import { StdFee } from "@cosmjs/stargate";

const fee: StdFee = {
  amount: [{ denom: "uixo", amount: "5000" }],   // total fee, in uixo
  gas: "300000",                                  // gas limit
};
const res = await signingClient.signAndBroadcast(address, [msg], fee);
// res.code === 0  => success; res.transactionHash; res.rawLog
```
You can batch multiple messages atomically by passing `[msg1, msg2, …]`. `authz`-wrapped
flows put a `MsgExec` (`/cosmos.authz.v1beta1.MsgExec`) in this array whose `msgs` contain the
inner ixo message (claims/entity-account delegation).

### 5. Query client (read-only)
```ts
import { createQueryClient } from "@ixo/impactxclient-sdk";
const queryClient = await createQueryClient(RPC_ENDPOINT);

const entities = await queryClient.ixo.entity.v1beta1.entityList();
const iidDoc   = await queryClient.ixo.iid.v1beta1.iidDocument({ id: did });
const balance  = await queryClient.cosmos.bank.v1beta1.balance({ address, denom: "uixo" });
```
The query namespace mirrors the gRPC service: `queryClient.ixo.<module>.v1beta1.<rpcName>(req)`.

### SDK version caveat
The published SDK can lag the chain proto. The known live example: `x/liquidstake` was
reworked for multi-pool in v7, but `@ixo/impactxclient-sdk@2.4.10` still ships the pre-v7
single-pool messages (no `poolId`, stale `MsgUpdateParams`). When a `fromPartial` helper is
missing a field this doc says exists, treat **the chain `tx.proto` as the source of truth**
and construct the `Any` manually (correct `typeUrl` + proto-encoded value). Details in
[liquidstake.md](liquidstake.md).

## Quick gotchas
- Amounts are integer strings in the micro denom (`uixo`), never decimals.
- A DID you create for an account must match your signer address exactly.
- Entity DIDs are returned by the chain, not chosen by you.
- "No query CLI" ≠ "no query" — use gRPC/REST/SDK for entity, claims, liquidstake.
- Gas: prefer `--gas auto --gas-prices 0.025uixo --gas-adjustment 1.3` on mainnet.
