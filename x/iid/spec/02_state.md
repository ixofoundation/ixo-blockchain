# State

This document describes the state pertaining to:

1. [IidDocument](./02_state.md#identifier)
2. [IidMetadata](./02_state.md#didmetadata)

## IidDocument
DidDocuments are stored in the state under the `0x61` key and are stored using their ids

- IidDocument: `0x61 | IidDocument.Id -> ProtocolBuffer(IidDocument)`


An iid document has the following fields:

```
type IidDocument struct {
	// @context is spec for did document.
	Context []*Context `protobuf:"bytes,1,rep,name=context,proto3" json:"@context,omitempty"`
	
	// id represents the id for the iid document.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	
	// A IID controller is an entity that is authorized to make changes to a IID document.
	Controller []string `protobuf:"bytes,3,rep,name=controller,proto3" json:"controller,omitempty"`
	
	// A IID document can express verification methods,
	// such as cryptographic public keys, which can be used
	// to authenticate or authorize interactions with the IID subject or associated parties.
	// https://www.w3.org/TR/did-core/#verification-methods
	VerificationMethod []*VerificationMethod `protobuf:"bytes,4,rep,name=verificationMethod,proto3" json:"verificationMethod,omitempty"`
	
	// Services are used in IID documents to express ways of communicating
	// with the IID subject or associated entities.
	// https://www.w3.org/TR/did-core/#services
	Service []*Service `protobuf:"bytes,5,rep,name=service,proto3" json:"service,omitempty"`
	
	// NOTE: below this line there are the relationships
	// Authentication represents public key associated with the did document.
	// cfr. https://www.w3.org/TR/did-core/#authentication
	Authentication []string `protobuf:"bytes,6,rep,name=authentication,proto3" json:"authentication,omitempty"`
	
	// Used to specify how the IID subject is expected to express claims,
	// such as for the purposes of issuing a Verifiable Credential.
	// cfr. https://www.w3.org/TR/did-core/#assertion
	AssertionMethod []string `protobuf:"bytes,7,rep,name=assertionMethod,proto3" json:"assertionMethod,omitempty"`
	
	// used to specify how an entity can generate encryption material
	// in order to transmit confidential information intended for the IID subject.
	// https://www.w3.org/TR/did-core/#key-agreement
	KeyAgreement []string `protobuf:"bytes,8,rep,name=keyAgreement,proto3" json:"keyAgreement,omitempty"`
	
	// Used to specify a verification method that might be used by the IID subject
	// to invoke a cryptographic capability, such as the authorization
	// to update the IID Document.
	// https://www.w3.org/TR/did-core/#capability-invocation
	CapabilityInvocation []string `protobuf:"bytes,9,rep,name=capabilityInvocation,proto3" json:"capabilityInvocation,omitempty"`
	
	// Used to specify a mechanism that might be used by the IID subject
	// to delegate a cryptographic capability to another party.
	// https://www.w3.org/TR/did-core/#capability-delegation
	CapabilityDelegation []string          `protobuf:"bytes,10,rep,name=capabilityDelegation,proto3" json:"capabilityDelegation,omitempty"`
	LinkedResource       []*LinkedResource `protobuf:"bytes,11,rep,name=linkedResource,proto3" json:"linkedResource,omitempty"`
	AccordedRight        []*AccordedRight  `protobuf:"bytes,12,rep,name=accordedRight,proto3" json:"accordedRight,omitempty"`
	LinkedEntity         []*LinkedEntity   `protobuf:"bytes,13,rep,name=linkedEntity,proto3" json:"linkedEntity,omitempty"`
	AlsoKnownAs          string            `protobuf:"bytes,14,opt,name=alsoKnownAs,proto3" json:"alsoKnownAs,omitempty"`
}
```

#### Source 
+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L22

### Verification Method

Verification methods are stored within an IidDocument and are used to store public key information as described in the [W3C DID recommendations](https://w3c.github.io/did-core/#verification-methods).

A verification method has the following fields:

- `id` - a string containing the unique identifier of the verification method (required).
- `type` - a string the type of verification method 
- `controller` - a string containing the did of the owner of the key 
- `verificationMaterial`: a string that is either   
  - `blockchainAccountID` - a string representing a cosmos based account address
  - `publicKeyHex` - a string representing a public key encoded as a hex string
  - `publicKeyMultibase` - a string representing a public key encoded according to the Multibase Data Format [Hexadecimal upper-case encoding](https://datatracker.ietf.org/doc/html/draft-multiformats-multibase#appendix-B.1)
  
#### Source 
+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L103


### Service
A Service is stored as a slice within a IidDocument data structure. Services are used to describe communication interfaces for an IID as described in the [W3C DID recommendations](https://w3c.github.io/did-core/#services)

A service has the following fields:

- `id` - a string representing the service id
- `type` - a string representing the type of the service (for example: IIDComm)
- `serviceEndpoint` - a string representing the endpoint of the service, such as an URL

#### Source 

+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L118

### Accorded Right
An Accorded Right is stored as a slice within a IidDocument data structure. Accorded Right are used to - shaun to provide

An Accorded Right has the following fields:

- `id` - a string representing the right id
- `type` - a string representing the type of the right 
- `mechanism` - a string representing the mechanism of the right,
- `message` - a string representing the message pertaining to the right, 
- `service` - a string representing the service this right describes

#### Source

+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L74

### Linked Resource
A Service is stored as a slice within a IidDocument data structure. Services are used to describe communication interfaces for an IID as described in the [W3C DID recommendations](https://w3c.github.io/did-core/#services)

A service has the following fields:

- `id` - a string representing the service id
- `type` - a string representing the type of linked resource
- `description ` - a string representing the description of this linked resource
- `mediaType ` - a string representing (Shaun to provide)
- `serviceEndpoint ` - a string representing (Shaun to provide)
- `proof ` - a string representing (Shaun to provide)
- `encrypted ` - a string representing (Shaun to provide)
- `right ` - a string representing (Shaun to provide)

#### Source

+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L84

### Linked Entity
A Service is stored as a slice within a IidDocument data structure. Services are used to describe communication interfaces for an IID as described in the [W3C DID recommendations](https://w3c.github.io/did-core/#services)

A service has the following fields:

- `id` - a string representing the linked entity id
- `relationship ` - a string representing (Shaun to provide)

#### Source

+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L96

## DidMetadata

A DidMetadata stores information relative to a DID document. The implementation supports the following fields: 

- `versionId` - the version of the last update operation for the did document
- `updated` - the timestamp of the last update operation for the did document 
- `created` - the timestamp of the create operation.
- `deactivated` - a boolean indicating if the iid has been deactivated
- `entityType ` - a string (shaun to provide)
- `startDate ` - a timestamp (shaun to provide)
- `endDate ` - a timestamp (shaun to provide)
- `status ` - an int32  (shaun to provide)
- `stage ` - a string (shaun to provide)
- `relayerNode ` - a string (shaun to provide)
- `verifiableCredential ` - a string (shaun to provide)
- `credentials ` - a repeated string of (shaun to provide)
#### Source 
+++ https://github.com/ixofoundation/ixo-blockchain/blob/devel/entity-module/proto/iid/iid.proto#L129


