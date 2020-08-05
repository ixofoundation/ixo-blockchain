# Messages

In this section we describe the processing of the bonds messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgAddDid

Bonds can be created by any address using `MsgAddDid`.

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

- another bond with this token is already registered, the token is the staking token, or the token is not a valid denomination

This message creates and stores the `Bond` object at appropriate indexes. Note that the sanity rate and sanity margin percentage are only used in the case of the `swapper_function`, but no error is raised if these are set for other function types.

## MsgAddCredential 

The owner of a bond can edit some of the bond's parameters using `MsgAddCredential`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| CredType               | `credType`         | The bond to be edited |
| Issuer                 | `issuer`           | 
| Issued                 | `issued`           | 
| ClaimID                | `did`              | 
| ClaimKYCValidated      | `true`             | 


This message is expected to fail if:
- any editable field violates the restrictions set for the same field in `MsgCreateBond`
- all editable fields are `"[do-not-modify]"`
- signers list is not equal to the bond's signers list


