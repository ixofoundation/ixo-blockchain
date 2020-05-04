#!/usr/bin/env bash

PASSWORD="12345678"

ixod init local --chain-id pandora-1

yes $PASSWORD | ixocli keys delete miguel --force
yes $PASSWORD | ixocli keys add miguel

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show miguel -a)" 100000000stake,1000000res,1000000rez,100000000000ixo

# Add DID-based genesis account
MIGUEL_ADDR="ixo1x2x0thq6x2rx7txl0ujpyg9rr0c8mc8ad904xw" # address from did:ixo:4XJLBfGtWSGKSz4BeRxdun
yes $PASSWORD | ixod add-genesis-account "$MIGUEL_ADDR" 100000000stake,1000000res,1000000rez,100000000000ixo

# Add genesis oracle
MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
yes $PASSWORD | ixod add-genesis-oracle "$MIGUEL_DID"

ixocli config chain-id pandora-1
ixocli config output json
ixocli config indent true
ixocli config trust-node true

yes $PASSWORD | ixod gentx --name miguel

ixod collect-gentxs
ixod validate-genesis

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
