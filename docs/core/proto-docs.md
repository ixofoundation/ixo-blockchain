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
<<<<<<< HEAD
            <a href="#ixo%2fbonds%2fv1beta1%2fbonds.proto">ixo/bonds/v1beta1/bonds.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.BaseOrder"><span class="badge">M</span>BaseOrder</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.Batch"><span class="badge">M</span>Batch</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.Bond"><span class="badge">M</span>Bond</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.BondDetails"><span class="badge">M</span>BondDetails</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.BuyOrder"><span class="badge">M</span>BuyOrder</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.FunctionParam"><span class="badge">M</span>FunctionParam</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.Params"><span class="badge">M</span>Params</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.SellOrder"><span class="badge">M</span>SellOrder</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.SwapOrder"><span class="badge">M</span>SwapOrder</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fbonds%2fv1beta1%2fgenesis.proto">ixo/bonds/v1beta1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.GenesisState"><span class="badge">M</span>GenesisState</a>
=======
            <a href="#bonds%2fgenesis.proto">bonds/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#bonds.GenesisState"><span class="badge">M</span>GenesisState</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fbonds%2fv1beta1%2fquery.proto">ixo/bonds/v1beta1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest"><span class="badge">M</span>QueryAlphaMaximumsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse"><span class="badge">M</span>QueryAlphaMaximumsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryAvailableReserveRequest"><span class="badge">M</span>QueryAvailableReserveRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryAvailableReserveResponse"><span class="badge">M</span>QueryAvailableReserveResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBatchRequest"><span class="badge">M</span>QueryBatchRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBatchResponse"><span class="badge">M</span>QueryBatchResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondRequest"><span class="badge">M</span>QueryBondRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondResponse"><span class="badge">M</span>QueryBondResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondsDetailedRequest"><span class="badge">M</span>QueryBondsDetailedRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondsDetailedResponse"><span class="badge">M</span>QueryBondsDetailedResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondsRequest"><span class="badge">M</span>QueryBondsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBondsResponse"><span class="badge">M</span>QueryBondsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBuyPriceRequest"><span class="badge">M</span>QueryBuyPriceRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryBuyPriceResponse"><span class="badge">M</span>QueryBuyPriceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCurrentPriceRequest"><span class="badge">M</span>QueryCurrentPriceRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCurrentPriceResponse"><span class="badge">M</span>QueryCurrentPriceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCurrentReserveRequest"><span class="badge">M</span>QueryCurrentReserveRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCurrentReserveResponse"><span class="badge">M</span>QueryCurrentReserveResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCustomPriceRequest"><span class="badge">M</span>QueryCustomPriceRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryCustomPriceResponse"><span class="badge">M</span>QueryCustomPriceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryLastBatchRequest"><span class="badge">M</span>QueryLastBatchRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryLastBatchResponse"><span class="badge">M</span>QueryLastBatchResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryParamsRequest"><span class="badge">M</span>QueryParamsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QueryParamsResponse"><span class="badge">M</span>QueryParamsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QuerySellReturnRequest"><span class="badge">M</span>QuerySellReturnRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QuerySellReturnResponse"><span class="badge">M</span>QuerySellReturnResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QuerySwapReturnRequest"><span class="badge">M</span>QuerySwapReturnRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.QuerySwapReturnResponse"><span class="badge">M</span>QuerySwapReturnResponse</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
                <li>
<<<<<<< HEAD
                  <a href="#ixo.bonds.v1beta1.Query"><span class="badge">S</span>Query</a>
=======
                  <a href="#bonds.Query"><span class="badge">S</span>Query</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fbonds%2fv1beta1%2ftx.proto">ixo/bonds/v1beta1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgBuy"><span class="badge">M</span>MsgBuy</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgBuyResponse"><span class="badge">M</span>MsgBuyResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgCreateBond"><span class="badge">M</span>MsgCreateBond</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgCreateBondResponse"><span class="badge">M</span>MsgCreateBondResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgEditBond"><span class="badge">M</span>MsgEditBond</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgEditBondResponse"><span class="badge">M</span>MsgEditBondResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgMakeOutcomePayment"><span class="badge">M</span>MsgMakeOutcomePayment</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse"><span class="badge">M</span>MsgMakeOutcomePaymentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSell"><span class="badge">M</span>MsgSell</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSellResponse"><span class="badge">M</span>MsgSellResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSetNextAlpha"><span class="badge">M</span>MsgSetNextAlpha</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSetNextAlphaResponse"><span class="badge">M</span>MsgSetNextAlphaResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSwap"><span class="badge">M</span>MsgSwap</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgSwapResponse"><span class="badge">M</span>MsgSwapResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgUpdateBondState"><span class="badge">M</span>MsgUpdateBondState</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgUpdateBondStateResponse"><span class="badge">M</span>MsgUpdateBondStateResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgWithdrawReserve"><span class="badge">M</span>MsgWithdrawReserve</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgWithdrawReserveResponse"><span class="badge">M</span>MsgWithdrawReserveResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgWithdrawShare"><span class="badge">M</span>MsgWithdrawShare</a>
                </li>
              
                <li>
                  <a href="#ixo.bonds.v1beta1.MsgWithdrawShareResponse"><span class="badge">M</span>MsgWithdrawShareResponse</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
                <li>
<<<<<<< HEAD
                  <a href="#ixo.bonds.v1beta1.Msg"><span class="badge">S</span>Msg</a>
=======
                  <a href="#bonds.Msg"><span class="badge">S</span>Msg</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fentity%2fv1beta1%2fentity.proto">ixo/entity/v1beta1/entity.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.entity.v1beta1.EntityDoc"><span class="badge">M</span>EntityDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.Params"><span class="badge">M</span>Params</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fentity%2fv1beta1%2fgenesis.proto">ixo/entity/v1beta1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.entity.v1beta1.GenesisState"><span class="badge">M</span>GenesisState</a>
=======
            <a href="#did%2fdiddoc.proto">did/diddoc.proto</a>
            <ul>
              
                <li>
                  <a href="#did.BaseDidDoc"><span class="badge">M</span>BaseDidDoc</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fentity%2fv1beta1%2fproposal.proto">ixo/entity/v1beta1/proposal.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.entity.v1beta1.InitializeNftContract"><span class="badge">M</span>InitializeNftContract</a>
=======
            <a href="#did%2fgenesis.proto">did/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#did.GenesisState"><span class="badge">M</span>GenesisState</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fentity%2fv1beta1%2fquery.proto">ixo/entity/v1beta1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityConfigRequest"><span class="badge">M</span>QueryEntityConfigRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityConfigResponse"><span class="badge">M</span>QueryEntityConfigResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityConfigResponse.MapEntry"><span class="badge">M</span>QueryEntityConfigResponse.MapEntry</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityDocRequest"><span class="badge">M</span>QueryEntityDocRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityDocResponse"><span class="badge">M</span>QueryEntityDocResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityListRequest"><span class="badge">M</span>QueryEntityListRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.QueryEntityListResponse"><span class="badge">M</span>QueryEntityListResponse</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
                <li>
<<<<<<< HEAD
                  <a href="#ixo.entity.v1beta1.Query"><span class="badge">S</span>Query</a>
=======
                  <a href="#did.Query"><span class="badge">S</span>Query</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fiid%2fv1beta1%2fiid.proto">ixo/iid/v1beta1/iid.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.iid.v1beta1.AccordedRight"><span class="badge">M</span>AccordedRight</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.Context"><span class="badge">M</span>Context</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.IidDocument"><span class="badge">M</span>IidDocument</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.IidMetadata"><span class="badge">M</span>IidMetadata</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.LinkedEntity"><span class="badge">M</span>LinkedEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.LinkedResource"><span class="badge">M</span>LinkedResource</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.Service"><span class="badge">M</span>Service</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.VerificationMethod"><span class="badge">M</span>VerificationMethod</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fiid%2fv1beta1%2ftx.proto">ixo/iid/v1beta1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddAccordedRight"><span class="badge">M</span>MsgAddAccordedRight</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddAccordedRightResponse"><span class="badge">M</span>MsgAddAccordedRightResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddController"><span class="badge">M</span>MsgAddController</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddControllerResponse"><span class="badge">M</span>MsgAddControllerResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddIidContext"><span class="badge">M</span>MsgAddIidContext</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddIidContextResponse"><span class="badge">M</span>MsgAddIidContextResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddLinkedEntity"><span class="badge">M</span>MsgAddLinkedEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddLinkedEntityResponse"><span class="badge">M</span>MsgAddLinkedEntityResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddLinkedResource"><span class="badge">M</span>MsgAddLinkedResource</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddLinkedResourceResponse"><span class="badge">M</span>MsgAddLinkedResourceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddService"><span class="badge">M</span>MsgAddService</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddServiceResponse"><span class="badge">M</span>MsgAddServiceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddVerification"><span class="badge">M</span>MsgAddVerification</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgAddVerificationResponse"><span class="badge">M</span>MsgAddVerificationResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgCreateIidDocument"><span class="badge">M</span>MsgCreateIidDocument</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgCreateIidDocumentResponse"><span class="badge">M</span>MsgCreateIidDocumentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeactivateIID"><span class="badge">M</span>MsgDeactivateIID</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeactivateIIDResponse"><span class="badge">M</span>MsgDeactivateIIDResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteAccordedRight"><span class="badge">M</span>MsgDeleteAccordedRight</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse"><span class="badge">M</span>MsgDeleteAccordedRightResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteController"><span class="badge">M</span>MsgDeleteController</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteControllerResponse"><span class="badge">M</span>MsgDeleteControllerResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteIidContext"><span class="badge">M</span>MsgDeleteIidContext</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteIidContextResponse"><span class="badge">M</span>MsgDeleteIidContextResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteLinkedEntity"><span class="badge">M</span>MsgDeleteLinkedEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse"><span class="badge">M</span>MsgDeleteLinkedEntityResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteLinkedResource"><span class="badge">M</span>MsgDeleteLinkedResource</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse"><span class="badge">M</span>MsgDeleteLinkedResourceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteService"><span class="badge">M</span>MsgDeleteService</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgDeleteServiceResponse"><span class="badge">M</span>MsgDeleteServiceResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgRevokeVerification"><span class="badge">M</span>MsgRevokeVerification</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgRevokeVerificationResponse"><span class="badge">M</span>MsgRevokeVerificationResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgSetVerificationRelationships"><span class="badge">M</span>MsgSetVerificationRelationships</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse"><span class="badge">M</span>MsgSetVerificationRelationshipsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgUpdateIidDocument"><span class="badge">M</span>MsgUpdateIidDocument</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse"><span class="badge">M</span>MsgUpdateIidDocumentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgUpdateIidMeta"><span class="badge">M</span>MsgUpdateIidMeta</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.MsgUpdateIidMetaResponse"><span class="badge">M</span>MsgUpdateIidMetaResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.Verification"><span class="badge">M</span>Verification</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.iid.v1beta1.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fentity%2fv1beta1%2ftx.proto">ixo/entity/v1beta1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgCreateEntity"><span class="badge">M</span>MsgCreateEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgCreateEntityResponse"><span class="badge">M</span>MsgCreateEntityResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgTransferEntity"><span class="badge">M</span>MsgTransferEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgTransferEntityResponse"><span class="badge">M</span>MsgTransferEntityResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgUpdateEntity"><span class="badge">M</span>MsgUpdateEntity</a>
                </li>
              
                <li>
                  <a href="#ixo.entity.v1beta1.MsgUpdateEntityResponse"><span class="badge">M</span>MsgUpdateEntityResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.entity.v1beta1.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fiid%2fv1beta1%2fevent.proto">ixo/iid/v1beta1/event.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.iid.v1beta1.IidDocumentCreatedEvent"><span class="badge">M</span>IidDocumentCreatedEvent</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.IidDocumentUpdatedEvent"><span class="badge">M</span>IidDocumentUpdatedEvent</a>
=======
            <a href="#payments%2fgenesis.proto">payments/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#payments.GenesisState"><span class="badge">M</span>GenesisState</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fiid%2fv1beta1%2fgenesis.proto">ixo/iid/v1beta1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.iid.v1beta1.GenesisState"><span class="badge">M</span>GenesisState</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fiid%2fv1beta1%2fquery.proto">ixo/iid/v1beta1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidDocumentRequest"><span class="badge">M</span>QueryIidDocumentRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidDocumentResponse"><span class="badge">M</span>QueryIidDocumentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidDocumentsRequest"><span class="badge">M</span>QueryIidDocumentsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidDocumentsResponse"><span class="badge">M</span>QueryIidDocumentsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidMetaDataRequest"><span class="badge">M</span>QueryIidMetaDataRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.iid.v1beta1.QueryIidMetaDataResponse"><span class="badge">M</span>QueryIidMetaDataResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.iid.v1beta1.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2flegacy%2fdid%2fdid.proto">ixo/legacy/did/did.proto</a>
            <ul>
              
                <li>
                  <a href="#legacydid.Claim"><span class="badge">M</span>Claim</a>
                </li>
              
                <li>
                  <a href="#legacydid.DidCredential"><span class="badge">M</span>DidCredential</a>
                </li>
              
                <li>
                  <a href="#legacydid.IxoDid"><span class="badge">M</span>IxoDid</a>
                </li>
              
                <li>
                  <a href="#legacydid.Secret"><span class="badge">M</span>Secret</a>
