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

| Type        | Attribute Key            | Attribute Value          |
|-------------|--------------------------|--------------------------|
| create_bond | bond                     | {token}                  |
| create_bond | name                     | {name}                   |
| create_bond | description              | {description}            |
| create_bond | function_type            | {functionType}           |
| create_bond | function_parameters [0]  | {functionParameters}     |
| create_bond | reserve_tokens [1]       | {reserveTokens}          |
| create_bond | tx_fee_percentage        | {txFeePercentage}        |
| create_bond | exit_fee_percentage      | {exitFeePercentage}      |
| create_bond | fee_address              | {feeAddress}             |
| create_bond | max_supply               | {maxSupply}              |
| create_bond | order_quantity_limits    | {orderQuantityLimits}    |
| create_bond | sanity_rate              | {sanityRate}             |
| create_bond | sanity_margin_percentage | {sanityMarginPercentage} |
| create_bond | allow_sells              | {allowSells}             |
| create_bond | signers [2]              | {signers}                |
| create_bond | batch_blocks             | {batchBlocks}            |
| create_bond | state                    | {state}                  |
| message     | module                   | bonds                    |
| message     | action                   | create_bond              |
| message     | sender                   | {senderAddress}          |

* [0] Example formatting: `"{m:12,n:2,c:100}"`
* [1] Example formatting: `"[res,rez]"`
* [2] Example formatting: `"[ADDR1,ADDR2]"`

### MsgEditBond

| Type      | Attribute Key            | Attribute Value          |
|-----------|--------------------------|--------------------------|
| edit_bond | bond                     | {token}                  |
| edit_bond | name                     | {name}                   |
| edit_bond | description              | {description}            |
| edit_bond | order_quantity_limits    | {orderQuantityLimits}    |
| edit_bond | sanity_rate              | {sanityRate}             |
| edit_bond | sanity_margin_percentage | {sanityMarginPercentage} |
| message   | module                   | bonds                    |
| message   | action                   | edit_bond                |
| message   | sender                   | {senderAddress}          |

### MsgBuy

#### First Buy for Swapper Function Bond

| Type         | Attribute Key  | Attribute Value |
|--------------|----------------|-----------------|
| init_swapper | bond           | {token}         |
| init_swapper | amount         | {amount}        |
| init_swapper | charged_prices | {chargedPrices} |
| message      | module         | bonds           |
| message      | action         | buy             |
| message      | sender         | {senderAddress} |

#### Otherwise

| Type         | Attribute Key | Attribute Value |
|--------------|---------------|-----------------|
| buy          | bond          | {token}         |
| buy          | amount        | {amount}        |
| buy          | max_prices    | {maxPrices}     |
| order_cancel | bond          | {token}         |
| order_cancel | order_type    | {orderType}     |
| order_cancel | address       | {address}       |
| order_cancel | cancel_reason | {cancelReason}  |
| message      | module        | bonds           |
| message      | action        | buy             |
| message      | sender        | {senderAddress} |

### MsgSell

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| sell    | bond          | {token}         |
| sell    | amount        | {amount}        |
| message | module        | bonds           |
| message | action        | buy             |
| message | sender        | {senderAddress} |

### MsgSwap

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| swap    | bond          | {token}         |
| swap    | amount        | {amount}        |
| swap    | from_token    | {fromToken}     |
| swap    | to_token      | {toToken}       |
| message | module        | bonds           |
| message | action        | swap            |
| message | sender        | {senderAddress} |
