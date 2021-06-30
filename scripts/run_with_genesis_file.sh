#!/usr/bin/env bash

echo "Backing up existing genesis file..."
cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup

echo "Copying new genesis file to $HOME/.ixod/config/genesis.json..."
cp genesis.json "$HOME"/.ixod/config/genesis.json

ixod unsafe-reset-all
ixod validate-genesis

# Enable REST API (assumed to be at line 104 of app.toml)
FROM="enable = false"
TO="enable = true"
sed -i "104s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Enable Swagger docs (assumed to be at line 107 of app.toml)
FROM="swagger = false"
TO="swagger = true"
sed -i "107s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.ixod/config/config.toml
#sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.ixod/config/config.toml

ixod start --pruning "nothing"
