# Bonds module specification

This document specifies the bonds module, a custom Ixo Cosmos SDK module.

The bonds module provides universal token bonding curve functions to mint, burn or swap any token in a Cosmos blockchain. Once the Inter-Blockchain Communication (IBC) protocol is available, this should enable cross-network exchanges of tokens at algorithmically-determined prices.

The bonds module can deliver applications such as:

- Automated market-makers (like [Uniswap](https://uniswap.io))
- Decentralised exchanges (like [Bancor](https://bancor.network))
- Curation markets (like [Relevant](https://github.com/relevant-community/contracts/tree/bondingCurves/contracts))
- Development Impact Bonds (like ixo alpha-Bonds)
- Continuous organisations (like [Moloch DAO](https://molochdao.com/))

Any Cosmos application chain that implements the Bonds module is able to perform functions such as:

- Issue a new token with custom parameters.
- Pool liquidity for reserves.
- Provide continuous funding.
- Automatically mint and burn tokens at deterministic prices.
- Swap tokens atomically within the same network.
- Exchange tokens across networks, with the IBC protocol.
- (Batch token transactions to prevent front-running)
- Launch a decentralised autonomous initial coin offerings ([DAICO](https://ethresear.ch/t/explanation-of-daicos/465))
- ..._other **DeFi**ant_ innovations.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Concepts](01_concepts.md#concepts)
     - [Token Bonding Curves](01_concepts.md#token-bonding-curves)
     - [Token Bonds Module](01_concepts.md#token-bonds-module)
     - [Batching](01_concepts.md#batching)

2. **[State](02_state.md)**

   - [State](02_state.md#state)
     - [Bonds](02_state.md#bonds)
     - [Batches](02_state.md#batches)
       - [Querying Batches](02_state.md#querying-batches)

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateBond](03_messages.md#msgcreatebond)
     - [MsgEditBond](03_messages.md#msgeditbond)
     - [MsgSetNextAlpha](03_messages.md#msgsetnextalpha)
     - [MsgUpdateBondState](03_messages.md#msgupdatebondstate)
     - [MsgBuy](03_messages.md#msgbuy)
       - [MsgBuy for Swapper Function Bonds](03_messages.md#msgbuy-for-swapper-function-bonds)
     - [MsgSell](03_messages.md#msgsell)
     - [MsgSwap](03_messages.md#msgswap)
     - [MsgMakeOutcomePayment](03_messages.md#msgmakeoutcomepayment)
     - [MsgWithdrawShare](03_messages.md#msgwithdrawshare)
     - [MsgWithdrawReserve](03_messages.md#msgwithdrawreserve)

4. **[End-Block](04_end_block.md)**

   - [End-Block](04_end_block.md#end-block)
     - [Buys](04_end_block.md#buys)
     - [Sells](04_end_block.md#sells)
     - [Swaps](04_end_block.md#swaps)
     - [Set Last Batch](04_end_block.md#set-last-batch)

5. **[Events](05_events.md)**

   - [Events](05_events.md#events)
     - [BondCreatedEvent](05_events.md#bondcreatedevent)
     - [BondUpdatedEvent](05_events.md#bondupdatedevent)
     - [BondSetNextAlphaEvent](05_events.md#bondsetnextalphaevent)
     - [BondBuyOrderEvent](05_events.md#bondbuyorderevent)
     - [BondSellOrderEvent](05_events.md#bondsellorderevent)
     - [BondSwapOrderEvent](05_events.md#bondswaporderevent)
     - [BondMakeOutcomePaymentEvent](05_events.md#bondmakeoutcomepaymentevent)
     - [BondWithdrawShareEvent](05_events.md#bondwithdrawshareevent)
     - [BondWithdrawReserveEvent](05_events.md#bondwithdrawreserveevent)
     - [BondEditAlphaSuccessEvent](05_events.md#bondeditalphasuccessevent)
     - [BondEditAlphaFailedEvent](05_events.md#bondeditalphafailedevent)
     - [BondBuyOrderFulfilledEvent](05_events.md#bondbuyorderfulfilledevent)
     - [BondSellOrderFulfilledEvent](05_events.md#bondsellorderfulfilledevent)
     - [BondSwapOrderFulfilledEvent](05_events.md#bondswaporderfulfilledevent)
     - [BondBuyOrderCancelledEvent](05_events.md#bondbuyordercancelledevent)

6. **[Parameters](06_params.md)**

7. **[Future Improvements](07_future_improvements.md)**

8. **[Functions Library](08_functions_library.ipynb)**
