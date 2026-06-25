# Bonds module — `x/bonds`

**Proto package:** `ixo.bonds.v1beta1` · **TS typeUrl prefix:** `/ixo.bonds.v1beta1.` · **CLI:** `ixod tx bonds …` / `ixod query bonds …`

## Purpose
The bonds module implements token bonding curves for automated market-making and continuous capital formation. Each bond declares its own token denomination whose price is a deterministic function of supply (or, for swappers, of the reserve ratio). Buys mint bond tokens against reserve tokens; sells burn them back for reserve. Bonds support alpha/outcome-payment models (augmented bonding curves) that let funders pay an outcome payment to settle the bond, after which holders withdraw a proportional share of the reserve.

## Concepts & state
- **Bond** — a bonding-curve instance (`bonds.proto` `Bond`). Keyed by `bond_did`; mints denom `token`. Holds `function_type`, `function_parameters`, `reserve_tokens`, fee percentages, `max_supply`, `current_supply`, `current_reserve`, `available_reserve`, `current_outcome_payment_reserve`, `state`, `creator_did`/`controller_did`/`oracle_did`.
- **Function types** (string constants, `x/bonds/types/bonds.go`): `power_function`, `sigmoid_function`, `swapper_function`, `augmented_function`, `bonding_function`. The type determines pricing math and required `function_parameters`.
- **FunctionParam** (`bonds.proto`) — `{ param: string, value: LegacyDec }`. Example power params `m:12,n:2,c:100`; augmented `d0,p0,theta,kappa`; swapper takes none.
- **Reserve tokens** — denominations the bond trades against. `swapper_function` requires exactly two; other types require one or more.
- **Batch processing** (`bonds.proto` `Batch`) — buy/sell/swap orders are NOT settled instantly. They accumulate in the bond's current `Batch` for `batch_blocks` blocks, then settle together at a common fair price (anti-front-running). Orders unfulfillable at batch end are cancelled and reverted.
- **BatchBlocks** — lifespan (in blocks) of each batch (`math.Uint`); set at creation, immutable.
- **Alpha / outcome payments** — alpha bonds (`alpha_bond=true`, only on `augmented_function`) carry a `public_alpha` adjustable by the oracle via `MsgSetNextAlpha` (queued into the batch). `outcome_payment` is the reserve amount a funder pays via `MsgMakeOutcomePayment` to move the bond to SETTLE.
- **Bond states** (`x/bonds/types/bonds.go`): `HATCH`, `OPEN`, `SETTLE`, `FAILED`. New bonds start `OPEN`, except `augmented_function`/`bonding_function` which start `HATCH`. State is not set by the creator.

## Messages
Source order from `tx.proto`. All amount/Dec/Int/Uint fields are proto `string` carrying the noted custom type.

