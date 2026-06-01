# State

## Collections

A Collections is stored in the state and is accessed by the identity of the collection that is fetched and incremented onchain using the `Param`: `CollectionSequence`.

- Collections: `0x01 | collectionId -> ProtocolBuffer(Collection)`

## Claims

A Claim is stored in the state and is accessed by the identity of the ClaimId(user provided).

- Claims: `0x02 | claimId -> ProtocolBuffer(Claim)`

## Disputes

A Dispute is stored in the state under its own `Data.Proof` CID. This preserves backward compatibility with legacy disputes that didn't carry a target role. A secondary index ([Dispute Subject Index](#dispute-subject-index)) is used to look up disputes by `(subject_id, target_role)` — the lookup path used by dispute filing and adjudication.

- Disputes: `0x03 | disputeProof(CID) -> ProtocolBuffer(Dispute)`

## Intents

An Intent is stored in the state under a composite key of agent address, collection id, and intent id.

- Intents: `0x04 | agentAddress + "/" + collectionId + "/" + intentId -> ProtocolBuffer(Intent)`

## Member Budgets

A MemberBudget is stored in the state under a composite key of collection id and member address. Storing each member budget under its own KV entry keeps reads and writes O(1) per member regardless of total team size — operations on one member do not require loading or rewriting other members' budgets, so per-operation gas is constant in team size.

- Member Budgets: `0x05 | collectionId + "/" + memberAddress -> ProtocolBuffer(MemberBudget)`

## Agent Deposit Balances

An AgentDepositBalance is the rolling performance-deposit balance held by a single agent on a single collection. Stored under a composite key of `(collectionId, agentAddress)` so reads / writes are O(1). The actual funds live inside the collection's existing `escrow_account`; this entry is the per-agent accounting record.

- Agent Deposit Balances: `0x06 | collectionId + "/" + agentAddress -> ProtocolBuffer(AgentDepositBalance)`

## Active Dispute Index

A presence-only secondary index used to answer "does this agent have any OPEN dispute targeting them on this collection?" in O(1) gas. The check is a prefix scan with limit 1 — used by `MsgSubmitClaim` / `MsgEvaluateClaim` / `MsgWithdrawPerformanceDeposit` to gate the actor while at least one OPEN dispute exists.

- Active Dispute Index: `0x07 | collectionId + "/" + agentAddress + "/" + subjectId -> []byte{1}` (presence-only)

