# Epochs module — `x/epochs`

**Proto package:** `ixo.epochs.v1beta1` · **CLI:** `ixod query epochs …` (queries only — no transactions)

## Purpose
`x/epochs` defines generic on-chain timers ("epochs") that tick at fixed time intervals, so other modules can register logic to run once per period (e.g. once a day or once a week) without managing their own scheduling. A timer ticks on the first block whose blocktime exceeds the timer's end time, then sets the new start to the prior end time (not the wall-clock block time) — so a chain that was down catches up with one tick per block. Other modules subscribe via the `EpochHooks` interface; in this chain `x/mint` (inflation) and `x/liquidstake` (autocompound/rebalance) are the consumers. Nothing in this module is directly user-callable.

## Concepts & state
- **`EpochInfo`** — one record per timer identifier, holding the timer's current state; modified only in begin-blockers. Fields (verbatim from `epoch.proto`):
  - `string identifier` — unique reference to this particular timer.
  - `google.protobuf.Timestamp start_time` (non-nullable, stdtime) — time at which the timer first ever ticks; if in the future, the epoch does not begin until then.
  - `google.protobuf.Duration duration` (non-nullable, stdduration) — time between epoch ticks; must be non-zero and should exceed expected block time.
  - `int64 current_epoch` — current epoch number (how many times the timer has ticked); first tick is `current_epoch=1`.
  - `google.protobuf.Timestamp current_epoch_start_time` (non-nullable, stdtime) — start time of the current interval `(current_epoch_start_time, current_epoch_start_time + duration]`; may diverge significantly from wall-clock time.
  - `bool epoch_counting_started` — whether this timer has begun yet.
  - `int64 current_epoch_start_height` — block height at which the current epoch started (height of the last tick). *(Field 7 is `reserved`.)*
- **Default epochs** (from `DefaultGenesis`) — three timers: `day` (24h), `hour` (1h), `week` (7d), each initialized with zero values and `epoch_counting_started=false`. (The `x/epochs/README.md` example shows only `day`/`week` and is stale; the genesis code defines all three. A commented-out `2min` test-only epoch exists for local liquidstake testing.)
- **`EpochHooks` interface** — the subscription contract other modules implement (`x/epochs/types/hooks.go`):
  - `BeforeEpochStart(ctx, epochIdentifier string, epochNumber int64) error` — called as a new epoch begins (the block after epoch end).
  - `AfterEpochEnd(ctx, epochIdentifier string, epochNumber int64) error` — called when an epoch ends (first block whose timestamp is after the duration).
  - `GetModuleName() string` — name of the implementing module.
  - Hooks are combined via `MultiEpochHooks`; each runs in registration order with **panic isolation** — if one hook errors/panics its state update is reverted, but remaining hooks still proceed.

## Messages
This module has no transactions (no `Msg` service). There is no `proto/ixo/epochs/v1beta1/tx.proto` and no `x/epochs/client/cli` package. Epoch timing is configured in genesis/params and advanced automatically each block by the begin-blocker — there is no runtime tx to create, edit, or trigger an epoch.

To **add an epoch**: include an additional `EpochInfo` in the module's genesis (`GenesisState.epochs`), e.g. via `types.NewGenesisEpochInfo(identifier, duration)`. To **hook into an epoch** (Go code): implement the `EpochHooks` interface on a module keeper, filter on `epochIdentifier` (typically a value from that module's params so governance can change it), and register the hooks in `app/keepers/keepers.go` via `EpochsKeeper.SetHooks(epochstypes.NewMultiEpochHooks(...))`.

## Queries
| Query | gRPC method | CLI | Args | Returns |
| --- | --- | --- | --- | --- |
| All running epoch infos | `EpochInfos` (`QueryEpochsInfoRequest`) | gRPC / REST only — no CLI command (no `client/cli` package) | none | `QueryEpochsInfoResponse { repeated EpochInfo epochs }` |
| Current epoch number for an identifier | `CurrentEpoch` (`QueryCurrentEpochRequest`) | gRPC / REST only — no CLI command | `identifier string` | `QueryCurrentEpochResponse { int64 current_epoch }` |

REST routes (from `query.proto`): `GET /ixo/epochs/v1beta1/epochs` and `GET /ixo/epochs/v1beta1/epochs/{identifier}`. Note: the `README.md` shows `ixod query epochs epoch-infos` / `current-epoch [identifier]`, but no CLI command package exists in this module, so those commands are not wired up here — use gRPC/REST.

## Events
Emitted as typed protobuf events (`EmitTypedEvent`, from `event.proto`) in the begin-blocker:
- **`EpochStartEvent`** — `int64 epoch_number` (epoch number, starting from 1) and `google.protobuf.Timestamp start_time` (start timestamp of the epoch); emitted when a new epoch begins.
- **`EpochEndEvent`** — `int64 epoch_number`; emitted when an epoch ends.

(The `README.md` documents legacy attribute-style `epoch_start`/`epoch_end` events with `epoch_number`/`start_time` keys; the current code emits the typed events above.)

## Module gotchas
- Periodic logic in other modules fires off these hooks: `x/mint` runs inflation/minting on its configured `EpochIdentifier` (default `"day"`), and `x/liquidstake` runs autocompound/rebalance on its configured epoch identifiers. Changing an epoch's `duration` or a consumer's identifier directly affects how often that downstream logic runs.
- Epoch **identifiers are referenced by name** from other modules' params (e.g. `x/mint` `EpochIdentifier`); a consumer hook only acts when `epochIdentifier` matches its configured value, so a typo or missing genesis epoch silently disables that downstream behavior.
- Ticks are driven by blocktime, not wall-clock: after downtime the timer catches up one tick per block, which can bunch multiple epochs (and their hooks) into consecutive blocks.
- Panic isolation means a downstream hook that depends on a prior hook's state must guard for the case where the prior hook reverted.
- Nothing here is directly user-callable — there are no transactions; users only read state via the two queries above.
