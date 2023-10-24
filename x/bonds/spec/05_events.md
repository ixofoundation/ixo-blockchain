# Events

The bonds module emits the following typed events:

### BondCreatedEvent

Emitted after a successfull `MsgCreateBond`

| **Field** | **Type** | **Description** |
| :-------- | :------- | :-------------- |
| Bond      | `*Bond`  |                 |

```go
type BondCreatedEvent struct {
	Bond *Bond
}
```

### BondUpdatedEvent

Emitted after a successfull `MsgEditBond`, `MsgUpdateBondState`

| **Field** | **Type** | **Description** |
| :-------- | :------- | :-------------- |
| Bond      | `*Bond`  |                 |

```go
type BondUpdatedEvent struct {
	Bond *Bond
}
```

### BondSetNextAlphaEvent

Emitted after a successfull `MsgSetNextAlpha`. Note this doesn't mean the alpha has been updated, it has only been added to the batch for next batch update.

| **Field** | **Type** | **Description** |
| :-------- | :------- | :-------------- |
| BondDid   | `string` |                 |
| NextAlpha | `string` |                 |
| Signer    | `string` |                 |

```go
type BondSetNextAlphaEvent struct {
	BondDid   string
	NextAlpha string
	Signer    string
}
```

### BondBuyOrderEvent

Emitted after a successfull `MsgBuy`. Note this doesn't mean the buy has been executed, it has only been added to the batch.

| **Field** | **Type**    | **Description** |
| :-------- | :---------- | :-------------- |
| BondDid   | `string`    |                 |
| Order     | `*BuyOrder` |                 |

```go
type BondBuyOrderEvent struct {
	Order   *BuyOrder
	BondDid string
}
```

### BondSellOrderEvent

Emitted after a successfull `MsgSell`. Note this doesn't mean the sell has been executed, it has only been added to the batch.

| **Field** | **Type**     | **Description** |
| :-------- | :----------- | :-------------- |
| BondDid   | `string`     |                 |
| Order     | `*SellOrder` |                 |

```go
type BondSellOrderEvent struct {
	Order   *SellOrder
	BondDid string
}
```

### BondSwapOrderEvent

Emitted after a successfull `MsgSwap`. Note this doesn't mean the swap has been executed, it has only been added to the batch.

| **Field** | **Type**     | **Description** |
| :-------- | :----------- | :-------------- |
| BondDid   | `string`     |                 |
| Order     | `*SwapOrder` |                 |

```go
type BondSwapOrderEvent struct {
	Order   *SwapOrder
	BondDid string
}
```

### BondMakeOutcomePaymentEvent

Emitted after a successfull `MsgMakeOutcomePayment`

| **Field**      | **Type**                                   | **Description** |
| :------------- | :----------------------------------------- | :-------------- |
| BondDid        | `string`                                   |                 |
| OutcomePayment | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| SenderDid      | `string`                                   |                 |
| SenderAddress  | `string`                                   |                 |

```go
type BondMakeOutcomePaymentEvent struct {
	BondDid        string
	OutcomePayment github_com_cosmos_cosmos_sdk_types.Coins
	SenderDid      string
	SenderAddress  string
}
```

### BondWithdrawShareEvent

Emitted after a successfull `MsgWithdrawShare`

| **Field**        | **Type**                                   | **Description** |
| :--------------- | :----------------------------------------- | :-------------- |
| BondDid          | `string`                                   |                 |
| WithdrawPayment  | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| RecipientDid     | `string`                                   |                 |
| RecipientAddress | `string`                                   |                 |

```go
type BondWithdrawShareEvent struct {
	BondDid          string
	WithdrawPayment  github_com_cosmos_cosmos_sdk_types.Coins
	RecipientDid     string
	RecipientAddress string
}
```

### BondWithdrawReserveEvent

Emitted after a successfull `MsgWithdrawReserve`

| **Field**                | **Type**                                   | **Description** |
| :----------------------- | :----------------------------------------- | :-------------- |
| BondDid                  | `string`                                   |                 |
| WithdrawAmount           | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| WithdrawerDid            | `string`                                   |                 |
| WithdrawerAddress        | `string`                                   |                 |
| ReserveWithdrawalAddress | `string`                                   |                 |

```go
type BondWithdrawReserveEvent struct {
	BondDid                  string
	WithdrawAmount           github_com_cosmos_cosmos_sdk_types.Coins
	WithdrawerDid            string
	WithdrawerAddress        string
	ReserveWithdrawalAddress string
}
```

