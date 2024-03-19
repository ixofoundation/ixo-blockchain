# Events

The entity module emits the following typed events:

### EntityCreatedEvent

Emitted after a successfull `MsgCreateEntity`

```go
type EntityCreatedEvent struct {
	Entity *Entity
  Signer  string
}
```

The field's descriptions is as follows:

- `entity` - the full [Entity](02_state.md#entity)
- `signer` - a string containing the did of the signer.

### EntityUpdatedEvent

Emitted after a successfull `MsgUpdateEntity`, `MsgUpdateEntityVerified`, `MsgTransferEntity`, `MsgCreateEntityAccount` whereby a field of the entity struct gets updated.

```go
type EntityUpdatedEvent struct {
	Entity *Entity
  Signer  string
}
```

The field's descriptions is as follows:

- `entity` - the full [Entity](02_state.md#entity)
- `signer` - a string containing the did of the signer or the address of the signer.

### EntityVerifiedUpdatedEvent

Emitted after a successfull `MsgUpdateEntityVerified`

```go
type EntityVerifiedUpdatedEvent struct {
	Id             string
	Signer         string
	EntityVerified bool
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `signer` - a string containing the did of the signer (will the `relayerNode` did in this case).
- `entityVerified` - a boolean indicating whether the entity is verified or not.

### EntityTransferredEvent

Emitted after a successfull `MsgTransferEntity`

```go
type EntityTransferredEvent struct {
	Id   string
	From string
	To   string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `to` - a string containing the did of the new owner.
- `from` - a string containing the did of the previous owner.

### EntityAccountCreatedEvent

Emitted after a successfull `MsgCreateEntityAccount`

```go
type EntityAccountCreatedEvent struct {
	Id             string
	Signer         string
	AccountName    string
	AccountAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `accountName` - a string containing the name of the new entity account.
- `accountAddress` - a string containing the address of the new entity account.
- `signer` - a string containing the address of the signer.

### EntityAccountAuthzCreatedEvent

Emitted after a successfull `MsgGrantEntityAccountAuthz`

```go
type EntityAccountAuthzCreatedEvent struct {
	Id          string
	Signer      string
	AccountName string
	Granter     string
	Grantee     string
	Grant       *Grant
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `signer` - a string containing the address of the signer.
- `accountName` - a string containing the name of the entity account the authz was granted.
- `granter` - a string containing the address of the entity account the authz was granted.
- `grantee` - a string containing the address of the grantee towards who the authz was granted.
- `grant` - the [Grant](https://docs.cosmos.network/main/build/modules/authz#grant) that was granted.

### EntityAccountAuthzRevokedEvent

Emitted after a successfull `MsgRevokeEntityAccountAuthz`

```go
type EntityAccountAuthzRevokedEvent struct {
	Id          string
	Signer      string
	AccountName string
	Granter     string
	Grantee     string
	MsgTypeUrl  string
}
```

The field's descriptions is as follows:

- `id` - a string containing the unique identifier of the entity.
- `signer` - a string containing the address of the signer.
- `accountName` - a string containing the name of the entity account the authz was granted before.
- `granter` - a string containing the address of the entity account the authz was granted before.
- `grantee` - a string containing the address of the grantee towards who the authz was granted before.
- `msgTypeUrl` - a string containing the message type url for the specific authz that was revoked.
