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
CHAIN_ID="pandora-4"

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
  # ixod q "$@" --output=json | jq .
  true
}

PROJECT_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
PROJECT_DID_FULL='{
  "did":"did:ixo:U7GK8p8rVhJMKhBVRCJJ8c",
  "verifyKey":"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW",
  "encryptionPublicKey":"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m",
  "secret":{
    "seed":"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053",
    "signKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC",
    "encryptionPrivateKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC"
  }
}'
PROJECT_INFO='{
  "nodeDid":"nodeDid",
  "requiredClaims":"5001",
  "serviceEndpoint":"serviceEndpoint",
  "createdOn":"2020-01-01T01:01:01.000Z",
  "createdBy":"Creator",
  "status":"",
  "fees":{
    "@context":"",
    "items": [
      {"@type":"OracleFee", "id":"payment:template:oracle-fee"},
      {"@type":"FeeForService", "id":"payment:template:fee-for-service"}
    ]
  }
}'

ORACLE_FEE_PAYMENT_TEMPLATE='{
  "id": "payment:template:oracle-fee",
  "payment_amount": [{"denom":"uixo", "amount":"5000000"}],
  "payment_minimum": [{"denom":"uixo", "amount":"5000000"}],
  "payment_maximum": [{"denom":"uixo", "amount":"50000000"}],
  "discounts": []
}'
FEE_FOR_SERVICE_PAYMENT_TEMPLATE='{
  "id": "payment:template:fee-for-service",
  "payment_amount": [{"denom":"uixo", "amount":"2000000"}],
  "payment_minimum": [{"denom":"uixo", "amount":"2000000"}],
  "payment_maximum": [{"denom":"uixo", "amount":"20000000"}],
  "discounts": []
}'

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

# Ledger DIDs
echo "Ledgering Miguel DID..."
ixod_tx iid create-iid-from-legacy-did "$MIGUEL_DID_FULL"
echo "Ledgering Francesco DID..."
ixod_tx iid create-iid-from-legacy-did "$FRANCESCO_DID_FULL"
echo "Ledgering Shaun DID..."
ixod_tx iid create-iid-from-legacy-did "$SHAUN_DID_FULL"

# Create oracle fee payment template
echo "Creating oracle fee payment template..."
CREATOR="$FRANCESCO_DID_FULL"
ixod_tx payments create-payment-template "$ORACLE_FEE_PAYMENT_TEMPLATE" "$CREATOR"

# Create fee-for-service payment template
echo "Creating fee-for-service payment template..."
CREATOR="$FRANCESCO_DID_FULL"
ixod_tx payments create-payment-template "$FEE_FOR_SERVICE_PAYMENT_TEMPLATE" "$CREATOR"

# Create project and progress status to PENDING
SENDER_DID="$SHAUN_DID"
echo "Creating project..."
ixod_tx project create-project "$SENDER_DID" "$PROJECT_INFO" "$PROJECT_DID_FULL"
echo "Updating project to CREATED..."
# # ixod_tx project update-project-status "$SENDER_DID" CREATED "$PROJECT_DID_FULL"
echo "Updating project to PENDING..."
ixod_tx project update-project-status "$SENDER_DID" PENDING "$PROJECT_DID_FULL"

# Updating project doc succeeds as project is in status PENDING and tx sender DID
# # is the same as project sender DID
PROJECT_INFO2='{
  "nodeDid":"nodeDid",
  "requiredClaims":"500",
  "serviceEndpoint":"serviceEndpoint",
  "createdOn":"2020-01-01T01:01:01.000Z",
  "createdBy":"Creator",
  "status":"",
  "fees":{
    "@context":"",
    "items": [
      {"@type":"OracleFee", "id":"payment:template:oracle-fee"},
      {"@type":"FeeForService", "id":"payment:template:fee-for-service"}
    ]
  },
  "newField":"someNewField"
}'
echo "Updating project doc..."
ixod_tx project update-project-doc "$SENDER_DID" "$PROJECT_INFO2" "$PROJECT_DID_FULL"

