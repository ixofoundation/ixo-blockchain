# Messages

In this section we describe the processing of the DID messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](01_state.md) section.

## MsgAddDid

This message is used to create a DidDoc with a given DID and an associated PubKey.

| **Field** | **Type**    | **Description** |
|:----------|:------------|:----------------|
| Did       | `Did`       | The DID of the DidDoc being added (e.g. `did:ixo:4XJLBfGtWSGKSz4BeRxdun`)
| PubKey    | `string`    | The PubKey to be associated with the DID (e.g. `2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt`)

```go
type MsgAddDid struct {
    Did    string
    PubKey string
}
```

This message is expected to fail if:
- the DID already exists

This message creates and stores the DID with its PubKey at appropriate indexes.

## MsgAddCredential 

The owner of a DidDoc can add a credential to their DidDoc using `MsgAddCredential`.

| **Field**     | **Type**        | **Description** |
|:--------------|:----------------|:----------------|
| DidCredential | `DidCredential` | Details of the credential being added, including which DID the credential is for and by which signer the credential is issued

```go
type MsgAddCredential struct {
    DidCredential DidCredential
}
```
