# Wasm artefacts for interchaintest

The `.wasm` files in this directory are the actual contracts the interchaintest
suite uploads to a real ixod chain to exercise the wasm-dependent modules.
They are sourced from `ixo-multiclient-sdk/assets/contracts/` and committed
to the repo so interchaintest runs offline (no GitHub release downloads, no
flaky network in CI).

## Inventory

| File | Source path in SDK | Used by | Code ID (devnet/testnet/mainnet) |
|---|---|---|---|
| `cw721.wasm` | `assets/contracts/ixo/cw721.wasm` | x/entity NFT contract | 1 / 1 / 1 |
| `ixo1155.wasm` | `assets/contracts/ixo/ixo1155.wasm` | x/token mint/transfer/retire/cancel/transferCredit | 2 / 2 / 2 |
| `cw20_base.wasm` | `assets/contracts/cosmwasm/cw20_base.wasm` | x/claims cw20 payment settlement | 25 / 25 / 25 |
| `cw721_base.wasm` | `assets/contracts/cosmwasm/cw721_base.wasm` | upstream cw721 base reference | 26 / 26 / 26 |
| `cw4_group.wasm` | `assets/contracts/cosmwasm/cw4_group.wasm` | reserved for group/auth tests | 24 / 24 / 24 |

Code IDs match the `contracts` array in
`ixo-multiclient-sdk/src/custom_queries/contract.constants.ts`. The IDs are
assigned in upload order during devnet bootstrap (`Proposals.instantiateModulesProposals`),
so test fixtures that hard-code code IDs against a freshly-bootstrapped
chain can use those numbers directly.

## Verifying

```bash
cd tests/interchaintest/contracts
./bootstrap.sh
```

`bootstrap.sh` runs a `shasum -a 256 -c checksums.txt` to confirm every
artefact is intact. CI runs it before the interchaintest E2E job so a
silent corruption can't slip through.

## Updating a contract

1. Drop the new `.wasm` into this directory (overwriting the old file).
2. Regenerate the checksum file:
   ```bash
   shasum -a 256 *.wasm > checksums.txt
   ```
3. Commit both the new `.wasm` and the updated `checksums.txt`.
4. Update the inventory table above if the source path or code-ID
   assignment changes.

## Why bundled and not downloaded

- **Reproducibility:** the contract bytecode that x/entity depends on is
  schema-locked. The CW20 / CW721 base in `ixo/` is not the upstream
  cw-plus binary — it is the version the ixo entity module's Wasm
  bindings expect.
- **Offline CI:** the `interchaintest-build` job runs on every PR; pulling
  300KB binaries from GitHub releases would fail half the time.
- **Single source of truth:** SDK and chain repo now resolve the same
  bytes. Bumping in one place + a sync commit is the explicit upgrade path.