# Fund project and progress status to FUNDED
FULL_PROJECT_ADDR=$(ixod q project get-project-accounts $PROJECT_DID | grep "$PROJECT_DID")
# Delete longest match of pattern ': ' from the beginning
PROJECT_ADDR=${FULL_PROJECT_ADDR##*: }
echo "Funding project at [$PROJECT_ADDR] with uixo from Francesco..."
# ixod_tx bank send francesco "$PROJECT_ADDR" 100000000uixo
echo "Updating project to FUNDED..."
SENDER_DID="$SHAUN_DID"
ixod_tx project update-project-status "$SENDER_DID" FUNDED "$PROJECT_DID_FULL"

# Progress project status to STARTED
SENDER_DID="$SHAUN_DID"
echo "Updating project to STARTED..."
ixod_tx project update-project-status "$SENDER_DID" STARTED "$PROJECT_DID_FULL"

# If we try updating the project-doc now, the tx fails as the project is in status STARTED
echo "Updating project doc fails since project is in status STARTED..."
ixod_tx project update-project-doc "$SENDER_DID" "$PROJECT_INFO" "$PROJECT_DID_FULL"

# # Create claim and evaluation
echo "Creating a claim in project..."
SENDER_DID="$SHAUN_DID"
ixod_tx project create-claim "tx_hash" "$SENDER_DID" "claim_id" "template_id" "$PROJECT_DID_FULL"
echo "Creating an evaluation in project..."
SENDER_DID="$MIGUEL_DID"
STATUS="1" # create-evaluation updates status of claim from 0 to 1
ixod_tx project create-evaluation "tx_hash" "$SENDER_DID" "claim_id" $STATUS "$PROJECT_DID_FULL"

# OracleFeePercentage:  0.1 (10%)
# NodeFeePercentage:    0.1 (10%)

# Fee for service:   2,000,000 uixo
# Oracle pay:        5,000,000 uixo

# Expected project account balances:
# - InitiatingNodePayFees:   50,000  # 0.1 of 0.1 of oracle pay
# - IxoFees:                      0
# - IxoPayFees:             450,000  # 0.9 of 0.1 of oracle pay
# - project:             93,000,000  # 100IXO - (5+2)IXO
# Expected external account balances:
# - Miguel:       1,000,004,495,000  # 1,000,000,000,000 original +  0.9 of oracle pay - 5,000 tx fee
# - Shaun:        1,000,000,995,000  # 1,000,000,000,000 original + 1.0 of fee-for-service - 1,000,000 project creation fee - 5,000 tx fee

# Sum of fee accounts: 500,000

# Progress project status to PAIDOUT
SENDER_DID="$SHAUN_DID"
echo "Updating project to STOPPED..."
ixod_tx project update-project-status "$SENDER_DID" STOPPED "$PROJECT_DID_FULL"
echo "Updating project to PAIDOUT..."
ixod_tx project update-project-status "$SENDER_DID" PAIDOUT "$PROJECT_DID_FULL"
echo "Project withdrawals query..."
ixod_q project get-project-txs $PROJECT_DID

# Expected withdrawals:
# - 500,000 to ixo (a.k.a Shaun) DID (did:ixo:U4tSpzzv91HHqWW1YmFkHJ)
# Expected project account balances:
# - InitiatingNodePayFees:        0
# - IxoFees:                      0
# - IxoPayFees:                   0
# - project:             93,000,000
# Expected external account balances:
# - Miguel:       1,000,004,495,000
# - Shaun:        1,000,001,495,000  # 500,000 withdrawal

echo "InitiatingNodePayFees"
ixod_q bank balances "ixo1xvjy68xrrtxnypwev9r8tmjys9wk0zkkspzjmq"
echo "IxoFees"
ixod_q bank balances "ixo1ff9we62w6eyes7wscjup3p40vy4uz0sa7j0ajc"
echo "IxoPayFees"
ixod_q bank balances "ixo1udgxtf6yd09mwnnd0ljpmeq4vnyhxdg03uvne3"
echo "(project) did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
ixod_q bank balances "ixo1rmkak6t606wczsps9ytpga3z4nre4z3nwc04p8"
echo "(Miguel) did:ixo:4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_FULL_ADDR="$(ixod q did get-address-from-did $MIGUEL_DID)"
MIGUEL_ADDR=${MIGUEL_FULL_ADDR##*: }
ixod_q bank balances "$MIGUEL_ADDR"
echo "(Shaun) did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
SHAUN_FULL_ADDR="$(ixod q did get-address-from-did $SHAUN_DID)"
SHAUN_ADDR=${SHAUN_FULL_ADDR##*: }
ixod_q bank balances "$SHAUN_ADDR"

# # Withdraw funds (from main project account, i.e. as refund)
# # --> FAIL since Miguel is not the project owner
echo "Withdraw project funds as Miguel (fail since Miguel is not the owner)..."
DATA="{\"projectDid\":\"$PROJECT_DID\",\"recipientDid\":\"$MIGUEL_DID\",\"amount\":\"100000000\",\"isRefund\":true}"
ixod_tx project withdraw-funds "$MIGUEL_DID_FULL" "$DATA"
echo "Project withdrawals query..."
ixod_q project get-project-txs $PROJECT_DID

# # Expected external account balances:
# - Miguel:       1,000,004,490,000 (5,000uixo tx fee deducted)

# Withdraw funds (from main project account, i.e. as refund)
# # --> SUCCESS since Shaun is the project owner
echo "Withdraw project funds as Shaun (success since Shaun is the owner)..."
DATA="{\"projectDid\":\"$PROJECT_DID\",\"recipientDid\":\"$SHAUN_DID\",\"amount\":\"1000000\",\"isRefund\":true}"
ixod_tx project withdraw-funds "$SHAUN_DID_FULL" "$DATA"
echo "Project withdrawals query..."
ixod_q project get-project-txs $PROJECT_DID

# Expected withdrawals:
# - 500,000 to ixo (a.k.a Shaun) DID (did:ixo:U4tSpzzv91HHqWW1YmFkHJ)
# - 1,000,000 to shaun DID (did:ixo:U4tSpzzv91HHqWW1YmFkHJ)
# Expected project account balances:
# - InitiatingNodePayFees:        0
# - IxoFees:                      0
# - IxoPayFees:                   0
# - project:             92,000,000  # 1,000,000 has been withdrawn
# Expected external account balances:
# - Miguel:       1,000,004,490,000
# - Shaun:        1,000,002,490,000  # + 1,000,000 withdrawal - 5,000 fee deducted
