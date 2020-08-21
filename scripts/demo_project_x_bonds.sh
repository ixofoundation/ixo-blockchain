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

SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
ORACLE_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
PROJECT_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
BOND_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"

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
  "evaluatorPayPerClaim":"50000000uixo",
  "claimerPayPerClaim":"1000000uxgbp",
  "serviceEndpoint":"serviceEndpoint",
  "createdOn":"2020-01-01T01:01:01.000Z",
  "createdBy":"Creator",
  "status":""
}'

SHAUN_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"

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

# Ledger DIDs
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
ixocli tx did add-did-doc "$SHAUN_DID_FULL" --gas-prices="$GAS_PRICES" -y --broadcast-mode block > /dev/null

# Fund oracle for fees
echo "Funding oracle for fees..."
yes $PASSWORD | ixocli tx send "$(ixocli keys show miguel -a)" "$(ixocli q did get-address-from-did $ORACLE_DID)" 1000000uixo --fees=5000uixo --broadcast-mode=block -y > /dev/null

# Fund DID accounts
echo "Funding DID accounts..."
ixocli tx treasury oracle-mint "$DID_1" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_2" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_3" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_4" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_5" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_6" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_7" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_8" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_9" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$DID_10" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx treasury oracle-mint "$SHAUN_DID" 10000000uixo "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Fund Shaun with 300xGBP (300000000uxgbp)
echo "Funding Shaun DID with 300xGBP (using treasury 'oracle-mint' using oracle)..."
ixocli tx treasury oracle-mint "$SHAUN_ADDR" 300000000uxgbp "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Create bond
echo "Creating bond..."
ixocli tx bonds create-bond \
  --token=uabc \
  --name="ABC" \
  --description="Description about ABC" \
  --function-type=augmented_function \
  --function-parameters="d0:10000000,p0:1,theta:0,kappa:3.0" \
  --reserve-tokens=uxgbp \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --max-supply=20000000uabc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --batch-blocks=1 \
  --outcome-payment="300000000uxgbp" \
  --bond-did="$BOND_DID" \
  --creator-did="$SHAUN_DID_FULL" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Create project and progress status to PENDING
SENDER_DID="$SHAUN_DID"
echo "Creating project..."
ixocli tx project create-project "$SENDER_DID" "$PROJECT_INFO" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to CREATED..."
ixocli tx project update-project-status "$SENDER_DID" CREATED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to PENDING..."
ixocli tx project update-project-status "$SENDER_DID" PENDING "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Fund project with 100xGBP (100000000uxgbp) and uixo (for fees)
PROJECT_ADDR=$(ixocli q project get-project-accounts $PROJECT_DID | grep "$PROJECT_DID" | cut -d \" -f 4)
echo "Funding project with uixo and uxgbp (using treasury 'oracle-mint' using oracle)..."
ixocli tx treasury oracle-mint "$PROJECT_ADDR" 10000000000uixo,100000000uxgbp "$ORACLE_DID_FULL" "dummy proof" --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# Progress project status to FUNDED and STARTED
SENDER_DID="$SHAUN_DID"
echo "Updating project to FUNDED..."
ixocli tx project update-project-status "$SENDER_DID" FUNDED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "Updating project to STARTED..."
ixocli tx project update-project-status "$SENDER_DID" STARTED "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Create claims
echo "Creating claims in project..."
ixocli tx project create-claim "tx_hash" "$DID_1" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_2" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_3" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_4" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_5" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_6" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_7" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_8" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_9" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null
ixocli tx project create-claim "tx_hash" "$DID_10" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" --gas=auto --gas-adjustment=1.5 -y > /dev/null

# Perform bond buys
echo "DID 1 buys 10ABC..."
ixocli tx bonds buy 10000000uabc 10000000uxgbp "$BOND_DID" "$DID_1_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 2 buys 2.6ABC..."
ixocli tx bonds buy 2599210uabc 10000000uxgbp "$BOND_DID" "$DID_2_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 3 buys 1.8ABC..."
ixocli tx bonds buy 1823285uabc 10000000uxgbp "$BOND_DID" "$DID_3_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 4 buys 1.4ABC..."
ixocli tx bonds buy 1451514uabc 10000000uxgbp "$BOND_DID" "$DID_4_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
echo "DID 5 buys 1.2ABC..."
ixocli tx bonds buy 1225749uabc 10000000uxgbp "$BOND_DID" "$DID_5_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Make outcome payment
echo "Shaun makes outcome payment..."
ixocli tx bonds make-outcome-payment "$BOND_DID" "$SHAUN_DID_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null

# Withdraw reserve shares
echo "Withdrawing shares (DID 1-5)..."
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_1_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_2_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_3_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_4_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
ixocli tx bonds withdraw-share "$BOND_DID" "$DID_5_FULL" --broadcast-mode block --gas-prices="$GAS_PRICES" -y > /dev/null
