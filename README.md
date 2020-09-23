# ixo Blockchain SDK

[![version](https://img.shields.io/github/tag/ixofoundation/ixo-blockchain.svg)](https://github.com/ixofoundation/ixo-blockchain/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ixofoundation/ixo-blockchain)](https://goreportcard.com/report/github.com/ixofoundation/ixo-blockchain)
[![LoC](https://tokei.rs/b1/github/ixofoundation/ixo-blockchain)](https://github.com/ixofoundation/ixo-blockchain)

This is the official repository for the Sustainability Hub (ixo-Hub)

> This document will have all details necessary to help getting started with ixo-Hub

## Documentation
- Guide for setting up a Relayer on the Pandora Test Network: [here](https://github.com/ixofoundation/docs/blob/master/developer-tools/test-networks/join-a-test-network.md)
- Blockchain Module Specifications can be found under `x/<module>/spec`

## Scripts
Quick-start:
```bash
cd ixo-blockchain/
bash ./scripts/clean_build.sh && bash ./scripts/run_with_some_data.sh  # Option 1
bash ./scripts/clean_build.sh && bash ./scripts/run_with_all_data.sh   # Option 2
```

To run without resetting data:
```bash
cd ixo-blockchain/
bash ./scripts/run_only.sh
```

(Optional) Once the chain has started, run one of the following:

- Add more data and activity:
```bash
cd ixo-blockchain/
bash ./scripts/add_dummy_testnet_data.sh
```

- Demos:
```bash
cd ixo-blockchain/
bash ./scripts/demo_bonds.sh     # Option 1
bash ./scripts/demo_payments.sh  # Option 2
bash ./scripts/demo_project.sh   # Option 3
...
# Look in the scripts folder for more options!
```
