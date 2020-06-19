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

PASSWORD="12345678"
GAS_PRICES="0.025uixo"
FEE=$(yes $PASSWORD | ixocli keys show fee -a)

BOND_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
#BOND_DID_FULL="{\"did\":\"did:ixo:U7GK8p8rVhJMKhBVRCJJ8c\",\"verifyKey\":\"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW\",\"encryptionPublicKey\":\"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m\",\"secret\":{\"seed\":\"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053\",\"signKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\",\"encryptionPrivateKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\"}}"

MIGUEL_ADDR="ixo1x2x0thq6x2rx7txl0ujpyg9rr0c8mc8ad904xw"
FRANCESCO_ADDR="ixo1nnxvyr6hs0sglppqzw4v5s9r5dwh89423xu5zp"
MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_DID_FULL="{\"did\":\"did:ixo:4XJLBfGtWSGKSz4BeRxdun\",\"verifyKey\":\"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt\",\"encryptionPublicKey\":\"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS\",\"secret\":{\"seed\":\"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180\",\"signKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\",\"encryptionPrivateKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\"}}"
FRANCESCO_DID_FULL="{\"did\":\"did:ixo:UKzkhVSHc3qEFva5EY2XHt\",\"verifyKey\":\"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej\",\"encryptionPublicKey\":\"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si\",\"secret\":{\"seed\":\"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de\",\"signKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\",\"encryptionPrivateKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\"}}"

# Ledger DIDs
echo "Ledgering DID 1/2..."
ixocli tx did add-did-doc "$MIGUEL_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Ledgering DID 2/2..."
ixocli tx did add-did-doc "$FRANCESCO_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

echo "Creating bond..."
ixocli tx bonds create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=power_function \
  --function-parameters="m:12,n:2,c:100" \
  --reserve-tokens=res \
  --tx-fee-percentage=0.5 \
  --exit-fee-percentage=0.1 \
  --fee-address="$FEE" \
  --max-supply=1000000abc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells=true \
  --batch-blocks=1 \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Created bond..."
ixocli q bonds bond "$BOND_DID"

echo "Editing bond..."
ixocli tx bonds edit-bond \
  --token=abc \
  --name="New A B C" \
  --bond-did="$BOND_DID" \
  --editor-did="$MIGUEL_DID_FULL" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Edited bond..."
ixocli q bonds bond "$BOND_DID"

echo "Miguel buys 10abc..."
ixocli tx bonds buy 10abc 1000000res "$BOND_DID" "$MIGUEL_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Miguel's account..."
ixocli q auth account "$MIGUEL_ADDR"

echo "Francesco buys 10abc..."
ixocli tx bonds buy 10abc 1000000res "$BOND_DID" "$FRANCESCO_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Francesco's account..."
ixocli q auth account "$FRANCESCO_ADDR"

echo "Miguel sells 10abc..."
ixocli tx bonds sell 10abc "$BOND_DID" "$MIGUEL_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Miguel's account..."
ixocli q auth account "$MIGUEL_ADDR"

echo "Francesco sells 10abc..."
ixocli tx bonds sell 10abc "$BOND_DID" "$FRANCESCO_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Francesco's account..."
ixocli q auth account "$FRANCESCO_ADDR"
