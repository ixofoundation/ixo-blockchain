# Messages

In this section we describe the processing of the entity messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreateEntity

A `MsgCreateEntity` creates and stores a new entity doc and a corresponding iid doc with the same id(DID) along with its nft at
appropriate indexes, which is a nft on a cw721 smart contract. It includes the [MsgCreateIidDocument](/x/iid/spec/03_messages.md#msgcreateiiddocument) fields also since creating an entity also creates a corresponding [IidDocument](/x/iid/spec/02_state.md#iiddocument) with the same id as the id generated for the entity.

```go
type MsgCreateEntity struct {
	EntityType string
	EntityStatus int32
	Controller     []string
	Context        []*types.Context
	Verification   []*types.Verification
	Service        []*types.Service
	AccordedRight  []*types.AccordedRight
	LinkedResource []*types.LinkedResource
	LinkedEntity   []*types.LinkedEntity
	StartDate *time.Time
	EndDate *time.Time
	RelayerNode string
	Credentials  []string
	OwnerDid     DIDFragment
	OwnerAddress string
	Data         encoding_json.RawMessage
	AlsoKnownAs  string
	LinkedClaim  []*types.LinkedClaim
}
```

The field's descriptions is as follows:

- `entityType` - a string representing the type of entity it is eg, `dao`, `protocol`, `asset/device`
- `entityStatus` - a integer representing the status of the entity, as defined by the implementer and interpreted by client applications
- `context` - a list of [Context](/x/iid/spec/02_state.md#context)
- `controllers` - a list of strings of controllers(DIDs). A DID controller is an entity that is authorized to make changes to a DID document. https://www.w3.org/TR/did-core/#did-controller
- `verifications` - a list of [Verifications](/x/iid/spec/02_state.md#verification)
- `service` - a list of [Service](/x/iid/spec/02_state.md#service)
- `linkedResource` - a list of [LinkedResource](/x/iid/spec/02_state.md#linked-resource)
- `accordedRight` - a list of [AccordedRight](/x/iid/spec/02_state.md#accorded-right)
- `linkedEntity` - a list of [LinkedEntity](/x/iid/spec/02_state.md#linked-entity)
- `linkedClaim` - a list of [LinkedClaim](/x/iid/spec/02_state.md#linked-claim)
- `alsoKnownAs` - a string. The assertion that two or more DIDs (or other types of URI) refer to the same DID subject can be made using the alsoKnownAs property. https://www.w3.org/TR/did-core/#also-known-as
- `ownerAddress` - a string containing the cosmos address of the private key signing the transaction
- `ownerDid` - a string containing the [MsgCreateIidDocument](/x/iid/spec/03_messages.md#msgcreateiiddocument) id (aka DID) of the owner which will be added to the [IidDocument](/x/iid/spec/02_state.md#iiddocument) list of controllers
- `startDate` - a timestamp of the start date for the entity, as defined by the implementer and interpreted by client applications
- `endDate` - a timestamp of the end date for the entity, as defined by the implementer and interpreted by client applications
- `relayerNode` - a string representing the did id of the operator through which the entity was created
- `credentials` - a list of string representing the credentials of the entity to be verified

## MsgUpdateEntity

The `MsgUpdateEntity` is used to update an entity. It updates the entity with all the fields, so if a field is empty it will be updated with default go type, aka never null. For this reason also provide the previous values for fields you do not wish to update.

```go
type MsgUpdateEntity struct {
	Id string
	EntityStatus int32
	StartDate *time.Time
	EndDate *time.Time
	Credentials       []string
	ControllerDid     DIDFragment
	ControllerAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `entityStatus` - a integer representing the status of the entity, as defined by the implementer and interpreted by client applications
- `startDate` - a timestamp of the start date for the entity, as defined by the implementer and interpreted by client applications
- `endDate` - a timestamp of the end date for the entity, as defined by the implementer and interpreted by client applications
- `credentials` - a list of string representing the credentials of the entity to be verified
- `ControllerAddress` - a string containing the cosmos address of the private key signing the transaction
- `ControllerDid` - a string containing the signers IidDocument Id. Must be in the IidDocument's controllers list to allow update and the `ControllerAddress` must be authorized to sign on behalf of the `ControllerDid`

## MsgUpdateEntityVerified

The `MsgUpdateEntityVerified` is used to update the `EntityVerified` field of an entity. Only the relayerNode for the entity can update it.

```go
type MsgUpdateEntityVerified struct {
	Id string
	EntityVerified     bool
	RelayerNodeDid     DIDFragment
	RelayerNodeAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `entityVerified` - a boolean indicating whether the entity is verified or not, can be based on credentials
- `relayerNodeAddress` - a string containing the cosmos address of the private key signing the transaction
- `relayerNodeDid` - a string containing the signers IidDocument Id. Must be in the `RelayerNode` of the entity and the `RelayerNodeAddress` must be authorized to sign on behalf of the `RelayerNodeDid`

## MsgTransferEntity

The `MsgTransferEntity` is used to transfer an entity to a new user. It will change the controllers on the [IidDocument](/x/iid/spec/02_state.md#iiddocument) to the `recipientDid` as well as remove all the [Verification Method](/x/iid/spec/02_state.md#verification-method) and add the the `recipientDid` as a new one. It also transfers the nft representing the entity on the cw721 smart contract to the recipient.

```go
type MsgTransferEntity struct {
	Id string
	OwnerDid DIDFragment
	OwnerAddress string
	RecipientDid DIDFragment
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `recipientDid` - a string containing the recipient IidDocument Id (Did).
- `ownerAddress` - a string containing the cosmos address of the private key signing the transaction.
- `ownerDid` - a string containing the signers IidDocument Id. Must be in the IidDocument's controllers list to allow update and the `ownerAddress` must be authorized to sign on behalf of the `ownerDid`

## MsgCreateEntityAccount

The `MsgCreateEntityAccount` is used to create additional [Entity Accounts](02_state.md#entityaccount).

```go
type MsgCreateEntityAccount struct {
	Id string
	Name string
	OwnerAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `name` - a string containing the name for the new entity account.
- `ownerAddress` - a string containing the cosmos address of the private key signing the transaction.

## MsgGrantEntityAccountAuthz

The `MsgGrantEntityAccountAuthz` is used to create an [Authz](https://docs.cosmos.network/main/build/modules/authz) grant from entity account (as granter) to the msg `GranteeAddress` for the specific authorization

```go
type MsgGrantEntityAccountAuthz struct {
	Id string
	Name string
	GranteeAddress string
	Grant github_com_cosmos_cosmos_sdk_x_authz.Grant
	OwnerAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `name` - a string containing the name for the entity account to use as granter for the authz grant.
- `granteeAddress` - a string containing the grantee address for the authz grant.
- `grant` - a [Grant](https://docs.cosmos.network/main/build/modules/authz#grant) that will be created.
- `ownerAddress` - a string containing the cosmos address of the private key signing the transaction.

## MsgRevokeEntityAccountAuthz

The `MsgRevokeEntityAccountAuthz` is used to revoke an [Authz](https://docs.cosmos.network/main/build/modules/authz) grant from entity account (as granter) to the msg `GranteeAddress` for the specific `MsgTypeUrl`

```go
type MsgRevokeEntityAccountAuthz struct {
	Id string
	Name string
	GranteeAddress string
	MsgTypeUrl string
	OwnerAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `name` - a string containing the name for the entity account to use as granter for the authz grant.
- `granteeAddress` - a string containing the grantee address for the authz grant.
- `msgTypeUrl` - a string containing the message type url for the specific authz to revoke.
- `ownerAddress` - a string containing the cosmos address of the private key signing the transaction.
