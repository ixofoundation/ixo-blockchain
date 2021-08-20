# Entity Accounts

The project module has a number of entity/project accounts created per
entity/project:

- `InitiatingNodePayFees`: when an oracle service is delivered, \[by default\] 10% of
  the oracle fee payment is sent to this address.
- `IxoPayFees`: when an oracle service is delivered, \[by default\] 80% of the
  oracle fee payment is sent to this address.
- `IxoFees`: at project payout, any amount in the `IxoPayFees` account is moved to this
  account address. In turn, any amount now in IxoFees is sent to pay the network fee to ixo (the account associated with a DID which is
  specified in the genesis file)
- `<projectDid>`: this is the project's funding account, from which tokens
  are paid to the oracle service.
- Additionally, one entity/project account for each agent is also created. Any
  payments intended for the particular agent is sent to their corresponding
  entity account. At project payout, the agent's tokens can be withdrawn.

The fee defaults can be configured from the Genesis file.

Example accounts for a project with DID `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c` and one
agent with DID `did:ixo:RYLHkfNpbA8Losy68jt4yF`:

```json
{
  "InitiatingNodePayFees": "ixo1xvjy68xrrtxnypwev9r8tmjys9wk0zkkspzjmq",
  "IxoPayFees": "ixo1udgxtf6yd09mwnnd0ljpmeq4vnyhxdg03uvne3",
  "IxoFees": "ixo1ff9we62w6eyes7wscjup3p40vy4uz0sa7j0ajc",
  "did:ixo:U7GK8p8rVhJMKhBVRCJJ8c": "ixo1rmkak6t606wczsps9ytpga3z4nre4z3nwc04p8",
  "did:ixo:RYLHkfNpbA8Losy68jt4yF": "ixo18nmp3w2xwz0rzkh8sdwkyz4fzjegemtx9vw3ky"
}
```
