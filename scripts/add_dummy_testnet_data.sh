PASSWORD="12345678"

FEE1=$(yes $PASSWORD | ixocli keys show fee -a)
FEE2=$(yes $PASSWORD | ixocli keys show fee2 -a)
FEE3=$(yes $PASSWORD | ixocli keys show fee3 -a)
FEE4=$(yes $PASSWORD | ixocli keys show fee4 -a)

BOND1_DID="{\"did\":\"U7GK8p8rVhJMKhBVRCJJ8c\",\"verifyKey\":\"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW\",\"encryptionPublicKey\":\"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m\",\"secret\":{\"seed\":\"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053\",\"signKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\",\"encryptionPrivateKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\"}}"
BOND2_DID="{\"did\":\"JHcN95bkS4aAWk3TKXapA2\",\"verifyKey\":\"ARTUGePyi4rm3ogq3kjp8dEAVq1RR7Z3HWwLze6Ey4qg\",\"encryptionPublicKey\":\"FjLXxW1N68XgBKekfB2isCLHPwbqhLiQLQFhLiiivXqP\",\"secret\":{\"seed\":\"9fc38edde7b9d7097b0aeefcb22a8fcccb6c0748fc3034eec5abdac0740339f7\",\"signKey\":\"Bken38cPuz2Mosb3poKQHd81Q2BiTWFznG1Wa7ZZkrni\",\"encryptionPrivateKey\":\"Bken38cPuz2Mosb3poKQHd81Q2BiTWFznG1Wa7ZZkrni\"}}"
BOND3_DID="{\"did\":\"48PVm1uyF6QVDSPdGRWw4T\",\"verifyKey\":\"2hs2cb232Ev97aSQLvrfK4q8ZceBR8cf33UTstWpKU9M\",\"encryptionPublicKey\":\"9k2THnNbTziXGRjn77tvWujffgigRPqPyKZUSdwjmfh2\",\"secret\":{\"seed\":\"82949a422215a5999846beaadf398659157c345564787993f92e91d192f2a9c5\",\"signKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\",\"encryptionPrivateKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\"}}"
BOND4_DID="{\"did\":\"RYLHkfNpbA8Losy68jt4yF\",\"verifyKey\":\"ENmMCsfNmjYoTRhNgnwXbQAw6p8JKH9DCJfGTPXNfsxW\",\"encryptionPublicKey\":\"5unQBt6JPW1pq9AqoRNhFJmibv8JqeoyyNvN3gF24EaU\",\"secret\":{\"seed\":\"d2c05b107acc2dfe3e9d67e98c993a9c03d227ed8f4505c43997cf4e7819bee2\",\"signKey\":\"FBgjjwoVPfd8ZUVs4nXVKtf4iV6xwnKhaMKBBEAHvGtH\",\"encryptionPrivateKey\":\"FBgjjwoVPfd8ZUVs4nXVKtf4iV6xwnKhaMKBBEAHvGtH\"}}"