Entries are written on `MsgDisputeClaim` (one per targeted role's agent) and deleted on `MsgAdjudicateDispute` when the dispute is resolved.

## Dispute Subject Index

A secondary index from `(subject_id, target_role)` to the most-recent dispute's proof CID. Walks back to the primary [Dispute](#dispute) record to read its status. Governs the "one OPEN dispute per (subject, role); AWARDED permanently blocks; DISMISSED allows new filings" rule.

- Dispute Subject Index: `0x08 | subjectId + "/" + targetRole(int) -> disputeProof`

# Types

### Collection

```go
type Collection struct {
	Id            string
	Entity        string
	Admin         string
	Protocol      string
	StartDate     *time.Time
	EndDate       *time.Time
	Quota         uint64
	Count         uint64
	Evaluated     uint64
	Approved      uint64
	Rejected      uint64
	Disputed      uint64
	Invalidated   uint64
	State         CollectionState
	Payments      *Payments
	EscrowAccount string
	Intents       CollectionIntentOptions
	Flagged       uint64
	FlaggedActive uint64
	// Dispute / performance-deposit config (all optional; opt-in).
	ServiceAgentDepositRequired  sdk.Coins
	EvaluatorDepositRequired     sdk.Coins
	DisputeDepositAmount         sdk.Coins
	Adjudicators                 []*AdjudicationDid
	PenaltyAmountPerDispute      sdk.Coins
	MinDepositPeriod             time.Duration
	// Dispute counters (internally calculated).
	DisputesOpen      uint64
	DisputesAwarded   uint64
	DisputesDismissed uint64
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
- `flagged` - a integer containing the cumulative number of times any claim in this collection has been flagged by an evaluator (internally calculated). This is an event-count metric — it increments on every flag, including re-flags by different agents on the same claim, and is **never decremented**.
- `flaggedActive` - a integer containing the number of claims currently in `FLAGGED` state (internally calculated). Increments on the first transition to `FLAGGED` for a claim, decrements when that claim transitions to a terminal evaluation status. Re-flags (`FLAGGED → FLAGGED` by a different agent) leave it unchanged.
- `serviceAgentDepositRequired` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) minimum performance-deposit balance a service agent must hold on this collection in order to `MsgSubmitClaim`. Empty / zero means no SA deposit gate.
- `evaluatorDepositRequired` - same as `serviceAgentDepositRequired` but for evaluators, gating `MsgEvaluateClaim`. Empty / zero means no EA deposit gate.
- `disputeDepositAmount` - the stake a disputer must attach to each `MsgDisputeClaim` (locked inline, refunded on `AWARDED`, forfeited as the penalty pot on `DISMISSED`). Empty / zero means no disputer stake required (legacy behavior).
- `adjudicators` - the whitelist of approved adjudicators, each entry an [AdjudicationDid](#adjudicationdid) pairing a DID with that adjudicator's own `reward_percentage`. Adjudicators self-set their fee, turning the whitelist into a competitive market (low-fee adjudicators may win on volume, high-fee adjudicators may trade on reputation). The chain doesn't enforce who adjudicates a given dispute — whichever whitelisted DID lands `MsgAdjudicateDispute` first wins, and that entry's `reward_percentage` is applied. Required to be non-empty if any of the deposit / penalty fields above is set; clearing it is blocked while `disputesOpen > 0`.
- `penaltyAmountPerDispute` - the fixed penalty applied on `AWARDED` adjudications. If empty, the adjudicator picks the penalty per-resolution (bounded by the loser's role deposit-required). At collection-validation time the fixed penalty must be ≤ each non-empty role deposit-required.
- `minDepositPeriod` - the minimum duration a performance deposit must remain locked after the most recent `MsgAddPerformanceDeposit` before `MsgWithdrawPerformanceDeposit` can succeed. Closes the in-same-tx exploit where an agent could deposit + submit + withdraw atomically and leave zero stake at dispute time. Zero disables the lock. Each top-up rolls `AgentDepositBalance.withdrawableAt` forward to `max(current, now + minDepositPeriod)`; the slash path is not gated by this lock.
- `disputesOpen` - the number of currently-OPEN disputes against any claim in this collection (internally calculated). Drives the "can the admin clear the adjudicator whitelist?" guard.
- `disputesAwarded` - cumulative number of `AWARDED` adjudications on this collection (internally calculated, never decremented).
- `disputesDismissed` - cumulative number of `DISMISSED` adjudications on this collection (internally calculated, never decremented).

### AdjudicationDid

An AdjudicationDid is a single entry on a collection's `adjudicators` whitelist — pairs an adjudicator DID with that adjudicator's own reward percentage. Lets adjudicators self-set fees so the whitelist functions as a competitive market.

```go
type AdjudicationDid struct {
	Did              string
	RewardPercentage math.LegacyDec // 0–100
}
```

The field's descriptions is as follows:

- `did` - the adjudicator's DID. Must appear on the `adjudicators` list to be allowed to settle disputes on the collection.
- `rewardPercentage` - the share (`LegacyDec`, range `[0, 100]`) of each actual penalty payout that goes to **this** adjudicator when they settle a dispute. The remainder goes to the dispute winner. `0` means the adjudicator works for free; `100` means they take the full pot.

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
	CW1155Payment         []CW1155Payment
	TimeoutNs             time.Duration
	CW20Payment           []CW20Payment
	IsOraclePayment       bool
}
```

The field's descriptions is as follows:

- `account` - a string containing the account address from which the payment will be made (must be an [EntityAccount](/x/entity/spec/02_state.md#entityaccount) of the `Entity` field for the Collection)
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on payment
- `contract_1155Payment` - a [Contract1155Payment]DEPRECATED, use [CW1155Payment](#cw1155payment) instead
- `timeoutNs` - a duration containing the timeout after claim/evaluation to create authZ for payment, if 0 then immediate direct payment is made
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the CW20 tokens to be paid
- `isOraclePayment` - a boolean indicating if the payment is for oracle payments, meaning it will go through network fees split. Only allowed for APPROVAL payment types. If true and the payment contains CW20 payments, the claim will only be successful if an intent exists to ensure immediate CW20 payment split, since there is no WithdrawalAuthorization to manage the CW20 payment split for delayed payments.
- `cw1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the CW1155 tokens and amount to be paid

Restrictions:

- If `isOraclePayment` is true, then `cw1155Payment` is not allowed (Since we can't split CW1155 payments)
- If `isOraclePayment` is true, then `cw20Payment` is only allowed if an intent exists to ensure immediate CW20 payment split, since there is no WithdrawalAuthorization to manage the CW20 payment split for delayed payments.
- For Evaluation payments, only native tokens are allowed. No CW20 or CW1155 payments are allowed.

### CW1155Payment

A CW1155Payment stores information about the payment to make if it is a cw1155 tokens payment.

```go
type CW1155Payment struct {
	Address string
	TokenId []string
	Amount  uint64
}
```

The field's descriptions is as follows:

- `address` - a string containing the smart contract address where the tokens can be transferred on (the cw1155 smart contract)
- `amount` - a integer indicating how many tokens must be transferred for the payment
- `tokenId` - an array of strings containing the `id` of the tokens on the cw1155 smart contract to transfer, if ids is provided,
  then we will add up the tokens in the provided list till the amount is reached, if not enough tokens, then we will throw an error
  If no ids are provided (empty array), then we take any of the tokens the account has and add them up till the amount is reached.
  So specifying a list of tokenIds limits the payment to only those tokens as allowed, empty list means any token id is allowed.

### CW1155IntentPayment

A CW1155IntentPayment stores information about the CW1155Payment made if an intent was used to submit a claim. Since the CW1155Payment only stores the total amount to be paid, but tokenIds can be variable, we need to store the individual tokenIds and amounts that was used to make the intent, so we can transfer the tokens to/from the escrow account.

```go
type CW1155IntentPayment struct {
	Address string
	Tokens  []CW1155IntentPaymentToken
}
type CW1155IntentPaymentToken struct {
	TokenId string
	Amount  uint64
}
```

The field's descriptions is as follows:

- `address` - a string containing the smart contract address where the tokens can be transferred on (the cw1155 smart contract)
- `tokens` - an array of CW1155IntentPaymentToken containing the individual tokenIds and amounts that was used to make the intent

### Contract1155Payment (DEPRECATED, please use CW1155Payment instead)

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
	CW1155Payment  []CW1155Payment
	CW1155IntentPayment []CW1155IntentPayment
	MemberAddress     string
	EvaluationHistory []*Evaluation
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
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified by service agent for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate amount wanted if no intent used)
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified by service agent for claim approval. If both amount and CW20 and CW1155 payments are empty, then collection default is used (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate cw20Payment wanted if no intent used)
- `cw1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the custom CW1155 payments specified by service agent for claim approval. If amount, CW20 and CW1155 payments are empty, then collection default is used (Note the Evaluation agent can still override this, this value is for whoever submits the claim to indicate cw1155Payment wanted if no intent used)
- `cw1155IntentPayment` - an array of [CW1155IntentPayment](#cw1155intentpayment) containing the custom CW1155 payments if an intent was used to submit the claim.
- `memberAddress` - a string containing the team member address this claim is on behalf of, if any. Copied from the intent (`Intent.MemberAddress`) when `useIntent` is true. Empty for individual subscriptions. Used to know which [MemberBudget](#memberbudget) to credit on rejection / dispute / invalidation.
- `evaluationHistory` - an array of prior [Evaluation](#evaluation) entries in chronological order (oldest first). The most recent evaluation always lives in `evaluation`; only superseded entries are appended here. Empty for claims that have been evaluated at most once. Populated when an evaluator transitions a claim out of `FLAGGED` (either by another flag in a re-flag chain, or by a terminal finalisation) — at that point the prior evaluation is moved into history and the new one becomes `evaluation`.

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
	CW1155Payment      []CW1155Payment
	CW1155IntentPayment []CW1155IntentPayment
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
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount to be paid on `Approval` payment if it is a custom amount and not the preset `Approval` from the [Collection](#collection) (Note if intent was used, then the amount and CW20 and CW1155 payments are ignored and overridden with intent amounts)
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified by evaluator for claim approval (Note if intent was used, then the amount and CW20 and CW1155 payments are ignored and overridden with intent amounts)
- `cw1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the custom CW1155 payments specified by evaluator for claim approval (Note if intent was used, then the amount and CW20 and CW1155 payments are ignored and overridden with intent amounts)
- `cw1155IntentPayment` - an array of [CW1155IntentPayment](#cw1155intentpayment) containing the custom CW1155 payments if an intent was used to submit the claim.

### Dispute

A Dispute stores information concerning the dispute made towards a [Claim](#claim). v7 extended the record with target-role + economic-stake fields, plus a `status` lifecycle that culminates in a populated `resolution`.

```go
type Dispute struct {
	SubjectId       string
	Type            int32
	Data            *DisputeData
	// Extended fields
	TargetRole      DisputeTargetRole
	DisputerAddress string
	DisputerDid     string
	DisputeDeposit  sdk.Coins
	SubmittedAt     *time.Time
	Status          DisputeStatus
	Resolution      *DisputeResolution
}
```

The field's descriptions is as follows:

- `subjectId` - a string containing the `id` of the claim the dispute is for.
- `type` - a integer interpreted by the client
- `data` - a [DisputeData](#disputedata)
- `targetRole` - a [DisputeTargetRole](#disputetargetrole) identifying which party of the claim is being disputed. `UNSPECIFIED` only appears on disputes migrated from pre-v7 state. The disputed agent's address is not stored on the record — it is derived at read time as `claim.agentAddress` for `SUBMITTER` or `claim.evaluation.agentAddress` for `EVALUATOR` (safe to re-derive because EVALUATOR disputes can only exist against terminal evaluations, which are immutable).
- `disputerAddress` - the account that filed the dispute and locked the dispute deposit.
- `disputerDid` - the DID of the disputer.
- `disputeDeposit` - the amount locked by the disputer at filing, equal to `collection.disputeDepositAmount` at that moment. Refunded on `AWARDED`; forfeited as the penalty pot on `DISMISSED`.
- `submittedAt` - the block time the dispute was filed.
- `status` - the current [DisputeStatus](#disputestatus).
- `resolution` - a [DisputeResolution](#disputeresolution), populated when the dispute is adjudicated.

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
- `proof ` - a string representing the proof to verify the data linked resource. This is the **primary key** of the Dispute record in state.
- `encrypted ` - a boolean value for whether this data linked resource is encrypted or not

### DisputeResolution

A DisputeResolution stores the outcome of `MsgAdjudicateDispute`. Tracks both the intended penalty (what the adjudicator / collection prescribed) and the actual penalty paid (capped at the loser's available deposit balance — may be less than intended if a prior dispute already drained the balance).

```go
type DisputeResolution struct {
	AdjudicatorDid           string
	AdjudicatorAddress       string
	AdjudicatorPayoutAddress string
	ResolvedAt               *time.Time
	Reason                   string
	IntendedPenalty          sdk.Coins
	ActualPenaltyPaid        sdk.Coins
	WinnerAmount             sdk.Coins
	AdjudicatorAmount        sdk.Coins
	WinnerAddress            string
	LoserAddress             string
}
```

The field's descriptions is as follows:

- `adjudicatorDid` - the DID that adjudicated; must be in `collection.adjudicators`.
- `adjudicatorAddress` - the signer of `MsgAdjudicateDispute`. Either an `EntityAccount` of the adjudicator DID, or a key registered on the DID document.
- `adjudicatorPayoutAddress` - where the adjudicator share was actually paid: the entity's `EntityAdjudicatorRevenueAccountName` account when `adjudicatorDid` resolves to an entity (auto-created on first adjudication), otherwise the signer address itself.
- `resolvedAt` - block time of adjudication.
- `reason` - free-form reason text from the adjudicator.
- `intendedPenalty` - the penalty the adjudicator selected (or the collection's fixed `penaltyAmountPerDispute` if set). May exceed what was actually paid if the loser's balance was insufficient.
- `actualPenaltyPaid` - what was actually slashed from the loser's balance (or paid out from the dispute deposit on `DISMISSED`). Always ≤ `intendedPenalty`.
- `winnerAmount` - the share of `actualPenaltyPaid` that went to the dispute winner (disputer on `AWARDED`, target agent on `DISMISSED`).
- `adjudicatorAmount` - the share of `actualPenaltyPaid` that went to the adjudicator.
- `winnerAddress` - the address that received `winnerAmount`.
- `loserAddress` - the address whose balance / deposit was slashed.

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
	CW1155Payment []CW1155Payment
	CreateDate    *time.Time
	ExpireDate    *time.Time
	Status        IntentStatus
	MemberAddress string
}
```

The field's descriptions is as follows:

- `id` - a string containing the intent's identifier, it is incremented on chain for each intent created
- `collectionId` - a string containing the Collection `id` this intent belongs to
- `agentDid` - a string containing the DID of the agent creating the intent
- `agentAddress` - a string containing the account address of the agent creating the intent
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the custom amount specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used for intent amount
- `cw20Payment` - an array of [CW20Payment](#cw20payment) containing the custom CW20 payments specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used for intent amount
- `cw1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the custom CW1155 payments specified for claim approval. If amount and CW20 and CW1155 payments are empty, then collection default is used for intent amount
- `createDate` - a timestamp of the date and time that the intent was created on-chain
- `expireDate` - a timestamp of the date and time that the intent will expire
- `status` - a [IntentStatus](#intentstatus) indicating the current status of the intent
- `memberAddress` - a string containing the team member this intent is on behalf of, if any. Required if the collection has [MemberBudgets](#memberbudget); empty for individual subscriptions. Validated against the oracle's [SubmitClaimConstraints](#submitclaimconstraints) `memberAddress` (the constraint must have been created by this same member) and used to deduct from the corresponding [MemberBudget](#memberbudget). Carried into the resulting [Claim](#claim) and used to restore the budget on rejection / dispute / invalidation / expiration.

### MemberBudget

A MemberBudget stores a team member's periodic spending budget on a [Collection](#collection). Member budgets are an opt-in mechanism that turns a collection into a "team / enterprise" collection — when one or more `MemberBudget` entries exist for a collection, all intents and claims on that collection must be attributed to a specific member, and per-member spending is enforced at intent creation time.

```go
type MemberBudget struct {
	CollectionId         string
	MemberAddress        string
	Period               time.Duration
	PeriodSpendLimit     github_com_cosmos_cosmos_sdk_types.Coins
	PeriodSpent          github_com_cosmos_cosmos_sdk_types.Coins
	PeriodCw20SpendLimit []CW20Payment
	PeriodCw20Spent      []CW20Payment
	PeriodResetAt        *time.Time
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` this budget belongs to
- `memberAddress` - a string containing the team member's blockchain address
- `period` - a duration for the budget reset cycle (e.g., 30 days). Must be at least 24 hours; shorter periods are rejected to prevent griefing through the lazy-reset loop
- `periodSpendLimit` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object specifying the maximum native coin spend per period
- `periodSpent` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object tracking native coin amounts already consumed (intented) in the current period
- `periodCw20SpendLimit` - an array of [CW20Payment](#cw20payment) specifying the maximum CW20 spend per period
- `periodCw20Spent` - an array of [CW20Payment](#cw20payment) tracking CW20 amounts already consumed in the current period
- `periodResetAt` - a timestamp of the next period reset boundary. Period reset is **lazy**: when an intent or restore operation runs and `now >= periodResetAt`, `periodSpent` and `periodCw20Spent` are zeroed and `periodResetAt` is rolled forward by `period` until it lands in the future. No background scheduler is involved.

Lifecycle:

- **Created** — when a member is added to a collection via `MsgSetCollectionMembers`. `periodSpent` starts empty and `periodResetAt = now + period`.
- **Updated (admin)** — when an existing member's budget is changed via `MsgSetCollectionMembers`. `periodSpent` and `periodResetAt` are preserved unless `resetPeriodSpent` is true.
- **Updated (intent)** — when a [MsgClaimIntent](03_messages.md#msgclaimintent) is created with a matching `memberAddress`, the intent amount is added to `periodSpent` (and CW20 equivalents).
- **Updated (restore)** — when a claim is rejected / disputed / invalidated, or an intent expires before being used, the corresponding amounts are subtracted back from `periodSpent`. If the period has reset between the original deduction and the restore, the restore is skipped (the old period's spend doesn't carry over into the new period).
- **Removed** — when removed via `MsgRemoveCollectionMembers`. Pending intents from a removed member still expire and refund escrow normally; budget restore is silently skipped.

### AgentDepositBalance

An AgentDepositBalance is the rolling performance-deposit balance held by an agent (service agent, evaluator, or third-party disputer using the same accounting bucket) on a single collection. The funds live in the collection's existing `escrow_account`; this record is the per-agent accounting.

```go
type AgentDepositBalance struct {
	CollectionId    string
	AgentAddress    string
	Amount          sdk.Coins
	WithdrawableAt  *time.Time
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` this balance is held on.
- `agentAddress` - a string containing the account address the balance is held for.
- `amount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) value of the current balance.
- `withdrawableAt` - the earliest block time at which `MsgWithdrawPerformanceDeposit` may be issued against this balance. Set on each `MsgAddPerformanceDeposit` to `max(current, now + collection.minDepositPeriod)`. Nil / zero on balances created under a collection with `minDepositPeriod == 0` (no lock). The slash path is not gated by this lock.

Lifecycle:

- **Created** — first `MsgAddPerformanceDeposit` for an `(collection, agent)` pair. Funds move from agent's wallet → collection escrow. Emits `AgentDepositBalanceCreatedEvent`.
- **Updated** — subsequent top-ups, partial withdrawals, or partial slashes on adjudicated dispute loss that leave a non-zero balance. Emits `AgentDepositBalanceUpdatedEvent`.
- **Removed** — full withdrawal or a slash that drains the balance to zero. The KV entry is deleted. Emits `AgentDepositBalanceRemovedEvent` with the final zero-amount balance.

`amount.Sort()` is always non-empty when the entry exists (no zero-amount denoms persist).

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

Defines the `status` of a [Evaluation](#evaluation) indicating the status and result of the evaluation. `APPROVED`, `REJECTED`, `DISPUTED`, and `INVALIDATED` are terminal — once a claim's `Evaluation` carries one of those, the claim is locked and cannot be re-evaluated. `FLAGGED` is non-terminal and explicitly designed to be re-evaluated to one of the terminal statuses (or re-flagged by a different agent).

```go
var EvaluationStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "REJECTED",
	3: "DISPUTED",
	4: "INVALIDATED",
	5: "FLAGGED"
}
```

`FLAGGED` semantics:

- **No payments fire.** Neither evaluator payment nor any approval / rejection payout is triggered.
- **Funds remain in escrow** if the claim was intent-funded — the intent escrow stays locked until a subsequent terminal evaluation moves it to the agent (`APPROVED`) or refunds the approval account (any other terminal status).
- **AgentQuota is consumed** the same as for any terminal evaluation — flagging counts against the evaluator's quota so an oracle cannot flag-bomb unboundedly. Only `INVALIDATED` skips quota decrement.
- **Re-evaluation is allowed.** A claim with `FLAGGED` status can be re-evaluated by:
  - the same agent that flagged (e.g. they later got more information and want to finalise), or
  - any other authorized evaluator (escalation to a human reviewer or a stricter oracle).
- **Re-flagging by the same agent is blocked.** An agent that flagged this claim — whether their flag is the current evaluation or sits in `evaluationHistory` after intervening flags from other agents — cannot flag the same claim again. The check covers both surfaces; flag-bombing across an intervening flag from another evaluator is rejected. Returns `ErrSelfReFlag`.
- **Counters.** On flag, `Collection.flagged++`. On the *first* transition into `FLAGGED` for a claim, `Collection.flaggedActive++`. On terminal finalisation of a flagged claim, `Collection.flaggedActive--` (with appropriate `Approved++` / `Rejected++` / `Invalidated++` increment for the terminal status). `Disputed++` does not apply because `DISPUTED` is recorded by `MsgDisputeClaim`, not by `MsgEvaluateClaim`.

**`DISPUTED` is deprecated for new transactions**. `MsgEvaluateClaim` rejects status `DISPUTED` with `ErrEvaluationStatusDisputedDeprecated`. The dispute lifecycle lives on the [Dispute](#dispute) record, not on the evaluation — existing on-chain evaluations with status `DISPUTED` from before v7 remain valid history.

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

### DisputeTargetRole

Defines which party of a claim a [Dispute](#dispute) is filed against. A dispute targets **exactly one** role; to dispute both the submitter and the evaluator of the same claim, file two separate disputes.

```go
var DisputeTargetRole_name = map[int32]string{
	0: "UNSPECIFIED", // Only appears on pre-v7 migrated disputes; rejected on new txs.
	1: "SUBMITTER",   // The service agent that submitted the claim.
	2: "EVALUATOR",   // The evaluation agent that evaluated the claim.
}
```

### DisputeStatus

The lifecycle state of a [Dispute](#dispute).

```go
var DisputeStatus_name = map[int32]string{
	0: "OPEN",      // Filed and awaiting adjudication. Targeted agent is blocked from new submissions / evaluations / withdrawal.
	1: "AWARDED",   // Disputer won. Loser's balance slashed (80% disputer / 20% adjudicator). Further disputes on the same (subject, role) are permanently blocked.
	2: "DISMISSED", // Disputer lost. Their dispute deposit becomes the pot (same 80/20 split, target agent is the winner). New disputes against the same (subject, role) may be filed with new evidence.
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
	CW1155Payment         []CW1155Payment
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
- `cw1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the CW1155 tokens and amount to be paid
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
	MaxCW1155Payment  []CW1155Payment
	IntentDurationNs  time.Duration
	MemberAddress     string
}
```

The field's descriptions is as follows:

- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz)
- `maxAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the maximum amount allowed to be specified by service agent for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `maxCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payments allowed to be specified by service agent for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `maxCW1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the maximum CW1155 payments allowed to be specified by service agent for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `intentDurationNs` - a duration for which the intent is active, after which it will expire (in nanoseconds)
- `memberAddress` - a string containing the team member who created this constraint via [MsgCreateClaimAuthorization](03_messages.md#msgcreateclaimauthorization). Empty for individual (non-team) subscriptions. The intent handler matches constraints by `(collectionId, memberAddress)` using strict equality (both empty for individual; both equal for team), so when multiple team members authorize the same oracle on the same collection, each member's intents consume their own constraint and `agentQuota` independently.

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
	MaxCustomCW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `claimIds` - a list of strings containing all the id's of the claimsthe grantee is allowed to evaluate, can be an empty list to allow any claim
- `collectionId` - a string containing the Collection `id` the constraints is for
- `agentQuota` - a integer containing the quota for amount of time the grantee can execute the given authorization(authz), note: it won't subtract one on evaluation if agent evaluates claim with status `invalidated`. All other statuses — `approved`, `rejected`, `disputed`, and `flagged` — consume one quota slot per call; an evaluator that flags a claim and later finalises it will burn two slots.
- `beforeDate` - a timestamp of the date after which the grantee can't execute this authz anymore, a cut off date
- `MaxCustomAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the coins and amount that indicates the maximum the evaluator is allowed to change the `APPROVED` payout to, since claims can be made for specific amount an evaluator is allowed to change the `APPROVED` payout amount.
- `MaxCustomCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payments allowed to be specified by evaluator for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.
- `MaxCustomCW1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the maximum CW1155 payments allowed to be specified by evaluator for claim approval. If empty then no custom amount is allowed, and default payments from Collection payments are used.

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
	MaxCW1155Payment     []CW1155Payment
	Expiration           *time.Time
	CollectionIds        []string
	AllowedAuthTypes     CreateClaimAuthorizationType
	MaxIntentDurationNs  time.Duration
	MemberAddress        string
}
```

The field's descriptions is as follows:

- `maxAuthorizations` - a integer containing the maximum number of authorizations that can be created through this meta-authorization. 0 means no quota.
- `maxAgentQuota` - a integer containing the maximum quota that can be set in created authorizations. 0 means no maximum quota per authorization.
- `maxAmount` - a [Coins](https://github.com/cosmos/cosmos-sdk/blob/main/types/coin.go#L180) object which denotes the maximum amount that can be set in created authorizations. If empty then any custom amount is allowed in the created authorizations. Explicitly set to 0 to disallow any custom amount in the created authorizations.
- `maxCW20Payment` - an array of [CW20Payment](#cw20payment) containing the maximum CW20 payment that can be set in created authorizations. If empty then any CW20 payment is allowed in the created authorizations. Explicitly set to 0 to disallow any CW20 payment in the created authorizations.
- `maxCW1155Payment` - an array of [CW1155Payment](#cw1155payment) containing the maximum CW1155 payment that can be set in created authorizations. If empty then any CW1155 payment is allowed in the created authorizations. Explicitly set to 0 to disallow any CW1155 payment in the created authorizations.
- `expiration` - a timestamp of the expiration of this meta-authorization(specific constraint). If not set then no expiration.
- `collectionIds` - a list of strings containing the Collection IDs the grantee can create authorizations for. If empty then all collections for the admin are allowed.
- `allowedAuthTypes` - a [CreateClaimAuthorizationType](#createclaimauthorizationtype) indicating the types of authorizations the grantee can create (submit, evaluate, or all/both).
- `maxIntentDurationNs` - a duration containing the maximum intent duration for the authorization allowed (for submit).
- `memberAddress` - a string containing the team member identity locked into this constraint by the admin at grant creation time. The Accept method enforces strict equality between this and the `MsgCreateClaimAuthorization.memberAddress` field — both empty for individual subscriptions, both equal to the same member address for team members. This is the **anti-spoofing primitive**: a grantee cannot create an authorization tagged with a member address other than the one the admin authorized.
