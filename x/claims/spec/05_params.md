# Parameters

The claims module contains the following parameter:

| **Key**              | **Type**                                 | **Description**                                                                                                                           |
| :------------------- | :--------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------- |
| CollectionSequence   | `uint64`                                 | An onchain sequence to generate collection ids                                                                                            |
| IxoAccount           | `string`                                 | The Ixo network account address that will receive it's `NetworkFeePercentage` for every claim evaluation                                  |
| NetworkFeePercentage | `github_com_cosmos_cosmos_sdk_types.Dec` | The percentage that the `IxoAccount` param account will receive for every claim evaluation                                                |
| NodeFeePercentage    | `github_com_cosmos_cosmos_sdk_types.Dec` | The percentage that the node relayer for the given marketplace that the collection was created in will receive for every claim evaluation |
