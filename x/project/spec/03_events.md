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

| Type                  | Attribute Key      | Attribute Value      |
|-----------------------|--------------------|----------------------|
| update_project_status | tx_hash            | {tx_hash}            |
| update_project_status | sender_did         | {sender_did}         |
| update_project_status | project_did        | {project_did}        |
| update_project_status | eth_funding_txn_id | {eth_funding_txn_id} |
| update_project_status | updated_status     | {status}             |

## MsgUpdateProjectDoc

| Type               | Attribute Key      | Attribute Value      |
|--------------------|--------------------|----------------------|
| update_project_doc | tx_hash            | {tx_hash}            |
| update_project_doc | sender_did         | {sender_did}         |
| update_project_doc | project_did        | {project_did}        |

## MsgCreateAgent

| Type         | Attribute Key | Attribute Value |
|--------------|---------------|-----------------|
| create_agent | tx_hash       | {tx_hash}       |
| create_agent | sender_did    | {sender_did}    |
| create_agent | project_did   | {project_did}   |
| create_agent | agent_did     | {agent_did}     |
| create_agent | role          | {role}          |

## MsgUpdateAgent

| Type         | Attribute Key  | Attribute Value |
|--------------|----------------|-----------------|
| update_agent | tx_hash        | {tx_hash}       |
| update_agent | sender_did     | {sender_did}    |
| update_agent | project_did    | {project_did}   |
| update_agent | agent_did      | {agent_did}     |
| update_agent | role           | {role}          |
| update_agent | updated_status | {status}        |

## MsgCreateClaim

| Type         | Attribute Key     | Attribute Value     |
|--------------|-------------------|---------------------|
| create_claim | tx_hash           | {tx_hash}           |
| create_claim | sender_did        | {sender_did}        |
| create_claim | project_did       | {project_did}       |
| create_claim | claim_id          | {claim_id}          |
| create_claim | claim_template_id | {claim_template_id} |

## MsgCreateEvaluation

| Type              | Attribute Key | Attribute Value |
|-------------------|---------------|-----------------|
| create_evaluation | tx_hash       | {tx_hash}       |
| create_evaluation | sender_did    | {sender_did}    |
| create_evaluation | project_did   | {project_did}   |
| create_evaluation | claim_id      | {claim_id}      |
| create_evaluation | claim_status  | {status}        |

## MsgWithdrawFund

| Type           | Attribute Key | Attribute Value |
|----------------|---------------|-----------------|
| withdraw_funds | sender_did    | {sender_did}    |
| withdraw_funds | project_did   | {project_did}   |
| withdraw_funds | recipient_did | {recipient_did} |
| withdraw_funds | amount        | {amount}        |
| withdraw_funds | is_refund     | {is_refund}     |