### MsgCreateBond
- **Purpose:** Create a new bonding-curve bond.
- **Signer / auth:** `creator_address` signs. No on-chain authority check at creation; `creator_did`/`controller_did`/`oracle_did` are recorded and gate later messages.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `bond_did` | string | yes | Bond DID (valid `did:ixo:` DID); must not already exist. |
| `token` | string | yes | Bond token denom to mint; must be valid, unused (zero bank supply, no denom metadata), not the staking denom, not reserved. |
| `name` | string | yes | Title. |
| `description` | string | yes | Description. |
| `function_type` | string | yes | One of `power_function`, `sigmoid_function`, `swapper_function`, `augmented_function`, `bonding_function`. |
| `function_parameters` | repeated FunctionParam (`FunctionParams`, non-nullable) | type-dep | Curve params; validated per `function_type`. Empty for swapper. |
| `creator_did` | string (DIDFragment) | yes | Creator DID. |
| `controller_did` | string (DIDFragment) | yes | Controller DID (authorizes alpha/state/withdraw-reserve). |
| `oracle_did` | string (DIDFragment) | yes* | Oracle DID. *ValidateBasic rejects empty; CLI flag is not marked required but a non-empty value is required. Authorizes `MsgSetNextAlpha`. |
| `reserve_tokens` | repeated string | yes | Reserve denoms (exactly 2 for swapper; ≥1 otherwise). |
| `tx_fee_percentage` | string (LegacyDec) | yes | Fee on buys/sells/swaps; ≥0. |
| `exit_fee_percentage` | string (LegacyDec) | yes | Extra fee on sells; ≥0; `tx_fee+exit_fee` must be `< 100`. |
| `fee_address` | string | yes | Bech32 account receiving fees; must not be a blocked address. |
| `reserve_withdrawal_address` | string | yes | Bech32 account reserve is withdrawn to. |
| `max_supply` | Coin (non-nullable) | yes | Max mintable; denom must equal `token`; amount > 0. |
| `order_quantity_limits` | repeated Coin (`Coins`, non-nullable) | no | Per-order max amounts; may be empty. |
| `sanity_rate` | string (LegacyDec) | no | Swapper r1/r2 sanity rate; ≥0; `0` disables. |
| `sanity_margin_percentage` | string (LegacyDec) | no | Allowed deviation; ≥0. |
| `allow_sells` | bool | no | Enable sells. Cannot both be true with `allow_reserve_withdrawals`. |
| `allow_reserve_withdrawals` | bool | no | Enable reserve withdrawals. |
| `alpha_bond` | bool | no | Mark as alpha bond; requires `augmented_function`. |
| `batch_blocks` | string (Uint) | yes | Batch lifespan in blocks; > 0. |
| `outcome_payment` | string (Int) | no | Outcome payment to settle; ≥0 (defaults to 0 via CLI). |
| `creator_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds create-bond [flags]` (no positional args; all values are flags). Required flags: `--token --name --description --function-type --function-parameters --reserve-tokens --tx-fee-percentage --exit-fee-percentage --fee-address --reserve-withdrawal-address --max-supply --order-quantity-limits --sanity-rate --sanity-margin-percentage --batch-blocks --bond-did --creator-did --controller-did`. Optional flags: `--oracle-did --allow-sells --allow-reserve-withdrawals --alpha-bond --outcome-payment`. (Despite ValidateBasic requiring a non-empty oracle DID, `--oracle-did` is not flagged required by the CLI; `--order-quantity-limits`/`--sanity-rate`/`--sanity-margin-percentage` are CLI-required even though the proto allows empty.)
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgCreateBond',
    value: ixo.bonds.v1beta1.MsgCreateBond.fromPartial({
      bondDid, token, name, description,
      functionType: 'augmented_function',
      functionParameters: [ /* ixo.bonds.v1beta1.FunctionParam.fromPartial({ param, value }) */ ],
      creatorDid, controllerDid, oracleDid,
      reserveTokens: ['uixo'],
      txFeePercentage: '0', exitFeePercentage: '0',
      feeAddress, reserveWithdrawalAddress,
      maxSupply: { denom: token, amount: '1000000' },
      sanityRate: '0', sanityMarginPercentage: '0',
      batchBlocks: '1',
      outcomePayment: '0',
      creatorAddress,
    }),
  };
  ```
- **Gotchas:** Token must be brand-new across the whole bank module — non-zero supply OR existing denom metadata is rejected (`ErrBondTokenAlreadyInUse`) to avoid colliding with e.g. a liquidstake LST. Errors: `ErrBondAlreadyExists`, `ErrBondTokenIsTaken`, `ErrBondTokenCannotBeStakingToken`, `ErrReservedBondToken`, `ErrFeesCannotBeOrExceed100Percent`, `ErrCannotAllowSellsAndWithdrawals`, `ErrMaxSupplyDenomDoesNotMatchTokenDenom`. For `augmented_function` the keeper auto-derives and appends `R0,S0,V0` (and for alpha bonds `I0,publicAlpha,systemAlpha`) to `function_parameters`. Augmented/bonding bonds force initial state `HATCH`.

### MsgEditBond
- **Purpose:** Edit mutable bond fields.
- **Signer / auth:** `editor_address` signs; `editor_did` must equal the bond's `creator_did` (note: creator, not controller).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `bond_did` | string | yes | Bond to edit. |
| `name` | string | yes | New name, or `[do-not-modify]`. |
| `description` | string | yes | New description, or `[do-not-modify]`. |
| `order_quantity_limits` | string | no | New limits string (e.g. `100res,200rez`), or `[do-not-modify]`; may be blank. |
| `sanity_rate` | string | yes | New sanity rate, or `[do-not-modify]`; `""` clears sanity to 0/0. |
| `sanity_margin_percentage` | string | yes | New margin, or `[do-not-modify]`. |
| `editor_did` | string (DIDFragment) | yes | Must be the bond creator. |
| `editor_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds edit-bond [flags]`. Required flags: `--bond-did --editor-did`. Editable flags default to `[do-not-modify]`: `--name --description --order-quantity-limits --sanity-rate --sanity-margin-percentage`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgEditBond',
    value: ixo.bonds.v1beta1.MsgEditBond.fromPartial({
      bondDid,
      name: '[do-not-modify]',
      description: '[do-not-modify]',
      orderQuantityLimits: '[do-not-modify]',
      sanityRate: '[do-not-modify]',
      sanityMarginPercentage: '[do-not-modify]',
      editorDid, editorAddress,
    }),
  };
  ```
