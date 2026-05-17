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
	// Dispute / performance-deposit config (all optional, opt-in).
	ServiceAgentDepositRequired  sdk.Coins
	EvaluatorDepositRequired     sdk.Coins
	DisputeDepositAmount         sdk.Coins
	Adjudicators                 []*AdjudicationDid
	PenaltyAmountPerDispute      sdk.Coins
	MinDepositPeriod             time.Duration
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
- `serviceAgentDepositRequired` - minimum performance-deposit balance an SA must hold on this collection (gates `MsgSubmitClaim`). Empty/zero means no gate.
- `evaluatorDepositRequired` - same for evaluators (gates `MsgEvaluateClaim`).
- `disputeDepositAmount` - stake the disputer must attach inline on `MsgDisputeClaim`.
- `adjudicators` - the whitelist of approved adjudicators ([AdjudicationDid](02_state.md#adjudicationdid) entries: `did` + per-adjudicator `reward_percentage`). Required to be non-empty if any deposit/penalty field is set.
- `penaltyAmountPerDispute` - fixed penalty on `AWARDED`. If empty the adjudicator picks per-resolution, bounded by the loser's role deposit-required. If set, must be ≤ each non-empty role deposit-required at validation time.
- `minDepositPeriod` - minimum duration a performance deposit must remain locked after the most recent top-up. Zero disables the lock (legacy behavior).

The dispute state shapes themselves (`Dispute`, `DisputeData`, `DisputeResolution`, `AdjudicationDid`, `AgentDepositBalance`, plus the active-dispute and dispute-subject indexes) are defined in [02_state.md](02_state.md); the per-message economics are documented under [MsgDisputeClaim](#msgdisputeclaim), [MsgAdjudicateDispute](#msgadjudicatedispute), [MsgUpdateCollectionDisputeConfig](#msgupdatecollectiondisputeconfig), [MsgAddPerformanceDeposit](#msgaddperformancedeposit), and [MsgWithdrawPerformanceDeposit](#msgwithdrawperformancedeposit) below.

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

## MsgUpdateCollectionQuota

A `MsgUpdateCollectionQuota` updates a Collection's `quota` (the maximum number of claims that may be submitted). Use `0` for unlimited.

```go
type MsgUpdateCollectionQuota struct {
	CollectionId string
	Quota        uint64
	AdminAddress string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` to make the update to
- `quota` - the new maximum claim count. `0` means unlimited. Must be `0` or `>= collection.count` (`ErrCollectionQuotaBelowCount`) — cannot retroactively cap below claims already submitted.
- `adminAddress` - a string containing the account address of the private key signing the transaction, must be same as [Collection](02_state.md#collection) admin field

## MsgSubmitClaim

A `MsgSubmitClaim` creates and stores a new Claim made towards a `Collection`. On Submission of claim `SUBMISSION` payments will be made if there is any defined in the [Collection](02_state.md#collection) `Payments`.

Gates (in addition to the existing collection-state / quota / intent checks):

- If the collection has `serviceAgentDepositRequired` set, the SA's [AgentDepositBalance](02_state.md#agentdepositbalance) on the collection must be ≥ the required amount (`ErrAgentDepositInsufficient`).
- The SA must have **no OPEN dispute** targeting their SUBMITTER role on the collection (`ErrAgentHasActiveDispute`).

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

A `MsgEvaluateClaim` updates the `Evaluation` for a claim. On evaluation payments will be made for both the evaluation of the claim (towards the agent or oracle) as well as if the claim was `APPROVED` then the claim submitter will also get a payment, which will the preset amount defined on th [Collection](02_state.md#collection) `Payments` field or the evaluator can also define a custom `Amount` that can be paid out on approval. Note: no payments will be made if agent evaluates claim with status `invalidated` or `flagged`.

Behavior:

- `status = DISPUTED` is **rejected** with `ErrEvaluationStatusDisputedDeprecated`. Disputes are recorded by [MsgDisputeClaim](#msgdisputeclaim) on the `Dispute` record, not on the evaluation.
- If the collection has `evaluatorDepositRequired` set, the EA's [AgentDepositBalance](02_state.md#agentdepositbalance) must be ≥ the required amount (`ErrAgentDepositInsufficient`).
- The EA must have **no OPEN dispute** targeting their EVALUATOR role on the collection (`ErrAgentHasActiveDispute`).

### Re-evaluation rules

By default a claim's evaluation is set once and locks the claim. The exception is `flagged`:

- A claim whose current evaluation status is `flagged` (5) **can** be re-evaluated by another authorized evaluator, or by the same evaluator that flagged it (e.g. once they have more information). The new evaluation may be another `flagged` (escalating to a third evaluator) or any terminal status.
- The same agent cannot flag a claim they have already flagged. The keeper checks both the current `Evaluation` and every entry in `Claim.evaluationHistory` and rejects the message with `ErrSelfReFlag` if the agent address appears in either.
- A claim with any *terminal* status (`approved`, `rejected`, `disputed`, `invalidated`) is locked. A second `MsgEvaluateClaim` against it returns `ErrClaimDuplicateEvaluation` regardless of the message status.
- On a successful re-evaluation, the prior `Evaluation` is appended to `Claim.evaluationHistory` and the new `Evaluation` becomes `Claim.evaluation`. Claims that have only ever been evaluated once leave `evaluationHistory` empty.

### `flagged` semantics

When `status` is `flagged` the keeper:

- skips the evaluator payment (the `Evaluation` payment for the collection does not fire),
- skips the approval / rejection payment branch,
- leaves any intent escrow funds in place (a subsequent terminal `approved` will pay the claim agent out of escrow; any other terminal status will refund the approval account),
- increments `Collection.flagged` (cumulative event counter, never decremented),
- increments `Collection.flaggedActive` only if the prior evaluation was not already `flagged` (re-flag chains do not double-count), and
- emits `ClaimEvaluatedEvent` and `ClaimUpdatedEvent` carrying the new `Evaluation` and the updated `Claim` (now including the latest `evaluationHistory`).

`AgentQuota` is decremented for `flagged` the same as for `approved` / `rejected` / `disputed`. Only `invalidated` skips the decrement.

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

Authorization: dispute filing is open to **anyone with a registered IID DID**. The chain-level IID ante still requires `agentAddress` to be a key on `agentDid`, but no module-level admin / controller / authz check applies. Spam is gated economically by the collection's `disputeDepositAmount` (the disputer must stake it inline; lost on `DISMISSED`).

```go
type MsgDisputeClaim struct {
	SubjectId string
	AgentDid    DIDFragment
	AgentAddress string
	DisputeType int32
	Data        *DisputeData
	TargetRole  DisputeTargetRole // v7: SUBMITTER or EVALUATOR (required)
}
```

The field's descriptions is as follows:

- `subjectId` - a string containing the `id` of the claim the dispute is for.
- `disputeType` - a integer interpreted by the client
- `data` - a [DisputeData](02_state.md#disputedata)
- `agentAddress` - a string containing the account address that is submitting the dispute
- `agentDid` - a string containing the Did of the agent that is submitting the dispute
- `targetRole` - a [DisputeTargetRole](02_state.md#disputetargetrole) (`SUBMITTER` or `EVALUATOR`; `UNSPECIFIED` is rejected). Identifies which party of the claim is being disputed.

v7 keeper behavior:

- The collection's `disputeDepositAmount` is debited from the disputer's wallet into the collection escrow and snapshotted on the `Dispute` record. If `disputeDepositAmount` is set, the collection's `adjudicators` list must be non-empty (`ErrAdjudicationNotConfigured`) — refuses to lock funds that can never be released through adjudication.
- The disputed agent is derived from `targetRole` + current claim state (`claim.agentAddress` for `SUBMITTER`, `claim.evaluation.agentAddress` for `EVALUATOR`) at both filing time (to write the active-dispute index entry) and adjudication time (to slash the right balance).
- Disputing the `EVALUATOR` role of a `FLAGGED` evaluation is rejected (`ErrDisputeTargetEvaluatorFlagged`) — re-evaluate to a terminal status first.
- One OPEN dispute per `(subject_id, target_role)`. AWARDED permanently blocks future disputes on that pair; DISMISSED allows new filings.
- The full AWARDED / DISMISSED economics — pot derivation, 80/20 split, penalty cap, payout routing — are documented under [MsgAdjudicateDispute](#msgadjudicatedispute) below.

## MsgAdjudicateDispute

A `MsgAdjudicateDispute` settles an OPEN dispute. `adjudicatorDid` must appear in `collection.adjudicators`. The signer (`adjudicatorAddress`) must be a registered verification method on that DID document under `Authentication` or `AssertionMethod` — the same chain-level rule the IID ante enforces on every DID-gated message in the module. The reward percentage applied to the penalty pot is the **matching adjudicator entry's** `reward_percentage` (not a collection-wide value), letting different adjudicators charge different fees. Payout routing is decided by whether `adjudicatorDid` resolves to an entity in the entity module: if yes, the adjudicator share is paid to the entity's `EntityAdjudicatorRevenueAccountName` account (auto-created); otherwise it goes directly to `adjudicatorAddress`.

```go
type MsgAdjudicateDispute struct {
	SubjectId           string
	TargetRole          DisputeTargetRole
	AdjudicatorDid      string
	AdjudicatorAddress  string
	Outcome             DisputeStatus  // AWARDED or DISMISSED
	Data                *DisputeData   // optional: adjudicator's opinion doc
	PenaltyAmount       sdk.Coins      // optional if collection has fixed penalty
}
```

The field's descriptions is as follows:

- `subjectId` - the `id` of the claim being disputed.
- `targetRole` - the [DisputeTargetRole](02_state.md#disputetargetrole) (`SUBMITTER` or `EVALUATOR`) identifying which dispute to adjudicate. Together with `subjectId` uniquely identifies the OPEN dispute.
- `adjudicatorDid` - the DID acting as adjudicator. Must appear in `collection.adjudicators`.
- `adjudicatorAddress` - the signer. Must be a registered key on `adjudicatorDid`'s DID document under `Authentication` or `AssertionMethod`. Enforced by the IID ante (`VerifyIidControllersAgainstSignature`) at tx time, plus a defense-in-depth re-check in the keeper.
- `outcome` - a [DisputeStatus](02_state.md#disputestatus); must be `AWARDED` or `DISMISSED` (`OPEN` is rejected).
- `data` - a [DisputeData](02_state.md#disputedata) payload (uri + proof + type + encrypted) the adjudicator wants pinned to the resolution. Symmetric with `MsgDisputeClaim.data` — lets the adjudicator publish a signed opinion document, declare its MIME type, and flag encryption. Optional; if supplied, all three string fields must be non-empty. Stored verbatim on `DisputeResolution.data`. Replaces the v6 free-form `reason` string.
- `penaltyAmount` - applied only on `AWARDED`. Ignored if `collection.penaltyAmountPerDispute` is set (the collection fixed value is used). Otherwise must be supplied and is capped at the loser's role deposit-required.

Outcome semantics:

- `AWARDED`: loser is the targeted agent. Penalty = `min(intended, targetAgent.balance)` is debited from their [AgentDepositBalance](02_state.md#agentdepositbalance). `1 − adjudicatorEntry.rewardPercentage` of the pot goes to the disputer; `adjudicatorEntry.rewardPercentage` goes to the adjudicator payout address. The disputer's `disputeDeposit` is refunded in full. `Collection.disputesAwarded++`.
- `DISMISSED`: loser is the disputer. Pot = `disputeDeposit`. Split 80/20 with the same percentage; winner = target agent (vindicated), adjudicator gets their share. `Collection.disputesDismissed++`.

In both cases `Collection.disputesOpen--`, the active-dispute index entry is removed (unblocking the targeted agent), and the resolution record is populated with both `intendedPenalty` and `actualPenaltyPaid`.

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

## MsgUpdateCollectionDisputeConfig

A `MsgUpdateCollectionDisputeConfig` replaces the dispute / performance-deposit configuration on a collection. All fields are full replacements (not merges) — the caller must send the desired full state. Signer is the collection admin (entity account; typically goes through `MsgExec` with the admin authz pattern).

```go
type MsgUpdateCollectionDisputeConfig struct {
	CollectionId                string
	AdminAddress                string
	ServiceAgentDepositRequired sdk.Coins
	EvaluatorDepositRequired    sdk.Coins
	DisputeDepositAmount        sdk.Coins
	Adjudicators                []*AdjudicationDid
	PenaltyAmountPerDispute     sdk.Coins
	MinDepositPeriod            time.Duration
}
```

Validation (in addition to admin-signer match):

- The cross-field invariants in `ValidateCollectionDisputeConfig` apply: penalty ≤ each non-empty role deposit-required; `adjudicators` must be non-empty if any deposit/penalty/disputer-stake field is set; each entry's DID is valid, DIDs unique within the list; each entry's `reward_percentage` in `[0, 100]`; `minDepositPeriod` non-negative.
- Clearing `adjudicators` is **rejected** while `collection.disputesOpen > 0` (`ErrAdjudicationNotConfigured`) — otherwise existing OPEN disputes would have no path to resolution.

In-flight disputes are unaffected: each `Dispute` snapshots `disputeDeposit` at filing and the adjudicator authorization / payout rules read collection state at adjudication time (which is by design — admins can adjust dispute economics for ongoing disputes; the actual slash is always bounded by available balance). Changing `minDepositPeriod` does not retroactively change `withdrawableAt` on existing balances; it only affects future top-ups.

## MsgAddPerformanceDeposit

A `MsgAddPerformanceDeposit` tops up an agent's [AgentDepositBalance](02_state.md#agentdepositbalance) on a collection. Funds move agent's wallet → collection escrow. Permitted regardless of whether the agent has active disputes — only withdrawal and new submissions are gated. Signer is the agent themselves (not via authz).

```go
type MsgAddPerformanceDeposit struct {
	CollectionId string
	AgentAddress string
	Amount       sdk.Coins
}
```

The field's descriptions is as follows:

- `collectionId` - the Collection `id` the deposit is held on.
- `agentAddress` - the agent the deposit is held for. Anyone may fund their own balance on any collection — no admin authz required.
- `amount` - the amount to add. Must be strictly positive.

On the first `MsgAddPerformanceDeposit` for an `(collection, agent)` pair, `AgentDepositBalanceCreatedEvent` fires; subsequent top-ups emit `AgentDepositBalanceUpdatedEvent`.

## MsgWithdrawPerformanceDeposit

A `MsgWithdrawPerformanceDeposit` pulls some / all of an agent's [AgentDepositBalance](02_state.md#agentdepositbalance) back to their wallet. Signed by the agent.

Rejected if:

- any OPEN dispute targets the agent on the collection (`ErrAgentDepositBalanceCannotWithdraw`); or
- the balance's `withdrawableAt` is still in the future (`ErrAgentDepositLocked`) — i.e. the most recent top-up is still inside the collection's `minDepositPeriod` window. This closes the deposit-then-immediately-withdraw exploit.

```go
type MsgWithdrawPerformanceDeposit struct {
	CollectionId string
	AgentAddress string
	Amount       sdk.Coins // empty means "withdraw full current balance"
}
```

The field's descriptions is as follows:

- `collectionId` - the Collection `id` the balance is held on.
- `agentAddress` - the owner of the balance (must equal the signer).
- `amount` - the amount to withdraw. Must be ≤ current balance. **Empty** withdraws the full current balance.

Emits `AgentDepositBalanceUpdatedEvent` on partial withdrawal, or `AgentDepositBalanceRemovedEvent` when the balance drains to zero and the KV entry is deleted.
