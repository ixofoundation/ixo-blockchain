# Entity module specification

This document specifies the entity module, a custom Ixo Cosmos SDK module.

The Entity Module introduces a holistic approach to NFT-backed identities, bridging the gap between decentralized identifiers and tangible assets. Upon entity creation, a symbiotic relationship forms between an IID Document, an NFT, and the Entity's metadata. Further enriched with the concept of Entity Accounts, this module ensures a seamless transition of ownership, while offering a robust framework for entities to operate within a decentralized landscape.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Entity Module](01_concepts.md#entity-module)
     - [Overview](01_concepts.md#overview)
     - [Key Components](01_concepts.md#key-components)
       - [Entity Data Structure](01_concepts.md#entity-data-structure)
         - [Entity Accounts](01_concepts.md#entity-accounts)
       - [NFT Creation](01_concepts.md#nft-creation)
       - [IID Document Creation](01_concepts.md#iid-document-creation)
       - [Advantages](01_concepts.md#advantages)

2. **[State](02_state.md)**

   - [State](02_state.md#state)
     - [Entities](02_state.md#entities)
   - [Types](02_state.md#types)
     - [Entity](02_state.md#entity)
     - [EntityAccount](02_state.md#entityaccount)
     - [EntityMetadata](02_state.md#entitymetadata)

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateEntity](03_messages.md#msgcreateentity)
     - [MsgUpdateEntity](03_messages.md#msgupdateentity)
     - [MsgUpdateEntityVerified](03_messages.md#msgupdateentityverified)
     - [MsgTransferEntity](03_messages.md#msgtransferentity)
     - [MsgCreateEntityAccount](03_messages.md#msgcreateentityaccount)
     - [MsgGrantEntityAccountAuthz](03_messages.md#msggrantentityaccountauthz)

4. **[Events](04_events.md)**

   - [Events](04_events.md#events)
     - [EntityCreatedEvent](04_events.md#entitycreatedevent)
     - [EntityUpdatedEvent](04_events.md#entityupdatedevent)
     - [EntityVerifiedUpdatedEvent](04_events.md#entityverifiedupdatedevent)
     - [EntityTransferredEvent](04_events.md#entitytransferredevent)
     - [EntityAccountCreatedEvent](04_events.md#entityaccountcreatedevent)
     - [EntityAccountAuthzCreatedEvent](04_events.md#entityaccountauthzcreatedevent)

5. **[Parameters](05_params.md)**

6. **[Future Improvements](06_future_improvements.md)**