- **Gotchas:** `name`, `description`, `sanity_rate`, `sanity_margin_percentage`, `editor_did` must be non-empty in ValidateBasic (pass `[do-not-modify]` to leave unchanged). Editor must be the creator (`ErrUnauthorized` "editor must be the creator of the bond"). Bond must exist (`ErrBondDoesNotExist`).

### MsgSetNextAlpha
- **Purpose:** Queue a new public alpha for an alpha bond into the current batch.
- **Signer / auth:** `oracle_address` signs; `oracle_did` must equal the bond's `oracle_did`. (The error string says "editor must be the controller" but the code compares against `oracle_did`.)
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `bond_did` | string | yes | Target bond. |
| `alpha` | string (LegacyDec) | yes | New public alpha; must satisfy `0.0001 ≤ alpha ≤ 0.9999` and differ from current. |
| `delta` | string (LegacyDec, nullable) | no | Optional delta (proto-nullable; not taken by the CLI). |
| `oracle_did` | string (DIDFragment) | yes | Must be the bond oracle. |
| `oracle_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds set-next-alpha [new-alpha] [bond-did] [editor-did]` (exactly 3 positional args; the third arg maps to `oracle_did`). No required flags beyond standard tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgSetNextAlpha',
    value: ixo.bonds.v1beta1.MsgSetNextAlpha.fromPartial({
      bondDid, alpha: '0.5', oracleDid, oracleAddress,
    }),
  };
  ```
- **Gotchas:** Bond must be `augmented_function` or `bonding_function` AND `alpha_bond=true`, AND state `OPEN`, else `ErrFunctionNotAvailableForFunctionType` / `ErrInvalidStateForAction`. Alpha equal to current, or system-alpha invariants `I0 > C*newSystemAlpha` / `R/C > Δalpha` violations, raise `ErrInvalidAlpha`. Only the next-alpha is stored on the batch; it applies at batch settlement.

### MsgUpdateBondState
- **Purpose:** Move a bond to SETTLE or FAILED.
- **Signer / auth:** `editor_address` signs; `editor_did` must equal the bond's `controller_did`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `bond_did` | string | yes | Target bond. |
| `state` | string | yes | Must be `SETTLE` or `FAILED` and a valid progression from current state. |
| `editor_did` | string (DIDFragment) | yes | Must be the bond controller. |
| `editor_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds update-bond-state [new-state] [bond-did] [editor-did]` (exactly 3 positional args).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgUpdateBondState',
    value: ixo.bonds.v1beta1.MsgUpdateBondState.fromPartial({
      bondDid, state: 'SETTLE', editorDid, editorAddress,
    }),
  };
  ```
- **Gotchas:** Keeper requires `function_type == augmented_function` (`ErrFunctionNotAvailableForFunctionType`). Transition must be valid (`ErrInvalidStateProgression`; valid: HATCH→OPEN/FAILED, OPEN→SETTLE/FAILED). Cannot settle/fail while the batch has outstanding orders (`ErrUnauthorized` "cannot update bond state ... while there are orders in the batch"). On SETTLE/FAILED the outcome-payment reserve is moved into the reserve and reserve balances are set to available reserve, enabling `MsgWithdrawShare`.

