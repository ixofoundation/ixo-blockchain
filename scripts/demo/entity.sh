#!/usr/bin/env bash

# Must be run from root path inside ixo-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

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
# echo $MIGUEL_DID_FULL | jq -rc .
# ixod_tx iid create-iid-from-legacy-did "$MIGUEL_DID_FULL"
# yes $PASSWORD | ixod_tx iid create-iid "$(echo $MIGUEL_DID_FULL | jq -rc .)" --from miguel
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
