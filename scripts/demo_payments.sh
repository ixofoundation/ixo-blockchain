#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(ixod status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(ixod status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

GAS_PRICES="0.025uixo"
PASSWORD="12345678"

MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"

MIGUEL_DID_FULL='{
  "did":"did:ixo:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'
FRANCESCO_DID_FULL='{
  "did":"did:ixo:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'
SHAUN_DID_FULL='{
  "did":"did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
  "verifyKey":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
  "encryptionPublicKey":"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i",
  "secret":{
    "seed":"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e",
    "signKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR",
    "encryptionPrivateKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR"
  }
}'
PAYMENT_RECIPIENTS='[
  {
    "address": "ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx",
    "percentage": "100"
  }
]'

# Ledger DIDs
echo "Ledgering Miguel DID..."
ixod tx did add-did-doc "$MIGUEL_DID_FULL" --gas-prices="$GAS_PRICES" -y
echo "Ledgering Francesco DID..."
ixod tx did add-did-doc "$FRANCESCO_DID_FULL" --gas-prices="$GAS_PRICES" -y
echo "Ledgering Shaun DID..."
ixod tx did add-did-doc "$SHAUN_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create payment template
echo "Creating payment template..."
PAYMENT_TEMPLATE='{
  "id": "payment:template:template1",
  "payment_amount": [
    {
      "denom": "uixo",
      "amount": "10"
    }
  ],
  "payment_minimum": [],
  "payment_maximum": [],
  "discounts": []
}'
CREATOR="$MIGUEL_DID_FULL"
ixod tx payments create-payment-template "$PAYMENT_TEMPLATE" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create payment contract
echo "Creating payment contract..."
PAYMENT_TEMPLATE_ID="payment:template:template1" # from PAYMENT_TEMPLATE
PAYMENT_CONTRACT_ID="payment:contract:contract1"
DISCOUNT_ID=0
CREATOR="$SHAUN_DID_FULL"
PAYER_ADDR="$(ixod q did get-address-from-did $FRANCESCO_DID)"
ixod tx payments create-payment-contract "$PAYMENT_CONTRACT_ID" "$PAYMENT_TEMPLATE_ID" "$PAYER_ADDR" "$PAYMENT_RECIPIENTS" True "$DISCOUNT_ID" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Authorise payment contract
echo "Authorising payment contract..."
PAYER="$FRANCESCO_DID_FULL"
ixod tx payments set-payment-contract-authorisation "$PAYMENT_CONTRACT_ID" True "$PAYER" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create subscription (with block period)
echo "Creating subscription 1/2 (with block period)..."
SUBSCRIPTION_ID="payment:subscription:subscription1"
PERIOD='{
  "type": "payments/BlockPeriod",
  "value": {
    "period_length": "3",
    "period_start_block": "5"
  }
}'
MAX_PERIODS=3
CREATOR="$SHAUN_DID_FULL"
ixod tx payments create-subscription "$SUBSCRIPTION_ID" "$PAYMENT_CONTRACT_ID" "$MAX_PERIODS" "$PERIOD" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

echo "Wait a few seconds for the subscription to get effected..."
sleep 6

# Deauthorise payment contract
echo "Deauthorising payment contract..."
PAYER="$FRANCESCO_DID_FULL"
ixod tx payments set-payment-contract-authorisation "$PAYMENT_CONTRACT_ID" False "$PAYER" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

echo "Now the subscription (block-period) will just accumulate periods and not charge anything."
echo ""

# Create subscription (with time period)
echo "Creating subscription 2/2 (with time-period)..."
SUBSCRIPTION_ID="payment:subscription:subscription2"
PERIOD='{
  "type": "payments/TimePeriod",
  "value": {
    "period_duration_ns": "6000000000",
    "period_start_time": "2020-06-03T13:00:00.00Z"
  }
}'
MAX_PERIODS=3
CREATOR="$SHAUN_DID_FULL"
ixod tx payments create-subscription "$SUBSCRIPTION_ID" "$PAYMENT_CONTRACT_ID" "$MAX_PERIODS" "$PERIOD" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

echo "The subscription (time-period) will just accumulate periods and not charge anything."
