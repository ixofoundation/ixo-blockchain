# Entity module specification

# IXO Entity Module: Digital Twin Domains in the Spatial Web

## Overview

The IXO **Entity** Module provides a cutting-edge approach to digital identity and real-world asset representation within the Spatial Web. By enabling sovereign digital twin domains, backed by NFTs and managed by actors with decentralized identifiers (DIDs), this module creates a network of interconnected nodes within the Spatial Web. These entities form a robust framework for decentralized operations, facilitating interactions that have tangible impacts in the real world.

## Key Features

### Sovereign Digital Twin Domains

Each entity represents a distinct, sovereign domain within the Spatial Web, including:

- **Controllers**: Authorities responsible for governing the domain.
- **Verifiable Credentials**: Cryptographic proofs for verifying attributes.
- **Linked Resources**: Digital or physical assets associated with the domain.
- **Accorded Rights**: Permissions and capabilities available within the Spatial Web.
- **Services**: Functionalities provided by the entity, including third-party web services that interface with the domain.
- **Linked Claims**: Assertions about or made by the entity.
- **Linked Entities**: Relationships with other domains in the Spatial Web.
- **Entity Accounts**: Any number of blockchain Accounts that may be used for different purposes.

### NFT-Backed Domains

- **NFT Integration**: Domains are represented as NFTs, integrated seamlessly with DIDs for enhanced security.
- **Tangible Assets**: Digital twin domains can be owned, transferred, financed, or used as collateral, similar to traditional assets.

### Interchain Identifier (IID) Document

- **Comprehensive Representation**: IIDs provide a complete representation of the identity domain.
- **Interoperability**: Supports interactions across multiple blockchain networks.

### Entity Metadata

- **Customizable Metadata**: Rich and customizable metadata allows detailed descriptions of each digital twin domain, enabling better discovery and context.

### Entity Accounts

- **Financial Management**: Dedicated accounts for managing resources, payments, and interactions between entities.

## Spatial Web Integration

The IXO Entity Module is a foundational element of the Spatial Web, enabling:

- **Interconnected Digital Twin Domains**: Digital entities that operate autonomously and are linked through the Spatial Web.
- **Graph-Based Relationships**: Establishing connections and relationships between entities, forming a robust network.
- **Decentralized Governance**: Sovereign entities enable new models of governance and ownership.

## Technical Implementation

The Entity Module is built using the IXO Cosmos SDK, leveraging blockchain technology for:

- **Immutability**: Ensuring that entity records are tamper-proof and reliable.
- **Decentralized Control**: Giving entities complete control over their domains without centralized authorities.
- **Interoperability**: Facilitating seamless interactions with other networks within the Cosmos ecosystem.

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
