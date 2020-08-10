# Messages

In this section we describe the processing of the project messages. 

## MsgCreateProject

Project can be created by any address using `MsgCreateProject`.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | Hash of the project request |
| SenderDid              | `did.Did`          | Sender account DID          |
| ProjectDid             | `did.Did`          | Sender's Project DID        |
| PubKey                 | `string`           | Pubkey of IXO account       |
| Data                   | `json.RawMessage`  | What the data is passing.   |


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

- senderDid is wrong  as input
- PubKey is wrong as input

This message creates and stores the `Project` object at appropriate indexes. 

## MsgUpdateProjectStatus


| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | Hash of the project request  |
| SenderDid              | `did.Did`          | Sender account DID   
| ProjectDid             | `did.Did`          | Sender's Project DID
| Data                   | `UpdateProjectStatusDoc`  |  Updated data to this project

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

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | Hash of the project request |
| SenderDid              | `did.Did`          | Sender account DID 
| ProjectDid             | `did.Did`          | Sender's Project DID
| Data                   | `UpdateAgentDoc`   | AgentDoc data

 

```go
type MsgUpdateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       UpdateAgentDoc `json:"data" yaml:"data"`
}
```

This message update the project  to new agent.

## MsgCreateClaim

This will create a claim for the project.

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | Hash of the project request |
| SenderDid              | `did.Did`          | Sender account DID 
| ProjectDid             | `did.Did`          | Sender's Project DID
| Data                   | `CreateClaimDoc`   |  Claim Doc for the project

This message is expected to fail if:
- sender-did is wrong
- project_did is wrong

```go
type MsgCreateClaim struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateClaimDoc `json:"data" yaml:"data"`
}
```


## MsgCreateEvaluation


| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| TxHash                 | `string`           | Hash of the project request  |
| SenderDid              | `did.Did`          | Sender account DID
| ProjectDid             | `did.Did`          | Sender's Project DID
| Data                   | `CreateEvaluationDoc`  | Evalution Doc for the project

This message is expected to fail if:
- sender-did is wrong
- project_did is wrong


```go
type MsgCreateEvaluation struct {
	TxHash     string              `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did             `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did             `json:"projectDid" yaml:"projectDid"`
	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
}
```


## MsgWithdrawFunds

| **Field**              | **Type**           | **Description**                                                                                               |
|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------------------|
| SenderDid              | `did.Did`          |  Hash of the project request
| Data                   | `WithdrawFundsDoc`  | Amount to which data is transferring


```go
type MsgWithdrawFunds struct {
	SenderDid did.Did          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}
```


