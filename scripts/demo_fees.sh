#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(ixocli status 2>&1)
    if [[ ($RET == ERROR*) || ($RET == *'"latest_block_height": "0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(ixocli status 2>&1)
if [[ ($RET == ERROR*) || ($RET == *'"latest_block_height": "0"'*) ]]; then
  wait
fi

FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
MIGUEL_DID_FULL="{\"did\":\"did:ixo:4XJLBfGtWSGKSz4BeRxdun\",\"verifyKey\":\"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt\",\"encryptionPublicKey\":\"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS\",\"secret\":{\"seed\":\"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180\",\"signKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\",\"encryptionPrivateKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\"}}"
FRANCESCO_DID_FULL="{\"did\":\"did:ixo:UKzkhVSHc3qEFva5EY2XHt\",\"verifyKey\":\"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej\",\"encryptionPublicKey\":\"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si\",\"secret\":{\"seed\":\"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de\",\"signKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\",\"encryptionPrivateKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\"}}"
SHAUN_DID_FULL="{\"did\":\"did:ixo:U4tSpzzv91HHqWW1YmFkHJ\",\"verifyKey\":\"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG\",\"encryptionPublicKey\":\"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i\",\"secret\":{\"seed\":\"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e\",\"signKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\",\"encryptionPrivateKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\"}}"

# Ledger DIDs
echo "Ledgering Miguel DID..."
ixocli tx did add-did-doc "$MIGUEL_DID_FULL" --broadcast-mode block -y
echo "Ledgering Francesco DID..."
ixocli tx did add-did-doc "$FRANCESCO_DID_FULL" --broadcast-mode block -y
echo "Ledgering Shaun DID..."
ixocli tx did add-did-doc "$SHAUN_DID_FULL" --broadcast-mode block -y

# Create fee
echo "Creating fee..."
FEE="$(sed 's/"/\"/g' samples/fee.json | tr -d '\n' | tr -d '[:blank:]')"
CREATOR="$MIGUEL_DID_FULL"
ixocli tx fees create-fee "$FEE" "$CREATOR" --broadcast-mode block -y

# Create fee contract
echo "Creating fee contract..."
FEE_ID="fee:fee1" # from FEE
FEE_CONTRACT_ID="fee:contract:fee1"
DISCOUNT_ID=0
CREATOR="$SHAUN_DID_FULL"
PAYER_ADDR="$(ixocli q did get-address-from-did $FRANCESCO_DID)"
ixocli tx fees create-fee-contract "$FEE_CONTRACT_ID" "$FEE_ID" "$PAYER_ADDR" True "$DISCOUNT_ID" "$CREATOR" --broadcast-mode block -y

# Authorise fee contract
echo "Authorising fee contract..."
PAYER="$FRANCESCO_DID_FULL"
ixocli tx fees set-fee-contract-authorisation "$FEE_CONTRACT_ID" True "$PAYER" --broadcast-mode block -y

# Create subscription (with block period)
echo "Creating subscription 1/2 (with block period)..."
SUBSCRIPTION_ID="fee:subscription:fee1"
PERIOD="$(sed 's/"/\"/g' samples/period_block.json | tr -d '\n' | tr -d '[:blank:]')"
MAX_PERIODS=3
CREATOR="$SHAUN_DID_FULL"
ixocli tx fees create-subscription "$SUBSCRIPTION_ID" "$FEE_CONTRACT_ID" "$MAX_PERIODS" "$PERIOD" "$CREATOR" --broadcast-mode block -y

echo "Wait a few seconds for the subscription to get charged..."
sleep 6

# Deauthorise fee contract
echo "Deauthorising fee contract..."
PAYER="$FRANCESCO_DID_FULL"
ixocli tx fees set-fee-contract-authorisation "$FEE_CONTRACT_ID" False "$PAYER" --broadcast-mode block -y

echo "Now the subscription (block-period) will just accumulate periods and not charge."
echo ""

# Create subscription (with time period)
echo "Creating subscription 2/2 (with time-period)..."
SUBSCRIPTION_ID="fee:subscription:fee2"
PERIOD="$(sed 's/"/\"/g' samples/period_time.json | tr -d '\n' | tr -d '[:blank:]')"
MAX_PERIODS=3
CREATOR="$SHAUN_DID_FULL"
ixocli tx fees create-subscription "$SUBSCRIPTION_ID" "$FEE_CONTRACT_ID" "$MAX_PERIODS" "$PERIOD" "$CREATOR" --broadcast-mode block -y

echo "The subscription (time-period) will just accumulate periods and not charge."
