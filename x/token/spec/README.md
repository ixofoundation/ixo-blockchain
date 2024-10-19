# Token module specification

This document specifies the token module, a custom Ixo Cosmos SDK module.

Embracing the versatility of the EIP-1155 standard, the Token Module offers a sophisticated mechanism for managing multi-token smart contracts. Whether you're dealing with fungible or non-fungible tokens, this module streamlines the process of creation, minting, and management. From defining token collections to ensuring transparent on-chain token attributes, the Token Module stands as a beacon of efficiency and flexibility in the decentralized token ecosystem.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Concepts](01_concepts.md#concepts)
     - [EIP-1155 Smart Contract Tokens](01_concepts.md#eip-1155-smart-contract-tokens)
   - [Token Module](01_concepts.md#token-module)
     - [Concepts](01_concepts.md#concepts-1)
       - [Token](01_concepts.md#token)
       - [TokenProperties](01_concepts.md#tokenproperties)
     - [Features](01_concepts.md#features)
       - [1. Token Collection Creation](01_concepts.md#1-token-collection-creation)
       - [2. Minting Authorizations](01_concepts.md#2-minting-authorizations)
       - [3. Batch Operations](01_concepts.md#3-batch-operations)
       - [4. Token Uniqueness and Fungibility](01_concepts.md#4-token-uniqueness-and-fungibility)
       - [5. Token Retirement](01_concepts.md#5-token-retirement)

2. **[State](02_state.md)**

   - [State](02_state.md#state)
     - [Tokens](02_state.md#tokens)
     - [TokenProperties](02_state.md#tokenproperties)
   - [Types](02_state.md#types)
     - [Token](02_state.md#token)
     - [TokensRetired](02_state.md#tokensretired)
     - [TokensCancelled](02_state.md#tokenscancelled)
     - [TokenProperties](02_state.md#tokenproperties-1)
     - [TokenData](02_state.md#tokendata)
     - [Authz Types](02_state.md#authz-types)
       - [MintAuthorization](02_state.md#mintauthorization)
       - [MintConstraints](02_state.md#mintconstraints)
   -

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateToken](03_messages.md#msgcreatetoken)
     - [MsgMintToken](03_messages.md#msgminttoken)
     - [MsgTransferToken](03_messages.md#msgtransfertoken)
     - [MsgRetireToken](03_messages.md#msgretiretoken)
     - [MsgCancelToken](03_messages.md#msgcanceltoken)
     - [MsgPauseToken](03_messages.md#msgpausetoken)
     - [MsgStopToken](03_messages.md#msgstoptoken)
     - [Generic types](03_messages.md#generic-types)
       - [MintBatch](03_messages.md#mintbatch)
       - [TokenBatch](03_messages.md#tokenbatch)

4. **[Events](04_events.md)**

   - [Events](04_events.md#events)
     - [TokenCreatedEvent](04_events.md#tokencreatedevent)
     - [TokenUpdatedEvent](04_events.md#tokenupdatedevent)
     - [TokenMintedEvent](04_events.md#tokenmintedevent)
     - [TokenTransferredEvent](04_events.md#tokentransferredevent)
     - [TokenCancelledEvent](04_events.md#tokencancelledevent)
     - [TokenRetiredEvent](04_events.md#tokenretiredevent)
     - [TokenPausedEvent](04_events.md#tokenpausedevent)
     - [TokenStoppedEvent](04_events.md#tokenstoppedevent)

5. **[Parameters](05_params.md)**

6. **[Future Improvements](06_future_improvements.md)**