### BondEditAlphaSuccessEvent

Emitted after a successfull update of the Alpha state which is done in `EndBlock`

| **Field**   | **Type** | **Description** |
| :---------- | :------- | :-------------- |
| BondDid     | `string` |                 |
| Token       | `string` |                 |
| PublicAlpha | `string` |                 |
| SystemAlpha | `string` |                 |

```go
type BondEditAlphaSuccessEvent struct {
	BondDid     string
	Token       string
	PublicAlpha string
	SystemAlpha string
}
```

### BondEditAlphaFailedEvent

Emitted if update of the Alpha state failed which is done in `EndBlock`

| **Field**    | **Type** | **Description** |
| :----------- | :------- | :-------------- |
| BondDid      | `string` |                 |
| Token        | `string` |                 |
| CancelReason | `string` |                 |

```go
type BondEditAlphaFailedEvent struct {
	BondDid     string
	Token       string
	CancelReason string
}
```

### BondBuyOrderFulfilledEvent

Emitted after a successfull `BuyOrder` has been executed which is done in `EndBlock`

| **Field**                   | **Type**                                   | **Description** |
| :-------------------------- | :----------------------------------------- | :-------------- |
| BondDid                     | `string`                                   |                 |
| Order                       | `*BuyOrder`                                |                 |
| ChargedPrices               | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| ChargedFees                 | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| ReturnedToAddress           | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| NewBondTokenBalance         | `github_com_cosmos_cosmos_sdk_types.Int`   |                 |
| ChargedPricesOfWhichReserve | `*github_com_cosmos_cosmos_sdk_types.Int`  |                 |
| ChargedPricesOfWhichFunding | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |

```go
type BondBuyOrderFulfilledEvent struct {
	BondDid                     string
	Order                       *BuyOrder
	ChargedPrices               github_com_cosmos_cosmos_sdk_types.Coins
	ChargedFees                 github_com_cosmos_cosmos_sdk_types.Coins
	ReturnedToAddress           github_com_cosmos_cosmos_sdk_types.Coins
	NewBondTokenBalance         github_com_cosmos_cosmos_sdk_types.Int
	ChargedPricesOfWhichReserve *github_com_cosmos_cosmos_sdk_types.Int
	ChargedPricesOfWhichFunding github_com_cosmos_cosmos_sdk_types.Coins
}
```

### BondSellOrderFulfilledEvent

Emitted after a successfull `SellOrder` has been executed which is done in `EndBlock`

| **Field**           | **Type**                                   | **Description** |
| :------------------ | :----------------------------------------- | :-------------- |
| BondDid             | `string`                                   |                 |
| Order               | `*SellOrder`                               |                 |
| ChargedFees         | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| ReturnedToAddress   | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| NewBondTokenBalance | `github_com_cosmos_cosmos_sdk_types.Int`   |                 |

```go
type BondSellOrderFulfilledEvent struct {
	BondDid             string
	Order               *SellOrder
	ChargedFees         github_com_cosmos_cosmos_sdk_types.Coins
	ReturnedToAddress   github_com_cosmos_cosmos_sdk_types.Coins
	NewBondTokenBalance github_com_cosmos_cosmos_sdk_types.Int
}
```

### BondSwapOrderFulfilledEvent

Emitted after a successfull `SwapOrder` has been executed which is done in `EndBlock`

| **Field**         | **Type**                                   | **Description** |
| :---------------- | :----------------------------------------- | :-------------- |
| BondDid           | `string`                                   |                 |
| Order             | `*SwapOrder`                               |                 |
| ChargedFee        | `types.Coin`                               |                 |
| ReturnedToAddress | `github_com_cosmos_cosmos_sdk_types.Coins` |                 |
| TokensSwapped     | `types.Coin`                               |                 |

```go
type BondSwapOrderFulfilledEvent struct {
	BondDid           string
	Order             *SwapOrder
	ChargedFee        types.Coin
	ReturnedToAddress github_com_cosmos_cosmos_sdk_types.Coins
	TokensSwapped     types.Coin
}
```

### BondBuyOrderCancelledEvent

Emitted when a `BuyOrder` has been cancelled as it is not eligible to be executed anymore.

| **Field** | **Type**    | **Description** |
| :-------- | :---------- | :-------------- |
| BondDid   | `string`    |                 |
| Order     | `*BuyOrder` |                 |

```go
type BondBuyOrderCancelledEvent struct {
	BondDid string
	Order   *BuyOrder
}
```
