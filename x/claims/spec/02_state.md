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
	EscrowAccount string
	Intents      CollectionIntentOptions
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
- `escrowAccount` - a string containing the escrow account address for this collection created at collection creation, used to transfer payments to escrow account for GUARANTEED payments through intents
- `intents` - a [CollectionIntentOptions](#collectionintentoptions) option for intents for this collection (allow, deny, required)

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
	CW20Payment           []CW20Payment
	IsOraclePayment       bool
}
```

The field's descriptions is as follows:

- `account` - a string containing the account address from which the payment will be made (must be an [EntityAccount](/x/entity/spec/02_state.md#entityaccount) of the `Entity` field for the Collection)
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on payment
- `contract_1155Payment` - a [Contract1155Payment](#contract1155payment), not allowed for `Evaluation` Payment
- `timeoutNs` - a duration containing the timeout after claim/evaluation to create authZ for payment, if 0 then immediate direct payment is made
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the CW20 tokens to be paid
- `isOraclePayment` - a boolean indicating if the payment is for oracle payments, meaning it will go through network fees split. Only allowed for APPROVAL payment types. If true and the payment contains CW20 payments, the claim will only be successful if an intent exists to ensure immediate CW20 payment split, since there is no WithdrawalAuthorization to manage the CW20 payment split for delayed payments.

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

- `address` - a string containing the smart contract address where the tokens can be transferred on (the cw1155 smart contract)
- `amount` - a integer indicating how many tokens must transferred for the payment
- `tokenId` - a string containing the `id` of the token on the cw1155 smart contract to transfer

### CW20Payment

A CW20Payment stores information about the payment to make if it is a CW20 token payment.

```go
type CW20Payment struct {
	Address string
	Amount  uint64
}
```

The field's descriptions is as follows:

- `address` - a string containing the contract address of the CW20 token
- `amount` - a uint64 containing the amount of CW20 tokens to transfer

### CW20Output

A CW20Output represents a CW20 token output for split payments.

```go
type CW20Output struct {
	Address         string
	ContractAddress string
	Amount          uint64
}
```

The field's descriptions is as follows:

- `address` - a string containing the address of the recipient
- `contractAddress` - a string containing the address of the CW20 contract
- `amount` - a uint64 containing the amount of the token to transfer

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
	UseIntent      bool
	Amount         github_com_cosmos_cosmos_sdk_types.Coins
	CW20Payment    []CW20Payment
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
- `useIntent` - a boolean indicating if this claim is using an intent. If true, then the amount and CW20 payment are ignored and overridden with intent amounts
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified by service agent for claim approval. If both amount and CW20 payment are empty, then collection default is used (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate amount wanted if no intent used)
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified by service agent for claim approval. If both amount and CW20 payment are empty, then collection default is used (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate cw20Payment wanted if no intent used)

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
	CW20Payment        []CW20Payment
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
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on `Approval` payment if it is a custom amount and not the preset `Approval` from the [Collection](#collection) (Note if intent was used, then the amount and CW20 payment are ignored and overridden with intent amounts)
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified by evaluator for claim approval (Note if intent was used, then the amount and CW20 payment are ignored and overridden with intent amounts)

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

### Intent

An Intent stores information about an intent to submit a claim to a collection.

```go
type Intent struct {
	Id            string
	CollectionId  string
	AgentDid      string
	AgentAddress  string
	Amount        github_com_cosmos_cosmos_sdk_types.Coins
	CW20Payment   []CW20Payment
	CreateDate    *time.Time
	ExpireDate    *time.Time
	Status        IntentStatus
}
```

The field's descriptions is as follows:

- `id` - a string containing the intent's identifier, it is incremented on chain for each intent created
- `collectionId` - a string containing the Collection `id` this intent belongs to
- `agentDid` - a string containing the DID of the agent creating the intent
- `agentAddress` - a string containing the account address of the agent creating the intent
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified for claim approval. If both amount and CW20 payment are empty, then collection default is used for intent amount
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified for claim approval. If both amount and CW20 payment are empty, then collection default is used for intent amount
- `createDate` - a timestamp of the date and time that the intent was created on-chain
- `expireDate` - a timestamp of the date and time that the intent will expire
- `status` - a [IntentStatus](#intentstatus) indicating the current status of the intent

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

### CollectionIntentOptions

Defines the options for intents for a [Collection](#collection), determining how intents are handled for claims in the collection.

```go
var CollectionIntentOptions_name = map[int32]string{
	0: "ALLOW",    // Allow: Intents can be made for claims, but claims can also be made without intents
	1: "DENY",     // Deny: Intents cannot be made for claims for the collection
	2: "REQUIRED", // Required: Claims cannot be made without an associated intent. An intent is mandatory before a claim can be submitted
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

### IntentStatus

Defines the status of an [Intent](#intent) indicating its current state.

```go
var IntentStatus_name = map[int32]string{
	0: "ACTIVE",    // Active: Intent is created and active, payments have been transferred to escrow if there is any
	1: "FULFILLED", // Fulfilled: Intent is fulfilled, was used to create a claim and funds will be released on claim APPROVAL, or funds will be reverted on claim REJECTION or DISPUTE
	2: "EXPIRED",   // Expired: Intent has expired, payments have been transferred back out of escrow
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
	3: "GUARANTEED",
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
	CW20Payment           []CW20Payment
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
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the CW20 tokens to be paid

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
	CollectionId      string
	AgentQuota        uint64
	MaxAmount         github_com_cosmos_cosmos_sdk_types.Coins
	MaxCW20Payment    []CW20Payment
	IntentDurationNs  time.Duration
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz)
- `maxAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the maximum amount allowed to be specified by service agent for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `maxCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payments allowed to be specified by service agent for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `intentDurationNs` - a duration for which the intent is active, after which it will expire (in nanoseconds)

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
	CollectionId          string
	ClaimIds              []string
	AgentQuota            uint64
	BeforeDate            *time.Time
	MaxCustomAmount       github_com_cosmos_cosmos_sdk_types.Coins
	MaxCustomCW20Payment  []CW20Payment
}
```

The field's descriptions is as follows:

- `claimIds` - a list of strings containing all the id's of the claimsthe grantee is allowed to evaluate, can be an empty list to allow any claim
- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz), note: it won't subtract one on evaluation if agent evaluates claim with status `invalidated`
- `beforeDate` - a timestamp of the date after which the grantee can't execute this authz anymore, a cut off date
- `MaxCustomAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount that indicates the maximum the evaluator is allowed to change the `APPROVED` payout to, since claims can be made for specific amount an evaluator is allowed to change the `APPROVED` payout amount.
- `MaxCustomCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payments allowed to be specified by evaluator for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.

### CreateClaimAuthorizationAuthorization

A CreateClaimAuthorizationAuthorization allows a grantee to create SubmitClaimAuthorization and EvaluateClaimAuthorization for specific collections(constraints).

```go
type CreateClaimAuthorizationAuthorization struct {
	Admin         string
	Constraints   []CreateClaimAuthorizationConstraints
}
```

The field's descriptions is as follows:

- `admin` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `constraints` - a list of [CreateClaimAuthorizationConstraints](#createclaimauthorizationconstraints)

### CreateClaimAuthorizationType

Defines the types of claim authorizations that can be created.

```go
var CreateClaimAuthorizationType_name = map[int32]string{
	0: "ALL",      // both submit and evaluate
	1: "SUBMIT",   // submit only
	2: "EVALUATE", // evaluate only
}
```

### CreateClaimAuthorizationConstraints

Constraints for creating claim authorizations through a [CreateClaimAuthorizationAuthorization](#createclaimauthorizationauthorization).

```go
type CreateClaimAuthorizationConstraints struct {
	MaxAuthorizations    uint64
	MaxAgentQuota        uint64
	MaxAmount            github_com_cosmos_cosmos_sdk_types.Coins
	MaxCW20Payment       []CW20Payment
	Expiration           *time.Time
	CollectionIds        []string
	AllowedAuthTypes     CreateClaimAuthorizationType
	MaxIntentDurationNs  time.Duration
}
```

The field's descriptions is as follows:

- `maxAuthorizations` - a integer containing the maximum number of authorizations that can be created through this meta-authorization. 0 means no quota.
- `maxAgentQuota` - a integer containing the maximum quota that can be set in created authorizations. 0 means no maximum quota per authorization.
- `maxAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the maximum amount that can be set in created authorizations. If empty then any custom amount is allowed in the created authorizations. Explicitly set to 0 to disallow any custom amount in the created authorizations.
- `maxCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payment that can be set in created authorizations. If empty then any CW20 payment is allowed in the created authorizations. Explicitly set to 0 to disallow any CW20 payment in the created authorizations.
- `expiration` - a timestamp of the expiration of this meta-authorization(specific constraint). If not set then no expiration.
- `collectionIds` - a list of strings containing the Collection IDs the grantee can create authorizations for. If empty then all collections for the admin are allowed.
- `allowedAuthTypes` - a [CreateClaimAuthorizationType](#createclaimauthorizationtype) indicating the types of authorizations the grantee can create (submit, evaluate, or all/both).
- `maxIntentDurationNs` - a duration containing the maximum intent duration for the authorization allowed (for submit).
