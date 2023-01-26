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
CHAIN_ID="pandora-4"

ixod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args
  # Get module ($1) and specific tx ($2), which forms the tx command
  cmd="$1 $2"
  shift 2

  ixod tx $cmd --gas-prices $GAS_PRICES --chain-id $CHAIN_ID --broadcast-mode block -y "$@" | jq .
    # The $@ adds any extra arguments to the end
}

ixod_q() {
  ixod q "$@" --output=json | jq .
}


MIGUEL_ADDR="ixo14q85xdkmg6j8nypzm0qclu0f2x8ya78k8f6kre"
MIGUEL_DID="did:x:zQ3shmDLZc6Xu2PLdNUEjRABGM8HeKYjzMtx5E5dFNCUSAeKz"
MIGUEL_DID_FULL='{
  "id": "did:x:zQ3shmDLZc6Xu2PLdNUEjRABGM8HeKYjzMtx5E5dFNCUSAeKz",
  "signer": "ixo14q85xdkmg6j8nypzm0qclu0f2x8ya78k8f6kre",
  "controllers": ["did:x:zQ3shmDLZc6Xu2PLdNUEjRABGM8HeKYjzMtx5E5dFNCUSAeKz"],
  "verifications": [
    {
      "method": {
        "id": "did:x:zQ3shmDLZc6Xu2PLdNUEjRABGM8HeKYjzMtx5E5dFNCUSAeKz",
        "type": "EcdsaSecp256k1VerificationKey2019",
        "controller": "did:x:zQ3shmDLZc6Xu2PLdNUEjRABGM8HeKYjzMtx5E5dFNCUSAeKz",
        "publicKeyBase58": "21GBYkx4Rhk7k8NZK35JDXXvCfnZ25LuYJ9sT4roJxAUG"
      },
      "relationships": ["authentication"],
      "contexts": []
    }
  ],
  "context": [],
  "services": [],
  "accorded_right": [],
  "linked_resources": [],
  "linked_entity": []
}'

# Ledger DIDs
echo "Ledgering DID 1/2..."
echo $MIGUEL_DID_FULL | jq -rc .
# ixod_tx iid create-iid-from-legacy-did "$MIGUEL_DID_FULL"
yes $PASSWORD | ixod_tx iid create-iid "$(echo $MIGUEL_DID_FULL | jq -rc .)" --from miguel
# echo "Ledgering DID 2/2..."
# ixod_tx iid create-iid-from-legacy-did "$FRANCESCO_DID_FULL"

ENTITY='{
"entity_type": "assets",
"entity_status": 1,
"owner_did": "did:ixo:4XJLBfGtWSGKSz4BeRxdun",
"owner_address": "ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"
}'
# echo $ENTITY | jq
# yes $PASSWORD | ixod_tx entity create-entity "$(echo $ENTITY | jq -rc .)" --from miguel

# ( echo 12345678; echo y ) | ixod keys delete miguel --force
# (ixod keys delete miguel --force < <(echo 12345678)) < <(echo y)
# yes 'y' | echo $PASSWORD | ixod keys delete miguel --force
# yes 'fall sound heavy fantasy start army shop license insane nuclear emotion execute' | yes $PASSWORD | ixod keys add miguel --recover
# {echo $PASSWORD;echo 'fall sound heavy fantasy start army shop license insane nuclear emotion execute'} | ixod keys add miguel --recover
