#!/usr/bin/env bash

# Must be run from root path inside ixo-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

ALICE_DID="did:x:zQ3shUQYdonbR4CaaWBmfYMHND7KEXY3tEStFndNGWp4njQ54"
ALICE_ADDR="ixo12am7v5xgjh72c7xujreyvtncqwue3w0v6ud3r4"
ALICE_PUBKEY58="iTPcxe7wjZMhG66FAGQKPWwnsmruwtqirhcUDUPfWvDL"

BOB_DID="did:x:zQ3shdv1QebkD12iJh2pNVsHtBZSW4prECtzwdtnj2eZKZ436"
BOB_ADDR="ixo13dy867pyn8jda82vnshy7jjjv42n69k7497jrh"
BOB_PUBKEY58="sxrPoTGjgPVRSw8x7nQqMy54R4fFvLxQhy2vjJtCLaBN"

CHARLIE_DID="did:x:zQ3shne9kEBA4Q8mQr5em3NF9mgE8zLXiHGMydLeDqPvbMo7c"
CHARLIE_ADDR="ixo1fewufqrjy0r8kercq3wazsr7v0cymhvgteq442"
CHARLIE_PUBKEY58="22gzjP2gb5VYXbyyLfHN6x5rhLaLjziKShQtRY4FU9KFt"

MIGUEL_DID="did:x:zQ3shoiydFD6jdTdXLPProPZWL6igg9bCyaJY6zEKqQoNE96C"
MIGUEL_ADDR="ixo1n8yrmeatsk74dw0zs95ess9sgzptd6thgjgcj2"
MIGUEL_PUBKEY58="23mpcQ4dGJpQe6HiSRJgTWWMF2PQEh2G1B4UXY58F1fEU"

FRANCESCO_DID="did:x:zQ3shMs8BxDXczP2i5KLybDTQHtwQ5Q5o6aKEgX8QF98TfkU4"
FRANCESCO_ADDR="ixo19hatudm363ufmyrrykxdmrf8f0884yq75expse"
FRANCESCO_PUBKEY58="buyB7549fjopqDfZD8aMUJZxRdtpp2GhkbNbwoTLTGcL"

SHAUN_DID="did:x:zQ3shQ79urtEmSDxdi7aQ5tGUkXFqdkQsXxsi3cGRo3mH3hPY"
SHAUN_ADDR="ixo1rngxtm5sapzqdtw3k3e2e9zkjxzgpxd6vw9pye"
SHAUN_PUBKEY58="e9zu1jmJ7ajkU1tyhoPRvvtPyzDuFQqB7gWdVi69qDXp"

# Ledger DIDs
echo "Ledgering DID for Alice"
DID=$(full_iid_doc $ALICE_DID $ALICE_ADDR $ALICE_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from alice

echo "Ledgering DID for Bob"
DID=$(full_iid_doc $BOB_DID $BOB_ADDR $BOB_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from bob

echo "Ledgering DID for Charlie"
DID=$(full_iid_doc $CHARLIE_DID $CHARLIE_ADDR $CHARLIE_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from charlie

echo "Ledgering DID for Miguel"
DID=$(full_iid_doc $MIGUEL_DID $MIGUEL_ADDR $MIGUEL_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from miguel

echo "Ledgering DID for Francesco"
DID=$(full_iid_doc $FRANCESCO_DID $FRANCESCO_ADDR $FRANCESCO_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from francesco

echo "Ledgering DID for Shaun"
DID=$(full_iid_doc $SHAUN_DID $SHAUN_ADDR $SHAUN_PUBKEY58)
ixod_tx iid create-iid "$(echo $DID | jq -rc .)" --from shaun

# Verifications
# echo "Add Verification for Alice"
# ixod_tx iid add-verification-method $ALICE_DID "$(echo $DID | jq -rc .)" --from alice