=======
            <a href="#project%2fgenesis.proto">project/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#project.GenesisState"><span class="badge">M</span>GenesisState</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2flegacy%2fdid%2fdiddoc.proto">ixo/legacy/did/diddoc.proto</a>
            <ul>
              
                <li>
                  <a href="#legacydid.BaseDidDoc"><span class="badge">M</span>BaseDidDoc</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fpayments%2fv1%2fpayments.proto">ixo/payments/v1/payments.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.payments.v1.BlockPeriod"><span class="badge">M</span>BlockPeriod</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.Discount"><span class="badge">M</span>Discount</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.DistributionShare"><span class="badge">M</span>DistributionShare</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.PaymentContract"><span class="badge">M</span>PaymentContract</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.PaymentTemplate"><span class="badge">M</span>PaymentTemplate</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.Subscription"><span class="badge">M</span>Subscription</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.TestPeriod"><span class="badge">M</span>TestPeriod</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.TimePeriod"><span class="badge">M</span>TimePeriod</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fpayments%2fv1%2fgenesis.proto">ixo/payments/v1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.payments.v1.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fpayments%2fv1%2fquery.proto">ixo/payments/v1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentContractRequest"><span class="badge">M</span>QueryPaymentContractRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentContractResponse"><span class="badge">M</span>QueryPaymentContractResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest"><span class="badge">M</span>QueryPaymentContractsByIdPrefixRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse"><span class="badge">M</span>QueryPaymentContractsByIdPrefixResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentTemplateRequest"><span class="badge">M</span>QueryPaymentTemplateRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QueryPaymentTemplateResponse"><span class="badge">M</span>QueryPaymentTemplateResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QuerySubscriptionRequest"><span class="badge">M</span>QuerySubscriptionRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.QuerySubscriptionResponse"><span class="badge">M</span>QuerySubscriptionResponse</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
                <li>
<<<<<<< HEAD
                  <a href="#ixo.payments.v1.Query"><span class="badge">S</span>Query</a>
=======
                  <a href="#project.Query"><span class="badge">S</span>Query</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
            </ul>
          </li>
        
          
          <li>
<<<<<<< HEAD
            <a href="#ixo%2fpayments%2fv1%2ftx.proto">ixo/payments/v1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreatePaymentContract"><span class="badge">M</span>MsgCreatePaymentContract</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreatePaymentContractResponse"><span class="badge">M</span>MsgCreatePaymentContractResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreatePaymentTemplate"><span class="badge">M</span>MsgCreatePaymentTemplate</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreatePaymentTemplateResponse"><span class="badge">M</span>MsgCreatePaymentTemplateResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreateSubscription"><span class="badge">M</span>MsgCreateSubscription</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgCreateSubscriptionResponse"><span class="badge">M</span>MsgCreateSubscriptionResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgEffectPayment"><span class="badge">M</span>MsgEffectPayment</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgEffectPaymentResponse"><span class="badge">M</span>MsgEffectPaymentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgGrantDiscount"><span class="badge">M</span>MsgGrantDiscount</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgGrantDiscountResponse"><span class="badge">M</span>MsgGrantDiscountResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgRevokeDiscount"><span class="badge">M</span>MsgRevokeDiscount</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgRevokeDiscountResponse"><span class="badge">M</span>MsgRevokeDiscountResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgSetPaymentContractAuthorisation"><span class="badge">M</span>MsgSetPaymentContractAuthorisation</a>
                </li>
              
                <li>
                  <a href="#ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse"><span class="badge">M</span>MsgSetPaymentContractAuthorisationResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.payments.v1.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fproject%2fv1%2fproject.proto">ixo/project/v1/project.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.project.v1.AccountMap"><span class="badge">M</span>AccountMap</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.AccountMap.MapEntry"><span class="badge">M</span>AccountMap.MapEntry</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.Claim"><span class="badge">M</span>Claim</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.Claims"><span class="badge">M</span>Claims</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.CreateAgentDoc"><span class="badge">M</span>CreateAgentDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.CreateClaimDoc"><span class="badge">M</span>CreateClaimDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.CreateEvaluationDoc"><span class="badge">M</span>CreateEvaluationDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.GenesisAccountMap"><span class="badge">M</span>GenesisAccountMap</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.GenesisAccountMap.MapEntry"><span class="badge">M</span>GenesisAccountMap.MapEntry</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.Params"><span class="badge">M</span>Params</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.ProjectDoc"><span class="badge">M</span>ProjectDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.UpdateAgentDoc"><span class="badge">M</span>UpdateAgentDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.UpdateProjectStatusDoc"><span class="badge">M</span>UpdateProjectStatusDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.WithdrawFundsDoc"><span class="badge">M</span>WithdrawFundsDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.WithdrawalInfoDoc"><span class="badge">M</span>WithdrawalInfoDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.WithdrawalInfoDocs"><span class="badge">M</span>WithdrawalInfoDocs</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fproject%2fv1%2fgenesis.proto">ixo/project/v1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.project.v1.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fproject%2fv1%2fquery.proto">ixo/project/v1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.project.v1.QueryParamsRequest"><span class="badge">M</span>QueryParamsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryParamsResponse"><span class="badge">M</span>QueryParamsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectAccountsRequest"><span class="badge">M</span>QueryProjectAccountsRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectAccountsResponse"><span class="badge">M</span>QueryProjectAccountsResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectDocRequest"><span class="badge">M</span>QueryProjectDocRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectDocResponse"><span class="badge">M</span>QueryProjectDocResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectTxRequest"><span class="badge">M</span>QueryProjectTxRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.QueryProjectTxResponse"><span class="badge">M</span>QueryProjectTxResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.project.v1.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2fproject%2fv1%2ftx.proto">ixo/project/v1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateAgent"><span class="badge">M</span>MsgCreateAgent</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateAgentResponse"><span class="badge">M</span>MsgCreateAgentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateClaim"><span class="badge">M</span>MsgCreateClaim</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateClaimResponse"><span class="badge">M</span>MsgCreateClaimResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateEvaluation"><span class="badge">M</span>MsgCreateEvaluation</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateEvaluationResponse"><span class="badge">M</span>MsgCreateEvaluationResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateProject"><span class="badge">M</span>MsgCreateProject</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgCreateProjectResponse"><span class="badge">M</span>MsgCreateProjectResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateAgent"><span class="badge">M</span>MsgUpdateAgent</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateAgentResponse"><span class="badge">M</span>MsgUpdateAgentResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateProjectDoc"><span class="badge">M</span>MsgUpdateProjectDoc</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateProjectDocResponse"><span class="badge">M</span>MsgUpdateProjectDocResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateProjectStatus"><span class="badge">M</span>MsgUpdateProjectStatus</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgUpdateProjectStatusResponse"><span class="badge">M</span>MsgUpdateProjectStatusResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgWithdrawFunds"><span class="badge">M</span>MsgWithdrawFunds</a>
                </li>
              
                <li>
                  <a href="#ixo.project.v1.MsgWithdrawFundsResponse"><span class="badge">M</span>MsgWithdrawFundsResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.project.v1.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2ftoken%2fv1beta1%2ftoken.proto">ixo/token/v1beta1/token.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.token.v1beta1.Params"><span class="badge">M</span>Params</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.TokenDoc"><span class="badge">M</span>TokenDoc</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2ftoken%2fv1beta1%2fgenesis.proto">ixo/token/v1beta1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.token.v1beta1.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2ftoken%2fv1beta1%2fproposal.proto">ixo/token/v1beta1/proposal.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.token.v1beta1.InitializeTokenContract"><span class="badge">M</span>InitializeTokenContract</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2ftoken%2fv1beta1%2fquery.proto">ixo/token/v1beta1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenConfigRequest"><span class="badge">M</span>QueryTokenConfigRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenConfigResponse"><span class="badge">M</span>QueryTokenConfigResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenConfigResponse.MapEntry"><span class="badge">M</span>QueryTokenConfigResponse.MapEntry</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenDocRequest"><span class="badge">M</span>QueryTokenDocRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenDocResponse"><span class="badge">M</span>QueryTokenDocResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenListRequest"><span class="badge">M</span>QueryTokenListRequest</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.QueryTokenListResponse"><span class="badge">M</span>QueryTokenListResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#ixo.token.v1beta1.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#ixo%2ftoken%2fv1beta1%2ftx.proto">ixo/token/v1beta1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#ixo.token.v1beta1.MsgCreateToken"><span class="badge">M</span>MsgCreateToken</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.MsgCreateTokenResponse"><span class="badge">M</span>MsgCreateTokenResponse</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.MsgTransferToken"><span class="badge">M</span>MsgTransferToken</a>
                </li>
              
                <li>
                  <a href="#ixo.token.v1beta1.MsgTransferTokenResponse"><span class="badge">M</span>MsgTransferTokenResponse</a>
=======
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
>>>>>>> upstream/devel/ben-alpha
                </li>
              
              
              
              
                <li>
<<<<<<< HEAD
                  <a href="#ixo.token.v1beta1.Msg"><span class="badge">S</span>Msg</a>
=======
                  <a href="#project.Msg"><span class="badge">S</span>Msg</a>
>>>>>>> upstream/devel/ben-alpha
                </li>
              
            </ul>
          </li>
        
        <li><a href="#scalar-value-types">Scalar Value Types</a></li>
      </ul>
    </div>

    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/bonds/v1beta1/bonds.proto">ixo/bonds/v1beta1/bonds.proto</h2><a href="#title">Top</a>
=======
        <h2 id="bonds/bonds.proto">bonds/bonds.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.BaseOrder">BaseOrder</h3>
=======
        <h3 id="bonds.BaseOrder">BaseOrder</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.Batch">Batch</h3>
=======
        <h3 id="bonds.Batch">Batch</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.BuyOrder">BuyOrder</a></td>
=======
                  <td><a href="#bonds.BuyOrder">BuyOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sells</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.SellOrder">SellOrder</a></td>
=======
                  <td><a href="#bonds.SellOrder">SellOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>swaps</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.SwapOrder">SwapOrder</a></td>
=======
                  <td><a href="#bonds.SwapOrder">SwapOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.Bond">Bond</h3>
=======
        <h3 id="bonds.Bond">Bond</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.FunctionParam">FunctionParam</a></td>
=======
                  <td><a href="#bonds.FunctionParam">FunctionParam</a></td>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.BondDetails">BondDetails</h3>
=======
        <h3 id="bonds.BondDetails">BondDetails</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.BuyOrder">BuyOrder</h3>
=======
        <h3 id="bonds.BuyOrder">BuyOrder</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>BuyOrder defines a type for submitting a buy order on a bond, together with the maximum</p><p>amount of reserve tokens the buyer is willing to pay.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.BaseOrder">BaseOrder</a></td>
=======
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.FunctionParam">FunctionParam</h3>
=======
        <h3 id="bonds.FunctionParam">FunctionParam</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.Params">Params</h3>
=======
        <h3 id="bonds.Params">Params</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.SellOrder">SellOrder</h3>
=======
        <h3 id="bonds.SellOrder">SellOrder</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>SellOrder defines a type for submitting a sell order on a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.BaseOrder">BaseOrder</a></td>
=======
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.SwapOrder">SwapOrder</h3>
=======
        <h3 id="bonds.SwapOrder">SwapOrder</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>SwapOrder defines a type for submitting a swap order between two tokens on a bond.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>base_order</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.BaseOrder">BaseOrder</a></td>
=======
                  <td><a href="#bonds.BaseOrder">BaseOrder</a></td>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/bonds/v1beta1/genesis.proto">ixo/bonds/v1beta1/genesis.proto</h2><a href="#title">Top</a>
=======
        <h2 id="bonds/genesis.proto">bonds/genesis.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.GenesisState">GenesisState</h3>
=======
        <h3 id="bonds.GenesisState">GenesisState</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>GenesisState defines the bonds module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bonds</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Bond">Bond</a></td>
=======
                  <td><a href="#bonds.Bond">Bond</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>batches</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Batch">Batch</a></td>
=======
                  <td><a href="#bonds.Batch">Batch</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Params">Params</a></td>
=======
                  <td><a href="#bonds.Params">Params</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/bonds/v1beta1/query.proto">ixo/bonds/v1beta1/query.proto</h2><a href="#title">Top</a>
=======
        <h2 id="bonds/query.proto">bonds/query.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</h3>
=======
        <h3 id="bonds.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</h3>
=======
        <h3 id="bonds.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryAvailableReserveRequest">QueryAvailableReserveRequest</h3>
=======
        <h3 id="bonds.QueryAvailableReserveRequest">QueryAvailableReserveRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryAvailableReserveResponse">QueryAvailableReserveResponse</h3>
=======
        <h3 id="bonds.QueryAvailableReserveResponse">QueryAvailableReserveResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBatchRequest">QueryBatchRequest</h3>
=======
        <h3 id="bonds.QueryBatchRequest">QueryBatchRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBatchResponse">QueryBatchResponse</h3>
