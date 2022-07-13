<!DOCTYPE html>

<html>
  <head>
    <title>Protocol Documentation</title>
    <meta charset="UTF-8">
    <link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Ubuntu:400,700,400italic"/>
    <style>
      body {
        width: 60em;
        margin: 1em auto;
        color: #222;
        font-family: "Ubuntu", sans-serif;
        padding-bottom: 4em;
      }

      h1 {
        font-weight: normal;
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
      }

      h2 {
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
        margin: 1.5em 0;
      }

      h3 {
        font-weight: normal;
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
      }

      a {
        text-decoration: none;
        color: #567e25;
      }

      table {
        width: 100%;
        font-size: 80%;
        border-collapse: collapse;
      }

      thead {
        font-weight: 700;
        background-color: #dcdcdc;
      }

      tbody tr:nth-child(even) {
        background-color: #fbfbfb;
      }

      td {
        border: 1px solid #ccc;
        padding: 0.5ex 2ex;
      }

      td p {
        text-indent: 1em;
        margin: 0;
      }

      td p:nth-child(1) {
        text-indent: 0;  
      }

       
      .field-table td:nth-child(1) {  
        width: 10em;
      }
      .field-table td:nth-child(2) {  
        width: 10em;
      }
      .field-table td:nth-child(3) {  
        width: 6em;
      }
      .field-table td:nth-child(4) {  
        width: auto;
      }

       
      .extension-table td:nth-child(1) {  
        width: 10em;
      }
      .extension-table td:nth-child(2) {  
        width: 10em;
      }
      .extension-table td:nth-child(3) {  
        width: 10em;
      }
      .extension-table td:nth-child(4) {  
        width: 5em;
      }
      .extension-table td:nth-child(5) {  
        width: auto;
      }

       
      .enum-table td:nth-child(1) {  
        width: 10em;
      }
      .enum-table td:nth-child(2) {  
        width: 10em;
      }
      .enum-table td:nth-child(3) {  
        width: auto;
      }

       
      .scalar-value-types-table tr {
        height: 3em;
      }

       
      #toc-container ul {
        list-style-type: none;
        padding-left: 1em;
        line-height: 180%;
        margin: 0;
      }
      #toc > li > a {
        font-weight: bold;
      }

       
      .file-heading {
        width: 100%;
        display: table;
        border-bottom: 1px solid #aaa;
        margin: 4em 0 1.5em 0;
      }
      .file-heading h2 {
        border: none;
        display: table-cell;
      }
      .file-heading a {
        text-align: right;
        display: table-cell;
      }

       
      .badge {
        width: 1.6em;
        height: 1.6em;
        display: inline-block;

        line-height: 1.6em;
        text-align: center;
        font-weight: bold;
        font-size: 60%;

        color: #89ba48;
        background-color: #dff0c8;

        margin: 0.5ex 1em 0.5ex -1em;
        border: 1px solid #fbfbfb;
        border-radius: 1ex;
      }
    </style>

    
    <link rel="stylesheet" type="text/css" href="stylesheet.css"/>
  </head>

  <body>

    <h1 id="title">Protocol Documentation</h1>

    <h2>Table of Contents</h2>

    <div id="toc-container">
      <ul id="toc">
        
          
          <li>
            <a href="#bonds%2fbonds.proto">bonds/bonds.proto</a>
            <ul>
              
                <li>
                  <a href="#bonds.BaseOrder"><span class="badge">M</span>BaseOrder</a>
                </li>
              
                <li>
                  <a href="#bonds.Batch"><span class="badge">M</span>Batch</a>
                </li>
              
                <li>
                  <a href="#bonds.Bond"><span class="badge">M</span>Bond</a>
                </li>
              
                <li>
                  <a href="#bonds.BondDetails"><span class="badge">M</span>BondDetails</a>
                </li>
              
                <li>
                  <a href="#bonds.BuyOrder"><span class="badge">M</span>BuyOrder</a>
                </li>
              
                <li>
                  <a href="#bonds.FunctionParam"><span class="badge">M</span>FunctionParam</a>
                </li>
              
                <li>
                  <a href="#bonds.Params"><span class="badge">M</span>Params</a>
                </li>
              
                <li>
                  <a href="#bonds.SellOrder"><span class="badge">M</span>SellOrder</a>
                </li>
              
                <li>
                  <a href="#bonds.SwapOrder"><span class="badge">M</span>SwapOrder</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bonds%2fgenesis.proto">bonds/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#bonds.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bonds%2fquery.proto">bonds/query.proto</a>
            <ul>
              
                <li>
                  <a href="#bonds.QueryAlphaMaximumsRequest"><span class="badge">M</span>QueryAlphaMaximumsRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryAlphaMaximumsResponse"><span class="badge">M</span>QueryAlphaMaximumsResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryAvailableReserveRequest"><span class="badge">M</span>QueryAvailableReserveRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryAvailableReserveResponse"><span class="badge">M</span>QueryAvailableReserveResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBatchRequest"><span class="badge">M</span>QueryBatchRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBatchResponse"><span class="badge">M</span>QueryBatchResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondRequest"><span class="badge">M</span>QueryBondRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondResponse"><span class="badge">M</span>QueryBondResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondsDetailedRequest"><span class="badge">M</span>QueryBondsDetailedRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondsDetailedResponse"><span class="badge">M</span>QueryBondsDetailedResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondsRequest"><span class="badge">M</span>QueryBondsRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBondsResponse"><span class="badge">M</span>QueryBondsResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBuyPriceRequest"><span class="badge">M</span>QueryBuyPriceRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryBuyPriceResponse"><span class="badge">M</span>QueryBuyPriceResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCurrentPriceRequest"><span class="badge">M</span>QueryCurrentPriceRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCurrentPriceResponse"><span class="badge">M</span>QueryCurrentPriceResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCurrentReserveRequest"><span class="badge">M</span>QueryCurrentReserveRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCurrentReserveResponse"><span class="badge">M</span>QueryCurrentReserveResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCustomPriceRequest"><span class="badge">M</span>QueryCustomPriceRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryCustomPriceResponse"><span class="badge">M</span>QueryCustomPriceResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryLastBatchRequest"><span class="badge">M</span>QueryLastBatchRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryLastBatchResponse"><span class="badge">M</span>QueryLastBatchResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryParamsRequest"><span class="badge">M</span>QueryParamsRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QueryParamsResponse"><span class="badge">M</span>QueryParamsResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QuerySellReturnRequest"><span class="badge">M</span>QuerySellReturnRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QuerySellReturnResponse"><span class="badge">M</span>QuerySellReturnResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.QuerySwapReturnRequest"><span class="badge">M</span>QuerySwapReturnRequest</a>
                </li>
              
                <li>
                  <a href="#bonds.QuerySwapReturnResponse"><span class="badge">M</span>QuerySwapReturnResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#bonds.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bonds%2ftx.proto">bonds/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#bonds.MsgBuy"><span class="badge">M</span>MsgBuy</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgBuyResponse"><span class="badge">M</span>MsgBuyResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgCreateBond"><span class="badge">M</span>MsgCreateBond</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgCreateBondResponse"><span class="badge">M</span>MsgCreateBondResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgEditBond"><span class="badge">M</span>MsgEditBond</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgEditBondResponse"><span class="badge">M</span>MsgEditBondResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgMakeOutcomePayment"><span class="badge">M</span>MsgMakeOutcomePayment</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgMakeOutcomePaymentResponse"><span class="badge">M</span>MsgMakeOutcomePaymentResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSell"><span class="badge">M</span>MsgSell</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSellResponse"><span class="badge">M</span>MsgSellResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSetNextAlpha"><span class="badge">M</span>MsgSetNextAlpha</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSetNextAlphaResponse"><span class="badge">M</span>MsgSetNextAlphaResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSwap"><span class="badge">M</span>MsgSwap</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgSwapResponse"><span class="badge">M</span>MsgSwapResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgUpdateBondState"><span class="badge">M</span>MsgUpdateBondState</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgUpdateBondStateResponse"><span class="badge">M</span>MsgUpdateBondStateResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgWithdrawReserve"><span class="badge">M</span>MsgWithdrawReserve</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgWithdrawReserveResponse"><span class="badge">M</span>MsgWithdrawReserveResponse</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgWithdrawShare"><span class="badge">M</span>MsgWithdrawShare</a>
                </li>
              
                <li>
                  <a href="#bonds.MsgWithdrawShareResponse"><span class="badge">M</span>MsgWithdrawShareResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#bonds.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#did%2fdid.proto">did/did.proto</a>
            <ul>
              
                <li>
                  <a href="#did.Claim"><span class="badge">M</span>Claim</a>
                </li>
              
                <li>
                  <a href="#did.DidCredential"><span class="badge">M</span>DidCredential</a>
                </li>
              
                <li>
                  <a href="#did.IxoDid"><span class="badge">M</span>IxoDid</a>
                </li>
              
                <li>
                  <a href="#did.Secret"><span class="badge">M</span>Secret</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#did%2fdiddoc.proto">did/diddoc.proto</a>
            <ul>
              
                <li>
                  <a href="#did.BaseDidDoc"><span class="badge">M</span>BaseDidDoc</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#did%2fgenesis.proto">did/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#did.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#did%2fquery.proto">did/query.proto</a>
            <ul>
              
                <li>
                  <a href="#did.QueryAddressFromBase58EncodedPubkeyRequest"><span class="badge">M</span>QueryAddressFromBase58EncodedPubkeyRequest</a>
                </li>
              
                <li>
                  <a href="#did.QueryAddressFromBase58EncodedPubkeyResponse"><span class="badge">M</span>QueryAddressFromBase58EncodedPubkeyResponse</a>
                </li>
              
                <li>
                  <a href="#did.QueryAddressFromDidRequest"><span class="badge">M</span>QueryAddressFromDidRequest</a>
                </li>
              
                <li>
                  <a href="#did.QueryAddressFromDidResponse"><span class="badge">M</span>QueryAddressFromDidResponse</a>
                </li>
              
                <li>
                  <a href="#did.QueryAllDidDocsRequest"><span class="badge">M</span>QueryAllDidDocsRequest</a>
                </li>
              
                <li>
                  <a href="#did.QueryAllDidDocsResponse"><span class="badge">M</span>QueryAllDidDocsResponse</a>
                </li>
              
                <li>
                  <a href="#did.QueryAllDidsRequest"><span class="badge">M</span>QueryAllDidsRequest</a>
                </li>
              
                <li>
                  <a href="#did.QueryAllDidsResponse"><span class="badge">M</span>QueryAllDidsResponse</a>
                </li>
              
                <li>
                  <a href="#did.QueryDidDocRequest"><span class="badge">M</span>QueryDidDocRequest</a>
                </li>
              
                <li>
                  <a href="#did.QueryDidDocResponse"><span class="badge">M</span>QueryDidDocResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#did.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#did%2ftx.proto">did/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#did.MsgAddCredential"><span class="badge">M</span>MsgAddCredential</a>
                </li>
              
                <li>
                  <a href="#did.MsgAddCredentialResponse"><span class="badge">M</span>MsgAddCredentialResponse</a>
                </li>
              
                <li>
                  <a href="#did.MsgAddDid"><span class="badge">M</span>MsgAddDid</a>
                </li>
              
                <li>
                  <a href="#did.MsgAddDidResponse"><span class="badge">M</span>MsgAddDidResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#did.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#payments%2fpayments.proto">payments/payments.proto</a>
            <ul>
              
                <li>
                  <a href="#payments.BlockPeriod"><span class="badge">M</span>BlockPeriod</a>
                </li>
              
                <li>
                  <a href="#payments.Discount"><span class="badge">M</span>Discount</a>
                </li>
              
                <li>
                  <a href="#payments.DistributionShare"><span class="badge">M</span>DistributionShare</a>
                </li>
              
                <li>
                  <a href="#payments.PaymentContract"><span class="badge">M</span>PaymentContract</a>
                </li>
              
                <li>
                  <a href="#payments.PaymentTemplate"><span class="badge">M</span>PaymentTemplate</a>
                </li>
              
                <li>
                  <a href="#payments.Subscription"><span class="badge">M</span>Subscription</a>
                </li>
              
                <li>
                  <a href="#payments.TestPeriod"><span class="badge">M</span>TestPeriod</a>
                </li>
              
                <li>
                  <a href="#payments.TimePeriod"><span class="badge">M</span>TimePeriod</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#payments%2fgenesis.proto">payments/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#payments.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#payments%2fquery.proto">payments/query.proto</a>
            <ul>
              
                <li>
                  <a href="#payments.QueryPaymentContractRequest"><span class="badge">M</span>QueryPaymentContractRequest</a>
                </li>
              
                <li>
                  <a href="#payments.QueryPaymentContractResponse"><span class="badge">M</span>QueryPaymentContractResponse</a>
                </li>
              
                <li>
                  <a href="#payments.QueryPaymentContractsByIdPrefixRequest"><span class="badge">M</span>QueryPaymentContractsByIdPrefixRequest</a>
                </li>
              
                <li>
                  <a href="#payments.QueryPaymentContractsByIdPrefixResponse"><span class="badge">M</span>QueryPaymentContractsByIdPrefixResponse</a>
                </li>
              
                <li>
                  <a href="#payments.QueryPaymentTemplateRequest"><span class="badge">M</span>QueryPaymentTemplateRequest</a>
                </li>
              
                <li>
                  <a href="#payments.QueryPaymentTemplateResponse"><span class="badge">M</span>QueryPaymentTemplateResponse</a>
                </li>
              
                <li>
                  <a href="#payments.QuerySubscriptionRequest"><span class="badge">M</span>QuerySubscriptionRequest</a>
                </li>
              
                <li>
                  <a href="#payments.QuerySubscriptionResponse"><span class="badge">M</span>QuerySubscriptionResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#payments.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#payments%2ftx.proto">payments/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#payments.MsgCreatePaymentContract"><span class="badge">M</span>MsgCreatePaymentContract</a>
                </li>
              
                <li>
                  <a href="#payments.MsgCreatePaymentContractResponse"><span class="badge">M</span>MsgCreatePaymentContractResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgCreatePaymentTemplate"><span class="badge">M</span>MsgCreatePaymentTemplate</a>
                </li>
              
                <li>
                  <a href="#payments.MsgCreatePaymentTemplateResponse"><span class="badge">M</span>MsgCreatePaymentTemplateResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgCreateSubscription"><span class="badge">M</span>MsgCreateSubscription</a>
                </li>
              
                <li>
                  <a href="#payments.MsgCreateSubscriptionResponse"><span class="badge">M</span>MsgCreateSubscriptionResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgEffectPayment"><span class="badge">M</span>MsgEffectPayment</a>
                </li>
              
                <li>
                  <a href="#payments.MsgEffectPaymentResponse"><span class="badge">M</span>MsgEffectPaymentResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgGrantDiscount"><span class="badge">M</span>MsgGrantDiscount</a>
                </li>
              
                <li>
                  <a href="#payments.MsgGrantDiscountResponse"><span class="badge">M</span>MsgGrantDiscountResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgRevokeDiscount"><span class="badge">M</span>MsgRevokeDiscount</a>
                </li>
              
                <li>
                  <a href="#payments.MsgRevokeDiscountResponse"><span class="badge">M</span>MsgRevokeDiscountResponse</a>
                </li>
              
                <li>
                  <a href="#payments.MsgSetPaymentContractAuthorisation"><span class="badge">M</span>MsgSetPaymentContractAuthorisation</a>
                </li>
              
                <li>
                  <a href="#payments.MsgSetPaymentContractAuthorisationResponse"><span class="badge">M</span>MsgSetPaymentContractAuthorisationResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#payments.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#project%2fproject.proto">project/project.proto</a>
            <ul>
              
                <li>
                  <a href="#project.AccountMap"><span class="badge">M</span>AccountMap</a>
                </li>
              
                <li>
                  <a href="#project.AccountMap.MapEntry"><span class="badge">M</span>AccountMap.MapEntry</a>
                </li>
              
                <li>
                  <a href="#project.Claim"><span class="badge">M</span>Claim</a>
                </li>
              
                <li>
                  <a href="#project.Claims"><span class="badge">M</span>Claims</a>
                </li>
              
                <li>
                  <a href="#project.CreateAgentDoc"><span class="badge">M</span>CreateAgentDoc</a>
                </li>
              
                <li>
                  <a href="#project.CreateClaimDoc"><span class="badge">M</span>CreateClaimDoc</a>
                </li>
              
                <li>
                  <a href="#project.CreateEvaluationDoc"><span class="badge">M</span>CreateEvaluationDoc</a>
                </li>
              
                <li>
                  <a href="#project.GenesisAccountMap"><span class="badge">M</span>GenesisAccountMap</a>
                </li>
              
                <li>
                  <a href="#project.GenesisAccountMap.MapEntry"><span class="badge">M</span>GenesisAccountMap.MapEntry</a>
                </li>
              
                <li>
                  <a href="#project.Params"><span class="badge">M</span>Params</a>
                </li>
              
                <li>
                  <a href="#project.ProjectDoc"><span class="badge">M</span>ProjectDoc</a>
                </li>
              
                <li>
                  <a href="#project.UpdateAgentDoc"><span class="badge">M</span>UpdateAgentDoc</a>
                </li>
              
                <li>
                  <a href="#project.UpdateProjectStatusDoc"><span class="badge">M</span>UpdateProjectStatusDoc</a>
                </li>
              
                <li>
                  <a href="#project.WithdrawFundsDoc"><span class="badge">M</span>WithdrawFundsDoc</a>
                </li>
              
                <li>
                  <a href="#project.WithdrawalInfoDoc"><span class="badge">M</span>WithdrawalInfoDoc</a>
                </li>
              
                <li>
                  <a href="#project.WithdrawalInfoDocs"><span class="badge">M</span>WithdrawalInfoDocs</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#project%2fgenesis.proto">project/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#project.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#project%2fquery.proto">project/query.proto</a>
            <ul>
              
                <li>
                  <a href="#project.QueryParamsRequest"><span class="badge">M</span>QueryParamsRequest</a>
                </li>
              
                <li>
                  <a href="#project.QueryParamsResponse"><span class="badge">M</span>QueryParamsResponse</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectAccountsRequest"><span class="badge">M</span>QueryProjectAccountsRequest</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectAccountsResponse"><span class="badge">M</span>QueryProjectAccountsResponse</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectDocRequest"><span class="badge">M</span>QueryProjectDocRequest</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectDocResponse"><span class="badge">M</span>QueryProjectDocResponse</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectTxRequest"><span class="badge">M</span>QueryProjectTxRequest</a>
                </li>
              
                <li>
                  <a href="#project.QueryProjectTxResponse"><span class="badge">M</span>QueryProjectTxResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#project.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#project%2ftx.proto">project/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#project.MsgCreateAgent"><span class="badge">M</span>MsgCreateAgent</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateAgentResponse"><span class="badge">M</span>MsgCreateAgentResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateClaim"><span class="badge">M</span>MsgCreateClaim</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateClaimResponse"><span class="badge">M</span>MsgCreateClaimResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateEvaluation"><span class="badge">M</span>MsgCreateEvaluation</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateEvaluationResponse"><span class="badge">M</span>MsgCreateEvaluationResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateProject"><span class="badge">M</span>MsgCreateProject</a>
                </li>
              
                <li>
                  <a href="#project.MsgCreateProjectResponse"><span class="badge">M</span>MsgCreateProjectResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateAgent"><span class="badge">M</span>MsgUpdateAgent</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateAgentResponse"><span class="badge">M</span>MsgUpdateAgentResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateProjectDoc"><span class="badge">M</span>MsgUpdateProjectDoc</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateProjectDocResponse"><span class="badge">M</span>MsgUpdateProjectDocResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateProjectStatus"><span class="badge">M</span>MsgUpdateProjectStatus</a>
                </li>
              
                <li>
                  <a href="#project.MsgUpdateProjectStatusResponse"><span class="badge">M</span>MsgUpdateProjectStatusResponse</a>
                </li>
              
                <li>
                  <a href="#project.MsgWithdrawFunds"><span class="badge">M</span>MsgWithdrawFunds</a>
                </li>
              
                <li>
                  <a href="#project.MsgWithdrawFundsResponse"><span class="badge">M</span>MsgWithdrawFundsResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#project.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
        <li><a href="#scalar-value-types">Scalar Value Types</a></li>
      </ul>
    </div>

    
      
      <div class="file-heading">
        <h2 id="bonds/bonds.proto">bonds/bonds.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bonds.BaseOrder">BaseOrder</h3>
        <p>BaseOrder defines a base order type. It contains all the necessary fields for specifying</p><p>the general details about a buy, sell, or swap order.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>account_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>cancelled</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>cancel_reason</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.Batch">Batch</h3>
        <p>Batch holds a collection of outstanding buy, sell, and swap orders on a particular bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>blocks_remaining</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>next_public_alpha</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_buy_amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_sell_amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>buy_prices</td>
                  <td><a href="#cosmos.base.v1beta1.DecCoin">cosmos.base.v1beta1.DecCoin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sell_prices</td>
                  <td><a href="#cosmos.base.v1beta1.DecCoin">cosmos.base.v1beta1.DecCoin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>buys</td>
                  <td><a href="#bonds.BuyOrder">BuyOrder</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sells</td>
                  <td><a href="#bonds.SellOrder">SellOrder</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>swaps</td>
                  <td><a href="#bonds.SwapOrder">SwapOrder</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.Bond">Bond</h3>
        <p>Bond defines a token bonding curve type with all of its parameters.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>token</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>name</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>description</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>creator_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>controller_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>function_type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>function_parameters</td>
                  <td><a href="#bonds.FunctionParam">FunctionParam</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reserve_tokens</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>tx_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>exit_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>fee_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reserve_withdrawal_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>order_quantity_limits</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_rate</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_margin_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>current_supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>current_reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>available_reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>current_outcome_payment_reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>allow_sells</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>allow_reserve_withdrawals</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>alpha_bond</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>batch_blocks</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>outcome_payment</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>state</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.BondDetails">BondDetails</h3>
        <p>BondDetails contains details about the current state of a given bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>spot_price</td>
                  <td><a href="#cosmos.base.v1beta1.DecCoin">cosmos.base.v1beta1.DecCoin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.BuyOrder">BuyOrder</h3>
        <p>BuyOrder defines a type for submitting a buy order on a bond, together with the maximum</p><p>amount of reserve tokens the buyer is willing to pay.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_prices</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.FunctionParam">FunctionParam</h3>
        <p>FunctionParam is a key-value pair used for specifying a specific bond parameter.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>param</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>value</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.Params">Params</h3>
        <p>Params defines the parameters for the bonds module.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reserved_bond_tokens</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.SellOrder">SellOrder</h3>
        <p>SellOrder defines a type for submitting a sell order on a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.SwapOrder">SwapOrder</h3>
        <p>SwapOrder defines a type for submitting a swap order between two tokens on a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>to_token</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bonds/genesis.proto">bonds/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bonds.GenesisState">GenesisState</h3>
        <p>GenesisState defines the bonds module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bonds</td>
                  <td><a href="#bonds.Bond">Bond</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>batches</td>
                  <td><a href="#bonds.Batch">Batch</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
                  <td><a href="#bonds.Params">Params</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bonds/query.proto">bonds/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bonds.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</h3>
        <p>QueryAlphaMaximumsRequest is the request type for the Query/AlphaMaximums RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</h3>
        <p>QueryAlphaMaximumsResponse is the response type for the Query/AlphaMaximums RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>max_system_alpha_increase</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_system_alpha</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryAvailableReserveRequest">QueryAvailableReserveRequest</h3>
        <p>QueryAvailableReserveRequest is the request type for the Query/AvailableReserve RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryAvailableReserveResponse">QueryAvailableReserveResponse</h3>
        <p>QueryAvailableReserveResponse is the response type for the Query/AvailableReserve RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>available_reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBatchRequest">QueryBatchRequest</h3>
        <p>QueryBatchRequest is the request type for the Query/Batch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBatchResponse">QueryBatchResponse</h3>
        <p>QueryBatchResponse is the response type for the Query/Batch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>batch</td>
                  <td><a href="#bonds.Batch">Batch</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBondRequest">QueryBondRequest</h3>
        <p>QueryBondRequest is the request type for the Query/Bond RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBondResponse">QueryBondResponse</h3>
        <p>QueryBondResponse is the response type for the Query/Bond RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond</td>
                  <td><a href="#bonds.Bond">Bond</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBondsDetailedRequest">QueryBondsDetailedRequest</h3>
        <p>QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC method.</p>

        

        
      
        <h3 id="bonds.QueryBondsDetailedResponse">QueryBondsDetailedResponse</h3>
        <p>QueryBondsDetailedResponse is the response type for the Query/BondsDetailed RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bonds_detailed</td>
                  <td><a href="#bonds.BondDetails">BondDetails</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBondsRequest">QueryBondsRequest</h3>
        <p>QueryBondsRequest is the request type for the Query/Bonds RPC method.</p>

        

        
      
        <h3 id="bonds.QueryBondsResponse">QueryBondsResponse</h3>
        <p>QueryBondsResponse is the response type for the Query/Bonds RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bonds</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBuyPriceRequest">QueryBuyPriceRequest</h3>
        <p>QueryCustomPriceRequest is the request type for the Query/BuyPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryBuyPriceResponse">QueryBuyPriceResponse</h3>
        <p>QueryCustomPriceResponse is the response type for the Query/BuyPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>adjusted_supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prices</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>tx_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_prices</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCurrentPriceRequest">QueryCurrentPriceRequest</h3>
        <p>QueryCurrentPriceRequest is the request type for the Query/CurrentPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCurrentPriceResponse">QueryCurrentPriceResponse</h3>
        <p>QueryCurrentPriceResponse is the response type for the Query/CurrentPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>current_price</td>
                  <td><a href="#cosmos.base.v1beta1.DecCoin">cosmos.base.v1beta1.DecCoin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCurrentReserveRequest">QueryCurrentReserveRequest</h3>
        <p>QueryCurrentReserveRequest is the request type for the Query/CurrentReserve RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCurrentReserveResponse">QueryCurrentReserveResponse</h3>
        <p>QueryCurrentReserveResponse is the response type for the Query/CurrentReserve RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>current_reserve</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCustomPriceRequest">QueryCustomPriceRequest</h3>
        <p>QueryCustomPriceRequest is the request type for the Query/CustomPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryCustomPriceResponse">QueryCustomPriceResponse</h3>
        <p>QueryCustomPriceResponse is the response type for the Query/CustomPrice RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>price</td>
                  <td><a href="#cosmos.base.v1beta1.DecCoin">cosmos.base.v1beta1.DecCoin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryLastBatchRequest">QueryLastBatchRequest</h3>
        <p>QueryLastBatchRequest is the request type for the Query/LastBatch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryLastBatchResponse">QueryLastBatchResponse</h3>
        <p>QueryLastBatchResponse is the response type for the Query/LastBatch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>last_batch</td>
                  <td><a href="#bonds.Batch">Batch</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QueryParamsRequest">QueryParamsRequest</h3>
        <p>QueryParamsRequest is the request type for the Query/Params RPC method.</p>

        

        
      
        <h3 id="bonds.QueryParamsResponse">QueryParamsResponse</h3>
        <p>QueryParamsResponse is the response type for the Query/Params RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
                  <td><a href="#bonds.Params">Params</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QuerySellReturnRequest">QuerySellReturnRequest</h3>
        <p>QuerySellReturnRequest is the request type for the Query/SellReturn RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QuerySellReturnResponse">QuerySellReturnResponse</h3>
        <p>QuerySellReturnResponse is the response type for the Query/SellReturn RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>adjusted_supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>returns</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>tx_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>exit_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_returns</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QuerySwapReturnRequest">QuerySwapReturnRequest</h3>
        <p>QuerySwapReturnRequest is the request type for the Query/SwapReturn RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>from_token_with_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>to_token</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.QuerySwapReturnResponse">QuerySwapReturnResponse</h3>
        <p>QuerySwapReturnResponse is the response type for the Query/SwapReturn RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>total_returns</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>total_fees</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="bonds.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>Bonds</td>
                <td><a href="#bonds.QueryBondsRequest">QueryBondsRequest</a></td>
                <td><a href="#bonds.QueryBondsResponse">QueryBondsResponse</a></td>
                <td><p>Bonds returns all existing bonds.</p></td>
              </tr>
            
              <tr>
                <td>BondsDetailed</td>
                <td><a href="#bonds.QueryBondsDetailedRequest">QueryBondsDetailedRequest</a></td>
                <td><a href="#bonds.QueryBondsDetailedResponse">QueryBondsDetailedResponse</a></td>
                <td><p>BondsDetailed returns a list of all existing bonds with some details about their current state.</p></td>
              </tr>
            
              <tr>
                <td>Params</td>
                <td><a href="#bonds.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#bonds.QueryParamsResponse">QueryParamsResponse</a></td>
                <td><p>Params queries the paramaters of x/bonds module.</p></td>
              </tr>
            
              <tr>
                <td>Bond</td>
                <td><a href="#bonds.QueryBondRequest">QueryBondRequest</a></td>
                <td><a href="#bonds.QueryBondResponse">QueryBondResponse</a></td>
                <td><p>Bond queries info of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>Batch</td>
                <td><a href="#bonds.QueryBatchRequest">QueryBatchRequest</a></td>
                <td><a href="#bonds.QueryBatchResponse">QueryBatchResponse</a></td>
                <td><p>Batch queries info of a specific bond&#39;s current batch.</p></td>
              </tr>
            
              <tr>
                <td>LastBatch</td>
                <td><a href="#bonds.QueryLastBatchRequest">QueryLastBatchRequest</a></td>
                <td><a href="#bonds.QueryLastBatchResponse">QueryLastBatchResponse</a></td>
                <td><p>LastBatch queries info of a specific bond&#39;s last batch.</p></td>
              </tr>
            
              <tr>
                <td>CurrentPrice</td>
                <td><a href="#bonds.QueryCurrentPriceRequest">QueryCurrentPriceRequest</a></td>
                <td><a href="#bonds.QueryCurrentPriceResponse">QueryCurrentPriceResponse</a></td>
                <td><p>CurrentPrice queries the current price/s of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>CurrentReserve</td>
                <td><a href="#bonds.QueryCurrentReserveRequest">QueryCurrentReserveRequest</a></td>
                <td><a href="#bonds.QueryCurrentReserveResponse">QueryCurrentReserveResponse</a></td>
                <td><p>CurrentReserve queries the current balance/s of the reserve pool for a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>AvailableReserve</td>
                <td><a href="#bonds.QueryAvailableReserveRequest">QueryAvailableReserveRequest</a></td>
                <td><a href="#bonds.QueryAvailableReserveResponse">QueryAvailableReserveResponse</a></td>
                <td><p>AvailableReserve queries current available balance/s of the reserve pool for a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>CustomPrice</td>
                <td><a href="#bonds.QueryCustomPriceRequest">QueryCustomPriceRequest</a></td>
                <td><a href="#bonds.QueryCustomPriceResponse">QueryCustomPriceResponse</a></td>
                <td><p>CustomPrice queries price/s of a specific bond at a specific supply.</p></td>
              </tr>
            
              <tr>
                <td>BuyPrice</td>
                <td><a href="#bonds.QueryBuyPriceRequest">QueryBuyPriceRequest</a></td>
                <td><a href="#bonds.QueryBuyPriceResponse">QueryBuyPriceResponse</a></td>
                <td><p>BuyPrice queries price/s of buying an amount of tokens from a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>SellReturn</td>
                <td><a href="#bonds.QuerySellReturnRequest">QuerySellReturnRequest</a></td>
                <td><a href="#bonds.QuerySellReturnResponse">QuerySellReturnResponse</a></td>
                <td><p>SellReturn queries return/s on selling an amount of tokens of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>SwapReturn</td>
                <td><a href="#bonds.QuerySwapReturnRequest">QuerySwapReturnRequest</a></td>
                <td><a href="#bonds.QuerySwapReturnResponse">QuerySwapReturnResponse</a></td>
                <td><p>SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>AlphaMaximums</td>
                <td><a href="#bonds.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</a></td>
                <td><a href="#bonds.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</a></td>
                <td><p>AlphaMaximums queries alpha maximums for a specific augmented bonding curve.</p></td>
              </tr>
            
          </tbody>
        </table>

        
          
          
          <h4>Methods with HTTP bindings</h4>
          <table>
            <thead>
              <tr>
                <td>Method Name</td>
                <td>Method</td>
                <td>Pattern</td>
                <td>Body</td>
              </tr>
            </thead>
            <tbody>
            
              
              
              <tr>
                <td>Bonds</td>
                <td>GET</td>
                <td>/ixo/bonds</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>BondsDetailed</td>
                <td>GET</td>
                <td>/ixo/bonds_detailed</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>Params</td>
                <td>GET</td>
                <td>/ixo/bonds/params</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>Bond</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>Batch</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/batch</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>LastBatch</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/last_batch</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>CurrentPrice</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/current_price</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>CurrentReserve</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/current_reserve</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AvailableReserve</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/available_reserve</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>CustomPrice</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/price/{bond_amount}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>BuyPrice</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/buy_price/{bond_amount}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>SellReturn</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/sell_return/{bond_amount}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>SwapReturn</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/swap_return/{from_token_with_amount}/{to_token}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AlphaMaximums</td>
                <td>GET</td>
                <td>/ixo/bonds/{bond_did}/alpha_maximums</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="bonds/tx.proto">bonds/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bonds.MsgBuy">MsgBuy</h3>
        <p>MsgBuy defines a message for buying from a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>buyer_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_prices</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgBuyResponse">MsgBuyResponse</h3>
        <p>MsgBuyResponse defines the Msg/Buy response type.</p>

        

        
      
        <h3 id="bonds.MsgCreateBond">MsgCreateBond</h3>
        <p>MsgCreateBond defines a message for creating a new bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>token</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>name</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>description</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>function_type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>function_parameters</td>
                  <td><a href="#bonds.FunctionParam">FunctionParam</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>creator_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>controller_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reserve_tokens</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>tx_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>exit_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>fee_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reserve_withdrawal_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_supply</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>order_quantity_limits</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_rate</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_margin_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>allow_sells</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>allow_reserve_withdrawals</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>alpha_bond</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>batch_blocks</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>outcome_payment</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgCreateBondResponse">MsgCreateBondResponse</h3>
        <p>MsgCreateBondResponse defines the Msg/CreateBond response type.</p>

        

        
      
        <h3 id="bonds.MsgEditBond">MsgEditBond</h3>
        <p>MsgEditBond defines a message for editing an existing bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>name</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>description</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>order_quantity_limits</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_rate</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sanity_margin_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>editor_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgEditBondResponse">MsgEditBondResponse</h3>
        <p>MsgEditBondResponse defines the Msg/EditBond response type.</p>

        

        
      
        <h3 id="bonds.MsgMakeOutcomePayment">MsgMakeOutcomePayment</h3>
        <p>MsgMakeOutcomePayment defines a message for making an outcome payment to a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</h3>
        <p>MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response type.</p>

        

        
      
        <h3 id="bonds.MsgSell">MsgSell</h3>
        <p>MsgSell defines a message for selling from a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>seller_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgSellResponse">MsgSellResponse</h3>
        <p>MsgSellResponse defines the Msg/Sell response type.</p>

        

        
      
        <h3 id="bonds.MsgSetNextAlpha">MsgSetNextAlpha</h3>
        <p>MsgSetNextAlpha defines a message for editing a bond's alpha parameter.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>alpha</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>editor_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</h3>
        <p></p>

        

        
      
        <h3 id="bonds.MsgSwap">MsgSwap</h3>
        <p>MsgSwap defines a message for swapping from one reserve bond token to another.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>swapper_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>from</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>to_token</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgSwapResponse">MsgSwapResponse</h3>
        <p>MsgSwapResponse defines the Msg/Swap response type.</p>

        

        
      
        <h3 id="bonds.MsgUpdateBondState">MsgUpdateBondState</h3>
        <p>MsgUpdateBondState defines a message for updating a bond's current state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>state</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>editor_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</h3>
        <p>MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.</p>

        

        
      
        <h3 id="bonds.MsgWithdrawReserve">MsgWithdrawReserve</h3>
        <p>MsgWithdrawReserve defines a message for withdrawing reserve from a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>withdrawer_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</h3>
        <p>MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.</p>

        

        
      
        <h3 id="bonds.MsgWithdrawShare">MsgWithdrawShare</h3>
        <p>MsgWithdrawShare defines a message for withdrawing a share from a bond that is in the SETTLE stage.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>recipient_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>bond_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bonds.MsgWithdrawShareResponse">MsgWithdrawShareResponse</h3>
        <p>MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.</p>

        

        
      

      

      

      
        <h3 id="bonds.Msg">Msg</h3>
        <p>Msg defines the bonds Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateBond</td>
                <td><a href="#bonds.MsgCreateBond">MsgCreateBond</a></td>
                <td><a href="#bonds.MsgCreateBondResponse">MsgCreateBondResponse</a></td>
                <td><p>CreateBond defines a method for creating a bond.</p></td>
              </tr>
            
              <tr>
                <td>EditBond</td>
                <td><a href="#bonds.MsgEditBond">MsgEditBond</a></td>
                <td><a href="#bonds.MsgEditBondResponse">MsgEditBondResponse</a></td>
                <td><p>EditBond defines a method for editing a bond.</p></td>
              </tr>
            
              <tr>
                <td>SetNextAlpha</td>
                <td><a href="#bonds.MsgSetNextAlpha">MsgSetNextAlpha</a></td>
                <td><a href="#bonds.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</a></td>
                <td><p>SetNextAlpha defines a method for editing a bond&#39;s alpha parameter.</p></td>
              </tr>
            
              <tr>
                <td>UpdateBondState</td>
                <td><a href="#bonds.MsgUpdateBondState">MsgUpdateBondState</a></td>
                <td><a href="#bonds.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</a></td>
                <td><p>UpdateBondState defines a method for updating a bond&#39;s current state.</p></td>
              </tr>
            
              <tr>
                <td>Buy</td>
                <td><a href="#bonds.MsgBuy">MsgBuy</a></td>
                <td><a href="#bonds.MsgBuyResponse">MsgBuyResponse</a></td>
                <td><p>Buy defines a method for buying from a bond.</p></td>
              </tr>
            
              <tr>
                <td>Sell</td>
                <td><a href="#bonds.MsgSell">MsgSell</a></td>
                <td><a href="#bonds.MsgSellResponse">MsgSellResponse</a></td>
                <td><p>Sell defines a method for selling from a bond.</p></td>
              </tr>
            
              <tr>
                <td>Swap</td>
                <td><a href="#bonds.MsgSwap">MsgSwap</a></td>
                <td><a href="#bonds.MsgSwapResponse">MsgSwapResponse</a></td>
                <td><p>Swap defines a method for swapping from one reserve bond token to another.</p></td>
              </tr>
            
              <tr>
                <td>MakeOutcomePayment</td>
                <td><a href="#bonds.MsgMakeOutcomePayment">MsgMakeOutcomePayment</a></td>
                <td><a href="#bonds.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</a></td>
                <td><p>MakeOutcomePayment defines a method for making an outcome payment to a bond.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawShare</td>
                <td><a href="#bonds.MsgWithdrawShare">MsgWithdrawShare</a></td>
                <td><a href="#bonds.MsgWithdrawShareResponse">MsgWithdrawShareResponse</a></td>
                <td><p>WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawReserve</td>
                <td><a href="#bonds.MsgWithdrawReserve">MsgWithdrawReserve</a></td>
                <td><a href="#bonds.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</a></td>
                <td><p>WithdrawReserve defines a method for withdrawing reserve from a bond.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
        <h2 id="did/did.proto">did/did.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="did.Claim">Claim</h3>
        <p>The claim section of a credential, indicating if the DID is KYC validated</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>KYC_validated</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.DidCredential">DidCredential</h3>
        <p>Digital identity credential issued to an ixo DID</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>cred_type</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>issuer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>issued</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>claim</td>
                  <td><a href="#did.Claim">Claim</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.IxoDid">IxoDid</h3>
        <p>An ixo DID with public and private keys, based on the Sovrin DID spec</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>verify_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>encryption_public_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>secret</td>
                  <td><a href="#did.Secret">Secret</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.Secret">Secret</h3>
        <p>The private section of an ixo DID, based on the Sovrin DID spec</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>seed</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sign_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>encryption_private_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="did/diddoc.proto">did/diddoc.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="did.BaseDidDoc">BaseDidDoc</h3>
        <p>BaseDidDoc defines a base DID document type. It implements the DidDoc interface.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pub_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>credentials</td>
                  <td><a href="#did.DidCredential">DidCredential</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="did/genesis.proto">did/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="did.GenesisState">GenesisState</h3>
        <p>GenesisState defines the did module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did_docs</td>
                  <td><a href="#google.protobuf.Any">google.protobuf.Any</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="did/query.proto">did/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="did.QueryAddressFromBase58EncodedPubkeyRequest">QueryAddressFromBase58EncodedPubkeyRequest</h3>
        <p>QueryAddressFromBase58EncodedPubkeyRequest is the request type for the Query/AddressFromBase58EncodedPubkey RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>pubKey</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryAddressFromBase58EncodedPubkeyResponse">QueryAddressFromBase58EncodedPubkeyResponse</h3>
        <p>QueryAddressFromBase58EncodedPubkeyResponse is the response type for the Query/AddressFromBase58EncodedPubkey RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryAddressFromDidRequest">QueryAddressFromDidRequest</h3>
        <p>QueryAddressFromDidRequest is the request type for the Query/AddressFromDid RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryAddressFromDidResponse">QueryAddressFromDidResponse</h3>
        <p>QueryAddressFromDidResponse is the response type for the Query/AddressFromDid RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryAllDidDocsRequest">QueryAllDidDocsRequest</h3>
        <p>QueryAllDidDocsRequest is the request type for the Query/AllDidDocs RPC method.</p>

        

        
      
        <h3 id="did.QueryAllDidDocsResponse">QueryAllDidDocsResponse</h3>
        <p>QueryAllDidDocsResponse is the response type for the Query/AllDidDocs RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>diddocs</td>
                  <td><a href="#google.protobuf.Any">google.protobuf.Any</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryAllDidsRequest">QueryAllDidsRequest</h3>
        <p>QueryAllDidsRequest is the request type for the Query/AllDids RPC method.</p>

        

        
      
        <h3 id="did.QueryAllDidsResponse">QueryAllDidsResponse</h3>
        <p>QueryAllDidsResponse is the response type for the Query/AllDids RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>dids</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryDidDocRequest">QueryDidDocRequest</h3>
        <p>QueryDidDocRequest is the request type for the Query/DidDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.QueryDidDocResponse">QueryDidDocResponse</h3>
        <p>QueryDidDocResponse is the response type for the Query/DidDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>diddoc</td>
                  <td><a href="#google.protobuf.Any">google.protobuf.Any</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="did.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>DidDoc</td>
                <td><a href="#did.QueryDidDocRequest">QueryDidDocRequest</a></td>
                <td><a href="#did.QueryDidDocResponse">QueryDidDocResponse</a></td>
                <td><p>DidDoc queries info of a specific DID&#39;s DidDoc.</p></td>
              </tr>
            
              <tr>
                <td>AllDids</td>
                <td><a href="#did.QueryAllDidsRequest">QueryAllDidsRequest</a></td>
                <td><a href="#did.QueryAllDidsResponse">QueryAllDidsResponse</a></td>
                <td><p>AllDids returns a list of all existing DIDs.</p></td>
              </tr>
            
              <tr>
                <td>AllDidDocs</td>
                <td><a href="#did.QueryAllDidDocsRequest">QueryAllDidDocsRequest</a></td>
                <td><a href="#did.QueryAllDidDocsResponse">QueryAllDidDocsResponse</a></td>
                <td><p>AllDidDocs returns a list of all existing DidDocs (i.e. all DIDs along with their DidDoc info).</p></td>
              </tr>
            
              <tr>
                <td>AddressFromDid</td>
                <td><a href="#did.QueryAddressFromDidRequest">QueryAddressFromDidRequest</a></td>
                <td><a href="#did.QueryAddressFromDidResponse">QueryAddressFromDidResponse</a></td>
                <td><p>AddressFromDid retrieves the cosmos address associated to an ixo DID.</p></td>
              </tr>
            
              <tr>
                <td>AddressFromBase58EncodedPubkey</td>
                <td><a href="#did.QueryAddressFromBase58EncodedPubkeyRequest">QueryAddressFromBase58EncodedPubkeyRequest</a></td>
                <td><a href="#did.QueryAddressFromBase58EncodedPubkeyResponse">QueryAddressFromBase58EncodedPubkeyResponse</a></td>
                <td><p>AddressFromBase58EncodedPubkey retrieves the cosmos address associated to an ixo DID&#39;s pubkey.</p></td>
              </tr>
            
          </tbody>
        </table>

        
          
          
          <h4>Methods with HTTP bindings</h4>
          <table>
            <thead>
              <tr>
                <td>Method Name</td>
                <td>Method</td>
                <td>Pattern</td>
                <td>Body</td>
              </tr>
            </thead>
            <tbody>
            
              
              
              <tr>
                <td>DidDoc</td>
                <td>GET</td>
                <td>/ixo/did/{did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AllDids</td>
                <td>GET</td>
                <td>/ixo/did</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AllDidDocs</td>
                <td>GET</td>
                <td>/ixo/allDidDocs</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AddressFromDid</td>
                <td>GET</td>
                <td>/ixo/didToAddr/{did=**}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AddressFromBase58EncodedPubkey</td>
                <td>GET</td>
                <td>/ixo/pubKeyToAddr/{pubKey}</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="did/tx.proto">did/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="did.MsgAddCredential">MsgAddCredential</h3>
        <p>MsgAddCredential defines a message for adding a credential to the signer's DID.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did_credential</td>
                  <td><a href="#did.DidCredential">DidCredential</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.MsgAddCredentialResponse">MsgAddCredentialResponse</h3>
        <p>MsgAddCredentialResponse defines the Msg/AddCredential response type.</p>

        

        
      
        <h3 id="did.MsgAddDid">MsgAddDid</h3>
        <p>MsgAddDid defines a message for adding a DID.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pubKey</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="did.MsgAddDidResponse">MsgAddDidResponse</h3>
        <p>MsgAddDidResponse defines the Msg/AddDid response type.</p>

        

        
      

      

      

      
        <h3 id="did.Msg">Msg</h3>
        <p>Msg defines the did Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>AddDid</td>
                <td><a href="#did.MsgAddDid">MsgAddDid</a></td>
                <td><a href="#did.MsgAddDidResponse">MsgAddDidResponse</a></td>
                <td><p>AddDid defines a method for adding a DID.</p></td>
              </tr>
            
              <tr>
                <td>AddCredential</td>
                <td><a href="#did.MsgAddCredential">MsgAddCredential</a></td>
                <td><a href="#did.MsgAddCredentialResponse">MsgAddCredentialResponse</a></td>
                <td><p>AddCredential defines a method for adding a credential to the signer&#39;s DID.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
        <h2 id="payments/payments.proto">payments/payments.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="payments.BlockPeriod">BlockPeriod</h3>
        <p>BlockPeriod implements the Period interface and specifies a period in terms of number</p><p>of blocks.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>period_length</td>
                  <td><a href="#int64">int64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>period_start_block</td>
                  <td><a href="#int64">int64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.Discount">Discount</h3>
        <p>Discount contains details about a discount which can be granted to payers.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>percent</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.DistributionShare">DistributionShare</h3>
        <p>DistributionShare specifies the share of a specific payment an address will receive.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.PaymentContract">PaymentContract</h3>
        <p>PaymentContract specifies an agreement between a payer and payee/s which can be invoked</p><p>once or multiple times to effect payment/s.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_template_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>recipients</td>
                  <td><a href="#payments.DistributionShare">DistributionShare</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>cumulative_pay</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>current_remainder</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>can_deauthorise</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>authorised</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>discount_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.PaymentTemplate">PaymentTemplate</h3>
        <p>PaymentTemplate contains details about a payment, with no info about the payer or payee.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_minimum</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_maximum</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>discounts</td>
                  <td><a href="#payments.Discount">Discount</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.Subscription">Subscription</h3>
        <p>Subscription specifies details of a payment to be effected periodically.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>periods_so_far</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_periods</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>periods_accumulated</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>period</td>
                  <td><a href="#google.protobuf.Any">google.protobuf.Any</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.TestPeriod">TestPeriod</h3>
        <p>TestPeriod implements the Period interface and is identical to BlockPeriod, except it</p><p>ignores the context in periodEnded() and periodStarted().</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>period_length</td>
                  <td><a href="#int64">int64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>period_start_block</td>
                  <td><a href="#int64">int64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.TimePeriod">TimePeriod</h3>
        <p>TimePeriod implements the Period interface and specifies a period in terms of time.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>period_duration_ns</td>
                  <td><a href="#google.protobuf.Duration">google.protobuf.Duration</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>period_start_time</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="payments/genesis.proto">payments/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="payments.GenesisState">GenesisState</h3>
        <p>GenesisState defines the payments module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_templates</td>
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contracts</td>
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>subscriptions</td>
                  <td><a href="#payments.Subscription">Subscription</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="payments/query.proto">payments/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="payments.QueryPaymentContractRequest">QueryPaymentContractRequest</h3>
        <p>QueryPaymentContractRequest is the request type for the Query/PaymentContract RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QueryPaymentContractResponse">QueryPaymentContractResponse</h3>
        <p>QueryPaymentContractResponse is the response type for the Query/PaymentContract RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contract</td>
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</h3>
        <p>QueryPaymentContractsByIdPrefixRequest is the request type for the Query/PaymentContractsByIdPrefix RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contracts_id_prefix</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</h3>
        <p>QueryPaymentContractsByIdPrefixResponse is the response type for the Query/PaymentContractsByIdPrefix RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contracts</td>
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</h3>
        <p>QueryPaymentTemplateRequest is the request type for the Query/PaymentTemplate RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_template_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</h3>
        <p>QueryPaymentTemplateResponse is the response type for the Query/PaymentTemplate RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_template</td>
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QuerySubscriptionRequest">QuerySubscriptionRequest</h3>
        <p>QuerySubscriptionRequest is the request type for the Query/Subscription RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>subscription_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.QuerySubscriptionResponse">QuerySubscriptionResponse</h3>
        <p>QuerySubscriptionResponse is the response type for the Query/Subscription RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>subscription</td>
                  <td><a href="#payments.Subscription">Subscription</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="payments.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>PaymentTemplate</td>
                <td><a href="#payments.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</a></td>
                <td><a href="#payments.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</a></td>
                <td><p>PaymentTemplate queries info of a specific payment template.</p></td>
              </tr>
            
              <tr>
                <td>PaymentContract</td>
                <td><a href="#payments.QueryPaymentContractRequest">QueryPaymentContractRequest</a></td>
                <td><a href="#payments.QueryPaymentContractResponse">QueryPaymentContractResponse</a></td>
                <td><p>PaymentContract queries info of a specific payment contract.</p></td>
              </tr>
            
              <tr>
                <td>PaymentContractsByIdPrefix</td>
                <td><a href="#payments.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</a></td>
                <td><a href="#payments.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</a></td>
                <td><p>PaymentContractsByIdPrefix lists all payment contracts having an id with a specific prefix.</p></td>
              </tr>
            
              <tr>
                <td>Subscription</td>
                <td><a href="#payments.QuerySubscriptionRequest">QuerySubscriptionRequest</a></td>
                <td><a href="#payments.QuerySubscriptionResponse">QuerySubscriptionResponse</a></td>
                <td><p>Subscription queries info of a specific Subscription.</p></td>
              </tr>
            
          </tbody>
        </table>

        
          
          
          <h4>Methods with HTTP bindings</h4>
          <table>
            <thead>
              <tr>
                <td>Method Name</td>
                <td>Method</td>
                <td>Pattern</td>
                <td>Body</td>
              </tr>
            </thead>
            <tbody>
            
              
              
              <tr>
                <td>PaymentTemplate</td>
                <td>GET</td>
                <td>/ixo/payments/templates/{payment_template_id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>PaymentContract</td>
                <td>GET</td>
                <td>/ixo/payments/contracts/{payment_contract_id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>PaymentContractsByIdPrefix</td>
                <td>GET</td>
                <td>/ixo/payments/contracts_by_id_prefix/{payment_contracts_id_prefix}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>Subscription</td>
                <td>GET</td>
                <td>/ixo/payments/subscriptions/{subscription_id}</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="payments/tx.proto">payments/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="payments.MsgCreatePaymentContract">MsgCreatePaymentContract</h3>
        <p>MsgCreatePaymentContract defines a message for creating a payment contract.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_template_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>recipients</td>
                  <td><a href="#payments.DistributionShare">DistributionShare</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>can_deauthorise</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>discount_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</h3>
        <p>MsgCreatePaymentContractResponse defines the Msg/CreatePaymentContract response type.</p>

        

        
      
        <h3 id="payments.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</h3>
        <p>MsgCreatePaymentTemplate defines a message for creating a payment template.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_template</td>
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</h3>
        <p>MsgCreatePaymentTemplateResponse defines the Msg/CreatePaymentTemplate response type.</p>

        

        
      
        <h3 id="payments.MsgCreateSubscription">MsgCreateSubscription</h3>
        <p>MsgCreateSubscription defines a message for creating a subscription.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>subscription_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>max_periods</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>period</td>
                  <td><a href="#google.protobuf.Any">google.protobuf.Any</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</h3>
        <p>MsgCreateSubscriptionResponse defines the Msg/CreateSubscription response type.</p>

        

        
      
        <h3 id="payments.MsgEffectPayment">MsgEffectPayment</h3>
        <p>MsgEffectPayment defines a message for putting a specific payment contract into effect.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgEffectPaymentResponse">MsgEffectPaymentResponse</h3>
        <p>MsgEffectPaymentResponse defines the Msg/EffectPayment response type.</p>

        

        
      
        <h3 id="payments.MsgGrantDiscount">MsgGrantDiscount</h3>
        <p>MsgGrantDiscount defines a message for granting a discount to a payer on a specific payment contract.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>discount_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>recipient</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgGrantDiscountResponse">MsgGrantDiscountResponse</h3>
        <p>MsgGrantDiscountResponse defines the Msg/GrantDiscount response type.</p>

        

        
      
        <h3 id="payments.MsgRevokeDiscount">MsgRevokeDiscount</h3>
        <p>MsgRevokeDiscount defines a message for revoking a discount previously granted to a payer.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>holder</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</h3>
        <p>MsgRevokeDiscountResponse defines the Msg/RevokeDiscount response type.</p>

        

        
      
        <h3 id="payments.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</h3>
        <p>MsgSetPaymentContractAuthorisation defines a message for authorising or deauthorising a payment contract.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contract_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payer_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>authorised</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="payments.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</h3>
        <p>MsgSetPaymentContractAuthorisationResponse defines the Msg/SetPaymentContractAuthorisation response type.</p>

        

        
      

      

      

      
        <h3 id="payments.Msg">Msg</h3>
        <p>Msg defines the payments Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>SetPaymentContractAuthorisation</td>
                <td><a href="#payments.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</a></td>
                <td><a href="#payments.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</a></td>
                <td><p>SetPaymentContractAuthorisation defines a method for authorising or deauthorising a payment contract.</p></td>
              </tr>
            
              <tr>
                <td>CreatePaymentTemplate</td>
                <td><a href="#payments.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</a></td>
                <td><a href="#payments.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</a></td>
                <td><p>CreatePaymentTemplate defines a method for creating a payment template.</p></td>
              </tr>
            
              <tr>
                <td>CreatePaymentContract</td>
                <td><a href="#payments.MsgCreatePaymentContract">MsgCreatePaymentContract</a></td>
                <td><a href="#payments.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</a></td>
                <td><p>CreatePaymentContract defines a method for creating a payment contract.</p></td>
              </tr>
            
              <tr>
                <td>CreateSubscription</td>
                <td><a href="#payments.MsgCreateSubscription">MsgCreateSubscription</a></td>
                <td><a href="#payments.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</a></td>
                <td><p>CreateSubscription defines a method for creating a subscription.</p></td>
              </tr>
            
              <tr>
                <td>GrantDiscount</td>
                <td><a href="#payments.MsgGrantDiscount">MsgGrantDiscount</a></td>
                <td><a href="#payments.MsgGrantDiscountResponse">MsgGrantDiscountResponse</a></td>
                <td><p>GrantDiscount defines a method for granting a discount to a payer on a specific payment contract.</p></td>
              </tr>
            
              <tr>
                <td>RevokeDiscount</td>
                <td><a href="#payments.MsgRevokeDiscount">MsgRevokeDiscount</a></td>
                <td><a href="#payments.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</a></td>
                <td><p>RevokeDiscount defines a method for revoking a discount previously granted to a payer.</p></td>
              </tr>
            
              <tr>
                <td>EffectPayment</td>
                <td><a href="#payments.MsgEffectPayment">MsgEffectPayment</a></td>
                <td><a href="#payments.MsgEffectPaymentResponse">MsgEffectPaymentResponse</a></td>
                <td><p>EffectPayment defines a method for putting a specific payment contract into effect.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
        <h2 id="project/project.proto">project/project.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="project.AccountMap">AccountMap</h3>
        <p>AccountMap maps a specific project's account names to the accounts' addresses.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
                  <td><a href="#project.AccountMap.MapEntry">AccountMap.MapEntry</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.AccountMap.MapEntry">AccountMap.MapEntry</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>value</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.Claim">Claim</h3>
        <p>Claim contains details required to start a claim on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>template_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>claimer_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.Claims">Claims</h3>
        <p>Claims contains a list of type Claim.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>claims_list</td>
                  <td><a href="#project.Claim">Claim</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.CreateAgentDoc">CreateAgentDoc</h3>
        <p>CreateAgentDoc contains details required to create an agent.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>agent_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>role</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.CreateClaimDoc">CreateClaimDoc</h3>
        <p>CreateClaimDoc contains details required to create a claim on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>claim_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>claim_template_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.CreateEvaluationDoc">CreateEvaluationDoc</h3>
        <p>CreateEvaluationDoc contains details required to create an evaluation for a specific claim on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>claim_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.GenesisAccountMap">GenesisAccountMap</h3>
        <p>GenesisAccountMap is a type used at genesis that maps a specific project's account names to the accounts' addresses.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
                  <td><a href="#project.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>value</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.Params">Params</h3>
        <p>Params defines the parameters for the project module.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>ixo_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_minimum_initial_funding</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>oracle_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>node_fee_percentage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.ProjectDoc">ProjectDoc</h3>
        <p>ProjectDoc defines a project (or entity) type with all of its parameters.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pub_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#bytes">bytes</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.UpdateAgentDoc">UpdateAgentDoc</h3>
        <p>UpdateAgentDoc contains details required to update an agent.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>role</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.UpdateProjectStatusDoc">UpdateProjectStatusDoc</h3>
        <p>UpdateProjectStatusDoc contains details required to update a project's status.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>eth_funding_txn_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.WithdrawFundsDoc">WithdrawFundsDoc</h3>
        <p>WithdrawFundsDoc contains details required to withdraw funds from a specific project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>recipient_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>is_refund</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.WithdrawalInfoDoc">WithdrawalInfoDoc</h3>
        <p>WithdrawalInfoDoc contains details required to withdraw from a specific project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>recipient_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.WithdrawalInfoDocs">WithdrawalInfoDocs</h3>
        <p>WithdrawalInfoDocs contains a list of type WithdrawalInfoDoc.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>docs_list</td>
                  <td><a href="#project.WithdrawalInfoDoc">WithdrawalInfoDoc</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="project/genesis.proto">project/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="project.GenesisState">GenesisState</h3>
        <p>GenesisState defines the project module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_docs</td>
                  <td><a href="#project.ProjectDoc">ProjectDoc</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>account_maps</td>
                  <td><a href="#project.GenesisAccountMap">GenesisAccountMap</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>withdrawals_infos</td>
                  <td><a href="#project.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>claims</td>
                  <td><a href="#project.Claims">Claims</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
                  <td><a href="#project.Params">Params</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="project/query.proto">project/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="project.QueryParamsRequest">QueryParamsRequest</h3>
        <p>QueryParamsRequest is the request type for the Query/Params RPC method.</p>

        

        
      
        <h3 id="project.QueryParamsResponse">QueryParamsResponse</h3>
        <p>QueryParamsResponse is the response type for the Query/Params RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
                  <td><a href="#project.Params">Params</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectAccountsRequest">QueryProjectAccountsRequest</h3>
        <p>QueryProjectAccountsRequest is the request type for the Query/ProjectAccounts RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectAccountsResponse">QueryProjectAccountsResponse</h3>
        <p>QueryProjectAccountsResponse is the response type for the Query/ProjectAccounts RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>account_map</td>
                  <td><a href="#project.AccountMap">AccountMap</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectDocRequest">QueryProjectDocRequest</h3>
        <p>QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectDocResponse">QueryProjectDocResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_doc</td>
                  <td><a href="#project.ProjectDoc">ProjectDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectTxRequest">QueryProjectTxRequest</h3>
        <p>QueryProjectTxRequest is the request type for the Query/ProjectTx RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.QueryProjectTxResponse">QueryProjectTxResponse</h3>
        <p>QueryProjectTxResponse is the response type for the Query/ProjectTx RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>txs</td>
                  <td><a href="#project.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="project.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>ProjectDoc</td>
                <td><a href="#project.QueryProjectDocRequest">QueryProjectDocRequest</a></td>
                <td><a href="#project.QueryProjectDocResponse">QueryProjectDocResponse</a></td>
                <td><p>ProjectDoc queries info of a specific project.</p></td>
              </tr>
            
              <tr>
                <td>ProjectAccounts</td>
                <td><a href="#project.QueryProjectAccountsRequest">QueryProjectAccountsRequest</a></td>
                <td><a href="#project.QueryProjectAccountsResponse">QueryProjectAccountsResponse</a></td>
                <td><p>ProjectAccounts lists a specific project&#39;s accounts.</p></td>
              </tr>
            
              <tr>
                <td>ProjectTx</td>
                <td><a href="#project.QueryProjectTxRequest">QueryProjectTxRequest</a></td>
                <td><a href="#project.QueryProjectTxResponse">QueryProjectTxResponse</a></td>
                <td><p>ProjectTx lists a specific project&#39;s transactions.</p></td>
              </tr>
            
              <tr>
                <td>Params</td>
                <td><a href="#project.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#project.QueryParamsResponse">QueryParamsResponse</a></td>
                <td><p>Params queries the paramaters of x/project module.</p></td>
              </tr>
            
          </tbody>
        </table>

        
          
          
          <h4>Methods with HTTP bindings</h4>
          <table>
            <thead>
              <tr>
                <td>Method Name</td>
                <td>Method</td>
                <td>Pattern</td>
                <td>Body</td>
              </tr>
            </thead>
            <tbody>
            
              
              
              <tr>
                <td>ProjectDoc</td>
                <td>GET</td>
                <td>/ixo/project/{project_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>ProjectAccounts</td>
                <td>GET</td>
                <td>/ixo/projectAccounts/{project_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>ProjectTx</td>
                <td>GET</td>
                <td>/ixo/projectTxs/{project_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>Params</td>
                <td>GET</td>
                <td>/ixo/projectParams</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="project/tx.proto">project/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="project.MsgCreateAgent">MsgCreateAgent</h3>
        <p>MsgCreateAgent defines a message for creating an agent on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.CreateAgentDoc">CreateAgentDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgCreateAgentResponse">MsgCreateAgentResponse</h3>
        <p>MsgCreateAgentResponse defines the Msg/CreateAgent response type.</p>

        

        
      
        <h3 id="project.MsgCreateClaim">MsgCreateClaim</h3>
        <p>MsgCreateClaim defines a message for creating a claim on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.CreateClaimDoc">CreateClaimDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgCreateClaimResponse">MsgCreateClaimResponse</h3>
        <p>MsgCreateClaimResponse defines the Msg/CreateClaim response type.</p>

        

        
      
        <h3 id="project.MsgCreateEvaluation">MsgCreateEvaluation</h3>
        <p>MsgCreateEvaluation defines a message for creating an evaluation for a specific claim on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.CreateEvaluationDoc">CreateEvaluationDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</h3>
        <p>MsgCreateEvaluationResponse defines the Msg/CreateEvaluation response type.</p>

        

        
      
        <h3 id="project.MsgCreateProject">MsgCreateProject</h3>
        <p>MsgCreateProject defines a message for creating a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pub_key</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#bytes">bytes</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgCreateProjectResponse">MsgCreateProjectResponse</h3>
        <p>MsgCreateProjectResponse defines the Msg/CreateProject response type.</p>

        

        
      
        <h3 id="project.MsgUpdateAgent">MsgUpdateAgent</h3>
        <p>MsgUpdateAgent defines a message for updating an agent on a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.UpdateAgentDoc">UpdateAgentDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgUpdateAgentResponse">MsgUpdateAgentResponse</h3>
        <p>MsgUpdateAgentResponse defines the Msg/UpdateAgent response type.</p>

        

        
      
        <h3 id="project.MsgUpdateProjectDoc">MsgUpdateProjectDoc</h3>
        <p>MsgUpdateProjectDoc defines a message for updating a project's data.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#bytes">bytes</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</h3>
        <p>MsgUpdateProjectDocResponse defines the Msg/UpdateProjectDoc response type.</p>

        

        
      
        <h3 id="project.MsgUpdateProjectStatus">MsgUpdateProjectStatus</h3>
        <p>MsgUpdateProjectStatus defines a message for updating a project's current status.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tx_hash</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.UpdateProjectStatusDoc">UpdateProjectStatusDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</h3>
        <p>MsgUpdateProjectStatusResponse defines the Msg/UpdateProjectStatus response type.</p>

        

        
      
        <h3 id="project.MsgWithdrawFunds">MsgWithdrawFunds</h3>
        <p>MsgWithdrawFunds defines a message for project agents to withdraw their funds from a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>sender_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#project.WithdrawFundsDoc">WithdrawFundsDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="project.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</h3>
        <p>MsgWithdrawFundsResponse defines the Msg/WithdrawFunds response type.</p>

        

        
      

      

      

      
        <h3 id="project.Msg">Msg</h3>
        <p>Msg defines the project Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateProject</td>
                <td><a href="#project.MsgCreateProject">MsgCreateProject</a></td>
                <td><a href="#project.MsgCreateProjectResponse">MsgCreateProjectResponse</a></td>
                <td><p>CreateProject defines a method for creating a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateProjectStatus</td>
                <td><a href="#project.MsgUpdateProjectStatus">MsgUpdateProjectStatus</a></td>
                <td><a href="#project.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</a></td>
                <td><p>UpdateProjectStatus defines a method for updating a project&#39;s current status.</p></td>
              </tr>
            
              <tr>
                <td>CreateAgent</td>
                <td><a href="#project.MsgCreateAgent">MsgCreateAgent</a></td>
                <td><a href="#project.MsgCreateAgentResponse">MsgCreateAgentResponse</a></td>
                <td><p>CreateAgent defines a method for creating an agent on a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateAgent</td>
                <td><a href="#project.MsgUpdateAgent">MsgUpdateAgent</a></td>
                <td><a href="#project.MsgUpdateAgentResponse">MsgUpdateAgentResponse</a></td>
                <td><p>UpdateAgent defines a method for updating an agent on a project.</p></td>
              </tr>
            
              <tr>
                <td>CreateClaim</td>
                <td><a href="#project.MsgCreateClaim">MsgCreateClaim</a></td>
                <td><a href="#project.MsgCreateClaimResponse">MsgCreateClaimResponse</a></td>
                <td><p>CreateClaim defines a method for creating a claim on a project.</p></td>
              </tr>
            
              <tr>
                <td>CreateEvaluation</td>
                <td><a href="#project.MsgCreateEvaluation">MsgCreateEvaluation</a></td>
                <td><a href="#project.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</a></td>
                <td><p>CreateEvaluation defines a method for creating an evaluation for a specific claim on a project.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawFunds</td>
                <td><a href="#project.MsgWithdrawFunds">MsgWithdrawFunds</a></td>
                <td><a href="#project.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</a></td>
                <td><p>WithdrawFunds defines a method for project agents to withdraw their funds from a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateProjectDoc</td>
                <td><a href="#project.MsgUpdateProjectDoc">MsgUpdateProjectDoc</a></td>
                <td><a href="#project.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</a></td>
                <td><p>UpdateProjectDoc defines a method for updating a project&#39;s data.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    

    <h2 id="scalar-value-types">Scalar Value Types</h2>
    <table class="scalar-value-types-table">
      <thead>
        <tr><td>.proto Type</td><td>Notes</td><td>C++</td><td>Java</td><td>Python</td><td>Go</td><td>C#</td><td>PHP</td><td>Ruby</td></tr>
      </thead>
      <tbody>
        
          <tr id="double">
            <td>double</td>
            <td></td>
            <td>double</td>
            <td>double</td>
            <td>float</td>
            <td>float64</td>
            <td>double</td>
            <td>float</td>
            <td>Float</td>
          </tr>
        
          <tr id="float">
            <td>float</td>
            <td></td>
            <td>float</td>
            <td>float</td>
            <td>float</td>
            <td>float32</td>
            <td>float</td>
            <td>float</td>
            <td>Float</td>
          </tr>
        
          <tr id="int32">
            <td>int32</td>
            <td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="int64">
            <td>int64</td>
            <td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="uint32">
            <td>uint32</td>
            <td>Uses variable-length encoding.</td>
            <td>uint32</td>
            <td>int</td>
            <td>int/long</td>
            <td>uint32</td>
            <td>uint</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="uint64">
            <td>uint64</td>
            <td>Uses variable-length encoding.</td>
            <td>uint64</td>
            <td>long</td>
            <td>int/long</td>
            <td>uint64</td>
            <td>ulong</td>
            <td>integer/string</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sint32">
            <td>sint32</td>
            <td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sint64">
            <td>sint64</td>
            <td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="fixed32">
            <td>fixed32</td>
            <td>Always four bytes. More efficient than uint32 if values are often greater than 2^28.</td>
            <td>uint32</td>
            <td>int</td>
            <td>int</td>
            <td>uint32</td>
            <td>uint</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="fixed64">
            <td>fixed64</td>
            <td>Always eight bytes. More efficient than uint64 if values are often greater than 2^56.</td>
            <td>uint64</td>
            <td>long</td>
            <td>int/long</td>
            <td>uint64</td>
            <td>ulong</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="sfixed32">
            <td>sfixed32</td>
            <td>Always four bytes.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sfixed64">
            <td>sfixed64</td>
            <td>Always eight bytes.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="bool">
            <td>bool</td>
            <td></td>
            <td>bool</td>
            <td>boolean</td>
            <td>boolean</td>
            <td>bool</td>
            <td>bool</td>
            <td>boolean</td>
            <td>TrueClass/FalseClass</td>
          </tr>
        
          <tr id="string">
            <td>string</td>
            <td>A string must always contain UTF-8 encoded or 7-bit ASCII text.</td>
            <td>string</td>
            <td>String</td>
            <td>str/unicode</td>
            <td>string</td>
            <td>string</td>
            <td>string</td>
            <td>String (UTF-8)</td>
          </tr>
        
          <tr id="bytes">
            <td>bytes</td>
            <td>May contain any arbitrary sequence of bytes.</td>
            <td>string</td>
            <td>ByteString</td>
            <td>str</td>
            <td>[]byte</td>
            <td>ByteString</td>
            <td>string</td>
            <td>String (ASCII-8BIT)</td>
          </tr>
        
      </tbody>
    </table>
  </body>
</html>

