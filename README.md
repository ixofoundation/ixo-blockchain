# ixo Blockchain SDK

[![version](https://img.shields.io/github/tag/ixofoundation/ixo-blockchain.svg)](https://github.com/ixofoundation/ixo-blockchain/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ixofoundation/ixo-blockchain)](https://goreportcard.com/report/github.com/ixofoundation/ixo-blockchain)
[![LoC](https://tokei.rs/b1/github/ixofoundation/ixo-blockchain)](https://github.com/ixofoundation/ixo-blockchain)

This is the official repository for the Impact Hub (ImpactHub)

## Notice
The IXO team is in the process of upgrading this repository better align with our new standards and goals. Exciting things will soon come, but not without first making some fundamental and overdue changes that might cause confusion for some.

### Mini Changelog
- Upgraded to cosmos-sdk 0.45
- Introduction of github actions to help automate some tasks. (Note this will be improved as we get more functionality in place)
- The `master` branch was renamed to `main` and will no longer represent the latest stable version but rather the next feature release. 
- As the project is still in active developmet, we thought it best to rename all are releases from version `v1.x.x` to `v0.x.x`. This would make [`v0.17.0`](https://github.com/ixofoundation/ixo-blockchain/releases/v0.17.0) our last stable version.
- Going forward all releases will follow clear samantic versioning guidelines and all stable releases will have a release branch dedicated to it. For example the release `v0.17.0` will associated the branch `release/v0.17.x` and all bugfixes for this release should be made against this branch as well as all upstream branches if deemed relevant.

---

> This document will have all details necessary to help getting started with ImpactHub

## Documentation
- Guide for setting up a Validator on the Pandora test network and Internet of Impact Hub main network: [here](https://github.com/ixofoundation/genesis)
- Swagger API documentation for ixo modules gRPC endpoints can be found at [client/docs/swagger-ui/swagger.yaml](client/docs/swagger-ui/swagger.yaml)
- Swagger API documentation for ixo modules legacy endpoints can be found at [client/docs/swagger-ui-legacy/swagger.yaml](client/docs/swagger-ui-legacy/swagger.yaml)
- Blockchain Module Specifications can be found under `x/<module>/spec`

## Building and Running

**Note**: Requires [Go 1.15+](https://golang.org/dl/)

To build and run the application:

```bash
make run
```

To build and run the application and also create some accounts:

```bash
make run_with_some_data  # Option 1
make run_with_all_data   # Option 2
```

(Optional) Once the chain has started, run one of the following:

- Add more data and activity:
```bash
bash ./scripts/add_dummy_testnet_data.sh
```

- Demos:
```bash
bash ./scripts/demo_bonds.sh     # Option 1
bash ./scripts/demo_payments.sh  # Option 2
bash ./scripts/demo_project.sh   # Option 3
...
# Look in the scripts folder for more options!
```

- To re-generate `.pb.go` and `.pb.gw.go` files from `.proto` files, as well as docs/core/proto-docs.md:
```bash
make proto-gen
```

- To re-generate API documentation (`swagger.yaml` file):
```bash
make proto-swagger-gen
```

- To build and run the application using Starport (demos will not work if the
  blockchain is started using this method, and the `./cmd/ixod` package has to
  be refactored to `./cmd/ixo-blockchaind`):

```bash
starport serve
```
