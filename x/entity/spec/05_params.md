# Parameters

The entity module contains the following parameter:

| **Key**            | **Type** | **Description**                                                                          |
| :----------------- | :------- | :--------------------------------------------------------------------------------------- |
| NftContractAddress | `string` | The cw721 smart contract address that is used to mint every entity as an nft on creation |
| NftContractMinter  | `string` | The address of the `NftContractAddress` minter                                           |
| CreateSequence     | `uint64` | An onchain sequence that is used to generate the entity id so it can be deterministic    |
