PASSWORD="12345678"

MIGUEL=$(yes $PASSWORD | ixocli keys show miguel -a)

FEE1=$(yes $PASSWORD | ixocli keys show fee -a)
FEE2=$(yes $PASSWORD | ixocli keys show fee2 -a)
FEE3=$(yes $PASSWORD | ixocli keys show fee3 -a)
FEE4=$(yes $PASSWORD | ixocli keys show fee4 -a)

# Power function with m:12,n:2,c:100, rez reserve, non-zero fees, and batch_blocks=1
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
  --signers="$MIGUEL" \
  --batch-blocks=1 \
  --from miguel -y \
  --bond-did="{\"did\":\"U7GK8p8rVhJMKhBVRCJJ8c\",\"verifyKey\":\"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW\",\"encryptionPublicKey\":\"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m\",\"secret\":{\"seed\":\"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053\",\"signKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\",\"encryptionPrivateKey\":\"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC\"}}"
sleep 6

# Power function with m:10,n:3,c:0, res reserve, zero fees, and batch_blocks=3
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
  --signers="$MIGUEL" \
  --batch-blocks=3 \
  --from miguel -y \
  --bond-did="{\"did\":\"JHcN95bkS4aAWk3TKXapA2\",\"verifyKey\":\"ARTUGePyi4rm3ogq3kjp8dEAVq1RR7Z3HWwLze6Ey4qg\",\"encryptionPublicKey\":\"FjLXxW1N68XgBKekfB2isCLHPwbqhLiQLQFhLiiivXqP\",\"secret\":{\"seed\":\"9fc38edde7b9d7097b0aeefcb22a8fcccb6c0748fc3034eec5abdac0740339f7\",\"signKey\":\"Bken38cPuz2Mosb3poKQHd81Q2BiTWFznG1Wa7ZZkrni\",\"encryptionPrivateKey\":\"Bken38cPuz2Mosb3poKQHd81Q2BiTWFznG1Wa7ZZkrni\"}}"
sleep 6

# Swapper function between res and rez with zero fees, and batch_blocks=2
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
  --signers="$MIGUEL" \
  --batch-blocks=2 \
  --from miguel -y \
  --bond-did="{\"did\":\"48PVm1uyF6QVDSPdGRWw4T\",\"verifyKey\":\"2hs2cb232Ev97aSQLvrfK4q8ZceBR8cf33UTstWpKU9M\",\"encryptionPublicKey\":\"9k2THnNbTziXGRjn77tvWujffgigRPqPyKZUSdwjmfh2\",\"secret\":{\"seed\":\"82949a422215a5999846beaadf398659157c345564787993f92e91d192f2a9c5\",\"signKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\",\"encryptionPrivateKey\":\"9njRge76sTYdfcpFfBG5p2NwbDXownFzUyTeN3iDQdjz\"}}"
sleep 6

# Swapper function between token1 and token2 with non-zero fees, and batch_blocks=1
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
  --signers="$MIGUEL" \
  --batch-blocks=1 \
  --from miguel -y \
  --bond-did="{\"did\":\"RYLHkfNpbA8Losy68jt4yF\",\"verifyKey\":\"ENmMCsfNmjYoTRhNgnwXbQAw6p8JKH9DCJfGTPXNfsxW\",\"encryptionPublicKey\":\"5unQBt6JPW1pq9AqoRNhFJmibv8JqeoyyNvN3gF24EaU\",\"secret\":{\"seed\":\"d2c05b107acc2dfe3e9d67e98c993a9c03d227ed8f4505c43997cf4e7819bee2\",\"signKey\":\"FBgjjwoVPfd8ZUVs4nXVKtf4iV6xwnKhaMKBBEAHvGtH\",\"encryptionPrivateKey\":\"FBgjjwoVPfd8ZUVs4nXVKtf4iV6xwnKhaMKBBEAHvGtH\"}}"
sleep 6

# Buy 5token1, 5token2 from Miguel
echo "Buying 5token1 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token1 "100000res" U7GK8p8rVhJMKhBVRCJJ8c --from miguel -y --broadcast-mode block
echo "Buying 5token2 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" JHcN95bkS4aAWk3TKXapA2 --from miguel -y --broadcast-mode block

# Buy token2 and token3 from Francesco and Shaun
echo "Buying 5token2 from Francesco"
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" JHcN95bkS4aAWk3TKXapA2 --from francesco -y --broadcast-mode block
echo "Buying 5token3 from Shaun"
yes $PASSWORD | ixocli tx bonds buy 5token3 "100res,100rez" 48PVm1uyF6QVDSPdGRWw4T --from shaun -y --broadcast-mode block

# Buy 5token4 from Miguel (using token1 and token2)
echo "Buying 5token4 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token4 "2token1,2token2" RYLHkfNpbA8Losy68jt4yF --from miguel -y --broadcast-mode block
