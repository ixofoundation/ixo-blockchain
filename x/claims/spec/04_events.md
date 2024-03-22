# Events

The bonds module emits the following typed events:

### CollectionCreatedEvent

Emitted after a successfull `MsgCreateCollection`

```go
type CollectionCreatedEvent struct {
	Collection *Collection
}
```

The field's descriptions is as follows:

- `collection` - the full [Collection](02_state.md#collection)

### CollectionUpdatedEvent

Emitted after a successfull collection update event or `MsgSubmitClaim` and `EvaluateClaim` since collection holds a count of claims

```go
type CollectionUpdatedEvent struct {
	Collection *Collection
}
```

The field's descriptions is as follows:

- `collection` - the full [Collection](02_state.md#collection)

### ClaimSubmittedEvent

Emitted after a successfull `MsgSubmitClaim`

```go
type ClaimSubmittedEvent struct {
	Claim *Claim
}
```

The field's descriptions is as follows:

- `claim` - the full [Claim](02_state.md#claim)

### ClaimUpdatedEvent

Emitted after a successfull `MsgEvaluateClaim` or when the state for the `ClaimPayments` changes.

```go
type ClaimUpdatedEvent struct {
	Claim *Claim
}
```

The field's descriptions is as follows:

- `claim` - the full [Claim](02_state.md#claim)

### ClaimEvaluatedEvent

Emitted after a successfull `MsgEvaluateClaim`

```go
type ClaimEvaluatedEvent struct {
	Evaluation *Evaluation
}
```

The field's descriptions is as follows:

- `evaluation` - the full [Evaluation](02_state.md#evaluation)

### ClaimDisputedEvent

Emitted after a successfull `MsgDisputeClaim`

```go
type ClaimDisputedEvent struct {
	Dispute *Dispute
}
```

The field's descriptions is as follows:

- `dispute` - the full [Dispute](02_state.md#dispute)

### PaymentWithdrawnEvent

Emitted after a successfull `MsgWithdrawPayment` or when a payment gets withdrawn through another trigger like payments on evaluation with 0 timeout.

```go
type PaymentWithdrawnEvent struct {
	Withdraw *WithdrawPaymentConstraints
}
```

The field's descriptions is as follows:

- `withdraw` - the full [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)

### PaymentWithdrawCreatedEvent

Emitted after a successfull payment withdrawal authz gets created on chain if the `Payments` `timeout_ns` is not 0, and a [WithdrawPaymentAuthorization](02_state.md#withdrawpaymentauthorization) is created with the receiver as the grantee

```go
type PaymentWithdrawCreatedEvent struct {
	Withdraw *WithdrawPaymentConstraints
}
```

The field's descriptions is as follows:

- `withdraw` - the full [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)
