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

GAS_PRICES="0.025ixo"
BONDDOC_FEE_FUND="15000ixo"  # 0.025 * 200000 * 3txs

BONDDOC_DID="did:ixo:48PVm1uyF6QVDSPdGRWw4T"
BONDDOC_DID_FULL="{\"did\":\"did:ixo:48PVm1uyF6QVDSPdGRWw4T\",\"verifyKey\":\"2hs2cb232Ev97aSQLvrfK4q8ZceBR8cf33UTstWpKU9M\",\"encryptionPublicKey\":\"9k2THnNbTziXGRjn77tvWujffgigRPqPyKZUSdwjmfh2\",\"secret\":{\"seed\":\"82949a422215a5999846beaadf398659157c345564787993f92e91d192f2a9c5\",\"signKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\",\"encryptionPrivateKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\"}}"
BONDDOC_INFO="{\"created_on\":\"created_on\",\"created_by\":\"created_by\"}"

SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
SHAUN_DID_FULL="{\"did\":\"did:ixo:U4tSpzzv91HHqWW1YmFkHJ\",\"verifyKey\":\"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG\",\"encryptionPublicKey\":\"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i\",\"secret\":{\"seed\":\"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e\",\"signKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\",\"encryptionPrivateKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\"}}"

# Ledger DIDs
echo "Ledgering Shaun DID..."
ixocli tx did add-did-doc "$SHAUN_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Fund bonddoc account
echo "Funding bonddoc account so it can pay fees..."
ixocli tx treasury send $BONDDOC_DID $BONDDOC_FEE_FUND "$SHAUN_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create bonddoc and progress status to OPEN
SENDER_DID="$SHAUN_DID"
echo "Creating bonddoc..."
ixocli tx bonddoc create-bond "$SENDER_DID" "$BONDDOC_INFO" "$BONDDOC_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Updating bonddoc to PREISSUANCE..."
ixocli tx bonddoc update-bond-status "$SENDER_DID" PREISSUANCE "$BONDDOC_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
echo "Updating bonddoc to OPEN..."
ixocli tx bonddoc update-bond-status "$SENDER_DID" OPEN "$BONDDOC_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y
