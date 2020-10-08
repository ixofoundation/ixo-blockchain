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
yes $PASSWORD | ixocli keys delete fee --force
yes $PASSWORD | ixocli keys add fee
FEE=$(yes $PASSWORD | ixocli keys show fee -a)

OWNER_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
IXO_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
ORACLE_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
PROJECT_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
BOND_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"

OWNER_DID_FULL='{
  "did":"did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
  "verifyKey":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
  "encryptionPublicKey":"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i",
  "secret":{
    "seed":"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e",
    "signKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR",
    "encryptionPrivateKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR"
  }
}'
IXO_DID_FULL='{
  "did":"did:ixo:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'
ORACLE_DID_FULL='{
  "did":"did:ixo:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'
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
  "payment_amount": [{"denom":"uxgbp", "amount":"1000000"}],
  "payment_minimum": [{"denom":"uxgbp", "amount":"1000000"}],
  "payment_maximum": [{"denom":"uxgbp", "amount":"10000000"}],
  "discounts": []
}'

OWNER_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"

# Generate DIDs
DID_1_FULL=$(node utils/did_gen.js)
DID_2_FULL=$(node utils/did_gen.js)
DID_3_FULL=$(node utils/did_gen.js)
DID_4_FULL=$(node utils/did_gen.js)
DID_5_FULL=$(node utils/did_gen.js)
DID_6_FULL=$(node utils/did_gen.js)
DID_7_FULL=$(node utils/did_gen.js)
DID_8_FULL=$(node utils/did_gen.js)
DID_9_FULL=$(node utils/did_gen.js)
DID_10_FULL=$(node utils/did_gen.js)
DID_1=$(echo "$DID_1_FULL" | cut -d \" -f 4)
DID_2=$(echo "$DID_2_FULL" | cut -d \" -f 4)
DID_3=$(echo "$DID_3_FULL" | cut -d \" -f 4)
DID_4=$(echo "$DID_4_FULL" | cut -d \" -f 4)
DID_5=$(echo "$DID_5_FULL" | cut -d \" -f 4)
DID_6=$(echo "$DID_6_FULL" | cut -d \" -f 4)
DID_7=$(echo "$DID_7_FULL" | cut -d \" -f 4)
DID_8=$(echo "$DID_8_FULL" | cut -d \" -f 4)
DID_9=$(echo "$DID_9_FULL" | cut -d \" -f 4)
DID_10=$(echo "$DID_10_FULL" | cut -d \" -f 4)

# Ledger the 10 DIDs, oracle, ixo, and owner DIDs
echo "Ledgering DIDs..."
ixocli tx did add-did-doc "$DID_1_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_2_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_3_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_4_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_5_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_6_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_7_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_8_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_9_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$DID_10_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$ORACLE_DID_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$IXO_DID_FULL" --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx did add-did-doc "$OWNER_DID_FULL" --gas-prices="$GAS_PRICES" -y --broadcast-mode block > /dev/null

# Fund oracle for gas fees and ixo DID for funding and gas fees
echo "Funding oracle and ixo DID..."
yes $PASSWORD | ixocli tx send "$(ixocli keys show miguel -a)" "$(ixocli q did get-address-from-did $ORACLE_DID)" 1000000uixo --fees=5000uixo --broadcast-mode=block -y > /dev/null
yes $PASSWORD | ixocli tx send "$(ixocli keys show miguel -a)" "$(ixocli q did get-address-from-did $IXO_DID)" 10000000000uixo --fees=5000uixo --broadcast-mode=block -y > /dev/null

# Fund DID accounts using ixo DID for gas fees
echo "Funding DID accounts..."
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_1" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_2" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_3" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_4" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_5" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_6" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_7" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_8" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_9" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$DID_10" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$OWNER_DID" 10000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Each DID now has 10IXO for gas fees
ixocli q account "$(ixocli q did get-address-from-did $DID_1)"

# Fund Owner with 300xGBP (300000000uxgbp)
echo "Funding Owner DID with 300xGBP (using treasury 'oracle-mint' using oracle)..."
ixocli tx treasury oracle-mint "$OWNER_ADDR" 300000000uxgbp "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Owner now has 10IXO and 300xGBP
ixocli q account "$(ixocli q did get-address-from-did $OWNER_DID)"

# Create bond
echo "Creating bond..."
ixocli tx bonds create-bond \
  --token=uabc \
  --name="ABC" \
  --description="Description about ABC" \
  --function-type=augmented_function \
  --function-parameters="d0:1000000,p0:1,theta:0,kappa:3.0" \
  --reserve-tokens=uxgbp \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --max-supply=20000000uabc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --batch-blocks=1 \
  --outcome-payment="300000000" \
  --bond-did="$BOND_DID" \
  --creator-did="$OWNER_DID_FULL" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create oracle fee payment template
echo "Creating oracle fee payment template..."
CREATOR="$IXO_DID_FULL"
ixocli tx payments create-payment-template "$ORACLE_FEE_PAYMENT_TEMPLATE" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create fee-for-service payment template
echo "Creating fee-for-service payment template..."
CREATOR="$IXO_DID_FULL"
ixocli tx payments create-payment-template "$FEE_FOR_SERVICE_PAYMENT_TEMPLATE" "$CREATOR" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create project and progress status to PENDING
SENDER_DID="$OWNER_DID"
echo "Creating project..."
ixocli tx project create-project "$SENDER_DID" "$PROJECT_INFO" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to CREATED..."
ixocli tx project update-project-status "$SENDER_DID" CREATED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to PENDING..."
ixocli tx project update-project-status "$SENDER_DID" PENDING "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Fund project with 100xGBP and 1000 IXO (for fees)
PROJECT_ADDR=$(ixocli q project get-project-accounts $PROJECT_DID | grep "$PROJECT_DID" | cut -d \" -f 4)
echo "Funding project with uixo and uxgbp (using treasury 'oracle-mint' and 'oracle-transfer' using oracle)..."
ixocli tx treasury oracle-mint "$PROJECT_ADDR" 100000000uxgbp "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-transfer "$IXO_DID" "$PROJECT_ADDR" 1000000000uixo "$ORACLE_DID_FULL" "proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Progress project status to FUNDED and STARTED
SENDER_DID="$OWNER_DID"
echo "Updating project to FUNDED..."
ixocli tx project update-project-status "$SENDER_DID" FUNDED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to STARTED..."
ixocli tx project update-project-status "$SENDER_DID" STARTED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Create claims
echo "Creating claims in project..."
ixocli tx project create-claim "tx_hash" "$DID_1" "claim1" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_2" "claim2" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_3" "claim3" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_4" "claim4" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_5" "claim5" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_6" "claim6" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_7" "claim7" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_8" "claim8" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_9" "claim9" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_10" "claim10" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Create evaluations
echo "Creating evaluations in project..."
STATUS="1"
ixocli tx project create-evaluation "tx_hash" "$DID_1" "claim1" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_2" "claim2" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_3" "claim3" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_4" "claim4" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_5" "claim5" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_6" "claim6" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_7" "claim7" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_8" "claim8" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_9" "claim9" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx project create-evaluation "tx_hash" "$DID_10" "claim10" "$STATUS" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Each of the 10 DIDs now has 1xGBP (1000000uxgbp)
ixocli q account "$(ixocli q did get-address-from-did "$DID_1")"

# Perform bond buys
echo "DID 1 buys 1ABC..."
ixocli tx bonds buy 1000000uabc 1000000uxgbp "$BOND_DID" "$DID_1_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 2 buys 0.26ABC..."
ixocli tx bonds buy 259921uabc 1000000uxgbp "$BOND_DID" "$DID_2_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 3 buys 0.18ABC..."
ixocli tx bonds buy 182328uabc 1000000uxgbp "$BOND_DID" "$DID_3_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 4 buys 0.14ABC..."
ixocli tx bonds buy 145151uabc 1000000uxgbp "$BOND_DID" "$DID_4_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 5 buys 0.12ABC..."
ixocli tx bonds buy 122574uabc 1000000uxgbp "$BOND_DID" "$DID_5_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Make outcome payment
echo "Owner makes outcome payment..."
ixocli tx bonds make-outcome-payment "$BOND_DID" "$OWNER_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Withdraw reserve shares
echo "Withdrawing shares (DID 1-5)..."
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_1_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_2_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_3_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_4_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_5_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
