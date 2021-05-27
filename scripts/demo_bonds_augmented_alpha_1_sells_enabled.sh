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

PASSWORD="12345678"
GAS_PRICES="0.025uixo"
CHAIN_ID="pandora-2"
FEE=$(yes $PASSWORD | ixod keys show fee -a)

ixod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  ixod tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

ixod_q() {
  ixod q "$@" --output=json | jq .
}

BOND_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
#BOND_DID_FULL='{
#  "did":"did:ixo:U7GK8p8rVhJMKhBVRCJJ8c",
#  "verifyKey":"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW",
#  "encryptionPublicKey":"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m",
#  "secret":{
#    "seed":"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053",
#    "signKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC",
#    "encryptionPrivateKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC"
#  }
#}'

MIGUEL_ADDR="ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"
FRANCESCO_ADDR="ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y"
SHAUN_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"
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
FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
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

# Ledger DIDs
echo "Ledgering DID 1/3..."
ixod_tx did add-did-doc "$MIGUEL_DID_FULL"
echo "Ledgering DID 2/3..."
ixod_tx did add-did-doc "$FRANCESCO_DID_FULL"
echo "Ledgering DID 3/3..."
ixod_tx did add-did-doc "$SHAUN_DID_FULL"

# d0 := 1000000 // initial raise (reserve)
# p0 := 1       // initial price (reserve per token)
# theta := 0    // initial allocation (percentage)
# kappa := 3    // degrees of polynomial (i.e. x^2, x^4, x^6)

# R0 = 1000000        // initial reserve (1-theta)*d0
# S0 = 1000000        // initial supply
# V0 = 1000000000000  // invariant

echo "Creating bond..."
ixod_tx bonds create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=augmented_function \
  --function-parameters="d0:1000000,p0:1,theta:0,kappa:3.0" \
  --reserve-tokens=res \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --max-supply=20000000abc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --alpha-bond \
  --batch-blocks=1 \
  --outcome-payment="300000000" \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --controller-did="$FRANCESCO_DID"
echo "Created bond..."
ixod_q bonds bond "$BOND_DID"

echo "Miguel buys 400000abc..."
ixod_tx bonds buy 400000abc 500000res "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Francesco buys 400000abc..."
ixod_tx bonds buy 400000abc 500000res "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixod_q bank balances "$FRANCESCO_ADDR"

echo "Shaun cannot buy 200001abc..."
ixod_tx bonds buy 200001abc 500000res "$BOND_DID" "$SHAUN_DID_FULL"
echo "Shaun cannot sell anything..."
ixod_tx bonds sell 20000abc "$BOND_DID" "$SHAUN_DID_FULL"
echo "Shaun can buy 200000abc..."
ixod_tx bonds buy 200000abc 500000res "$BOND_DID" "$SHAUN_DID_FULL"
echo "Shaun's account..."
ixod_q bank balances "$SHAUN_ADDR"

echo "Bond state is now open..."  # since 1000000 (S0) reached
ixod_q bonds bond "$BOND_DID"

echo "Current price is 3..."
ixod_q bonds current-price "$BOND_DID"

echo "Changing public alpha 0.5->0.51..."
NEW_ALPHA="0.51"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Current price is now approx 1.85..."
ixod_q bonds current-price "$BOND_DID"

echo "Changing public alpha 0.51->0.4..."
NEW_ALPHA="0.4"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Current price is now approx 1.86..."
ixod_q bonds current-price "$BOND_DID"

echo "Cannot change public alpha 0.4->0.6..."
NEW_ALPHA="0.6"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel sells 400000abc..."
ixod_tx bonds sell 400000abc "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Francesco makes outcome payment of 50000000 [1]..."
ixod_tx bonds make-outcome-payment "$BOND_DID" "50000000" "$FRANCESCO_DID_FULL"
echo "Francesco makes outcome payment of 100000000 [2]..."
ixod_tx bonds make-outcome-payment "$BOND_DID" "100000000" "$FRANCESCO_DID_FULL"
echo "Francesco makes outcome payment of 150000000 [3]..."
ixod_tx bonds make-outcome-payment "$BOND_DID" "150000000" "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixod_q bank balances "$FRANCESCO_ADDR"
echo "Bond outcome payment reserve is now 300000000..."
ixod_q bonds bond "$BOND_DID"

echo "Francesco updates the bond state to SETTLE"
ixod_tx bonds update-bond-state "SETTLE" "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Bond outcome payment reserve is now empty (moved to main reserve)..."
ixod_q bonds bond "$BOND_DID"

echo "Francesco withdraws share..."
ixod_tx bonds withdraw-share "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixod_q bank balances "$FRANCESCO_ADDR"

echo "Shaun withdraws share..."
ixod_tx bonds withdraw-share "$BOND_DID" "$SHAUN_DID_FULL"
echo "Shaun's account..."
ixod_q bank balances "$SHAUN_ADDR"

echo "Bond reserve is now empty and supply is 0..."
ixod_q bonds bond "$BOND_DID"
