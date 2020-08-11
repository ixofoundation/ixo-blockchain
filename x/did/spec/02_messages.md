# Messages

In this section we describe the processing of the DID messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](01_state.md) section.

## MsgAddDid

This message is used to create a DID with an associated PubKey.

| **Field** | **Type**    | **Description** |
|:----------|:------------|:----------------|
| Did       | `did.DID`   | The DID being added 
| PubKey    | `publicKey` | The PubKey to be associated with the DID

```go
type MsgAddDid struct {
	Did    exported.Did
	PubKey string
}
```

This message is expected to fail if:
- the DID already exists

This message creates and stores the DID with its PubKey at appropriate indexes.

## MsgAddCredential 

The owner of a DID can add a credential to their DID using `MsgAddCredential`.

| **Field**     | **Type**                 | **Description** |
|:--------------|:-------------------------|:----------------|
| DidCredential | `exported.DidCredential` | The credential being added 

```go
type MsgAddCredential struct {
	DidCredential exported.DidCredential
}
```
