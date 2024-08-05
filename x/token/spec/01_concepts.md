# Concepts

### EIP-1155 Smart Contract Tokens

[EIP-1155](https://eips.ethereum.org/EIPS/eip-1155) represents a multi-token standard that supports fungible, semi-fungible, and non-fungible tokens within a single smart contract. The distinction is crucial as it allows for both unique and fungible token behaviors within the same contract. By adopting this standard, the `Token` module provides a flexible solution for varied token use-cases, including those in DeFi, gaming, art, and beyond.

# Token Module

The Token module offers an advanced management solution for the creation, minting, and management of EIP-1155 smart contract tokens within the decentralized ecosystem.

## Concepts

### Token

A [Token](02_state.md#token) can also be thought of as a token collection and it serves as an umbrella for individual tokens and consists of the following attributes:

- **Minter:** The primary entity authorized to mint tokens within the collection.
- **Name (Namespace):** A unique identifier for the token collection.
- **Description:** A textual overview of the token collection.
- **Image:** A visual representation or logo for the token collection.
- **Cap:** The upper limit on the number of unique EIP-1155 tokens that can be minted within the collection.
- **Current Supply Count:** A real-time count of minted tokens within the collection.
- **Retired and Cancelled Tokens List:** A registry for tracking tokens that have been retired or canceled.

Upon the creation of a token collection, a corresponding CW1155 smart contract is initiated, with its address stored within the collection for reference.

### TokenProperties

Every time a token is minted, a [TokenProperties](02_state.md#tokenproperties) object gets created and stored in the on-chain key-value store. This ensures the persistence of crucial token attributes and data, granting transparency and integrity.

## Features

### 1. Token Collection Creation

Minting starts with the establishment of a token collection, complete with all necessary attributes and an associated CW1155 smart contract.

### 2. Minting Authorizations

The primary minter has the authority to delegate minting capabilities to other entities through [authz grants](02_state.md#authz-types), further decentralizing and democratizing the minting process.

### 3. Batch Operations

Tokens can be minted, transferred and retired in batches, allowing for efficient bulk operations while maintaining the unique attributes of each individual token.

### 4. Token Uniqueness and Fungibility

Leveraging the benefits of EIP-1155, tokens within the same namespace (like "CARBON") can have distinct IDs, making them non-fungible. However, tokens sharing an ID remain fungible among themselves. This dual nature offers unparalleled versatility in token operations and value calculations.

### 5. Token Retirement

Token holders can retire their tokens as a means to "offset their footprint," offering an environmentally-conscious angle to token utility.
