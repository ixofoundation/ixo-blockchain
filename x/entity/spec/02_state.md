# State

## Entities

A Entity is stored in the state and is accessed by the identity of the entity(DID).

- Entities: `0x01 | entityId(DID) -> ProtocolBuffer(Entity)`

# Types

### Entity

```go
type Entity struct {
	Id              string
	Type            string
	StartDate       *time.Time
	EndDate         *time.Time
	Status          int32
	RelayerNode     string
	Credentials     []string
	EntityVerified  bool
	Accounts        []*EntityAccount
	Metadata        *EntityMetadata
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity
- `type` - a string representing the type of entity it is eg, `dao`, `protocol`, `asset/device`
- `startDate` - a timestamp of the start date for the entity, as defined by the implementer and interpreted by client applications
- `endDate` - a timestamp of the end date for the entity, as defined by the implementer and interpreted by client applications
- `status` - a integer representing the status of the entity, as defined by the implementer and interpreted by client applications
- `relayerNode` - a string representing the did id of the operator through which the entity was created
- `credentials` - a list of string representing the credentials of the entity to be verified
- `entityVerified` - a boolean used as check whether the credentials of entity is verified
- `accounts` - a list of [EntityAccount](#entityaccount)
- `metadata` - a [EntityMetadata](#entitymetadata)

### EntityAccount

A EntityAccount is a module account generated deterministically using the entity id and name for the account and can only be controlled through the entity owner.

```go
type EntityAccount struct {
	Name         string
	Address      string
}
```

The field's descriptions is as follows:

- `name` - a string representing the name for the [Entity Account](01_concepts.md#entity-accounts)
- `address` - a string containing the address of the [Entity Account](01_concepts.md#entity-accounts)

### EntityMetadata

A EntityMetadata stores information relative to a Entity such as versionId, created, and updated.

```go
type EntityMetadata struct {
	VersionId   string
	Created     *time.Time
	Updated     *time.Time
}
```

The field's descriptions is as follows:

- `versionId` - the version of the last update operation for the entity
- `updated` - the timestamp of the last update operation for the entity
- `created` - the timestamp of the create operation
