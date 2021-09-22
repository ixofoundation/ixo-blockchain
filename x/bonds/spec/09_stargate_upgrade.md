# Stargate Upgrade

The bonds module has been upgraded to be compatible with the Cosmos SDK v0.40 (a.k.a. Stargate) release. The main 
changes are listed below.

## Protocol buffers (Protobuf)
One of the most significant improvements introduced in Stargate is the use of protocol buffers as the new standard 
serialization format for blockchain state & wire communication. Protobuf definitions are organized into packages in the 
new `proto/` directory. The three categories of types that must be converted to Protobuf messages are:
* client-facing types: Msgs, query requests and responses. This is because the client will send these types over the 
  wire to the app.
* objects that are stored in state. This is because the SDK stores the binary representation of these types in state.
* genesis types. These are used when importing and exporting state snapshots during chain upgrades.

Once all Protobuf messages are defined in the `proto/` directory, we can generate the `*.pb.go` files by running `sudo 
make proto-gen`. This in turn uses the `scripts/protocgen.sh` script.

Existing Amino REST endpoints are all preserved, although they are planned to be deprecated in a future release. New 
routes have been added via gRPC-gateway. gRPC-gateway exposes gRPC endpoints as REST endpoints. For each RPC endpoint 
defined in a Protobuf service, the SDK offers a REST equivalent, which is defined as an option. For example, in
`query.proto` we have the following.
```
rpc Bond(QueryBondRequest) returns (QueryBondResponse) {
    option (google.api.http).get = "/ixo/bonds/{bond_did}";
}
```
To query the bond `did:ixo:U7GK8p8rVhJMKhBVRCJJ8c` using the legacy Amino REST endpoint, use 
`/bonds/did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`. To query the bond using the new gRPC-gateway REST 
endpoint, use `/ixo/bonds/did:ixo:U7GK8p8rVhJMKhBVRCJJ8c`.

## Single application binary
ixo-blockchain now compiles to a single application binary, as opposed to separate binaries for running a node and one 
for the CLI & REST server. In practice, this means we no longer have an `ixocli` command and now only use `ixod`. 

Note: There is currently no way of configuring the `ixod` command, which means we have to add flags such as the chain 
ID every time we use `ixod`. This has been reported and (at the time of writing) is an open issue, available here: 
https://github.com/cosmos/cosmos-sdk/issues/8529. 
