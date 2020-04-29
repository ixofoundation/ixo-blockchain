#!/usr/bin/env bash

echo "Exporting app state to genesis file..."
ixod export >genesis.json

echo "Fixing genesis file..."
sed -i 's/"genutil":null/"genutil":{"gentxs":null}/g' genesis.json
# https://github.com/cosmos/cosmos-sdk/issues/5086

echo "Backing up existing genesis file..."
cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup

echo "Moving new genesis file to $HOME/.ixod/config/genesis.json..."
mv genesis.json "$HOME"/.ixod/config/genesis.json

ixod unsafe-reset-all
ixod validate-genesis

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
