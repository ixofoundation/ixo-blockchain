

# `x/iid`

## Contents

## Abstract

`x/iid` is an implementation of a custom Cosmos SDK module, that allows you to create an on chain IID identifier which is an extension of the did spec as noted [here](https://www.w3.org/TR/did-core/) , CRUD functions exist to add or delete objects from the identifier.

* [Concept](#concepts)
    * [Structure](#structure)
    * [IID](#IID)
      *[State](./spec/02_state.md)
      *[Transitions](./spec/03_state_transitions.md) 
      *[Messages](./spec/04_messages.md)

# Concepts

## Structure

`x/iid` module defines a struct `IidDocument` to house the contents of the identifier and its linked resources and capabilities.
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

## IID

shaun to provide
# State

## IidDocument

an IID document is composed of the following fields
* Context: `Context []*Context`
* Id: `Id string`
* Controller: `Controller []string`
* Verification Method: `VerificationMethod []*VerificationMethod`
* Service: `Service []*Service`
* Authentication: `Authentication []string`
* Assertion Method: `AssertionMethod []string`
* Key Agreement: `KeyAgreement []string`
* Capability Invocation: `CapabilityInvocation []string`
* Capability Delegation: `CapabilityDelegation []string`
* Linked Resource: `LinkedResource []*LinkedResource`
* Accorded Right: `AccordedRight        []*AccordedRight`
* Linked Entity: `LinkedEntity         []*LinkedEntity`
* Also known as: `AlsoKnownAs          string`