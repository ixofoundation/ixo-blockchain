#!/usr/bin/env bash
PASSWORD=“12345678”
# ixod init local --chain-id devnet-1
/app/ixod init local --chain-id devnet-1
# yes ‘y’ | ixod keys delete miguel --force
# yes ‘y’ | ixod keys delete francesco --force
# yes ‘y’ | ixod keys delete shaun --force
# yes ‘y’ | ixod keys delete fee --force
# yes ‘y’ | ixod keys delete fee2 --force
# yes ‘y’ | ixod keys delete fee3 --force
# yes ‘y’ | ixod keys delete fee4 --force
# yes ‘y’ | ixod keys delete fee5 --force
# yes ‘y’ | ixod keys delete reserveOut --force
# yes $PASSWORD | ixod keys add miguel
# yes $PASSWORD | ixod keys add francesco
# yes $PASSWORD | ixod keys add shaun
# yes $PASSWORD | ixod keys add fee
# yes $PASSWORD | ixod keys add fee2
# yes $PASSWORD | ixod keys add fee3
# yes $PASSWORD | ixod keys add fee4
# yes $PASSWORD | ixod keys add fee5
# yes $PASSWORD | ixod keys add reserveOut
yes 12345678 | /app/ixod keys add dev-main
# Note: important to add ‘miguel’ as a genesis-account since this is the chain’s validator
yes 12345678 | /app/ixod add-genesis-account $(/app/ixod keys show dev-main -a) 100000000000000000uixo

# yes $PASSWORD | ixod add-genesis-account “$(ixod keys show miguel -a)” 1000000000000uixo,1000000000000res,1000000000000rez,1000000000000uxgbp
# yes $PASSWORD | ixod add-genesis-account “$(ixod keys show francesco -a)” 1000000000000uixo,1000000000000res,1000000000000rez
# yes $PASSWORD | ixod add-genesis-account “$(ixod keys show shaun -a)” 1000000000000uixo,1000000000000res,1000000000000rez
# Add pubkey-based genesis accounts
# MIGUEL_ADDR=“ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx”    # address from did:ixo:4XJLBfGtWSGKSz4BeRxdun’s pubkey
# FRANCESCO_ADDR=“ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y” # address from did:ixo:UKzkhVSHc3qEFva5EY2XHt’s pubkey
# SHAUN_ADDR=“ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un”     # address from did:ixo:U4tSpzzv91HHqWW1YmFkHJ’s pubkey
# yes $PASSWORD | ixod add-genesis-account “$MIGUEL_ADDR” 1000000000000uixo,1000000000000res,1000000000000rez
# yes $PASSWORD | ixod add-genesis-account “$FRANCESCO_ADDR” 1000000000000uixo,1000000000000res,1000000000000rez
# yes $PASSWORD | ixod add-genesis-account “$SHAUN_ADDR” 1000000000000uixo,1000000000000res,1000000000000rez
# yes $PASSWORD | ixod add-genesis-account “ixo19h3lqj50uhzdrv8mkafnp55nqmz4ghc2sd3m48” 1000000000000uixo,1000000000000res,1000000000000rez
# Add ixo did
yes 12345678 | /app/ixod gentx dev-main 10000000000uixo --chain-id devnet-1
HOME=/root
/app/ixod collect-gentxs

# IXO_DID=“did:ixo:U4tSpzzv91HHqWW1YmFkHJ”
# export FROM=\”ixo_did\“: \“\”
# export TO=“\”ixo_did\“: \“$IXO_DID\“”
sed -i 's/"ixo_did" : ""/"ixo_did" : ""/;s/"bond_denom" : "stake"/"bond_denom" : "uixo"/;s/"mint_denom" : "stake"/"mint_denom" : "uixo"/;s/stake/uixo/;s/"Reserved_bond_tokens" : "\[\]"/"Reserved_bond_tokens" : "\[\]"/;s/"minimum-gas-prices" : ""/"minimum-gas-prices" : "0.025uixo"/;s/"enable" : "false"/"enable" : "true"/;s/"swagger" : "false"/"swagger" : "true"/;' $HOME/.ixod/config/genesis.json
MAX_VOTING_PERIOD="30s" # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.ixod/config/genesis.json
# sed -i '/ixo_did/c\   \"i xo_did\" : \"did:ixo:aaaaaaaaaaa\",' $HOME/.ixod/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
# export STAKING_TOKEN=“uixo”
# export FROM=“\”bond_denom\“: \“stake\“”
# export TO=“\”bond_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/bond_denom/c\   \"bond_denom\" : \"uixo\",' $HOME/.ixod/config/genesis.json

# export FROM=“\”mint_denom\“: \“stake\“”
# export TO=“\”mint_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/mint_denom/c\   \"mint_denom\" : \"uixo\",' $HOME/.ixod/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
# export FEE_TOKEN=“uixo”
# export FROM=“\”stake\“”
# export TO=“\”$FEE_TOKEN\“”
# sed 's/stake/uixo' $HOME/.ixod/config/genesis.json

# Set reserved bond tokens
# export RESERVED_BOND_TOKENS=“”  # example: ” \“abc\“, \“def\“, \“ghi\” ”
# export FROM=“\”reserved_bond_tokens\“: \[\]”
# export TO=“\”reserved_bond_tokens\“: \[$RESERVED_BOND_TOKENS\]”
# sed -i '/reserved_bond_tokens/c\   \"Reserved_bond_tokens\" : \"[]\",' $HOME/.ixod/config/genesis.json

# Set min-gas-prices (using fee token)
# export FROM=“minimum-gas-prices = \“\”
# export TO=“minimum-gas-prices = \“0.025$FEE_TOKEN\“”
# sed -i '/minimum-gas-prices/c\   \"minimum-gas-prices\" : \"0.025uixo\",' $HOME/.ixod/config/genesis.json
#ixod config chain-id devnet-1
#ixod config output jsonW
#ixod config indent true
#ixod config trust-node true
# ixod gentx miguel 1000000uixo --chain-id devnet-1
/app/ixod validate-genesis
# Enable REST API (assumed to be at line 104 of app.toml)
# export FROM=“enable = false”
# TO=“enable = true”
# sed -i '/enable/c\   enable = true' $HOME/.ixod/config/genesis.json
# Enable Swagger docs (assumed to be at line 107 of app.toml)
# export FROM=“swagger = false”
# export TO=“swagger = true”
# sed -i '/swagger/c\   swagger = true' $HOME/.ixod/config/genesis.json
# Uncomment the below to broadcast node RPC endpoint
#FROM=“laddr = \“tcp:\/\/127.0.0.1:26657\“”
#TO=“laddr = \“tcp:\/\/0.0.0.0:26657\“”
#sed -i “s/$FROM/$TO/” “$HOME”/.ixod/config/config.toml
# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i ‘s/timeout_commit = “5s”/timeout_commit = “1s”/g’ “$HOME”/.ixod/config/config.toml
#sed -i ‘s/timeout_propose = “3s”/timeout_propose = “1s”/g’ “$HOME”/.ixod/config/config.toml
# ixod start --pruning “nothing”