MIGUEL_DID="4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_DID_FULL="{\"did\":\"4XJLBfGtWSGKSz4BeRxdun\",\"verifyKey\":\"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt\",\"encryptionPublicKey\":\"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS\",\"secret\":{\"seed\":\"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180\",\"signKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\",\"encryptionPrivateKey\":\"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh\"}}"
FRANCESCO_DID_FULL="{\"did\":\"UKzkhVSHc3qEFva5EY2XHt\",\"verifyKey\":\"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej\",\"encryptionPublicKey\":\"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si\",\"secret\":{\"seed\":\"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de\",\"signKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\",\"encryptionPrivateKey\":\"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM\"}}"
SHAUN_DID_FULL="{\"did\":\"U4tSpzzv91HHqWW1YmFkHJ\",\"verifyKey\":\"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG\",\"encryptionPublicKey\":\"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i\",\"secret\":{\"seed\":\"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e\",\"signKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\",\"encryptionPrivateKey\":\"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR\"}}"

# Ledger DIDs
echo "Ledgering DID 1/3..."
ixocli tx did addDidDoc "$MIGUEL_DID_FULL" --broadcast-mode block
echo "Ledgering DID 2/3..."
ixocli tx did addDidDoc "$FRANCESCO_DID_FULL" --broadcast-mode block
echo "Ledgering DID 3/3..."
ixocli tx did addDidDoc "$SHAUN_DID_FULL" --broadcast-mode block

# Power function with m:12,n:2,c:100, rez reserve, non-zero fees, and batch_blocks=1
echo "Creating bond 1/4..."
yes $PASSWORD | ixocli tx bonds create-bond \
  --token=token1 \
  --name="Test Token 1" \
  --description="Power function with non-zero fees and batch_blocks=1" \
  --function-type=power_function \
  --function-parameters="m:12,n:2,c:100" \
  --reserve-tokens=res \
  --tx-fee-percentage=0.5 \
  --exit-fee-percentage=0.1 \
  --fee-address="$FEE1" \
  --max-supply=1000000token1 \
  --order-quantity-limits="" \
  --sanity-rate="" \
  --sanity-margin-percentage="" \
  --allow-sells=true \
  --batch-blocks=1 \
  --bond-did="$BOND1_DID" \
  --creator-did="$MIGUEL_DID" \
  --broadcast-mode block

# Power function with m:10,n:3,c:0, res reserve, zero fees, and batch_blocks=3
echo "Creating bond 2/4..."
yes $PASSWORD | ixocli tx bonds create-bond \
  --token=token2 \
  --name="Test Token 2" \
  --description="Power function with zero fees and batch_blocks=4" \
  --function-type=power_function \
  --function-parameters="m:10,n:3,c:1" \
  --reserve-tokens=res \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE2" \
  --max-supply=1000000token2 \
  --order-quantity-limits="" \
  --sanity-rate="" \
  --sanity-margin-percentage="" \
  --allow-sells=true \
  --batch-blocks=3 \
  --bond-did="$BOND2_DID" \
  --creator-did="$MIGUEL_DID" \
  --broadcast-mode block

# Swapper function between res and rez with zero fees, and batch_blocks=2
echo "Creating bond 3/4..."
yes $PASSWORD | ixocli tx bonds create-bond \
  --token=token3 \
  --name="Test Token 3" \
  --description="Swapper function between res and rez" \
  --function-type=swapper_function \
  --function-parameters="" \
  --reserve-tokens="res,rez" \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE3" \
  --max-supply=1000000token3 \
  --order-quantity-limits="" \
  --sanity-rate="" \
  --sanity-margin-percentage="" \
  --allow-sells=true \
  --batch-blocks=2 \
  --bond-did="$BOND3_DID" \
  --creator-did="$MIGUEL_DID" \
  --broadcast-mode block

# Swapper function between token1 and token2 with non-zero fees, and batch_blocks=1
echo "Creating bond 4/4..."
yes $PASSWORD | ixocli tx bonds create-bond \
  --token=token4 \
  --name="Test Token 4" \
  --description="Swapper function between res and rez" \
  --function-type=swapper_function \
  --function-parameters="" \
  --reserve-tokens="token1,token2" \
  --tx-fee-percentage=2.5 \
  --exit-fee-percentage=5 \
  --fee-address="$FEE4" \
  --max-supply=1000000token4 \
  --order-quantity-limits="" \
  --sanity-rate="" \
  --sanity-margin-percentage="" \
  --allow-sells=true \
  --batch-blocks=1 \
  --bond-did="$BOND4_DID" \
  --creator-did="$MIGUEL_DID" \
  --broadcast-mode block

# Buy 5token1, 5token2 from Miguel
echo "Buying 5token1 from Miguel..."
yes $PASSWORD | ixocli tx bonds buy 5token1 "100000res" U7GK8p8rVhJMKhBVRCJJ8c "$MIGUEL_DID_FULL" --broadcast-mode block
echo "Buying 5token2 from Miguel..."
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" JHcN95bkS4aAWk3TKXapA2 "$MIGUEL_DID_FULL" --broadcast-mode block

# Buy token2 and token3 from Francesco and Shaun
echo "Buying 5token2 from Francesco..."
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" JHcN95bkS4aAWk3TKXapA2 "$FRANCESCO_DID_FULL" --broadcast-mode block
echo "Buying 5token3 from Shaun..."
yes $PASSWORD | ixocli tx bonds buy 5token3 "100res,100rez" 48PVm1uyF6QVDSPdGRWw4T "$SHAUN_DID_FULL" --broadcast-mode block

# Buy 5token4 from Miguel (using token1 and token2)
echo "Buying 5token4 from Miguel..."
yes $PASSWORD | ixocli tx bonds buy 5token4 "2token1,2token2" RYLHkfNpbA8Losy68jt4yF "$MIGUEL_DID_FULL" --broadcast-mode block
