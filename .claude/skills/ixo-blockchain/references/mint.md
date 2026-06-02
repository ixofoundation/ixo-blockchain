# Mint module — `x/mint`

**Proto package:** `ixo.mint.v1beta1` · **CLI:** `ixod query mint …` (queries only; no transactions — there is no Msg service)

## Purpose
IXO's `x/mint` is a custom, epoch-driven minter that replaces the standard cosmos-sdk mint. Instead of computing a per-block inflation rate from a bonded-ratio target, it mints a fixed amount of `mint_denom` once per epoch (default `day`, signalled by `x/epochs`) and reduces that amount geometrically over time. Every `reduction_period_in_epochs` (default 365 epochs ≈ 1 year) the per-epoch provision is multiplied by `reduction_factor` (default `0.666…`), a generalized Bitcoin-style halvening that gives the token a finite total supply. Each epoch's minted coins are split by `distribution_proportions` between staking rewards (via the fee collector), impact-rewards receivers, and the community pool.

## Concepts & state
- **Minter** (`MinterKey = 0x00`): single record holding `epoch_provisions` (`LegacyDec`) — the reward amount minted for the current epoch. Initialized to `0` (`InitialMinter`); seeded from `genesis_epoch_provisions` is not automatic — at the start epoch the minter is set up by genesis/params.
- **LastReductionEpoch** (`LastReductionEpochKey = 0x03`): the epoch number at which the last reduction occurred; used to decide when the next reduction is due. Stored as a big-endian uint64.
- **Params**: held in the legacy `x/params` subspace named `mint` (not a module-local collection). Govern denom, the genesis provision, epoch cadence, reduction schedule, and distribution split.
- **Reduction recalculation** (`Minter.NextEpochProvisions`): `next = epoch_provisions * reduction_factor`. Applied in `AfterEpochEnd` once `epochNumber >= reduction_period_in_epochs + last_reduction_epoch`, then `last_reduction_epoch` is updated.
- **Per-epoch provision** (`Minter.EpochProvision`): `sdk.NewCoin(mint_denom, epoch_provisions.TruncateInt())` — the integer coin actually minted each epoch.
- **No bonded-ratio / annual-provisions logic**: there is no `Inflation`, `AnnualProvisions`, `goal_bonded`, `blocks_per_year`, or `max/min inflation` field. The rate is set purely by the provision amount and reduction schedule.

### `Params` fields (verbatim from `mint.proto`)
- `string mint_denom` — denom of the coin to mint.
- `genesis_epoch_provisions` (`cosmossdk.io/math.LegacyDec`, non-nullable) — epoch provisions from the first epoch.
- `string epoch_identifier` — mint epoch identifier, e.g. `day`, `week` (must be a valid `x/epochs` identifier).
- `int64 reduction_period_in_epochs` — number of epochs between reward reductions.
- `reduction_factor` (`cosmossdk.io/math.LegacyDec`, non-nullable) — multiplier applied at the end of each reduction period.
- `DistributionProportions distribution_proportions` (non-nullable) — how minted denom is split.
- `repeated WeightedAddress weighted_impact_rewards_receivers` (non-nullable) — impact-reward receivers; each gets `epoch_provisions * distribution_proportions.impact_rewards * weight`.
- `int64 minting_rewards_distribution_start_epoch` — first epoch at which minting/distribution begins.

**`DistributionProportions`** (all `LegacyDec`, non-nullable; must sum to `1`): `staking` (field 1), `impact_rewards` (field 3), `community_pool` (field 4). (Note: proto field number 2 is unused/reserved.)

**`WeightedAddress`**: `string address` (field 1, may be empty `""` → funds community pool), `weight` (field 2, `LegacyDec`, non-nullable). Weights across receivers must sum to `1`.

## Messages
No user transactions. There is **no `tx.proto` and no `Msg` service** — `RegisterServices` only registers the query server, and `codec.go` registers no messages. Inflation runs automatically on each epoch. Params are not changed via a `MsgUpdateParams`; they live in the legacy `x/params` subspace (`paramsKeeper.Subspace("mint")`) and are updated through governance using the legacy `ParameterChangeProposal` (x/gov param-change), not a module message.

## Queries
| Query | gRPC method | CLI | Args | Returns |
| --- | --- | --- | --- | --- |
| Params | `ixo.mint.v1beta1.Query/Params` | `ixod query mint params` | none | `Params` (full mint param set) |
| EpochProvisions | `ixo.mint.v1beta1.Query/EpochProvisions` | `ixod query mint epoch-provisions` | none | `epoch_provisions` (`LegacyDec`) — current per-epoch minting amount |

REST: `GET /ixo/mint/v1beta1/params`, `GET /ixo/mint/v1beta1/epoch_provisions`. Both CLI commands are wired via `autocli.go`. There is no `Inflation` or `AnnualProvisions` query (those standard-mint queries do not exist here).

## Events
- `MintEpochProvisionsMintedEvent` (from `event.proto`, emitted via `EmitTypedEvent` in `keeper/hooks.go` `AfterEpochEnd` after a successful mint+distribute) — attributes: `epoch_number`, `epoch_provisions`, `amount`. The legacy README's `mint` typed-event table (Type `mint`, keys `epoch_number` / `epoch_provisions` / `amount`) describes this same event.

## Module gotchas
- **Per-epoch, not per-block.** There is no `abci.go` and no `BeginBlocker`/`EndBlocker` minting. Minting is driven entirely by the `x/epochs` hook `AfterEpochEnd`; `BeforeEpochStart` is a no-op. Mint's `Hooks()` is registered into the epochs keeper (`app/keepers/keepers.go`), and mint hooks must be set before epochs hooks because epochs calls into them.
- **Gated start.** Nothing is minted while `epochNumber < minting_rewards_distribution_start_epoch`; at that exact epoch `last_reduction_epoch` is seeded. Minting only fires when the incoming `epochIdentifier` equals `params.epoch_identifier`.
- **Distribution path** (`DistributeMintedCoin`): coins are minted to the `mint` module account, then `staking` proportion is sent to the `auth` fee collector (`feeCollectorName`, picked up by distribution on the next begin-block), `impact_rewards` is split to `weighted_impact_rewards_receivers` (empty list or empty address → community pool), and the remainder goes to the community pool. A `mint` hook `AfterDistributeMintedCoin` fires at the end.
- **Default proportions** in code: `staking = 1.0`, `impact_rewards = 0.0`, `community_pool = 0.0` (the README's 0.9/0.1 and 45/25/5 tables are illustrative, not the in-code defaults).
- **Not user-callable.** Validators/users cannot trigger or change minting directly; only governance param changes affect the schedule.
