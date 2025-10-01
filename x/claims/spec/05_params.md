# Parameters

The claims module contains the following parameter:

| Key                    | Type   | Example                      |
| ---------------------- | ------ | ---------------------------- |
| collection_sequence    | uint64 | 0                            |
| ixo_account            | string | "ixo1...address"             |
| network_fee_percentage | Dec    | "0.100000000000000000" (10%) |
| node_fee_percentage    | Dec    | "0.100000000000000000" (10%) |
| intent_sequence        | uint64 | 0                            |

## CollectionSequence

The collection_sequence parameter is a counter used to generate unique identifiers for new claim collections. The sequence is incremented each time a new collection is created.

## IxoAccount

The ixo_account parameter contains the account address of the network DAO operated by the Impacts Venture Cooperative that receives a portion of Oracle payment fees.

## NetworkFeePercentage

The network_fee_percentage parameter defines the percentage of Oracle payments that go to the network DAO (identified by ixo_account). This is expressed as a decimal value, where "0.100000000000000000" represents 10%.

## NodeFeePercentage

The node_fee_percentage parameter defines the percentage of Oracle payments that go to the Market Relayer that connects Oracle services to users. This is expressed as a decimal value, where "0.100000000000000000" represents 10%.

## IntentSequence

The intent_sequence parameter is a counter used to generate unique identifiers for new claim intents. The sequence is incremented each time a new intent is created.
