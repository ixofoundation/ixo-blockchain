# Future Improvements

- **Order processing and front-running prevention**: Improved order fulfillment procedure with less cancellations and more options for the user when buying/selling/swapping, such as minimum returns, specifying amount to be spent rather than bought, etc. The intention is primarily to improve user experience. The main challenge lies in doing this without compromising on front-running prevention and order batching in general. More options for the user means more ways in which an order can be cancelled, and any cancelled order will affect the fulfillability of other orders, which may in turn get cancelled, and so on. One option would be to have an exchange-like behaviour and postpone orders that cannot be fulfilled to the next batch, which then runs into complications of dealing with stale orders. On a similar note, work can be done towards implementing front-running prevention for swap orders [1].
- **Bond creation and function types**: More function types and an improved bond creation process, with more options for the creator and smarter parameter restrictions. An interesting function type that can be implemented is a rule-based function [2].
- **IBC**: The availability of Inter-Blockchain Communication will unlock the full potential of the bonds module. On top of being able to create any bond, one will be able to use tokens from other chains as reserve tokens for the created bonds and transfer the bond tokens across chains. Further work would need to be done to ensure compatibility with IBC.

## References

1. https://ethresear.ch/t/improving-front-running-resistance-of-x-y-k-market-makers/1281
2. https://medium.com/thoughtchains/on-single-bonding-curves-for-continuous-token-models-a167f5ffef89