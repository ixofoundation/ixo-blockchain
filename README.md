# ixo Blockchain SDK

[![version](https://img.shields.io/github/tag/ixofoundation/ixo-blockchain.svg)](https://github.com/ixofoundation/ixo-blockchain/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ixofoundation/ixo-blockchain)](https://goreportcard.com/report/github.com/ixofoundation/ixo-blockchain)
[![LoC](https://tokei.rs/b1/github/ixofoundation/ixo-blockchain)](https://github.com/ixofoundation/ixo-blockchain)

This is the official repository for the Sustainability Hub (ixo-Hub)

> This document will have all details necessary to help getting started with ixo-Hub

## Documentation
- Guide for setting up a Relayer on the Pandora test network and Internet of Impact Hub main network: [here](https://github.com/ixofoundation/genesis)
- Swagger API documentation can be found under `client/docs/swagger-ui` and `client/docs/swagger-ui-legacy`
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

- To re-generate `.pb.go` and `.pb.gw.go` files from `.proto` files:
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
