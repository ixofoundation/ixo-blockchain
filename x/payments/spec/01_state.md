# State

The payments module stores three lists of the following three types of data:

1. [Payment templates](#payment-templates)
1. [Payment contracts](#payment-contracts)
1. [Subscriptions](#subscriptions)

## Payment Templates

```go
type PaymentTemplate struct {
	Id             string
	PaymentAmount  sdk.Coins
	PaymentMinimum sdk.Coins
	PaymentMaximum sdk.Coins
	Discounts      []Discount
}
```

A payment template contains basic details about a payment, without any info
about the payer or the payee. Apart from an ID identifying the template, the
template stores the payment amount, a minimum and maximum payment, and a list of
discounts.

**Minimum payment**: if this is higher than the payment amount, the first
payment effected is increased up to the minimum amount. In other words, this
minimum specifies a floor for the cumulative payment amount.

**Maximum payment**: this specifies a ceiling for the cumulative payment amount.
If the cumulative payment amount exceeds this value, the amount is decreased
down to the maximum amount, so that the payer never pays more than the max.

Together, these two values enforce a range of payment value that ensures the
payer never pays less or more than the limits, in total.

**Discounts**: a list of discounts, identified by an ID, each of which have a
percentage value. This allows payment contracts (further on) to grant discounts
to certain payers by specifying the ID of the discount in the contract.

```go
type Discounts []Discount

type Discount struct {
	Id      sdk.Uint
	Percent sdk.Dec
}
```

## Payment Contracts

```go
type PaymentContract struct {
	Id                string
	PaymentTemplateId string
	Creator           string
	Payer             string
	Recipients        []DistributionShare
	CumulativePay     sdk.Coins
	CurrentRemainder  sdk.Coins
	CanDeauthorise    bool
	Authorised        bool
	DiscountId        sdk.Uint
}
```

A payment contract is a concrete agreement between a payer and payee(s) which
can be invoked once or multiple times to effect payment(s). Like the payment
template, it has an ID, but it also points to a payment template by its ID. This
means that contracts build upon the information provided in templates.

The contract identifies the contract creator, payer, and recipients:

- **Creator**: creator of a payment contract is able to grant and revoke
  discounts to payers, invoke the contract to effect payment, and create a
  subscription agreement based on the payment contract (discussed later on).
- **Payer**: address from which tokens for payment are subtracted. If the payer
  does not have enough tokens, then the contract payment cannot be effected.
- **Recipients**: a list of distribution shares, which specifies how the payment
  from the payer will be split and to which addresses. The share percentages add
  up to 100% (i.e. all of the payment is distributed).

```go
type Distribution []DistributionShare

type DistributionShare struct {
	Address    string
	Percentage sdk.Dec
}
```

The contract keeps track of the cumulative payment and a remainder:

- **CumulativePay**: since a contract can be invoked more than once, this value
  keeps track of the total value paid via this contract. This value is compared
  with the payment template's minimum and maximum to ensure that the payment
  amount always stays within these limits.
- **CurrentRemainder**: any remainder after all payees have been paid, due to
  rounding when multiplying percentages with pay amount. This remainder is
  discounted from any subsequent payment made by the payer.

In terms of authorisation, the contract stores two values:

- **CanDeauthorise**: indicates whether or not a contract can be de-authorised
  once it has been authorised. This is specified by the creator of the contract.
  The payer should pay attention since if this value is _False_, then the payer
  cannot decide to halt any payments due to this contract in the future.
- **Authorised**: by default a contract is not authorised (otherwise the
  contract creator can just steal money from any address). The contract's payer
  needs to explicitly authorise the contract in a subsequent transaction.

Lastly, the contract can point to a discount by ID, as specified in the payment
template:

- **DiscountId**: discount specified by the contract creator at creation, but
  can also be granted and revoked at any point by the contract creator. The
  discount can also be swapped for any another discount specified in the
  template. Any discount is applied when the contract payment is being effected.

## Subscriptions

```go
type Subscription struct {
	Id                 string
	PaymentContractId  string
	PeriodsSoFar       sdk.Uint
	MaxPeriods         sdk.Uint
	PeriodsAccumulated sdk.Uint
	Period             *sdk.Any
}
```

Sometimes we want payments to be effected periodically. Subscriptions build on
top of payment contracts (by linking to a payment contract ID), adding info that
describes a periodic payment. All active subscriptions are checked at the end of
each block and subscriptions for which enough time has passed are invoked.

- **PeriodsSoFar**: just keeps track of the number of periods that have elapsed
  (i.e. how many payments should have been effected). This is initialised to 0.
- **MaxPeriods**: places a limit on the number of periods, otherwise the
  subscription will go on forever, effecting payment with each period. However,
  the minimum and maximum payment specified by the payment template are still
  enforced and are able to override this max periods amount. Similarly, the max
  periods can override the minimum and maximum if it is more restrictive.
- **PeriodsAccumulated**: number of unpaid periods. If the subscription payment
  fails to get effected because of insufficient funds, the `PeriodsSoFar` value
  is still updated, but we also increment the `PeriodsAccumulated` value. The
  subscription only revisits these unpaid periods once the maximum number of
  periods has been reached.
- **Period**: describes the start of the subscription and the length of each
  period. It can either be a `BlockPeriod` (period specified in terms of blocks)
  or a `TimePeriod` (period specified in terms of time).

```go
type Period interface {
	proto.Message 
	
	GetPeriodUnit() string
	Validate() error
	periodStarted(ctx sdk.Context) bool
	periodEnded(ctx sdk.Context) bool
	nextPeriod() Period
}

type BlockPeriod struct {
	PeriodLength     int64 `json:"period_length" yaml:"period_length"`
	PeriodStartBlock int64 `json:"period_start_block" yaml:"period_start_block"`
}

type TimePeriod struct {
	PeriodDurationNs time.Duration `json:"period_duration_ns" yaml:"period_duration_ns"`
	PeriodStartTime  time.Time     `json:"period_start_time" yaml:"period_start_time"`
}
```
