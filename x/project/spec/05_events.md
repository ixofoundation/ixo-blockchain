# Events

The project module emits the following events:

## EndBlocker

| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| create_project  | tx_hash                | {tx_hash}           |
| create_project  | sender_did      | {sender_did} |
| create_project  | project_did               | {project_did}          |
| create_project  | pub_key            | {pub_key}        |
| update_project_status | TxHash                     | {token}               |
| update_project_status | SenderDid               | {orderType}           |
| update_project_status | ProjectDid                  | {address}             |
| update_project_status | EthFundingTxnID             | {tokensMinted}        |
| update_project_status | Status            | {chargedPrices}       |


