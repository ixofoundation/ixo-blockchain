#!/usr/bin/env bash

# Must be run from root path inside ixo-blockchain for source to work
source ./scripts/constants.sh

ixod init local --chain-id $CHAIN_ID

# When incorporating another genesis file
# cp "$HOME"/.ixod/config/genesis.json "$HOME"/.ixod/config/genesis.json.backup #Backing up
# cp genesis.json "$HOME"/.ixod/config/genesis.json #Copy over new genesis in root

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
  # add as genesis-account with fees
  yes $PASSWORD | ixod genesis add-genesis-account $(ixod keys show ${USERS[i]} -a) 1000000000000uixo,1000000000000res,1000000000000rez,1000000000000uxgbp
done

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

# Set reserved bond tokens
RESERVED_BOND_TOKENS="" # example: " \"abc\", \"def\", \"ghi\" "
FROM="\"reserved_bond_tokens\": \[\]"
TO="\"reserved_bond_tokens\": \[$RESERVED_BOND_TOKENS\]"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

# Set bank denom metadata
DENOM_METADATA="\{\"base\":\"uixo\",\"denom_units\":\[\{\"aliases\":\[\"microixo\"\],\"denom\":\"uixo\",\"exponent\":0\},\{\"aliases\":\[\"milliixo\"\],\"denom\":\"mixo\",\"exponent\":3\},\{\"aliases\":\[\],\"denom\":\"ixo\",\"exponent\":6\}\],\"description\":\"The native staking token of ixo.\",\"display\":\"ixo\",\"name\":\"ixo\",\"symbol\":\"ixo\"\}"
FROM="\"denom_metadata\": \[\]"
TO="\"denom_metadata\": \[$DENOM_METADATA\]"
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/app.toml

# Make voting period short
MAX_VOTING_PERIOD="90s" # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json
EXPEDITED_VOTING_PERIOD="60s" # example: "172800s"
FROM="\"expedited_voting_period\": \"86400s\""
TO="\"expedited_voting_period\": \"$EXPEDITED_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json

yes $PASSWORD | ixod genesis gentx miguel 1000000uixo --chain-id $CHAIN_ID
ixod genesis collect-gentxs
ixod genesis validate-genesis

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
ixod start --pruning "nothing"
