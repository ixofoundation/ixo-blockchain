# Messages

In this section we describe the processing of the claims messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state.md) section.

## MsgCreateCollection

A `MsgCreateCollection` creates and stores a new Collection which defines protocols, quotas and payments for claim submissions and evaluations

```go
type MsgCreateCollection struct {
	Entity string
	Signer string
	Protocol string
	StartDate *time.Time
	EndDate *time.Time
	Quota uint64
	State CollectionState
	Payments *Payments
	Intents CollectionIntentOptions
}
```

The field's descriptions is as follows:

- `entity` - a string containing the DID of the entity for which the claims are being created
- `signer` - a string containing the account address of the private key signing the transaction
- `protocol` - a string containing the DID of the claim protocol
- `startDate` - a timestamp of the start date for the collection, after which claims may be submitted
- `endDate` - a timestamp of the end date for the collection, after which no more claims may be submitted (no endDate is allowed)
- `quota` - a integer containing the maximum number of claims that may be submitted, 0 is unlimited
- `state` - a [CollectionState](02_state.md#collectionstate)
- `payments` - a [Payments](02_state.md#payments)
- `intents` - a [CollectionIntentOptions](02_state.md#collectionintentoptions) option for intents for this collection (allow, deny, required)

## MsgUpdateCollectionState

A `MsgUpdateCollectionState` updates a Collection's `state` field.

```go
type MsgUpdateCollectionState struct {
	CollectionId string
	State CollectionState
	AdminAddress string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to make the update to
- `state` - a [CollectionState](02_state.md#collectionstate)
- `adminAddress` - a string containing the account address of the private key signing the transaction, must be same as [Collection](02_state.md#collection) admin field

## MsgUpdateCollectionDates

A `MsgUpdateCollectionDates` updates a Collection's `startDate` and `endDate` fields.

```go
type MsgUpdateCollectionDates struct {
	CollectionId string
  StartDate *time.Time
  EndDate *time.Time
	AdminAddress string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to make the update to
- `startDate` - a timestamp of the start date for the collection, after which claims may be submitted
- `endDate` - a timestamp of the end date for the collection, after which no more claims may be submitted (no endDate is allowed)
- `adminAddress` - a string containing the account address of the private key signing the transaction, must be same as [Collection](02_state.md#collection) admin field

## MsgUpdateCollectionPayments

A `MsgUpdateCollectionPayments` updates a Collection's `state` field.

```go
type MsgUpdateCollectionPayments struct {
	CollectionId string
	Payments *Payments
	AdminAddress string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to make the update to
- `payments` - a [Payments](02_state.md#payments)
- `adminAddress` - a string containing the account address of the private key signing the transaction, must be same as [Collection](02_state.md#collection) admin field

## MsgUpdateCollectionIntents

A `MsgUpdateCollectionIntents` updates a Collection's `intents` field.

```go
type MsgUpdateCollectionIntents struct {
	CollectionId string
	Intents CollectionIntentOptions
	AdminAddress string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to make the update to
- `intents` - a [CollectionIntentOptions](02_state.md#collectionintentoptions) option for intents for this collection (allow, deny, required)
- `adminAddress` - a string containing the account address of the private key signing the transaction, must be same as [Collection](02_state.md#collection) admin field

## MsgSubmitClaim

A `MsgSubmitClaim` creates and stores a new Claim made towards a `Collection`. On Submission of claim `SUBMISSION` payments will be made if there is any defined in the [Collection](02_state.md#collection) `Payments`.

```go
type MsgSubmitClaim struct {
	CollectionId string
	ClaimId string
	AgentDid  DIDFragment
	AgentAddress string
	AdminAddress string
	UseIntent bool
	Amount github_com_cosmos_cosmos_sdk_types.Coins
	CW20Payment []CW20Payment
	CW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` this claim belongs
- `claimId` - a string containing the unique identifier of the claim in the cid hash format
- `agentAddress` - a string containing the account address that is submitting the claim
- `agentDid` - a string containing the Did of the agent that is submitting the claim
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `useIntent` - a boolean indicating if this claim is using an intent. If true, then the amount and CW20 payment are ignored and overridden with intent amounts. If true and there is no active intent then will error.
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified by service agent for claim approval. If both amount and CW20 payment are empty, then collection default is used. (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate amount wanted if no intent used)
- `cw20Payment` - an array of [CW20Payment](02_state.md#cw20payment) containing the custom CW20 payments specified by service agent for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used. (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate cw20Payment wanted if no intent used)
- `cw1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) containing the custom CW1155 payments specified by service agent for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used. (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate cw1155Payment wanted if no intent used)
- `memberAddress` - an optional string containing the team member this claim is on behalf of. Required when `useIntent` is true and the originating intent has a `memberAddress`; the values must match exactly (strict equality, including both being empty). If `useIntent` is false, `memberAddress` must be empty — without an intent there is no member context to attribute the claim to. The matching constraint on the oracle's [SubmitClaimAuthorization](02_state.md#submitclaimauthorization) is selected by `(collectionId, memberAddress)` strict equality.

## MsgEvaluateClaim

A `MsgEvaluateClaim` updates the `Evaluation` for a claim. On evaluation payments will be made for both the evaluation of the claim (towards the agent or oracle) as well as if the claim was `APPROVED` then the claim submitter will also get a payment, which will the preset amount defined on th [Collection](02_state.md#collection) `Payments` field or the evaluator can also define a custom `Amount` that can be paid out on approval. Note: no payments will be made if agent evaluates claim with status `invalidated`

```go
type MsgEvaluateClaim struct {
  ClaimId string
	CollectionId string
	Oracle string
	AgentDid  DIDFragment
	AgentAddress string
	AdminAddress string
	Status EvaluationStatus
	Reason uint32
	VerificationProof string
	Amount github_com_cosmos_cosmos_sdk_types.Coins
	CW20Payment []CW20Payment
	CW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `claimId` - a string containing the `id` of the claim the evaluation is for
- `collectionId` - a string containing the Collection `id` the claim this evaluation is for belongs to
- `oracle` - a string containing the DID of the Oracle entity that is evaluating the claim
- `agentAddress` - a string containing the account address that is submitting the evaluation
- `agentDid` - a string containing the Did of the agent that is submitting the evaluation
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `status` - a [EvaluationStatus](#evaluationstatus)
- `reason` - a integer for why the evaluation result was given (codes defined by evaluator)
- `verificationProof` - a string containing the proof to verify the linked resource (eg. the cid of the evaluation Verfiable Credential)
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on `Approval` payment if it is a custom amount and not the preset `Approval` from the [Collection](#collection). NOTE: if claim is using intent, then amount and CW20 and CW1155 payments are ignored and overridden with intent amounts. NOTE: if amount and CW20 and CW1155 payments are empty then collection default is used.
- `cw20Payment` - an array of [CW20Payment](02_state.md#cw20payment) containing the custom CW20 payments specified by evaluator for claim approval. NOTE: if claim is using intent, then amount and CW20 and CW1155 payments are ignored and overridden with intent amounts. NOTE: if amount and CW20 and CW1155 payments are empty then collection default is used.
- `cw1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) containing the custom CW1155 payments specified by evaluator for claim approval. NOTE: if claim is using intent, then amount and CW20 and CW1155 payments are ignored and overridden with intent amounts. NOTE: if amount and CW20 and CW1155 payments are empty then collection default is used.

## MsgDisputeClaim

A `MsgDisputeClaim` creates and stores a new Dispute.

```go
type MsgDisputeClaim struct {
	SubjectId string
	AgentDid    DIDFragment
	AgentAddress string
	DisputeType int32
	Data        *DisputeData
}
```

The field's descriptions is as follows:

- `subjectId` - a string containing the `id` of the claim the dispute is for. A unique field
- `disputeType` - a integer interpreted by the client
- `data` - a [DisputeData](02_state.md#disputedata)
- `agentAddress` - a string containing the account address that is submitting the dispute
- `agentDid` - a string containing the Did of the agent that is submitting the dispute

## MsgWithdrawPayment

A `MsgWithdrawPayment` creates and stores a new Withdrawal Payment which contains details about a payment that can be made if the `TimeoutNs` in a [Payment](02_state.md#payment) is non 0, which means the payout won't happen on submission/evaluation but an authz gets granted with the receiver as grantee that allows the execution of this message after the `ReleaseDate`. This will allow the receiver to execute the payout through a [MsgExec](https://docs.cosmos.network/main/build/modules/authz#msgexec) to execute the authz.

```go
type MsgWithdrawPayment struct {
	ClaimId string
	Inputs []github_com_cosmos_cosmos_sdk_x_bank_types.Input
	Outputs []github_com_cosmos_cosmos_sdk_x_bank_types.Output
	PaymentType PaymentType
	Contract_1155Payment *Contract1155Payment
	ToAddress string
	FromAddress string
	ReleaseDate *time.Time
	AdminAddress string
	CW20Payment []CW20Payment
	CW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `claimId` - a string containing the `id` of the claim the withdrawal is for
- `inputs` - a list of cosmos defined `Input` to pass to the the multisend tx to run to withdraw payment
- `outputs` - a list of cosmos defined `Output` to pass to the the multisend tx to run to withdraw payment
- `paymentType` - a [PaymentType](02_state.md#paymenttype)
- `contract_1155Payment` - a [Contract1155Payment](02_state.md#contract1155payment) DEPRECATED, use [CW1155Payment](02_state.md#cw1155payment) instead
- `toAddress` - a string containing the account address to make the payment to
- `fromAddress` - a string containing the account address to make the payment from
- `releaseDate` - a timestamp of the date that grantee can execute authorization to make the withdrawal payment, calculated from created date plus the timeout on [Collection](02_state.md#collection) `Payments`
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `cw20Payment` - an array of [CW20Payment](02_state.md#cw20payment) containing the CW20 tokens to be paid
- `cw1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) containing the CW1155 tokens to be paid

## MsgClaimIntent

A `MsgClaimIntent` creates a new intent to claim on a collection, with an optional custom payment amount. An agent must have a valid [SubmitClaimAuthorization](02_state.md#SubmitClaimAuthorization) to create and intent.
On successful intent creation the claim amount is transferred from the deed admin account to the Collections Escrow account.
Once a claim is created using the intent then the claim payment has a status of GUARANTEED.
Once the intent expires without any claim made then the funds is transferred back out of the Escrow account.

```go
type MsgClaimIntent struct {
	AgentDid     DIDFragment
	AgentAddress string
	CollectionId string
	Amount       github_com_cosmos_cosmos_sdk_types.Coins
	CW20Payment  []CW20Payment
	CW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `agentDid` - a string containing the DID of the agent creating the intent
- `agentAddress` - a string containing the account address of the agent creating the intent
- `collectionId` - a string containing the Collection `id` this intent belongs to
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used.
- `cw20Payment` - an array of [CW20Payment](02_state.md#cw20payment) containing the custom CW20 payments specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used.
- `cw1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) containing the custom CW1155 payments specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used.
- `memberAddress` - an optional string containing the team member this intent is on behalf of. **Required** when the collection has [MemberBudgets](02_state.md#memberbudget); **must be empty** when the collection does not. When provided, the handler:
  1. Looks up the matching [SubmitClaimConstraints](02_state.md#submitclaimconstraints) by `(collectionId, memberAddress)` strict equality — the oracle must hold a constraint that this specific member created.
  2. Loads the member's [MemberBudget](02_state.md#memberbudget), lazily resets the period if expired, and verifies the intent amount fits within the remaining budget.
  3. Deducts the intent amount from `periodSpent` (and CW20 equivalents).

## MsgCreateClaimAuthorization

A `MsgCreateClaimAuthorization` creates a new claim authorization on behalf of an entity admin account (SubmitClaimAuthorization or EvaluateClaimAuthorization). The creator (one executing this message), must have a valid [CreateClaimAuthorizationAuthorization](02_state.md#CreateClaimAuthorizationAuthorization) with at least one of the [constraints](02_state.md#CreateClaimAuthorizationConstraints) permitting the Authorization the creator is attempting to create.

```go
type MsgCreateClaimAuthorization struct {
	CreatorAddress string
	CreatorDid     DIDFragment
	GranteeAddress string
	AdminAddress   string
	CollectionId   string
	AuthType       CreateClaimAuthorizationType
	AgentQuota     uint64
	MaxAmount      github_com_cosmos_cosmos_sdk_types.Coins
	MaxCW20Payment []CW20Payment
	MaxCW1155Payment []CW1155Payment
	Expiration     *time.Time
	IntentDurationNs time.Duration
	BeforeDate     *time.Time
}
```

The field's descriptions is as follows:

- `creatorAddress` - a string containing the address of the creator (user with meta-authorization)
- `creatorDid` - a string containing the DID of the agent creating the authorization
- `granteeAddress` - a string containing the address of the grantee (who will receive the authorization)
- `adminAddress` - a string containing the account address defined in the [Collection](02_state.md#collection) `admin` field
- `collectionId` - a string containing the Collection ID the authorization applies to (for both submit and evaluate)
- `authType` - a [CreateClaimAuthorizationType](02_state.md#createclaimauthorizationtype) indicating the type of authorization to create (submit or evaluate, can't create both in a single request)
- `agentQuota` - a integer containing the quota for the created authorization (for both submit and evaluate)
- `maxAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the maximum amount that can be specified in the authorization (for both submit and evaluate)
- `maxCW20Payment` - an array of [CW20Payment](02_state.md#cw20payment) containing the maximum CW20 payment that can be specified in the authorization (for both submit and evaluate)
- `maxCW1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) containing the maximum CW1155 payment that can be specified in the authorization (for both submit and evaluate)
- `expiration` - a timestamp of the expiration time for the authorization. Be careful with this as it is the expiration of the authorization itself, not the constraints, meaning if the authorization expires all constraints will be removed with the authorization (standard authz behavior).
- `intentDurationNs` - a duration containing the maximum intent duration for the authorization allowed (for submit)
- `beforeDate` - a timestamp after which the grantee can't execute this authz anymore, a cut off date (for evaluate). If null then no before_date validation done.
- `memberAddress` - an optional string identifying the team member this authorization is being created for. The grantee's [CreateClaimAuthorizationAuthorization](02_state.md#createclaimauthorizationauthorization) constraint must have a matching `memberAddress` — strict equality (both empty for individual; both equal for team). The matched value is propagated onto the resulting [SubmitClaimConstraints](02_state.md#submitclaimconstraints) (or evaluate constraint) so the oracle's authorization carries member attribution downstream. This is the on-chain anti-spoofing primitive: a grantee cannot mint authorizations tagged with a member address other than the one the admin authorized for them.

## MsgSetCollectionMembers

A `MsgSetCollectionMembers` adds or updates one or more [MemberBudget](02_state.md#memberbudget) entries on a collection in a single transaction. The signer must be the collection's admin. Adding the first member to a collection turns it into a "team / enterprise" collection and from that point on every intent on the collection must carry a `memberAddress`.

```go
type MsgSetCollectionMembers struct {
	CollectionId string
	AdminAddress string
	Members      []*CollectionMemberInput
}

type CollectionMemberInput struct {
	MemberAddress         string
	Period                time.Duration
	PeriodSpendLimit      github_com_cosmos_cosmos_sdk_types.Coins
	PeriodCw20SpendLimit  []CW20Payment
	ResetPeriodSpent      bool
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to add / update members on
- `adminAddress` - a string containing the account address of the signer, must equal the Collection `admin` field
- `members` - a non-empty list of `CollectionMemberInput` entries (no duplicate `memberAddress` allowed within a single message). For each entry:
  - `memberAddress` - the team member's account address
  - `period` - the budget reset duration. Must be **at least 24 hours** (`MinMemberBudgetPeriod`); shorter durations are rejected to bound the cost of the lazy-reset roll-forward loop
  - `periodSpendLimit` - the maximum native coin spend per period. Must be sorted (standard Cosmos `Coins` invariant)
  - `periodCw20SpendLimit` - the maximum CW20 spend per period
  - At least one of `periodSpendLimit` or `periodCw20SpendLimit` must be non-empty / non-zero — a budget with no spend allowance is rejected (use `MsgRemoveCollectionMembers` instead to remove a member)
  - `resetPeriodSpent` - if true and the member already exists, `periodSpent` is cleared and `periodResetAt` is set to `now + period` (a manual mid-period reset). If false (or the member is new), an existing member's `periodSpent` and `periodResetAt` are preserved across the update

Behaviour per member:

- If the member does **not** exist on the collection: a new [MemberBudget](02_state.md#memberbudget) is created with `periodSpent = empty`, `periodResetAt = now + period`. Emits `MemberBudgetCreatedEvent`.
- If the member already exists and `resetPeriodSpent = false`: `periodSpendLimit` / `periodCw20SpendLimit` / `period` are updated; `periodSpent` / `periodResetAt` are preserved. The new limits apply to the rest of the current period. Emits `MemberBudgetUpdatedEvent`.
- If the member already exists and `resetPeriodSpent = true`: as above, plus `periodSpent` is zeroed and `periodResetAt = now + period`. Emits `MemberBudgetUpdatedEvent`.

## MsgRemoveCollectionMembers

A `MsgRemoveCollectionMembers` removes one or more [MemberBudget](02_state.md#memberbudget) entries from a collection in a single transaction. The signer must be the collection's admin.

This message does **not** revoke any existing claim authorizations the removed members granted to oracles — the admin should do that separately if needed. New intents from the removed members will fail at the budget lookup. Pending intents from removed members continue to expire and refund escrow normally; the budget restore is silently skipped because the budget no longer exists.

```go
type MsgRemoveCollectionMembers struct {
	CollectionId    string
	AdminAddress    string
	MemberAddresses []string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to remove members from
- `adminAddress` - a string containing the account address of the signer, must equal the Collection `admin` field
- `memberAddresses` - a non-empty list of member addresses to remove (no duplicates). Each address must currently exist as a member on the collection — the message fails if any one of them does not, and the entire transaction is rolled back atomically

Emits one `MemberBudgetRemovedEvent` per removed member, carrying the final budget state at the time of removal.
