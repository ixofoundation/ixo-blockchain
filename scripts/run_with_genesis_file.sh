#!/usr/bin/env bash

echo "Backing up existing genesis file..."
cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup

echo "Copying new genesis file to $HOME/.ixod/config/genesis.json..."
cp genesis.json "$HOME"/.ixod/config/genesis.json

ixod unsafe-reset-all
ixod validate-genesis

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

ixod start --pruning "nothing" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
