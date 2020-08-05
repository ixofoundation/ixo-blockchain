# State

## Bonds

The instance of a bond is stored with its bond-specific parameters. This record is accessed by the identity of a token that represents the bond.

- Bonds: `0x00 | tokenHash -> amino(Bond)`

## Batches

As a protection against front-runnning orders, a batching mechanism creates a cache of orders and combines these into a single transaction when the batch conditions have been met.
The state of 2 consecutive batches is held for both the current and last (previous) batch. 
This enables querying the final state of a batch before the orders were fulfilled, after the transaction has completed. 
The temporary state of a batch in the current block is not observable. This batch is cleared as soon as the batch transaction has completed.

### Querying Batches

Batches are accessed by the identity token of the bond.

- Current Batches: `0x01 | tokenHash -> amino(Batch) `

- Last Batches: `0x02 | tokenHash -> amino(Batch) `