### MsgBuy
- **Purpose:** Submit a buy order (mints bond tokens at batch settlement).
- **Signer / auth:** `buyer_address` signs; `buyer_did`'s IID doc must resolve to that blockchain address.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `buyer_did` | string (DIDFragment) | yes | Buyer DID; must have an IID doc with matching verification method. |
| `amount` | Coin (non-nullable) | yes | Bond tokens to buy; denom must equal the bond's `token`. |
| `max_prices` | repeated Coin (`Coins`, non-nullable) | yes | Max reserve to pay; denoms must equal the bond's reserve tokens; locked until settlement. |
| `bond_did` | string | yes | Target bond. |
| `buyer_address` | string (jsontag `buyer_address`) | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds buy [bond-token-with-amount] [max-prices] [bond-did] [buyer-did]` (exactly 4 positional args; e.g. `buy 10abc 1000res1,1000res2 <bond-did> <buyer-did>`).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgBuy',
    value: ixo.bonds.v1beta1.MsgBuy.fromPartial({
      buyerDid,
      amount: { denom: 'abc', amount: '10' },
      maxPrices: [{ denom: 'uixo', amount: '1000' }],
      bondDid, buyerAddress,
    }),
  };
  ```
- **Gotchas:** Bond state must be `HATCH` or `OPEN` (`ErrInvalidStateForAction`). `max_prices` denoms must match reserve tokens (`ErrReserveDenomsMismatch`) and not exceed balance (sending to module enforces this). Order-limit breach → `ErrOrderQuantityLimitExceeded`; wrong token denom → `ErrBondTokenDoesNotMatchBond`. First buy of an empty `swapper_function` bond is special: `max_prices` become the actual price (no fee), one set of reserves is deposited, and it must not violate the sanity rate (`ErrValuesViolateSanityRate`). Settlement happens later in EndBlock, not in this call.

### MsgSell
- **Purpose:** Submit a sell order (burns bond tokens, returns reserve at settlement).
- **Signer / auth:** `seller_address` signs; `seller_did` IID doc must resolve to that address.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `seller_did` | string (DIDFragment) | yes | Seller DID with matching IID doc. |
| `amount` | Coin (non-nullable) | yes | Bond tokens to sell; denom must equal the bond's `token`. |
| `bond_did` | string | yes | Target bond. |
| `seller_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds sell [bond-token-with-amount] [bond-did] [seller-did]` (exactly 3 positional args).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgSell',
    value: ixo.bonds.v1beta1.MsgSell.fromPartial({
      sellerDid,
      amount: { denom: 'abc', amount: '10' },
      bondDid, sellerAddress,
    }),
  };
  ```
- **Gotchas:** Bond must have `allow_sells=true` (`ErrBondDoesNotAllowSelling`) and state `OPEN` (`ErrInvalidStateForAction`; augmented bonds in HATCH cannot sell). Order-limit and token-denom checks as for buy. Tokens are burned immediately; reserve returned (minus tx + exit fees) at batch settlement. Sell orders cannot be cancelled.

### MsgSwap
- **Purpose:** Submit a swap order between the two reserve tokens of a swapper bond.
- **Signer / auth:** `swapper_address` signs; `swapper_did` IID doc must resolve to that address.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `swapper_did` | string (DIDFragment) | yes | Swapper DID with matching IID doc. |
| `bond_did` | string | yes | Target bond. |
| `from` | Coin (non-nullable) | yes | Reserve token + amount to swap in. |
| `to_token` | string | yes | Reserve denom to receive; must differ from `from`. |
| `swapper_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds swap [from-amount] [from-token] [to-token] [bond-did] [swapper-did]` (exactly 5 positional args; e.g. `swap 100 res1 res2 <bond-did> <swapper-did>`).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgSwap',
    value: ixo.bonds.v1beta1.MsgSwap.fromPartial({
      swapperDid,
      bondDid,
      from: { denom: 'res1', amount: '100' },
      toToken: 'res2',
      swapperAddress,
    }),
  };
  ```
- **Gotchas:** Only `swapper_function` bonds in state `OPEN` (`ErrFunctionNotAvailableForFunctionType` / `ErrInvalidStateForAction`). `from` denom + `to_token` must exactly be the bond's two reserve tokens (`ErrReserveDenomsMismatch`); identical from/to is rejected. Order-limit applies. `from` is taken immediately into the batch intermediary account; settles at batch end (minus fee).

