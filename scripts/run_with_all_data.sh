#!/usr/bin/env bash

PASSWORD="12345678"

ixod init local --chain-id pandora-1

yes 'y' | ixod keys delete miguel --force
yes 'y' | ixod keys delete francesco --force
yes 'y' | ixod keys delete shaun --force
yes 'y' | ixod keys delete fee --force
yes 'y' | ixod keys delete fee2 --force
yes 'y' | ixod keys delete fee3 --force
yes 'y' | ixod keys delete fee4 --force
yes 'y' | ixod keys delete fee5 --force

yes $PASSWORD | ixod keys add miguel
yes $PASSWORD | ixod keys add francesco
yes $PASSWORD | ixod keys add shaun
yes $PASSWORD | ixod keys add fee
yes $PASSWORD | ixod keys add fee2
yes $PASSWORD | ixod keys add fee3
yes $PASSWORD | ixod keys add fee4
yes $PASSWORD | ixod keys add fee5

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | ixod add-genesis-account "$(ixod keys show miguel -a)" 100000000000uixo,100000000000res,100000000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixod keys show francesco -a)" 100000000000uixo,100000000000res,100000000000rez
yes $PASSWORD | ixod add-genesis-account "$(ixod keys show shaun -a)" 100000000000uixo,100000000000res,100000000000rez

# Add pubkey-based genesis accounts
MIGUEL_ADDR="ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"    # address from did:ixo:4XJLBfGtWSGKSz4BeRxdun's pubkey
FRANCESCO_ADDR="ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y" # address from did:ixo:UKzkhVSHc3qEFva5EY2XHt's pubkey
SHAUN_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"     # address from did:ixo:U4tSpzzv91HHqWW1YmFkHJ's pubkey
yes $PASSWORD | ixod add-genesis-account "$MIGUEL_ADDR" 100000000000uixo,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$FRANCESCO_ADDR" 100000000000uixo,1000000res,1000000rez
yes $PASSWORD | ixod add-genesis-account "$SHAUN_ADDR" 100000000000uixo,1000000res,1000000rez

# Add genesis oracles
MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
SHAUN_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
yes $PASSWORD | ixod add-genesis-oracle "$MIGUEL_DID" "uixo:mint"
yes $PASSWORD | ixod add-genesis-oracle "$FRANCESCO_DID" "uixo:mint/burn/transfer,uxgbp:mint/burn/transfer"
yes $PASSWORD | ixod add-genesis-oracle "$SHAUN_DID" "res:transfer,rez:transfer"

# Add ixo did
IXO_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
FROM="\"ixo_did\": \"\""
TO="\"ixo_did\": \"$IXO_DID\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
STAKING_TOKEN="uixo"
FROM="\"bond_denom\": \"stake\""
TO="\"bond_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json
FROM="\"mint_denom\": \"stake\""
TO="\"mint_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
FEE_TOKEN="uixo"
FROM="\"stake\""
TO="\"$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
ixod config chain-id pandora-1
ixod config output json
ixod config indent true
ixod config trust-node true

#ixod gentx miguel --amount 1000000uixo --chain-id pandora-1
ixod gentx miguel 1000000uixo --chain-id pandora-1

ixod collect-gentxs
ixod validate-genesis

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

# Uncomment the below to broadcast REST endpoint
# Do not forget to comment the bottom lines !!
#ixod start --pruning "syncable" &
#ixod rest-server --chain-id pandora-1 --laddr="tcp://0.0.0.0:1317" --trust-node && fg

ixod start --pruning "everything" &
ixod rest-server --chain-id pandora-1 --trust-node && fg