=======
        <h3 id="bonds.QueryBatchResponse">QueryBatchResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryBatchResponse is the response type for the Query/Batch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>batch</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Batch">Batch</a></td>
=======
                  <td><a href="#bonds.Batch">Batch</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondRequest">QueryBondRequest</h3>
=======
        <h3 id="bonds.QueryBondRequest">QueryBondRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondResponse">QueryBondResponse</h3>
=======
        <h3 id="bonds.QueryBondResponse">QueryBondResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryBondResponse is the response type for the Query/Bond RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bond</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Bond">Bond</a></td>
=======
                  <td><a href="#bonds.Bond">Bond</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondsDetailedRequest">QueryBondsDetailedRequest</h3>
=======
        <h3 id="bonds.QueryBondsDetailedRequest">QueryBondsDetailedRequest</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC method.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondsDetailedResponse">QueryBondsDetailedResponse</h3>
=======
        <h3 id="bonds.QueryBondsDetailedResponse">QueryBondsDetailedResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryBondsDetailedResponse is the response type for the Query/BondsDetailed RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>bonds_detailed</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.BondDetails">BondDetails</a></td>
=======
                  <td><a href="#bonds.BondDetails">BondDetails</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondsRequest">QueryBondsRequest</h3>
=======
        <h3 id="bonds.QueryBondsRequest">QueryBondsRequest</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryBondsRequest is the request type for the Query/Bonds RPC method.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBondsResponse">QueryBondsResponse</h3>
=======
        <h3 id="bonds.QueryBondsResponse">QueryBondsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBuyPriceRequest">QueryBuyPriceRequest</h3>
=======
        <h3 id="bonds.QueryBuyPriceRequest">QueryBuyPriceRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryBuyPriceResponse">QueryBuyPriceResponse</h3>
=======
        <h3 id="bonds.QueryBuyPriceResponse">QueryBuyPriceResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCurrentPriceRequest">QueryCurrentPriceRequest</h3>
=======
        <h3 id="bonds.QueryCurrentPriceRequest">QueryCurrentPriceRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCurrentPriceResponse">QueryCurrentPriceResponse</h3>
=======
        <h3 id="bonds.QueryCurrentPriceResponse">QueryCurrentPriceResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCurrentReserveRequest">QueryCurrentReserveRequest</h3>
=======
        <h3 id="bonds.QueryCurrentReserveRequest">QueryCurrentReserveRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCurrentReserveResponse">QueryCurrentReserveResponse</h3>
=======
        <h3 id="bonds.QueryCurrentReserveResponse">QueryCurrentReserveResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCustomPriceRequest">QueryCustomPriceRequest</h3>
=======
        <h3 id="bonds.QueryCustomPriceRequest">QueryCustomPriceRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryCustomPriceResponse">QueryCustomPriceResponse</h3>
=======
        <h3 id="bonds.QueryCustomPriceResponse">QueryCustomPriceResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryLastBatchRequest">QueryLastBatchRequest</h3>
=======
        <h3 id="bonds.QueryLastBatchRequest">QueryLastBatchRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryLastBatchResponse">QueryLastBatchResponse</h3>
=======
        <h3 id="bonds.QueryLastBatchResponse">QueryLastBatchResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryLastBatchResponse is the response type for the Query/LastBatch RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>last_batch</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Batch">Batch</a></td>
=======
                  <td><a href="#bonds.Batch">Batch</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryParamsRequest">QueryParamsRequest</h3>
=======
        <h3 id="bonds.QueryParamsRequest">QueryParamsRequest</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryParamsRequest is the request type for the Query/Params RPC method.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QueryParamsResponse">QueryParamsResponse</h3>
=======
        <h3 id="bonds.QueryParamsResponse">QueryParamsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryParamsResponse is the response type for the Query/Params RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.Params">Params</a></td>
=======
                  <td><a href="#bonds.Params">Params</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QuerySellReturnRequest">QuerySellReturnRequest</h3>
=======
        <h3 id="bonds.QuerySellReturnRequest">QuerySellReturnRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QuerySellReturnResponse">QuerySellReturnResponse</h3>
=======
        <h3 id="bonds.QuerySellReturnResponse">QuerySellReturnResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QuerySwapReturnRequest">QuerySwapReturnRequest</h3>
=======
        <h3 id="bonds.QuerySwapReturnRequest">QuerySwapReturnRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.QuerySwapReturnResponse">QuerySwapReturnResponse</h3>
=======
        <h3 id="bonds.QuerySwapReturnResponse">QuerySwapReturnResponse</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.Query">Query</h3>
=======
        <h3 id="bonds.Query">Query</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>Bonds</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryBondsRequest">QueryBondsRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryBondsResponse">QueryBondsResponse</a></td>
=======
                <td><a href="#bonds.QueryBondsRequest">QueryBondsRequest</a></td>
                <td><a href="#bonds.QueryBondsResponse">QueryBondsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Bonds returns all existing bonds.</p></td>
              </tr>
            
              <tr>
                <td>BondsDetailed</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryBondsDetailedRequest">QueryBondsDetailedRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryBondsDetailedResponse">QueryBondsDetailedResponse</a></td>
=======
                <td><a href="#bonds.QueryBondsDetailedRequest">QueryBondsDetailedRequest</a></td>
                <td><a href="#bonds.QueryBondsDetailedResponse">QueryBondsDetailedResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>BondsDetailed returns a list of all existing bonds with some details about their current state.</p></td>
              </tr>
            
              <tr>
                <td>Params</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryParamsResponse">QueryParamsResponse</a></td>
=======
                <td><a href="#bonds.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#bonds.QueryParamsResponse">QueryParamsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Params queries the paramaters of x/bonds module.</p></td>
              </tr>
            
              <tr>
                <td>Bond</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryBondRequest">QueryBondRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryBondResponse">QueryBondResponse</a></td>
=======
                <td><a href="#bonds.QueryBondRequest">QueryBondRequest</a></td>
                <td><a href="#bonds.QueryBondResponse">QueryBondResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Bond queries info of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>Batch</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryBatchRequest">QueryBatchRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryBatchResponse">QueryBatchResponse</a></td>
=======
                <td><a href="#bonds.QueryBatchRequest">QueryBatchRequest</a></td>
                <td><a href="#bonds.QueryBatchResponse">QueryBatchResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Batch queries info of a specific bond&#39;s current batch.</p></td>
              </tr>
            
              <tr>
                <td>LastBatch</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryLastBatchRequest">QueryLastBatchRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryLastBatchResponse">QueryLastBatchResponse</a></td>
=======
                <td><a href="#bonds.QueryLastBatchRequest">QueryLastBatchRequest</a></td>
                <td><a href="#bonds.QueryLastBatchResponse">QueryLastBatchResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>LastBatch queries info of a specific bond&#39;s last batch.</p></td>
              </tr>
            
              <tr>
                <td>CurrentPrice</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryCurrentPriceRequest">QueryCurrentPriceRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryCurrentPriceResponse">QueryCurrentPriceResponse</a></td>
=======
                <td><a href="#bonds.QueryCurrentPriceRequest">QueryCurrentPriceRequest</a></td>
                <td><a href="#bonds.QueryCurrentPriceResponse">QueryCurrentPriceResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CurrentPrice queries the current price/s of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>CurrentReserve</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryCurrentReserveRequest">QueryCurrentReserveRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryCurrentReserveResponse">QueryCurrentReserveResponse</a></td>
=======
                <td><a href="#bonds.QueryCurrentReserveRequest">QueryCurrentReserveRequest</a></td>
                <td><a href="#bonds.QueryCurrentReserveResponse">QueryCurrentReserveResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CurrentReserve queries the current balance/s of the reserve pool for a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>AvailableReserve</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryAvailableReserveRequest">QueryAvailableReserveRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryAvailableReserveResponse">QueryAvailableReserveResponse</a></td>
=======
                <td><a href="#bonds.QueryAvailableReserveRequest">QueryAvailableReserveRequest</a></td>
                <td><a href="#bonds.QueryAvailableReserveResponse">QueryAvailableReserveResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>AvailableReserve queries current available balance/s of the reserve pool for a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>CustomPrice</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryCustomPriceRequest">QueryCustomPriceRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryCustomPriceResponse">QueryCustomPriceResponse</a></td>
=======
                <td><a href="#bonds.QueryCustomPriceRequest">QueryCustomPriceRequest</a></td>
                <td><a href="#bonds.QueryCustomPriceResponse">QueryCustomPriceResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CustomPrice queries price/s of a specific bond at a specific supply.</p></td>
              </tr>
            
              <tr>
                <td>BuyPrice</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryBuyPriceRequest">QueryBuyPriceRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryBuyPriceResponse">QueryBuyPriceResponse</a></td>
=======
                <td><a href="#bonds.QueryBuyPriceRequest">QueryBuyPriceRequest</a></td>
                <td><a href="#bonds.QueryBuyPriceResponse">QueryBuyPriceResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>BuyPrice queries price/s of buying an amount of tokens from a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>SellReturn</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QuerySellReturnRequest">QuerySellReturnRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QuerySellReturnResponse">QuerySellReturnResponse</a></td>
=======
                <td><a href="#bonds.QuerySellReturnRequest">QuerySellReturnRequest</a></td>
                <td><a href="#bonds.QuerySellReturnResponse">QuerySellReturnResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>SellReturn queries return/s on selling an amount of tokens of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>SwapReturn</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QuerySwapReturnRequest">QuerySwapReturnRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QuerySwapReturnResponse">QuerySwapReturnResponse</a></td>
=======
                <td><a href="#bonds.QuerySwapReturnRequest">QuerySwapReturnRequest</a></td>
                <td><a href="#bonds.QuerySwapReturnResponse">QuerySwapReturnResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond.</p></td>
              </tr>
            
              <tr>
                <td>AlphaMaximums</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</a></td>
                <td><a href="#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</a></td>
=======
                <td><a href="#bonds.QueryAlphaMaximumsRequest">QueryAlphaMaximumsRequest</a></td>
                <td><a href="#bonds.QueryAlphaMaximumsResponse">QueryAlphaMaximumsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/bonds/v1beta1/tx.proto">ixo/bonds/v1beta1/tx.proto</h2><a href="#title">Top</a>
=======
        <h2 id="bonds/tx.proto">bonds/tx.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgBuy">MsgBuy</h3>
=======
        <h3 id="bonds.MsgBuy">MsgBuy</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>buyer_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgBuyResponse">MsgBuyResponse</h3>
=======
        <h3 id="bonds.MsgBuyResponse">MsgBuyResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgBuyResponse defines the Msg/Buy response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgCreateBond">MsgCreateBond</h3>
=======
        <h3 id="bonds.MsgCreateBond">MsgCreateBond</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.bonds.v1beta1.FunctionParam">FunctionParam</a></td>
=======
                  <td><a href="#bonds.FunctionParam">FunctionParam</a></td>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>creator_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgCreateBondResponse">MsgCreateBondResponse</h3>
=======
        <h3 id="bonds.MsgCreateBondResponse">MsgCreateBondResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateBondResponse defines the Msg/CreateBond response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgEditBond">MsgEditBond</h3>
=======
        <h3 id="bonds.MsgEditBond">MsgEditBond</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>editor_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgEditBondResponse">MsgEditBondResponse</h3>
=======
        <h3 id="bonds.MsgEditBondResponse">MsgEditBondResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgEditBondResponse defines the Msg/EditBond response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgMakeOutcomePayment">MsgMakeOutcomePayment</h3>
=======
        <h3 id="bonds.MsgMakeOutcomePayment">MsgMakeOutcomePayment</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>sender_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</h3>
=======
        <h3 id="bonds.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSell">MsgSell</h3>
=======
        <h3 id="bonds.MsgSell">MsgSell</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>seller_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSellResponse">MsgSellResponse</h3>
=======
        <h3 id="bonds.MsgSellResponse">MsgSellResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgSellResponse defines the Msg/Sell response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSetNextAlpha">MsgSetNextAlpha</h3>
=======
        <h3 id="bonds.MsgSetNextAlpha">MsgSetNextAlpha</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>editor_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</h3>
=======
        <h3 id="bonds.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p></p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSwap">MsgSwap</h3>
=======
        <h3 id="bonds.MsgSwap">MsgSwap</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>swapper_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgSwapResponse">MsgSwapResponse</h3>
=======
        <h3 id="bonds.MsgSwapResponse">MsgSwapResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgSwapResponse defines the Msg/Swap response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgUpdateBondState">MsgUpdateBondState</h3>
=======
        <h3 id="bonds.MsgUpdateBondState">MsgUpdateBondState</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>editor_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</h3>
=======
        <h3 id="bonds.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgWithdrawReserve">MsgWithdrawReserve</h3>
=======
        <h3 id="bonds.MsgWithdrawReserve">MsgWithdrawReserve</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>withdrawer_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</h3>
=======
        <h3 id="bonds.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgWithdrawShare">MsgWithdrawShare</h3>
=======
        <h3 id="bonds.MsgWithdrawShare">MsgWithdrawShare</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>recipient_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.MsgWithdrawShareResponse">MsgWithdrawShareResponse</h3>
