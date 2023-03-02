# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [ixo/bonds/v1beta1/bonds.proto](#ixo/bonds/v1beta1/bonds.proto)
    - [BaseOrder](#ixo.bonds.v1beta1.BaseOrder)
    - [Batch](#ixo.bonds.v1beta1.Batch)
    - [Bond](#ixo.bonds.v1beta1.Bond)
    - [BondDetails](#ixo.bonds.v1beta1.BondDetails)
    - [BuyOrder](#ixo.bonds.v1beta1.BuyOrder)
    - [FunctionParam](#ixo.bonds.v1beta1.FunctionParam)
    - [Params](#ixo.bonds.v1beta1.Params)
    - [SellOrder](#ixo.bonds.v1beta1.SellOrder)
    - [SwapOrder](#ixo.bonds.v1beta1.SwapOrder)
  
- [ixo/bonds/v1beta1/genesis.proto](#ixo/bonds/v1beta1/genesis.proto)
    - [GenesisState](#ixo.bonds.v1beta1.GenesisState)
  
- [ixo/bonds/v1beta1/query.proto](#ixo/bonds/v1beta1/query.proto)
    - [QueryAlphaMaximumsRequest](#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest)
    - [QueryAlphaMaximumsResponse](#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse)
    - [QueryAvailableReserveRequest](#ixo.bonds.v1beta1.QueryAvailableReserveRequest)
    - [QueryAvailableReserveResponse](#ixo.bonds.v1beta1.QueryAvailableReserveResponse)
    - [QueryBatchRequest](#ixo.bonds.v1beta1.QueryBatchRequest)
    - [QueryBatchResponse](#ixo.bonds.v1beta1.QueryBatchResponse)
    - [QueryBondRequest](#ixo.bonds.v1beta1.QueryBondRequest)
    - [QueryBondResponse](#ixo.bonds.v1beta1.QueryBondResponse)
    - [QueryBondsDetailedRequest](#ixo.bonds.v1beta1.QueryBondsDetailedRequest)
    - [QueryBondsDetailedResponse](#ixo.bonds.v1beta1.QueryBondsDetailedResponse)
    - [QueryBondsRequest](#ixo.bonds.v1beta1.QueryBondsRequest)
    - [QueryBondsResponse](#ixo.bonds.v1beta1.QueryBondsResponse)
    - [QueryBuyPriceRequest](#ixo.bonds.v1beta1.QueryBuyPriceRequest)
    - [QueryBuyPriceResponse](#ixo.bonds.v1beta1.QueryBuyPriceResponse)
    - [QueryCurrentPriceRequest](#ixo.bonds.v1beta1.QueryCurrentPriceRequest)
    - [QueryCurrentPriceResponse](#ixo.bonds.v1beta1.QueryCurrentPriceResponse)
    - [QueryCurrentReserveRequest](#ixo.bonds.v1beta1.QueryCurrentReserveRequest)
    - [QueryCurrentReserveResponse](#ixo.bonds.v1beta1.QueryCurrentReserveResponse)
    - [QueryCustomPriceRequest](#ixo.bonds.v1beta1.QueryCustomPriceRequest)
    - [QueryCustomPriceResponse](#ixo.bonds.v1beta1.QueryCustomPriceResponse)
    - [QueryLastBatchRequest](#ixo.bonds.v1beta1.QueryLastBatchRequest)
    - [QueryLastBatchResponse](#ixo.bonds.v1beta1.QueryLastBatchResponse)
    - [QueryParamsRequest](#ixo.bonds.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.bonds.v1beta1.QueryParamsResponse)
    - [QuerySellReturnRequest](#ixo.bonds.v1beta1.QuerySellReturnRequest)
    - [QuerySellReturnResponse](#ixo.bonds.v1beta1.QuerySellReturnResponse)
    - [QuerySwapReturnRequest](#ixo.bonds.v1beta1.QuerySwapReturnRequest)
    - [QuerySwapReturnResponse](#ixo.bonds.v1beta1.QuerySwapReturnResponse)
  
    - [Query](#ixo.bonds.v1beta1.Query)
  
- [ixo/bonds/v1beta1/tx.proto](#ixo/bonds/v1beta1/tx.proto)
    - [MsgBuy](#ixo.bonds.v1beta1.MsgBuy)
    - [MsgBuyResponse](#ixo.bonds.v1beta1.MsgBuyResponse)
    - [MsgCreateBond](#ixo.bonds.v1beta1.MsgCreateBond)
    - [MsgCreateBondResponse](#ixo.bonds.v1beta1.MsgCreateBondResponse)
    - [MsgEditBond](#ixo.bonds.v1beta1.MsgEditBond)
    - [MsgEditBondResponse](#ixo.bonds.v1beta1.MsgEditBondResponse)
    - [MsgMakeOutcomePayment](#ixo.bonds.v1beta1.MsgMakeOutcomePayment)
    - [MsgMakeOutcomePaymentResponse](#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse)
    - [MsgSell](#ixo.bonds.v1beta1.MsgSell)
    - [MsgSellResponse](#ixo.bonds.v1beta1.MsgSellResponse)
    - [MsgSetNextAlpha](#ixo.bonds.v1beta1.MsgSetNextAlpha)
    - [MsgSetNextAlphaResponse](#ixo.bonds.v1beta1.MsgSetNextAlphaResponse)
    - [MsgSwap](#ixo.bonds.v1beta1.MsgSwap)
    - [MsgSwapResponse](#ixo.bonds.v1beta1.MsgSwapResponse)
    - [MsgUpdateBondState](#ixo.bonds.v1beta1.MsgUpdateBondState)
    - [MsgUpdateBondStateResponse](#ixo.bonds.v1beta1.MsgUpdateBondStateResponse)
    - [MsgWithdrawReserve](#ixo.bonds.v1beta1.MsgWithdrawReserve)
    - [MsgWithdrawReserveResponse](#ixo.bonds.v1beta1.MsgWithdrawReserveResponse)
    - [MsgWithdrawShare](#ixo.bonds.v1beta1.MsgWithdrawShare)
    - [MsgWithdrawShareResponse](#ixo.bonds.v1beta1.MsgWithdrawShareResponse)
  
    - [Msg](#ixo.bonds.v1beta1.Msg)
  
- [ixo/claims/v1beta1/claims.proto](#ixo/claims/v1beta1/claims.proto)
    - [Claim](#ixo.claims.v1beta1.Claim)
    - [ClaimPayments](#ixo.claims.v1beta1.ClaimPayments)
    - [Collection](#ixo.claims.v1beta1.Collection)
    - [Dispute](#ixo.claims.v1beta1.Dispute)
    - [DisputeData](#ixo.claims.v1beta1.DisputeData)
    - [Evaluation](#ixo.claims.v1beta1.Evaluation)
    - [Params](#ixo.claims.v1beta1.Params)
    - [Payment](#ixo.claims.v1beta1.Payment)
    - [Payments](#ixo.claims.v1beta1.Payments)
  
    - [CollectionState](#ixo.claims.v1beta1.CollectionState)
    - [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus)
    - [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus)
    - [PaymentType](#ixo.claims.v1beta1.PaymentType)
  
- [ixo/claims/v1beta1/cosmos.proto](#ixo/claims/v1beta1/cosmos.proto)
    - [Input](#ixo.claims.v1beta1.Input)
    - [Output](#ixo.claims.v1beta1.Output)
  
- [ixo/claims/v1beta1/authz.proto](#ixo/claims/v1beta1/authz.proto)
    - [EvaluateClaimAuthorization](#ixo.claims.v1beta1.EvaluateClaimAuthorization)
    - [EvaluateClaimConstraints](#ixo.claims.v1beta1.EvaluateClaimConstraints)
    - [SubmitClaimAuthorization](#ixo.claims.v1beta1.SubmitClaimAuthorization)
    - [SubmitClaimConstraints](#ixo.claims.v1beta1.SubmitClaimConstraints)
    - [WithdrawPaymentAuthorization](#ixo.claims.v1beta1.WithdrawPaymentAuthorization)
    - [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints)
  
- [ixo/claims/v1beta1/event.proto](#ixo/claims/v1beta1/event.proto)
    - [ClaimDisputedEvent](#ixo.claims.v1beta1.ClaimDisputedEvent)
    - [ClaimEvaluatedEvent](#ixo.claims.v1beta1.ClaimEvaluatedEvent)
    - [ClaimSubmittedEvent](#ixo.claims.v1beta1.ClaimSubmittedEvent)
    - [ClaimUpdatedEvent](#ixo.claims.v1beta1.ClaimUpdatedEvent)
    - [CollectionCreatedEvent](#ixo.claims.v1beta1.CollectionCreatedEvent)
    - [CollectionUpdatedEvent](#ixo.claims.v1beta1.CollectionUpdatedEvent)
    - [PaymentWithdrawCreatedEvent](#ixo.claims.v1beta1.PaymentWithdrawCreatedEvent)
    - [PaymentWithdrawnEvent](#ixo.claims.v1beta1.PaymentWithdrawnEvent)
  
- [ixo/claims/v1beta1/genesis.proto](#ixo/claims/v1beta1/genesis.proto)
    - [GenesisState](#ixo.claims.v1beta1.GenesisState)
  
- [ixo/claims/v1beta1/query.proto](#ixo/claims/v1beta1/query.proto)
    - [QueryClaimListRequest](#ixo.claims.v1beta1.QueryClaimListRequest)
    - [QueryClaimListResponse](#ixo.claims.v1beta1.QueryClaimListResponse)
    - [QueryClaimRequest](#ixo.claims.v1beta1.QueryClaimRequest)
    - [QueryClaimResponse](#ixo.claims.v1beta1.QueryClaimResponse)
    - [QueryCollectionListRequest](#ixo.claims.v1beta1.QueryCollectionListRequest)
    - [QueryCollectionListResponse](#ixo.claims.v1beta1.QueryCollectionListResponse)
    - [QueryCollectionRequest](#ixo.claims.v1beta1.QueryCollectionRequest)
    - [QueryCollectionResponse](#ixo.claims.v1beta1.QueryCollectionResponse)
    - [QueryDisputeListRequest](#ixo.claims.v1beta1.QueryDisputeListRequest)
    - [QueryDisputeListResponse](#ixo.claims.v1beta1.QueryDisputeListResponse)
    - [QueryDisputeRequest](#ixo.claims.v1beta1.QueryDisputeRequest)
    - [QueryDisputeResponse](#ixo.claims.v1beta1.QueryDisputeResponse)
    - [QueryParamsRequest](#ixo.claims.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.claims.v1beta1.QueryParamsResponse)
  
    - [Query](#ixo.claims.v1beta1.Query)
  
- [ixo/claims/v1beta1/tx.proto](#ixo/claims/v1beta1/tx.proto)
    - [MsgCreateCollection](#ixo.claims.v1beta1.MsgCreateCollection)
    - [MsgCreateCollectionResponse](#ixo.claims.v1beta1.MsgCreateCollectionResponse)
    - [MsgDisputeClaim](#ixo.claims.v1beta1.MsgDisputeClaim)
    - [MsgDisputeClaimResponse](#ixo.claims.v1beta1.MsgDisputeClaimResponse)
    - [MsgEvaluateClaim](#ixo.claims.v1beta1.MsgEvaluateClaim)
    - [MsgEvaluateClaimResponse](#ixo.claims.v1beta1.MsgEvaluateClaimResponse)
    - [MsgSubmitClaim](#ixo.claims.v1beta1.MsgSubmitClaim)
    - [MsgSubmitClaimResponse](#ixo.claims.v1beta1.MsgSubmitClaimResponse)
    - [MsgWithdrawPayment](#ixo.claims.v1beta1.MsgWithdrawPayment)
    - [MsgWithdrawPaymentResponse](#ixo.claims.v1beta1.MsgWithdrawPaymentResponse)
  
    - [Msg](#ixo.claims.v1beta1.Msg)
  
- [ixo/iid/v1beta1/types.proto](#ixo/iid/v1beta1/types.proto)
    - [AccordedRight](#ixo.iid.v1beta1.AccordedRight)
    - [Context](#ixo.iid.v1beta1.Context)
    - [IidMetadata](#ixo.iid.v1beta1.IidMetadata)
    - [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim)
    - [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity)
    - [LinkedResource](#ixo.iid.v1beta1.LinkedResource)
    - [Service](#ixo.iid.v1beta1.Service)
    - [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod)
  
- [ixo/iid/v1beta1/iid.proto](#ixo/iid/v1beta1/iid.proto)
    - [IidDocument](#ixo.iid.v1beta1.IidDocument)
  
- [ixo/entity/v1beta1/entity.proto](#ixo/entity/v1beta1/entity.proto)
    - [Entity](#ixo.entity.v1beta1.Entity)
    - [EntityMetadata](#ixo.entity.v1beta1.EntityMetadata)
    - [Params](#ixo.entity.v1beta1.Params)
  
- [ixo/entity/v1beta1/event.proto](#ixo/entity/v1beta1/event.proto)
    - [EntityCreatedEvent](#ixo.entity.v1beta1.EntityCreatedEvent)
    - [EntityTransferredEvent](#ixo.entity.v1beta1.EntityTransferredEvent)
    - [EntityUpdatedEvent](#ixo.entity.v1beta1.EntityUpdatedEvent)
    - [EntityVerifiedUpdatedEvent](#ixo.entity.v1beta1.EntityVerifiedUpdatedEvent)
  
- [ixo/entity/v1beta1/genesis.proto](#ixo/entity/v1beta1/genesis.proto)
    - [GenesisState](#ixo.entity.v1beta1.GenesisState)
  
- [ixo/entity/v1beta1/proposal.proto](#ixo/entity/v1beta1/proposal.proto)
    - [InitializeNftContract](#ixo.entity.v1beta1.InitializeNftContract)
  
- [ixo/entity/v1beta1/query.proto](#ixo/entity/v1beta1/query.proto)
    - [QueryEntityIidDocumentRequest](#ixo.entity.v1beta1.QueryEntityIidDocumentRequest)
    - [QueryEntityIidDocumentResponse](#ixo.entity.v1beta1.QueryEntityIidDocumentResponse)
    - [QueryEntityListRequest](#ixo.entity.v1beta1.QueryEntityListRequest)
    - [QueryEntityListResponse](#ixo.entity.v1beta1.QueryEntityListResponse)
    - [QueryEntityMetadataRequest](#ixo.entity.v1beta1.QueryEntityMetadataRequest)
    - [QueryEntityMetadataResponse](#ixo.entity.v1beta1.QueryEntityMetadataResponse)
    - [QueryEntityRequest](#ixo.entity.v1beta1.QueryEntityRequest)
    - [QueryEntityResponse](#ixo.entity.v1beta1.QueryEntityResponse)
    - [QueryEntityVerifiedRequest](#ixo.entity.v1beta1.QueryEntityVerifiedRequest)
    - [QueryEntityVerifiedResponse](#ixo.entity.v1beta1.QueryEntityVerifiedResponse)
  
    - [Query](#ixo.entity.v1beta1.Query)
  
- [ixo/iid/v1beta1/tx.proto](#ixo/iid/v1beta1/tx.proto)
    - [MsgAddAccordedRight](#ixo.iid.v1beta1.MsgAddAccordedRight)
    - [MsgAddAccordedRightResponse](#ixo.iid.v1beta1.MsgAddAccordedRightResponse)
    - [MsgAddController](#ixo.iid.v1beta1.MsgAddController)
    - [MsgAddControllerResponse](#ixo.iid.v1beta1.MsgAddControllerResponse)
    - [MsgAddIidContext](#ixo.iid.v1beta1.MsgAddIidContext)
    - [MsgAddIidContextResponse](#ixo.iid.v1beta1.MsgAddIidContextResponse)
    - [MsgAddLinkedClaim](#ixo.iid.v1beta1.MsgAddLinkedClaim)
    - [MsgAddLinkedClaimResponse](#ixo.iid.v1beta1.MsgAddLinkedClaimResponse)
    - [MsgAddLinkedEntity](#ixo.iid.v1beta1.MsgAddLinkedEntity)
    - [MsgAddLinkedEntityResponse](#ixo.iid.v1beta1.MsgAddLinkedEntityResponse)
    - [MsgAddLinkedResource](#ixo.iid.v1beta1.MsgAddLinkedResource)
    - [MsgAddLinkedResourceResponse](#ixo.iid.v1beta1.MsgAddLinkedResourceResponse)
    - [MsgAddService](#ixo.iid.v1beta1.MsgAddService)
    - [MsgAddServiceResponse](#ixo.iid.v1beta1.MsgAddServiceResponse)
    - [MsgAddVerification](#ixo.iid.v1beta1.MsgAddVerification)
    - [MsgAddVerificationResponse](#ixo.iid.v1beta1.MsgAddVerificationResponse)
    - [MsgCreateIidDocument](#ixo.iid.v1beta1.MsgCreateIidDocument)
    - [MsgCreateIidDocumentResponse](#ixo.iid.v1beta1.MsgCreateIidDocumentResponse)
    - [MsgDeactivateIID](#ixo.iid.v1beta1.MsgDeactivateIID)
    - [MsgDeactivateIIDResponse](#ixo.iid.v1beta1.MsgDeactivateIIDResponse)
    - [MsgDeleteAccordedRight](#ixo.iid.v1beta1.MsgDeleteAccordedRight)
    - [MsgDeleteAccordedRightResponse](#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse)
    - [MsgDeleteController](#ixo.iid.v1beta1.MsgDeleteController)
    - [MsgDeleteControllerResponse](#ixo.iid.v1beta1.MsgDeleteControllerResponse)
    - [MsgDeleteIidContext](#ixo.iid.v1beta1.MsgDeleteIidContext)
    - [MsgDeleteIidContextResponse](#ixo.iid.v1beta1.MsgDeleteIidContextResponse)
    - [MsgDeleteLinkedClaim](#ixo.iid.v1beta1.MsgDeleteLinkedClaim)
    - [MsgDeleteLinkedClaimResponse](#ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse)
    - [MsgDeleteLinkedEntity](#ixo.iid.v1beta1.MsgDeleteLinkedEntity)
    - [MsgDeleteLinkedEntityResponse](#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse)
    - [MsgDeleteLinkedResource](#ixo.iid.v1beta1.MsgDeleteLinkedResource)
    - [MsgDeleteLinkedResourceResponse](#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse)
    - [MsgDeleteService](#ixo.iid.v1beta1.MsgDeleteService)
    - [MsgDeleteServiceResponse](#ixo.iid.v1beta1.MsgDeleteServiceResponse)
    - [MsgRevokeVerification](#ixo.iid.v1beta1.MsgRevokeVerification)
    - [MsgRevokeVerificationResponse](#ixo.iid.v1beta1.MsgRevokeVerificationResponse)
    - [MsgSetVerificationRelationships](#ixo.iid.v1beta1.MsgSetVerificationRelationships)
    - [MsgSetVerificationRelationshipsResponse](#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse)
    - [MsgUpdateIidDocument](#ixo.iid.v1beta1.MsgUpdateIidDocument)
    - [MsgUpdateIidDocumentResponse](#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse)
    - [Verification](#ixo.iid.v1beta1.Verification)
  
    - [Msg](#ixo.iid.v1beta1.Msg)
  
- [ixo/entity/v1beta1/tx.proto](#ixo/entity/v1beta1/tx.proto)
    - [MsgCreateEntity](#ixo.entity.v1beta1.MsgCreateEntity)
    - [MsgCreateEntityResponse](#ixo.entity.v1beta1.MsgCreateEntityResponse)
    - [MsgTransferEntity](#ixo.entity.v1beta1.MsgTransferEntity)
    - [MsgTransferEntityResponse](#ixo.entity.v1beta1.MsgTransferEntityResponse)
    - [MsgUpdateEntity](#ixo.entity.v1beta1.MsgUpdateEntity)
    - [MsgUpdateEntityResponse](#ixo.entity.v1beta1.MsgUpdateEntityResponse)
    - [MsgUpdateEntityVerified](#ixo.entity.v1beta1.MsgUpdateEntityVerified)
    - [MsgUpdateEntityVerifiedResponse](#ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse)
  
    - [Msg](#ixo.entity.v1beta1.Msg)
  
- [ixo/iid/v1beta1/event.proto](#ixo/iid/v1beta1/event.proto)
    - [IidDocumentCreatedEvent](#ixo.iid.v1beta1.IidDocumentCreatedEvent)
    - [IidDocumentUpdatedEvent](#ixo.iid.v1beta1.IidDocumentUpdatedEvent)
  
- [ixo/iid/v1beta1/genesis.proto](#ixo/iid/v1beta1/genesis.proto)
    - [GenesisState](#ixo.iid.v1beta1.GenesisState)
  
- [ixo/iid/v1beta1/query.proto](#ixo/iid/v1beta1/query.proto)
    - [QueryIidDocumentRequest](#ixo.iid.v1beta1.QueryIidDocumentRequest)
    - [QueryIidDocumentResponse](#ixo.iid.v1beta1.QueryIidDocumentResponse)
    - [QueryIidDocumentsRequest](#ixo.iid.v1beta1.QueryIidDocumentsRequest)
    - [QueryIidDocumentsResponse](#ixo.iid.v1beta1.QueryIidDocumentsResponse)
  
    - [Query](#ixo.iid.v1beta1.Query)
  
- [ixo/legacy/did/did.proto](#ixo/legacy/did/did.proto)
    - [Claim](#legacydid.Claim)
    - [DidCredential](#legacydid.DidCredential)
    - [IxoDid](#legacydid.IxoDid)
    - [Secret](#legacydid.Secret)
  
- [ixo/legacy/did/diddoc.proto](#ixo/legacy/did/diddoc.proto)
    - [BaseDidDoc](#legacydid.BaseDidDoc)
  
- [ixo/payments/v1/payments.proto](#ixo/payments/v1/payments.proto)
    - [BlockPeriod](#ixo.payments.v1.BlockPeriod)
    - [Discount](#ixo.payments.v1.Discount)
    - [DistributionShare](#ixo.payments.v1.DistributionShare)
    - [PaymentContract](#ixo.payments.v1.PaymentContract)
    - [PaymentTemplate](#ixo.payments.v1.PaymentTemplate)
    - [Subscription](#ixo.payments.v1.Subscription)
    - [TestPeriod](#ixo.payments.v1.TestPeriod)
    - [TimePeriod](#ixo.payments.v1.TimePeriod)
  
- [ixo/payments/v1/genesis.proto](#ixo/payments/v1/genesis.proto)
    - [GenesisState](#ixo.payments.v1.GenesisState)
  
- [ixo/payments/v1/query.proto](#ixo/payments/v1/query.proto)
    - [QueryPaymentContractRequest](#ixo.payments.v1.QueryPaymentContractRequest)
    - [QueryPaymentContractResponse](#ixo.payments.v1.QueryPaymentContractResponse)
    - [QueryPaymentContractsByIdPrefixRequest](#ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest)
    - [QueryPaymentContractsByIdPrefixResponse](#ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse)
    - [QueryPaymentTemplateRequest](#ixo.payments.v1.QueryPaymentTemplateRequest)
    - [QueryPaymentTemplateResponse](#ixo.payments.v1.QueryPaymentTemplateResponse)
    - [QuerySubscriptionRequest](#ixo.payments.v1.QuerySubscriptionRequest)
    - [QuerySubscriptionResponse](#ixo.payments.v1.QuerySubscriptionResponse)
  
    - [Query](#ixo.payments.v1.Query)
  
- [ixo/payments/v1/tx.proto](#ixo/payments/v1/tx.proto)
    - [MsgCreatePaymentContract](#ixo.payments.v1.MsgCreatePaymentContract)
    - [MsgCreatePaymentContractResponse](#ixo.payments.v1.MsgCreatePaymentContractResponse)
    - [MsgCreatePaymentTemplate](#ixo.payments.v1.MsgCreatePaymentTemplate)
    - [MsgCreatePaymentTemplateResponse](#ixo.payments.v1.MsgCreatePaymentTemplateResponse)
    - [MsgCreateSubscription](#ixo.payments.v1.MsgCreateSubscription)
    - [MsgCreateSubscriptionResponse](#ixo.payments.v1.MsgCreateSubscriptionResponse)
    - [MsgEffectPayment](#ixo.payments.v1.MsgEffectPayment)
    - [MsgEffectPaymentResponse](#ixo.payments.v1.MsgEffectPaymentResponse)
    - [MsgGrantDiscount](#ixo.payments.v1.MsgGrantDiscount)
    - [MsgGrantDiscountResponse](#ixo.payments.v1.MsgGrantDiscountResponse)
    - [MsgRevokeDiscount](#ixo.payments.v1.MsgRevokeDiscount)
    - [MsgRevokeDiscountResponse](#ixo.payments.v1.MsgRevokeDiscountResponse)
    - [MsgSetPaymentContractAuthorisation](#ixo.payments.v1.MsgSetPaymentContractAuthorisation)
    - [MsgSetPaymentContractAuthorisationResponse](#ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse)
  
    - [Msg](#ixo.payments.v1.Msg)
  
- [ixo/project/v1/project.proto](#ixo/project/v1/project.proto)
    - [AccountMap](#ixo.project.v1.AccountMap)
    - [AccountMap.MapEntry](#ixo.project.v1.AccountMap.MapEntry)
    - [Claim](#ixo.project.v1.Claim)
    - [Claims](#ixo.project.v1.Claims)
    - [CreateAgentDoc](#ixo.project.v1.CreateAgentDoc)
    - [CreateClaimDoc](#ixo.project.v1.CreateClaimDoc)
    - [CreateEvaluationDoc](#ixo.project.v1.CreateEvaluationDoc)
    - [GenesisAccountMap](#ixo.project.v1.GenesisAccountMap)
    - [GenesisAccountMap.MapEntry](#ixo.project.v1.GenesisAccountMap.MapEntry)
    - [Params](#ixo.project.v1.Params)
    - [ProjectDoc](#ixo.project.v1.ProjectDoc)
    - [UpdateAgentDoc](#ixo.project.v1.UpdateAgentDoc)
    - [UpdateProjectStatusDoc](#ixo.project.v1.UpdateProjectStatusDoc)
    - [WithdrawFundsDoc](#ixo.project.v1.WithdrawFundsDoc)
    - [WithdrawalInfoDoc](#ixo.project.v1.WithdrawalInfoDoc)
    - [WithdrawalInfoDocs](#ixo.project.v1.WithdrawalInfoDocs)
  
- [ixo/project/v1/genesis.proto](#ixo/project/v1/genesis.proto)
    - [GenesisState](#ixo.project.v1.GenesisState)
  
- [ixo/project/v1/query.proto](#ixo/project/v1/query.proto)
    - [QueryParamsRequest](#ixo.project.v1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.project.v1.QueryParamsResponse)
    - [QueryProjectAccountsRequest](#ixo.project.v1.QueryProjectAccountsRequest)
    - [QueryProjectAccountsResponse](#ixo.project.v1.QueryProjectAccountsResponse)
    - [QueryProjectDocRequest](#ixo.project.v1.QueryProjectDocRequest)
    - [QueryProjectDocResponse](#ixo.project.v1.QueryProjectDocResponse)
    - [QueryProjectTxRequest](#ixo.project.v1.QueryProjectTxRequest)
    - [QueryProjectTxResponse](#ixo.project.v1.QueryProjectTxResponse)
  
    - [Query](#ixo.project.v1.Query)
  
- [ixo/project/v1/tx.proto](#ixo/project/v1/tx.proto)
    - [MsgCreateAgent](#ixo.project.v1.MsgCreateAgent)
    - [MsgCreateAgentResponse](#ixo.project.v1.MsgCreateAgentResponse)
    - [MsgCreateClaim](#ixo.project.v1.MsgCreateClaim)
    - [MsgCreateClaimResponse](#ixo.project.v1.MsgCreateClaimResponse)
    - [MsgCreateEvaluation](#ixo.project.v1.MsgCreateEvaluation)
    - [MsgCreateEvaluationResponse](#ixo.project.v1.MsgCreateEvaluationResponse)
    - [MsgCreateProject](#ixo.project.v1.MsgCreateProject)
    - [MsgCreateProjectResponse](#ixo.project.v1.MsgCreateProjectResponse)
    - [MsgUpdateAgent](#ixo.project.v1.MsgUpdateAgent)
    - [MsgUpdateAgentResponse](#ixo.project.v1.MsgUpdateAgentResponse)
    - [MsgUpdateProjectDoc](#ixo.project.v1.MsgUpdateProjectDoc)
    - [MsgUpdateProjectDocResponse](#ixo.project.v1.MsgUpdateProjectDocResponse)
    - [MsgUpdateProjectStatus](#ixo.project.v1.MsgUpdateProjectStatus)
    - [MsgUpdateProjectStatusResponse](#ixo.project.v1.MsgUpdateProjectStatusResponse)
    - [MsgWithdrawFunds](#ixo.project.v1.MsgWithdrawFunds)
    - [MsgWithdrawFundsResponse](#ixo.project.v1.MsgWithdrawFundsResponse)
  
    - [Msg](#ixo.project.v1.Msg)
  
- [ixo/token/v1beta1/token.proto](#ixo/token/v1beta1/token.proto)
    - [Params](#ixo.token.v1beta1.Params)
    - [Token](#ixo.token.v1beta1.Token)
    - [TokenData](#ixo.token.v1beta1.TokenData)
    - [TokenProperties](#ixo.token.v1beta1.TokenProperties)
    - [TokensCancelled](#ixo.token.v1beta1.TokensCancelled)
    - [TokensRetired](#ixo.token.v1beta1.TokensRetired)
  
- [ixo/token/v1beta1/authz.proto](#ixo/token/v1beta1/authz.proto)
    - [MintAuthorization](#ixo.token.v1beta1.MintAuthorization)
    - [MintConstraints](#ixo.token.v1beta1.MintConstraints)
  
- [ixo/token/v1beta1/tx.proto](#ixo/token/v1beta1/tx.proto)
    - [MintBatch](#ixo.token.v1beta1.MintBatch)
    - [MsgCancelToken](#ixo.token.v1beta1.MsgCancelToken)
    - [MsgCancelTokenResponse](#ixo.token.v1beta1.MsgCancelTokenResponse)
    - [MsgCreateToken](#ixo.token.v1beta1.MsgCreateToken)
    - [MsgCreateTokenResponse](#ixo.token.v1beta1.MsgCreateTokenResponse)
    - [MsgMintToken](#ixo.token.v1beta1.MsgMintToken)
    - [MsgMintTokenResponse](#ixo.token.v1beta1.MsgMintTokenResponse)
    - [MsgPauseToken](#ixo.token.v1beta1.MsgPauseToken)
    - [MsgPauseTokenResponse](#ixo.token.v1beta1.MsgPauseTokenResponse)
    - [MsgRetireToken](#ixo.token.v1beta1.MsgRetireToken)
    - [MsgRetireTokenResponse](#ixo.token.v1beta1.MsgRetireTokenResponse)
    - [MsgStopToken](#ixo.token.v1beta1.MsgStopToken)
    - [MsgStopTokenResponse](#ixo.token.v1beta1.MsgStopTokenResponse)
    - [MsgTransferToken](#ixo.token.v1beta1.MsgTransferToken)
    - [MsgTransferTokenResponse](#ixo.token.v1beta1.MsgTransferTokenResponse)
    - [TokenBatch](#ixo.token.v1beta1.TokenBatch)
  
    - [Msg](#ixo.token.v1beta1.Msg)
  
- [ixo/token/v1beta1/event.proto](#ixo/token/v1beta1/event.proto)
    - [TokenCancelledEvent](#ixo.token.v1beta1.TokenCancelledEvent)
    - [TokenCreatedEvent](#ixo.token.v1beta1.TokenCreatedEvent)
    - [TokenMintedEvent](#ixo.token.v1beta1.TokenMintedEvent)
    - [TokenPausedEvent](#ixo.token.v1beta1.TokenPausedEvent)
    - [TokenRetiredEvent](#ixo.token.v1beta1.TokenRetiredEvent)
    - [TokenStoppedEvent](#ixo.token.v1beta1.TokenStoppedEvent)
    - [TokenTransferredEvent](#ixo.token.v1beta1.TokenTransferredEvent)
    - [TokenUpdatedEvent](#ixo.token.v1beta1.TokenUpdatedEvent)
  
- [ixo/token/v1beta1/genesis.proto](#ixo/token/v1beta1/genesis.proto)
    - [GenesisState](#ixo.token.v1beta1.GenesisState)
  
- [ixo/token/v1beta1/proposal.proto](#ixo/token/v1beta1/proposal.proto)
    - [SetTokenContractCodes](#ixo.token.v1beta1.SetTokenContractCodes)
  
- [ixo/token/v1beta1/query.proto](#ixo/token/v1beta1/query.proto)
    - [QueryTokenDocRequest](#ixo.token.v1beta1.QueryTokenDocRequest)
    - [QueryTokenDocResponse](#ixo.token.v1beta1.QueryTokenDocResponse)
    - [QueryTokenListRequest](#ixo.token.v1beta1.QueryTokenListRequest)
    - [QueryTokenListResponse](#ixo.token.v1beta1.QueryTokenListResponse)
    - [QueryTokenMetadataRequest](#ixo.token.v1beta1.QueryTokenMetadataRequest)
    - [QueryTokenMetadataResponse](#ixo.token.v1beta1.QueryTokenMetadataResponse)
    - [TokenMetadataProperties](#ixo.token.v1beta1.TokenMetadataProperties)
  
    - [Query](#ixo.token.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="ixo/bonds/v1beta1/bonds.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/bonds.proto



<a name="ixo.bonds.v1beta1.BaseOrder"></a>

### BaseOrder
BaseOrder defines a base order type. It contains all the necessary fields for
specifying the general details about a buy, sell, or swap order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| cancelled | [bool](#bool) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.Batch"></a>

### Batch
Batch holds a collection of outstanding buy, sell, and swap orders on a
particular bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| blocks_remaining | [string](#string) |  |  |
| next_public_alpha | [string](#string) |  |  |
| total_buy_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| total_sell_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| buy_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| sell_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| buys | [BuyOrder](#ixo.bonds.v1beta1.BuyOrder) | repeated |  |
| sells | [SellOrder](#ixo.bonds.v1beta1.SellOrder) | repeated |  |
| swaps | [SwapOrder](#ixo.bonds.v1beta1.SwapOrder) | repeated |  |






<a name="ixo.bonds.v1beta1.Bond"></a>

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
| function_parameters | [FunctionParam](#ixo.bonds.v1beta1.FunctionParam) | repeated |  |
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
| oracle_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondDetails"></a>

### BondDetails
BondDetails contains details about the current state of a given bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| spot_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.BuyOrder"></a>

### BuyOrder
BuyOrder defines a type for submitting a buy order on a bond, together with
the maximum amount of reserve tokens the buyer is willing to pay.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.FunctionParam"></a>

### FunctionParam
FunctionParam is a key-value pair used for specifying a specific bond
parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| param | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.Params"></a>

### Params
Params defines the parameters for the bonds module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reserved_bond_tokens | [string](#string) | repeated |  |






<a name="ixo.bonds.v1beta1.SellOrder"></a>

### SellOrder
SellOrder defines a type for submitting a sell order on a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |






<a name="ixo.bonds.v1beta1.SwapOrder"></a>

### SwapOrder
SwapOrder defines a type for submitting a swap order between two tokens on a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |
| to_token | [string](#string) |  |  |





 

 

 

 



<a name="ixo/bonds/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/genesis.proto



<a name="ixo.bonds.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the bonds module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [Bond](#ixo.bonds.v1beta1.Bond) | repeated |  |
| batches | [Batch](#ixo.bonds.v1beta1.Batch) | repeated |  |
| params | [Params](#ixo.bonds.v1beta1.Params) |  |  |





 

 

 

 



<a name="ixo/bonds/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/query.proto



<a name="ixo.bonds.v1beta1.QueryAlphaMaximumsRequest"></a>

### QueryAlphaMaximumsRequest
QueryAlphaMaximumsRequest is the request type for the Query/AlphaMaximums RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAlphaMaximumsResponse"></a>

### QueryAlphaMaximumsResponse
QueryAlphaMaximumsResponse is the response type for the Query/AlphaMaximums
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_system_alpha_increase | [string](#string) |  |  |
| max_system_alpha | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAvailableReserveRequest"></a>

### QueryAvailableReserveRequest
QueryAvailableReserveRequest is the request type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAvailableReserveResponse"></a>

### QueryAvailableReserveResponse
QueryAvailableReserveResponse is the response type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBatchRequest"></a>

### QueryBatchRequest
QueryBatchRequest is the request type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBatchResponse"></a>

### QueryBatchResponse
QueryBatchResponse is the response type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#ixo.bonds.v1beta1.Batch) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondRequest"></a>

### QueryBondRequest
QueryBondRequest is the request type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondResponse"></a>

### QueryBondResponse
QueryBondResponse is the response type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#ixo.bonds.v1beta1.Bond) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondsDetailedRequest"></a>

### QueryBondsDetailedRequest
QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC
method.






<a name="ixo.bonds.v1beta1.QueryBondsDetailedResponse"></a>

### QueryBondsDetailedResponse
QueryBondsDetailedResponse is the response type for the Query/BondsDetailed
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds_detailed | [BondDetails](#ixo.bonds.v1beta1.BondDetails) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBondsRequest"></a>

### QueryBondsRequest
QueryBondsRequest is the request type for the Query/Bonds RPC method.






<a name="ixo.bonds.v1beta1.QueryBondsResponse"></a>

### QueryBondsResponse
QueryBondsResponse is the response type for the Query/Bonds RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [string](#string) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBuyPriceRequest"></a>

### QueryBuyPriceRequest
QueryCustomPriceRequest is the request type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBuyPriceResponse"></a>

### QueryBuyPriceResponse
QueryCustomPriceResponse is the response type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest
QueryCurrentPriceRequest is the request type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse
QueryCurrentPriceResponse is the response type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest
QueryCurrentReserveRequest is the request type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse
QueryCurrentReserveResponse is the response type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCustomPriceRequest"></a>

### QueryCustomPriceRequest
QueryCustomPriceRequest is the request type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCustomPriceResponse"></a>

### QueryCustomPriceResponse
QueryCustomPriceResponse is the response type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryLastBatchRequest"></a>

### QueryLastBatchRequest
QueryLastBatchRequest is the request type for the Query/LastBatch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryLastBatchResponse"></a>

### QueryLastBatchResponse
QueryLastBatchResponse is the response type for the Query/LastBatch RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_batch | [Batch](#ixo.bonds.v1beta1.Batch) |  |  |






<a name="ixo.bonds.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="ixo.bonds.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.bonds.v1beta1.Params) |  |  |






<a name="ixo.bonds.v1beta1.QuerySellReturnRequest"></a>

### QuerySellReturnRequest
QuerySellReturnRequest is the request type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QuerySellReturnResponse"></a>

### QuerySellReturnResponse
QuerySellReturnResponse is the response type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| exit_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QuerySwapReturnRequest"></a>

### QuerySwapReturnRequest
QuerySwapReturnRequest is the request type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| from_token_with_amount | [string](#string) |  |  |
| to_token | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QuerySwapReturnResponse"></a>

### QuerySwapReturnResponse
QuerySwapReturnResponse is the response type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 


<a name="ixo.bonds.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bonds | [QueryBondsRequest](#ixo.bonds.v1beta1.QueryBondsRequest) | [QueryBondsResponse](#ixo.bonds.v1beta1.QueryBondsResponse) | Bonds returns all existing bonds. |
| BondsDetailed | [QueryBondsDetailedRequest](#ixo.bonds.v1beta1.QueryBondsDetailedRequest) | [QueryBondsDetailedResponse](#ixo.bonds.v1beta1.QueryBondsDetailedResponse) | BondsDetailed returns a list of all existing bonds with some details about their current state. |
| Params | [QueryParamsRequest](#ixo.bonds.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.bonds.v1beta1.QueryParamsResponse) | Params queries the paramaters of x/bonds module. |
| Bond | [QueryBondRequest](#ixo.bonds.v1beta1.QueryBondRequest) | [QueryBondResponse](#ixo.bonds.v1beta1.QueryBondResponse) | Bond queries info of a specific bond. |
| Batch | [QueryBatchRequest](#ixo.bonds.v1beta1.QueryBatchRequest) | [QueryBatchResponse](#ixo.bonds.v1beta1.QueryBatchResponse) | Batch queries info of a specific bond&#39;s current batch. |
| LastBatch | [QueryLastBatchRequest](#ixo.bonds.v1beta1.QueryLastBatchRequest) | [QueryLastBatchResponse](#ixo.bonds.v1beta1.QueryLastBatchResponse) | LastBatch queries info of a specific bond&#39;s last batch. |
| CurrentPrice | [QueryCurrentPriceRequest](#ixo.bonds.v1beta1.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#ixo.bonds.v1beta1.QueryCurrentPriceResponse) | CurrentPrice queries the current price/s of a specific bond. |
| CurrentReserve | [QueryCurrentReserveRequest](#ixo.bonds.v1beta1.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#ixo.bonds.v1beta1.QueryCurrentReserveResponse) | CurrentReserve queries the current balance/s of the reserve pool for a specific bond. |
| AvailableReserve | [QueryAvailableReserveRequest](#ixo.bonds.v1beta1.QueryAvailableReserveRequest) | [QueryAvailableReserveResponse](#ixo.bonds.v1beta1.QueryAvailableReserveResponse) | AvailableReserve queries current available balance/s of the reserve pool for a specific bond. |
| CustomPrice | [QueryCustomPriceRequest](#ixo.bonds.v1beta1.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#ixo.bonds.v1beta1.QueryCustomPriceResponse) | CustomPrice queries price/s of a specific bond at a specific supply. |
| BuyPrice | [QueryBuyPriceRequest](#ixo.bonds.v1beta1.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#ixo.bonds.v1beta1.QueryBuyPriceResponse) | BuyPrice queries price/s of buying an amount of tokens from a specific bond. |
| SellReturn | [QuerySellReturnRequest](#ixo.bonds.v1beta1.QuerySellReturnRequest) | [QuerySellReturnResponse](#ixo.bonds.v1beta1.QuerySellReturnResponse) | SellReturn queries return/s on selling an amount of tokens of a specific bond. |
| SwapReturn | [QuerySwapReturnRequest](#ixo.bonds.v1beta1.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#ixo.bonds.v1beta1.QuerySwapReturnResponse) | SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond. |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse) | AlphaMaximums queries alpha maximums for a specific augmented bonding curve. |

 



<a name="ixo/bonds/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/tx.proto



<a name="ixo.bonds.v1beta1.MsgBuy"></a>

### MsgBuy
MsgBuy defines a message for buying from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| buyer_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse defines the Msg/Buy response type.






<a name="ixo.bonds.v1beta1.MsgCreateBond"></a>

### MsgCreateBond
MsgCreateBond defines a message for creating a new bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#ixo.bonds.v1beta1.FunctionParam) | repeated |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
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
| creator_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgCreateBondResponse"></a>

### MsgCreateBondResponse
MsgCreateBondResponse defines the Msg/CreateBond response type.






<a name="ixo.bonds.v1beta1.MsgEditBond"></a>

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
| editor_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgEditBondResponse"></a>

### MsgEditBondResponse
MsgEditBondResponse defines the Msg/EditBond response type.






<a name="ixo.bonds.v1beta1.MsgMakeOutcomePayment"></a>

### MsgMakeOutcomePayment
MsgMakeOutcomePayment defines a message for making an outcome payment to a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse"></a>

### MsgMakeOutcomePaymentResponse
MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response
type.






<a name="ixo.bonds.v1beta1.MsgSell"></a>

### MsgSell
MsgSell defines a message for selling from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seller_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| bond_did | [string](#string) |  |  |
| seller_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse defines the Msg/Sell response type.






<a name="ixo.bonds.v1beta1.MsgSetNextAlpha"></a>

### MsgSetNextAlpha
MsgSetNextAlpha defines a message for editing a bond&#39;s alpha parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| alpha | [string](#string) |  |  |
| delta | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
| oracle_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSetNextAlphaResponse"></a>

### MsgSetNextAlphaResponse







<a name="ixo.bonds.v1beta1.MsgSwap"></a>

### MsgSwap
MsgSwap defines a message for swapping from one reserve bond token to
another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swapper_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| from | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| to_token | [string](#string) |  |  |
| swapper_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSwapResponse"></a>

### MsgSwapResponse
MsgSwapResponse defines the Msg/Swap response type.






<a name="ixo.bonds.v1beta1.MsgUpdateBondState"></a>

### MsgUpdateBondState
MsgUpdateBondState defines a message for updating a bond&#39;s current state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| state | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |
| editor_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgUpdateBondStateResponse"></a>

### MsgUpdateBondStateResponse
MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.






<a name="ixo.bonds.v1beta1.MsgWithdrawReserve"></a>

### MsgWithdrawReserve
MsgWithdrawReserve defines a message for withdrawing reserve from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdrawer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| withdrawer_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgWithdrawReserveResponse"></a>

### MsgWithdrawReserveResponse
MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.






<a name="ixo.bonds.v1beta1.MsgWithdrawShare"></a>

### MsgWithdrawShare
MsgWithdrawShare defines a message for withdrawing a share from a bond that
is in the SETTLE stage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| recipient_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgWithdrawShareResponse"></a>

### MsgWithdrawShareResponse
MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.





 

 

 


<a name="ixo.bonds.v1beta1.Msg"></a>

### Msg
Msg defines the bonds Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBond | [MsgCreateBond](#ixo.bonds.v1beta1.MsgCreateBond) | [MsgCreateBondResponse](#ixo.bonds.v1beta1.MsgCreateBondResponse) | CreateBond defines a method for creating a bond. |
| EditBond | [MsgEditBond](#ixo.bonds.v1beta1.MsgEditBond) | [MsgEditBondResponse](#ixo.bonds.v1beta1.MsgEditBondResponse) | EditBond defines a method for editing a bond. |
| SetNextAlpha | [MsgSetNextAlpha](#ixo.bonds.v1beta1.MsgSetNextAlpha) | [MsgSetNextAlphaResponse](#ixo.bonds.v1beta1.MsgSetNextAlphaResponse) | SetNextAlpha defines a method for editing a bond&#39;s alpha parameter. |
| UpdateBondState | [MsgUpdateBondState](#ixo.bonds.v1beta1.MsgUpdateBondState) | [MsgUpdateBondStateResponse](#ixo.bonds.v1beta1.MsgUpdateBondStateResponse) | UpdateBondState defines a method for updating a bond&#39;s current state. |
| Buy | [MsgBuy](#ixo.bonds.v1beta1.MsgBuy) | [MsgBuyResponse](#ixo.bonds.v1beta1.MsgBuyResponse) | Buy defines a method for buying from a bond. |
| Sell | [MsgSell](#ixo.bonds.v1beta1.MsgSell) | [MsgSellResponse](#ixo.bonds.v1beta1.MsgSellResponse) | Sell defines a method for selling from a bond. |
| Swap | [MsgSwap](#ixo.bonds.v1beta1.MsgSwap) | [MsgSwapResponse](#ixo.bonds.v1beta1.MsgSwapResponse) | Swap defines a method for swapping from one reserve bond token to another. |
| MakeOutcomePayment | [MsgMakeOutcomePayment](#ixo.bonds.v1beta1.MsgMakeOutcomePayment) | [MsgMakeOutcomePaymentResponse](#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse) | MakeOutcomePayment defines a method for making an outcome payment to a bond. |
| WithdrawShare | [MsgWithdrawShare](#ixo.bonds.v1beta1.MsgWithdrawShare) | [MsgWithdrawShareResponse](#ixo.bonds.v1beta1.MsgWithdrawShareResponse) | WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage. |
| WithdrawReserve | [MsgWithdrawReserve](#ixo.bonds.v1beta1.MsgWithdrawReserve) | [MsgWithdrawReserveResponse](#ixo.bonds.v1beta1.MsgWithdrawReserveResponse) | WithdrawReserve defines a method for withdrawing reserve from a bond. |

 



<a name="ixo/claims/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/claims.proto



<a name="ixo.claims.v1beta1.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| submission_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submissionDate is the date and time that the claim was submitted on-chain |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| evaluation | [Evaluation](#ixo.claims.v1beta1.Evaluation) |  | evaluation is the result of one or more claim evaluations |
| payments_status | [ClaimPayments](#ixo.claims.v1beta1.ClaimPayments) |  |  |






<a name="ixo.claims.v1beta1.ClaimPayments"></a>

### ClaimPayments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| evaluation | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| approval | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| rejection | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  | PaymentStatus penalty = 5; |






<a name="ixo.claims.v1beta1.Collection"></a>

### Collection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | collection id is the incremented internal id for the collection of claims |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| admin | [string](#string) |  | admin is the account address that will authorize or revoke agents and payments (the grantor) |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| count | [uint64](#uint64) |  | count is the number of claims already submitted (internally calculated) |
| evaluated | [uint64](#uint64) |  | evaluated is the number of claims that have been evaluated (internally calculated) |
| approved | [uint64](#uint64) |  | approved is the number of claims that have been evaluated and approved (internally calculated) |
| rejected | [uint64](#uint64) |  | rejected is the number of claims that have been evaluated and rejected (internally calculated) |
| disputed | [uint64](#uint64) |  | disputed is the number of claims that have disputed status (internally calculated) |
| state | [CollectionState](#ixo.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#ixo.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |






<a name="ixo.claims.v1beta1.Dispute"></a>

### Dispute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  |  |
| type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#ixo.claims.v1beta1.DisputeData) |  |  |






<a name="ixo.claims.v1beta1.DisputeData"></a>

### DisputeData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | dispute link ***.ipfs |
| type | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [bool](#bool) |  |  |






<a name="ixo.claims.v1beta1.Evaluation"></a>

### Evaluation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id indicates which Claim this evaluation is for |
| collection_id | [string](#string) |  | collection_id indicates to which Collection the claim being evaluated belongs to |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| status | [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verfiable Credential |
| evaluation_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | evaluationDate is the date and time that the claim evaluation was submitted on-chain |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval, if empty list then use default by Collection |






<a name="ixo.claims.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_sequence | [uint64](#uint64) |  |  |
| ixo_account | [string](#string) |  |  |
| network_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the entity account address from which the payment will be made |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| timeout_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | timeout after claim/evaluation to create authZ for payment, if 0 then immidiate direct payment |






<a name="ixo.claims.v1beta1.Payments"></a>

### Payments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| evaluation | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| approval | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| rejection | [Payment](#ixo.claims.v1beta1.Payment) |  | Payment penalty = 5; |





 


<a name="ixo.claims.v1beta1.CollectionState"></a>

### CollectionState


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPEN | 0 |  |
| PAUSED | 1 |  |
| CLOSED | 2 |  |



<a name="ixo.claims.v1beta1.EvaluationStatus"></a>

### EvaluationStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| APPROVED | 1 |  |
| REJECTED | 2 |  |
| DISPUTED | 3 |  |



<a name="ixo.claims.v1beta1.PaymentStatus"></a>

### PaymentStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| NO_PAYMENT | 0 |  |
| PROMISED | 1 | agent is contracted to receive payment |
| AUTHORIZED | 2 | authz set up, no guarantee |
| GAURANTEED | 3 | escrow set up with funds blocked |
| PAID | 4 |  |
| FAILED | 5 |  |
| DISPUTED | 6 |  |



<a name="ixo.claims.v1beta1.PaymentType"></a>

### PaymentType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBMISSION | 0 |  |
| APPROVAL | 1 |  |
| EVALUATION | 2 |  |
| REJECTION | 3 |  |


 

 

 



<a name="ixo/claims/v1beta1/cosmos.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/cosmos.proto



<a name="ixo.claims.v1beta1.Input"></a>

### Input
Input models transaction input.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.claims.v1beta1.Output"></a>

### Output
Output models transaction outputs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 

 



<a name="ixo/claims/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/authz.proto



<a name="ixo.claims.v1beta1.EvaluateClaimAuthorization"></a>

### EvaluateClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [EvaluateClaimConstraints](#ixo.claims.v1beta1.EvaluateClaimConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.EvaluateClaimConstraints"></a>

### EvaluateClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_ids | [string](#string) | repeated | either collection_id or claim_ids is needed |
| agent_quota | [uint64](#uint64) |  |  |
| before_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | if zero then no before_date validation done |
| max_custom_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | max custom amount evaluator can change, if empty list must use amount defined in Token payments |






<a name="ixo.claims.v1beta1.SubmitClaimAuthorization"></a>

### SubmitClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [SubmitClaimConstraints](#ixo.claims.v1beta1.SubmitClaimConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.SubmitClaimConstraints"></a>

### SubmitClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_quota | [uint64](#uint64) |  |  |






<a name="ixo.claims.v1beta1.WithdrawPaymentAuthorization"></a>

### WithdrawPaymentAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.WithdrawPaymentConstraints"></a>

### WithdrawPaymentConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [Input](#ixo.claims.v1beta1.Input) | repeated | Inputs to the multisend tx to run to withdraw payment |
| outputs | [Output](#ixo.claims.v1beta1.Output) | repeated | Outputs for the multisend tx to run to withdraw payment |
| payment_type | [PaymentType](#ixo.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments |





 

 

 

 



<a name="ixo/claims/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/event.proto



<a name="ixo.claims.v1beta1.ClaimDisputedEvent"></a>

### ClaimDisputedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#ixo.claims.v1beta1.Dispute) |  |  |






<a name="ixo.claims.v1beta1.ClaimEvaluatedEvent"></a>

### ClaimEvaluatedEvent
ClaimEvaluatedEvent is an event triggered on a Claim evaluation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| evaluation | [Evaluation](#ixo.claims.v1beta1.Evaluation) |  |  |






<a name="ixo.claims.v1beta1.ClaimSubmittedEvent"></a>

### ClaimSubmittedEvent
CollectionCreatedEvent is an event triggered on a Claim submission


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.ClaimUpdatedEvent"></a>

### ClaimUpdatedEvent
ClaimUpdatedEvent is an event triggered on a Claim update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.CollectionCreatedEvent"></a>

### CollectionCreatedEvent
CollectionCreatedEvent is an event triggered on a Collection creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.CollectionUpdatedEvent"></a>

### CollectionUpdatedEvent
CollectionUpdatedEvent is an event triggered on a Collection update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.PaymentWithdrawCreatedEvent"></a>

### PaymentWithdrawCreatedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) |  |  |






<a name="ixo.claims.v1beta1.PaymentWithdrawnEvent"></a>

### PaymentWithdrawnEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) |  |  |





 

 

 

 



<a name="ixo/claims/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/genesis.proto



<a name="ixo.claims.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the claims module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.claims.v1beta1.Params) |  |  |
| collections | [Collection](#ixo.claims.v1beta1.Collection) | repeated |  |
| claims | [Claim](#ixo.claims.v1beta1.Claim) | repeated |  |





 

 

 

 



<a name="ixo/claims/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/query.proto



<a name="ixo.claims.v1beta1.QueryClaimListRequest"></a>

### QueryClaimListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimListResponse"></a>

### QueryClaimListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims | [Claim](#ixo.claims.v1beta1.Claim) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimRequest"></a>

### QueryClaimRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimResponse"></a>

### QueryClaimResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionListRequest"></a>

### QueryCollectionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionListResponse"></a>

### QueryCollectionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [Collection](#ixo.claims.v1beta1.Collection) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionRequest"></a>

### QueryCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionResponse"></a>

### QueryCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeListRequest"></a>

### QueryDisputeListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeListResponse"></a>

### QueryDisputeListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disputes | [Dispute](#ixo.claims.v1beta1.Dispute) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeRequest"></a>

### QueryDisputeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proof | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeResponse"></a>

### QueryDisputeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#ixo.claims.v1beta1.Dispute) |  |  |






<a name="ixo.claims.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="ixo.claims.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.claims.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="ixo.claims.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.claims.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.claims.v1beta1.QueryParamsResponse) | Parameters queries the parameters of the module. |
| Collection | [QueryCollectionRequest](#ixo.claims.v1beta1.QueryCollectionRequest) | [QueryCollectionResponse](#ixo.claims.v1beta1.QueryCollectionResponse) |  |
| CollectionList | [QueryCollectionListRequest](#ixo.claims.v1beta1.QueryCollectionListRequest) | [QueryCollectionListResponse](#ixo.claims.v1beta1.QueryCollectionListResponse) |  |
| Claim | [QueryClaimRequest](#ixo.claims.v1beta1.QueryClaimRequest) | [QueryClaimResponse](#ixo.claims.v1beta1.QueryClaimResponse) |  |
| ClaimList | [QueryClaimListRequest](#ixo.claims.v1beta1.QueryClaimListRequest) | [QueryClaimListResponse](#ixo.claims.v1beta1.QueryClaimListResponse) |  |
| Dispute | [QueryDisputeRequest](#ixo.claims.v1beta1.QueryDisputeRequest) | [QueryDisputeResponse](#ixo.claims.v1beta1.QueryDisputeResponse) |  |
| DisputeList | [QueryDisputeListRequest](#ixo.claims.v1beta1.QueryDisputeListRequest) | [QueryDisputeListResponse](#ixo.claims.v1beta1.QueryDisputeListResponse) |  |

 



<a name="ixo/claims/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/tx.proto



<a name="ixo.claims.v1beta1.MsgCreateCollection"></a>

### MsgCreateCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| admin | [string](#string) |  | admin is the account address that will authorize or revoke agents and payments (the grantor), signer for tx |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| state | [CollectionState](#ixo.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#ixo.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |






<a name="ixo.claims.v1beta1.MsgCreateCollectionResponse"></a>

### MsgCreateCollectionResponse







<a name="ixo.claims.v1beta1.MsgDisputeClaim"></a>

### MsgDisputeClaim
Agent laying dispute must be admin for Collection, or controller on
Collection entity, or have authz cap, aka is agent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  | subject_id for which this dispute is against, for now can only lay disputes against claims |
| agent_did | [string](#string) |  | agent is the DID of the agent disputing the claim, agent detials wont be saved in kvStore |
| agent_address | [string](#string) |  |  |
| dispute_type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#ixo.claims.v1beta1.DisputeData) |  |  |






<a name="ixo.claims.v1beta1.MsgDisputeClaimResponse"></a>

### MsgDisputeClaimResponse







<a name="ixo.claims.v1beta1.MsgEvaluateClaim"></a>

### MsgEvaluateClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim to make evaluation against |
| collection_id | [string](#string) |  | claimID is the unique identifier of the claim to make evaluation against |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| status | [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verfiable Credential |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval, if empty list then use default by Collection |






<a name="ixo.claims.v1beta1.MsgEvaluateClaimResponse"></a>

### MsgEvaluateClaimResponse







<a name="ixo.claims.v1beta1.MsgSubmitClaim"></a>

### MsgSubmitClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgSubmitClaimResponse"></a>

### MsgSubmitClaimResponse







<a name="ixo.claims.v1beta1.MsgWithdrawPayment"></a>

### MsgWithdrawPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [Input](#ixo.claims.v1beta1.Input) | repeated | Inputs to the multisend tx to run to withdraw payment |
| outputs | [Output](#ixo.claims.v1beta1.Output) | repeated | Outputs for the multisend tx to run to withdraw payment |
| payment_type | [PaymentType](#ixo.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgWithdrawPaymentResponse"></a>

### MsgWithdrawPaymentResponse






 

 

 


<a name="ixo.claims.v1beta1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCollection | [MsgCreateCollection](#ixo.claims.v1beta1.MsgCreateCollection) | [MsgCreateCollectionResponse](#ixo.claims.v1beta1.MsgCreateCollectionResponse) |  |
| SubmitClaim | [MsgSubmitClaim](#ixo.claims.v1beta1.MsgSubmitClaim) | [MsgSubmitClaimResponse](#ixo.claims.v1beta1.MsgSubmitClaimResponse) |  |
| EvaluateClaim | [MsgEvaluateClaim](#ixo.claims.v1beta1.MsgEvaluateClaim) | [MsgEvaluateClaimResponse](#ixo.claims.v1beta1.MsgEvaluateClaimResponse) |  |
| DisputeClaim | [MsgDisputeClaim](#ixo.claims.v1beta1.MsgDisputeClaim) | [MsgDisputeClaimResponse](#ixo.claims.v1beta1.MsgDisputeClaimResponse) |  |
| WithdrawPayment | [MsgWithdrawPayment](#ixo.claims.v1beta1.MsgWithdrawPayment) | [MsgWithdrawPaymentResponse](#ixo.claims.v1beta1.MsgWithdrawPaymentResponse) |  |

 



<a name="ixo/iid/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/types.proto



<a name="ixo.iid.v1beta1.AccordedRight"></a>

### AccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| mechanism | [string](#string) |  |  |
| message | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.Context"></a>

### Context



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| val | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.IidMetadata"></a>

### IidMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versionId | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deactivated | [bool](#bool) |  |  |






<a name="ixo.iid.v1beta1.LinkedClaim"></a>

### LinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| mediaType | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.LinkedEntity"></a>

### LinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| relationship | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.LinkedResource"></a>

### LinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| mediaType | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.Service"></a>

### Service



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.VerificationMethod"></a>

### VerificationMethod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| controller | [string](#string) |  |  |
| blockchainAccountID | [string](#string) |  |  |
| publicKeyHex | [string](#string) |  |  |
| publicKeyMultibase | [string](#string) |  |  |
| publicKeyBase58 | [string](#string) |  |  |





 

 

 

 



<a name="ixo/iid/v1beta1/iid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/iid.proto



<a name="ixo.iid.v1beta1.IidDocument"></a>

### IidDocument
type entity account
relationship entity account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated | @context is spec for did document. |
| id | [string](#string) |  | id represents the id for the did document. |
| controller | [string](#string) | repeated | A DID controller is an entity that is authorized to make changes to a DID document. cfr. https://www.w3.org/TR/did-core/#did-controller |
| verificationMethod | [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod) | repeated | A DID document can express verification methods, such as cryptographic public keys, which can be used to authenticate or authorize interactions with the DID subject or associated parties. https://www.w3.org/TR/did-core/#verification-methods |
| service | [Service](#ixo.iid.v1beta1.Service) | repeated | Services are used in DID documents to express ways of communicating with the DID subject or associated entities. https://www.w3.org/TR/did-core/#services |
| authentication | [string](#string) | repeated | NOTE: below this line there are the relationships Authentication represents public key associated with the did document. cfr. https://www.w3.org/TR/did-core/#authentication |
| assertionMethod | [string](#string) | repeated | Used to specify how the DID subject is expected to express claims, such as for the purposes of issuing a Verifiable Credential. cfr. https://www.w3.org/TR/did-core/#assertion |
| keyAgreement | [string](#string) | repeated | used to specify how an entity can generate encryption material in order to transmit confidential information intended for the DID subject. https://www.w3.org/TR/did-core/#key-agreement |
| capabilityInvocation | [string](#string) | repeated | Used to specify a verification method that might be used by the DID subject to invoke a cryptographic capability, such as the authorization to update the DID Document. https://www.w3.org/TR/did-core/#capability-invocation |
| capabilityDelegation | [string](#string) | repeated | Used to specify a mechanism that might be used by the DID subject to delegate a cryptographic capability to another party. https://www.w3.org/TR/did-core/#capability-delegation |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| metadata | [IidMetadata](#ixo.iid.v1beta1.IidMetadata) |  | Metadata concerning the IidDocument such as versionId, created, updated and deactivated |





 

 

 

 



<a name="ixo/entity/v1beta1/entity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/entity.proto



<a name="ixo.entity.v1beta1.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id represents the id for the entity document. |
| type | [string](#string) |  | Type of entity, eg protocol or asset |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Address of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Credentials of the enitity to be verified |
| entity_verified | [bool](#bool) |  | Used as check whether the credentials of entity is verified |
| metadata | [EntityMetadata](#ixo.entity.v1beta1.EntityMetadata) |  | Metadata concerning the Entity such as versionId, created, updated and deactivated |






<a name="ixo.entity.v1beta1.EntityMetadata"></a>

### EntityMetadata
EntityMetadata defines metadata associated to a entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version_id | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ixo.entity.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nftContractAddress | [string](#string) |  |  |
| nftContractMinter | [string](#string) |  |  |
| createSequence | [uint64](#uint64) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/event.proto



<a name="ixo.entity.v1beta1.EntityCreatedEvent"></a>

### EntityCreatedEvent
EntityCreatedEvent is an event triggered on a Entity creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| owner | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityTransferredEvent"></a>

### EntityTransferredEvent
EntityTransferredEvent is an event triggered on a entity transfer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityUpdatedEvent"></a>

### EntityUpdatedEvent
EntityUpdatedEvent is an event triggered on a entity document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| owner | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityVerifiedUpdatedEvent"></a>

### EntityVerifiedUpdatedEvent
EntityVerifiedUpdatedEvent is an event triggered on a entity verified
document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| entity_verified | [bool](#bool) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/genesis.proto



<a name="ixo.entity.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the project module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#ixo.entity.v1beta1.Entity) | repeated |  |
| params | [Params](#ixo.entity.v1beta1.Params) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/proposal.proto



<a name="ixo.entity.v1beta1.InitializeNftContract"></a>

### InitializeNftContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| NftContractCodeId | [uint64](#uint64) |  |  |
| NftMinterAddress | [string](#string) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/query.proto



<a name="ixo.entity.v1beta1.QueryEntityIidDocumentRequest"></a>

### QueryEntityIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityIidDocumentResponse"></a>

### QueryEntityIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [ixo.iid.v1beta1.IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityListRequest"></a>

### QueryEntityListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | string type = 2; string status = 3; |






<a name="ixo.entity.v1beta1.QueryEntityListResponse"></a>

### QueryEntityListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#ixo.entity.v1beta1.Entity) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityMetadataRequest"></a>

### QueryEntityMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityMetadataResponse"></a>

### QueryEntityMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityRequest"></a>

### QueryEntityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityResponse"></a>

### QueryEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| iidDocument | [ixo.iid.v1beta1.IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityVerifiedRequest"></a>

### QueryEntityVerifiedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityVerifiedResponse"></a>

### QueryEntityVerifiedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_verified | [bool](#bool) |  |  |





 

 

 


<a name="ixo.entity.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Entity | [QueryEntityRequest](#ixo.entity.v1beta1.QueryEntityRequest) | [QueryEntityResponse](#ixo.entity.v1beta1.QueryEntityResponse) |  |
| EntityMetaData | [QueryEntityMetadataRequest](#ixo.entity.v1beta1.QueryEntityMetadataRequest) | [QueryEntityMetadataResponse](#ixo.entity.v1beta1.QueryEntityMetadataResponse) |  |
| EntityIidDocument | [QueryEntityIidDocumentRequest](#ixo.entity.v1beta1.QueryEntityIidDocumentRequest) | [QueryEntityIidDocumentResponse](#ixo.entity.v1beta1.QueryEntityIidDocumentResponse) |  |
| EntityVerified | [QueryEntityVerifiedRequest](#ixo.entity.v1beta1.QueryEntityVerifiedRequest) | [QueryEntityVerifiedResponse](#ixo.entity.v1beta1.QueryEntityVerifiedResponse) |  |
| EntityList | [QueryEntityListRequest](#ixo.entity.v1beta1.QueryEntityListRequest) | [QueryEntityListResponse](#ixo.entity.v1beta1.QueryEntityListResponse) |  |

 



<a name="ixo/iid/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/tx.proto



<a name="ixo.iid.v1beta1.MsgAddAccordedRight"></a>

### MsgAddAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) |  | the Accorded right to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddAccordedRightResponse"></a>

### MsgAddAccordedRightResponse







<a name="ixo.iid.v1beta1.MsgAddController"></a>

### MsgAddController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to add as a controller of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddControllerResponse"></a>

### MsgAddControllerResponse







<a name="ixo.iid.v1beta1.MsgAddIidContext"></a>

### MsgAddIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| context | [Context](#ixo.iid.v1beta1.Context) |  | the context to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddIidContextResponse"></a>

### MsgAddIidContextResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedClaim"></a>

### MsgAddLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) |  | the claim to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedClaimResponse"></a>

### MsgAddLinkedClaimResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedEntity"></a>

### MsgAddLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) |  | the entity to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedEntityResponse"></a>

### MsgAddLinkedEntityResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedResource"></a>

### MsgAddLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedResourceResponse"></a>

### MsgAddLinkedResourceResponse







<a name="ixo.iid.v1beta1.MsgAddService"></a>

### MsgAddService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_data | [Service](#ixo.iid.v1beta1.Service) |  | the service data to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddServiceResponse"></a>

### MsgAddServiceResponse







<a name="ixo.iid.v1beta1.MsgAddVerification"></a>

### MsgAddVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| verification | [Verification](#ixo.iid.v1beta1.Verification) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddVerificationResponse"></a>

### MsgAddVerificationResponse







<a name="ixo.iid.v1beta1.MsgCreateIidDocument"></a>

### MsgCreateIidDocument
MsgCreateDidDocument defines a SDK message for creating a new did.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#ixo.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#ixo.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="ixo.iid.v1beta1.MsgCreateIidDocumentResponse"></a>

### MsgCreateIidDocumentResponse







<a name="ixo.iid.v1beta1.MsgDeactivateIID"></a>

### MsgDeactivateIID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| state | [bool](#bool) |  |  |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeactivateIIDResponse"></a>

### MsgDeactivateIIDResponse







<a name="ixo.iid.v1beta1.MsgDeleteAccordedRight"></a>

### MsgDeleteAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| right_id | [string](#string) |  | the Accorded right id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteAccordedRightResponse"></a>

### MsgDeleteAccordedRightResponse







<a name="ixo.iid.v1beta1.MsgDeleteController"></a>

### MsgDeleteController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to remove from the list of controllers of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteControllerResponse"></a>

### MsgDeleteControllerResponse







<a name="ixo.iid.v1beta1.MsgDeleteIidContext"></a>

### MsgDeleteIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| contextKey | [string](#string) |  | the context key |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteIidContextResponse"></a>

### MsgDeleteIidContextResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedClaim"></a>

### MsgDeleteLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| claim_id | [string](#string) |  | the claim id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse"></a>

### MsgDeleteLinkedClaimResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedEntity"></a>

### MsgDeleteLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| entity_id | [string](#string) |  | the entity id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse"></a>

### MsgDeleteLinkedEntityResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedResource"></a>

### MsgDeleteLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| resource_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse"></a>

### MsgDeleteLinkedResourceResponse







<a name="ixo.iid.v1beta1.MsgDeleteService"></a>

### MsgDeleteService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteServiceResponse"></a>

### MsgDeleteServiceResponse







<a name="ixo.iid.v1beta1.MsgRevokeVerification"></a>

### MsgRevokeVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgRevokeVerificationResponse"></a>

### MsgRevokeVerificationResponse







<a name="ixo.iid.v1beta1.MsgSetVerificationRelationships"></a>

### MsgSetVerificationRelationships



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| relationships | [string](#string) | repeated | the list of relationships to set |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse"></a>

### MsgSetVerificationRelationshipsResponse







<a name="ixo.iid.v1beta1.MsgUpdateIidDocument"></a>

### MsgUpdateIidDocument
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#ixo.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#ixo.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="ixo.iid.v1beta1.MsgUpdateIidDocumentResponse"></a>

### MsgUpdateIidDocumentResponse







<a name="ixo.iid.v1beta1.Verification"></a>

### Verification
Verification is a message that allows to assign a verification method
to one or more verification relationships


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relationships | [string](#string) | repeated | verificationRelationships defines which relationships are allowed to use the verification method

relationships that the method is allowed into. |
| method | [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod) |  | public key associated with the did document. |
| context | [string](#string) | repeated | additional contexts (json ld schemas) |





 

 

 


<a name="ixo.iid.v1beta1.Msg"></a>

### Msg
Msg defines the identity Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateIidDocument | [MsgCreateIidDocument](#ixo.iid.v1beta1.MsgCreateIidDocument) | [MsgCreateIidDocumentResponse](#ixo.iid.v1beta1.MsgCreateIidDocumentResponse) | CreateDidDocument defines a method for creating a new identity. |
| UpdateIidDocument | [MsgUpdateIidDocument](#ixo.iid.v1beta1.MsgUpdateIidDocument) | [MsgUpdateIidDocumentResponse](#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse) | UpdateDidDocument defines a method for updating an identity. |
| AddVerification | [MsgAddVerification](#ixo.iid.v1beta1.MsgAddVerification) | [MsgAddVerificationResponse](#ixo.iid.v1beta1.MsgAddVerificationResponse) | AddVerificationMethod adds a new verification method |
| RevokeVerification | [MsgRevokeVerification](#ixo.iid.v1beta1.MsgRevokeVerification) | [MsgRevokeVerificationResponse](#ixo.iid.v1beta1.MsgRevokeVerificationResponse) | RevokeVerification remove the verification method and all associated verification Relations |
| SetVerificationRelationships | [MsgSetVerificationRelationships](#ixo.iid.v1beta1.MsgSetVerificationRelationships) | [MsgSetVerificationRelationshipsResponse](#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse) | SetVerificationRelationships overwrite current verification relationships |
| AddService | [MsgAddService](#ixo.iid.v1beta1.MsgAddService) | [MsgAddServiceResponse](#ixo.iid.v1beta1.MsgAddServiceResponse) | AddService add a new service |
| DeleteService | [MsgDeleteService](#ixo.iid.v1beta1.MsgDeleteService) | [MsgDeleteServiceResponse](#ixo.iid.v1beta1.MsgDeleteServiceResponse) | DeleteService delete an existing service |
| AddController | [MsgAddController](#ixo.iid.v1beta1.MsgAddController) | [MsgAddControllerResponse](#ixo.iid.v1beta1.MsgAddControllerResponse) | AddService add a new service |
| DeleteController | [MsgDeleteController](#ixo.iid.v1beta1.MsgDeleteController) | [MsgDeleteControllerResponse](#ixo.iid.v1beta1.MsgDeleteControllerResponse) | DeleteService delete an existing service |
| AddLinkedResource | [MsgAddLinkedResource](#ixo.iid.v1beta1.MsgAddLinkedResource) | [MsgAddLinkedResourceResponse](#ixo.iid.v1beta1.MsgAddLinkedResourceResponse) | Add / Delete Linked Resource |
| DeleteLinkedResource | [MsgDeleteLinkedResource](#ixo.iid.v1beta1.MsgDeleteLinkedResource) | [MsgDeleteLinkedResourceResponse](#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse) |  |
| AddLinkedClaim | [MsgAddLinkedClaim](#ixo.iid.v1beta1.MsgAddLinkedClaim) | [MsgAddLinkedClaimResponse](#ixo.iid.v1beta1.MsgAddLinkedClaimResponse) | Add / Delete Linked Claims |
| DeleteLinkedClaim | [MsgDeleteLinkedClaim](#ixo.iid.v1beta1.MsgDeleteLinkedClaim) | [MsgDeleteLinkedClaimResponse](#ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse) |  |
| AddLinkedEntity | [MsgAddLinkedEntity](#ixo.iid.v1beta1.MsgAddLinkedEntity) | [MsgAddLinkedEntityResponse](#ixo.iid.v1beta1.MsgAddLinkedEntityResponse) | Add / Delete Linked Entity |
| DeleteLinkedEntity | [MsgDeleteLinkedEntity](#ixo.iid.v1beta1.MsgDeleteLinkedEntity) | [MsgDeleteLinkedEntityResponse](#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse) |  |
| AddAccordedRight | [MsgAddAccordedRight](#ixo.iid.v1beta1.MsgAddAccordedRight) | [MsgAddAccordedRightResponse](#ixo.iid.v1beta1.MsgAddAccordedRightResponse) | Add / Delete Accorded Right |
| DeleteAccordedRight | [MsgDeleteAccordedRight](#ixo.iid.v1beta1.MsgDeleteAccordedRight) | [MsgDeleteAccordedRightResponse](#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse) |  |
| AddIidContext | [MsgAddIidContext](#ixo.iid.v1beta1.MsgAddIidContext) | [MsgAddIidContextResponse](#ixo.iid.v1beta1.MsgAddIidContextResponse) | Add / Delete Context |
| DeactivateIID | [MsgDeactivateIID](#ixo.iid.v1beta1.MsgDeactivateIID) | [MsgDeactivateIIDResponse](#ixo.iid.v1beta1.MsgDeactivateIIDResponse) |  |
| DeleteIidContext | [MsgDeleteIidContext](#ixo.iid.v1beta1.MsgDeleteIidContext) | [MsgDeleteIidContextResponse](#ixo.iid.v1beta1.MsgDeleteIidContextResponse) |  |

 



<a name="ixo/entity/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/tx.proto



<a name="ixo.entity.v1beta1.MsgCreateEntity"></a>

### MsgCreateEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_type | [string](#string) |  | An Entity Type as defined by the implementer |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| controller | [string](#string) | repeated | the list of controller DIDs |
| context | [ixo.iid.v1beta1.Context](#ixo.iid.v1beta1.Context) | repeated | JSON-LD contexts |
| verification | [ixo.iid.v1beta1.Verification](#ixo.iid.v1beta1.Verification) | repeated | Verification Methods and Verification Relationships |
| service | [ixo.iid.v1beta1.Service](#ixo.iid.v1beta1.Service) | repeated | Service endpoints |
| accorded_right | [ixo.iid.v1beta1.AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated | Legal or Electronic Rights and associated Object Capabilities |
| linked_resource | [ixo.iid.v1beta1.LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated | Digital resources associated with the Subject |
| linked_entity | [ixo.iid.v1beta1.LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated | DID of a linked Entity and its relationship with the Subject |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Address of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| owner_did | [string](#string) |  | Owner of the Entity NFT | The ownersdid used to sign this transaction. |
| owner_address | [string](#string) |  | The ownersdid address used to sign this transaction. |
| data | [bytes](#bytes) |  | Extention data |
| alsoKnownAs | [string](#string) |  |  |
| linked_claim | [ixo.iid.v1beta1.LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated | Digital claims associated with the Subject |






<a name="ixo.entity.v1beta1.MsgCreateEntityResponse"></a>

### MsgCreateEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_id | [string](#string) |  |  |
| entity_type | [string](#string) |  |  |
| entity_status | [int32](#int32) |  |  |






<a name="ixo.entity.v1beta1.MsgTransferEntity"></a>

### MsgTransferEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner_did | [string](#string) |  | The owner_did used to sign this transaction. |
| owner_address | [string](#string) |  | The owner_address used to sign this transaction. |
| recipient_did | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.MsgTransferEntityResponse"></a>

### MsgTransferEntityResponse







<a name="ixo.entity.v1beta1.MsgUpdateEntity"></a>

### MsgUpdateEntity
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| controller_did | [string](#string) |  | The controllerDid used to sign this transaction. |
| controller_address | [string](#string) |  | The controllerAddress used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgUpdateEntityResponse"></a>

### MsgUpdateEntityResponse







<a name="ixo.entity.v1beta1.MsgUpdateEntityVerified"></a>

### MsgUpdateEntityVerified
Only relayer nodes can update entity field &#39;entityVerified&#39;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_verified | [bool](#bool) |  | Whether entity is verified or not based on credentials |
| relayer_node_did | [string](#string) |  | The relayer node&#39;s did used to sign this transaction. |
| relayer_node_address | [string](#string) |  | The relayer node&#39;s address used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse"></a>

### MsgUpdateEntityVerifiedResponse






 

 

 


<a name="ixo.entity.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEntity | [MsgCreateEntity](#ixo.entity.v1beta1.MsgCreateEntity) | [MsgCreateEntityResponse](#ixo.entity.v1beta1.MsgCreateEntityResponse) | CreateEntity defines a method for creating a entity. |
| UpdateEntity | [MsgUpdateEntity](#ixo.entity.v1beta1.MsgUpdateEntity) | [MsgUpdateEntityResponse](#ixo.entity.v1beta1.MsgUpdateEntityResponse) | UpdateEntity defines a method for updating a entity |
| UpdateEntityVerified | [MsgUpdateEntityVerified](#ixo.entity.v1beta1.MsgUpdateEntityVerified) | [MsgUpdateEntityVerifiedResponse](#ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse) | UpdateEntityVerified defines a method for updating if an entity is verified |
| TransferEntity | [MsgTransferEntity](#ixo.entity.v1beta1.MsgTransferEntity) | [MsgTransferEntityResponse](#ixo.entity.v1beta1.MsgTransferEntityResponse) | Transfers an entity and its nft to the recipient |

 



<a name="ixo/iid/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/event.proto



<a name="ixo.iid.v1beta1.IidDocumentCreatedEvent"></a>

### IidDocumentCreatedEvent
IidDocumentCreatedEvent is triggered when a new IidDocument is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.iid.v1beta1.IidDocumentUpdatedEvent"></a>

### IidDocumentUpdatedEvent
DidDocumentUpdatedEvent is an event triggered on a DID document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |





 

 

 

 



<a name="ixo/iid/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/genesis.proto



<a name="ixo.iid.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iid_docs | [IidDocument](#ixo.iid.v1beta1.IidDocument) | repeated |  |





 

 

 

 



<a name="ixo/iid/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/query.proto



<a name="ixo.iid.v1beta1.QueryIidDocumentRequest"></a>

### QueryIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | did id of iid document querying |






<a name="ixo.iid.v1beta1.QueryIidDocumentResponse"></a>

### QueryIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.iid.v1beta1.QueryIidDocumentsRequest"></a>

### QueryIidDocumentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="ixo.iid.v1beta1.QueryIidDocumentsResponse"></a>

### QueryIidDocumentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocuments | [IidDocument](#ixo.iid.v1beta1.IidDocument) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 

 

 


<a name="ixo.iid.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IidDocuments | [QueryIidDocumentsRequest](#ixo.iid.v1beta1.QueryIidDocumentsRequest) | [QueryIidDocumentsResponse](#ixo.iid.v1beta1.QueryIidDocumentsResponse) | IidDocuments queries all iid documents that match the given status. |
| IidDocument | [QueryIidDocumentRequest](#ixo.iid.v1beta1.QueryIidDocumentRequest) | [QueryIidDocumentResponse](#ixo.iid.v1beta1.QueryIidDocumentResponse) | IidDocument queries a iid documents with an id. |

 



<a name="ixo/legacy/did/did.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/legacy/did/did.proto



<a name="legacydid.Claim"></a>

### Claim
The claim section of a credential, indicating if the DID is KYC validated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| KYC_validated | [bool](#bool) |  |  |






<a name="legacydid.DidCredential"></a>

### DidCredential
Digital identity credential issued to an ixo DID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cred_type | [string](#string) | repeated |  |
| issuer | [string](#string) |  |  |
| issued | [string](#string) |  |  |
| claim | [Claim](#legacydid.Claim) |  |  |






<a name="legacydid.IxoDid"></a>

### IxoDid
An ixo DID with public and private keys, based on the Sovrin DID spec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| verify_key | [string](#string) |  |  |
| encryption_public_key | [string](#string) |  |  |
| secret | [Secret](#legacydid.Secret) |  |  |






<a name="legacydid.Secret"></a>

### Secret
The private section of an ixo DID, based on the Sovrin DID spec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seed | [string](#string) |  |  |
| sign_key | [string](#string) |  |  |
| encryption_private_key | [string](#string) |  |  |





 

 

 

 



<a name="ixo/legacy/did/diddoc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/legacy/did/diddoc.proto



<a name="legacydid.BaseDidDoc"></a>

### BaseDidDoc
BaseDidDoc defines a base DID document type. It implements the DidDoc
interface.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| credentials | [DidCredential](#legacydid.DidCredential) | repeated |  |





 

 

 

 



<a name="ixo/payments/v1/payments.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/payments/v1/payments.proto



<a name="ixo.payments.v1.BlockPeriod"></a>

### BlockPeriod
BlockPeriod implements the Period interface and specifies a period in terms
of number of blocks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="ixo.payments.v1.Discount"></a>

### Discount
Discount contains details about a discount which can be granted to payers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| percent | [string](#string) |  |  |






<a name="ixo.payments.v1.DistributionShare"></a>

### DistributionShare
DistributionShare specifies the share of a specific payment an address will
receive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| percentage | [string](#string) |  |  |






<a name="ixo.payments.v1.PaymentContract"></a>

### PaymentContract
PaymentContract specifies an agreement between a payer and payee/s which can
be invoked once or multiple times to effect payment/s.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_template_id | [string](#string) |  |  |
| creator | [string](#string) |  |  |
| payer | [string](#string) |  |  |
| recipients | [DistributionShare](#ixo.payments.v1.DistributionShare) | repeated |  |
| cumulative_pay | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_remainder | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| can_deauthorise | [bool](#bool) |  |  |
| authorised | [bool](#bool) |  |  |
| discount_id | [string](#string) |  |  |






<a name="ixo.payments.v1.PaymentTemplate"></a>

### PaymentTemplate
PaymentTemplate contains details about a payment, with no info about the
payer or payee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| payment_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_minimum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| payment_maximum | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| discounts | [Discount](#ixo.payments.v1.Discount) | repeated |  |






<a name="ixo.payments.v1.Subscription"></a>

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






<a name="ixo.payments.v1.TestPeriod"></a>

### TestPeriod
TestPeriod implements the Period interface and is identical to BlockPeriod,
except it ignores the context in periodEnded() and periodStarted().


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_length | [int64](#int64) |  |  |
| period_start_block | [int64](#int64) |  |  |






<a name="ixo.payments.v1.TimePeriod"></a>

### TimePeriod
TimePeriod implements the Period interface and specifies a period in terms of
time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| period_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| period_start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="ixo/payments/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/payments/v1/genesis.proto



<a name="ixo.payments.v1.GenesisState"></a>

### GenesisState
GenesisState defines the payments module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_templates | [PaymentTemplate](#ixo.payments.v1.PaymentTemplate) | repeated |  |
| payment_contracts | [PaymentContract](#ixo.payments.v1.PaymentContract) | repeated |  |
| subscriptions | [Subscription](#ixo.payments.v1.Subscription) | repeated |  |





 

 

 

 



<a name="ixo/payments/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/payments/v1/query.proto



<a name="ixo.payments.v1.QueryPaymentContractRequest"></a>

### QueryPaymentContractRequest
QueryPaymentContractRequest is the request type for the Query/PaymentContract
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |






<a name="ixo.payments.v1.QueryPaymentContractResponse"></a>

### QueryPaymentContractResponse
QueryPaymentContractResponse is the response type for the
Query/PaymentContract RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract | [PaymentContract](#ixo.payments.v1.PaymentContract) |  |  |






<a name="ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest"></a>

### QueryPaymentContractsByIdPrefixRequest
QueryPaymentContractsByIdPrefixRequest is the request type for the
Query/PaymentContractsByIdPrefix RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts_id_prefix | [string](#string) |  |  |






<a name="ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse"></a>

### QueryPaymentContractsByIdPrefixResponse
QueryPaymentContractsByIdPrefixResponse is the response type for the
Query/PaymentContractsByIdPrefix RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contracts | [PaymentContract](#ixo.payments.v1.PaymentContract) | repeated |  |






<a name="ixo.payments.v1.QueryPaymentTemplateRequest"></a>

### QueryPaymentTemplateRequest
QueryPaymentTemplateRequest is the request type for the Query/PaymentTemplate
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template_id | [string](#string) |  |  |






<a name="ixo.payments.v1.QueryPaymentTemplateResponse"></a>

### QueryPaymentTemplateResponse
QueryPaymentTemplateResponse is the response type for the
Query/PaymentTemplate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_template | [PaymentTemplate](#ixo.payments.v1.PaymentTemplate) |  |  |






<a name="ixo.payments.v1.QuerySubscriptionRequest"></a>

### QuerySubscriptionRequest
QuerySubscriptionRequest is the request type for the Query/Subscription RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |






<a name="ixo.payments.v1.QuerySubscriptionResponse"></a>

### QuerySubscriptionResponse
QuerySubscriptionResponse is the response type for the Query/Subscription RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#ixo.payments.v1.Subscription) |  |  |





 

 

 


<a name="ixo.payments.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PaymentTemplate | [QueryPaymentTemplateRequest](#ixo.payments.v1.QueryPaymentTemplateRequest) | [QueryPaymentTemplateResponse](#ixo.payments.v1.QueryPaymentTemplateResponse) | PaymentTemplate queries info of a specific payment template. |
| PaymentContract | [QueryPaymentContractRequest](#ixo.payments.v1.QueryPaymentContractRequest) | [QueryPaymentContractResponse](#ixo.payments.v1.QueryPaymentContractResponse) | PaymentContract queries info of a specific payment contract. |
| PaymentContractsByIdPrefix | [QueryPaymentContractsByIdPrefixRequest](#ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest) | [QueryPaymentContractsByIdPrefixResponse](#ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse) | PaymentContractsByIdPrefix lists all payment contracts having an id with a specific prefix. |
| Subscription | [QuerySubscriptionRequest](#ixo.payments.v1.QuerySubscriptionRequest) | [QuerySubscriptionResponse](#ixo.payments.v1.QuerySubscriptionResponse) | Subscription queries info of a specific Subscription. |

 



<a name="ixo/payments/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/payments/v1/tx.proto



<a name="ixo.payments.v1.MsgCreatePaymentContract"></a>

### MsgCreatePaymentContract
MsgCreatePaymentContract defines a message for creating a payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| payment_template_id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| payer | [string](#string) |  |  |
| recipients | [DistributionShare](#ixo.payments.v1.DistributionShare) | repeated |  |
| can_deauthorise | [bool](#bool) |  |  |
| discount_id | [string](#string) |  |  |
| creator_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgCreatePaymentContractResponse"></a>

### MsgCreatePaymentContractResponse
MsgCreatePaymentContractResponse defines the Msg/CreatePaymentContract
response type.






<a name="ixo.payments.v1.MsgCreatePaymentTemplate"></a>

### MsgCreatePaymentTemplate
MsgCreatePaymentTemplate defines a message for creating a payment template.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| payment_template | [PaymentTemplate](#ixo.payments.v1.PaymentTemplate) |  |  |
| creator_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgCreatePaymentTemplateResponse"></a>

### MsgCreatePaymentTemplateResponse
MsgCreatePaymentTemplateResponse defines the Msg/CreatePaymentTemplate
response type.






<a name="ixo.payments.v1.MsgCreateSubscription"></a>

### MsgCreateSubscription
MsgCreateSubscription defines a message for creating a subscription.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_did | [string](#string) |  |  |
| subscription_id | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| max_periods | [string](#string) |  |  |
| period | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| creator_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgCreateSubscriptionResponse"></a>

### MsgCreateSubscriptionResponse
MsgCreateSubscriptionResponse defines the Msg/CreateSubscription response
type.






<a name="ixo.payments.v1.MsgEffectPayment"></a>

### MsgEffectPayment
MsgEffectPayment defines a message for putting a specific payment contract
into effect.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| partial_payment_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgEffectPaymentResponse"></a>

### MsgEffectPaymentResponse
MsgEffectPaymentResponse defines the Msg/EffectPayment response type.






<a name="ixo.payments.v1.MsgGrantDiscount"></a>

### MsgGrantDiscount
MsgGrantDiscount defines a message for granting a discount to a payer on a
specific payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| discount_id | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgGrantDiscountResponse"></a>

### MsgGrantDiscountResponse
MsgGrantDiscountResponse defines the Msg/GrantDiscount response type.






<a name="ixo.payments.v1.MsgRevokeDiscount"></a>

### MsgRevokeDiscount
MsgRevokeDiscount defines a message for revoking a discount previously
granted to a payer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| payment_contract_id | [string](#string) |  |  |
| holder | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgRevokeDiscountResponse"></a>

### MsgRevokeDiscountResponse
MsgRevokeDiscountResponse defines the Msg/RevokeDiscount response type.






<a name="ixo.payments.v1.MsgSetPaymentContractAuthorisation"></a>

### MsgSetPaymentContractAuthorisation
MsgSetPaymentContractAuthorisation defines a message for authorising or
deauthorising a payment contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_contract_id | [string](#string) |  |  |
| payer_did | [string](#string) |  |  |
| authorised | [bool](#bool) |  |  |
| payer_address | [string](#string) |  |  |






<a name="ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse"></a>

### MsgSetPaymentContractAuthorisationResponse
MsgSetPaymentContractAuthorisationResponse defines the
Msg/SetPaymentContractAuthorisation response type.





 

 

 


<a name="ixo.payments.v1.Msg"></a>

### Msg
Msg defines the payments Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetPaymentContractAuthorisation | [MsgSetPaymentContractAuthorisation](#ixo.payments.v1.MsgSetPaymentContractAuthorisation) | [MsgSetPaymentContractAuthorisationResponse](#ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse) | SetPaymentContractAuthorisation defines a method for authorising or deauthorising a payment contract. |
| CreatePaymentTemplate | [MsgCreatePaymentTemplate](#ixo.payments.v1.MsgCreatePaymentTemplate) | [MsgCreatePaymentTemplateResponse](#ixo.payments.v1.MsgCreatePaymentTemplateResponse) | CreatePaymentTemplate defines a method for creating a payment template. |
| CreatePaymentContract | [MsgCreatePaymentContract](#ixo.payments.v1.MsgCreatePaymentContract) | [MsgCreatePaymentContractResponse](#ixo.payments.v1.MsgCreatePaymentContractResponse) | CreatePaymentContract defines a method for creating a payment contract. |
| CreateSubscription | [MsgCreateSubscription](#ixo.payments.v1.MsgCreateSubscription) | [MsgCreateSubscriptionResponse](#ixo.payments.v1.MsgCreateSubscriptionResponse) | CreateSubscription defines a method for creating a subscription. |
| GrantDiscount | [MsgGrantDiscount](#ixo.payments.v1.MsgGrantDiscount) | [MsgGrantDiscountResponse](#ixo.payments.v1.MsgGrantDiscountResponse) | GrantDiscount defines a method for granting a discount to a payer on a specific payment contract. |
| RevokeDiscount | [MsgRevokeDiscount](#ixo.payments.v1.MsgRevokeDiscount) | [MsgRevokeDiscountResponse](#ixo.payments.v1.MsgRevokeDiscountResponse) | RevokeDiscount defines a method for revoking a discount previously granted to a payer. |
| EffectPayment | [MsgEffectPayment](#ixo.payments.v1.MsgEffectPayment) | [MsgEffectPaymentResponse](#ixo.payments.v1.MsgEffectPaymentResponse) | EffectPayment defines a method for putting a specific payment contract into effect. |

 



<a name="ixo/project/v1/project.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/project/v1/project.proto



<a name="ixo.project.v1.AccountMap"></a>

### AccountMap
AccountMap maps a specific project&#39;s account names to the accounts&#39;
addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| map | [AccountMap.MapEntry](#ixo.project.v1.AccountMap.MapEntry) | repeated |  |






<a name="ixo.project.v1.AccountMap.MapEntry"></a>

### AccountMap.MapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="ixo.project.v1.Claim"></a>

### Claim
Claim contains details required to start a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| template_id | [string](#string) |  |  |
| claimer_did | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="ixo.project.v1.Claims"></a>

### Claims
Claims contains a list of type Claim.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims_list | [Claim](#ixo.project.v1.Claim) | repeated |  |






<a name="ixo.project.v1.CreateAgentDoc"></a>

### CreateAgentDoc
CreateAgentDoc contains details required to create an agent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_did | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="ixo.project.v1.CreateClaimDoc"></a>

### CreateClaimDoc
CreateClaimDoc contains details required to create a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| claim_template_id | [string](#string) |  |  |






<a name="ixo.project.v1.CreateEvaluationDoc"></a>

### CreateEvaluationDoc
CreateEvaluationDoc contains details required to create an evaluation for a
specific claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="ixo.project.v1.GenesisAccountMap"></a>

### GenesisAccountMap
GenesisAccountMap is a type used at genesis that maps a specific project&#39;s
account names to the accounts&#39; addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| map | [GenesisAccountMap.MapEntry](#ixo.project.v1.GenesisAccountMap.MapEntry) | repeated |  |






<a name="ixo.project.v1.GenesisAccountMap.MapEntry"></a>

### GenesisAccountMap.MapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="ixo.project.v1.Params"></a>

### Params
Params defines the parameters for the project module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo_did | [string](#string) |  |  |
| project_minimum_initial_funding | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| oracle_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |






<a name="ixo.project.v1.ProjectDoc"></a>

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






<a name="ixo.project.v1.UpdateAgentDoc"></a>

### UpdateAgentDoc
UpdateAgentDoc contains details required to update an agent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| did | [string](#string) |  |  |
| status | [string](#string) |  |  |
| role | [string](#string) |  |  |






<a name="ixo.project.v1.UpdateProjectStatusDoc"></a>

### UpdateProjectStatusDoc
UpdateProjectStatusDoc contains details required to update a project&#39;s
status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| eth_funding_txn_id | [string](#string) |  |  |






<a name="ixo.project.v1.WithdrawFundsDoc"></a>

### WithdrawFundsDoc
WithdrawFundsDoc contains details required to withdraw funds from a specific
project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| is_refund | [bool](#bool) |  |  |






<a name="ixo.project.v1.WithdrawalInfoDoc"></a>

### WithdrawalInfoDoc
WithdrawalInfoDoc contains details required to withdraw from a specific
project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |
| recipient_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ixo.project.v1.WithdrawalInfoDocs"></a>

### WithdrawalInfoDocs
WithdrawalInfoDocs contains a list of type WithdrawalInfoDoc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs_list | [WithdrawalInfoDoc](#ixo.project.v1.WithdrawalInfoDoc) | repeated |  |





 

 

 

 



<a name="ixo/project/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/project/v1/genesis.proto



<a name="ixo.project.v1.GenesisState"></a>

### GenesisState
GenesisState defines the project module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_docs | [ProjectDoc](#ixo.project.v1.ProjectDoc) | repeated |  |
| account_maps | [GenesisAccountMap](#ixo.project.v1.GenesisAccountMap) | repeated |  |
| withdrawals_infos | [WithdrawalInfoDocs](#ixo.project.v1.WithdrawalInfoDocs) | repeated |  |
| claims | [Claims](#ixo.project.v1.Claims) | repeated |  |
| params | [Params](#ixo.project.v1.Params) |  |  |





 

 

 

 



<a name="ixo/project/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/project/v1/query.proto



<a name="ixo.project.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="ixo.project.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.project.v1.Params) |  |  |






<a name="ixo.project.v1.QueryProjectAccountsRequest"></a>

### QueryProjectAccountsRequest
QueryProjectAccountsRequest is the request type for the Query/ProjectAccounts
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="ixo.project.v1.QueryProjectAccountsResponse"></a>

### QueryProjectAccountsResponse
QueryProjectAccountsResponse is the response type for the
Query/ProjectAccounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_map | [AccountMap](#ixo.project.v1.AccountMap) |  |  |






<a name="ixo.project.v1.QueryProjectDocRequest"></a>

### QueryProjectDocRequest
QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="ixo.project.v1.QueryProjectDocResponse"></a>

### QueryProjectDocResponse
QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_doc | [ProjectDoc](#ixo.project.v1.ProjectDoc) |  |  |






<a name="ixo.project.v1.QueryProjectTxRequest"></a>

### QueryProjectTxRequest
QueryProjectTxRequest is the request type for the Query/ProjectTx RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_did | [string](#string) |  |  |






<a name="ixo.project.v1.QueryProjectTxResponse"></a>

### QueryProjectTxResponse
QueryProjectTxResponse is the response type for the Query/ProjectTx RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txs | [WithdrawalInfoDocs](#ixo.project.v1.WithdrawalInfoDocs) |  |  |





 

 

 


<a name="ixo.project.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ProjectDoc | [QueryProjectDocRequest](#ixo.project.v1.QueryProjectDocRequest) | [QueryProjectDocResponse](#ixo.project.v1.QueryProjectDocResponse) | ProjectDoc queries info of a specific project. |
| ProjectAccounts | [QueryProjectAccountsRequest](#ixo.project.v1.QueryProjectAccountsRequest) | [QueryProjectAccountsResponse](#ixo.project.v1.QueryProjectAccountsResponse) | ProjectAccounts lists a specific project&#39;s accounts. |
| ProjectTx | [QueryProjectTxRequest](#ixo.project.v1.QueryProjectTxRequest) | [QueryProjectTxResponse](#ixo.project.v1.QueryProjectTxResponse) | ProjectTx lists a specific project&#39;s transactions. |
| Params | [QueryParamsRequest](#ixo.project.v1.QueryParamsRequest) | [QueryParamsResponse](#ixo.project.v1.QueryParamsResponse) | Params queries the paramaters of x/project module. |

 



<a name="ixo/project/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/project/v1/tx.proto



<a name="ixo.project.v1.MsgCreateAgent"></a>

### MsgCreateAgent
MsgCreateAgent defines a message for creating an agent on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateAgentDoc](#ixo.project.v1.CreateAgentDoc) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgCreateAgentResponse"></a>

### MsgCreateAgentResponse
MsgCreateAgentResponse defines the Msg/CreateAgent response type.






<a name="ixo.project.v1.MsgCreateClaim"></a>

### MsgCreateClaim
MsgCreateClaim defines a message for creating a claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateClaimDoc](#ixo.project.v1.CreateClaimDoc) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgCreateClaimResponse"></a>

### MsgCreateClaimResponse
MsgCreateClaimResponse defines the Msg/CreateClaim response type.






<a name="ixo.project.v1.MsgCreateEvaluation"></a>

### MsgCreateEvaluation
MsgCreateEvaluation defines a message for creating an evaluation for a
specific claim on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [CreateEvaluationDoc](#ixo.project.v1.CreateEvaluationDoc) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgCreateEvaluationResponse"></a>

### MsgCreateEvaluationResponse
MsgCreateEvaluationResponse defines the Msg/CreateEvaluation response type.






<a name="ixo.project.v1.MsgCreateProject"></a>

### MsgCreateProject
MsgCreateProject defines a message for creating a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| pub_key | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgCreateProjectResponse"></a>

### MsgCreateProjectResponse
MsgCreateProjectResponse defines the Msg/CreateProject response type.






<a name="ixo.project.v1.MsgUpdateAgent"></a>

### MsgUpdateAgent
MsgUpdateAgent defines a message for updating an agent on a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateAgentDoc](#ixo.project.v1.UpdateAgentDoc) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgUpdateAgentResponse"></a>

### MsgUpdateAgentResponse
MsgUpdateAgentResponse defines the Msg/UpdateAgent response type.






<a name="ixo.project.v1.MsgUpdateProjectDoc"></a>

### MsgUpdateProjectDoc
MsgUpdateProjectDoc defines a message for updating a project&#39;s data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgUpdateProjectDocResponse"></a>

### MsgUpdateProjectDocResponse
MsgUpdateProjectDocResponse defines the Msg/UpdateProjectDoc response type.






<a name="ixo.project.v1.MsgUpdateProjectStatus"></a>

### MsgUpdateProjectStatus
MsgUpdateProjectStatus defines a message for updating a project&#39;s current
status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_hash | [string](#string) |  |  |
| sender_did | [string](#string) |  |  |
| project_did | [string](#string) |  |  |
| data | [UpdateProjectStatusDoc](#ixo.project.v1.UpdateProjectStatusDoc) |  |  |
| project_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgUpdateProjectStatusResponse"></a>

### MsgUpdateProjectStatusResponse
MsgUpdateProjectStatusResponse defines the Msg/UpdateProjectStatus response
type.






<a name="ixo.project.v1.MsgWithdrawFunds"></a>

### MsgWithdrawFunds
MsgWithdrawFunds defines a message for project agents to withdraw their funds
from a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| data | [WithdrawFundsDoc](#ixo.project.v1.WithdrawFundsDoc) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.project.v1.MsgWithdrawFundsResponse"></a>

### MsgWithdrawFundsResponse
MsgWithdrawFundsResponse defines the Msg/WithdrawFunds response type.





 

 

 


<a name="ixo.project.v1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProject | [MsgCreateProject](#ixo.project.v1.MsgCreateProject) | [MsgCreateProjectResponse](#ixo.project.v1.MsgCreateProjectResponse) | CreateProject defines a method for creating a project. |
| UpdateProjectStatus | [MsgUpdateProjectStatus](#ixo.project.v1.MsgUpdateProjectStatus) | [MsgUpdateProjectStatusResponse](#ixo.project.v1.MsgUpdateProjectStatusResponse) | UpdateProjectStatus defines a method for updating a project&#39;s current status. |
| CreateAgent | [MsgCreateAgent](#ixo.project.v1.MsgCreateAgent) | [MsgCreateAgentResponse](#ixo.project.v1.MsgCreateAgentResponse) | CreateAgent defines a method for creating an agent on a project. |
| UpdateAgent | [MsgUpdateAgent](#ixo.project.v1.MsgUpdateAgent) | [MsgUpdateAgentResponse](#ixo.project.v1.MsgUpdateAgentResponse) | UpdateAgent defines a method for updating an agent on a project. |
| CreateClaim | [MsgCreateClaim](#ixo.project.v1.MsgCreateClaim) | [MsgCreateClaimResponse](#ixo.project.v1.MsgCreateClaimResponse) | CreateClaim defines a method for creating a claim on a project. |
| CreateEvaluation | [MsgCreateEvaluation](#ixo.project.v1.MsgCreateEvaluation) | [MsgCreateEvaluationResponse](#ixo.project.v1.MsgCreateEvaluationResponse) | CreateEvaluation defines a method for creating an evaluation for a specific claim on a project. |
| WithdrawFunds | [MsgWithdrawFunds](#ixo.project.v1.MsgWithdrawFunds) | [MsgWithdrawFundsResponse](#ixo.project.v1.MsgWithdrawFundsResponse) | WithdrawFunds defines a method for project agents to withdraw their funds from a project. |
| UpdateProjectDoc | [MsgUpdateProjectDoc](#ixo.project.v1.MsgUpdateProjectDoc) | [MsgUpdateProjectDocResponse](#ixo.project.v1.MsgUpdateProjectDocResponse) | UpdateProjectDoc defines a method for updating a project&#39;s data. |

 



<a name="ixo/token/v1beta1/token.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/token.proto



<a name="ixo.token.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo1155_contract_code | [uint64](#uint64) |  |  |






<a name="ixo.token.v1beta1.Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  | generated on token intiation through MsgSetupMinter |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| type | [string](#string) |  | type is the token type (eg ixo1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |
| supply | [string](#string) |  | how much has already been minted for this Token type, aka the supply |
| paused | [bool](#bool) |  | stop allowance of token minter temporarily |
| stopped | [bool](#bool) |  | stop allowance of token minter permanently |
| retired | [TokensRetired](#ixo.token.v1beta1.TokensRetired) | repeated | tokens that has been retired for this Token with specific name and contract address |
| cancelled | [TokensCancelled](#ixo.token.v1beta1.TokensCancelled) | repeated | tokens that has been cancelled for this Token with specific name and contract address |






<a name="ixo.token.v1beta1.TokenData"></a>

### TokenData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | media type value should always be &#34;application/json&#34;

credential link ***.ipfs |
| encrypted | [bool](#bool) |  |  |
| proof | [string](#string) |  |  |
| type | [string](#string) |  |  |
| id | [string](#string) |  | did of entity to map token to |






<a name="ixo.token.v1beta1.TokenProperties"></a>

### TokenProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| name | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="ixo.token.v1beta1.TokensCancelled"></a>

### TokensCancelled



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="ixo.token.v1beta1.TokensRetired"></a>

### TokensRetired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| jurisdiction | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/authz.proto



<a name="ixo.token.v1beta1.MintAuthorization"></a>

### MintAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| constraints | [MintConstraints](#ixo.token.v1beta1.MintConstraints) | repeated |  |






<a name="ixo.token.v1beta1.MintConstraints"></a>

### MintConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |





 

 

 

 



<a name="ixo/token/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/tx.proto



<a name="ixo.token.v1beta1.MintBatch"></a>

### MintBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to mint |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| token_data | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="ixo.token.v1beta1.MsgCancelToken"></a>

### MsgCancelToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="ixo.token.v1beta1.MsgCancelTokenResponse"></a>

### MsgCancelTokenResponse







<a name="ixo.token.v1beta1.MsgCreateToken"></a>

### MsgCreateToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| token_type | [string](#string) |  | type is the token type (eg ixo1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |






<a name="ixo.token.v1beta1.MsgCreateTokenResponse"></a>

### MsgCreateTokenResponse







<a name="ixo.token.v1beta1.MsgMintToken"></a>

### MsgMintToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| owner | [string](#string) |  | address of owner to mint for |
| mint_batch | [MintBatch](#ixo.token.v1beta1.MintBatch) | repeated |  |






<a name="ixo.token.v1beta1.MsgMintTokenResponse"></a>

### MsgMintTokenResponse







<a name="ixo.token.v1beta1.MsgPauseToken"></a>

### MsgPauseToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  | pause or unpause Token Minting allowance |






<a name="ixo.token.v1beta1.MsgPauseTokenResponse"></a>

### MsgPauseTokenResponse







<a name="ixo.token.v1beta1.MsgRetireToken"></a>

### MsgRetireToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| jurisdiction | [string](#string) |  | jurisdiction is the jurisdiction of the token owner. A jurisdiction has the format: &lt;country-code&gt;[-&lt;sub-national-code&gt;[ &lt;postal-code&gt;]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision. See the valid format for this below. |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="ixo.token.v1beta1.MsgRetireTokenResponse"></a>

### MsgRetireTokenResponse







<a name="ixo.token.v1beta1.MsgStopToken"></a>

### MsgStopToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |






<a name="ixo.token.v1beta1.MsgStopTokenResponse"></a>

### MsgStopTokenResponse







<a name="ixo.token.v1beta1.MsgTransferToken"></a>

### MsgTransferToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| recipient | [string](#string) |  | address of receiver |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | all tokens must be in same smart contract |






<a name="ixo.token.v1beta1.MsgTransferTokenResponse"></a>

### MsgTransferTokenResponse







<a name="ixo.token.v1beta1.TokenBatch"></a>

### TokenBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to transfer |





 

 

 


<a name="ixo.token.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateToken | [MsgCreateToken](#ixo.token.v1beta1.MsgCreateToken) | [MsgCreateTokenResponse](#ixo.token.v1beta1.MsgCreateTokenResponse) |  |
| MintToken | [MsgMintToken](#ixo.token.v1beta1.MsgMintToken) | [MsgMintTokenResponse](#ixo.token.v1beta1.MsgMintTokenResponse) |  |
| TransferToken | [MsgTransferToken](#ixo.token.v1beta1.MsgTransferToken) | [MsgTransferTokenResponse](#ixo.token.v1beta1.MsgTransferTokenResponse) |  |
| RetireToken | [MsgRetireToken](#ixo.token.v1beta1.MsgRetireToken) | [MsgRetireTokenResponse](#ixo.token.v1beta1.MsgRetireTokenResponse) |  |
| CancelToken | [MsgCancelToken](#ixo.token.v1beta1.MsgCancelToken) | [MsgCancelTokenResponse](#ixo.token.v1beta1.MsgCancelTokenResponse) |  |
| PauseToken | [MsgPauseToken](#ixo.token.v1beta1.MsgPauseToken) | [MsgPauseTokenResponse](#ixo.token.v1beta1.MsgPauseTokenResponse) |  |
| StopToken | [MsgStopToken](#ixo.token.v1beta1.MsgStopToken) | [MsgStopTokenResponse](#ixo.token.v1beta1.MsgStopTokenResponse) |  |

 



<a name="ixo/token/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/event.proto



<a name="ixo.token.v1beta1.TokenCancelledEvent"></a>

### TokenCancelledEvent
TokenCancelledEvent is an event triggered on a Token cancel execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenCreatedEvent"></a>

### TokenCreatedEvent
TokenCreatedEvent is an event triggered on a Token creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#ixo.token.v1beta1.Token) |  |  |






<a name="ixo.token.v1beta1.TokenMintedEvent"></a>

### TokenMintedEvent
TokenMintedEvent is an event triggered on a Token mint execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | the contract address of token contract being initialized |
| minter | [string](#string) |  | the token minter |
| owner | [string](#string) |  | the new tokens owner |
| amount | [string](#string) |  |  |
| tokenProperties | [TokenProperties](#ixo.token.v1beta1.TokenProperties) |  |  |






<a name="ixo.token.v1beta1.TokenPausedEvent"></a>

### TokenPausedEvent
TokenPausedEvent is an event triggered on a Token pause/unpause execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  |  |






<a name="ixo.token.v1beta1.TokenRetiredEvent"></a>

### TokenRetiredEvent
TokenRetiredEvent is an event triggered on a Token retire execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenStoppedEvent"></a>

### TokenStoppedEvent
TokenStoppedEvent is an event triggered on a Token stopped execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| stopped | [bool](#bool) |  |  |






<a name="ixo.token.v1beta1.TokenTransferredEvent"></a>

### TokenTransferredEvent
TokenTransferedEvent is an event triggered on a Token transfer execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the old token owner |
| recipient | [string](#string) |  | the new tokens owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenUpdatedEvent"></a>

### TokenUpdatedEvent
TokenUpdatedEvent is an event triggered on a Token update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#ixo.token.v1beta1.Token) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/genesis.proto



<a name="ixo.token.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.token.v1beta1.Params) |  |  |
| tokens | [Token](#ixo.token.v1beta1.Token) | repeated |  |
| token_properties | [TokenProperties](#ixo.token.v1beta1.TokenProperties) | repeated |  |





 

 

 

 



<a name="ixo/token/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/proposal.proto



<a name="ixo.token.v1beta1.SetTokenContractCodes"></a>

### SetTokenContractCodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo1155_contract_code | [uint64](#uint64) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/query.proto



<a name="ixo.token.v1beta1.QueryTokenDocRequest"></a>

### QueryTokenDocRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | minter address to get Token Doc for |
| contract_address | [string](#string) |  |  |






<a name="ixo.token.v1beta1.QueryTokenDocResponse"></a>

### QueryTokenDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokenDoc | [Token](#ixo.token.v1beta1.Token) |  |  |






<a name="ixo.token.v1beta1.QueryTokenListRequest"></a>

### QueryTokenListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |
| minter | [string](#string) |  | minter address to get list for |






<a name="ixo.token.v1beta1.QueryTokenListResponse"></a>

### QueryTokenListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |
| tokenDocs | [Token](#ixo.token.v1beta1.Token) | repeated |  |






<a name="ixo.token.v1beta1.QueryTokenMetadataRequest"></a>

### QueryTokenMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.token.v1beta1.QueryTokenMetadataResponse"></a>

### QueryTokenMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| decimals | [string](#string) |  |  |
| image | [string](#string) |  |  |
| index | [string](#string) |  |  |
| properties | [TokenMetadataProperties](#ixo.token.v1beta1.TokenMetadataProperties) |  |  |






<a name="ixo.token.v1beta1.TokenMetadataProperties"></a>

### TokenMetadataProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [string](#string) |  |  |
| collection | [string](#string) |  |  |
| cap | [string](#string) |  |  |
| linkedResources | [TokenData](#ixo.token.v1beta1.TokenData) | repeated |  |





 

 

 


<a name="ixo.token.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TokenList | [QueryTokenListRequest](#ixo.token.v1beta1.QueryTokenListRequest) | [QueryTokenListResponse](#ixo.token.v1beta1.QueryTokenListResponse) |  |
| TokenDoc | [QueryTokenDocRequest](#ixo.token.v1beta1.QueryTokenDocRequest) | [QueryTokenDocResponse](#ixo.token.v1beta1.QueryTokenDocResponse) |  |
| TokenMetadata | [QueryTokenMetadataRequest](#ixo.token.v1beta1.QueryTokenMetadataRequest) | [QueryTokenMetadataResponse](#ixo.token.v1beta1.QueryTokenMetadataResponse) |  |

 



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

