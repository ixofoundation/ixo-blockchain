# Events

In this section we describe the processing of the claims events. All events here are listed so SDK clients can build adequate user interface to reflect the results of executing a Message.

### CollectionCreatedEvent

This event is emitted when a [Collection](./02_state.md#collection) is created using the [MsgCreateCollection](./03_messages.md#msgcreatecollection) message.

```go
type CollectionCreatedEvent struct {
	Collection *Collection
}
```

The field's descriptions is as follows:

- `collection` - the full [Collection](02_state.md#collection)

### CollectionUpdatedEvent

This event is emitted when a [Collection](./02_state.md#collection) is updated using the [MsgUpdateCollectionState](./03_messages.md#msgupdatecollectionstate), [MsgUpdateCollectionDates](./03_messages.md#msgupdatecollectiondates), [MsgUpdateCollectionPayments](./03_messages.md#msgupdatecollectionpayments), or [MsgUpdateCollectionIntents](./03_messages.md#msgupdatecollectionintents) message.

```go
type CollectionUpdatedEvent struct {
	Collection *Collection
}
```

The field's descriptions is as follows:

- `collection` - the full [Collection](02_state.md#collection)

### ClaimSubmittedEvent

This event is emitted when a [Claim](./02_state.md#claim) is submitted using the [MsgSubmitClaim](./03_messages.md#msgsubmitclaim) message.

```go
type ClaimSubmittedEvent struct {
	Claim *Claim
}
```

The field's descriptions is as follows:

- `claim` - the full [Claim](02_state.md#claim)

### ClaimUpdatedEvent

This event is emitted when a [Claim](./02_state.md#claim) is updated, typically when the payment status changes or when the claim is evaluated (including a `flagged` evaluation or a re-evaluation that transitions a flagged claim out of `FLAGGED`). On a re-evaluation the carried `Claim` includes the updated `evaluationHistory` (the prior evaluation appended) and the new `evaluation`.

```go
type ClaimUpdatedEvent struct {
	Claim *Claim
}
```

The field's descriptions is as follows:

- `claim` - the full [Claim](02_state.md#claim)

### ClaimEvaluatedEvent

This event is emitted when a [Claim](./02_state.md#claim) is Evaluated using the [MsgEvaluateClaim](./03_messages.md#msgevaluateclaim) message. Fires for every successful call to that message — including evaluations with status `flagged` (non-terminal) and re-evaluations that transition a flagged claim to a terminal status. Indexers can distinguish flag events from terminal events by inspecting `evaluation.status`.

```go
type ClaimEvaluatedEvent struct {
	Evaluation *Evaluation
}
```

The field's descriptions is as follows:

- `evaluation` - the full [Evaluation](02_state.md#evaluation)

### ClaimDisputedEvent

This event is emitted when a [Claim](./02_state.md#claim) is disputed using the [MsgDisputeClaim](./03_messages.md#msgdisputeclaim) message.

```go
type ClaimDisputedEvent struct {
	Dispute *Dispute
}
```

The field's descriptions is as follows:

- `dispute` - the full [Dispute](02_state.md#dispute)

### PaymentWithdrawnEvent

This event is emitted when a payment is withdrawn using the [MsgWithdrawPayment](./03_messages.md#msgwithdrawpayment) message.

```go
type PaymentWithdrawnEvent struct {
	Withdraw *WithdrawPaymentConstraints
  CW20Output []CW20Output
  CW1155Payment []CW1155Payment
}
```

The field's descriptions is as follows:

- `withdraw` - the full [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)
- `cw20Output` - an array of [CW20Output](02_state.md#CW20Output) that defines the CW20Payment payment split outputs if any.
- `cw1155Payment` - an array of [CW1155Payment](02_state.md#cw1155payment) that defines the CW1155Payment payments to know amount per token id.

### PaymentWithdrawCreatedEvent

Emitted after a successful payment withdrawal authz gets created on chain if the `Payments` `timeout_ns` is not 0, and a [WithdrawPaymentAuthorization](02_state.md#withdrawpaymentauthorization) is created with the receiver as the grantee

```go
type PaymentWithdrawCreatedEvent struct {
  Withdraw *WithdrawPaymentConstraints
}
```

The field's descriptions is as follows:

- `withdraw` - the full [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)

## IntentSubmittedEvent

This event is emitted when an [Intent](./02_state.md#intent) is created using the [MsgClaimIntent](./03_messages.md#msgclaimintent) message.

```go
type IntentSubmittedEvent struct {
	Intent *Intent
}
```

The field's descriptions is as follows:

- `intent` - a full [Intent](./02_state.md#intent)

## IntentUpdatedEvent

This event is emitted when an [Intent](./02_state.md#intent) is updated by submitting a claim with the intent and it's status changes or when an intent
expires and is cleaned up, it then also gets a status change to EXPIRED.

```go
type IntentUpdatedEvent struct {
	Intent *Intent
}
```

The field's descriptions is as follows:

- `intent` - a full [Intent](./02_state.md#intent)

## ClaimAuthorizationCreatedEvent

This event is emitted when a claim authorization is created using the [MsgCreateClaimAuthorization](./03_messages.md#msgcreateclaimauthorization) message.

```go
type ClaimAuthorizationCreatedEvent struct {
  Creator string
  CreatorDid string
	Grantee string
	Admin string
	CollectionId string
	AuthType CreateClaimAuthorizationType
}
```

The field's descriptions is as follows:

- `creator` - a string containing the address of the creator who executed the meta-authorization and created the SubmitClaimAuthorization/EvaluateClaimAuthorization
- `creatorDid` - a string containing the did of the creator as defined above
- `grantee` - a string containing the address of the account receiving the authorization
- `admin` - a string containing the address of the admin account on the collection from which the authz is given through the meta-authorization
- `collectionId` - a string containing the id of the collection the authorization applies to
- `authorizationType` - a [CreateClaimAuthorizationType](./02_state.md#createclaimauthorizationtype) indicating the type of authorization created

## MemberBudgetCreatedEvent

This event is emitted when a [MemberBudget](./02_state.md#memberbudget) is added to a collection for the first time via [MsgSetCollectionMembers](./03_messages.md#msgsetcollectionmembers). Carries the full initial budget state so an indexer can insert a fresh row without re-querying the chain.

```go
type MemberBudgetCreatedEvent struct {
	Budget *MemberBudget
}
```

The field's descriptions is as follows:

- `budget` - the full [MemberBudget](./02_state.md#memberbudget)

## MemberBudgetUpdatedEvent

This event is emitted on **every state change** to an existing [MemberBudget](./02_state.md#memberbudget). Carries the post-update full budget state so an indexer can `UPDATE` the corresponding row directly.

Triggered by:

- An admin update to an existing member's limits via [MsgSetCollectionMembers](./03_messages.md#msgsetcollectionmembers) (preserving or resetting `periodSpent` per `resetPeriodSpent`).
- A `periodSpent` deduction during [MsgClaimIntent](./03_messages.md#msgclaimintent) (covers both the lazy period reset and the deduction in one event).
- A `periodSpent` restoration on claim rejection / dispute / invalidation via [MsgEvaluateClaim](./03_messages.md#msgevaluateclaim). Note: a `flagged` evaluation does **not** restore the budget — the spend stays held until the claim reaches a terminal evaluation status (the budget is restored at that point if the terminal status is non-approved).
- A `periodSpent` restoration on intent expiration in the EndBlocker.
- A lazy period reset triggered during a restore operation when the period had elapsed (the restore early-returns after persisting the reset state).

```go
type MemberBudgetUpdatedEvent struct {
	Budget *MemberBudget
}
```

The field's descriptions is as follows:

- `budget` - the full updated [MemberBudget](./02_state.md#memberbudget)

## MemberBudgetRemovedEvent

This event is emitted when a [MemberBudget](./02_state.md#memberbudget) is removed via [MsgRemoveCollectionMembers](./03_messages.md#msgremovecollectionmembers). Carries the final budget state at time of removal so an indexer can audit / archive the final values before deleting (or marking removed) the row.

```go
type MemberBudgetRemovedEvent struct {
	Budget *MemberBudget
}
```

The field's descriptions is as follows:

- `budget` - the full [MemberBudget](./02_state.md#memberbudget) state at time of removal
