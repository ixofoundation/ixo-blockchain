#!/usr/bin/env bash

# For development purposes this script assumes you already ran run_with_all_data for app.toml changes and 1s block time

echo "Exporting app state to genesis file..."
ixod export >genesis.json

echo "Backing up existing genesis file..."
cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup

echo "Moving new genesis file to $HOME/.ixod/config/genesis.json..."
mv genesis.json "$HOME"/.ixod/config/genesis.json

ixod tendermint unsafe-reset-all
ixod validate-genesis

ixod start --pruning "nothing"
