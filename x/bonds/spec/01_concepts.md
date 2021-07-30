# Concepts

## Token Bonding Curves

Bonding curves are continuous liquidity mechanisms which are used in market design for cryptographically-supported token economies. Tokens are atomic units of state information which are cryptographically verifiable in peer-to-peer networks. Bonding curves are an example of an enforceable mechanism through which participating agents influence this state. By designing such mechanisms, an engineer may establish the topological structure of a token economy without presupposing the utilities or associated actions of the agents within that economy.

Token Bonding Curves are therefore an important crypto-economic mechanism for building a wide range of capabilities directly into decentralised applications. They can function simultaneously as means of decentralised capital formation, liquidity provision and autonomous market maker.
Bonding curves are powerful tools because the tokens they issue can represent rights - including
* rights of access
* rights of use
* rights of ownership, and 
* voting rights. 

In the case of continuous organizations, tokens issued through bonding curves embody rights to the future revenues of a startup. 
In the augmented bonding curve, tokens can embody the rights to govern how funds are spent by a not-for-profit organization. 
In an Alpha-Bond, tokens can give holders the rights to future outcomes payments and performance incentive bonuses.

## Token Bonds Module

The Token Bonds Cosmos SDK Module enables applications that use token bonding curves to be created on-the-fly. 
Each new Token B instance declares a new token denomination in the application, with a set of parameters.
The module stores the current state of all tokens that have been created using this module.
Changes in state occur through transactions that are instructed by valid *buy, sell, and swap* messages.

**Buy** instructions cause bond tokens to be minted during a state transition. This increases the total supply balance of tokens.
**Sell** instructions burn bond tokens during a state transition that decreases the total supply balance of tokens.
**Limits** are set for the maximum numbers of bond tokens that can exist at any point in state.

Bond Tokens trade against pairings in their Bond Reserve.

The bonding curve forms an interface between the reserve token quantity and the bond token price (in the Reserve currency).

Bonding curves are defined by their mathematical properties. This is determined by the type of curve function and by the function parameters that are set. Generally these parameters are chosen to best-fit empiricially-observed market dynamics of supply and demand. 
External parameters, such as market supply and demand, are complex and typically hard to predict. 

*****

Pricing is defined by the function type and function parameters, which can define either the pricing function of the bond as a function of the supply, or simply indicate that the bond is a token swapper, where pricing is instead defined by the first buyer and any swaps performed thereafter.

A bond may also specify non-zero fees, which are calculated based on the size of an order and sent to the specified fee address, order quantity limits to limit the size of orders, disable the ability to sell tokens, specify multiple signers that will need to sign for any editing of the bond details, and in the case of swapper bonds, sanity values to set a range of valid exchange rate between the two reserve tokens. Lastly, a bond has a string state value, which in most cases is _open_, but in certain function types it has more meaning, such as for augmented bonding curves, in which case it can be _open_ \[for open phase\] and _hatch_ \[for hatch phase\]. This state is _not_ specified by the creator during bond creation.

```go
type Bond struct {
    Token                        string                                   
    Name                         string                                   
    Description                  string                                   
    CreatorDid                   string                                   
    ControllerDid                string                                   
    FunctionType                 string                                   
    FunctionParameters           FunctionParams                           
    ReserveTokens                []string                                 
    TxFeePercentage              sdk.Dec   
    ExitFeePercentage            sdk.Dec   
    FeeAddress                   string
	ReserveWithdrawalAddress     string
    MaxSupply                    sdk.Coin                               
    OrderQuantityLimits          sdk.Coins 
    SanityRate                   sdk.Dec   
    SanityMarginPercentage       sdk.Dec   
    CurrentSupply                sdk.Coin                               
    CurrentReserve               sdk.Coins
    AvailableReserve             sdk.Coins
    CurrentOutcomePaymentReserve sdk.Coins 
    AllowSells                   bool
    AllowReserveWithdrawals      bool
    AlphaBond                    bool                                     
    BatchBlocks                  sdk.Uint  
    OutcomePayment               sdk.Int   
    State                        string                                   
    BondDid                      string                                   
}
```

## Batching

For each bond, a single corresponding batch holds a collection of outstanding buy, sell, and swap orders. The lifespan of a batch, in terms of the number of blocks, is defined in the corresponding bond (`BatchBlocks`).

Orders can be added to the current batch at any point in time. Any order that is not cancelled by the end of the batch's lifespan is eligible to get fulfilled. Otherwise, the order is discarded and any actions that were already performed are reverted.

The primary task of the batching mechanism is to find a common price for all of the buys and sells submitted to the batch by summing up all of the buys and sells, thus ignoring their order, and matching-up the total buy and sell amounts to give balanced and fair global buy and sell prices.

For alpha bonds, the batch also stores the next alpha value, if it was changed throughout the lifetime of the batch. If alpha has not changed, then it will show up as `-1`.

```go
type Batch struct {
	BondDid         string
	BlocksRemaining sdk.Uint
	NextPublicAlpha sdk.Dec
	TotalBuyAmount  sdk.Coin
	TotalSellAmount sdk.Coin
	BuyPrices       sdk.DecCoins
	SellPrices      sdk.DecCoins
	Buys            []BuyOrder
	Sells           []SellOrder
	Swaps           []SwapOrder
}
```
