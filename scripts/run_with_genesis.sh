#!/usr/bin/env bash

echo "Exporting app state to genesis file..."
ixod export >genesis.json

echo "Backing up existing genesis file..."
cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup

echo "Moving new genesis file to $HOME/.ixod/config/genesis.json..."
mv genesis.json "$HOME"/.ixod/config/genesis.json

ixod unsafe-reset-all
ixod validate-genesis

# Enable REST API > app.toml
FROM="enable = false"
TO="enable = true"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Enable Swagger docs > app.toml
FROM="swagger = false"
TO="swagger = true"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Broadcast node RPC endpoint
FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

# Set timeouts to 1s for shorter block times
#sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.ixod/config/config.toml
#sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.ixod/config/config.toml

ixod start --pruning "nothing"
