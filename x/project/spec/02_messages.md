# Messages

In this section we describe the processing of the project messages and the
corresponding updates to the state. All created/modified state objects specified
by each message are defined within the [state](01_state.md) section.

## MsgCreateProject

This message creates and stores a new project doc with arbitrary `Data` at
appropriate indexes. Refer to [01_state.md](./01_state.md) for information
about project docs.

| **Field**  | **Type**          | **Description** |
|:-----------|:------------------|:----------------|
| TxHash     | `string`          | Hash of the project request
| SenderDid  | `did.Did`         | Sender account DID
| ProjectDid | `did.Did`         | Sender's Project DID
| PubKey     | `string`          | PubKey of ixo account
| Data       | `json.RawMessage` | What the data is passing

```go
type MsgCreateProject struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    PubKey     string
    Data       json.RawMessage
}
```

This message is expected to fail if:
- A project doc with DID ProjectDid already exists
- SenderDid is empty or invalid
- PubKey is empty or does not match ProjectDid
- Data is unmarshallable to `map[string]json.RawMessage`
- Project fees in Data are invalid

## MsgUpdateProjectStatus

This message updates a project's current status.

| **Field**  | **Type**                 | **Description** |
|:-----------|:-------------------------|:----------------|
| TxHash     | `string`                 | Hash of the project request
| SenderDid  | `did.Did`                | Sender account DID
| ProjectDid | `did.Did`                | Sender's Project DID
| Data       | `UpdateProjectStatusDoc` | Updated data to this project

```go
type MsgUpdateProjectStatus struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       UpdateProjectStatusDoc
}
```

An `UpdateProjectStatusDoc` contains the new project status as well as TODO.

```go
type UpdateProjectStatusDoc struct {
    Status          string
    EthFundingTxnId string
}
```

This message is expected to fail if:
- Project doc having DID ProjectDID does not exist
- SenderDid is empty or invalid
- ProjectDid is empty or invalid
- The status change constitutes an invalid status progression
- The new status is FUNDED and the project has not yet reached minimum funding
- The new status is PAIDOUT and paying out fees fails

## MsgUpdateProjectDoc

| **Field**  | **Type**          | **Description** |
|:-----------|:------------------|:----------------|
| TxHash     | `string`          | Hash of the project request
| SenderDid  | `did.Did`         | Sender account DID
| ProjectDid | `did.Did`         | Sender's Project DID
| Data       | `json.RawMessage` | What the data is passing

```go
type MsgUpdateProjectDoc struct {
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  string          `json:"senderDid" yaml:"senderDid"`
	ProjectDid string          `json:"projectDid" yaml:"projectDid"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}
```

This message is expected to fail if:
- senderDid is wrong
- projectDid is wrong
- data is unmarshallable to `map[string]json.RawMessage`

This message stores the updated `MsgUpdateProjectDoc` object.

## MsgCreateAgent

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID
| ProjectDid | `did.Did`        | Sender's Project DID
| Data       | `CreateAgentDoc` | AgentDoc data

```go
type MsgCreateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  string         `json:"senderDid" yaml:"senderDid"`
	ProjectDid string         `json:"projectDid" yaml:"projectDid"`
	Data       CreateAgentDoc `json:"data" yaml:"data"`
}
```

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
	SenderDid  string         `json:"senderDid" yaml:"senderDid"`
	ProjectDid string         `json:"projectDid" yaml:"projectDid"`
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
- claim ID (in Data) is wrong
- claim template ID (in Data) is wrong

```go
type MsgCreateClaim struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  string         `json:"senderDid" yaml:"senderDid"`
	ProjectDid string         `json:"projectDid" yaml:"projectDid"`
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
	SenderDid  string              `json:"senderDid" yaml:"senderDid"`
	ProjectDid string              `json:"projectDid" yaml:"projectDid"`
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
	SenderDid string          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}
```