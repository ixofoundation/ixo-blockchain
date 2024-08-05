# Concepts

## DID Documents

A DID (Decentralized Identifier) Document is a standardized format that describes how to reach a DID subject. This format is inherently decentralized, meaning it isn't dependent on a centralized authority or a single point of failure.

### Key Components of DID Documents:

- **DID Subject:** The entity that the DID Document describes. It is identified by the DID itself.
- **Public Keys:** Used for authenticating the DID subject and, in some cases, for recovering lost or compromised DIDs.
- **Authentication Methods:** Specify the mechanisms that can be used to authenticate as the DID subject.
- **Service Endpoints:** These enable the discovery of services that can be utilized to interact with the DID subject or associated entities.

DID Documents in our system also adhere to the JSON-LD (JSON Linked Data) protocols, which allow data interchange and integration of linked data using JSON. JSON-LD ensures a standardized, extensible, and context-aware representation of data in JSON format.

### Advantages:

- **Decentralization:** DID Documents leverage decentralized systems, negating the vulnerabilities of central points of failure.
- **Interoperability:** Their standardized format ensures DID Documents can be consistently interpreted across varying systems and platforms.
- **Security and Privacy:** Their decentralized nature guarantees users retain full control over their identifiers, freeing them from reliance on potentially vulnerable centralized systems.

For a comprehensive understanding and detailed specifications of DID Documents, refer to the official [W3C DID specification](https://www.w3.org/TR/did-core/). For specifics on JSON-LD, consult the [W3C JSON-LD specification](https://www.w3.org/TR/json-ld/).

![DID-Document](./assets/did_doc.svg)

# IID Module

<div  style=' background-color:#444;'>
The `iid` module encapsulates our implementation of DID Document management on our blockchain. Built atop the sturdy foundation of the Cosmos SDK, this module supports the creation, reading, and updating of DID Documents, ensuring these operations are in compliance with W3C's official standards. By following both the DID and JSON-LD specifications, we provide a decentralized platform optimized for dependable DID management.
</div>
<br />

Every digital asset in the Cosmos context should have a universally addressable Interchain Identifier (IID).
Interchain Identifiers are a standards-compliant mechanism for uniquely identifying and referring to digital assets within chain namespaces.
IIDs also enable (off-chain) assertions to be made about (on-chain) digital assets – for instance, in the form of Verifiable Credentials.
Each IID is associated with an IID Document, which contains all the data needed to compose and interact with an asset's properties and services.

## Context

Applications using this module will enable assets on Cosmos networks to be interoperable with the systems and tooling for Decentralised Identifiers (DIDs), which conform with the [W3C DID Core](https://w3c.github.io/did-core/) and related family of specifications.
The types of assets for which this is relevant includes Non-Fungible Tokens (NFTs), Fungible Tokens, Wallets Accounts, Self-sovereign Digital Identifiers, Name Records, or any other uniquely identifiable asset type.

Any module which performs IID registry functions may implement the same methods as the IID module. Or application modules may use the services of the IID Module to performs these functions.

Integrating IID registry functions within application-specific modules has the advantage of reducing redundancies by having the application module as a context.

For application chains which have multiple modules that use IIDs (e.g. an NFT module plus Fungible Tokens module, plus Self-sovereign Identity Module), developers might find it more convenient to include the IID Module to service all their application-specific modules.

## Concepts

### DIDs and IIDs

Decentralized Identifiers (DIDs) are the [W3C specification](https://w3c.github.io/did-core/) for identifying any subject in the physical or digital realm. DIDs implement standardised DID Methods to produce fully-qualified Universal Resource Identifier (URI), as defined by RFC3986.

Interchain Identifiers (IIDs) are a DID Method for identifying on-chain assets – such as NFTs, fungible tokens, namespace records and account wallets.

### IID Document

Properties of an IID are conceptually stored in the format of an IID Document object. Which contains core properties (as defined by [W3C DID Core](https://w3c.github.io/did-core/)):

- Identifiers
  - DID Subject
  - DID Controller
  - Also Known As
- Verification Methods
  - Cryptographic material
- Verification Relationships
  - Authentication
  - Assertion
  - Key Agreement
  - Capability Invocation
  - Capability Delegation
- Service

As well as additional property sets which are unique to digital assets:

- Linked Resources
- Linked Claims
- Linked Entities
- Accorded Rights

Property extensions may be added by application developers who need to implement their own [DID Method](https://w3c.github.io/did-core/#method-syntax) for a specific use case (although it is anticipated that the IID method should serve most).

### IID Registry

The IID Module is a [Verifiable Data Registry](https://w3c.github.io/did-core/#dfn-verifiable-data-registry) system to CRU(without the D) decentralized identifiers and IID documents.

Resolving a given IID using the IID Module services returns the data necessary to produce an IID document in a [DID-conformant format](https://w3c.github.io/did-core/#dfn-did-documents), which can be serialized as JSON-LD.

### IID Method

The IID Method defines a standard way to:

- Create an IID
- Set the properties associated with the IID
- Read (resolve) an IID Document to produce a conformant JSON-LD representation.
- Update IID Document properties
- Deactivate an IID

### IID Resolver

The IID resolver is a [DID resolver](https://w3c.github.io/did-core/#dfn-did-resolvers) service that takes an IID as input and produces an IID Document as output (which conforms to the [W3C DID documents](https://w3c.github.io/did-core/#dfn-did-documents) format).

### IID Deactivation

If an IID has been [deactivated](https://w3c.github.io/did-core/#method-operations), the IID document metadata includes a property `deactivated` with the boolean value `true`.
