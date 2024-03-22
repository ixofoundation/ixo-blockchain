# State

## Iids

An IidDocument is stored in the state and is accessed by the id of the IidDocument(user provided).

- Iids: `0x01 | iidId(DID) -> ProtocolBuffer(IidDocument)`

# Types

### IidDocument

```go
type IidDocument struct {
	Context              []*Context
	Id                   string
	Controller           []string
	VerificationMethod   []*VerificationMethod
	Service              []*Service
	Authentication       []string
	AssertionMethod      []string
	KeyAgreement         []string
	CapabilityInvocation []string
	CapabilityDelegation []string
	LinkedResource       []*LinkedResource
	AccordedRight        []*AccordedRight
	LinkedEntity         []*LinkedEntity
	LinkedClaim          []*LinkedClaim
	AlsoKnownAs          string
  IidMetadata          *IidMetadata
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the DID document.
- `context` - a list of [Context](#context)
- `controller` - a list of strings of controllers(DIDs). A DID controller is an entity that is authorized to make changes to a DID document. https://www.w3.org/TR/did-core/#did-controller
- `verificationMethod` - a list of [VerificationMethod](#verification-method)
- `service` - a list of [Service](#service)
- `linkedResource` - a list of [LinkedResource](#linked-resource)
- `accordedRight` - a list of [AccordedRight](#accorded-right)
- `linkedEntity` - a list of [LinkedEntity](#linked-entity)
- `linkedClaim` - a list of [LinkedClaim](#linked-claim)
- `authentication` - a list of strings. Authentication represents public keys associated with the did document. https://www.w3.org/TR/did-core/#authentication
- `assertionMethod` - a list of strings. Used to specify how the DID subject is expected to express claims, such as for the purposes of issuing a Verifiable Credential. https://www.w3.org/TR/did-core/#assertion
- `keyAgreement` - a list of strings. Used to specify how an entity can generate encryption material in order to transmit confidential information intended for the DID subject. https://www.w3.org/TR/did-core/#key-agreement
- `capabilityInvocation` - a list of strings. Used to specify a verification method that might be used by the DID subject to invoke a cryptographic capability, such as the authorization to update the DID Document. https://www.w3.org/TR/did-core/#capability-invocation
- `capabilityDelegation` - a list of strings. Used to specify a mechanism that might be used by the DID subject to delegate a cryptographic capability to another party. https://www.w3.org/TR/did-core/#capability-delegation
- `alsoKnownAs` - a string. The assertion that two or more DIDs (or other types of URI) refer to the same DID subject can be made using the alsoKnownAs property. https://www.w3.org/TR/did-core/#also-known-as
- `iidMetadata` - a [IidMetadata](#iidmetadata)

### Verification Method

A DID document can express verification methods, such as cryptographic public keys, which can be used to authenticate or authorize interactions with the DID subject or associated parties. https://www.w3.org/TR/did-core/#verification-methods

```go
type VerificationMethod struct {
	Id         string
	Type       string
	Controller string
	// Types that are valid to be assigned to VerificationMaterial:
	//	*VerificationMethod_BlockchainAccountID
	//	*VerificationMethod_PublicKeyHex
	//	*VerificationMethod_PublicKeyMultibase
	//	*VerificationMethod_PublicKeyBase58
	VerificationMaterial isVerificationMethod_VerificationMaterial
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the verification method (required).
- `type` - a string the type of verification method
- `controller` - a string containing the did of the owner of the key
- `verificationMaterial`: a string that is either
  - `blockchainAccountID` - a string representing a cosmos based account address
  - `publicKeyHex` - a string representing a public key encoded as a hex string
  - `publicKeyMultibase` - a string representing a public key encoded according to the Multibase Data Format [Hexadecimal upper-case encoding](https://datatracker.ietf.org/doc/html/draft-multiformats-multibase#appendix-B.1)
  - `publicKeyBase58` - a string representing a public key encoded as a base58 string

### Service

Services are used in DID documents to express ways of communicating with the DID subject or associated entities. https://www.w3.org/TR/did-core/#services

```go
type Service struct {
	Id              string
	Type            string
	ServiceEndpoint string
}
```

The field's descriptions is as follows:

- `id` - a string representing the service id
- `type` - a string representing the type of the service (for example: IIDComm)
- `serviceEndpoint` - a string representing the endpoint of the service, such as an URL

### Accorded Right

An Accorded Right is stored in a list within the IidDocument data structure. Accorded Right are used to - <!-- TODO  -->

```go
type AccordedRight struct {
	Type      string
	Id        string
	Mechanism string
	Message   string
	Service   string
}
```

The field's descriptions is as follows:

- `id` - a string representing the right id
- `type` - a string representing the type of the right
- `mechanism` - a string representing the mechanism of the right,
- `message` - a string representing the message pertaining to the right,
- `service` - a string representing the service this right describes

### Linked Resource

An Linked Resource is stored in a list within the IidDocument data structure. Linked Resource are used to - <!-- TODO  -->

```go
type LinkedResource struct {
	Type            string
	Id              string
	Description     string
	MediaType       string
	ServiceEndpoint string
	Proof           string
	Encrypted       string
	Right           string
}
```

The field's descriptions is as follows:

- `id` - a string representing the unique service id
- `type` - a string representing the type of linked resource
- `description ` - a string representing the description of this linked resource
- `mediaType ` - a string representing the [MIME](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types) type of the linked resource
- `serviceEndpoint ` - a string representing the endpoint of the service, such as an URL
- `proof ` - a string representing the proof to verify the linked resource
- `encrypted ` - a string representing the boolean value for whether this linked resource is encrypted or not
- `right ` - a string <!-- TODO  -->

### Linked Entity

An Linked Entity is stored in a list within the IidDocument data structure. Linked Entity are used to - <!-- TODO  -->

```go
type LinkedEntity struct {
	Type         string
	Id           string
	Relationship string
	Service      string
}
```

The field's descriptions is as follows:

- `id` - a string representing the unique linked entity id
- `relationship ` - a string representing the relationship for the linked Entity eg. `subsidiary`
- `type ` - a string representing the type of linked entity eg. `BlockchainAccount`, `Group`
- `service ` - a string representing the service (possibly uri) that the linked entity is on.

### Linked Claim

An Linked Claim is stored in a list within the IidDocument data structure. Linked Claim are used to - <!-- TODO  -->

```go
type LinkedClaim struct {
	Type            string
	Id              string
	Description     string
	ServiceEndpoint string
	Proof           string
	Encrypted       string
	Right           string
}
```

The field's descriptions is as follows:

- `id` - a string representing the unique service id
- `type` - a string representing the type of linked resource
- `description ` - a string representing the description of this linked resource
- `serviceEndpoint ` - a string representing the endpoint of the service, such as an URL
- `proof ` - a string representing the proof to verify the linked resource.
- `encrypted ` - a string representing the boolean value for whether this linked resource is encrypted or not.
- `right ` - a string <!-- TODO  -->

### Context

A Context is stored in a list within the IidDocument data structure. Context are used to define the spec for the DID Documents since it follows the jsonld protocol. https://www.w3.org/TR/did-core/#json-ld

```go
type Context struct {
	Key string
	Val string
}
```

The field's descriptions is as follows:

- `key` - a string identifying the context mainly for storage purposes
- `value ` - a string representing the context

## IidMetadata

A IidMetadata stores information relative to a DID document such as versionId, created, updated and deactivated.

```go
type IidMetadata struct {
	VersionId   string
	Created     *time.Time
	Updated     *time.Time
	Deactivated bool
}
```

The field's descriptions is as follows:

- `versionId` - the version of the last update operation for the did document
- `updated` - the timestamp of the last update operation for the did document
- `created` - the timestamp of the create operation
- `deactivated` - a boolean indicating if the iid has been deactivated

## Verification

A verification message represent a combination of a verification method and a set of verification relationships. It has the following fields:

- `relationships` - a list of strings identifying the verification relationship for the verification method
- `method` - a [verification method object](#verification-method)
- `context` - a list of strings identifying additional [json ld contexts](https://json-ld.org/spec/latest/json-ld/#the-context)
