# Events

The treasury module emits the following events:

## Handler

## MsgSend

| Type           | Attribute Key            | Attribute Value       |
|----------------|--------------------------|-----------------------|
| send           | tx_hash_from_did         | {tx_hash_from_did}    |
| send           | to_did_or_addr           | {to_did_or_addr}      |
| send           | amount                   | {amount}              |

## MsgOracleTransfer
| Type           | Attribute Key            | Attribute Value       |
|----------------|--------------------------|-----------------------|
| oracle_transfer| oracle_did               | {oracle_did}          |
| oracle_transfer| from_did                 | {from_did}            |
| oracle_transfer| to_did_or_addr           | {to_did_or_addr}      |
| oracle_transfer| amount                   | {amount}              |
| oracle_transfer| proof                    | {proof}               |

## MsgOracleBurn
| Type           | Attribute Key            | Attribute Value       |
|----------------|--------------------------|-----------------------|
| oracle_burn    | tx_hash_from_did         | {tx_hash_from_did}    |
| oracle_burn    | oracle_did               | {oracle_did}          |
| oracle_burn    | amount                   | {amount}              |
| oracle_burn    | proof                    | {proof}               |



