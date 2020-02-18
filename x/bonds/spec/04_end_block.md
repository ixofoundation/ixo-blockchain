# End-Block

At the end of each block, any batch of orders that has reached the end of its lifespan, measured in number of blocks, is cleared. For the rest of the batches, their blocks remaining value is decremented by 1. Orders are performed in the following order:
1. Buys
2. Sells
3. Swaps

Since the buy and sell prices are pre-calculated from when the buy and sell orders were added to the batch, there is no additional cancellations of buys or sells that will take place at this stage. However, swaps are processed on a first come first served basis and a swap is cancelled if it violates the sanity rates.

## Buys

Using the buy price stored in the batch, the following steps are followed for each buy order:
1. Mint and send `n` bond tokens to the buyer
2. Calculate total price`total = r + f` in reserve tokens
   1. `r` is the price of buying `n` bond tokens
   2. `f` is the transactional fee based on `r`
3. Send `r` to the reserve address
4. Send `f` to the fee address
5. Send unused reserve tokens (`maxPrices-total`) back to buyer
6. Increase bond's current supply by `n`

Note: the `maxPrices` reserve tokens were locked upon submitting the buy order.

## Sells

Using the sell price stored in the batch, the following steps are followed for each sell order:
1. Calculate total returns `total = r - f` in reserve tokens
   1. `r` is the return for selling `n` bond tokens
   2. `f` is the transactional and exit fees based on `r`
2. Send `total` to the seller
3. Send `f` to the fee address
4. Decrease bond's current supply by `n`

Note: the `n` bond tokens were burned upon submitting the sell order.

## Swaps

The following steps are followed for each swap order:
1. Calculate the transactional fee `f` based on `t1` reserve tokens
2. Calculate the return `t2` for swapping `t1-f` reserve tokens
3. Check whether the swap violates the sanity rate
   1. Calculate the new reserve balances as a result of the swap
   2. Cancel the swap if the new balances violate the sanity rate
4. Send `t2` to the swapper
5. Send `t1-f` to the reserve address
6. Send `f` to the fee address

Note: the `t1` reserve tokens were locked upon submitting the swap order. If a swap order is cancelled, the `t1` tokens are immediately returned back to the swapper.

## Set Last Batch

Once all orders have been processed, the last batch is set as the current batch and the current batch is cleared in preparation for a new list of orders.