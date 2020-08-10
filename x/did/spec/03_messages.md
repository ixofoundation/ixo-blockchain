# Messages

In this section we describe the processing of the did messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgAddDid

Dids can be created by any address using `MsgAddDid`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-----------------  |:--------------------------------------------------------------------------------------------------------------|
| Did                    | `did.DID`          | 
| PubKey                 | `publicKey`        | 

```go
type MsgAddDid struct {
	Did    exported.Did `json:"did" yaml:"did"`
	PubKey string       `json:"pubKey" yaml:"pubKey"`
}
```

This message is expected to fail if:

- if the same DID created.

This message creates and stores the `did` object at appropriate indexes. 

## MsgAddCredential 

The owner of a did can edit some of the did's parameters using `MsgAddCredential`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CredType               | `string`           | Type of the credential |
| Issuer                 | `string`           | Who is the issuer
| Issued                 | `string`           | What is been issued
| ClaimID                | `string`           | ClaimID supporting to the credential
| ClaimKYCValidated      | `bool`             | Validation of the ClaimID


