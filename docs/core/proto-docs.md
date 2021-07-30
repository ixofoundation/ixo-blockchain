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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| cancelled | [bool](#bool) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="bonds.Batch"></a>

### Batch



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| spot_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.BuyOrder"></a>

### BuyOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.FunctionParam"></a>

### FunctionParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| param | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="bonds.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reserved_bond_tokens | [string](#string) | repeated |  |






<a name="bonds.SellOrder"></a>

### SellOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |






<a name="bonds.SwapOrder"></a>

### SwapOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#bonds.BaseOrder) |  |  |
| to_token | [string](#string) |  |  |





 

 

 

 



<a name="bonds/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/genesis.proto



<a name="bonds.GenesisState"></a>

### GenesisState



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryAlphaMaximumsResponse"></a>

### QueryAlphaMaximumsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_system_alpha_increase | [string](#string) |  |  |
| max_system_alpha | [string](#string) |  |  |






<a name="bonds.QueryAvailableReserveRequest"></a>

### QueryAvailableReserveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryAvailableReserveResponse"></a>

### QueryAvailableReserveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryBatchRequest"></a>

### QueryBatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryBatchResponse"></a>

### QueryBatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#bonds.Batch) |  |  |






<a name="bonds.QueryBondRequest"></a>

### QueryBondRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryBondResponse"></a>

### QueryBondResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#bonds.Bond) |  |  |






<a name="bonds.QueryBondsDetailedRequest"></a>

### QueryBondsDetailedRequest







<a name="bonds.QueryBondsDetailedResponse"></a>

### QueryBondsDetailedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds_detailed | [BondDetails](#bonds.BondDetails) | repeated |  |






<a name="bonds.QueryBondsRequest"></a>

### QueryBondsRequest







<a name="bonds.QueryBondsResponse"></a>

### QueryBondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [string](#string) | repeated |  |






<a name="bonds.QueryBuyPriceRequest"></a>

### QueryBuyPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QueryBuyPriceResponse"></a>

### QueryBuyPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="bonds.QueryCustomPriceRequest"></a>

### QueryCustomPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QueryCustomPriceResponse"></a>

### QueryCustomPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryLastBatchRequest"></a>

### QueryLastBatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryLastBatchResponse"></a>

### QueryLastBatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_batch | [Batch](#bonds.Batch) |  |  |






<a name="bonds.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="bonds.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#bonds.Params) |  |  |






<a name="bonds.QuerySellReturnRequest"></a>

### QuerySellReturnRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="bonds.QuerySellReturnResponse"></a>

### QuerySellReturnResponse



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| from_token_with_amount | [string](#string) |  |  |
| to_token | [string](#string) |  |  |






<a name="bonds.QuerySwapReturnResponse"></a>

### QuerySwapReturnResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 


<a name="bonds.Query"></a>

### Query
To get a list of all module queries, go to the module&#39;s keeper/querier.go
and check all cases in NewQuerier(). REST endpoints taken from bonds/client/rest/query.go

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bonds | [QueryBondsRequest](#bonds.QueryBondsRequest) | [QueryBondsResponse](#bonds.QueryBondsResponse) |  |
| BondsDetailed | [QueryBondsDetailedRequest](#bonds.QueryBondsDetailedRequest) | [QueryBondsDetailedResponse](#bonds.QueryBondsDetailedResponse) |  |
| Params | [QueryParamsRequest](#bonds.QueryParamsRequest) | [QueryParamsResponse](#bonds.QueryParamsResponse) |  |
| Bond | [QueryBondRequest](#bonds.QueryBondRequest) | [QueryBondResponse](#bonds.QueryBondResponse) |  |
| Batch | [QueryBatchRequest](#bonds.QueryBatchRequest) | [QueryBatchResponse](#bonds.QueryBatchResponse) |  |
| LastBatch | [QueryLastBatchRequest](#bonds.QueryLastBatchRequest) | [QueryLastBatchResponse](#bonds.QueryLastBatchResponse) |  |
| CurrentPrice | [QueryCurrentPriceRequest](#bonds.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#bonds.QueryCurrentPriceResponse) |  |
| CurrentReserve | [QueryCurrentReserveRequest](#bonds.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#bonds.QueryCurrentReserveResponse) |  |
| AvailableReserve | [QueryAvailableReserveRequest](#bonds.QueryAvailableReserveRequest) | [QueryAvailableReserveResponse](#bonds.QueryAvailableReserveResponse) |  |
| CustomPrice | [QueryCustomPriceRequest](#bonds.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#bonds.QueryCustomPriceResponse) |  |
| BuyPrice | [QueryBuyPriceRequest](#bonds.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#bonds.QueryBuyPriceResponse) |  |
| SellReturn | [QuerySellReturnRequest](#bonds.QuerySellReturnRequest) | [QuerySellReturnResponse](#bonds.QuerySellReturnResponse) |  |
| SwapReturn | [QuerySwapReturnRequest](#bonds.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#bonds.QuerySwapReturnResponse) |  |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#bonds.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#bonds.QueryAlphaMaximumsResponse) |  |

 



<a name="bonds/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bonds/tx.proto



<a name="bonds.MsgBuy"></a>

### MsgBuy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgBuyResponse"></a>

### MsgBuyResponse







<a name="bonds.MsgCreateBond"></a>

### MsgCreateBond



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







<a name="bonds.MsgEditBond"></a>

### MsgEditBond



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







<a name="bonds.MsgMakeOutcomePayment"></a>

### MsgMakeOutcomePayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgMakeOutcomePaymentResponse"></a>

### MsgMakeOutcomePaymentResponse







<a name="bonds.MsgSell"></a>

### MsgSell



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seller_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgSellResponse"></a>

### MsgSellResponse







<a name="bonds.MsgSetNextAlpha"></a>

### MsgSetNextAlpha



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| alpha | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |






<a name="bonds.MsgSetNextAlphaResponse"></a>

### MsgSetNextAlphaResponse







<a name="bonds.MsgSwap"></a>

### MsgSwap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swapper_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| from | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| to_token | [string](#string) |  |  |






<a name="bonds.MsgSwapResponse"></a>

### MsgSwapResponse







<a name="bonds.MsgUpdateBondState"></a>

### MsgUpdateBondState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| state | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |






<a name="bonds.MsgUpdateBondStateResponse"></a>

### MsgUpdateBondStateResponse







<a name="bonds.MsgWithdrawReserve"></a>

### MsgWithdrawReserve



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdrawer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgWithdrawReserveResponse"></a>

### MsgWithdrawReserveResponse







<a name="bonds.MsgWithdrawShare"></a>

### MsgWithdrawShare



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |






<a name="bonds.MsgWithdrawShareResponse"></a>

### MsgWithdrawShareResponse






 

 

 


<a name="bonds.Msg"></a>

### Msg
To get a list of all module messages, go to your module&#39;s handler.go and
check all cases in NewHandler().

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBond | [MsgCreateBond](#bonds.MsgCreateBond) | [MsgCreateBondResponse](#bonds.MsgCreateBondResponse) |  |
| EditBond | [MsgEditBond](#bonds.MsgEditBond) | [MsgEditBondResponse](#bonds.MsgEditBondResponse) |  |
| SetNextAlpha | [MsgSetNextAlpha](#bonds.MsgSetNextAlpha) | [MsgSetNextAlphaResponse](#bonds.MsgSetNextAlphaResponse) |  |
| UpdateBondState | [MsgUpdateBondState](#bonds.MsgUpdateBondState) | [MsgUpdateBondStateResponse](#bonds.MsgUpdateBondStateResponse) |  |
| Buy | [MsgBuy](#bonds.MsgBuy) | [MsgBuyResponse](#bonds.MsgBuyResponse) |  |
| Sell | [MsgSell](#bonds.MsgSell) | [MsgSellResponse](#bonds.MsgSellResponse) |  |
| Swap | [MsgSwap](#bonds.MsgSwap) | [MsgSwapResponse](#bonds.MsgSwapResponse) |  |
| MakeOutcomePayment | [MsgMakeOutcomePayment](#bonds.MsgMakeOutcomePayment) | [MsgMakeOutcomePaymentResponse](#bonds.MsgMakeOutcomePaymentResponse) |  |
| WithdrawShare | [MsgWithdrawShare](#bonds.MsgWithdrawShare) | [MsgWithdrawShareResponse](#bonds.MsgWithdrawShareResponse) |  |
| WithdrawReserve | [MsgWithdrawReserve](#bonds.MsgWithdrawReserve) | [MsgWithdrawReserveResponse](#bonds.MsgWithdrawReserveResponse) |  |

 



<a name="did/did.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/did.proto



<a name="did.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| KYC_validated | [bool](#bool) |  |  |






<a name="did.DidCredential"></a>

### DidCredential



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cred_type | [string](#string) | repeated |  |
| issuer | [string](#string) |  |  |
| issued | [string](#string) |  |  |
| claim | [Claim](#did.Claim) |  |  |






<a name="did.IxoDid"></a>

### IxoDid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| verify_key | [string](#string) |  |  |
| encryption_public_key | [string](#string) |  |  |
| secret | [Secret](#did.Secret) |  |  |






<a name="did.Secret"></a>

### Secret



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did_docs | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |





 

 

 

 



<a name="did/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/query.proto



<a name="did.QueryAddressFromBase58EncodedPubkeyRequest"></a>

### QueryAddressFromBase58EncodedPubkeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubKey | [string](#string) |  |  |






<a name="did.QueryAddressFromBase58EncodedPubkeyResponse"></a>

### QueryAddressFromBase58EncodedPubkeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="did.QueryAddressFromDidRequest"></a>

### QueryAddressFromDidRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |






<a name="did.QueryAddressFromDidResponse"></a>

### QueryAddressFromDidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="did.QueryAllDidDocsRequest"></a>

### QueryAllDidDocsRequest







<a name="did.QueryAllDidDocsResponse"></a>

### QueryAllDidDocsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="did.QueryAllDidsRequest"></a>

### QueryAllDidsRequest







<a name="did.QueryAllDidsResponse"></a>

### QueryAllDidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dids | [string](#string) | repeated |  |






<a name="did.QueryDidDocRequest"></a>

### QueryDidDocRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |






<a name="did.QueryDidDocResponse"></a>

### QueryDidDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddoc | [google.protobuf.Any](#google.protobuf.Any) |  |  |





 

 

 


<a name="did.Query"></a>

### Query
To get a list of all module queries, go to your module&#39;s keeper/querier.go
and check all cases in NewQuerier(). REST endpoints taken from did/client/rest/query.go

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DidDoc | [QueryDidDocRequest](#did.QueryDidDocRequest) | [QueryDidDocResponse](#did.QueryDidDocResponse) |  |
| AllDids | [QueryAllDidsRequest](#did.QueryAllDidsRequest) | [QueryAllDidsResponse](#did.QueryAllDidsResponse) |  |
| AllDidDocs | [QueryAllDidDocsRequest](#did.QueryAllDidDocsRequest) | [QueryAllDidDocsResponse](#did.QueryAllDidDocsResponse) |  |
| AddressFromDid | [QueryAddressFromDidRequest](#did.QueryAddressFromDidRequest) | [QueryAddressFromDidResponse](#did.QueryAddressFromDidResponse) |  |
| AddressFromBase58EncodedPubkey | [QueryAddressFromBase58EncodedPubkeyRequest](#did.QueryAddressFromBase58EncodedPubkeyRequest) | [QueryAddressFromBase58EncodedPubkeyResponse](#did.QueryAddressFromBase58EncodedPubkeyResponse) |  |

 



<a name="did/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/tx.proto



<a name="did.MsgAddCredential"></a>

### MsgAddCredential



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did_credential | [DidCredential](#did.DidCredential) |  |  |






<a name="did.MsgAddCredentialResponse"></a>

### MsgAddCredentialResponse







<a name="did.MsgAddDid"></a>

### MsgAddDid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pubKey | [string](#string) |  |  |






<a name="did.MsgAddDidResponse"></a>

### MsgAddDidResponse






 

 

 


<a name="did.Msg"></a>

### Msg
To get a list of all module messages, go to your module&#39;s handler.go and
check all cases in NewHandler().

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddDid | [MsgAddDid](#did.MsgAddDid) | [MsgAddDidResponse](#did.MsgAddDidResponse) |  |
| AddCredential | [MsgAddCredential](#did.MsgAddCredential) | [MsgAddCredentialResponse](#did.MsgAddCredentialResponse) |  |

 



<a name="payments/payments.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/payments.proto



<a name="payments.BlockPeriod"></a>

### BlockPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="payments.Discount"></a>

### Discount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| percent | [string](#string) |  |  |






<a name="payments.DistributionShare"></a>

### DistributionShare



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| percentage | [string](#string) |  |  |






<a name="payments.PaymentContract"></a>

### PaymentContract



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_minimum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_maximum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| discounts | [Discount](#payments.Discount) | repeated |  |






<a name="payments.Subscription"></a>

### Subscription



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="payments.TimePeriod"></a>

### TimePeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| period_start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="payments/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/genesis.proto



<a name="payments.GenesisState"></a>

### GenesisState



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |






<a name="payments.QueryPaymentContractResponse"></a>

### QueryPaymentContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract | [PaymentContract](#payments.PaymentContract) |  |  |






<a name="payments.QueryPaymentContractsByIdPrefixRequest"></a>

### QueryPaymentContractsByIdPrefixRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts_id_prefix | [string](#string) |  |  |






<a name="payments.QueryPaymentContractsByIdPrefixResponse"></a>

### QueryPaymentContractsByIdPrefixResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts | [PaymentContract](#payments.PaymentContract) | repeated |  |






<a name="payments.QueryPaymentTemplateRequest"></a>

### QueryPaymentTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template_id | [string](#string) |  |  |






<a name="payments.QueryPaymentTemplateResponse"></a>

### QueryPaymentTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template | [PaymentTemplate](#payments.PaymentTemplate) |  |  |






<a name="payments.QuerySubscriptionRequest"></a>

### QuerySubscriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |






<a name="payments.QuerySubscriptionResponse"></a>

### QuerySubscriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#payments.Subscription) |  |  |





 

 

 


<a name="payments.Query"></a>

### Query
To get a list of all module queries, go to your module&#39;s keeper/querier.go
and check all cases in NewQuerier(). REST endpoints taken from payments/client/rest/query.go

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PaymentTemplate | [QueryPaymentTemplateRequest](#payments.QueryPaymentTemplateRequest) | [QueryPaymentTemplateResponse](#payments.QueryPaymentTemplateResponse) |  |
| PaymentContract | [QueryPaymentContractRequest](#payments.QueryPaymentContractRequest) | [QueryPaymentContractResponse](#payments.QueryPaymentContractResponse) |  |
| PaymentContractsByIdPrefix | [QueryPaymentContractsByIdPrefixRequest](#payments.QueryPaymentContractsByIdPrefixRequest) | [QueryPaymentContractsByIdPrefixResponse](#payments.QueryPaymentContractsByIdPrefixResponse) |  |
| Subscription | [QuerySubscriptionRequest](#payments.QuerySubscriptionRequest) | [QuerySubscriptionResponse](#payments.QuerySubscriptionResponse) |  |

 



<a name="payments/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments/tx.proto



<a name="payments.MsgCreatePaymentContract"></a>

### MsgCreatePaymentContract



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







<a name="payments.MsgCreatePaymentTemplate"></a>

### MsgCreatePaymentTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| payment_template | [PaymentTemplate](#payments.PaymentTemplate) |  |  |






<a name="payments.MsgCreatePaymentTemplateResponse"></a>

### MsgCreatePaymentTemplateResponse







<a name="payments.MsgCreateSubscription"></a>

### MsgCreateSubscription



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| subscription_id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| max_periods | [string](#string) |  |  |
| period | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="payments.MsgCreateSubscriptionResponse"></a>

### MsgCreateSubscriptionResponse







<a name="payments.MsgEffectPayment"></a>

### MsgEffectPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |






<a name="payments.MsgEffectPaymentResponse"></a>

### MsgEffectPaymentResponse







<a name="payments.MsgGrantDiscount"></a>

### MsgGrantDiscount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| discount_id | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="payments.MsgGrantDiscountResponse"></a>

### MsgGrantDiscountResponse







<a name="payments.MsgRevokeDiscount"></a>

### MsgRevokeDiscount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| holder | [string](#string) |  |  |






<a name="payments.MsgRevokeDiscountResponse"></a>

### MsgRevokeDiscountResponse







<a name="payments.MsgSetPaymentContractAuthorisation"></a>

### MsgSetPaymentContractAuthorisation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |
| payer_did | [string](#string) |  |  |
| authorised | [bool](#bool) |  |  |






<a name="payments.MsgSetPaymentContractAuthorisationResponse"></a>

### MsgSetPaymentContractAuthorisationResponse






 

 

 


<a name="payments.Msg"></a>

### Msg
To get a list of all module messages, go to your module&#39;s handler.go and
check all cases in NewHandler().

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetPaymentContractAuthorisation | [MsgSetPaymentContractAuthorisation](#payments.MsgSetPaymentContractAuthorisation) | [MsgSetPaymentContractAuthorisationResponse](#payments.MsgSetPaymentContractAuthorisationResponse) |  |
| CreatePaymentTemplate | [MsgCreatePaymentTemplate](#payments.MsgCreatePaymentTemplate) | [MsgCreatePaymentTemplateResponse](#payments.MsgCreatePaymentTemplateResponse) |  |
| CreatePaymentContract | [MsgCreatePaymentContract](#payments.MsgCreatePaymentContract) | [MsgCreatePaymentContractResponse](#payments.MsgCreatePaymentContractResponse) |  |
| CreateSubscription | [MsgCreateSubscription](#payments.MsgCreateSubscription) | [MsgCreateSubscriptionResponse](#payments.MsgCreateSubscriptionResponse) |  |
| GrantDiscount | [MsgGrantDiscount](#payments.MsgGrantDiscount) | [MsgGrantDiscountResponse](#payments.MsgGrantDiscountResponse) |  |
| RevokeDiscount | [MsgRevokeDiscount](#payments.MsgRevokeDiscount) | [MsgRevokeDiscountResponse](#payments.MsgRevokeDiscountResponse) |  |
| EffectPayment | [MsgEffectPayment](#payments.MsgEffectPayment) | [MsgEffectPaymentResponse](#payments.MsgEffectPaymentResponse) |  |

 



<a name="project/project.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/project.proto



<a name="project.AccountMap"></a>

### AccountMap



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| template_id | [string](#string) |  |  |
| claimer_did | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="project.Claims"></a>

### Claims



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims_list | [Claim](#project.Claim) | repeated |  |






<a name="project.CreateAgentDoc"></a>

### CreateAgentDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_did | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="project.CreateClaimDoc"></a>

### CreateClaimDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| claim_template_id | [string](#string) |  |  |






<a name="project.CreateEvaluationDoc"></a>

### CreateEvaluationDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="project.GenesisAccountMap"></a>

### GenesisAccountMap



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo_did | [string](#string) |  |  |
| project_minimum_initial_funding | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| oracle_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |






<a name="project.ProjectDoc"></a>

### ProjectDoc



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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| status | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="project.UpdateProjectStatusDoc"></a>

### UpdateProjectStatusDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| eth_funding_txn_id | [string](#string) |  |  |






<a name="project.WithdrawFundsDoc"></a>

### WithdrawFundsDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| is_refund | [bool](#bool) |  |  |






<a name="project.WithdrawalInfoDoc"></a>

### WithdrawalInfoDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="project.WithdrawalInfoDocs"></a>

### WithdrawalInfoDocs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs_list | [WithdrawalInfoDoc](#project.WithdrawalInfoDoc) | repeated |  |





 

 

 

 



<a name="project/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/genesis.proto



<a name="project.GenesisState"></a>

### GenesisState



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







<a name="project.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#project.Params) |  |  |






<a name="project.QueryProjectAccountsRequest"></a>

### QueryProjectAccountsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectAccountsResponse"></a>

### QueryProjectAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_map | [AccountMap](#project.AccountMap) |  |  |






<a name="project.QueryProjectDocRequest"></a>

### QueryProjectDocRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectDocResponse"></a>

### QueryProjectDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_doc | [ProjectDoc](#project.ProjectDoc) |  |  |






<a name="project.QueryProjectTxRequest"></a>

### QueryProjectTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="project.QueryProjectTxResponse"></a>

### QueryProjectTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txs | [WithdrawalInfoDocs](#project.WithdrawalInfoDocs) |  |  |





 

 

 


<a name="project.Query"></a>

### Query


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ProjectDoc | [QueryProjectDocRequest](#project.QueryProjectDocRequest) | [QueryProjectDocResponse](#project.QueryProjectDocResponse) |  |
| ProjectAccounts | [QueryProjectAccountsRequest](#project.QueryProjectAccountsRequest) | [QueryProjectAccountsResponse](#project.QueryProjectAccountsResponse) |  |
| ProjectTx | [QueryProjectTxRequest](#project.QueryProjectTxRequest) | [QueryProjectTxResponse](#project.QueryProjectTxResponse) |  |
| Params | [QueryParamsRequest](#project.QueryParamsRequest) | [QueryParamsResponse](#project.QueryParamsResponse) |  |

 



<a name="project/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project/tx.proto



<a name="project.MsgCreateAgent"></a>

### MsgCreateAgent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateAgentDoc](#project.CreateAgentDoc) |  |  |






<a name="project.MsgCreateAgentResponse"></a>

### MsgCreateAgentResponse







<a name="project.MsgCreateClaim"></a>

### MsgCreateClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateClaimDoc](#project.CreateClaimDoc) |  |  |






<a name="project.MsgCreateClaimResponse"></a>

### MsgCreateClaimResponse







<a name="project.MsgCreateEvaluation"></a>

### MsgCreateEvaluation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateEvaluationDoc](#project.CreateEvaluationDoc) |  |  |






<a name="project.MsgCreateEvaluationResponse"></a>

### MsgCreateEvaluationResponse







<a name="project.MsgCreateProject"></a>

### MsgCreateProject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="project.MsgCreateProjectResponse"></a>

### MsgCreateProjectResponse







<a name="project.MsgUpdateAgent"></a>

### MsgUpdateAgent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateAgentDoc](#project.UpdateAgentDoc) |  |  |






<a name="project.MsgUpdateAgentResponse"></a>

### MsgUpdateAgentResponse







<a name="project.MsgUpdateProjectDoc"></a>

### MsgUpdateProjectDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="project.MsgUpdateProjectDocResponse"></a>

### MsgUpdateProjectDocResponse







<a name="project.MsgUpdateProjectStatus"></a>

### MsgUpdateProjectStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateProjectStatusDoc](#project.UpdateProjectStatusDoc) |  |  |






<a name="project.MsgUpdateProjectStatusResponse"></a>

### MsgUpdateProjectStatusResponse







<a name="project.MsgWithdrawFunds"></a>

### MsgWithdrawFunds



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| data | [WithdrawFundsDoc](#project.WithdrawFundsDoc) |  |  |






<a name="project.MsgWithdrawFundsResponse"></a>

### MsgWithdrawFundsResponse






 

 

 


<a name="project.Msg"></a>

### Msg
To get a list of all module messages, go to your module&#39;s handler.go and
check all cases in NewHandler().

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProject | [MsgCreateProject](#project.MsgCreateProject) | [MsgCreateProjectResponse](#project.MsgCreateProjectResponse) |  |
| UpdateProjectStatus | [MsgUpdateProjectStatus](#project.MsgUpdateProjectStatus) | [MsgUpdateProjectStatusResponse](#project.MsgUpdateProjectStatusResponse) |  |
| CreateAgent | [MsgCreateAgent](#project.MsgCreateAgent) | [MsgCreateAgentResponse](#project.MsgCreateAgentResponse) |  |
| UpdateAgent | [MsgUpdateAgent](#project.MsgUpdateAgent) | [MsgUpdateAgentResponse](#project.MsgUpdateAgentResponse) |  |
| CreateClaim | [MsgCreateClaim](#project.MsgCreateClaim) | [MsgCreateClaimResponse](#project.MsgCreateClaimResponse) |  |
| CreateEvaluation | [MsgCreateEvaluation](#project.MsgCreateEvaluation) | [MsgCreateEvaluationResponse](#project.MsgCreateEvaluationResponse) |  |
| WithdrawFunds | [MsgWithdrawFunds](#project.MsgWithdrawFunds) | [MsgWithdrawFundsResponse](#project.MsgWithdrawFundsResponse) |  |
| UpdateProjectDoc | [MsgUpdateProjectDoc](#project.MsgUpdateProjectDoc) | [MsgUpdateProjectDocResponse](#project.MsgUpdateProjectDocResponse) |  |

 



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

