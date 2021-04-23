# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [cosmos/base/v1beta1/coin.proto](#cosmos/base/v1beta1/coin.proto)
    - [Coin](#cosmos.base.v1beta1.Coin)
    - [DecCoin](#cosmos.base.v1beta1.DecCoin)
    - [DecProto](#cosmos.base.v1beta1.DecProto)
    - [IntProto](#cosmos.base.v1beta1.IntProto)
  
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
    - [MsgWithdrawShare](#bonds.MsgWithdrawShare)
    - [MsgWithdrawShareResponse](#bonds.MsgWithdrawShareResponse)
  
    - [Msg](#bonds.Msg)
  
- [cosmos/base/abci/v1beta1/abci.proto](#cosmos/base/abci/v1beta1/abci.proto)
    - [ABCIMessageLog](#cosmos.base.abci.v1beta1.ABCIMessageLog)
    - [Attribute](#cosmos.base.abci.v1beta1.Attribute)
    - [GasInfo](#cosmos.base.abci.v1beta1.GasInfo)
    - [MsgData](#cosmos.base.abci.v1beta1.MsgData)
    - [Result](#cosmos.base.abci.v1beta1.Result)
    - [SearchTxsResult](#cosmos.base.abci.v1beta1.SearchTxsResult)
    - [SimulationResponse](#cosmos.base.abci.v1beta1.SimulationResponse)
    - [StringEvent](#cosmos.base.abci.v1beta1.StringEvent)
    - [TxMsgData](#cosmos.base.abci.v1beta1.TxMsgData)
    - [TxResponse](#cosmos.base.abci.v1beta1.TxResponse)
  
- [cosmos/base/kv/v1beta1/kv.proto](#cosmos/base/kv/v1beta1/kv.proto)
    - [Pair](#cosmos.base.kv.v1beta1.Pair)
    - [Pairs](#cosmos.base.kv.v1beta1.Pairs)
  
- [cosmos/base/query/v1beta1/pagination.proto](#cosmos/base/query/v1beta1/pagination.proto)
    - [PageRequest](#cosmos.base.query.v1beta1.PageRequest)
    - [PageResponse](#cosmos.base.query.v1beta1.PageResponse)
  
- [cosmos/base/reflection/v1beta1/reflection.proto](#cosmos/base/reflection/v1beta1/reflection.proto)
    - [ListAllInterfacesRequest](#cosmos.base.reflection.v1beta1.ListAllInterfacesRequest)
    - [ListAllInterfacesResponse](#cosmos.base.reflection.v1beta1.ListAllInterfacesResponse)
    - [ListImplementationsRequest](#cosmos.base.reflection.v1beta1.ListImplementationsRequest)
    - [ListImplementationsResponse](#cosmos.base.reflection.v1beta1.ListImplementationsResponse)
  
    - [ReflectionService](#cosmos.base.reflection.v1beta1.ReflectionService)
  
- [cosmos/base/snapshots/v1beta1/snapshot.proto](#cosmos/base/snapshots/v1beta1/snapshot.proto)
    - [Metadata](#cosmos.base.snapshots.v1beta1.Metadata)
    - [Snapshot](#cosmos.base.snapshots.v1beta1.Snapshot)
  
- [cosmos/base/store/v1beta1/commit_info.proto](#cosmos/base/store/v1beta1/commit_info.proto)
    - [CommitID](#cosmos.base.store.v1beta1.CommitID)
    - [CommitInfo](#cosmos.base.store.v1beta1.CommitInfo)
    - [StoreInfo](#cosmos.base.store.v1beta1.StoreInfo)
  
- [cosmos/base/store/v1beta1/snapshot.proto](#cosmos/base/store/v1beta1/snapshot.proto)
    - [SnapshotIAVLItem](#cosmos.base.store.v1beta1.SnapshotIAVLItem)
    - [SnapshotItem](#cosmos.base.store.v1beta1.SnapshotItem)
    - [SnapshotStoreItem](#cosmos.base.store.v1beta1.SnapshotStoreItem)
  
- [cosmos/base/tendermint/v1beta1/query.proto](#cosmos/base/tendermint/v1beta1/query.proto)
    - [GetBlockByHeightRequest](#cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest)
    - [GetBlockByHeightResponse](#cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse)
    - [GetLatestBlockRequest](#cosmos.base.tendermint.v1beta1.GetLatestBlockRequest)
    - [GetLatestBlockResponse](#cosmos.base.tendermint.v1beta1.GetLatestBlockResponse)
    - [GetLatestValidatorSetRequest](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest)
    - [GetLatestValidatorSetResponse](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse)
    - [GetNodeInfoRequest](#cosmos.base.tendermint.v1beta1.GetNodeInfoRequest)
    - [GetNodeInfoResponse](#cosmos.base.tendermint.v1beta1.GetNodeInfoResponse)
    - [GetSyncingRequest](#cosmos.base.tendermint.v1beta1.GetSyncingRequest)
    - [GetSyncingResponse](#cosmos.base.tendermint.v1beta1.GetSyncingResponse)
    - [GetValidatorSetByHeightRequest](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest)
    - [GetValidatorSetByHeightResponse](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse)
    - [Module](#cosmos.base.tendermint.v1beta1.Module)
    - [Validator](#cosmos.base.tendermint.v1beta1.Validator)
    - [VersionInfo](#cosmos.base.tendermint.v1beta1.VersionInfo)
  
    - [Service](#cosmos.base.tendermint.v1beta1.Service)
  
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
  
- [Scalar Value Types](#scalar-value-types)



<a name="cosmos/base/v1beta1/coin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/v1beta1/coin.proto



<a name="cosmos.base.v1beta1.Coin"></a>

### Coin
Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |
| amount | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.DecCoin"></a>

### DecCoin
DecCoin defines a token with a denomination and a decimal amount.

NOTE: The amount field is an Dec which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |
| amount | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.DecProto"></a>

### DecProto
DecProto defines a Protobuf wrapper around a Dec object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dec | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.IntProto"></a>

### IntProto
IntProto defines a Protobuf wrapper around an Int object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| int | [string](#string) |  |  |





 

 

 

 



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
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| current_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_outcome_payment_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| allow_sells | [bool](#bool) |  |  |
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
GenesisState defines the did module&#39;s genesis state.


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
Request/response types from old x/bonds/client/cli/query.go and
x/bonds/client/rest/query.go






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
| buy_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






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
| dec_coins | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="bonds.QueryLastBatchRequest"></a>

### QueryLastBatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="bonds.QueryLastBatchResponse"></a>

### QueryLastBatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#bonds.Batch) |  |  |






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
| Bond | [QueryBondRequest](#bonds.QueryBondRequest) | [QueryBondResponse](#bonds.QueryBondResponse) |  |
| Batch | [QueryBatchRequest](#bonds.QueryBatchRequest) | [QueryBatchResponse](#bonds.QueryBatchResponse) |  |
| LastBatch | [QueryLastBatchRequest](#bonds.QueryLastBatchRequest) | [QueryLastBatchResponse](#bonds.QueryLastBatchResponse) |  |
| CurrentPrice | [QueryCurrentPriceRequest](#bonds.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#bonds.QueryCurrentPriceResponse) |  |
| CurrentReserve | [QueryCurrentReserveRequest](#bonds.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#bonds.QueryCurrentReserveResponse) |  |
| CustomPrice | [QueryCustomPriceRequest](#bonds.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#bonds.QueryCustomPriceResponse) |  |
| BuyPrice | [QueryBuyPriceRequest](#bonds.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#bonds.QueryBuyPriceResponse) |  |
| SellReturn | [QuerySellReturnRequest](#bonds.QuerySellReturnRequest) | [QuerySellReturnResponse](#bonds.QuerySellReturnResponse) |  |
| SwapReturn | [QuerySwapReturnRequest](#bonds.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#bonds.QuerySwapReturnResponse) |  |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#bonds.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#bonds.QueryAlphaMaximumsResponse) |  |
| Params | [QueryParamsRequest](#bonds.QueryParamsRequest) | [QueryParamsResponse](#bonds.QueryParamsResponse) |  |

 



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
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| allow_sells | [bool](#bool) |  |  |
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

 



<a name="cosmos/base/abci/v1beta1/abci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/abci/v1beta1/abci.proto



<a name="cosmos.base.abci.v1beta1.ABCIMessageLog"></a>

### ABCIMessageLog
ABCIMessageLog defines a structure containing an indexed tx ABCI message log.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg_index | [uint32](#uint32) |  |  |
| log | [string](#string) |  |  |
| events | [StringEvent](#cosmos.base.abci.v1beta1.StringEvent) | repeated | Events contains a slice of Event objects that were emitted during some execution. |






<a name="cosmos.base.abci.v1beta1.Attribute"></a>

### Attribute
Attribute defines an attribute wrapper where the key and value are
strings instead of raw bytes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cosmos.base.abci.v1beta1.GasInfo"></a>

### GasInfo
GasInfo defines tx execution gas context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gas_wanted | [uint64](#uint64) |  | GasWanted is the maximum units of work we allow this tx to perform. |
| gas_used | [uint64](#uint64) |  | GasUsed is the amount of gas actually consumed. |






<a name="cosmos.base.abci.v1beta1.MsgData"></a>

### MsgData
MsgData defines the data returned in a Result object during message
execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg_type | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="cosmos.base.abci.v1beta1.Result"></a>

### Result
Result is the union of ResponseFormat and ResponseCheckTx.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Data is any data returned from message or handler execution. It MUST be length prefixed in order to separate data from multiple message executions. |
| log | [string](#string) |  | Log contains the log information from message or handler execution. |
| events | [tendermint.abci.Event](#tendermint.abci.Event) | repeated | Events contains a slice of Event objects that were emitted during message or handler execution. |






<a name="cosmos.base.abci.v1beta1.SearchTxsResult"></a>

### SearchTxsResult
SearchTxsResult defines a structure for querying txs pageable


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_count | [uint64](#uint64) |  | Count of all txs |
| count | [uint64](#uint64) |  | Count of txs in current page |
| page_number | [uint64](#uint64) |  | Index of current page, start from 1 |
| page_total | [uint64](#uint64) |  | Count of total pages |
| limit | [uint64](#uint64) |  | Max count txs per page |
| txs | [TxResponse](#cosmos.base.abci.v1beta1.TxResponse) | repeated | List of txs in current page |






<a name="cosmos.base.abci.v1beta1.SimulationResponse"></a>

### SimulationResponse
SimulationResponse defines the response generated when a transaction is
successfully simulated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gas_info | [GasInfo](#cosmos.base.abci.v1beta1.GasInfo) |  |  |
| result | [Result](#cosmos.base.abci.v1beta1.Result) |  |  |






<a name="cosmos.base.abci.v1beta1.StringEvent"></a>

### StringEvent
StringEvent defines en Event object wrapper where all the attributes
contain key/value pairs that are strings instead of raw bytes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| attributes | [Attribute](#cosmos.base.abci.v1beta1.Attribute) | repeated |  |






<a name="cosmos.base.abci.v1beta1.TxMsgData"></a>

### TxMsgData
TxMsgData defines a list of MsgData. A transaction will have a MsgData object
for each message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [MsgData](#cosmos.base.abci.v1beta1.MsgData) | repeated |  |






<a name="cosmos.base.abci.v1beta1.TxResponse"></a>

### TxResponse
TxResponse defines a structure containing relevant tx data and metadata. The
tags are stringified and the log is JSON decoded.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [int64](#int64) |  | The block height |
| txhash | [string](#string) |  | The transaction hash. |
| codespace | [string](#string) |  | Namespace for the Code |
| code | [uint32](#uint32) |  | Response code. |
| data | [string](#string) |  | Result bytes, if any. |
| raw_log | [string](#string) |  | The output of the application&#39;s logger (raw string). May be non-deterministic. |
| logs | [ABCIMessageLog](#cosmos.base.abci.v1beta1.ABCIMessageLog) | repeated | The output of the application&#39;s logger (typed). May be non-deterministic. |
| info | [string](#string) |  | Additional information. May be non-deterministic. |
| gas_wanted | [int64](#int64) |  | Amount of gas requested for transaction. |
| gas_used | [int64](#int64) |  | Amount of gas consumed by transaction. |
| tx | [google.protobuf.Any](#google.protobuf.Any) |  | The request transaction bytes. |
| timestamp | [string](#string) |  | Time of the previous block. For heights &gt; 1, it&#39;s the weighted median of the timestamps of the valid votes in the block.LastCommit. For height == 1, it&#39;s genesis time. |





 

 

 

 



<a name="cosmos/base/kv/v1beta1/kv.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/kv/v1beta1/kv.proto



<a name="cosmos.base.kv.v1beta1.Pair"></a>

### Pair
Pair defines a key/value bytes tuple.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [bytes](#bytes) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="cosmos.base.kv.v1beta1.Pairs"></a>

### Pairs
Pairs defines a repeated slice of Pair objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pairs | [Pair](#cosmos.base.kv.v1beta1.Pair) | repeated |  |





 

 

 

 



<a name="cosmos/base/query/v1beta1/pagination.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/query/v1beta1/pagination.proto



<a name="cosmos.base.query.v1beta1.PageRequest"></a>

### PageRequest
PageRequest is to be embedded in gRPC request messages for efficient
pagination. Ex:

 message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [bytes](#bytes) |  | key is a value returned in PageResponse.next_key to begin querying the next page most efficiently. Only one of offset or key should be set. |
| offset | [uint64](#uint64) |  | offset is a numeric offset that can be used when key is unavailable. It is less efficient than using key. Only one of offset or key should be set. |
| limit | [uint64](#uint64) |  | limit is the total number of results to be returned in the result page. If left empty it will default to a value to be set by each app. |
| count_total | [bool](#bool) |  | count_total is set to true to indicate that the result set should include a count of the total number of items available for pagination in UIs. count_total is only respected when offset is used. It is ignored when key is set. |






<a name="cosmos.base.query.v1beta1.PageResponse"></a>

### PageResponse
PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_key | [bytes](#bytes) |  | next_key is the key to be passed to PageRequest.key to query the next page most efficiently |
| total | [uint64](#uint64) |  | total is total number of results available if PageRequest.count_total was set, its value is undefined otherwise |





 

 

 

 



<a name="cosmos/base/reflection/v1beta1/reflection.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/reflection/v1beta1/reflection.proto



<a name="cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"></a>

### ListAllInterfacesRequest
ListAllInterfacesRequest is the request type of the ListAllInterfaces RPC.






<a name="cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"></a>

### ListAllInterfacesResponse
ListAllInterfacesResponse is the response type of the ListAllInterfaces RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_names | [string](#string) | repeated | interface_names is an array of all the registered interfaces. |






<a name="cosmos.base.reflection.v1beta1.ListImplementationsRequest"></a>

### ListImplementationsRequest
ListImplementationsRequest is the request type of the ListImplementations
RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interface_name | [string](#string) |  | interface_name defines the interface to query the implementations for. |






<a name="cosmos.base.reflection.v1beta1.ListImplementationsResponse"></a>

### ListImplementationsResponse
ListImplementationsResponse is the response type of the ListImplementations
RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| implementation_message_names | [string](#string) | repeated |  |





 

 

 


<a name="cosmos.base.reflection.v1beta1.ReflectionService"></a>

### ReflectionService
ReflectionService defines a service for interface reflection.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListAllInterfaces | [ListAllInterfacesRequest](#cosmos.base.reflection.v1beta1.ListAllInterfacesRequest) | [ListAllInterfacesResponse](#cosmos.base.reflection.v1beta1.ListAllInterfacesResponse) | ListAllInterfaces lists all the interfaces registered in the interface registry. |
| ListImplementations | [ListImplementationsRequest](#cosmos.base.reflection.v1beta1.ListImplementationsRequest) | [ListImplementationsResponse](#cosmos.base.reflection.v1beta1.ListImplementationsResponse) | ListImplementations list all the concrete types that implement a given interface. |

 



<a name="cosmos/base/snapshots/v1beta1/snapshot.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/snapshots/v1beta1/snapshot.proto



<a name="cosmos.base.snapshots.v1beta1.Metadata"></a>

### Metadata
Metadata contains SDK-specific snapshot metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chunk_hashes | [bytes](#bytes) | repeated | SHA-256 chunk hashes |






<a name="cosmos.base.snapshots.v1beta1.Snapshot"></a>

### Snapshot
Snapshot contains Tendermint state sync snapshot info.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| format | [uint32](#uint32) |  |  |
| chunks | [uint32](#uint32) |  |  |
| hash | [bytes](#bytes) |  |  |
| metadata | [Metadata](#cosmos.base.snapshots.v1beta1.Metadata) |  |  |





 

 

 

 



<a name="cosmos/base/store/v1beta1/commit_info.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/store/v1beta1/commit_info.proto



<a name="cosmos.base.store.v1beta1.CommitID"></a>

### CommitID
CommitID defines the committment information when a specific store is
committed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [int64](#int64) |  |  |
| hash | [bytes](#bytes) |  |  |






<a name="cosmos.base.store.v1beta1.CommitInfo"></a>

### CommitInfo
CommitInfo defines commit information used by the multi-store when committing
a version/height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [int64](#int64) |  |  |
| store_infos | [StoreInfo](#cosmos.base.store.v1beta1.StoreInfo) | repeated |  |






<a name="cosmos.base.store.v1beta1.StoreInfo"></a>

### StoreInfo
StoreInfo defines store-specific commit information. It contains a reference
between a store name and the commit ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| commit_id | [CommitID](#cosmos.base.store.v1beta1.CommitID) |  |  |





 

 

 

 



<a name="cosmos/base/store/v1beta1/snapshot.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/store/v1beta1/snapshot.proto



<a name="cosmos.base.store.v1beta1.SnapshotIAVLItem"></a>

### SnapshotIAVLItem
SnapshotIAVLItem is an exported IAVL node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [bytes](#bytes) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| height | [int32](#int32) |  |  |






<a name="cosmos.base.store.v1beta1.SnapshotItem"></a>

### SnapshotItem
SnapshotItem is an item contained in a rootmulti.Store snapshot.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [SnapshotStoreItem](#cosmos.base.store.v1beta1.SnapshotStoreItem) |  |  |
| iavl | [SnapshotIAVLItem](#cosmos.base.store.v1beta1.SnapshotIAVLItem) |  |  |






<a name="cosmos.base.store.v1beta1.SnapshotStoreItem"></a>

### SnapshotStoreItem
SnapshotStoreItem contains metadata about a snapshotted store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |





 

 

 

 



<a name="cosmos/base/tendermint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/tendermint/v1beta1/query.proto



<a name="cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"></a>

### GetBlockByHeightRequest
GetBlockByHeightRequest is the request type for the Query/GetBlockByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [int64](#int64) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"></a>

### GetBlockByHeightResponse
GetBlockByHeightResponse is the response type for the Query/GetBlockByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_id | [tendermint.types.BlockID](#tendermint.types.BlockID) |  |  |
| block | [tendermint.types.Block](#tendermint.types.Block) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"></a>

### GetLatestBlockRequest
GetLatestBlockRequest is the request type for the Query/GetLatestBlock RPC method.






<a name="cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"></a>

### GetLatestBlockResponse
GetLatestBlockResponse is the response type for the Query/GetLatestBlock RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_id | [tendermint.types.BlockID](#tendermint.types.BlockID) |  |  |
| block | [tendermint.types.Block](#tendermint.types.Block) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"></a>

### GetLatestValidatorSetRequest
GetLatestValidatorSetRequest is the request type for the Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"></a>

### GetLatestValidatorSetResponse
GetLatestValidatorSetResponse is the response type for the Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [int64](#int64) |  |  |
| validators | [Validator](#cosmos.base.tendermint.v1beta1.Validator) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"></a>

### GetNodeInfoRequest
GetNodeInfoRequest is the request type for the Query/GetNodeInfo RPC method.






<a name="cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"></a>

### GetNodeInfoResponse
GetNodeInfoResponse is the request type for the Query/GetNodeInfo RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| default_node_info | [tendermint.p2p.DefaultNodeInfo](#tendermint.p2p.DefaultNodeInfo) |  |  |
| application_version | [VersionInfo](#cosmos.base.tendermint.v1beta1.VersionInfo) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetSyncingRequest"></a>

### GetSyncingRequest
GetSyncingRequest is the request type for the Query/GetSyncing RPC method.






<a name="cosmos.base.tendermint.v1beta1.GetSyncingResponse"></a>

### GetSyncingResponse
GetSyncingResponse is the response type for the Query/GetSyncing RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| syncing | [bool](#bool) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"></a>

### GetValidatorSetByHeightRequest
GetValidatorSetByHeightRequest is the request type for the Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [int64](#int64) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"></a>

### GetValidatorSetByHeightResponse
GetValidatorSetByHeightResponse is the response type for the Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [int64](#int64) |  |  |
| validators | [Validator](#cosmos.base.tendermint.v1beta1.Validator) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.base.tendermint.v1beta1.Module"></a>

### Module
Module is the type for VersionInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | module path |
| version | [string](#string) |  | module version |
| sum | [string](#string) |  | checksum |






<a name="cosmos.base.tendermint.v1beta1.Validator"></a>

### Validator
Validator is the type for the validator-set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| pub_key | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| voting_power | [int64](#int64) |  |  |
| proposer_priority | [int64](#int64) |  |  |






<a name="cosmos.base.tendermint.v1beta1.VersionInfo"></a>

### VersionInfo
VersionInfo is the type for the GetNodeInfoResponse message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| app_name | [string](#string) |  |  |
| version | [string](#string) |  |  |
| git_commit | [string](#string) |  |  |
| build_tags | [string](#string) |  |  |
| go_version | [string](#string) |  |  |
| build_deps | [Module](#cosmos.base.tendermint.v1beta1.Module) | repeated |  |





 

 

 


<a name="cosmos.base.tendermint.v1beta1.Service"></a>

### Service
Service defines the gRPC querier service for tendermint queries.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetNodeInfo | [GetNodeInfoRequest](#cosmos.base.tendermint.v1beta1.GetNodeInfoRequest) | [GetNodeInfoResponse](#cosmos.base.tendermint.v1beta1.GetNodeInfoResponse) | GetNodeInfo queries the current node info. |
| GetSyncing | [GetSyncingRequest](#cosmos.base.tendermint.v1beta1.GetSyncingRequest) | [GetSyncingResponse](#cosmos.base.tendermint.v1beta1.GetSyncingResponse) | GetSyncing queries node syncing. |
| GetLatestBlock | [GetLatestBlockRequest](#cosmos.base.tendermint.v1beta1.GetLatestBlockRequest) | [GetLatestBlockResponse](#cosmos.base.tendermint.v1beta1.GetLatestBlockResponse) | GetLatestBlock returns the latest block. |
| GetBlockByHeight | [GetBlockByHeightRequest](#cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest) | [GetBlockByHeightResponse](#cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse) | GetBlockByHeight queries block for given height. |
| GetLatestValidatorSet | [GetLatestValidatorSetRequest](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest) | [GetLatestValidatorSetResponse](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse) | GetLatestValidatorSet queries latest validator-set. |
| GetValidatorSetByHeight | [GetValidatorSetByHeightRequest](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest) | [GetValidatorSetByHeightResponse](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse) | GetValidatorSetByHeight queries validator-set at a given height. |

 



<a name="did/did.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/did.proto



<a name="did.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| KYCvalidated | [bool](#bool) |  |  |






<a name="did.DidCredential"></a>

### DidCredential



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credtype | [string](#string) | repeated |  |
| issuer | [string](#string) |  |  |
| issued | [string](#string) |  |  |
| claim | [Claim](#did.Claim) |  |  |






<a name="did.IxoDid"></a>

### IxoDid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| verifyKey | [string](#string) |  |  |
| encryptionPublicKey | [string](#string) |  |  |
| secret | [Secret](#did.Secret) |  |  |






<a name="did.Secret"></a>

### Secret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seed | [string](#string) |  |  |
| signKey | [string](#string) |  |  |
| encryptionPrivateKey | [string](#string) |  |  |





 

 

 

 



<a name="did/diddoc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/diddoc.proto



<a name="did.BaseDidDoc"></a>

### BaseDidDoc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pubKey | [string](#string) |  |  |
| credentials | [DidCredential](#did.DidCredential) | repeated |  |





 

 

 

 



<a name="did/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/genesis.proto



<a name="did.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated | DidDoc is an interface so we use Any here, like evidence GenesisState |





 

 

 

 



<a name="did/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## did/query.proto



<a name="did.QueryAddressFromBase58EncodedPubkeyRequest"></a>

### QueryAddressFromBase58EncodedPubkeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubKey | [string](#string) |  | pubKey defines the PubKey for the requested address |






<a name="did.QueryAddressFromBase58EncodedPubkeyResponse"></a>

### QueryAddressFromBase58EncodedPubkeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address returns the address for a given PubKey |






<a name="did.QueryAddressFromDidRequest"></a>

### QueryAddressFromDidRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  | did defines the DID for the requested address |






<a name="did.QueryAddressFromDidResponse"></a>

### QueryAddressFromDidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address returns the address for a given DID |






<a name="did.QueryAllDidDocsRequest"></a>

### QueryAllDidDocsRequest
no input needed






<a name="did.QueryAllDidDocsResponse"></a>

### QueryAllDidDocsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddocs | [google.protobuf.Any](#google.protobuf.Any) | repeated | diddocs returns a list of all DidDocs |






<a name="did.QueryAllDidsRequest"></a>

### QueryAllDidsRequest
no input needed






<a name="did.QueryAllDidsResponse"></a>

### QueryAllDidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dids | [string](#string) | repeated | dids returns a list of all DIDs |






<a name="did.QueryDidDocRequest"></a>

### QueryDidDocRequest
Request/response types from old x/did/client/cli/query.go and
x/did/client/rest/query.go


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  | did defines the DID for the requested DidDoc |






<a name="did.QueryDidDocResponse"></a>

### QueryDidDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| diddoc | [google.protobuf.Any](#google.protobuf.Any) |  | diddoc returns the requested DidDoc |





 

 

 


<a name="did.Query"></a>

### Query
To get a list of all module queries, go to your module&#39;s keeper/querier.go
and check all cases in NewQuerier(). REST endpoints taken from previous
did/client/rest/query.go

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
| didCredential | [DidCredential](#did.DidCredential) |  |  |






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

