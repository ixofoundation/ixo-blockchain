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

This event is emitted when a [Claim](./02_state.md#claim) is updated, typically when the payment status changes.

```go
type ClaimUpdatedEvent struct {
	Claim *Claim
}
```

The field's descriptions is as follows:

- `claim` - the full [Claim](02_state.md#claim)

### ClaimEvaluatedEvent

This event is emitted when a [Claim](./02_state.md#claim) is Evaluated using the [MsgEvaluateClaim](./03_messages.md#msgevaluateclaim) message.

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
}
```

The field's descriptions is as follows:

- `withdraw` - the full [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)
- `cw20Output` - an array of [CW20Output](02_state.md#CW20Output) that defines the CW20Payment payment split outputs if any.

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
