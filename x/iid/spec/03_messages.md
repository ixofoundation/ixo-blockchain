# Messages

In this section we describe the processing of the iid messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreateIidDocument

A `MsgCreateIidDocument` is used to create a new iid document.

```go
type MsgCreateIidDocument struct {
	Id             string
	Controllers    []string
	Context        []*Context
	Verifications  []*Verification
	Services       []*Service
	AccordedRight  []*AccordedRight
	LinkedResource []*LinkedResource
	LinkedEntity   []*LinkedEntity
	AlsoKnownAs    string
	Signer         string
	LinkedClaim    []*LinkedClaim
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the DID document.
- `context` - a list of [Context](02_state.md#context)
- `controllers` - a list of strings of controllers(DIDs). A DID controller is an entity that is authorized to make changes to a DID document. https://www.w3.org/TR/did-core/#did-controller
- `verifications` - a list of [Verifications](02_state.md#verification)
- `service` - a list of [Service](02_state.md#service)
- `linkedResource` - a list of [LinkedResource](02_state.md#linked-resource)
- `accordedRight` - a list of [AccordedRight](02_state.md#accorded-right)
- `linkedEntity` - a list of [LinkedEntity](02_state.md#linked-entity)
- `linkedClaim` - a list of [LinkedClaim](02_state.md#linked-claim)
- `alsoKnownAs` - a string. The assertion that two or more DIDs (or other types of URI) refer to the same DID subject can be made using the alsoKnownAs property. https://www.w3.org/TR/did-core/#also-known-as
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgUpdateIidDocument

The `MsgUpdateIidDocument` is used to update an iid document. It updates the iid document with all the fields, so if a field is empty it will be updated with default go type, aka never null. For this reason also provide the previous values for fields you do not wish to update.

```go
type MsgUpdateIidDocument struct {
	Id             string
	Controllers    []string
	Context        []*Context
	Verifications  []*Verification
	Services       []*Service
	AccordedRight  []*AccordedRight
	LinkedResource []*LinkedResource
	LinkedEntity   []*LinkedEntity
	AlsoKnownAs    string
	Signer         string
	LinkedClaim    []*LinkedClaim
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the DID document.
- `context` - a list of [Context](02_state.md#context)
- `controllers` - a list of strings of controllers(DIDs). A DID controller is an entity that is authorized to make changes to a DID document. https://www.w3.org/TR/did-core/#did-controller
- `verifications` - a list of [Verifications](02_state.md#verification)
- `service` - a list of [Service](02_state.md#service)
- `linkedResource` - a list of [LinkedResource](02_state.md#linked-resource)
- `accordedRight` - a list of [AccordedRight](02_state.md#accorded-right)
- `linkedEntity` - a list of [LinkedEntity](02_state.md#linked-entity)
- `linkedClaim` - a list of [LinkedClaim](02_state.md#linked-claim)
- `alsoKnownAs` - a string. The assertion that two or more DIDs (or other types of URI) refer to the same DID subject can be made using the alsoKnownAs property. https://www.w3.org/TR/did-core/#also-known-as
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddVerification

The `MsgAddVerification` is used to add new [verification methods](https://w3c.github.io/did-core/#verification-methods) and [verification relationships](https://w3c.github.io/did-core/#verification-relationships) to an iid document.

```go
type MsgAddVerification struct {
	Id           string
	Verification *Verification
	Signer       string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `verification` - the [verification](02_state.md#verification) to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgSetVerificationRelationships

The `MsgSetVerificationRelationships` is used to overwrite the [verification relationships](https://w3c.github.io/did-core/#verification-relationships) for a [verification methods](https://w3c.github.io/did-core/#verification-methods) of an iid document.

```go
type MsgSetVerificationRelationships struct {
	Id            string
	MethodId      string
	Relationships []string
	Signer        string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `methodId` - a string containing the unique identifier of the verification method within the iid document.
- `relationships` - a list of strings identifying the verification relationship for the verification method
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgRevokeVerification

The `MsgRevokeVerification` is used to remove a [verification method](https://w3c.github.io/did-core/#verification-methods) and related [verification relationships](https://w3c.github.io/did-core/#verification-relationships) from an iid document.

```go
type MsgRevokeVerification struct {
	Id       string
	MethodId string
	Signer   string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `methodId` - a string containing the unique identifier of the verification method within the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddService

The `MsgAddService` is used to add a [service](https://w3c.github.io/did-core/#services) to an iid document.

```go
type MsgAddService struct {
	Id          string
	ServiceData *Service
	Signer      string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `service_data` - the [service](02_state.md#service) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteService

The `MsgDeleteService` is used to remove a [service](https://w3c.github.io/did-core/#services) from an iid document.

```go
type MsgDeleteService struct {
	Id        string
	ServiceId string
	Signer    string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `serviceData` - the [service](02_state.md#service) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddController

The `MsgAddController` is used to add a controller to an iid document.

```go
type MsgAddController struct {
	Id            string
	ControllerDid string
	Signer        string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `controllerDid` - the did string to add to the iid document's controllers
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteController

The `MsgDeleteController` is used to remove a controller from an iid document.

```go
type MsgDeleteController struct {
	Id            string
	ControllerDid string
	Signer        string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `controllerDid` - the controller did object to remove from the iid document's controllers
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddLinkedResource

The `MsgAddLinkedResource` is used to add a [LinkedResource](02_state.md#linked-resource) to an iid document.

```go
type MsgAddLinkedResource struct {
	Id             string
	LinkedResource *LinkedResource
	Signer         string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `linkedResource` - the [LinkedResource](02_state.md#linked-resource) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteLinkedResource

The `MsgDeleteLinkedResource` is used to remove a [LinkedResource](02_state.md#linked-resource) from an iid document.

```go
type MsgDeleteLinkedResource struct {
	Id         string
	ResourceId string
	Signer     string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `resourceId` - the unique id of the [LinkedResource](02_state.md#linked-resource) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddLinkedClaim

The `MsgAddLinkedClaim` is used to add a [LinkedClaim](02_state.md#linked-claim) to an iid document.

```go
type MsgAddLinkedClaim struct {
	Id          string
	LinkedClaim *LinkedClaim
	Signer      string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `linkedClaim` - the [LinkedClaim](02_state.md#linked-claim) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteLinkedClaim

The `MsgDeleteLinkedClaim` is used to remove a [LinkedClaim](02_state.md#linked-claim) from an iid document.

```go
type MsgDeleteLinkedClaim struct {
	Id      string
	ClaimId string
	Signer  string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `claimId` - the unique id of the [LinkedClaim](02_state.md#linked-claim) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddLinkedEntity

The `MsgAddLinkedEntity` is used to add a [LinkedEntity](02_state.md#linked-entity) to an iid document.

```go
type MsgAddLinkedEntity struct {
	Id           string
	LinkedEntity *LinkedEntity
	Signer       string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `linkedEntity` - the [LinkedEntity](02_state.md#linked-entity) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteLinkedEntity

The `MsgDeleteLinkedEntity` is used to remove a [LinkedEntity](02_state.md#linked-entity) from an iid document.

```go
type MsgDeleteLinkedEntity struct {
	Id       string
	EntityId string
	Signer   string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `entityId` - the unique id of the [LinkedEntity](02_state.md#linked-entity) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddAccordedRight

The `MsgAddAccordedRight` is used to add an [AccordedRight](02_state.md#accorded-right) to an iid document.

```go
type MsgAddAccordedRight struct {
	Id            string
	AccordedRight *AccordedRight
	Signer        string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `accordedRight` - the [AccordedRight](02_state.md#accorded-right) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteAccordedRight

The `MsgDeleteAccordedRight` is used to remove a [AccordedRight](02_state.md#accorded-right) from an iid document.

```go
type MsgDeleteAccordedRight struct {
	Id      string
	RightId string
	Signer  string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `rightId` - the unique id of the [AccordedRight](02_state.md#accorded-right) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgAddIidContext

The `MsgAddIidContext` is used to add a [Context](02_state.md#context) to an iid document.

```go
type MsgAddIidContext struct {
	Id      string
	Context *Context
	Signer  string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `context ` - the [Context](02_state.md#context) object to add to the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeleteIidContext

The `MsgDeleteIidContext` is used to remove a [Context](02_state.md#context) from an iid document.

```go
type MsgDeleteIidContext struct {
	Id         string
	ContextKey string
	Signer     string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `contextKey ` - the unique key of the [Context](02_state.md#context) in the iid document
- `signer` - a string containing the cosmos address of the private key signing the transaction

## MsgDeactivateIID

The `MsgDeactivateIID` is used to change the state of the [IidDocument](02_state.md#iiddocument) by changing the `Deactivated` field inside the [IidMetadata](02_state.md#iidmetadata) to true.

```go
type MsgDeactivateIID struct {
	Id     string
	State  bool
	Signer string
}
```

The field's descriptions is as follows:

- `id` - the iid string identifying the iid document
- `state` - the string representing the boolean value to change the [IidMetadata](02_state.md#iidmetadata) `Deactivated` field to, currently ignored and only changed to `false`
- `signer` - a string containing the cosmos address of the private key signing the transaction
