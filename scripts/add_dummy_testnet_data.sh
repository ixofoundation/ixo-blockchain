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
  --broadcast-mode block

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
  --broadcast-mode block

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
  --broadcast-mode block

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
  --broadcast-mode block

# Buy 5token1, 5token2 from Miguel
echo "Buying 5token1 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token1 "100000res" --from miguel -y --broadcast-mode block
echo "Buying 5token2 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" --from miguel -y --broadcast-mode block

# Buy token2 and token3 from Francesco and Shaun
echo "Buying 5token2 from Francesco"
yes $PASSWORD | ixocli tx bonds buy 5token2 "100000res" --from francesco -y --broadcast-mode block
echo "Buying 5token3 from Shaun"
yes $PASSWORD | ixocli tx bonds buy 5token3 "100res,100rez" --from shaun -y --broadcast-mode block

# Buy 5token4 from Miguel (using token1 and token2)
echo "Buying 5token4 from Miguel"
yes $PASSWORD | ixocli tx bonds buy 5token4 "2token1,2token2" --from miguel -y --broadcast-mode block