=======
        <h3 id="bonds.MsgWithdrawShareResponse">MsgWithdrawShareResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.</p>

        

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.bonds.v1beta1.Msg">Msg</h3>
=======
        <h3 id="bonds.Msg">Msg</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Msg defines the bonds Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateBond</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgCreateBond">MsgCreateBond</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgCreateBondResponse">MsgCreateBondResponse</a></td>
=======
                <td><a href="#bonds.MsgCreateBond">MsgCreateBond</a></td>
                <td><a href="#bonds.MsgCreateBondResponse">MsgCreateBondResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateBond defines a method for creating a bond.</p></td>
              </tr>
            
              <tr>
                <td>EditBond</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgEditBond">MsgEditBond</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgEditBondResponse">MsgEditBondResponse</a></td>
=======
                <td><a href="#bonds.MsgEditBond">MsgEditBond</a></td>
                <td><a href="#bonds.MsgEditBondResponse">MsgEditBondResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>EditBond defines a method for editing a bond.</p></td>
              </tr>
            
              <tr>
                <td>SetNextAlpha</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgSetNextAlpha">MsgSetNextAlpha</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</a></td>
=======
                <td><a href="#bonds.MsgSetNextAlpha">MsgSetNextAlpha</a></td>
                <td><a href="#bonds.MsgSetNextAlphaResponse">MsgSetNextAlphaResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>SetNextAlpha defines a method for editing a bond&#39;s alpha parameter.</p></td>
              </tr>
            
              <tr>
                <td>UpdateBondState</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgUpdateBondState">MsgUpdateBondState</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</a></td>
=======
                <td><a href="#bonds.MsgUpdateBondState">MsgUpdateBondState</a></td>
                <td><a href="#bonds.MsgUpdateBondStateResponse">MsgUpdateBondStateResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>UpdateBondState defines a method for updating a bond&#39;s current state.</p></td>
              </tr>
            
              <tr>
                <td>Buy</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgBuy">MsgBuy</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgBuyResponse">MsgBuyResponse</a></td>
=======
                <td><a href="#bonds.MsgBuy">MsgBuy</a></td>
                <td><a href="#bonds.MsgBuyResponse">MsgBuyResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Buy defines a method for buying from a bond.</p></td>
              </tr>
            
              <tr>
                <td>Sell</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgSell">MsgSell</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgSellResponse">MsgSellResponse</a></td>
=======
                <td><a href="#bonds.MsgSell">MsgSell</a></td>
                <td><a href="#bonds.MsgSellResponse">MsgSellResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Sell defines a method for selling from a bond.</p></td>
              </tr>
            
              <tr>
                <td>Swap</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgSwap">MsgSwap</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgSwapResponse">MsgSwapResponse</a></td>
=======
                <td><a href="#bonds.MsgSwap">MsgSwap</a></td>
                <td><a href="#bonds.MsgSwapResponse">MsgSwapResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>Swap defines a method for swapping from one reserve bond token to another.</p></td>
              </tr>
            
              <tr>
                <td>MakeOutcomePayment</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgMakeOutcomePayment">MsgMakeOutcomePayment</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</a></td>
=======
                <td><a href="#bonds.MsgMakeOutcomePayment">MsgMakeOutcomePayment</a></td>
                <td><a href="#bonds.MsgMakeOutcomePaymentResponse">MsgMakeOutcomePaymentResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>MakeOutcomePayment defines a method for making an outcome payment to a bond.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawShare</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgWithdrawShare">MsgWithdrawShare</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgWithdrawShareResponse">MsgWithdrawShareResponse</a></td>
=======
                <td><a href="#bonds.MsgWithdrawShare">MsgWithdrawShare</a></td>
                <td><a href="#bonds.MsgWithdrawShareResponse">MsgWithdrawShareResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawReserve</td>
<<<<<<< HEAD
                <td><a href="#ixo.bonds.v1beta1.MsgWithdrawReserve">MsgWithdrawReserve</a></td>
                <td><a href="#ixo.bonds.v1beta1.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</a></td>
=======
                <td><a href="#bonds.MsgWithdrawReserve">MsgWithdrawReserve</a></td>
                <td><a href="#bonds.MsgWithdrawReserveResponse">MsgWithdrawReserveResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>WithdrawReserve defines a method for withdrawing reserve from a bond.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/entity/v1beta1/entity.proto">ixo/entity/v1beta1/entity.proto</h2><a href="#title">Top</a>
