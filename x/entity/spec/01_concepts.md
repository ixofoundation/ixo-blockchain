# Entity Module

The `Entity` module is a pivotal component of our blockchain system, facilitating the unified management of both non-fungible tokens (NFTs) and their associated Interchain Identifiers (IIDs). By seamlessly intertwining NFTs with IIDs, the `Entity` module establishes a new paradigm for decentralized asset management, bolstered by the credibility and dependability of DID Documents.

## Overview

When an `Entity` is created on the blockchain, a trifecta of operations is set into motion:

1. An IID Document is birthed through the `iid` module, cementing the foundational identity and associated metadata of the entity.
2. An NFT is minted using the CW721 smart contract, capturing the unique and immutable essence of the entity in tokenized form.
3. A dedicated `Entity` data structure is etched into the KV store. This is not merely a reflection of the IID Document; it captures supplementary metadata exclusive to the `Entity`, enhancing its depth and versatility.

The `id` for all 3 of the different data stored above is deterministically determined on-chain on Entity creation.

## Key Components

### Entity Data Structure

Stored within the KV store, the `Entity` data structure houses metadata that isn't contained within the IID Document. It's instrumental in furnishing a holistic view of the `Entity`, encapsulating both inherent and appended attributes. This data provides richer context, driving interoperability and comprehensive understanding across platforms and applications.

#### Entity Accounts

An Entity Account is a Cosmos Module Account that can only be created and controlled by an Entity Owner. An Entity may have any number of Entity Accounts. Each Entity Account has a unique name. For instance “Savings”. The account address is derived from this name and the DID of the Entity. The Entity Owner may authorise any other account to perform transactions for an Entity Account, using Authz. When ownership of an Entity is transferred, the new owner gains control of the Entity Accounts. All residual balances in an Entity Account get transferred together with the Entity NFT when ownership is transferred.

On Entity creation a default Entity Account gets created with the name `admin`.

### NFT Creation

The minting of an NFT via the CW721 smart contract reaffirms the distinctiveness of each `Entity`. This tokenized representation not only encapsulates the essence of the entity but also offers a myriad of possibilities in terms of trade, ownership verification, and utilization in decentralized applications.

### IID Document Creation

The automatic generation of an IID Document ensures that every `Entity` is backed by a standardized, robust, and decentralized identity mechanism. This identity is enriched with crucial metadata and, being compliant with W3C's DID specifications, offers global compatibility and recognition.

### Advantages

- **Unified Management:** The `Entity` module simplifies the intricacies of managing decentralized assets by unifying IID creation, NFT minting, and metadata storage.
- **Enhanced Metadata:** By preserving additional metadata in the `Entity` data structure, a comprehensive and nuanced understanding of each entity is achieved.
- **Interoperability:** Rooted in globally recognized standards, `Entities` foster seamless interactions across various decentralized platforms and applications.

- **Decentralization and Security:** The tandem of IID Documents and NFTs ensures a decentralized, tamper-proof, and transparent representation of entities.

For a deeper dive into the foundational principles of IIDs, consult the [`iid` module documentation](/x/iid/spec/README.md). For a comprehensive understanding of NFTs and the CW721 standard, refer to the [official CW721 documentation](https://github.com/CosmWasm/cw-nfts/blob/main/packages/cw721/README.md) which follows the [eip721 standard](https://eips.ethereum.org/EIPS/eip-721).
