#!/usr/bin/env bash

# Must be run from root path inside ixo-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

# Create multi-sig key using Miguel, Francesco, and Shaun, but with a 2 key
# threshold, meaning that only two people have to sign any transactions.
ixod keys delete mfs-multisig -y  # remove just in case it already exists
ixod keys add mfs-multisig \
  --multisig miguel,francesco,shaun \
  --multisig-threshold 2

MULTISIG_ADDR=$(yes $PASSWORD | ixod keys show mfs-multisig -a)
MIGUEL_ADDR=$(yes $PASSWORD | ixod keys show miguel -a)
SHAUN_ADDR=$(yes $PASSWORD | ixod keys show shaun -a)

# Send tokens to the multi-sig address
ixod_tx bank send miguel "$MULTISIG_ADDR" 1000000uixo
# Check balance of the multi-sig address
ixod_q bank balances "$MULTISIG_ADDR"

# Send some tokens back to Miguel
#
# ...part 1: generate transaction
ixod tx bank send "$MULTISIG_ADDR" "$MIGUEL_ADDR" 123uixo \
  --generate-only \
  --gas-prices="$GAS_PRICES" \
  --chain-id="$CHAIN_ID" \
  > tx.json
#
# ...part 2: miguel signs tx.json and produces tx-signed-miguel.json
ixod tx sign tx.json \
  --from "$MIGUEL_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-miguel.json
#
# ...part 2: shaun signs tx.json and produces tx-signed-shaun.json
ixod tx sign tx.json \
  --from "$SHAUN_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-shaun.json
#
# ...part 3: join signatures
ixod tx multisign tx.json mfs-multisig tx-signed-miguel.json tx-signed-shaun.json \
  --from mfs-multisig \
  --chain-id="$CHAIN_ID" \
  > tx_ms.json
#
# ...part 4: broadcast
ixod_tx broadcast ./tx_ms.json

# Check Miguel's balance
ixod_q bank balances "$MIGUEL_ADDR"
# Check balance of the multi-sig address
ixod_q bank balances "$MULTISIG_ADDR"

# Cleanup
ixod keys delete mfs-multisig -y
rm tx.json tx-signed-miguel.json tx-signed-shaun.json tx_msg.json