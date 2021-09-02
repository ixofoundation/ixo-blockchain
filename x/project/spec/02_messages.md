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
| SenderDid  | `did.Did`         | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`         | New project DID (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| PubKey     | `string`          | PubKey of project's ixo account (e.g. `FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW`)
| Data       | `json.RawMessage` | Data relevant to the project

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
| SenderDid  | `did.Did`                | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`                | Project DID whose status is to be changed (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `UpdateProjectStatusDoc` | Updated status data to this project

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
- Project doc having DID ProjectDid does not exist
- SenderDid is empty or invalid
- ProjectDid is empty or invalid
- The status change constitutes an invalid status progression
- The new status is FUNDED and the project has not yet reached minimum funding
- The new status is PAIDOUT and paying out fees fails

## MsgUpdateProjectDoc

This message updates a project's Data field.

| **Field**  | **Type**          | **Description** |
|:-----------|:------------------|:----------------|
| TxHash     | `string`          | Hash of the project request
| SenderDid  | `did.Did`         | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`         | Project DID whose Data field is to be updated (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `json.RawMessage` | Updated data relevant to the project

```go
type MsgUpdateProjectDoc struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       json.RawMessage
}
```

This message is expected to fail if:
- SenderDid is empty or invalid
- SenderDid does not match project creator DID
- ProjectDid is empty or invalid
- Project doc having DID ProjectDid does not exist
- Project is in status STARTED, STOPPED, or PAIDOUT
- Data is unmarshallable to `map[string]json.RawMessage`
- Project fees in updated Data are invalid

## MsgCreateAgent

This message creates an agent on a specified project. An agent is any account
interacting with the project. Each agent has a project account.

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`        | Project DID to which agent will be added (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `CreateAgentDoc` | AgentDoc data

```go
type MsgCreateAgent struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       CreateAgentDoc
}
```

```go
type CreateAgentDoc struct {
    AgentDid string
    Role     string
}
```

The role of an agent must be one of `SA` (can list claims), `EA` (can evaluate
claims), or `IA`. TODO

## MsgUpdateAgent

This message updates the status of an agent on a project.

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`        | Project DID whose agent status will be updated (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `UpdateAgentDoc` | AgentDoc data

```go
type MsgUpdateAgent struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       UpdateAgentDoc
}
```

```go
type UpdateAgentDoc struct {
    Did    string
    Status string
    Role   string
}
```

Similarly to `MsgCreateAgent`, the role of an agent must be one of `SA`, `EA`,
or `IA`. The status must be one of `0` (Pending), `1` (Approved), or `2` (Revoked).

## MsgCreateClaim

This message creates a claim for a specified project. Refer to
[01_state.md](./01_state.md) for information about claims.

| **Field**  | **Type**         | **Description** |
|:-----------|:-----------------|:----------------|
| TxHash     | `string`         | Hash of the project request
| SenderDid  | `did.Did`        | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`        | Project DID on which a claim is to be created (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `CreateClaimDoc` | Claim Doc for the project

```go
type MsgCreateClaim struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       CreateClaimDoc
}
```

A `CreateClaimDoc` contains an ID uniquely identifying the claim, and TODO. Upon
creating a claim, its default status is `0` i.e. Pending.

```go
type CreateClaimDoc struct {
    ClaimId         string
    ClaimTemplateId string
}
```

This message is expected to fail if:
- Project doc having DID ProjectDid does not exist
- Project is not in status STARTED
- SenderDid is empty or invalid
- ProjectDid is empty
- ClaimId (in Data) already exists
- ClaimTemplateId (in Data) is empty

## MsgCreateEvaluation

This message creates an evaluation for a specified claim on a specified project.

| **Field**  | **Type**              | **Description** |
|:-----------|:----------------------|:----------------|
| TxHash     | `string`              | Hash of the project request
| SenderDid  | `did.Did`             | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| ProjectDid | `did.Did`             | Project DID (e.g. `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`)
| Data       | `CreateEvaluationDoc` | Evalution Doc for the project

```go
type MsgCreateEvaluation struct {
    TxHash     string
    SenderDid  string
    ProjectDid string
    Data       CreateEvaluationDoc
}
```

A `CreateEvaluationDoc` contains the claim ID of the claim being evaluated, and a
new status indicating whether the claim is accepted (status `1`) or rejected
(status `2`).

```go
type CreateEvaluationDoc struct {
    ClaimId string
    Status  string
}
```

This message is expected to fail if:
- Project doc having DID ProjectDid does not exist
- Project is not in status STARTED
- Claim with ClaimId (in Data) does not exist
- Claim with ClaimId (in Data) is not in status Pending (status `0`)
- Oracle fee is present in project fees map, and ixo address, node (relayer)
  address, or sender (oracle) address cannot be found, or if payment cannot be
  processed
- SenderDid is empty or invalid
- ProjectDid is empty or invalid

## MsgWithdrawFunds

This message allows project agents to withdraw their funds from the project.

| **Field** | **Type**           | **Description** |
|:----------|:-------------------|:----------------|
| SenderDid | `did.Did`          | Sender account DID (e.g. `did:ixo:U4tSpzzv91HHqWW1YmFkHJ`)
| Data      | `WithdrawFundsDoc` | Details about the funds to be withdrawn

```go
type MsgWithdrawFunds struct {
    SenderDid string
    Data      WithdrawFundsDoc
}
```

The `WithdrawFundsDoc` specifies the project DID from which funds are to be withdrawn,
the recipient of the funds, the amount of funds to be withdrawn, and whether the
withdrawal is a refund to be sent to the project creator.

```go
type WithdrawFundsDoc struct {
    ProjectDid   string
    RecipientDid string
    Amount       sdk.Int
    IsRefund     bool
}
```

This message is expected to fail if:
- Project doc having DID ProjectDid does not exist
- Project is not in status PAIDOUT
- IsRefund is set to true and RecipientDid is not the project creator
- Project account has insufficient funds
- Sending funds to recipient fails
- SenderDid is empty or invalid
- SenderDid does not match RecipientDid
- RecipientDid (in Data) is empty or invalid
- Amount (in Data) is not positive
