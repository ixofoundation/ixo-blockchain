# Messages

In this section we describe the processing of the project messages. 

## MsgCreateProject

This message creates a project with arbitrary `Data`.

| **Field**  | **Type**          | **Description** |
|:-----------|:------------------|:----------------|
| TxHash     | `string`          | Hash of the project request
| SenderDid  | `did.Did`         | Sender account DID
| ProjectDid | `did.Did`         | Sender's Project DID
| PubKey     | `string`          | PubKey of ixo account
| Data       | `json.RawMessage` | What the data is passing

```go
type MsgCreateProject struct {
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}
```

This message is expected to fail if:
- senderDid is incorrect
- PubKey is incorrect

This message creates and stores the `Project` object at appropriate indexes.

### Non-Arbitrary Project Data

Despite being mostly arbitrary, a project's `Data ("data")` field is in some cases expected to follow concrete formats. Currently, there is only one such case, which is when we want to specify payment templates to be used when charging project-related fees.

The two (optional) project-related fees currently supported are:
- Oracle Fee (`OracleFee`)
- Fee for Service (`FeeForService`)

The following is an example where both an `OracleFee` and `FeeForService` are specified:
```json
"data": {
    ...
    "fees": {
        "@context": "...",
        "items": [
            {
                "@type": "OracleFee",
                "id":"payment:template:oracle-fee-template-1"
            },
            {
                "@type": "FeeForService", 
                "id":"payment:template:fee-for-service-template-1"
            }
        ]
    }
    ...
}
```

If we do not specify fees, a blank `items` array is required:
```json
"data": {
    ...
    "fees": {
        "@context": "...",
        "items": []
    }
    ...
}
```

The payment templates (e.g. `payment:template:oracle-fee-template-1`) are expected to exist before the project is created (refer to payments module for payment template creation).

For information around how these payment templates are used, refer to the [Fees page](04_fees.md) of this module's spec.

## MsgUpdateProjectStatus

| **Field**  | **Type**                 | **Description** |
|:-----------|:-------------------------|:----------------|
| TxHash     | `string`                 | Hash of the project request
| SenderDid  | `did.Did`                | Sender account DID
| ProjectDid | `did.Did`                | Sender's Project DID
| Data       | `UpdateProjectStatusDoc` |  Updated data to this project

This message is expected to fail if:
- senderDid is wrong
- projectDid is wrong

```go
type MsgUpdateProjectStatus struct {
	TxHash     string                 `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                `json:"projectDid" yaml:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data" yaml:"data"`
}
```

This message stores the updated `MsgUpdateProjectStatus` object.

## MsgUpdateAgent

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID
| ProjectDid | `did.Did`        | Sender's Project DID
| Data       | `UpdateAgentDoc` | AgentDoc data

```go
type MsgUpdateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       UpdateAgentDoc `json:"data" yaml:"data"`
}
```

## MsgCreateClaim

This will create a claim for the specified project.

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID
| ProjectDid | `did.Did`        | Sender's Project DID
| Data       | `CreateClaimDoc` |  Claim Doc for the project

This message is expected to fail if:
- senderDid is wrong
- projectDid is wrong

```go
type MsgCreateClaim struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateClaimDoc `json:"data" yaml:"data"`
}
```

## MsgCreateEvaluation

This will create a claim evaluation for the specified project.

| **Field**  | **Type**              | **Description** |
|:-----------|:----------------------|:----------------|
| TxHash     | `string`              | Hash of the project request
| SenderDid  | `did.Did`             | Sender account DID
| ProjectDid | `did.Did`             | Sender's Project DID
| Data       | `CreateEvaluationDoc` | Evalution Doc for the project

This message is expected to fail if:
- senderDid is wrong
- projectDid is wrong

```go
type MsgCreateEvaluation struct {
	TxHash     string              `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did             `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did             `json:"projectDid" yaml:"projectDid"`
	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
}
```

## MsgWithdrawFunds

This is used by project agents to withdraw their funds from the project.

| **Field** | **Type**           | **Description** |
|:----------|:-------------------|:----------------|
| SenderDid | `did.Did`          |  Hash of the project request
| Data      | `WithdrawFundsDoc` | Amount to which data is transferring

```go
type MsgWithdrawFunds struct {
	SenderDid did.Did          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}
```
