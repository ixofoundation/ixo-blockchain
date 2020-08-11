# Events

The project module emits the following events:

## Handlers

## MsgCreateProject

| Type           | Attribute Key | Attribute Value |
|----------------|---------------|-----------------|
| create_project | tx_hash       | {tx_hash}       |
| create_project | sender_did    | {sender_did}    |
| create_project | project_did   | {project_did}   |
| create_project | pub_key       | {pub_key}       |

## MsgUpdateProjectStatus

| Type                  | Attribute Key   | Attribute Value      |
|-----------------------|-----------------|----------------------|
| update_project_status | TxHash          | {token}              |
| update_project_status | SenderDid       | {orderType}          |
| update_project_status | ProjectDid      | {address}            |
| update_project_status | EthFundingTxnID | {eth_funding_txn_id} |
| update_project_status | Status          | {status}             |

## MsgCreateEvaluation

| Type              | Attribute Key | Attribute Value |
|-------------------|---------------|-----------------|
| create_evaluation | TxHash        | {token}         |
| create_evaluation | SenderDid     | {orderType}     |
| create_evaluation | ProjectDid    | {address}       |
| create_evaluation | ClaimID       | {claim_id}      |
| create_evaluation | Status        | {status}        |

## MsgWithdrawFund

| Type           | Attribute Key | Attribute Value |
|----------------|---------------|-----------------|
| withdraw_funds | RecipientDid  | {recipient_did} |
| withdraw_funds | SenderDid     | {orderType}     |
| withdraw_funds | ProjectDid    | {address}       |
| withdraw_funds | Amount        | {amount}        |
| withdraw_funds | IsRefund      | {is_refund}     |
