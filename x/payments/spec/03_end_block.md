# End-Block

At the end of each block, any subscription that satisfies a set of criteria
(the `ShouldEffect` function) gets effected. These criteria are that (i)
the subscription has started, and that (iiA) the maximum number of periods has
not been reached and the period has ended, or (iiB) the max has been reached,
but accumulated (unpaid) periods exist.

To effect a subscription payment, the subscription is loaded and the payment
contract that it points to is in turn effected. If the payment does not trigger
errors, we can proceed. If the maximum number of periods had not been reached,
then we can move to the next period (note: even if payment was not effected).
Otherwise, we are trying to effect a previously unpaid payment, and so if this
was successful, we deduct one from the accumulated (unpaid) payments.

Moving to the next period involves three general steps:

1. Increment the number of periods so far
1. If payment was _not_ effected, increment the number of accumulated periods
1. The period start is set to the period end (old end is the new start)

Feel free to refer to [01_state.md](./01_state.md) for more general information
about subscriptions.
