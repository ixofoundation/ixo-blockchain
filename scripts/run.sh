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
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show miguel -a)" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show francesco -a)" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixocli keys show shaun -a)" 100000000stake,1000000res,1000000rez

# Add DID-based genesis accounts
MADDR="cosmos1yhfwr0u62dpy35hl4e6nxarysuuwrwhxmle95m"  # address from 4XJLBfGtWSGKSz4BeRxdun
FADDR="cosmos16ne06jpdr3eu2a3uv3etwxavw0q03sqscjdcj5"  # address from UKzkhVSHc3qEFva5EY2XHt
SADDR="cosmos1vc9v45u8rc946tn3j06c0glhx9q0llx2u3uuan"  # address from U4tSpzzv91HHqWW1YmFkHJ
yes $PASSWORD | ixod add-genesis-account "$MADDR" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$FADDR" 100000000stake,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$SADDR" 100000000stake,1000000res,1000000rez

ixocli config chain-id pandora-1
ixocli config output json
ixocli config indent true
ixocli config trust-node true

yes $PASSWORD | ixod gentx --name miguel

ixod collect-gentxs
ixod validate-genesis

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
