#!/usr/bin/env bash

# This is to spin up another node for development, run this to get everything ready, run this then copy over the genesis file

# Must be run from root path inside ixo-blockchain for source to work
source ./scripts/constants.sh

ixod init local2 --chain-id $CHAIN_ID

# first add and remove a dummy user so PASSWORD can be set in keychain
yes $PASSWORD | ixod keys add dummy &>/dev/null
yes $PASSWORD | ixod keys delete ${USERS[i]} -y &>/dev/null

for ((i = 0; i < ${#USERS[@]}; ++i)); do
  # delete key if exists
  yes $PASSWORD | ixod keys delete ${USERS[i]} -y 2>/dev/null
  # create key with constant mnemonic in /scripts/constants.sh
  (
    echo ${MNEMONICS[i]}
    echo $PASSWORD
  ) | ixod keys add ${USERS[i]} --recover
done

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Enable REST API, RPC, and gRPC
FROM="enable = false"
TO="enable = true"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml
FROM="address = \"tcp:\/\/localhost:1317\""
TO="address = \"tcp:\/\/0.0.0.0:1317\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml
FROM="address = \"localhost:9090\""
TO="address = \"0.0.0.0:9090\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Enable cors
FROM="enabled-unsafe-cors = false"
TO="enabled-unsafe-cors = true"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml
FROM="cors_allowed_origins = \[\]"
TO="cors_allowed_origins = \[\"*\"\]"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

# Enable Swagger docs
FROM="swagger = false"
TO="swagger = true"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Broadcast node RPC endpoint
FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/config.toml

# Set timeouts to 1s for shorter block times
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.ixod/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.ixod/config/config.toml

# ixod start --pruning "nothing" --log_level "trace" --trace
# ixod start --pruning "nothing"
