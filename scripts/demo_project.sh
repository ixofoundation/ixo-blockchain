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

PROJECT_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
PROJECT_DID_FULL="{\"did\":\"did:ixo:U7GK8p8rVhJMKhBVRCJJ8c\",\"verifyKey\":\"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW\",\"encryptionPublicKey\":\"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m\",\"secret\":{\"seed\":\"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053\",\"signKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\",\"encryptionPrivateKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\"}}"
PROJECT_INFO="{\"nodeDid\":\"nodeDid\",\"requiredClaims\":\"500\",\"evaluatorPayPerClaim\":\"50\",\"serviceEndpoint\":\"serviceEndpoint\",\"createdOn\":\"2020-01-01T01:01:01.000Z\",\"createdBy\":\"Miguel\",\"status\":\"\"}"

MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_DID_FULL="{\"did\":\"did:ixo:4XJLBfGtWSGKSz4BeRxdun\",\"verifyKey\":\"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt\",\"encryptionPublicKey\":\"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS\",\"secret\":{\"seed\":\"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180\",\"signKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\",\"encryptionPrivateKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\"}}"
FRANCESCO_DID_FULL="{\"did\":\"did:ixo:UKzkhVSHc3qEFva5EY2XHt\",\"verifyKey\":\"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej\",\"encryptionPublicKey\":\"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si\",\"secret\":{\"seed\":\"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de\",\"signKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\",\"encryptionPrivateKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\"}}"
SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"

# Ledger DIDs
echo "Ledgering Miguel DID..."
ixocli tx did addDidDoc "$MIGUEL_DID_FULL"
echo "Ledgering Francesco DID..."
ixocli tx did addDidDoc "$FRANCESCO_DID_FULL" --broadcast-mode block

# Create project and progress status to PENDING
SENDER_DID="$SHAUN_DID"
echo "Creating project..."
ixocli tx project createProject "$SENDER_DID" "$PROJECT_INFO" "$PROJECT_DID_FULL" --broadcast-mode block
echo "Updating project to CREATED..."
ixocli tx project updateProjectStatus "$SENDER_DID" CREATED "$PROJECT_DID_FULL" --broadcast-mode block
echo "Updating project to PENDING..."
ixocli tx project updateProjectStatus "$SENDER_DID" PENDING "$PROJECT_DID_FULL" --broadcast-mode block

# Fund project and progress status to FUNDED
echo "Funding project (using treasury 'oracle-transfer' from Miguel using Francesco oracle)..."
SENDER_DID="$SHAUN_DID"
ixocli tx treasury oracle-transfer "$MIGUEL_DID" "$PROJECT_DID/$PROJECT_DID" 10000000000ixo "$FRANCESCO_DID_FULL" "dummy proof" --broadcast-mode block
echo "Updating project to FUNDED..."
ixocli tx project updateProjectStatus "$SENDER_DID" FUNDED "$PROJECT_DID_FULL" --broadcast-mode block

# Create claim and evaluation
echo "Creating a claim in project..."
SENDER_DID="$SHAUN_DID"
ixocli tx project createClaim "tx_hash" "$SENDER_DID" "claim_id" "$PROJECT_DID_FULL" --broadcast-mode block
echo "Creating an evaluation in project..."
SENDER_DID="$MIGUEL_DID"
STATUS="1" # createEvaluation updates status of claim from 0 to 1 implicitly (explicitly in blocksync)
ixocli tx project createEvaluation "tx_hash" "$SENDER_DID" "claim_id" $STATUS "$PROJECT_DID_FULL" --broadcast-mode block

# Expected InitiatingNodePayFees:  100000000
# Expected IxoFees:                 50000000
# Expected IxoPayFees:             400000000
# Expected ValidatingNodeSetFees:   50000000
# Expected evaluator (Miguel):    4500000000
# Expected project:               4900000000

# Progress project status to PAIDOUT
SENDER_DID="$SHAUN_DID"
echo "Updating project to STARTED..."
ixocli tx project updateProjectStatus "$SENDER_DID" STARTED "$PROJECT_DID_FULL" --broadcast-mode block
echo "Updating project to STOPPED..."
ixocli tx project updateProjectStatus "$SENDER_DID" STOPPED "$PROJECT_DID_FULL" --broadcast-mode block
echo "Updating project to PAIDOUT..."
ixocli tx project updateProjectStatus "$SENDER_DID" PAIDOUT "$PROJECT_DID_FULL" --broadcast-mode block
echo "Project withdrawals query..."
ixocli q project getProjectTxs $PROJECT_DID

# Expected InitiatingNodePayFees:          0
# Expected IxoFees:                        0
# Expected IxoPayFees:                     0
# Expected ValidatingNodeSetFees:          0
# Expected evaluator (Miguel):    4500000000
# Expected project:               4900000000

echo "InitiatingNodePayFees"
ixocli q auth account "ixo1xvjy68xrrtxnypwev9r8tmjys9wk0zkkspzjmq"
echo "IxoFees"
ixocli q auth account "ixo1ff9we62w6eyes7wscjup3p40vy4uz0sa7j0ajc"
echo "IxoPayFees"
ixocli q auth account "ixo1udgxtf6yd09mwnnd0ljpmeq4vnyhxdg03uvne3"
echo "ValidatingNodeSetFees"
ixocli q auth account "ixo16dxuhf9hzqtfks2s0azwetl59ldwp5tl5kcxa4"
echo "did:ixo:4XJLBfGtWSGKSz4BeRxdun"
ixocli q auth account "ixo1545vwm70g6tws43a3epakzwyhpj7hxcs7qgpcv"
echo "did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
ixocli q auth account "ixo1rmkak6t606wczsps9ytpga3z4nre4z3nwc04p8"

# Withdraw funds
DATA="{\"projectDid\":\"$PROJECT_DID\",\"recipientDid\":\"$MIGUEL_DID\",\"amount\":\"100000000\",\"isRefund\":false}"
ixocli tx project withdraw-funds "$MIGUEL_DID_FULL" "$DATA" --broadcast-mode block
echo "Project withdrawals query..."
ixocli q project getProjectTxs $PROJECT_DID

# Expected InitiatingNodePayFees:          0
# Expected IxoFees:                        0
# Expected IxoPayFees:                     0
# Expected ValidatingNodeSetFees:          0
# Expected evaluator (Miguel):    4400000000
# Expected project:               4900000000

# Withdraw funds
DATA="{\"projectDid\":\"$PROJECT_DID\",\"recipientDid\":\"$MIGUEL_DID\",\"amount\":\"100000000\",\"isRefund\":true}"
ixocli tx project withdraw-funds "$MIGUEL_DID_FULL" "$DATA" --broadcast-mode block
echo "Project withdrawals query..."
ixocli q project getProjectTxs $PROJECT_DID

# Expected InitiatingNodePayFees:          0
# Expected IxoFees:                        0
# Expected IxoPayFees:                     0
# Expected ValidatingNodeSetFees:          0
# Expected evaluator (Miguel):    4400000000
# Expected project:               4800000000
