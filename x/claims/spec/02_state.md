# State

## Collections

A Collections is stored in the state and is accessed by the identity of the collection that is fetched and incremented onchain using the `Param`: `CollectionSequence`.

- Collections: `0x01 | collectionId -> ProtocolBuffer(Collection)`

## Claims

A Claim is stored in the state and is accessed by the identity of the ClaimId(user provided).

- Claims: `0x02 | claimId -> ProtocolBuffer(Claim)`

## Disputes

A Dispute is stored in the state and is accessed by the SubjectId of the dispute(user provided).

- Disputes: `0x03 | disputeSubjectId(DID) -> ProtocolBuffer(Dispute)`

# Types

### Collection

```go
type Collection struct {
	Id           string
	Entity       string
	Admin        string
	Protocol     string
	StartDate    *time.Time
	EndDate      *time.Time
	Quota        uint64
	Count        uint64
	Evaluated    uint64
	Approved     uint64
	Rejected     uint64
	Disputed     uint64
	Invalidated  uint64
	State        CollectionState
	Payments     *Payments
}
```

The field's descriptions is as follows:

- `id` - a string containing the collections identifier, it is incremented on chain for the collection of claims
- `entity` - a string containing the DID of the entity for which the claims are being created
- `admin` - a string containing the account address that will authorize or revoke agents and payments (the granter). It is the `Entity`'s [EntityAccount](/x/entity/spec/02_state.md#entityaccount) named `admin`
- `protocol` - a string containing the DID of the claim protocol
- `startDate` - a timestamp of the start date for the collection, after which claims may be submitted
- `endDate` - a timestamp of the end date for the collection, after which no more claims may be submitted (no endDate is allowed)
- `quota` - a integer containing the maximum number of claims that may be submitted, 0 is unlimited
- `count` - a integer containing the number of claims already submitted (internally calculated)
- `evaluated` - a integer containing the number of claims that have been evaluated (internally calculated)
- `approved` - a integer containing the number of claims that have been evaluated and approved (internally calculated)
- `rejected` - a integer containing the number of claims that have been evaluated and rejected (internally calculated)
- `disputed` - a integer containing the number of claims that have disputed status (internally calculated)
- `invalidated` - a integer containing the number of claims that have invalidated status (internally calculated)
- `state` - a [CollectionState](#collectionstate)
- `payments` - a [Payments](#payments)

### Payments

A Payments stores [Payment](#payment) for the claim submission, evaluation, approval, or rejection payments made towards the collection

```go
type Payments struct {
	Submission *Payment
	Evaluation *Payment
	Approval   *Payment
	Rejection  *Payment
}
```

The field's descriptions is as follows:

- `submission` - a [Payment](#payment)
- `evaluation` - a [Payment](#payment)
- `approval` - a [Payment](#payment)
- `rejection` - a [Payment](#payment)

### Payment

A Payment stores information about the amount paid for claim submission, evaluation, approval, or rejection

```go
type Payment struct {
	Account               string
	Amount                github_com_cosmos_cosmos_sdk_types.Coins
	Contract_1155Payment  *Contract1155Payment
	TimeoutNs             time.Duration
}
```

The field's descriptions is as follows:

- `account` - a string containing the account address from which the payment will be made (ideally a [EntityAccount](/x/entity/spec/02_state.md#entityaccount))
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on payment
- `contract_1155Payment` - a [Contract1155Payment](#contract1155payment)
- `timeoutNs` - a duration containing the timeout after claim/evaluation to create authZ for payment, if 0 then immidiate direct payment is made

### Contract1155Payment

A Contract1155Payment stores information about the payment to make if it is a cw1155 tokens payment.

```go
type Contract1155Payment struct {
	Address string
	TokenId string
	Amount  uint32
}
```

The field's descriptions is as follows:

- `address` - a string containing the smart contract address where the tokens can be transfered on (the cw1155 smart contract)
- `amount` - a integer indicating how many tokens must transfered for the payment
- `tokenId` - a string containing the `id` of the token on the cw1155 smart contract to transfer

### Claim

A Claim stores information about a claim that was made towards a [Collection](#collection)

```go
type Claim struct {
	CollectionId   string
	AgentDid       string
	AgentAddress   string
	SubmissionDate *time.Time
	ClaimId        string
	Evaluation     *Evaluation
	PaymentsStatus *ClaimPayments
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` this claim belongs
- `agentAddress` - a string containing the account address that submitted the claim
- `agentDid` - a string containing the Did of the agent that submitted the claim
- `submissionDate` - the timestamp of the date and time that the claim was submitted on-chain
- `claimId` - a string containing the unique identifier of the claim (eg. cid hash of file is good identifier)
- `evaluation` - a [Evaluation](#evaluation)
- `paymentsStatus` - a [ClaimPayments](#claimpayments)

### ClaimPayments

A ClaimPayments stores an enum for the status for the claim submission, evaluation, approval, or rejection payments

```go
type ClaimPayments struct {
	Submission PaymentStatus
	Evaluation PaymentStatus
	Approval   PaymentStatus
	Rejection  PaymentStatus
}
```

The field's descriptions is as follows:

- `submission` - a [PaymentStatus](#paymentstatus)
- `evaluation` - a [PaymentStatus](#paymentstatus)
- `approval` - a [PaymentStatus](#paymentstatus)
- `rejection` - a [PaymentStatus](#paymentstatus)

### Evaluation

A Evaluation stores information concerning the evaluation of a [Claim](#claim) made

```go
type Evaluation struct {
	ClaimId            string
	CollectionId       string
	Oracle             string
	AgentDid           string
	AgentAddress       string
	Status             EvaluationStatus
	Reason             uint32
	VerificationProof  string
	EvaluationDate     *time.Time
	Amount             github_com_cosmos_cosmos_sdk_types.Coins
}
```

The field's descriptions is as follows:

- `claimId` - a string containing the `id` of the claim the evaluation is for
- `collectionId` - a string containing the Collection `id` the claim this evaluation is for belongs to
- `oracle` - a string containing the DID of the Oracle entity that evaluates the claim
- `agentAddress` - a string containing the account address that submitted the evaluation
- `agentDid` - a string containing the Did of the agent that submitted the evaluation
- `status` - a [EvaluationStatus](#evaluationstatus)
- `reason` - a integer for why the evaluation result was given (codes defined by evaluator)
- `verificationProof` - a string containing the proof to verify the linked resource (eg. the cid of the evaluation Verfiable Credential)
- `evaluationDate` - the timestamp of the date and time that the claim evaluation was submitted on-chain
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on `Approval` payment if it is a custom amount and not the preset `Approval` from the [Collection](#collection)

### Dispute

A Dispute stores information concerning the dispute made towards a [Claim](#claim)

```go
type Dispute struct {
	SubjectId  string
	Type       int32
	Data       *DisputeData
}
```

The field's descriptions is as follows:

- `subjectId` - a string containing the `id` of the claim the dispute is for. A unique field
- `type` - a integer interpreted by the client
- `data` - a [DisputeData](#disputedata)

### DisputeData

A DisputeData stores information concerning the data for a dispute made towards a [Claim](#claim)

```go
type DisputeData struct {
	Uri       string
	Type      string
	Proof     string
	Encrypted bool
}
```

The field's descriptions is as follows:

- `uri` - a string representing the endpoint of the data linked resource
- `type ` - a string representing the [MIME](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types) type of the data linked resource
- `proof ` - a string representing the proof to verify the data linked resource
- `encrypted ` - a boolean value for whether this data linked resource is encrypted or not

## Enums

### CollectionState

Defines the `state` of a [Collection](#collection) denoting whether claims is allowed to be submitted or not. Already submitted claims is still allowed to be evaluated regardless of the `state`

```go
var CollectionState_name = map[int32]string{
	0: "OPEN",
	1: "PAUSED",
	2: "CLOSED",
}
```

### EvaluationStatus

Defines the `status` of a [Evaluation](#evaluation) indicating the status and result of the evaluation.

```go
var EvaluationStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "REJECTED",
	3: "DISPUTED",
  4: "INVALIDATED"
}
```

### PaymentType

Defines the type of [Payment](#payment) used to keep track for payment withdrawals to update `ClaimPayments` accordingly

```go
var PaymentType_name = map[int32]string{
	0: "SUBMISSION",
	1: "APPROVAL",
	2: "EVALUATION",
	3: "REJECTION",
}
```

### PaymentStatus

Defines the status of the payment types for `ClaimPayments`

```go
var PaymentStatus_name = map[int32]string{
	0: "NO_PAYMENT",
	1: "PROMISED",
	2: "AUTHORIZED",
	3: "GAURANTEED",
	4: "PAID",
	5: "FAILED",
	6: "DISPUTED",
}
```

## Authz Types

### WithdrawPaymentAuthorization

A WithdrawPaymentAuthorization is an authz authorization that can be granted to allow the grantee to make a withdrawal payout to receive his payment for claims submitted, evaluated or that was approved

```go
type WithdrawPaymentAuthorization struct {
	Admin         string
	Constraints   []*WithdrawPaymentConstraints
}
```

The field's descriptions is as follows:

- `admin` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `constraints` - a list of [WithdrawPaymentConstraints](#withdrawpaymentconstraints)

### WithdrawPaymentConstraints

A WithdrawPaymentConstraints stores information about authorization given to make a withdrawal payment through a [WithdrawPaymentAuthorization](#withdrawpaymentauthorization)

```go
type WithdrawPaymentConstraints struct {
	ClaimId               string
	Inputs                []github_com_cosmos_cosmos_sdk_x_bank_types.Input
	Outputs               []github_com_cosmos_cosmos_sdk_x_bank_types.Output
	PaymentType           PaymentType
	Contract_1155Payment  *Contract1155Payment
	ToAddress             string
	FromAddress           string
	ReleaseDate           *time.Time
}
```

The field's descriptions is as follows:

- `claimId` - a string containing the `id` of the claim the withdrawal is for
- `inputs` - a list of cosmos defined `Input` to pass to the the multisend tx to run to withdraw payment
- `outputs` - a list of cosmos defined `Output` to pass to the the multisend tx to run to withdraw payment
- `paymentType` - a [PaymentType](02_state.md#paymenttype)
- `contract_1155Payment` - a [Contract1155Payment](02_state.md#contract1155payment)
- `toAddress` - a string containing the account address to make the payment to
- `fromAddress` - a string containing the account address to make the payment from
- `releaseDate` - a timestamp of the date that grantee can execute authorization to make the withdrawal payment, calculated from created date plus the timeout on [Collection](02_state.md#collection) `Payments`
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field

### SubmitClaimAuthorization

A SubmitClaimAuthorization is an authz authorization that can be granted to allow the grantee to submit claims for a specified collection

```go
type SubmitClaimAuthorization struct {
	Admin         string
	Constraints   []*SubmitClaimConstraints
}
```

The field's descriptions is as follows:

- `admin` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `constraints` - a list of [SubmitClaimConstraints](#submitclaimconstraints)

### SubmitClaimConstraints

A SubmitClaimConstraints stores information about authorization given to submit claims through a [SubmitClaimAuthorization](#submitclaimauthorization)

```go
type SubmitClaimConstraints struct {
	CollectionId   string
	AgentQuota     uint64
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz)

### EvaluateClaimAuthorization

A EvaluateClaimAuthorization is an authz authorization that can be granted to allow the grantee to evaluate claims for a specified collection

```go
type EvaluateClaimAuthorization struct {
	Admin         string
	Constraints   []*EvaluateClaimConstraints
}
```

The field's descriptions is as follows:

- `admin` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `constraints` - a list of [EvaluateClaimConstraints](#evaluateclaimconstraints)

### EvaluateClaimConstraints

A EvaluateClaimConstraints stores information about authorization given to evaluate claims through a [EvaluateClaimAuthorization](#evaluateclaimauthorization)

```go
type EvaluateClaimConstraints struct {
	CollectionId     string
	ClaimIds         []string
	AgentQuota       uint64
	BeforeDate       *time.Time
	MaxCustomAmount  github_com_cosmos_cosmos_sdk_types.Coins
}
```

The field's descriptions is as follows:

- `claimIds` - a list of strings containing all the id's of the claimsthe grantee is allowed to evaluate, can be an empty list to allow any claim
- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz), note: it won't subtract one on evaluation if agent evaluates claim with status `invalidated`
- `beforeDate` - a timestamp of the date after which the grantee can't execute this authz anymore, a cut off date
- `MaxCustomAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount that indicates the maximum the evaluator is allowed to change the `APPROVED` payout to, since claims can be made for specific amount an evaluator is allowed to change the `APPROVED` payout amount.
