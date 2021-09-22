# State

The project module stores four lists of the following four types of data, as
well as module parameters (see the [Params spec page](06_params.md)):

1. [Project docs](#project-docs)
2. [Genesis account maps](#genesis-account-maps)
3. [Withdrawal info docs](#withdrawal-info-docs)
4. [Claims](#claims)

## Project Docs

```go
type ProjectDoc struct {
    TxHash     string
    ProjectDid string
    SenderDid  string
    PubKey     string
    Status     string
    Data       encoding_json.RawMessage
}
```

A project doc contains basic details about a project. Apart from the project DID
identifying the project doc, the project doc stores a transaction hash, the
project creator's DID, the project's public key (associated with the project's
DID), the project's current status, and other user-specified data associated to
the project.

**TxHash**: a value usually set by ixo-cellnode if it is the component that is
creating the project. If creating a project using `ixod`, the value of this
field is not very relevant.

**Status**: indicates the project's current status, which in turn dictates the
functionalities that are currently available on the project. The possible
statuses a project can be in are:

- the null status (`""`)
- `"CREATED"`
- `"PENDING"`
- `"FUNDED"`
- `"STARTED"`
- `"STOPPED"`
- `"PAIDOUT"`

Newly created projects are in the null status by default. The following are the
valid transitions between the states.
<pre>
- Null    -> CREATED
- CREATED -> PENDING
- PENDING -> CREATED or FUNDED
- FUNDED  -> STARTED
- STARTED -> STOPPED
- STOPPED -> PAIDOUT
</pre>

**Data**: a JSON object containing data relevant to the project. Despite being
mostly arbitrary, a project's `Data ("data")` field is in some cases expected to
follow concrete formats. Currently, there is only one such case, which is when
we want to specify payment templates to be used when charging project-related
fees.

The two (optional) project-related fees currently supported are:

- Oracle Fee (`OracleFee`)
- Fee for Service (`FeeForService`)

The following is an example where both an `OracleFee` and `FeeForService` are
specified:

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

The payment templates (e.g. `payment:template:oracle-fee-template-1`) are
expected to exist before the project is created (refer to payments module for
payment template creation).

For information around how these payment templates are used, refer to
the [Fees page](04_fees.md) of this module's spec.

## Genesis Account Maps

```go
type GenesisAccountMap struct {
    Map map[string]string
}
```

A genesis account map maps a project's account names to the accounts' addresses.
It keeps track of the addresses for the `InitiatingNodePayFees`, `IxoPayFees`,
and `IxoFees` accounts, as well as the accounts of the project's agents. For
more detail on project accounts, refer to the
[Entity Accounts page](05_entity_accounts.md) of this module's spec.

## Withdrawal Info Docs

```go
type WithdrawalInfoDocs struct {
    DocsList []WithdrawalInfoDoc
}
```

```go
type WithdrawalInfoDoc struct {
    ProjectDid   string
    RecipientDid string
    Amount       types.Coin
}
```

A withdrawal info doc stores details of any withdrawal performed on a project
and stores the project's DID, the recipient's DID, and the amount withdrawn.
`WithdrawalInfoDocs` stores a list of these withdrawal docs for a specific
project.

## Claims

```go
type Claims struct {
    ClaimsList []Claim
}
```

```go
type Claim struct {
    Id         string
    TemplateId string
    ClaimerDid string
    Status     string
}
```

A claim can be created on a project that is in status STARTED. Apart from an ID
identifying the claim, a claim also stores a template ID, the claimer's DID, and
the claim's current status. `Claims` stores a list of all claims for a specific
project.

**TemplateId**: indicates the ID of the claim template that the claim is based
on. More information about schema templates (and specifically claim templates)
in [this repository](https://github.com/ixofoundation/schema).

**Status**: indicates the claim's current status, which can be one of three:

- Pending (status `0`)
- Approved (status `1`)
- Rejected (status `2`)
