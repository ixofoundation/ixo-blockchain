# Events

The bonds module emits the following events:

## EndBlocker

| Type          | Attribute Key     | Attribute Value     |
|---------------|-------------------|---------------------|
| order_cancel  | bond              | {token}             |
| order_cancel  | order_type        | {orderType}         |
| order_cancel  | address           | {address}           |
| order_cancel  | cancel_reason     | {cancelReason}      |
| order_fulfill | bond              | {token}             |
| order_fulfill | order_type        | {orderType}         |
| order_fulfill | address           | {address}           |
| order_fulfill | tokensMinted      | {tokensMinted}      |
| order_fulfill | chargedPrices     | {chargedPrices}     |
| order_fulfill | chargedFees       | {chargedFees}       |
| order_fulfill | returnedToAddress | {returnedToAddress} |
| state_change  | bond              | {token}             |
| state_change  | old_state         | {oldState}          |
| state_change  | new_state         | {newState}          |

## Handlers

### MsgCreateBond

| Type        | Attribute Key              | Attribute Value            |
|-------------|----------------------------|----------------------------|
| create_bond | bond_did                   | {bondDid}                  |
| create_bond | token                      | {token}                    |
| create_bond | name                       | {name}                     |
| create_bond | description                | {description}              |
| create_bond | function_type              | {functionType}             |
| create_bond | function_parameters [0]    | {functionParameters}       |
| create_bond | creator_did                | {creatorDid}               |
| create_bond | controller_did             | {controllerDid}            |
| create_bond | reserve_tokens [1]         | {reserveTokens}            |
| create_bond | tx_fee_percentage          | {txFeePercentage}          |
| create_bond | exit_fee_percentage        | {exitFeePercentage}        |
| create_bond | fee_address                | {feeAddress}               |
| create_bond | reserve_withdrawal_address | {reserveWithdrawalAddress} |
| create_bond | max_supply                 | {maxSupply}                |
| create_bond | order_quantity_limits      | {orderQuantityLimits}      |
| create_bond | sanity_rate                | {sanityRate}               |
| create_bond | sanity_margin_percentage   | {sanityMarginPercentage}   |
| create_bond | allow_sells                | {allowSells}               |
| create_bond | allow_reserve_withdrawals  | {allowReserveWithdrawals}  |
| create_bond | alpha_bond                 | {alphaBond}                |
| create_bond | batch_blocks               | {batchBlocks}              |
| create_bond | outcome_payment            | {outcomePayment}           |
| create_bond | state                      | {state}                    |
| message     | module                     | bonds                      |
| message     | sender                     | {creatorDid}               |

* [0] Example formatting: `"{m:12,n:2,c:100}"`
* [1] Example formatting: `"[res,rez]"`

### MsgEditBond

| Type      | Attribute Key            | Attribute Value          |
|-----------|--------------------------|--------------------------|
| edit_bond | bond_did                 | {bondDid}                |
| edit_bond | name                     | {name}                   |
| edit_bond | description              | {description}            |
| edit_bond | order_quantity_limits    | {orderQuantityLimits}    |
| edit_bond | sanity_rate              | {sanityRate}             |
| edit_bond | sanity_margin_percentage | {sanityMarginPercentage} |
| message   | module                   | bonds                    |
| message   | sender                   | {editorDid}              |

### MsgSetNextAlpha

| Type           | Attribute Key            | Attribute Value          |
|----------------|--------------------------|--------------------------|
| set_next_alpha | bond_did                 | {bondDid}                |
| set_next_alpha | public_alpha             | {name}                   |
| message        | module                   | bonds                    |
| message        | sender                   | {setterDid}              |

### MsgUpdateBondState

| Type      | Attribute Key            | Attribute Value          |
|-----------|--------------------------|--------------------------|
| message   | module                   | bonds                    |
| message   | sender                   | {editorDid}              |

### MsgBuy

#### First Buy for Swapper Function Bond

| Type         | Attribute Key  | Attribute Value |
|--------------|----------------|-----------------|
| init_swapper | bond_did       | {bondDid}       |
| init_swapper | amount         | {amount}        |
| init_swapper | charged_prices | {chargedPrices} |
| message      | module         | bonds           |
| message      | sender         | {buyerDid}      |

#### Otherwise

| Type         | Attribute Key | Attribute Value |
|--------------|---------------|-----------------|
| buy          | bond_did      | {bondDid}       |
| buy          | amount        | {amount}        |
| buy          | max_prices    | {maxPrices}     |
| order_cancel | bond          | {token}         |
| order_cancel | order_type    | {orderType}     |
| order_cancel | address       | {address}       |
| order_cancel | cancel_reason | {cancelReason}  |
| message      | module        | bonds           |
| message      | sender        | {buyerDid}      |

### MsgSell

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| sell    | bond_did      | {bondDid}       |
| sell    | amount        | {amount}        |
| message | module        | bonds           |
| message | sender        | {sellerDid}     |

### MsgSwap

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| swap    | bond_did      | {bondDid}       |
| swap    | amount        | {amount}        |
| swap    | from_token    | {fromToken}     |
| swap    | to_token      | {toToken}       |
| message | module        | bonds           |
| message | sender        | {swapperDid}    |

### MsgMakeOutcomePayment

| Type                 | Attribute Key | Attribute Value      |
|----------------------|---------------|----------------------|
| make_outcome_payment | bond_did      | {bondDid}            |
| make_outcome_payment | amount        | {amount}             |
| make_outcome_payment | address       | {senderAddress}      |
| message              | module        | bonds                |
| message              | sender        | {senderDid}          |

### MsgWithdrawShare

| Type           | Attribute Key | Attribute Value    |
|----------------|---------------|--------------------|
| withdraw_share | bond_did      | {bondDid}          |
| withdraw_share | address       | {recipientAddress} |
| withdraw_share | amount        | {reserveOwed}      |
| message        | module        | bonds              |
| message        | sender        | {recipientDid}     |

### MsgWithdrawReserve

| Type             | Attribute Key | Attribute Value            |
|------------------|---------------|----------------------------|
| withdraw_reserve | bond_did      | {bondDid}                  |
| withdraw_reserve | address       | {reserveWithdrawalAddress} |
| withdraw_reserve | amount        | {amountWithdrawn}          |
| message          | module        | bonds                      |
| message          | sender        | {withdrawerDid}            |
