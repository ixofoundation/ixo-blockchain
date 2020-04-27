#!/usr/bin/env bash

PASSWORD="12345678"

ixod init local --chain-id pandora-1

yes $PASSWORD | ixocli keys delete miguel --force
yes $PASSWORD | ixocli keys delete francesco --force
yes $PASSWORD | ixocli keys delete shaun --force
yes $PASSWORD | ixocli keys delete fee --force
yes $PASSWORD | ixocli keys delete fee2 --force
yes $PASSWORD | ixocli keys delete fee3 --force
yes $PASSWORD | ixocli keys delete fee4 --force
yes $PASSWORD | ixocli keys delete fee5 --force

yes $PASSWORD | ixocli keys add miguel
yes $PASSWORD | ixocli keys add francesco
yes $PASSWORD | ixocli keys add shaun
yes $PASSWORD | ixocli keys add fee
yes $PASSWORD | ixocli keys add fee2
yes $PASSWORD | ixocli keys add fee3
yes $PASSWORD | ixocli keys add fee4
yes $PASSWORD | ixocli keys add fee5

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show miguel -a)" 100000000stake,1000000res,1000000rez,100000000000ixo
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show francesco -a)" 100000000stake,1000000res,1000000rez,100000000000ixo
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show shaun -a)" 100000000stake,1000000res,1000000rez,100000000000ixo

# Add DID-based genesis accounts
MIGUEL_ADDR="ixo1x2x0thq6x2rx7txl0ujpyg9rr0c8mc8ad904xw"    # address from did:ixo:4XJLBfGtWSGKSz4BeRxdun
FRANCESCO_ADDR="ixo1nnxvyr6hs0sglppqzw4v5s9r5dwh89423xu5zp" # address from did:ixo:UKzkhVSHc3qEFva5EY2XHt
SHAUN_ADDR="ixo1vnsjples23f6ggtsvu5s24vv86ue0hl2hvnnsd"     # address from did:ixo:U4tSpzzv91HHqWW1YmFkHJ
yes $PASSWORD | ixod add-genesis-account "$MIGUEL_ADDR" 100000000stake,1000000res,1000000rez,100000000000ixo
yes $PASSWORD | ixod add-genesis-account "$FRANCESCO_ADDR" 100000000stake,1000000res,1000000rez,100000000000ixo
yes $PASSWORD | ixod add-genesis-account "$SHAUN_ADDR" 100000000stake,1000000res,1000000rez,100000000000ixo

# Add genesis oracles
MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
yes $PASSWORD | ixod add-genesis-oracle "$MIGUEL_DID" "ixo:mint"
yes $PASSWORD | ixod add-genesis-oracle "$FRANCESCO_DID" "ixo:mint/burn/transfer"
yes $PASSWORD | ixod add-genesis-oracle "$SHAUN_DID" "res:transfer,rez:transfer"

ixocli config chain-id pandora-1
ixocli config output json
ixocli config indent true
ixocli config trust-node true

yes $PASSWORD | ixod gentx --name miguel

ixod collect-gentxs
ixod validate-genesis

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