=======
        <h2 id="did/did.proto">did/did.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.entity.v1beta1.EntityDoc">EntityDoc</h3>
        <p>ProjectDoc defines a project (or entity) type with all of its parameters.</p>

        

        
      
        <h3 id="ixo.entity.v1beta1.Params">Params</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>NftContractAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>NftContractMinter</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/entity/v1beta1/genesis.proto">ixo/entity/v1beta1/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.entity.v1beta1.GenesisState">GenesisState</h3>
        <p>GenesisState defines the project module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entity_docs</td>
                  <td><a href="#ixo.entity.v1beta1.EntityDoc">EntityDoc</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
                  <td><a href="#ixo.entity.v1beta1.Params">Params</a></td>
                  <td></td>
                  <td><p>repeated GenesisAccountMap account_maps       = 2 [(gogoproto.nullable) = false, (gogoproto.moretags) = &#34;yaml:\&#34;account_maps\&#34;&#34;]; </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/entity/v1beta1/proposal.proto">ixo/entity/v1beta1/proposal.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.entity.v1beta1.InitializeNftContract">InitializeNftContract</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>NftContractCodeId</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>NftMinterAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/entity/v1beta1/query.proto">ixo/entity/v1beta1/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.entity.v1beta1.QueryEntityConfigRequest">QueryEntityConfigRequest</h3>
        <p></p>

        

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityConfigResponse">QueryEntityConfigResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
                  <td><a href="#ixo.entity.v1beta1.QueryEntityConfigResponse.MapEntry">QueryEntityConfigResponse.MapEntry</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityConfigResponse.MapEntry">QueryEntityConfigResponse.MapEntry</h3>
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

          

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityDocRequest">QueryEntityDocRequest</h3>
        <p>QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entity_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityDocResponse">QueryEntityDocResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityListRequest">QueryEntityListRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entity_type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>entity_status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.QueryEntityListResponse">QueryEntityListResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        

        
      

      

      

      
        <h3 id="ixo.entity.v1beta1.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>EntityList</td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityListRequest">QueryEntityListRequest</a></td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityListResponse">QueryEntityListResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>EntityDoc</td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityDocRequest">QueryEntityDocRequest</a></td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityDocResponse">QueryEntityDocResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>EntityConfig</td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityConfigRequest">QueryEntityConfigRequest</a></td>
                <td><a href="#ixo.entity.v1beta1.QueryEntityConfigResponse">QueryEntityConfigResponse</a></td>
                <td><p></p></td>
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
                <td>EntityList</td>
                <td>GET</td>
                <td>/ixo/entity</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>EntityDoc</td>
                <td>GET</td>
                <td>/ixo/entity/{entity_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>EntityConfig</td>
                <td>GET</td>
                <td>/ixo/entity/config</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="ixo/iid/v1beta1/iid.proto">ixo/iid/v1beta1/iid.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.iid.v1beta1.AccordedRight">AccordedRight</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>mechanism</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>message</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>service</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.Context">Context</h3>
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
                  <td>val</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.IidDocument">IidDocument</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>context</td>
                  <td><a href="#ixo.iid.v1beta1.Context">Context</a></td>
                  <td>repeated</td>
                  <td><p>@context is spec for did document. </p></td>
                </tr>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>id represents the id for the did document. </p></td>
                </tr>
              
                <tr>
                  <td>controller</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>A DID controller is an entity that is authorized to make changes to a DID document.
cfr. https://www.w3.org/TR/did-core/#did-controller </p></td>
                </tr>
              
                <tr>
                  <td>verificationMethod</td>
                  <td><a href="#ixo.iid.v1beta1.VerificationMethod">VerificationMethod</a></td>
                  <td>repeated</td>
                  <td><p>A DID document can express verification methods, 
such as cryptographic public keys, which can be used 
to authenticate or authorize interactions with the DID subject or associated parties.
https://www.w3.org/TR/did-core/#verification-methods </p></td>
                </tr>
              
                <tr>
                  <td>service</td>
                  <td><a href="#ixo.iid.v1beta1.Service">Service</a></td>
                  <td>repeated</td>
                  <td><p>Services are used in DID documents to express ways of communicating 
with the DID subject or associated entities.
https://www.w3.org/TR/did-core/#services </p></td>
                </tr>
              
                <tr>
                  <td>authentication</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>NOTE: below this line there are the relationships
Authentication represents public key associated with the did document.
cfr. https://www.w3.org/TR/did-core/#authentication </p></td>
                </tr>
              
                <tr>
                  <td>assertionMethod</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>Used to specify how the DID subject is expected to express claims, 
such as for the purposes of issuing a Verifiable Credential.
cfr. https://www.w3.org/TR/did-core/#assertion </p></td>
                </tr>
              
                <tr>
                  <td>keyAgreement</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>used to specify how an entity can generate encryption material 
in order to transmit confidential information intended for the DID subject.
https://www.w3.org/TR/did-core/#key-agreement </p></td>
                </tr>
              
                <tr>
                  <td>capabilityInvocation</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>Used to specify a verification method that might be used by the DID subject
to invoke a cryptographic capability, such as the authorization 
to update the DID Document.
https://www.w3.org/TR/did-core/#capability-invocation </p></td>
                </tr>
              
                <tr>
                  <td>capabilityDelegation</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>Used to specify a mechanism that might be used by the DID subject 
to delegate a cryptographic capability to another party.
https://www.w3.org/TR/did-core/#capability-delegation </p></td>
                </tr>
              
                <tr>
                  <td>linkedResource</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedResource">LinkedResource</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>accordedRight</td>
                  <td><a href="#ixo.iid.v1beta1.AccordedRight">AccordedRight</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>linkedEntity</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedEntity">LinkedEntity</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>alsoKnownAs</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.IidMetadata">IidMetadata</h3>
        <p>DidMetadata defines metadata associated to a did document such as </p><p>the status of the DID document</p>

        
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
                  <td>versionId</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>created</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>updated</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>deactivated</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>entityType</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>startDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>endDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>status</td>
                  <td><a href="#int32">int32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>stage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>relayerNode</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>verifiableCredential</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>credentials</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.LinkedEntity">LinkedEntity</h3>
        <p></p>

        
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
                  <td>relationship</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.LinkedResource">LinkedResource</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>id</td>
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
                  <td>mediaType</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>serviceEndpoint</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>proof</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>encrypted</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>right</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.Service">Service</h3>
        <p>Service defines how to find data associated with a identifer</p>

        
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
                  <td>type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>serviceEndpoint</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.VerificationMethod">VerificationMethod</h3>
        <p></p>

        
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
                  <td>type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>controller</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>blockchainAccountID</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>publicKeyHex</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>publicKeyMultibase</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/iid/v1beta1/tx.proto">ixo/iid/v1beta1/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.iid.v1beta1.MsgAddAccordedRight">MsgAddAccordedRight</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>accordedRight</td>
                  <td><a href="#ixo.iid.v1beta1.AccordedRight">AccordedRight</a></td>
                  <td></td>
                  <td><p>the Accorded right to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddAccordedRightResponse">MsgAddAccordedRightResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddController">MsgAddController</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did of the document </p></td>
                </tr>
              
                <tr>
                  <td>controller_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did to add as a controller of the did document </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddControllerResponse">MsgAddControllerResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddIidContext">MsgAddIidContext</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>context</td>
                  <td><a href="#ixo.iid.v1beta1.Context">Context</a></td>
                  <td></td>
                  <td><p>the context to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddIidContextResponse">MsgAddIidContextResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddLinkedEntity">MsgAddLinkedEntity</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the iid </p></td>
                </tr>
              
                <tr>
                  <td>linkedEntity</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedEntity">LinkedEntity</a></td>
                  <td></td>
                  <td><p>the entity to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddLinkedEntityResponse">MsgAddLinkedEntityResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddLinkedResource">MsgAddLinkedResource</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>linkedResource</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedResource">LinkedResource</a></td>
                  <td></td>
                  <td><p>the verification to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddLinkedResourceResponse">MsgAddLinkedResourceResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddService">MsgAddService</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>service_data</td>
                  <td><a href="#ixo.iid.v1beta1.Service">Service</a></td>
                  <td></td>
                  <td><p>the service data to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddServiceResponse">MsgAddServiceResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddVerification">MsgAddVerification</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>verification</td>
                  <td><a href="#ixo.iid.v1beta1.Verification">Verification</a></td>
                  <td></td>
                  <td><p>the verification to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgAddVerificationResponse">MsgAddVerificationResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgCreateIidDocument">MsgCreateIidDocument</h3>
        <p>MsgCreateDidDocument defines a SDK message for creating a new did.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>controllers</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>the list of controller DIDs </p></td>
                </tr>
              
                <tr>
                  <td>context</td>
                  <td><a href="#ixo.iid.v1beta1.Context">Context</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>verifications</td>
                  <td><a href="#ixo.iid.v1beta1.Verification">Verification</a></td>
                  <td>repeated</td>
                  <td><p>the list of verification methods and relationships </p></td>
                </tr>
              
                <tr>
                  <td>services</td>
                  <td><a href="#ixo.iid.v1beta1.Service">Service</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>accordedRight</td>
                  <td><a href="#ixo.iid.v1beta1.AccordedRight">AccordedRight</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>linkedResource</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedResource">LinkedResource</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>linkedEntity</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedEntity">LinkedEntity</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgCreateIidDocumentResponse">MsgCreateIidDocumentResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeactivateIID">MsgDeactivateIID</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>state</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeactivateIIDResponse">MsgDeactivateIIDResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteAccordedRight">MsgDeleteAccordedRight</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>right_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the service id </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteAccordedRightResponse">MsgDeleteAccordedRightResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteController">MsgDeleteController</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did of the document </p></td>
                </tr>
              
                <tr>
                  <td>controller_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did to remove from the list of controllers of the did document </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteControllerResponse">MsgDeleteControllerResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteIidContext">MsgDeleteIidContext</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>contextKey</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the context key </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteIidContextResponse">MsgDeleteIidContextResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteLinkedEntity">MsgDeleteLinkedEntity</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the iid </p></td>
                </tr>
              
                <tr>
                  <td>entity_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the entity id </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse">MsgDeleteLinkedEntityResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteLinkedResource">MsgDeleteLinkedResource</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>resource_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the service id </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse">MsgDeleteLinkedResourceResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteService">MsgDeleteService</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>service_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the service id </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgDeleteServiceResponse">MsgDeleteServiceResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgRevokeVerification">MsgRevokeVerification</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>method_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the verification method id </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgRevokeVerificationResponse">MsgRevokeVerificationResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgSetVerificationRelationships">MsgSetVerificationRelationships</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>method_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the verification method id </p></td>
                </tr>
              
                <tr>
                  <td>relationships</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>the list of relationships to set </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse">MsgSetVerificationRelationshipsResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgUpdateIidDocument">MsgUpdateIidDocument</h3>
        <p>MsgUpdateDidDocument replace an existing did document with a new version</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>doc</td>
                  <td><a href="#ixo.iid.v1beta1.IidDocument">IidDocument</a></td>
                  <td></td>
                  <td><p>the did document to replace </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgUpdateIidDocumentResponse">MsgUpdateIidDocumentResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.MsgUpdateIidMeta">MsgUpdateIidMeta</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did </p></td>
                </tr>
              
                <tr>
                  <td>meta</td>
                  <td><a href="#ixo.iid.v1beta1.IidMetadata">IidMetadata</a></td>
                  <td></td>
                  <td><p>the context to add </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>address of the account signing the message </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.MsgUpdateIidMetaResponse">MsgUpdateIidMetaResponse</h3>
        <p></p>

        

        
      
        <h3 id="ixo.iid.v1beta1.Verification">Verification</h3>
        <p>Verification is a message that allows to assign a verification method</p><p>to one or more verification relationships</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>relationships</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>verificationRelationships defines which relationships
are allowed to use the verification method

relationships that the method is allowed into. </p></td>
                </tr>
              
                <tr>
                  <td>method</td>
                  <td><a href="#ixo.iid.v1beta1.VerificationMethod">VerificationMethod</a></td>
                  <td></td>
                  <td><p>public key associated with the did document. </p></td>
                </tr>
              
                <tr>
                  <td>context</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>additional contexts (json ld schemas) </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="ixo.iid.v1beta1.Msg">Msg</h3>
        <p>Msg defines the identity Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateIidDocument</td>
                <td><a href="#ixo.iid.v1beta1.MsgCreateIidDocument">MsgCreateIidDocument</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgCreateIidDocumentResponse">MsgCreateIidDocumentResponse</a></td>
                <td><p>CreateDidDocument defines a method for creating a new identity.</p></td>
              </tr>
            
              <tr>
                <td>UpdateIidDocument</td>
                <td><a href="#ixo.iid.v1beta1.MsgUpdateIidDocument">MsgUpdateIidDocument</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse">MsgUpdateIidDocumentResponse</a></td>
                <td><p>UpdateDidDocument defines a method for updating an identity.</p></td>
              </tr>
            
              <tr>
                <td>AddVerification</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddVerification">MsgAddVerification</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddVerificationResponse">MsgAddVerificationResponse</a></td>
                <td><p>AddVerificationMethod adds a new verification method</p></td>
              </tr>
            
              <tr>
                <td>RevokeVerification</td>
                <td><a href="#ixo.iid.v1beta1.MsgRevokeVerification">MsgRevokeVerification</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgRevokeVerificationResponse">MsgRevokeVerificationResponse</a></td>
                <td><p>RevokeVerification remove the verification method and all associated verification Relations</p></td>
              </tr>
            
              <tr>
                <td>SetVerificationRelationships</td>
                <td><a href="#ixo.iid.v1beta1.MsgSetVerificationRelationships">MsgSetVerificationRelationships</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse">MsgSetVerificationRelationshipsResponse</a></td>
                <td><p>SetVerificationRelationships overwrite current verification relationships</p></td>
              </tr>
            
              <tr>
                <td>AddService</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddService">MsgAddService</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddServiceResponse">MsgAddServiceResponse</a></td>
                <td><p>AddService add a new service</p></td>
              </tr>
            
              <tr>
                <td>DeleteService</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteService">MsgDeleteService</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteServiceResponse">MsgDeleteServiceResponse</a></td>
                <td><p>DeleteService delete an existing service</p></td>
              </tr>
            
              <tr>
                <td>AddController</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddController">MsgAddController</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddControllerResponse">MsgAddControllerResponse</a></td>
                <td><p>AddService add a new service</p></td>
              </tr>
            
              <tr>
                <td>DeleteController</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteController">MsgDeleteController</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteControllerResponse">MsgDeleteControllerResponse</a></td>
                <td><p>DeleteService delete an existing service</p></td>
              </tr>
            
              <tr>
                <td>AddLinkedResource</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddLinkedResource">MsgAddLinkedResource</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddLinkedResourceResponse">MsgAddLinkedResourceResponse</a></td>
                <td><p>Add / Delete Linked Resource</p></td>
              </tr>
            
              <tr>
                <td>DeleteLinkedResource</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteLinkedResource">MsgDeleteLinkedResource</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse">MsgDeleteLinkedResourceResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>AddLinkedEntity</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddLinkedEntity">MsgAddLinkedEntity</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddLinkedEntityResponse">MsgAddLinkedEntityResponse</a></td>
                <td><p>Add / Delete Linked Entity</p></td>
              </tr>
            
              <tr>
                <td>DeleteLinkedEntity</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteLinkedEntity">MsgDeleteLinkedEntity</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse">MsgDeleteLinkedEntityResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>AddAccordedRight</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddAccordedRight">MsgAddAccordedRight</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddAccordedRightResponse">MsgAddAccordedRightResponse</a></td>
                <td><p>Add / Delete Accorded Right</p></td>
              </tr>
            
              <tr>
                <td>DeleteAccordedRight</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteAccordedRight">MsgDeleteAccordedRight</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse">MsgDeleteAccordedRightResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>AddIidContext</td>
                <td><a href="#ixo.iid.v1beta1.MsgAddIidContext">MsgAddIidContext</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgAddIidContextResponse">MsgAddIidContextResponse</a></td>
                <td><p>Add / Delete Context</p></td>
              </tr>
            
              <tr>
                <td>DeactivateIID</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeactivateIID">MsgDeactivateIID</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeactivateIIDResponse">MsgDeactivateIIDResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>DeleteIidContext</td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteIidContext">MsgDeleteIidContext</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgDeleteIidContextResponse">MsgDeleteIidContextResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>UpdateMetaData</td>
                <td><a href="#ixo.iid.v1beta1.MsgUpdateIidMeta">MsgUpdateIidMeta</a></td>
                <td><a href="#ixo.iid.v1beta1.MsgUpdateIidMetaResponse">MsgUpdateIidMetaResponse</a></td>
                <td><p>Update META</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
        <h2 id="ixo/entity/v1beta1/tx.proto">ixo/entity/v1beta1/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.entity.v1beta1.MsgCreateEntity">MsgCreateEntity</h3>
        <p>MsgCreateEntity defines a message for creating a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entityType</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>An Entity Type as defined by the implementer </p></td>
                </tr>
              
                <tr>
                  <td>entityStatus</td>
                  <td><a href="#int32">int32</a></td>
                  <td></td>
                  <td><p>Status of the Entity as defined by the implementer and interpreted by Client applications </p></td>
                </tr>
              
                <tr>
                  <td>controller</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>the list of controller DIDs </p></td>
                </tr>
              
                <tr>
                  <td>context</td>
                  <td><a href="#ixo.iid.v1beta1.Context">ixo.iid.v1beta1.Context</a></td>
                  <td>repeated</td>
                  <td><p>JSON-LD contexts </p></td>
                </tr>
              
                <tr>
                  <td>verification</td>
                  <td><a href="#ixo.iid.v1beta1.Verification">ixo.iid.v1beta1.Verification</a></td>
                  <td>repeated</td>
                  <td><p>Verification Methods and Verification Relationships </p></td>
                </tr>
              
                <tr>
                  <td>service</td>
                  <td><a href="#ixo.iid.v1beta1.Service">ixo.iid.v1beta1.Service</a></td>
                  <td>repeated</td>
                  <td><p>Service endpoints </p></td>
                </tr>
              
                <tr>
                  <td>accordedRight</td>
                  <td><a href="#ixo.iid.v1beta1.AccordedRight">ixo.iid.v1beta1.AccordedRight</a></td>
                  <td>repeated</td>
                  <td><p>Legal or Electronic Rights and associated Object Capabilities </p></td>
                </tr>
              
                <tr>
                  <td>linkedResource</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedResource">ixo.iid.v1beta1.LinkedResource</a></td>
                  <td>repeated</td>
                  <td><p>Digital resources associated with the Subject </p></td>
                </tr>
              
                <tr>
                  <td>linkedEntity</td>
                  <td><a href="#ixo.iid.v1beta1.LinkedEntity">ixo.iid.v1beta1.LinkedEntity</a></td>
                  <td>repeated</td>
                  <td><p>DID of a linked Entity and its relationship with the Subject </p></td>
                </tr>
              
                <tr>
                  <td>deactivated</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p>Operational status of the Entity </p></td>
                </tr>
              
                <tr>
                  <td>startDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p>Start Date of the Entity as defined by the implementer and interpreted by Client applications

address of the account signing the message </p></td>
                </tr>
              
                <tr>
                  <td>endDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p>End Date of the Entity as defined by the implementer and interpreted by Client applications

address of the account signing the message </p></td>
                </tr>
              
                <tr>
                  <td>stage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>State of the Entity as defined by the implementer and interpreted by Client applications </p></td>
                </tr>
              
                <tr>
                  <td>relayerNode</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>DID of the operator through which the Entity was created </p></td>
                </tr>
              
                <tr>
                  <td>verificationStatus</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>Public proof that the Entity is verified </p></td>
                </tr>
              
                <tr>
                  <td>verifiableCredential</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p>Content ID or Hash of public Verifiable Credentials associated with the  subject </p></td>
                </tr>
              
                <tr>
                  <td>ownerDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>Owner of the Entity NFT | The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>ownerAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid address used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>data</td>
                  <td><a href="#bytes">bytes</a></td>
                  <td></td>
                  <td><p>Extention data </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.MsgCreateEntityResponse">MsgCreateEntityResponse</h3>
        <p>MsgCreateProjectResponse defines the Msg/CreateProject response type.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entityId</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>entityType</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>entityStatus</td>
                  <td><a href="#int32">int32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.MsgTransferEntity">MsgTransferEntity</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>entityDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>ownerDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>ownerAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>recipientDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.MsgTransferEntityResponse">MsgTransferEntityResponse</h3>
        <p>MsgUpdateProjectStatusResponse defines the Msg/UpdateEntityStatus response type.</p>

        

        
      
        <h3 id="ixo.entity.v1beta1.MsgUpdateEntity">MsgUpdateEntity</h3>
        <p>MsgUpdateEntityStatus defines a message for updating a entity's current status.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>status</td>
                  <td><a href="#int32">int32</a></td>
                  <td></td>
                  <td><p>The status of the entity. Should represent an enum in the client. </p></td>
                </tr>
              
                <tr>
                  <td>deactivated</td>
                  <td><a href="#bool">bool</a></td>
                  <td></td>
                  <td><p>Whether this entity is enabled ot not, basically a soft delete. </p></td>
                </tr>
              
                <tr>
                  <td>startDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p>refer to iid module for more information

address of the account signing the message </p></td>
                </tr>
              
                <tr>
                  <td>endDate</td>
                  <td><a href="#google.protobuf.Timestamp">google.protobuf.Timestamp</a></td>
                  <td></td>
                  <td><p>refer to iid module for more information

address of the account signing the message </p></td>
                </tr>
              
                <tr>
                  <td>stage</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>refer to iid module meta data for more information </p></td>
                </tr>
              
                <tr>
                  <td>relayerNode</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>refer to iid module for more information </p></td>
                </tr>
              
                <tr>
                  <td>verifiableCredential</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>refer to iid module for more information </p></td>
                </tr>
              
                <tr>
                  <td>controllerDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>controllerAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.entity.v1beta1.MsgUpdateEntityResponse">MsgUpdateEntityResponse</h3>
        <p>MsgUpdateProjectStatusResponse defines the Msg/UpdateEntityStatus response type.</p>

        

        
      

      

      

      
        <h3 id="ixo.entity.v1beta1.Msg">Msg</h3>
        <p>Msg defines the project Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateEntity</td>
                <td><a href="#ixo.entity.v1beta1.MsgCreateEntity">MsgCreateEntity</a></td>
                <td><a href="#ixo.entity.v1beta1.MsgCreateEntityResponse">MsgCreateEntityResponse</a></td>
                <td><p>CreateProject defines a method for creating a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateEntity</td>
                <td><a href="#ixo.entity.v1beta1.MsgUpdateEntity">MsgUpdateEntity</a></td>
                <td><a href="#ixo.entity.v1beta1.MsgUpdateEntityResponse">MsgUpdateEntityResponse</a></td>
                <td><p>UpdateEntityStatus defines a method for updating a entity&#39;s current status.</p></td>
              </tr>
            
              <tr>
                <td>TransferEntity</td>
                <td><a href="#ixo.entity.v1beta1.MsgTransferEntity">MsgTransferEntity</a></td>
                <td><a href="#ixo.entity.v1beta1.MsgTransferEntityResponse">MsgTransferEntityResponse</a></td>
                <td><p>Transfers an entity and its nft to the recipient</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
        <h2 id="ixo/iid/v1beta1/event.proto">ixo/iid/v1beta1/event.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.iid.v1beta1.IidDocumentCreatedEvent">IidDocumentCreatedEvent</h3>
        <p>DidDocumentCreatedEvent is an event triggered on a DID document creation</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did being created </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the signer account creating the did </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.IidDocumentUpdatedEvent">IidDocumentUpdatedEvent</h3>
        <p>DidDocumentUpdatedEvent is an event triggered on a DID document update</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the did being updated </p></td>
                </tr>
              
                <tr>
                  <td>signer</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the signer account of the change </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/iid/v1beta1/genesis.proto">ixo/iid/v1beta1/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.iid.v1beta1.GenesisState">GenesisState</h3>
        <p>GenesisState defines the did module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>iid_docs</td>
                  <td><a href="#ixo.iid.v1beta1.IidDocument">IidDocument</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>iid_meta</td>
                  <td><a href="#ixo.iid.v1beta1.IidMetadata">IidMetadata</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/iid/v1beta1/query.proto">ixo/iid/v1beta1/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.iid.v1beta1.QueryIidDocumentRequest">QueryIidDocumentRequest</h3>
        <p>QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>status enables to query for validators matching a given status. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.QueryIidDocumentResponse">QueryIidDocumentResponse</h3>
        <p>QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>iidDocument</td>
                  <td><a href="#ixo.iid.v1beta1.IidDocument">IidDocument</a></td>
                  <td></td>
                  <td><p>validators contains all the queried validators.

DidMetadata didMetadata = 2  [(gogoproto.nullable) = false]; </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.QueryIidDocumentsRequest">QueryIidDocumentsRequest</h3>
        <p>QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>status enables to query for validators matching a given status. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.QueryIidDocumentsResponse">QueryIidDocumentsResponse</h3>
        <p>QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>iidDocuments</td>
                  <td><a href="#ixo.iid.v1beta1.IidDocument">IidDocument</a></td>
                  <td>repeated</td>
                  <td><p>validators contains all the queried validators. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.QueryIidMetaDataRequest">QueryIidMetaDataRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>status enables to query for validators matching a given status. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.iid.v1beta1.QueryIidMetaDataResponse">QueryIidMetaDataResponse</h3>
        <p>this line is used by starport scaffolding # 3</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>didMetadata</td>
                  <td><a href="#ixo.iid.v1beta1.IidMetadata">IidMetadata</a></td>
                  <td></td>
                  <td><p>validators contains all the queried validators.
IidDocument iidDocument = 1  [(gogoproto.nullable) = false]; </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="ixo.iid.v1beta1.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>IidDocuments</td>
                <td><a href="#ixo.iid.v1beta1.QueryIidDocumentsRequest">QueryIidDocumentsRequest</a></td>
                <td><a href="#ixo.iid.v1beta1.QueryIidDocumentsResponse">QueryIidDocumentsResponse</a></td>
                <td><p>IidDocuments queries all iid documents that match the given status.</p></td>
              </tr>
            
              <tr>
                <td>IidDocument</td>
                <td><a href="#ixo.iid.v1beta1.QueryIidDocumentRequest">QueryIidDocumentRequest</a></td>
                <td><a href="#ixo.iid.v1beta1.QueryIidDocumentResponse">QueryIidDocumentResponse</a></td>
                <td><p>IidDocument queries a iid documents with an id.</p></td>
              </tr>
            
              <tr>
                <td>MetaData</td>
                <td><a href="#ixo.iid.v1beta1.QueryIidMetaDataRequest">QueryIidMetaDataRequest</a></td>
                <td><a href="#ixo.iid.v1beta1.QueryIidMetaDataResponse">QueryIidMetaDataResponse</a></td>
                <td><p>MetaData queries a iid documents with an id.</p></td>
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
                <td>IidDocuments</td>
                <td>GET</td>
                <td>/ixo/did/dids</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>IidDocument</td>
                <td>GET</td>
                <td>/ixo/did/dids/{id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>MetaData</td>
                <td>GET</td>
                <td>/ixo/did/dids/{id}</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="ixo/legacy/did/did.proto">ixo/legacy/did/did.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="legacydid.Claim">Claim</h3>
=======
        <h3 id="did.Claim">Claim</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="legacydid.DidCredential">DidCredential</h3>
=======
        <h3 id="did.DidCredential">DidCredential</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#legacydid.Claim">Claim</a></td>
=======
                  <td><a href="#did.Claim">Claim</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="legacydid.IxoDid">IxoDid</h3>
=======
        <h3 id="did.IxoDid">IxoDid</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#legacydid.Secret">Secret</a></td>
=======
                  <td><a href="#did.Secret">Secret</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="legacydid.Secret">Secret</h3>
=======
        <h3 id="did.Secret">Secret</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/legacy/did/diddoc.proto">ixo/legacy/did/diddoc.proto</h2><a href="#title">Top</a>
=======
        <h2 id="did/diddoc.proto">did/diddoc.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="legacydid.BaseDidDoc">BaseDidDoc</h3>
=======
        <h3 id="did.BaseDidDoc">BaseDidDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#legacydid.DidCredential">DidCredential</a></td>
=======
                  <td><a href="#did.DidCredential">DidCredential</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/payments/v1/payments.proto">ixo/payments/v1/payments.proto</h2><a href="#title">Top</a>
=======
        <h2 id="did/genesis.proto">did/genesis.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.BlockPeriod">BlockPeriod</h3>
=======
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
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.Discount">Discount</h3>
=======
        <h3 id="payments.Discount">Discount</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.DistributionShare">DistributionShare</h3>
=======
        <h3 id="payments.DistributionShare">DistributionShare</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.PaymentContract">PaymentContract</h3>
=======
        <h3 id="payments.PaymentContract">PaymentContract</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.DistributionShare">DistributionShare</a></td>
=======
                  <td><a href="#payments.DistributionShare">DistributionShare</a></td>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.PaymentTemplate">PaymentTemplate</h3>
=======
        <h3 id="payments.PaymentTemplate">PaymentTemplate</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.Discount">Discount</a></td>
=======
                  <td><a href="#payments.Discount">Discount</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.Subscription">Subscription</h3>
=======
        <h3 id="payments.Subscription">Subscription</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.TestPeriod">TestPeriod</h3>
=======
        <h3 id="payments.TestPeriod">TestPeriod</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.TimePeriod">TimePeriod</h3>
=======
        <h3 id="payments.TimePeriod">TimePeriod</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/payments/v1/genesis.proto">ixo/payments/v1/genesis.proto</h2><a href="#title">Top</a>
=======
        <h2 id="payments/genesis.proto">payments/genesis.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.GenesisState">GenesisState</h3>
=======
        <h3 id="payments.GenesisState">GenesisState</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>GenesisState defines the payments module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_templates</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentTemplate">PaymentTemplate</a></td>
=======
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payment_contracts</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentContract">PaymentContract</a></td>
=======
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>subscriptions</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.Subscription">Subscription</a></td>
=======
                  <td><a href="#payments.Subscription">Subscription</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/payments/v1/query.proto">ixo/payments/v1/query.proto</h2><a href="#title">Top</a>
=======
        <h2 id="payments/query.proto">payments/query.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentContractRequest">QueryPaymentContractRequest</h3>
=======
        <h3 id="payments.QueryPaymentContractRequest">QueryPaymentContractRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentContractResponse">QueryPaymentContractResponse</h3>
=======
        <h3 id="payments.QueryPaymentContractResponse">QueryPaymentContractResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryPaymentContractResponse is the response type for the Query/PaymentContract RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contract</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentContract">PaymentContract</a></td>
=======
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</h3>
=======
        <h3 id="payments.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</h3>
=======
        <h3 id="payments.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryPaymentContractsByIdPrefixResponse is the response type for the Query/PaymentContractsByIdPrefix RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_contracts</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentContract">PaymentContract</a></td>
=======
                  <td><a href="#payments.PaymentContract">PaymentContract</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</h3>
=======
        <h3 id="payments.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</h3>
=======
        <h3 id="payments.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryPaymentTemplateResponse is the response type for the Query/PaymentTemplate RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>payment_template</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentTemplate">PaymentTemplate</a></td>
=======
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QuerySubscriptionRequest">QuerySubscriptionRequest</h3>
=======
        <h3 id="payments.QuerySubscriptionRequest">QuerySubscriptionRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.QuerySubscriptionResponse">QuerySubscriptionResponse</h3>
=======
        <h3 id="payments.QuerySubscriptionResponse">QuerySubscriptionResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QuerySubscriptionResponse is the response type for the Query/Subscription RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>subscription</td>
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.Subscription">Subscription</a></td>
=======
                  <td><a href="#payments.Subscription">Subscription</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.Query">Query</h3>
=======
        <h3 id="payments.Query">Query</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>PaymentTemplate</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</a></td>
                <td><a href="#ixo.payments.v1.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</a></td>
=======
                <td><a href="#payments.QueryPaymentTemplateRequest">QueryPaymentTemplateRequest</a></td>
                <td><a href="#payments.QueryPaymentTemplateResponse">QueryPaymentTemplateResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>PaymentTemplate queries info of a specific payment template.</p></td>
              </tr>
            
              <tr>
                <td>PaymentContract</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.QueryPaymentContractRequest">QueryPaymentContractRequest</a></td>
                <td><a href="#ixo.payments.v1.QueryPaymentContractResponse">QueryPaymentContractResponse</a></td>
=======
                <td><a href="#payments.QueryPaymentContractRequest">QueryPaymentContractRequest</a></td>
                <td><a href="#payments.QueryPaymentContractResponse">QueryPaymentContractResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>PaymentContract queries info of a specific payment contract.</p></td>
              </tr>
            
              <tr>
                <td>PaymentContractsByIdPrefix</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</a></td>
                <td><a href="#ixo.payments.v1.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</a></td>
=======
                <td><a href="#payments.QueryPaymentContractsByIdPrefixRequest">QueryPaymentContractsByIdPrefixRequest</a></td>
                <td><a href="#payments.QueryPaymentContractsByIdPrefixResponse">QueryPaymentContractsByIdPrefixResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>PaymentContractsByIdPrefix lists all payment contracts having an id with a specific prefix.</p></td>
              </tr>
            
              <tr>
                <td>Subscription</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.QuerySubscriptionRequest">QuerySubscriptionRequest</a></td>
                <td><a href="#ixo.payments.v1.QuerySubscriptionResponse">QuerySubscriptionResponse</a></td>
=======
                <td><a href="#payments.QuerySubscriptionRequest">QuerySubscriptionRequest</a></td>
                <td><a href="#payments.QuerySubscriptionResponse">QuerySubscriptionResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/payments/v1/tx.proto">ixo/payments/v1/tx.proto</h2><a href="#title">Top</a>
=======
        <h2 id="payments/tx.proto">payments/tx.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreatePaymentContract">MsgCreatePaymentContract</h3>
=======
        <h3 id="payments.MsgCreatePaymentContract">MsgCreatePaymentContract</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.DistributionShare">DistributionShare</a></td>
=======
                  <td><a href="#payments.DistributionShare">DistributionShare</a></td>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>creator_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</h3>
=======
        <h3 id="payments.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreatePaymentContractResponse defines the Msg/CreatePaymentContract response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</h3>
=======
        <h3 id="payments.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.payments.v1.PaymentTemplate">PaymentTemplate</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>creator_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#payments.PaymentTemplate">PaymentTemplate</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</h3>
=======
        <h3 id="payments.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreatePaymentTemplateResponse defines the Msg/CreatePaymentTemplate response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreateSubscription">MsgCreateSubscription</h3>
=======
        <h3 id="payments.MsgCreateSubscription">MsgCreateSubscription</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>creator_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</h3>
=======
        <h3 id="payments.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateSubscriptionResponse defines the Msg/CreateSubscription response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgEffectPayment">MsgEffectPayment</h3>
=======
        <h3 id="payments.MsgEffectPayment">MsgEffectPayment</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>partial_payment_amount</td>
                  <td><a href="#cosmos.base.v1beta1.Coin">cosmos.base.v1beta1.Coin</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgEffectPaymentResponse">MsgEffectPaymentResponse</h3>
=======
        <h3 id="payments.MsgEffectPaymentResponse">MsgEffectPaymentResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgEffectPaymentResponse defines the Msg/EffectPayment response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgGrantDiscount">MsgGrantDiscount</h3>
=======
        <h3 id="payments.MsgGrantDiscount">MsgGrantDiscount</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>sender_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgGrantDiscountResponse">MsgGrantDiscountResponse</h3>
=======
        <h3 id="payments.MsgGrantDiscountResponse">MsgGrantDiscountResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgGrantDiscountResponse defines the Msg/GrantDiscount response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgRevokeDiscount">MsgRevokeDiscount</h3>
=======
        <h3 id="payments.MsgRevokeDiscount">MsgRevokeDiscount</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>sender_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</h3>
=======
        <h3 id="payments.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgRevokeDiscountResponse defines the Msg/RevokeDiscount response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</h3>
=======
        <h3 id="payments.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>payer_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</h3>
=======
        <h3 id="payments.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgSetPaymentContractAuthorisationResponse defines the Msg/SetPaymentContractAuthorisation response type.</p>

        

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.payments.v1.Msg">Msg</h3>
=======
        <h3 id="payments.Msg">Msg</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Msg defines the payments Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>SetPaymentContractAuthorisation</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</a></td>
                <td><a href="#ixo.payments.v1.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</a></td>
=======
                <td><a href="#payments.MsgSetPaymentContractAuthorisation">MsgSetPaymentContractAuthorisation</a></td>
                <td><a href="#payments.MsgSetPaymentContractAuthorisationResponse">MsgSetPaymentContractAuthorisationResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>SetPaymentContractAuthorisation defines a method for authorising or deauthorising a payment contract.</p></td>
              </tr>
            
              <tr>
                <td>CreatePaymentTemplate</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</a></td>
                <td><a href="#ixo.payments.v1.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</a></td>
=======
                <td><a href="#payments.MsgCreatePaymentTemplate">MsgCreatePaymentTemplate</a></td>
                <td><a href="#payments.MsgCreatePaymentTemplateResponse">MsgCreatePaymentTemplateResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreatePaymentTemplate defines a method for creating a payment template.</p></td>
              </tr>
            
              <tr>
                <td>CreatePaymentContract</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgCreatePaymentContract">MsgCreatePaymentContract</a></td>
                <td><a href="#ixo.payments.v1.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</a></td>
=======
                <td><a href="#payments.MsgCreatePaymentContract">MsgCreatePaymentContract</a></td>
                <td><a href="#payments.MsgCreatePaymentContractResponse">MsgCreatePaymentContractResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreatePaymentContract defines a method for creating a payment contract.</p></td>
              </tr>
            
              <tr>
                <td>CreateSubscription</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgCreateSubscription">MsgCreateSubscription</a></td>
                <td><a href="#ixo.payments.v1.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</a></td>
=======
                <td><a href="#payments.MsgCreateSubscription">MsgCreateSubscription</a></td>
                <td><a href="#payments.MsgCreateSubscriptionResponse">MsgCreateSubscriptionResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateSubscription defines a method for creating a subscription.</p></td>
              </tr>
            
              <tr>
                <td>GrantDiscount</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgGrantDiscount">MsgGrantDiscount</a></td>
                <td><a href="#ixo.payments.v1.MsgGrantDiscountResponse">MsgGrantDiscountResponse</a></td>
=======
                <td><a href="#payments.MsgGrantDiscount">MsgGrantDiscount</a></td>
                <td><a href="#payments.MsgGrantDiscountResponse">MsgGrantDiscountResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>GrantDiscount defines a method for granting a discount to a payer on a specific payment contract.</p></td>
              </tr>
            
              <tr>
                <td>RevokeDiscount</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgRevokeDiscount">MsgRevokeDiscount</a></td>
                <td><a href="#ixo.payments.v1.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</a></td>
=======
                <td><a href="#payments.MsgRevokeDiscount">MsgRevokeDiscount</a></td>
                <td><a href="#payments.MsgRevokeDiscountResponse">MsgRevokeDiscountResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>RevokeDiscount defines a method for revoking a discount previously granted to a payer.</p></td>
              </tr>
            
              <tr>
                <td>EffectPayment</td>
<<<<<<< HEAD
                <td><a href="#ixo.payments.v1.MsgEffectPayment">MsgEffectPayment</a></td>
                <td><a href="#ixo.payments.v1.MsgEffectPaymentResponse">MsgEffectPaymentResponse</a></td>
=======
                <td><a href="#payments.MsgEffectPayment">MsgEffectPayment</a></td>
                <td><a href="#payments.MsgEffectPaymentResponse">MsgEffectPaymentResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>EffectPayment defines a method for putting a specific payment contract into effect.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/project/v1/project.proto">ixo/project/v1/project.proto</h2><a href="#title">Top</a>
=======
        <h2 id="project/project.proto">project/project.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.AccountMap">AccountMap</h3>
=======
        <h3 id="project.AccountMap">AccountMap</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>AccountMap maps a specific project's account names to the accounts' addresses.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.AccountMap.MapEntry">AccountMap.MapEntry</a></td>
=======
                  <td><a href="#project.AccountMap.MapEntry">AccountMap.MapEntry</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.AccountMap.MapEntry">AccountMap.MapEntry</h3>
=======
        <h3 id="project.AccountMap.MapEntry">AccountMap.MapEntry</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.Claim">Claim</h3>
=======
        <h3 id="project.Claim">Claim</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.Claims">Claims</h3>
=======
        <h3 id="project.Claims">Claims</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Claims contains a list of type Claim.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>claims_list</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.Claim">Claim</a></td>
=======
                  <td><a href="#project.Claim">Claim</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.CreateAgentDoc">CreateAgentDoc</h3>
=======
        <h3 id="project.CreateAgentDoc">CreateAgentDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.CreateClaimDoc">CreateClaimDoc</h3>
=======
        <h3 id="project.CreateClaimDoc">CreateClaimDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.CreateEvaluationDoc">CreateEvaluationDoc</h3>
=======
        <h3 id="project.CreateEvaluationDoc">CreateEvaluationDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.GenesisAccountMap">GenesisAccountMap</h3>
=======
        <h3 id="project.GenesisAccountMap">GenesisAccountMap</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>GenesisAccountMap is a type used at genesis that maps a specific project's account names to the accounts' addresses.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</a></td>
=======
                  <td><a href="#project.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</h3>
=======
        <h3 id="project.GenesisAccountMap.MapEntry">GenesisAccountMap.MapEntry</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.Params">Params</h3>
=======
        <h3 id="project.Params">Params</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.ProjectDoc">ProjectDoc</h3>
=======
        <h3 id="project.ProjectDoc">ProjectDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.UpdateAgentDoc">UpdateAgentDoc</h3>
=======
        <h3 id="project.UpdateAgentDoc">UpdateAgentDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.UpdateProjectStatusDoc">UpdateProjectStatusDoc</h3>
=======
        <h3 id="project.UpdateProjectStatusDoc">UpdateProjectStatusDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.WithdrawFundsDoc">WithdrawFundsDoc</h3>
=======
        <h3 id="project.WithdrawFundsDoc">WithdrawFundsDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.WithdrawalInfoDoc">WithdrawalInfoDoc</h3>
=======
        <h3 id="project.WithdrawalInfoDoc">WithdrawalInfoDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.WithdrawalInfoDocs">WithdrawalInfoDocs</h3>
=======
        <h3 id="project.WithdrawalInfoDocs">WithdrawalInfoDocs</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>WithdrawalInfoDocs contains a list of type WithdrawalInfoDoc.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>docs_list</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.WithdrawalInfoDoc">WithdrawalInfoDoc</a></td>
=======
                  <td><a href="#project.WithdrawalInfoDoc">WithdrawalInfoDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/project/v1/genesis.proto">ixo/project/v1/genesis.proto</h2><a href="#title">Top</a>
=======
        <h2 id="project/genesis.proto">project/genesis.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.GenesisState">GenesisState</h3>
=======
        <h3 id="project.GenesisState">GenesisState</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>GenesisState defines the project module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_docs</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.ProjectDoc">ProjectDoc</a></td>
=======
                  <td><a href="#project.ProjectDoc">ProjectDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>account_maps</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.GenesisAccountMap">GenesisAccountMap</a></td>
=======
                  <td><a href="#project.GenesisAccountMap">GenesisAccountMap</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>withdrawals_infos</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
=======
                  <td><a href="#project.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>claims</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.Claims">Claims</a></td>
=======
                  <td><a href="#project.Claims">Claims</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.Params">Params</a></td>
=======
                  <td><a href="#project.Params">Params</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
<<<<<<< HEAD
        <h2 id="ixo/project/v1/query.proto">ixo/project/v1/query.proto</h2><a href="#title">Top</a>
=======
        <h2 id="project/query.proto">project/query.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryParamsRequest">QueryParamsRequest</h3>
=======
        <h3 id="project.QueryParamsRequest">QueryParamsRequest</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryParamsRequest is the request type for the Query/Params RPC method.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryParamsResponse">QueryParamsResponse</h3>
=======
        <h3 id="project.QueryParamsResponse">QueryParamsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryParamsResponse is the response type for the Query/Params RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.Params">Params</a></td>
=======
                  <td><a href="#project.Params">Params</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectAccountsRequest">QueryProjectAccountsRequest</h3>
=======
        <h3 id="project.QueryProjectAccountsRequest">QueryProjectAccountsRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectAccountsResponse">QueryProjectAccountsResponse</h3>
=======
        <h3 id="project.QueryProjectAccountsResponse">QueryProjectAccountsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryProjectAccountsResponse is the response type for the Query/ProjectAccounts RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>account_map</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.AccountMap">AccountMap</a></td>
=======
                  <td><a href="#project.AccountMap">AccountMap</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectDocRequest">QueryProjectDocRequest</h3>
=======
        <h3 id="project.QueryProjectDocRequest">QueryProjectDocRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectDocResponse">QueryProjectDocResponse</h3>
=======
        <h3 id="project.QueryProjectDocResponse">QueryProjectDocResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>project_doc</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.ProjectDoc">ProjectDoc</a></td>
=======
                  <td><a href="#project.ProjectDoc">ProjectDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectTxRequest">QueryProjectTxRequest</h3>
=======
        <h3 id="project.QueryProjectTxRequest">QueryProjectTxRequest</h3>
>>>>>>> upstream/devel/ben-alpha
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

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.QueryProjectTxResponse">QueryProjectTxResponse</h3>
=======
        <h3 id="project.QueryProjectTxResponse">QueryProjectTxResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>QueryProjectTxResponse is the response type for the Query/ProjectTx RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>txs</td>
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
=======
                  <td><a href="#project.WithdrawalInfoDocs">WithdrawalInfoDocs</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.Query">Query</h3>
=======
        <h3 id="project.Query">Query</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>ProjectDoc</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.QueryProjectDocRequest">QueryProjectDocRequest</a></td>
                <td><a href="#ixo.project.v1.QueryProjectDocResponse">QueryProjectDocResponse</a></td>
=======
                <td><a href="#project.QueryProjectDocRequest">QueryProjectDocRequest</a></td>
                <td><a href="#project.QueryProjectDocResponse">QueryProjectDocResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>ProjectDoc queries info of a specific project.</p></td>
              </tr>
            
              <tr>
                <td>ProjectAccounts</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.QueryProjectAccountsRequest">QueryProjectAccountsRequest</a></td>
                <td><a href="#ixo.project.v1.QueryProjectAccountsResponse">QueryProjectAccountsResponse</a></td>
=======
                <td><a href="#project.QueryProjectAccountsRequest">QueryProjectAccountsRequest</a></td>
                <td><a href="#project.QueryProjectAccountsResponse">QueryProjectAccountsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>ProjectAccounts lists a specific project&#39;s accounts.</p></td>
              </tr>
            
              <tr>
                <td>ProjectTx</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.QueryProjectTxRequest">QueryProjectTxRequest</a></td>
                <td><a href="#ixo.project.v1.QueryProjectTxResponse">QueryProjectTxResponse</a></td>
=======
                <td><a href="#project.QueryProjectTxRequest">QueryProjectTxRequest</a></td>
                <td><a href="#project.QueryProjectTxResponse">QueryProjectTxResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>ProjectTx lists a specific project&#39;s transactions.</p></td>
              </tr>
            
              <tr>
                <td>Params</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#ixo.project.v1.QueryParamsResponse">QueryParamsResponse</a></td>
=======
                <td><a href="#project.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#project.QueryParamsResponse">QueryParamsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
        <h2 id="ixo/project/v1/tx.proto">ixo/project/v1/tx.proto</h2><a href="#title">Top</a>
=======
        <h2 id="project/tx.proto">project/tx.proto</h2><a href="#title">Top</a>
>>>>>>> upstream/devel/ben-alpha
      </div>
      <p></p>

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateAgent">MsgCreateAgent</h3>
=======
        <h3 id="project.MsgCreateAgent">MsgCreateAgent</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.CreateAgentDoc">CreateAgentDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.CreateAgentDoc">CreateAgentDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateAgentResponse">MsgCreateAgentResponse</h3>
=======
        <h3 id="project.MsgCreateAgentResponse">MsgCreateAgentResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateAgentResponse defines the Msg/CreateAgent response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateClaim">MsgCreateClaim</h3>
=======
        <h3 id="project.MsgCreateClaim">MsgCreateClaim</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.CreateClaimDoc">CreateClaimDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.CreateClaimDoc">CreateClaimDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateClaimResponse">MsgCreateClaimResponse</h3>
=======
        <h3 id="project.MsgCreateClaimResponse">MsgCreateClaimResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateClaimResponse defines the Msg/CreateClaim response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateEvaluation">MsgCreateEvaluation</h3>
=======
        <h3 id="project.MsgCreateEvaluation">MsgCreateEvaluation</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.CreateEvaluationDoc">CreateEvaluationDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.CreateEvaluationDoc">CreateEvaluationDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</h3>
=======
        <h3 id="project.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateEvaluationResponse defines the Msg/CreateEvaluation response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateProject">MsgCreateProject</h3>
=======
        <h3 id="project.MsgCreateProject">MsgCreateProject</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgCreateProjectResponse">MsgCreateProjectResponse</h3>
=======
        <h3 id="project.MsgCreateProjectResponse">MsgCreateProjectResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgCreateProjectResponse defines the Msg/CreateProject response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateAgent">MsgUpdateAgent</h3>
=======
        <h3 id="project.MsgUpdateAgent">MsgUpdateAgent</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.UpdateAgentDoc">UpdateAgentDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.UpdateAgentDoc">UpdateAgentDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateAgentResponse">MsgUpdateAgentResponse</h3>
=======
        <h3 id="project.MsgUpdateAgentResponse">MsgUpdateAgentResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgUpdateAgentResponse defines the Msg/UpdateAgent response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateProjectDoc">MsgUpdateProjectDoc</h3>
=======
        <h3 id="project.MsgUpdateProjectDoc">MsgUpdateProjectDoc</h3>
>>>>>>> upstream/devel/ben-alpha
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
              
<<<<<<< HEAD
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
=======
>>>>>>> upstream/devel/ben-alpha
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</h3>
=======
        <h3 id="project.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgUpdateProjectDocResponse defines the Msg/UpdateProjectDoc response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateProjectStatus">MsgUpdateProjectStatus</h3>
=======
        <h3 id="project.MsgUpdateProjectStatus">MsgUpdateProjectStatus</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.UpdateProjectStatusDoc">UpdateProjectStatusDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>project_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.UpdateProjectStatusDoc">UpdateProjectStatusDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</h3>
=======
        <h3 id="project.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgUpdateProjectStatusResponse defines the Msg/UpdateProjectStatus response type.</p>

        

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgWithdrawFunds">MsgWithdrawFunds</h3>
=======
        <h3 id="project.MsgWithdrawFunds">MsgWithdrawFunds</h3>
>>>>>>> upstream/devel/ben-alpha
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
<<<<<<< HEAD
                  <td><a href="#ixo.project.v1.WithdrawFundsDoc">WithdrawFundsDoc</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>sender_address</td>
                  <td><a href="#string">string</a></td>
=======
                  <td><a href="#project.WithdrawFundsDoc">WithdrawFundsDoc</a></td>
>>>>>>> upstream/devel/ben-alpha
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</h3>
=======
        <h3 id="project.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>MsgWithdrawFundsResponse defines the Msg/WithdrawFunds response type.</p>

        

        
      

      

      

      
<<<<<<< HEAD
        <h3 id="ixo.project.v1.Msg">Msg</h3>
=======
        <h3 id="project.Msg">Msg</h3>
>>>>>>> upstream/devel/ben-alpha
        <p>Msg defines the project Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateProject</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgCreateProject">MsgCreateProject</a></td>
                <td><a href="#ixo.project.v1.MsgCreateProjectResponse">MsgCreateProjectResponse</a></td>
=======
                <td><a href="#project.MsgCreateProject">MsgCreateProject</a></td>
                <td><a href="#project.MsgCreateProjectResponse">MsgCreateProjectResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateProject defines a method for creating a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateProjectStatus</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgUpdateProjectStatus">MsgUpdateProjectStatus</a></td>
                <td><a href="#ixo.project.v1.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</a></td>
=======
                <td><a href="#project.MsgUpdateProjectStatus">MsgUpdateProjectStatus</a></td>
                <td><a href="#project.MsgUpdateProjectStatusResponse">MsgUpdateProjectStatusResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>UpdateProjectStatus defines a method for updating a project&#39;s current status.</p></td>
              </tr>
            
              <tr>
                <td>CreateAgent</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgCreateAgent">MsgCreateAgent</a></td>
                <td><a href="#ixo.project.v1.MsgCreateAgentResponse">MsgCreateAgentResponse</a></td>
=======
                <td><a href="#project.MsgCreateAgent">MsgCreateAgent</a></td>
                <td><a href="#project.MsgCreateAgentResponse">MsgCreateAgentResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateAgent defines a method for creating an agent on a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateAgent</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgUpdateAgent">MsgUpdateAgent</a></td>
                <td><a href="#ixo.project.v1.MsgUpdateAgentResponse">MsgUpdateAgentResponse</a></td>
=======
                <td><a href="#project.MsgUpdateAgent">MsgUpdateAgent</a></td>
                <td><a href="#project.MsgUpdateAgentResponse">MsgUpdateAgentResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>UpdateAgent defines a method for updating an agent on a project.</p></td>
              </tr>
            
              <tr>
                <td>CreateClaim</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgCreateClaim">MsgCreateClaim</a></td>
                <td><a href="#ixo.project.v1.MsgCreateClaimResponse">MsgCreateClaimResponse</a></td>
=======
                <td><a href="#project.MsgCreateClaim">MsgCreateClaim</a></td>
                <td><a href="#project.MsgCreateClaimResponse">MsgCreateClaimResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateClaim defines a method for creating a claim on a project.</p></td>
              </tr>
            
              <tr>
                <td>CreateEvaluation</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgCreateEvaluation">MsgCreateEvaluation</a></td>
                <td><a href="#ixo.project.v1.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</a></td>
=======
                <td><a href="#project.MsgCreateEvaluation">MsgCreateEvaluation</a></td>
                <td><a href="#project.MsgCreateEvaluationResponse">MsgCreateEvaluationResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>CreateEvaluation defines a method for creating an evaluation for a specific claim on a project.</p></td>
              </tr>
            
              <tr>
                <td>WithdrawFunds</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgWithdrawFunds">MsgWithdrawFunds</a></td>
                <td><a href="#ixo.project.v1.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</a></td>
=======
                <td><a href="#project.MsgWithdrawFunds">MsgWithdrawFunds</a></td>
                <td><a href="#project.MsgWithdrawFundsResponse">MsgWithdrawFundsResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>WithdrawFunds defines a method for project agents to withdraw their funds from a project.</p></td>
              </tr>
            
              <tr>
                <td>UpdateProjectDoc</td>
<<<<<<< HEAD
                <td><a href="#ixo.project.v1.MsgUpdateProjectDoc">MsgUpdateProjectDoc</a></td>
                <td><a href="#ixo.project.v1.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</a></td>
=======
                <td><a href="#project.MsgUpdateProjectDoc">MsgUpdateProjectDoc</a></td>
                <td><a href="#project.MsgUpdateProjectDocResponse">MsgUpdateProjectDocResponse</a></td>
>>>>>>> upstream/devel/ben-alpha
                <td><p>UpdateProjectDoc defines a method for updating a project&#39;s data.</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
<<<<<<< HEAD
      
      <div class="file-heading">
        <h2 id="ixo/token/v1beta1/token.proto">ixo/token/v1beta1/token.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.token.v1beta1.Params">Params</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>NftContractAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>NftContractMinter</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.TokenDoc">TokenDoc</h3>
        <p>ProjectDoc defines a project (or token) type with all of its parameters.</p>

        

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/token/v1beta1/genesis.proto">ixo/token/v1beta1/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.token.v1beta1.GenesisState">GenesisState</h3>
        <p>GenesisState defines the project module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>token_docs</td>
                  <td><a href="#ixo.token.v1beta1.TokenDoc">TokenDoc</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>params</td>
                  <td><a href="#ixo.token.v1beta1.Params">Params</a></td>
                  <td></td>
                  <td><p>repeated GenesisAccountMap account_maps       = 2 [(gogoproto.nullable) = false, (gogoproto.moretags) = &#34;yaml:\&#34;account_maps\&#34;&#34;]; </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/token/v1beta1/proposal.proto">ixo/token/v1beta1/proposal.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.token.v1beta1.InitializeTokenContract">InitializeTokenContract</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>NftContractCodeId</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>NftMinterAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="ixo/token/v1beta1/query.proto">ixo/token/v1beta1/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.token.v1beta1.QueryTokenConfigRequest">QueryTokenConfigRequest</h3>
        <p></p>

        

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenConfigResponse">QueryTokenConfigResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>map</td>
                  <td><a href="#ixo.token.v1beta1.QueryTokenConfigResponse.MapEntry">QueryTokenConfigResponse.MapEntry</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenConfigResponse.MapEntry">QueryTokenConfigResponse.MapEntry</h3>
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

          

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenDocRequest">QueryTokenDocRequest</h3>
        <p>QueryProjectDocRequest is the request type for the Query/ProjectDoc RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>token_did</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenDocResponse">QueryTokenDocResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenListRequest">QueryTokenListRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>token_type</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>token_status</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.QueryTokenListResponse">QueryTokenListResponse</h3>
        <p>QueryProjectDocResponse is the response type for the Query/ProjectDoc RPC method.</p>

        

        
      

      

      

      
        <h3 id="ixo.token.v1beta1.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>TokenList</td>
                <td><a href="#ixo.token.v1beta1.QueryTokenListRequest">QueryTokenListRequest</a></td>
                <td><a href="#ixo.token.v1beta1.QueryTokenListResponse">QueryTokenListResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>TokenDoc</td>
                <td><a href="#ixo.token.v1beta1.QueryTokenDocRequest">QueryTokenDocRequest</a></td>
                <td><a href="#ixo.token.v1beta1.QueryTokenDocResponse">QueryTokenDocResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>TokenConfig</td>
                <td><a href="#ixo.token.v1beta1.QueryTokenConfigRequest">QueryTokenConfigRequest</a></td>
                <td><a href="#ixo.token.v1beta1.QueryTokenConfigResponse">QueryTokenConfigResponse</a></td>
                <td><p></p></td>
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
                <td>TokenList</td>
                <td>GET</td>
                <td>/ixo/token</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>TokenDoc</td>
                <td>GET</td>
                <td>/ixo/token/{token_did}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>TokenConfig</td>
                <td>GET</td>
                <td>/ixo/token/config</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="ixo/token/v1beta1/tx.proto">ixo/token/v1beta1/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="ixo.token.v1beta1.MsgCreateToken">MsgCreateToken</h3>
        <p>MsgCreateToken defines a message for creating a project.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>ownerDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>An Token Type as defined by the implementer

Owner of the Token NFT | The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>ownerAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid address used to sign this transaction. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.MsgCreateTokenResponse">MsgCreateTokenResponse</h3>
        <p>MsgCreateProjectResponse defines the Msg/CreateProject response type.</p>

        

        
      
        <h3 id="ixo.token.v1beta1.MsgTransferToken">MsgTransferToken</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>tokenDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>ownerDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>ownerAddress</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>The ownersdid used to sign this transaction. </p></td>
                </tr>
              
                <tr>
                  <td>recipientDid</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="ixo.token.v1beta1.MsgTransferTokenResponse">MsgTransferTokenResponse</h3>
        <p>MsgUpdateProjectStatusResponse defines the Msg/UpdateTokenStatus response type.</p>

        

        
      

      

      

      
        <h3 id="ixo.token.v1beta1.Msg">Msg</h3>
        <p>Msg defines the project Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateToken</td>
                <td><a href="#ixo.token.v1beta1.MsgCreateToken">MsgCreateToken</a></td>
                <td><a href="#ixo.token.v1beta1.MsgCreateTokenResponse">MsgCreateTokenResponse</a></td>
                <td><p>CreateProject defines a method for creating a project.</p></td>
              </tr>
            
              <tr>
                <td>TransferToken</td>
                <td><a href="#ixo.token.v1beta1.MsgTransferToken">MsgTransferToken</a></td>
                <td><a href="#ixo.token.v1beta1.MsgTransferTokenResponse">MsgTransferTokenResponse</a></td>
                <td><p>Transfers an token and its nft to the recipient</p></td>
              </tr>
            
          </tbody>
        </table>

        
    
=======
>>>>>>> upstream/devel/ben-alpha

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

