# State

## ProjectDoc

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

A project doc contains basic details about a project. Apart from the ProjectDid
identifying the project doc, the project doc stores a tx hash, the project creator's
DID, the project's public key, the project's current status, and other user-specified
data associated to the project.

**TxHash**: TODO

**Status**: indicates the project's current status, which in turn dictates the
functionalities that are currently available on the project. The possible
statuses a project can be in are:
- the null status
- CREATED
- PENDING
- FUNDED
- STARTED
- STOPPED
- PAIDOUT.

Newly created projects are in the null status by default.

**Data**: a JSON object containing data relevant to the project. Despite being
mostly arbitrary, a project's `Data ("data")` field is in some cases expected to
follow concrete formats. Currently, there is only one such case, which is when we
want to specify payment templates to be used when charging project-related fees.

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

## Claim

```go
type Claim struct {
    Id         string
    TemplateId string
    ClaimerDid string
    Status     string
}
```

A claim can be created on a project that is in status STARTED. Apart from an ID
identifying the claim, a claim also stores a template ID, the claimer's DID, and the
claim's current status.

**TemplateId**: TODO

**Status**: indicates the claim's current status, which can be one of three:
- Pending (status `0`)
- Approved (status `1`)
- Rejected (status `2`)
