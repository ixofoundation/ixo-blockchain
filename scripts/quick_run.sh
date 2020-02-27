#!/usr/bin/env bash

PASSWORD="12345678"

ixod init local --chain-id pandora-test-1

yes $PASSWORD | ixocli keys delete miguel --force
yes $PASSWORD | ixocli keys delete francesco --force
yes $PASSWORD | ixocli keys delete shaun --force
yes $PASSWORD | ixocli keys delete fee --force

yes $PASSWORD | ixocli keys add miguel
yes $PASSWORD | ixocli keys add francesco
yes $PASSWORD | ixocli keys add shaun
yes $PASSWORD | ixocli keys add fee

yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show miguel -a)" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show francesco -a)" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show shaun -a)" 100000000stake,1000000res,1000000rez

ixocli config chain-id bondschain-1
ixocli config output json
ixocli config indent true
ixocli config trust-node true

yes $PASSWORD | ixod gentx --name miguel

ixod collect-gentxs
ixod validate-genesis

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-test-1 --trust-node && fg