### MsgMakeOutcomePayment
- **Purpose:** Pay the outcome payment into the bond's outcome reserve.
- **Signer / auth:** `sender_address` signs; `sender_did` IID doc must resolve to that address.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `sender_did` | string (DIDFragment) | yes | Payer DID with matching IID doc. |
| `amount` | string (Int) | yes | Payment amount (converted to reserve coins by the bond). |
| `bond_did` | string | yes | Target bond. |
| `sender_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds make-outcome-payment [bond-did] [amount] [sender-did]` (exactly 3 positional args; arg order is bond-did, amount, sender-did).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgMakeOutcomePayment',
    value: ixo.bonds.v1beta1.MsgMakeOutcomePayment.fromPartial({
      senderDid, amount: '100', bondDid, senderAddress,
    }),
  };
  ```
- **Gotchas:** Bond must be state `OPEN` (`ErrInvalidStateForAction`). Amount must be a valid integer (CLI: `ErrArgumentMustBeInteger`) and the sender must afford it. Payment is deposited into the outcome-payment reserve (it does not itself flip state to SETTLE in the current keeper; settlement is driven by `MsgUpdateBondState`).

### MsgWithdrawShare
- **Purpose:** Burn held bond tokens for a proportional share of the reserve after settlement.
- **Signer / auth:** `recipient_address` signs; `recipient_did` IID doc must resolve to that address.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `recipient_did` | string (DIDFragment) | yes | Holder DID with matching IID doc. |
| `bond_did` | string | yes | Target bond. |
| `recipient_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds withdraw-share [bond-did] [recipient-did]` (exactly 2 positional args).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgWithdrawShare',
    value: ixo.bonds.v1beta1.MsgWithdrawShare.fromPartial({
      recipientDid, bondDid, recipientAddress,
    }),
  };
  ```
- **Gotchas:** Bond must be in `SETTLE` or `FAILED` (`ErrInvalidStateForAction`) — note FAILED is allowed even though the CLI help says "settlement state". Recipient must hold bond tokens (`ErrNoBondTokensOwned`). Share = ownedTokens / currentSupply of the remaining reserve, truncated down (later withdrawers absorb rounding).

### MsgWithdrawReserve
- **Purpose:** Controller withdraws available reserve out of the bond (alpha bonds).
- **Signer / auth:** `withdrawer_address` signs; `withdrawer_did` must equal the bond's `controller_did`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| `withdrawer_did` | string (DIDFragment) | yes | Must be the bond controller. |
| `amount` | repeated Coin (`Coins`, non-nullable) | yes | Reserve to withdraw; must be ≤ `available_reserve`. |
| `bond_did` | string | yes | Target bond. |
| `withdrawer_address` | string | yes | Signer bech32 address. |

