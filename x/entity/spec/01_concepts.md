# Entity Module: Digital Twin Domains in the Spatial Web

## Overview

The Entity Module introduces a revolutionary approach to digital identity and asset representation in the Spatial Web. When an `Entity` is created, it establishes a sovereign digital twin domain through a trifecta of operations:

1. An Interoperable Identifier (IID) Document is generated via the `iid` module, establishing the foundational decentralized identity and associated metadata of the entity.
2. A Non-Fungible Token (NFT) is minted using the [CW721](https://github.com/CosmWasm/cw-nfts/blob/main/packages/cw721/README.md) smart contract, capturing the unique and immutable essence of the entity in tokenized form.
3. A dedicated `Entity` data structure is recorded in the key-value (KV) store, encompassing supplementary metadata exclusive to the `Entity`.

The `id` for all three data components is deterministically generated on-chain during Entity creation, ensuring consistency and traceability.

## Key Components

### Entity Data Structure

The `Entity` data structure, stored within the KV store, serves as a comprehensive digital twin domain representation, including:

- **Controllers**: Governing authorities of the domain.
- **Verifiable Credentials**: Cryptographic proofs for verifying attributes, bassed on [W3C standards](https://w3c.github.io/vc-overview/)
- **Linked Resources**: Digital or physical assets associated with the domain.
- **Accorded Rights**: Permissions and capabilities available within the Spatial Web.
- **Services**: Functionalities provided by the entity, including third-party web services that interface with the domain.
- **Linked Claims**: Assertions about or made by the entity.
- **Linked Entities**: Relationships with other domains in the Spatial Web graph.

### Entity Accounts

Entity Accounts are Cosmos Module Accounts uniquely associated with and controlled by an Entity Owner. Key features include:

- **Multiple Accounts**: Each Entity may have multiple accounts, each with a unique name (e.g., 'Savings'). 
- **Account Address Generation**: Account addresses are derived from the account name and Entity's DID.
- **Authz Delegation**: Authz-based delegation allows transaction permissions to be granted to other accounts, using the [Authz Module](https://docs.cosmos.network/main/modules/authz/)
- **Ownership Transfer**: Control and balances are automatically transferred when Entity ownership changes.
- **Default Admin Account**: A default 'admin' account is created when the Entity is first established.

### NFT Representation

The minted NFT via the CW721 smart contract embodies the distinct essence of each `Entity`, offering:

- **Tokenized representation** for trade and ownership verification.
- **Utilization in decentralized applications** within the Spatial Web.
- **Immutable link** to the entity's digital twin domain.

### IID Document

The automatically generated IID Document ensures each `Entity` has a standardized, robust, and decentralized identity, featuring:

- **Compliance with W3C Specifications**: Each IID complies with W3C's [DID Core](https://www.w3.org/TR/did-core/) specification for global compatibility.
- **Metadata Enrichment**: Enrichment with crucial metadata provides a comprehensive representation of the identity.

## Advantages

- **Sovereign Digital Twin Domains**: Each entity represents a self-sovereign domain in the Spatial Web, with full control over its digital representation and associated assets.
- **Enhanced Interoperability**: Rooted in globally recognized standards, `Entities` foster seamless interactions across various decentralized platforms and applications in the Spatial Web ecosystem.
- **Comprehensive Metadata Management**: The unified approach to metadata storage provides a nuanced and rich understanding of each entity's characteristics and relationships.
- **Decentralization and Security**: The combination of IID Documents and NFTs ensures a decentralized, tamper-proof, and transparent representation of entities within the Spatial Web.
- **Flexible Governance**: Entity Accounts and delegation mechanisms allow for sophisticated control and management of digital assets and interactions.

This Entity Module serves as a crucial building block in the Spatial Web, enabling interconnected digital twin domains, graph-based relationships between entities, and decentralized governance models.
