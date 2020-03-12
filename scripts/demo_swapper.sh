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

tx() {
  cmd=$1
  shift
  ixocli tx bonds "$cmd" --broadcast-mode block "$@"
}

RET=$(ixocli status 2>&1)
if [[ ($RET == ERROR*) || ($RET == *'"latest_block_height": "0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
FEE=$(yes $PASSWORD | ixocli keys show fee -a)

BOND_DID="{\"did\":\"U7GK8p8rVhJMKhBVRCJJ8c\",\"verifyKey\":\"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW\",\"encryptionPublicKey\":\"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m\",\"secret\":{\"seed\":\"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053\",\"signKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\",\"encryptionPrivateKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\"}}"

MIGUEL_ADDR="cosmos1yhfwr0u62dpy35hl4e6nxarysuuwrwhxmle95m"
FRANCESCO_ADDR="cosmos16ne06jpdr3eu2a3uv3etwxavw0q03sqscjdcj5"
MIGUEL_DID="4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_DID_FULL="{\"did\":\"4XJLBfGtWSGKSz4BeRxdun\",\"verifyKey\":\"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt\",\"encryptionPublicKey\":\"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS\",\"secret\":{\"seed\":\"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180\",\"signKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\",\"encryptionPrivateKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\"}}"
FRANCESCO_DID_FULL="{\"did\":\"UKzkhVSHc3qEFva5EY2XHt\",\"verifyKey\":\"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej\",\"encryptionPublicKey\":\"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si\",\"secret\":{\"seed\":\"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de\",\"signKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\",\"encryptionPrivateKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\"}}"

# Ledger DIDs
echo "Ledgering DID 1/2..."
ixocli tx did addDidDoc "$MIGUEL_DID_FULL"
echo "Ledgering DID 2/2..."
ixocli tx did addDidDoc "$FRANCESCO_DID_FULL"

echo "Creating bond..."
ixocli tx bonds create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=swapper_function \
  --function-parameters="" \
  --reserve-tokens=res,rez \
  --tx-fee-percentage=0.5 \
  --exit-fee-percentage=0.1 \
  --fee-address="$FEE" \
  --max-supply=1000000abc \
  --order-quantity-limits="10abc,5000res,5000rez" \
  --sanity-rate="0.5" \
  --sanity-margin-percentage="20" \
  --allow-sells=true \
  --batch-blocks=1 \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID" \
  --broadcast-mode block
echo "Created bond..."
ixocli query bonds bond U7GK8p8rVhJMKhBVRCJJ8c

echo "Miguel buys 1abc..."
tx buy 1abc 500res,1000rez U7GK8p8rVhJMKhBVRCJJ8c "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixocli query auth account "$MIGUEL_ADDR"

echo "Francesco buys 10abc..."
tx buy 10abc 10100res,10100rez U7GK8p8rVhJMKhBVRCJJ8c "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixocli query auth account "$FRANCESCO_ADDR"

echo "Miguel swap 500 res to rez..."
tx swap 500 res rez U7GK8p8rVhJMKhBVRCJJ8c "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixocli query auth account "$MIGUEL_ADDR"

echo "Francesco swap 500 rez to res..."
tx swap 500 rez res U7GK8p8rVhJMKhBVRCJJ8c "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixocli query auth account "$FRANCESCO_ADDR"

echo "Miguel swaps above order limit..."
tx swap 5001 res rez U7GK8p8rVhJMKhBVRCJJ8c "$MIGUEL_DID_FULL"
echo "Miguel's account (no  changes)..."
ixocli query auth account "$MIGUEL_ADDR"

echo "Francesco swaps to violate sanity..."
tx swap 5000 rez res U7GK8p8rVhJMKhBVRCJJ8c "$FRANCESCO_DID_FULL"
echo "Francesco's account (no changes)..."
ixocli query auth account "$FRANCESCO_ADDR"

echo "Miguel sells 1abc..."
tx sell 1abc U7GK8p8rVhJMKhBVRCJJ8c "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixocli query auth account "$MIGUEL_ADDR"

echo "Francesco sells 10abc..."
tx sell 10abc U7GK8p8rVhJMKhBVRCJJ8c "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixocli query auth account "$FRANCESCO_ADDR"
