# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [bonds/bonds.proto](#bonds/bonds.proto)
    - [BaseOrder](#bonds.BaseOrder)
    - [Batch](#bonds.Batch)
    - [Bond](#bonds.Bond)
    - [BondDetails](#bonds.BondDetails)
    - [BuyOrder](#bonds.BuyOrder)
    - [FunctionParam](#bonds.FunctionParam)
    - [Params](#bonds.Params)
    - [SellOrder](#bonds.SellOrder)
    - [SwapOrder](#bonds.SwapOrder)
  
- [bonds/genesis.proto](#bonds/genesis.proto)
    - [GenesisState](#bonds.GenesisState)
  
- [bonds/query.proto](#bonds/query.proto)
    - [QueryAlphaMaximumsRequest](#bonds.QueryAlphaMaximumsRequest)
    - [QueryAlphaMaximumsResponse](#bonds.QueryAlphaMaximumsResponse)
    - [QueryAvailableReserveRequest](#bonds.QueryAvailableReserveRequest)
    - [QueryAvailableReserveResponse](#bonds.QueryAvailableReserveResponse)
    - [QueryBatchRequest](#bonds.QueryBatchRequest)
    - [QueryBatchResponse](#bonds.QueryBatchResponse)
    - [QueryBondRequest](#bonds.QueryBondRequest)
    - [QueryBondResponse](#bonds.QueryBondResponse)
    - [QueryBondsDetailedRequest](#bonds.QueryBondsDetailedRequest)
    - [QueryBondsDetailedResponse](#bonds.QueryBondsDetailedResponse)
    - [QueryBondsRequest](#bonds.QueryBondsRequest)
    - [QueryBondsResponse](#bonds.QueryBondsResponse)
    - [QueryBuyPriceRequest](#bonds.QueryBuyPriceRequest)
    - [QueryBuyPriceResponse](#bonds.QueryBuyPriceResponse)
    - [QueryCurrentPriceRequest](#bonds.QueryCurrentPriceRequest)
    - [QueryCurrentPriceResponse](#bonds.QueryCurrentPriceResponse)
    - [QueryCurrentReserveRequest](#bonds.QueryCurrentReserveRequest)
    - [QueryCurrentReserveResponse](#bonds.QueryCurrentReserveResponse)
    - [QueryCustomPriceRequest](#bonds.QueryCustomPriceRequest)
    - [QueryCustomPriceResponse](#bonds.QueryCustomPriceResponse)
    - [QueryLastBatchRequest](#bonds.QueryLastBatchRequest)
    - [QueryLastBatchResponse](#bonds.QueryLastBatchResponse)
    - [QueryParamsRequest](#bonds.QueryParamsRequest)
    - [QueryParamsResponse](#bonds.QueryParamsResponse)
    - [QuerySellReturnRequest](#bonds.QuerySellReturnRequest)
    - [QuerySellReturnResponse](#bonds.QuerySellReturnResponse)
    - [QuerySwapReturnRequest](#bonds.QuerySwapReturnRequest)
    - [QuerySwapReturnResponse](#bonds.QuerySwapReturnResponse)
  
    - [Query](#bonds.Query)
  
- [bonds/tx.proto](#bonds/tx.proto)
    - [MsgBuy](#bonds.MsgBuy)
    - [MsgBuyResponse](#bonds.MsgBuyResponse)
    - [MsgCreateBond](#bonds.MsgCreateBond)
    - [MsgCreateBondResponse](#bonds.MsgCreateBondResponse)
    - [MsgEditBond](#bonds.MsgEditBond)
    - [MsgEditBondResponse](#bonds.MsgEditBondResponse)
    - [MsgMakeOutcomePayment](#bonds.MsgMakeOutcomePayment)
    - [MsgMakeOutcomePaymentResponse](#bonds.MsgMakeOutcomePaymentResponse)
    - [MsgSell](#bonds.MsgSell)
    - [MsgSellResponse](#bonds.MsgSellResponse)
    - [MsgSetNextAlpha](#bonds.MsgSetNextAlpha)
    - [MsgSetNextAlphaResponse](#bonds.MsgSetNextAlphaResponse)
    - [MsgSwap](#bonds.MsgSwap)
    - [MsgSwapResponse](#bonds.MsgSwapResponse)
    - [MsgUpdateBondState](#bonds.MsgUpdateBondState)
    - [MsgUpdateBondStateResponse](#bonds.MsgUpdateBondStateResponse)
    - [MsgWithdrawReserve](#bonds.MsgWithdrawReserve)
    - [MsgWithdrawReserveResponse](#bonds.MsgWithdrawReserveResponse)
    - [MsgWithdrawShare](#bonds.MsgWithdrawShare)
    - [MsgWithdrawShareResponse](#bonds.MsgWithdrawShareResponse)
  
    - [Msg](#bonds.Msg)
  
- [did/did.proto](#did/did.proto)
    - [Claim](#did.Claim)
    - [DidCredential](#did.DidCredential)
    - [IxoDid](#did.IxoDid)
    - [Secret](#did.Secret)
  
- [did/diddoc.proto](#did/diddoc.proto)
    - [BaseDidDoc](#did.BaseDidDoc)
  
- [did/genesis.proto](#did/genesis.proto)
    - [GenesisState](#did.GenesisState)
  
- [did/query.proto](#did/query.proto)
    - [QueryAddressFromBase58EncodedPubkeyRequest](#did.QueryAddressFromBase58EncodedPubkeyRequest)
    - [QueryAddressFromBase58EncodedPubkeyResponse](#did.QueryAddressFromBase58EncodedPubkeyResponse)
    - [QueryAddressFromDidRequest](#did.QueryAddressFromDidRequest)
    - [QueryAddressFromDidResponse](#did.QueryAddressFromDidResponse)
    - [QueryAllDidDocsRequest](#did.QueryAllDidDocsRequest)
    - [QueryAllDidDocsResponse](#did.QueryAllDidDocsResponse)
    - [QueryAllDidsRequest](#did.QueryAllDidsRequest)
    - [QueryAllDidsResponse](#did.QueryAllDidsResponse)
    - [QueryDidDocRequest](#did.QueryDidDocRequest)
    - [QueryDidDocResponse](#did.QueryDidDocResponse)
  
    - [Query](#did.Query)
  
- [did/tx.proto](#did/tx.proto)
    - [MsgAddCredential](#did.MsgAddCredential)
    - [MsgAddCredentialResponse](#did.MsgAddCredentialResponse)
    - [MsgAddDid](#did.MsgAddDid)
    - [MsgAddDidResponse](#did.MsgAddDidResponse)
  
    - [Msg](#did.Msg)
  
- [payments/payments.proto](#payments/payments.proto)
    - [BlockPeriod](#payments.BlockPeriod)
    - [Discount](#payments.Discount)
    - [DistributionShare](#payments.DistributionShare)
    - [PaymentContract](#payments.PaymentContract)
    - [PaymentTemplate](#payments.PaymentTemplate)
    - [Subscription](#payments.Subscription)
    - [TestPeriod](#payments.TestPeriod)
    - [TimePeriod](#payments.TimePeriod)
  
- [payments/genesis.proto](#payments/genesis.proto)
    - [GenesisState](#payments.GenesisState)
  
- [payments/query.proto](#payments/query.proto)
    - [QueryPaymentContractRequest](#payments.QueryPaymentContractRequest)
    - [QueryPaymentContractResponse](#payments.QueryPaymentContractResponse)
    - [QueryPaymentContractsByIdPrefixRequest](#payments.QueryPaymentContractsByIdPrefixRequest)
    - [QueryPaymentContractsByIdPrefixResponse](#payments.QueryPaymentContractsByIdPrefixResponse)
    - [QueryPaymentTemplateRequest](#payments.QueryPaymentTemplateRequest)
    - [QueryPaymentTemplateResponse](#payments.QueryPaymentTemplateResponse)
    - [QuerySubscriptionRequest](#payments.QuerySubscriptionRequest)
    - [QuerySubscriptionResponse](#payments.QuerySubscriptionResponse)
  
    - [Query](#payments.Query)
  
- [payments/tx.proto](#payments/tx.proto)
    - [MsgCreatePaymentContract](#payments.MsgCreatePaymentContract)
    - [MsgCreatePaymentContractResponse](#payments.MsgCreatePaymentContractResponse)
    - [MsgCreatePaymentTemplate](#payments.MsgCreatePaymentTemplate)
    - [MsgCreatePaymentTemplateResponse](#payments.MsgCreatePaymentTemplateResponse)
    - [MsgCreateSubscription](#payments.MsgCreateSubscription)
    - [MsgCreateSubscriptionResponse](#payments.MsgCreateSubscriptionResponse)
    - [MsgEffectPayment](#payments.MsgEffectPayment)
    - [MsgEffectPaymentResponse](#payments.MsgEffectPaymentResponse)
    - [MsgGrantDiscount](#payments.MsgGrantDiscount)
    - [MsgGrantDiscountResponse](#payments.MsgGrantDiscountResponse)
    - [MsgRevokeDiscount](#payments.MsgRevokeDiscount)
    - [MsgRevokeDiscountResponse](#payments.MsgRevokeDiscountResponse)
    - [MsgSetPaymentContractAuthorisation](#payments.MsgSetPaymentContractAuthorisation)
    - [MsgSetPaymentContractAuthorisationResponse](#payments.MsgSetPaymentContractAuthorisationResponse)
  
    - [Msg](#payments.Msg)
  
- [project/project.proto](#project/project.proto)
    - [AccountMap](#project.AccountMap)
    - [AccountMap.MapEntry](#project.AccountMap.MapEntry)
    - [Claim](#project.Claim)
    - [Claims](#project.Claims)
    - [CreateAgentDoc](#project.CreateAgentDoc)
    - [CreateClaimDoc](#project.CreateClaimDoc)
    - [CreateEvaluationDoc](#project.CreateEvaluationDoc)
    - [GenesisAccountMap](#project.GenesisAccountMap)
    - [GenesisAccountMap.MapEntry](#project.GenesisAccountMap.MapEntry)
    - [Params](#project.Params)
    - [ProjectDoc](#project.ProjectDoc)
    - [UpdateAgentDoc](#project.UpdateAgentDoc)
    - [UpdateProjectStatusDoc](#project.UpdateProjectStatusDoc)
    - [WithdrawFundsDoc](#project.WithdrawFundsDoc)
    - [WithdrawalInfoDoc](#project.WithdrawalInfoDoc)
    - [WithdrawalInfoDocs](#project.WithdrawalInfoDocs)
  
- [project/genesis.proto](#project/genesis.proto)
    - [GenesisState](#project.GenesisState)
  
- [project/query.proto](#project/query.proto)
    - [QueryParamsRequest](#project.QueryParamsRequest)
    - [QueryParamsResponse](#project.QueryParamsResponse)
    - [QueryProjectAccountsRequest](#project.QueryProjectAccountsRequest)
    - [QueryProjectAccountsResponse](#project.QueryProjectAccountsResponse)
    - [QueryProjectDocRequest](#project.QueryProjectDocRequest)
    - [QueryProjectDocResponse](#project.QueryProjectDocResponse)
    - [QueryProjectTxRequest](#project.QueryProjectTxRequest)
    - [QueryProjectTxResponse](#project.QueryProjectTxResponse)
  
    - [Query](#project.Query)
  
- [project/tx.proto](#project/tx.proto)
    - [MsgCreateAgent](#project.MsgCreateAgent)
    - [MsgCreateAgentResponse](#project.MsgCreateAgentResponse)
    - [MsgCreateClaim](#project.MsgCreateClaim)
    - [MsgCreateClaimResponse](#project.MsgCreateClaimResponse)
    - [MsgCreateEvaluation](#project.MsgCreateEvaluation)
    - [MsgCreateEvaluationResponse](#project.MsgCreateEvaluationResponse)
    - [MsgCreateProject](#project.MsgCreateProject)
    - [MsgCreateProjectResponse](#project.MsgCreateProjectResponse)
    - [MsgUpdateAgent](#project.MsgUpdateAgent)
    - [MsgUpdateAgentResponse](#project.MsgUpdateAgentResponse)
    - [MsgUpdateProjectDoc](#project.MsgUpdateProjectDoc)
    - [MsgUpdateProjectDocResponse](#project.MsgUpdateProjectDocResponse)
    - [MsgUpdateProjectStatus](#project.MsgUpdateProjectStatus)
    - [MsgUpdateProjectStatusResponse](#project.MsgUpdateProjectStatusResponse)
    - [MsgWithdrawFunds](#project.MsgWithdrawFunds)
    - [MsgWithdrawFundsResponse](#project.MsgWithdrawFundsResponse)
  
    - [Msg](#project.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="bonds/bonds.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/bonds.proto



<a name="bonds.BaseOrder"></a>

### BaseOrder
BaseOrder defines a base order type. It contains all the necessary fields for specifying
a buy, sell, or swap order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| cancelled | [bool](#bool) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="bonds.Batch"></a>

### Batch
Batch holds a collection of outstanding buy, sell, and swap orders on a particular bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| blocks_remaining | [string](#string) |  |  |
| next_public_alpha | [string](#string) |  |  |
| total_buy_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| total_sell_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| buy_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| sell_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| buys | [BuyOrder](#bonds.BuyOrder) | repeated |  |
| sells | [SellOrder](#bonds.SellOrder) | repeated |  |
| swaps | [SwapOrder](#bonds.SwapOrder) | repeated |  |






<a name="bonds.Bond"></a>

### Bond
Bond defines a token bonding curve type with all of its parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#bonds.FunctionParam) | repeated |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| current_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_outcome_payment_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |
| state | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.BondDetails"></a>

### BondDetails
BondDetails contains details about the current state of a given bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| spot_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.BuyOrder"></a>

### BuyOrder
BuyOrder defines a type for submitting a buy order on a bond, together with the maximum
amount of reserve tokens the buyer is willing to pay.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.FunctionParam"></a>

### FunctionParam
FunctionParam is a key-value pair used for specifying a specific bond parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| param | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="bonds.Params"></a>

### Params
Params defines the parameters for the bonds module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reserved_bond_tokens | [string](#string) | repeated |  |






<a name="bonds.SellOrder"></a>

### SellOrder
SellOrder defines a type for submitting a sell order on a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |






<a name="bonds.SwapOrder"></a>

### SwapOrder
SwapOrder defines a type for submitting a swap order between two tokens on a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |
| to_token | [string](#string) |  |  |





 

 

 

 



<a name="bonds/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/genesis.proto



<a name="bonds.GenesisState"></a>

### GenesisState
GenesisState defines the bonds module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [Bond](#bonds.Bond) | repeated |  |
| batches | [Batch](#bonds.Batch) | repeated |  |
| params | [Params](#bonds.Params) |  |  |





 

 

 

 



<a name="bonds/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/query.proto



<a name="bonds.QueryAlphaMaximumsRequest"></a>

### QueryAlphaMaximumsRequest
QueryAlphaMaximumsRequest is the request type for the Query/AlphaMaximums RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryAlphaMaximumsResponse"></a>

### QueryAlphaMaximumsResponse
QueryAlphaMaximumsResponse is the response type for the Query/AlphaMaximums RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_system_alpha_increase | [string](#string) |  |  |
| max_system_alpha | [string](#string) |  |  |






<a name="bonds.QueryAvailableReserveRequest"></a>

### QueryAvailableReserveRequest
QueryAvailableReserveRequest is the request type for the Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryAvailableReserveResponse"></a>

### QueryAvailableReserveResponse
QueryAvailableReserveResponse is the response type for the Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryBatchRequest"></a>

### QueryBatchRequest
QueryBatchRequest is the request type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryBatchResponse"></a>

### QueryBatchResponse
QueryBatchResponse is the response type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#bonds.Batch) |  |  |






<a name="bonds.QueryBondRequest"></a>

### QueryBondRequest
QueryBondRequest is the request type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryBondResponse"></a>

### QueryBondResponse
QueryBondResponse is the response type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#bonds.Bond) |  |  |






<a name="bonds.QueryBondsDetailedRequest"></a>

### QueryBondsDetailedRequest
QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC method.






<a name="bonds.QueryBondsDetailedResponse"></a>

### QueryBondsDetailedResponse
QueryBondsDetailedResponse is the response type for the Query/BondsDetailed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds_detailed | [BondDetails](#bonds.BondDetails) | repeated |  |






<a name="bonds.QueryBondsRequest"></a>

### QueryBondsRequest
QueryBondsRequest is the request type for the Query/Bonds RPC method.






<a name="bonds.QueryBondsResponse"></a>

### QueryBondsResponse
QueryBondsResponse is the response type for the Query/Bonds RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [string](#string) | repeated |  |






<a name="bonds.QueryBuyPriceRequest"></a>

### QueryBuyPriceRequest
QueryCustomPriceRequest is the request type for the Query/BuyPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QueryBuyPriceResponse"></a>

### QueryBuyPriceResponse
QueryCustomPriceResponse is the response type for the Query/BuyPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest
QueryCurrentPriceRequest is the request type for the Query/CurrentPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse
QueryCurrentPriceResponse is the response type for the Query/CurrentPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest
QueryCurrentReserveRequest is the request type for the Query/CurrentReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse
QueryCurrentReserveResponse is the response type for the Query/CurrentReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryCustomPriceRequest"></a>

### QueryCustomPriceRequest
QueryCustomPriceRequest is the request type for the Query/CustomPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QueryCustomPriceResponse"></a>

### QueryCustomPriceResponse
QueryCustomPriceResponse is the response type for the Query/CustomPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryLastBatchRequest"></a>

### QueryLastBatchRequest
QueryLastBatchRequest is the request type for the Query/LastBatch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryLastBatchResponse"></a>

### QueryLastBatchResponse
QueryLastBatchResponse is the response type for the Query/LastBatch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_batch | [Batch](#bonds.Batch) |  |  |






<a name="bonds.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="bonds.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#bonds.Params) |  |  |






<a name="bonds.QuerySellReturnRequest"></a>

### QuerySellReturnRequest
QuerySellReturnRequest is the request type for the Query/SellReturn RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QuerySellReturnResponse"></a>

### QuerySellReturnResponse
QuerySellReturnResponse is the response type for the Query/SellReturn RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| exit_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QuerySwapReturnRequest"></a>

### QuerySwapReturnRequest
QuerySwapReturnRequest is the request type for the Query/SwapReturn RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| from_token_with_amount | [string](#string) |  |  |
| to_token | [string](#string) |  |  |






<a name="bonds.QuerySwapReturnResponse"></a>

### QuerySwapReturnResponse
QuerySwapReturnResponse is the response type for the Query/SwapReturn RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 


<a name="bonds.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bonds | [QueryBondsRequest](#bonds.QueryBondsRequest) | [QueryBondsResponse](#bonds.QueryBondsResponse) | Bonds returns all existing bonds. |
| BondsDetailed | [QueryBondsDetailedRequest](#bonds.QueryBondsDetailedRequest) | [QueryBondsDetailedResponse](#bonds.QueryBondsDetailedResponse) | BondsDetailed returns a list of all existing bonds with details about their current state. |
| Params | [QueryParamsRequest](#bonds.QueryParamsRequest) | [QueryParamsResponse](#bonds.QueryParamsResponse) | Params queries the paramaters of x/bonds module. |
| Bond | [QueryBondRequest](#bonds.QueryBondRequest) | [QueryBondResponse](#bonds.QueryBondResponse) | Bond queries info of a specific bond. |
| Batch | [QueryBatchRequest](#bonds.QueryBatchRequest) | [QueryBatchResponse](#bonds.QueryBatchResponse) | Batch queries info of a specific bond&#39;s current batch. |
| LastBatch | [QueryLastBatchRequest](#bonds.QueryLastBatchRequest) | [QueryLastBatchResponse](#bonds.QueryLastBatchResponse) | LastBatch queries info of a specific bond&#39;s last batch. |
| CurrentPrice | [QueryCurrentPriceRequest](#bonds.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#bonds.QueryCurrentPriceResponse) | CurrentPrice queries the current price/s of a specific bond. |
| CurrentReserve | [QueryCurrentReserveRequest](#bonds.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#bonds.QueryCurrentReserveResponse) | CurrentReserve queries the current balance/s of the reserve pool for a specific bond. |
| AvailableReserve | [QueryAvailableReserveRequest](#bonds.QueryAvailableReserveRequest) | [QueryAvailableReserveResponse](#bonds.QueryAvailableReserveResponse) | AvailableReserve queries current available balance/s of the reserve pool for a specific bond. |
| CustomPrice | [QueryCustomPriceRequest](#bonds.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#bonds.QueryCustomPriceResponse) | CustomPrice queries price/s of a specific bond at a specific supply. |
| BuyPrice | [QueryBuyPriceRequest](#bonds.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#bonds.QueryBuyPriceResponse) | BuyPrice queries price/s of buying an amount of tokens from a specific bond. |
| SellReturn | [QuerySellReturnRequest](#bonds.QuerySellReturnRequest) | [QuerySellReturnResponse](#bonds.QuerySellReturnResponse) | SellReturn queries return/s on selling an amount of tokens of a specific bond. |
| SwapReturn | [QuerySwapReturnRequest](#bonds.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#bonds.QuerySwapReturnResponse) | SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond. |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#bonds.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#bonds.QueryAlphaMaximumsResponse) | AlphaMaximums queries alpha maximums for a specific augmented bonding curve. |

 



<a name="bonds/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/tx.proto



<a name="bonds.MsgBuy"></a>

### MsgBuy
MsgBuy defines a message for buying from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse defines the Msg/Buy response type.






<a name="bonds.MsgCreateBond"></a>

### MsgCreateBond
MsgCreateBond defines a message for creating a new bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#bonds.FunctionParam) | repeated |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |






<a name="bonds.MsgCreateBondResponse"></a>

### MsgCreateBondResponse
MsgCreateBondResponse defines the Msg/CreateBond response type.






<a name="bonds.MsgEditBond"></a>

### MsgEditBond
MsgEditBond defines a message for editing an existing bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| order_quantity_limits | [string](#string) |  |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |






<a name="bonds.MsgEditBondResponse"></a>

### MsgEditBondResponse
MsgEditBondResponse defines the Msg/EditBond response type.






<a name="bonds.MsgMakeOutcomePayment"></a>

### MsgMakeOutcomePayment
MsgMakeOutcomePayment defines a message for making an outcome payment to a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgMakeOutcomePaymentResponse"></a>

### MsgMakeOutcomePaymentResponse
MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response type.






<a name="bonds.MsgSell"></a>

### MsgSell
MsgSell defines a message for selling from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seller_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse defines the Msg/Sell response type.






<a name="bonds.MsgSetNextAlpha"></a>

### MsgSetNextAlpha
MsgSetNextAlpha defines a message for editing a bond&#39;s alpha parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| alpha | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |






<a name="bonds.MsgSetNextAlphaResponse"></a>

### MsgSetNextAlphaResponse







<a name="bonds.MsgSwap"></a>

### MsgSwap
MsgSwap defines a message for swapping from one reserve bond token to another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swapper_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| from | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| to_token | [string](#string) |  |  |






<a name="bonds.MsgSwapResponse"></a>

### MsgSwapResponse
MsgSwapResponse defines the Msg/Swap response type.






<a name="bonds.MsgUpdateBondState"></a>

### MsgUpdateBondState
MsgUpdateBondState defines a message for updating a bond&#39;s current state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| state | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |






<a name="bonds.MsgUpdateBondStateResponse"></a>

### MsgUpdateBondStateResponse
MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.






<a name="bonds.MsgWithdrawReserve"></a>

### MsgWithdrawReserve
MsgWithdrawReserve defines a message for withdrawing reserve from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdrawer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgWithdrawReserveResponse"></a>

### MsgWithdrawReserveResponse
MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.






<a name="bonds.MsgWithdrawShare"></a>

### MsgWithdrawShare
MsgWithdrawShare defines a message for withdrawing a share from a bond that is in the SETTLE stage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgWithdrawShareResponse"></a>

### MsgWithdrawShareResponse
MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.





 

 

 


<a name="bonds.Msg"></a>

### Msg
Msg defines the bonds Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBond | [MsgCreateBond](#bonds.MsgCreateBond) | [MsgCreateBondResponse](#bonds.MsgCreateBondResponse) | CreateBond defines a method for creating a bond. |
| EditBond | [MsgEditBond](#bonds.MsgEditBond) | [MsgEditBondResponse](#bonds.MsgEditBondResponse) | EditBond defines a method for editing a bond. |
| SetNextAlpha | [MsgSetNextAlpha](#bonds.MsgSetNextAlpha) | [MsgSetNextAlphaResponse](#bonds.MsgSetNextAlphaResponse) | SetNextAlpha defines a method for editing a bond&#39;s alpha parameter. |
| UpdateBondState | [MsgUpdateBondState](#bonds.MsgUpdateBondState) | [MsgUpdateBondStateResponse](#bonds.MsgUpdateBondStateResponse) | UpdateBondState defines a method for updating a bond&#39;s current state. |
| Buy | [MsgBuy](#bonds.MsgBuy) | [MsgBuyResponse](#bonds.MsgBuyResponse) | Buy defines a method for buying from a bond. |
| Sell | [MsgSell](#bonds.MsgSell) | [MsgSellResponse](#bonds.MsgSellResponse) | Sell defines a method for selling from a bond. |
| Swap | [MsgSwap](#bonds.MsgSwap) | [MsgSwapResponse](#bonds.MsgSwapResponse) | Swap defines a method for swapping from one reserve bond token to another. |
| MakeOutcomePayment | [MsgMakeOutcomePayment](#bonds.MsgMakeOutcomePayment) | [MsgMakeOutcomePaymentResponse](#bonds.MsgMakeOutcomePaymentResponse) | MakeOutcomePayment defines a method for making an outcome payment to a bond. |
| WithdrawShare | [MsgWithdrawShare](#bonds.MsgWithdrawShare) | [MsgWithdrawShareResponse](#bonds.MsgWithdrawShareResponse) | WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage. |
| WithdrawReserve | [MsgWithdrawReserve](#bonds.MsgWithdrawReserve) | [MsgWithdrawReserveResponse](#bonds.MsgWithdrawReserveResponse) | WithdrawReserve defines a method for withdrawing reserve from a bond. |

 



<a name="did/did.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/did.proto



<a name="did.Claim"></a>

### Claim
TODO


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| KYC_validated | [bool](#bool) |  |  |






<a name="did.DidCredential"></a>

### DidCredential
TODO


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cred_type | [string](#string) | repeated |  |
| issuer | [string](#string) |  |  |
| issued | [string](#string) |  |  |
| claim | [Claim](#did.Claim) |  |  |






<a name="did.IxoDid"></a>

### IxoDid
TODO


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| verify_key | [string](#string) |  |  |
| encryption_public_key | [string](#string) |  |  |
| secret | [Secret](#did.Secret) |  |  |






<a name="did.Secret"></a>

### Secret
TODO


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seed | [string](#string) |  |  |
| sign_key | [string](#string) |  |  |
| encryption_private_key | [string](#string) |  |  |





 

 

 

 



<a name="did/diddoc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/diddoc.proto



<a name="did.BaseDidDoc"></a>

### BaseDidDoc
BaseDidDoc defines a base DID document type. It implements the DidDoc interface.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| credentials | [DidCredential](#did.DidCredential) | repeated |  |





 

 

 

 



<a name="did/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/genesis.proto



<a name="did.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did_docs | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |





 

 

 

 



<a name="did/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/query.proto



<a name="did.QueryAddressFromBase58EncodedPubkeyRequest"></a>

### QueryAddressFromBase58EncodedPubkeyRequest
QueryAddressFromBase58EncodedPubkeyRequest is the request type for the Query/AddressFromBase58EncodedPubkey RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubKey | [string](#string) |  |  |






<a name="did.QueryAddressFromBase58EncodedPubkeyResponse"></a>

### QueryAddressFromBase58EncodedPubkeyResponse
QueryAddressFromBase58EncodedPubkeyResponse is the response type for the Query/AddressFromBase58EncodedPubkey RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="did.QueryAddressFromDidRequest"></a>

### QueryAddressFromDidRequest
QueryAddressFromDidRequest is the request type for the Query/AddressFromDid RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |






<a name="did.QueryAddressFromDidResponse"></a>

### QueryAddressFromDidResponse
QueryAddressFromDidResponse is the response type for the Query/AddressFromDid RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="did.QueryAllDidDocsRequest"></a>

### QueryAllDidDocsRequest
QueryAllDidDocsRequest is the request type for the Query/AllDidDocs RPC method.






<a name="did.QueryAllDidDocsResponse"></a>

### QueryAllDidDocsResponse
QueryAllDidDocsResponse is the response type for the Query/AllDidDocs RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="did.QueryAllDidsRequest"></a>

### QueryAllDidsRequest
QueryAllDidsRequest is the request type for the Query/AllDids RPC method.






<a name="did.QueryAllDidsResponse"></a>

### QueryAllDidsResponse
QueryAllDidsResponse is the response type for the Query/AllDids RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dids | [string](#string) | repeated |  |






<a name="did.QueryDidDocRequest"></a>

### QueryDidDocRequest
QueryDidDocRequest is the request type for the Query/DidDoc RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |






<a name="did.QueryDidDocResponse"></a>

### QueryDidDocResponse
QueryDidDocResponse is the response type for the Query/DidDoc RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddoc | [google.protobuf.Any](#google.protobuf.Any) |  |  |





 

 

 


<a name="did.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DidDoc | [QueryDidDocRequest](#did.QueryDidDocRequest) | [QueryDidDocResponse](#did.QueryDidDocResponse) | DidDoc queries info of a specific DID&#39;s DidDoc. |
| AllDids | [QueryAllDidsRequest](#did.QueryAllDidsRequest) | [QueryAllDidsResponse](#did.QueryAllDidsResponse) | AllDids returns a list of all existing DIDs. |
| AllDidDocs | [QueryAllDidDocsRequest](#did.QueryAllDidDocsRequest) | [QueryAllDidDocsResponse](#did.QueryAllDidDocsResponse) | AllDidDocs returns a list of all existing DidDocs (i.e. all DIDs along with their DidDoc info). |
| AddressFromDid | [QueryAddressFromDidRequest](#did.QueryAddressFromDidRequest) | [QueryAddressFromDidResponse](#did.QueryAddressFromDidResponse) | AddressFromDid retrieves the cosmos address associated to an ixo DID. |
| AddressFromBase58EncodedPubkey | [QueryAddressFromBase58EncodedPubkeyRequest](#did.QueryAddressFromBase58EncodedPubkeyRequest) | [QueryAddressFromBase58EncodedPubkeyResponse](#did.QueryAddressFromBase58EncodedPubkeyResponse) | AddressFromBase58EncodedPubkey retrieves the cosmos address associated to an ixo DID&#39;s pubkey. |

 



<a name="did/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/tx.proto



<a name="did.MsgAddCredential"></a>

### MsgAddCredential
MsgAddCredential defines a message for adding a credential to the signer&#39;s DID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did_credential | [DidCredential](#did.DidCredential) |  |  |






<a name="did.MsgAddCredentialResponse"></a>

### MsgAddCredentialResponse
MsgAddCredentialResponse defines the Msg/AddCredential response type.






<a name="did.MsgAddDid"></a>

### MsgAddDid
MsgAddDid defines a message for adding a DID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pubKey | [string](#string) |  |  |






<a name="did.MsgAddDidResponse"></a>

### MsgAddDidResponse
MsgAddDidResponse defines the Msg/AddDid response type.





 

 

 


<a name="did.Msg"></a>

### Msg
Msg defines the did Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddDid | [MsgAddDid](#did.MsgAddDid) | [MsgAddDidResponse](#did.MsgAddDidResponse) | AddDid defines a method for adding a DID. |
| AddCredential | [MsgAddCredential](#did.MsgAddCredential) | [MsgAddCredentialResponse](#did.MsgAddCredentialResponse) | AddCredential defines a method for adding a credential to the signer&#39;s DID. |

 



<a name="payments/payments.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/payments.proto



<a name="payments.BlockPeriod"></a>

### BlockPeriod
BlockPeriod implements the Period interface and specifies a period in terms of number
of blocks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="payments.Discount"></a>

### Discount
Discount contains details about a discount which can be granted to payers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| percent | [string](#string) |  |  |






<a name="payments.DistributionShare"></a>

### DistributionShare
DistributionShare specifies the share of a specific payment an address will receive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| percentage | [string](#string) |  |  |






<a name="payments.PaymentContract"></a>

### PaymentContract
PaymentContract specifies an agreement between a payer and payee/s which can be invoked
once or multiple times to effect payment/s.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_template_id | [string](#string) |  |  |
| creator | [string](#string) |  |  |
| payer | [string](#string) |  |  |
| recipients | [DistributionShare](#payments.DistributionShare) | repeated |  |
| cumulative_pay | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_remainder | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| can_deauthorise | [bool](#bool) |  |  |
| authorised | [bool](#bool) |  |  |
| discount_id | [string](#string) |  |  |






<a name="payments.PaymentTemplate"></a>

### PaymentTemplate
PaymentTemplate contains details about a payment, with no info about the payer or payee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_minimum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_maximum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| discounts | [Discount](#payments.Discount) | repeated |  |






<a name="payments.Subscription"></a>

### Subscription
Subscription specifies details of a payment to be effected periodically.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| periods_so_far | [string](#string) |  |  |
| max_periods | [string](#string) |  |  |
| periods_accumulated | [string](#string) |  |  |
| period | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="payments.TestPeriod"></a>

### TestPeriod
TestPeriod implements the Period interface and is identical to BlockPeriod, except it
ignores the context in periodEnded() and periodStarted().


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="payments.TimePeriod"></a>

### TimePeriod
TimePeriod implements the Period interface and specifies a period in terms of time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| period_start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="payments/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/genesis.proto



<a name="payments.GenesisState"></a>

### GenesisState
GenesisState defines the payments module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_templates | [PaymentTemplate](#payments.PaymentTemplate) | repeated |  |
| payment_contracts | [PaymentContract](#payments.PaymentContract) | repeated |  |
| subscriptions | [Subscription](#payments.Subscription) | repeated |  |





 

 

 

 



<a name="payments/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/query.proto



<a name="payments.QueryPaymentContractRequest"></a>

### QueryPaymentContractRequest
QueryPaymentContractRequest is the request type for the Query/PaymentContract RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |






<a name="payments.QueryPaymentContractResponse"></a>

### QueryPaymentContractResponse
QueryPaymentContractResponse is the response type for the Query/PaymentContract RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract | [PaymentContract](#payments.PaymentContract) |  |  |






<a name="payments.QueryPaymentContractsByIdPrefixRequest"></a>

### QueryPaymentContractsByIdPrefixRequest
QueryPaymentContractsByIdPrefixRequest is the request type for the Query/PaymentContractsByIdPrefix RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts_id_prefix | [string](#string) |  |  |






<a name="payments.QueryPaymentContractsByIdPrefixResponse"></a>

### QueryPaymentContractsByIdPrefixResponse
QueryPaymentContractsByIdPrefixResponse is the response type for the Query/PaymentContractsByIdPrefix RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts | [PaymentContract](#payments.PaymentContract) | repeated |  |






<a name="payments.QueryPaymentTemplateRequest"></a>

### QueryPaymentTemplateRequest
QueryPaymentTemplateRequest is the request type for the Query/PaymentTemplate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template_id | [string](#string) |  |  |






<a name="payments.QueryPaymentTemplateResponse"></a>

### QueryPaymentTemplateResponse
QueryPaymentTemplateResponse is the response type for the Query/PaymentTemplate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template | [PaymentTemplate](#payments.PaymentTemplate) |  |  |






<a name="payments.QuerySubscriptionRequest"></a>

### QuerySubscriptionRequest
QuerySubscriptionRequest is the request type for the Query/Subscription RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |






<a name="payments.QuerySubscriptionResponse"></a>

### QuerySubscriptionResponse
QuerySubscriptionResponse is the response type for the Query/Subscription RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#payments.Subscription) |  |  |





 

 

 


<a name="payments.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PaymentTemplate | [QueryPaymentTemplateRequest](#payments.QueryPaymentTemplateRequest) | [QueryPaymentTemplateResponse](#payments.QueryPaymentTemplateResponse) | PaymentTemplate queries info of a specific payment template. |
| PaymentContract | [QueryPaymentContractRequest](#payments.QueryPaymentContractRequest) | [QueryPaymentContractResponse](#payments.QueryPaymentContractResponse) | PaymentContract queries info of a specific payment contract. |
| PaymentContractsByIdPrefix | [QueryPaymentContractsByIdPrefixRequest](#payments.QueryPaymentContractsByIdPrefixRequest) | [QueryPaymentContractsByIdPrefixResponse](#payments.QueryPaymentContractsByIdPrefixResponse) | PaymentContractsByIdPrefix lists all payment contracts having an id with a specific prefix. |
| Subscription | [QuerySubscriptionRequest](#payments.QuerySubscriptionRequest) | [QuerySubscriptionResponse](#payments.QuerySubscriptionResponse) | Subscription queries info of a specific Subscription. |

 



<a name="payments/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/tx.proto



<a name="payments.MsgCreatePaymentContract"></a>

### MsgCreatePaymentContract
MsgCreatePaymentContract defines a message for creating a payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| payment_template_id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| payer | [string](#string) |  |  |
| recipients | [DistributionShare](#payments.DistributionShare) | repeated |  |
| can_deauthorise | [bool](#bool) |  |  |
| discount_id | [string](#string) |  |  |






<a name="payments.MsgCreatePaymentContractResponse"></a>

### MsgCreatePaymentContractResponse
MsgCreatePaymentContractResponse defines the Msg/CreatePaymentContract response type.






<a name="payments.MsgCreatePaymentTemplate"></a>

### MsgCreatePaymentTemplate
MsgCreatePaymentTemplate defines a message for creating a payment template.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| payment_template | [PaymentTemplate](#payments.PaymentTemplate) |  |  |






<a name="payments.MsgCreatePaymentTemplateResponse"></a>

### MsgCreatePaymentTemplateResponse
MsgCreatePaymentTemplateResponse defines the Msg/CreatePaymentTemplate response type.






<a name="payments.MsgCreateSubscription"></a>

### MsgCreateSubscription
MsgCreateSubscription defines a message for creating a subscription.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| subscription_id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| max_periods | [string](#string) |  |  |
| period | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="payments.MsgCreateSubscriptionResponse"></a>

### MsgCreateSubscriptionResponse
MsgCreateSubscriptionResponse defines the Msg/CreateSubscription response type.






<a name="payments.MsgEffectPayment"></a>

### MsgEffectPayment
MsgEffectPayment defines a message for putting a specific payment contract into effect.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |






<a name="payments.MsgEffectPaymentResponse"></a>

### MsgEffectPaymentResponse
MsgEffectPaymentResponse defines the Msg/EffectPayment response type.






<a name="payments.MsgGrantDiscount"></a>

### MsgGrantDiscount
MsgGrantDiscount defines a message for granting a discount to a payer on a specific payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| discount_id | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="payments.MsgGrantDiscountResponse"></a>

### MsgGrantDiscountResponse
MsgGrantDiscountResponse defines the Msg/GrantDiscount response type.






<a name="payments.MsgRevokeDiscount"></a>

### MsgRevokeDiscount
MsgRevokeDiscount defines a message for revoking a discount previously granted to a payer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| holder | [string](#string) |  |  |






<a name="payments.MsgRevokeDiscountResponse"></a>

### MsgRevokeDiscountResponse
MsgRevokeDiscountResponse defines the Msg/RevokeDiscount response type.






<a name="payments.MsgSetPaymentContractAuthorisation"></a>

### MsgSetPaymentContractAuthorisation
MsgSetPaymentContractAuthorisation defines a message for authorising or deauthorising a payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |
| payer_did | [string](#string) |  |  |
| authorised | [bool](#bool) |  |  |






<a name="payments.MsgSetPaymentContractAuthorisationResponse"></a>

### MsgSetPaymentContractAuthorisationResponse
MsgSetPaymentContractAuthorisationResponse defines the Msg/SetPaymentContractAuthorisation response type.





 

 

 


<a name="payments.Msg"></a>

### Msg
Msg defines the payments Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetPaymentContractAuthorisation | [MsgSetPaymentContractAuthorisation](#payments.MsgSetPaymentContractAuthorisation) | [MsgSetPaymentContractAuthorisationResponse](#payments.MsgSetPaymentContractAuthorisationResponse) | SetPaymentContractAuthorisation defines a method for authorising or deauthorising a payment contract. |
| CreatePaymentTemplate | [MsgCreatePaymentTemplate](#payments.MsgCreatePaymentTemplate) | [MsgCreatePaymentTemplateResponse](#payments.MsgCreatePaymentTemplateResponse) | CreatePaymentTemplate defines a method for creating a payment template. |
| CreatePaymentContract | [MsgCreatePaymentContract](#payments.MsgCreatePaymentContract) | [MsgCreatePaymentContractResponse](#payments.MsgCreatePaymentContractResponse) | CreatePaymentContract defines a method for creating a payment contract. |
| CreateSubscription | [MsgCreateSubscription](#payments.MsgCreateSubscription) | [MsgCreateSubscriptionResponse](#payments.MsgCreateSubscriptionResponse) | CreateSubscription defines a method for creating a subscription. |
| GrantDiscount | [MsgGrantDiscount](#payments.MsgGrantDiscount) | [MsgGrantDiscountResponse](#payments.MsgGrantDiscountResponse) | GrantDiscount defines a method for granting a discount to a payer on a specific payment contract. |
| RevokeDiscount | [MsgRevokeDiscount](#payments.MsgRevokeDiscount) | [MsgRevokeDiscountResponse](#payments.MsgRevokeDiscountResponse) | RevokeDiscount defines a method for revoking a discount previously granted to a payer. |
| EffectPayment | [MsgEffectPayment](#payments.MsgEffectPayment) | [MsgEffectPaymentResponse](#payments.MsgEffectPaymentResponse) | EffectPayment defines a method for putting a specific payment contract into effect. |

 



<a name="project/project.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/project.proto



<a name="project.AccountMap"></a>

### AccountMap
AccountMap maps a specific project&#39;s account names to the accounts&#39; addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| map | [AccountMap.MapEntry](#project.AccountMap.MapEntry) | repeated |  |






<a name="project.AccountMap.MapEntry"></a>

### AccountMap.MapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="project.Claim"></a>

### Claim
Claim contains details required to start a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| template_id | [string](#string) |  |  |
| claimer_did | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="project.Claims"></a>

### Claims
Claims contains a list of type Claim.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims_list | [Claim](#project.Claim) | repeated |  |






<a name="project.CreateAgentDoc"></a>

### CreateAgentDoc
CreateAgentDoc contains details required to create an agent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_did | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="project.CreateClaimDoc"></a>

### CreateClaimDoc
CreateClaimDoc contains details required to create a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| claim_template_id | [string](#string) |  |  |






<a name="project.CreateEvaluationDoc"></a>

### CreateEvaluationDoc
CreateEvaluationDoc contains details required to create an evaluation for a specific claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="project.GenesisAccountMap"></a>

### GenesisAccountMap
GenesisAccountMap is a type used at genesis that maps a specific project&#39;s account names to the accounts&#39; addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| map | [GenesisAccountMap.MapEntry](#project.GenesisAccountMap.MapEntry) | repeated |  |






<a name="project.GenesisAccountMap.MapEntry"></a>

### GenesisAccountMap.MapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="project.Params"></a>

### Params
Params defines the parameters for the project module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo_did | [string](#string) |  |  |
| project_minimum_initial_funding | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| oracle_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |






<a name="project.ProjectDoc"></a>

### ProjectDoc
ProjectDoc defines a project (or entity) type with all of its parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| status | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="project.UpdateAgentDoc"></a>

### UpdateAgentDoc
UpdateAgentDoc contains details required to update an agent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| status | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="project.UpdateProjectStatusDoc"></a>

### UpdateProjectStatusDoc
UpdateProjectStatusDoc contains details required to update a project&#39;s status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| eth_funding_txn_id | [string](#string) |  |  |






<a name="project.WithdrawFundsDoc"></a>

### WithdrawFundsDoc
WithdrawFundsDoc contains details required to withdraw funds from a specific project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| is_refund | [bool](#bool) |  |  |






<a name="project.WithdrawalInfoDoc"></a>

### WithdrawalInfoDoc
WithdrawalInfoDoc contains details required to withdraw from a specific project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="project.WithdrawalInfoDocs"></a>

### WithdrawalInfoDocs
WithdrawalInfoDocs contains a list of type WithdrawalInfoDoc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs_list | [WithdrawalInfoDoc](#project.WithdrawalInfoDoc) | repeated |  |





 

 

 

 



<a name="project/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/genesis.proto



<a name="project.GenesisState"></a>

### GenesisState
GenesisState defines the project module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_docs | [ProjectDoc](#project.ProjectDoc) | repeated |  |
| account_maps | [GenesisAccountMap](#project.GenesisAccountMap) | repeated |  |
| withdrawals_infos | [WithdrawalInfoDocs](#project.WithdrawalInfoDocs) | repeated |  |
| claims | [Claims](#project.Claims) | repeated |  |
| params | [Params](#project.Params) |  |  |





 

 

 

 



<a name="project/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/query.proto



<a name="project.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="project.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#project.Params) |  |  |






<a name="project.QueryProjectAccountsRequest"></a>

### QueryProjectAccountsRequest
QueryProjectAccountsRequest is the request type for the Query/ProjectAccounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectAccountsResponse"></a>

### QueryProjectAccountsResponse
QueryProjectAccountsResponse is the response type for the Query/ProjectAccounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_map | [AccountMap](#project.AccountMap) |  |  |






<a name="project.QueryProjectDocRequest"></a>

### QueryProjectDocRequest
QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectDocResponse"></a>

### QueryProjectDocResponse
QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_doc | [ProjectDoc](#project.ProjectDoc) |  |  |






<a name="project.QueryProjectTxRequest"></a>

### QueryProjectTxRequest
QueryProjectTxRequest is the request type for the Query/ProjectTx RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectTxResponse"></a>

### QueryProjectTxResponse
QueryProjectTxResponse is the response type for the Query/ProjectTx RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txs | [WithdrawalInfoDocs](#project.WithdrawalInfoDocs) |  |  |





 

 

 


<a name="project.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ProjectDoc | [QueryProjectDocRequest](#project.QueryProjectDocRequest) | [QueryProjectDocResponse](#project.QueryProjectDocResponse) | ProjectDoc queries info of a specific project. |
| ProjectAccounts | [QueryProjectAccountsRequest](#project.QueryProjectAccountsRequest) | [QueryProjectAccountsResponse](#project.QueryProjectAccountsResponse) | ProjectAccounts lists a specific project&#39;s accounts. |
| ProjectTx | [QueryProjectTxRequest](#project.QueryProjectTxRequest) | [QueryProjectTxResponse](#project.QueryProjectTxResponse) | ProjectTx lists a specific project&#39;s transactions. |
| Params | [QueryParamsRequest](#project.QueryParamsRequest) | [QueryParamsResponse](#project.QueryParamsResponse) | Params queries the paramaters of x/project module. |

 



<a name="project/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/tx.proto



<a name="project.MsgCreateAgent"></a>

### MsgCreateAgent
MsgCreateAgent defines a message for creating an agent on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateAgentDoc](#project.CreateAgentDoc) |  |  |






<a name="project.MsgCreateAgentResponse"></a>

### MsgCreateAgentResponse
MsgCreateAgentResponse defines the Msg/CreateAgent response type.






<a name="project.MsgCreateClaim"></a>

### MsgCreateClaim
MsgCreateClaim defines a message for creating a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateClaimDoc](#project.CreateClaimDoc) |  |  |






<a name="project.MsgCreateClaimResponse"></a>

### MsgCreateClaimResponse
MsgCreateClaimResponse defines the Msg/CreateClaim response type.






<a name="project.MsgCreateEvaluation"></a>

### MsgCreateEvaluation
MsgCreateEvaluation defines a message for creating an evaluation for a specific claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateEvaluationDoc](#project.CreateEvaluationDoc) |  |  |






<a name="project.MsgCreateEvaluationResponse"></a>

### MsgCreateEvaluationResponse
MsgCreateEvaluationResponse defines the Msg/CreateEvaluation response type.






<a name="project.MsgCreateProject"></a>

### MsgCreateProject
MsgCreateProject defines a message for creating a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="project.MsgCreateProjectResponse"></a>

### MsgCreateProjectResponse
MsgCreateProjectResponse defines the Msg/CreateProject response type.






<a name="project.MsgUpdateAgent"></a>

### MsgUpdateAgent
MsgUpdateAgent defines a message for updating an agent on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateAgentDoc](#project.UpdateAgentDoc) |  |  |






<a name="project.MsgUpdateAgentResponse"></a>

### MsgUpdateAgentResponse
MsgUpdateAgentResponse defines the Msg/UpdateAgent response type.






<a name="project.MsgUpdateProjectDoc"></a>

### MsgUpdateProjectDoc
MsgUpdateProjectDoc defines a message for updating a project&#39;s data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="project.MsgUpdateProjectDocResponse"></a>

### MsgUpdateProjectDocResponse
MsgUpdateProjectDocResponse defines the Msg/UpdateProjectDoc response type.






<a name="project.MsgUpdateProjectStatus"></a>

### MsgUpdateProjectStatus
MsgUpdateProjectStatus defines a message for updating a project&#39;s current status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateProjectStatusDoc](#project.UpdateProjectStatusDoc) |  |  |






<a name="project.MsgUpdateProjectStatusResponse"></a>

### MsgUpdateProjectStatusResponse
MsgUpdateProjectStatusResponse defines the Msg/UpdateProjectStatus response type.






<a name="project.MsgWithdrawFunds"></a>

### MsgWithdrawFunds
MsgWithdrawFunds defines a message for project agents to withdraw their funds from a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| data | [WithdrawFundsDoc](#project.WithdrawFundsDoc) |  |  |






<a name="project.MsgWithdrawFundsResponse"></a>

### MsgWithdrawFundsResponse
MsgWithdrawFundsResponse defines the Msg/WithdrawFunds response type.





 

 

 


<a name="project.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProject | [MsgCreateProject](#project.MsgCreateProject) | [MsgCreateProjectResponse](#project.MsgCreateProjectResponse) | CreateProject defines a method for creating a project. |
| UpdateProjectStatus | [MsgUpdateProjectStatus](#project.MsgUpdateProjectStatus) | [MsgUpdateProjectStatusResponse](#project.MsgUpdateProjectStatusResponse) | UpdateProjectStatus defines a method for updating a project&#39;s current status. |
| CreateAgent | [MsgCreateAgent](#project.MsgCreateAgent) | [MsgCreateAgentResponse](#project.MsgCreateAgentResponse) | CreateAgent defines a method for creating an agent on a project. |
| UpdateAgent | [MsgUpdateAgent](#project.MsgUpdateAgent) | [MsgUpdateAgentResponse](#project.MsgUpdateAgentResponse) | UpdateAgent defines a method for updating an agent on a project. |
| CreateClaim | [MsgCreateClaim](#project.MsgCreateClaim) | [MsgCreateClaimResponse](#project.MsgCreateClaimResponse) | CreateClaim defines a method for creating a claim on a project. |
| CreateEvaluation | [MsgCreateEvaluation](#project.MsgCreateEvaluation) | [MsgCreateEvaluationResponse](#project.MsgCreateEvaluationResponse) | CreateEvaluation defines a method for creating an evaluation for a specific claim on a project. |
| WithdrawFunds | [MsgWithdrawFunds](#project.MsgWithdrawFunds) | [MsgWithdrawFundsResponse](#project.MsgWithdrawFundsResponse) | WithdrawFunds defines a method for project agents to withdraw their funds from a project. |
| UpdateProjectDoc | [MsgUpdateProjectDoc](#project.MsgUpdateProjectDoc) | [MsgUpdateProjectDocResponse](#project.MsgUpdateProjectDocResponse) | UpdateProjectDoc defines a method for updating a project&#39;s data. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

