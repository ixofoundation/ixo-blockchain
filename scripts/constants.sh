#!/usr/bin/env bash

PASSWORD="12345678"
GAS_PRICES="0.025uixo"
CHAIN_ID="devnet-1"
NODE="https://devnet.ixo.earth:443/rpc/"

USERS=(alice bob charlie miguel francesco shaun fee fee2 fee3 fee4 fee5 reserveOut)
MNEMONICS=(
  'fall sound heavy fantasy start army shop license insane nuclear emotion execute'
  'genuine suspect someone trip school order amateur heart cheap similar creek turn'
  'faith game good hover hope area detect scout easily filter piece measure'
  'jungle brave person inmate dirt upset try rotate twin fossil grid border'
  'solution fame hundred price survey quantum swear grass opinion spot image figure'
  'pool announce mandate pride pill virus remind valid sunny length embrace avoid'
  'pilot sniff focus engine gym puppy special cat surround decline buzz morning'
  'perfect topic area embark lawsuit list polar solar special brief gas latin'
  'soccer walk grain purse ankle chef fade conduct pepper evil expect super'
  'rare wing climb shadow casino original such fade film rough egg frown'
  'slush slight giraffe table strong vintage media nut just shoe emerge point'
  'current little pave couple concert much success wreck price found gadget effort'
)

wait_chain_start() {
  RET=$(ixod status 2>&1)
  if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
    while :; do
      RET=$(ixod status 2>&1)
      if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
        sleep 1
      else
        echo "A few more seconds..."
        sleep 7
        break
      fi
    done
  fi
}

ixod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args
  # Get module ($1) and specific tx ($2), which forms the tx command
  cmd="$1 $2"
  shift 2

  yes $PASSWORD | ixod tx $cmd \
    --gas-prices $GAS_PRICES \
    --chain-id $CHAIN_ID \
    --broadcast-mode block \
    -y \
    "$@" | jq .
  # The $@ adds any extra arguments to the end
  # --node="$NODE" \
}

ixod_q() {
  ixod q \
    "$@" \
    --output=json | jq .
  # --node="$NODE" \
}

# TRY CATCH implementation following link below
# https://www.xmodulo.com/catch-handle-errors-bash.html
try() {
  [[ $- = *e* ]]
  SAVED_OPT_E=$?
  set +e
}
throw() {
  exit $1
}
catch() {
  export exception_code=$?
  (($SAVED_OPT_E)) && set +e
  return $exception_code
}

full_iid_doc() {
  # Helper function to create a full iid doc => full_iid_doc did address pubkeyBase58
  local DID=$1
  local ADDRESS=$2
  local PUBKEY=$3

  local DID_FULL=$(
    cat <<-END
    {
      "id": "${DID}",
      "controllers": ["${DID}"],
      "verifications": [
        {
          "method": {
            "id": "${DID}",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "controller": "${DID}",
            "publicKeyBase58": "${PUBKEY}"
          },
          "relationships": ["authentication"],
          "context": []
        }
      ],
      "context": [],
      "services": [],
      "accorded_right": [],
      "linked_resources": [],
      "linked_entity": []
    }
END
  )

  echo $DID_FULL
}
