# Bonds module specification

## Abstract

This document specifies the bonds bodule; a custom Cosmos SDK module.

The bonds module provides universal token bonding curve functions to mint, burn or swap any token in a Cosmos blockchain. Once the Inter-Blockchain Communication (IBC) protocol is available, this should enable cross-network exchanges of tokens at algorithmically-determined prices.

The bonds module can be deployed through Cosmos Hubs and Zones to deliver applications such as:
* Automated market-makers (like [Uniswap](https://uniswap.io))
* Decentralised exchanges (like [Bancor](https://bancor.network))
* Curation markets (like [Relevant](https://github.com/relevant-community/contracts/tree/bondingCurves/contracts))
* Development Impact Bonds (like ixo alpha-Bonds)
* Continuous organisations (like [Moloch DAO](https://molochdao.com/))

Any Cosmos application chain that implements the Bonds module is able to perform functions such as:
* Issue a new token with custom parameters.
* Pool liquidity for reserves.
* Provide continuous funding.
* Automatically mint and burn tokens at deterministic prices.
* Swap tokens atomically within the same network.
* Exchange tokens across networks, with the IBC protocol.
* (Batch token transactions to prevent front-running)
* Launch a decentralised autonomous initial coin offerings ([DAICO](https://ethresear.ch/t/explanation-of-daicos/465))
* ...*other **DeFi**ant* innovations.

## Contents

1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
    - [Bonds](02_state.md#bonds)
    - [Batches](02_state.md#batches)
3. **[Messages](03_messages.md)**
    - [MsgCreateBond](03_messages.md#msgcreatebond)
    - [MsgEditBond](03_messages.md#msgeditbond)
    - [MsgBuy](03_messages.md#msgbuy)
    - [MsgSell](03_messages.md#msgsell)
    - [MsgSwap](03_messages.md#msgswap)
4. **[End-Block](04_end_block.md)**
    - [Buys](04_end_block.md#buys)
    - [Sells](04_end_block.md#sells)
    - [Swaps](04_end_block.md#swaps)
    - [Set Last Batch](04_end_block.md#set-last-batch)
5. **[Events](05_events.md)**
    - [EndBlocker](05_events.md#endblocker)
    - [Handlers](05_events.md#handlers)
6. **[Future Improvements](06_future_improvements.md)**
7. **[Functions Library](07_functions_library.ipynb)**