- **CLI:** `ixod tx bonds withdraw-reserve [bond-did] [amount] [withdrawer-did]` (exactly 3 positional args; e.g. `withdraw-reserve <bond-did> 1000res <withdrawer-did>`).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.bonds.v1beta1.MsgWithdrawReserve',
    value: ixo.bonds.v1beta1.MsgWithdrawReserve.fromPartial({
      withdrawerDid,
      amount: [{ denom: 'res', amount: '1000' }],
      bondDid, withdrawerAddress,
    }),
  };
  ```
- **Gotchas:** Requires `function_type == augmented_function`, `alpha_bond == true`, state `OPEN`, and `allow_reserve_withdrawals == true` (`ErrFunctionNotAvailableForFunctionType` / `ErrInvalidStateForAction` / `ErrUnauthorized`). Withdrawer must be the controller. Amount must be ≤ `available_reserve` (`ErrInsufficientReserveForWithdraw`). Funds go to the bond's `reserve_withdrawal_address`. Reduces `available_reserve` only; the virtual `current_reserve` is unchanged.

## Queries
CLI from `client/cli/query.go`; gRPC from `query.proto` (`ixo.bonds.v1beta1.Query`).

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| All bond DIDs | `Bonds` | `ixod query bonds bonds-list` | — | `repeated string bonds` |
| All bonds w/ state | `BondsDetailed` | `ixod query bonds bonds-list-detailed` | — | `repeated BondDetails` |
| Module params | `Params` | `ixod query bonds params` | — | `Params` |
| One bond | `Bond` | `ixod query bonds bond [bond-did]` | bond_did | `Bond` |
| Current batch | `Batch` | `ixod query bonds batch [bond-did]` | bond_did | `Batch` |
| Last batch | `LastBatch` | `ixod query bonds last-batch [bond-did]` | bond_did | `Batch` |
| Current price | `CurrentPrice` | `ixod query bonds current-price [bond-did]` | bond_did | `DecCoins current_price` |
| Current reserve | `CurrentReserve` | `ixod query bonds current-reserve [bond-did]` | bond_did | `Coins current_reserve` |
| Available reserve | `AvailableReserve` | `ixod query bonds available-reserve [bond-did]` | bond_did | `Coins available_reserve` |
| Price at supply | `CustomPrice` | `ixod query bonds price [bond-token-with-amount] [bond-did]` | bond_did, bond_amount | `DecCoins price` |
| Buy price | `BuyPrice` | `ixod query bonds buy-price [bond-token-with-amount] [bond-did]` | bond_did, bond_amount | adjusted_supply, prices, tx_fees, total_prices, total_fees |
| Sell return | `SellReturn` | `ixod query bonds sell-return [bond-token-with-amount] [bond-did]` | bond_did, bond_amount | adjusted_supply, returns, tx_fees, exit_fees, total_returns, total_fees |
| Swap return | `SwapReturn` | `ixod query bonds swap-return [bond-did] [from-token-with-amount] [to-token]` | bond_did, from_token_with_amount, to_token | total_returns, total_fees |
| Alpha maximums | `AlphaMaximums` | `ixod query bonds alpha-maximums [bond-did]` | bond_did | max_system_alpha_increase, max_system_alpha |

Note: for `price`/`buy-price`/`sell-return` the CLI accepts a coin string (e.g. `10abc`) but only the numeric amount is sent as `bond_amount`.

## Events
Typed events from `event.proto` (emitted via `EmitTypedEvents`):
- `BondCreatedEvent` — on bond creation; carries the new `Bond`.
- `BondUpdatedEvent` — on `MsgEditBond`; carries the updated `Bond`.
- `BondSetNextAlphaEvent` — when next batch alpha is set (`bond_did`, `next_alpha`, `signer`).
- `BondBuyOrderEvent` — a buy order added to the batch (`order`, `bond_did`).
- `BondSellOrderEvent` — a sell order added to the batch.
- `BondSwapOrderEvent` — a swap order added to the batch.
- `BondMakeOutcomePaymentEvent` — outcome payment made (`bond_did`, `outcome_payment`, `sender_did`, `sender_address`).
- `BondWithdrawShareEvent` — share withdrawn (`bond_did`, `withdraw_payment`, `recipient_did`, `recipient_address`).
- `BondWithdrawReserveEvent` — reserve withdrawn (`bond_did`, `withdraw_amount`, `withdrawer_did`, `withdrawer_address`, `reserve_withdrawal_address`).
- `BondEditAlphaSuccessEvent` / `BondEditAlphaFailedEvent` — alpha edit applied at batch settlement succeeded/failed.
- `BondBuyOrderFulfilledEvent` / `BondSellOrderFulfilledEvent` / `BondSwapOrderFulfilledEvent` — order settled in a batch (charged prices/fees, returned amounts, new balances).
- `BondBuyOrderCancelledEvent` — buy order cancelled (unfulfillable / max-price breach).

## Module gotchas
- **Batch settlement, not instant:** buy/sell/swap only register orders; minting/burning and final pricing happen when the batch closes (`batch_blocks` later) in EndBlock. Prices are computed across all orders in the batch to deter front-running.
- **Price protection:** buys must supply `max_prices` (locked, denom-checked against reserves); buyers/sellers can be order-quantity-limited. Sell orders are irreversible; buy orders may be auto-cancelled if max price is breached.
- **DID-tied:** every order/admin message carries a DID. Buy/sell/swap/outcome/withdraw-share require the signer DID to have an IID document whose verification method resolves to the signing blockchain address (`signer must be payment contract payer` / "Address not found in iid doc" on failure). Admin messages compare `editor_did`/`oracle_did`/`withdrawer_did` against the bond's creator/oracle/controller respectively.
- **Function type drives math & params:** `swapper_function` needs exactly two reserve tokens and no params (first buy seeds the price); `augmented_function` needs `d0,p0,theta,kappa` and supports alpha + reserve withdrawals + outcome settlement; `power_function`/`sigmoid_function` are supply-priced curves.
- **Fees & reserve:** `tx_fee_percentage` (all trades) + `exit_fee_percentage` (sells) must sum to `< 100%`; fees go to `fee_address`. `current_reserve` is the virtual curve reserve; `available_reserve` is what the controller may still withdraw (alpha bonds only).
- **Token denom must be globally fresh** at creation — non-zero bank supply or registered denom metadata for the symbol is rejected; it also cannot be the staking denom or a reserved bond token.
