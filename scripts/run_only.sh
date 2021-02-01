#!/usr/bin/env bash

# Uncomment the below to broadcast REST endpoint
# Do not forget to comment the bottom lines !!
#ixod start --pruning "everything" &
#ixocli rest-server --chain-id pandora-1 --laddr="tcp://0.0.0.0:1317" --trust-node && fg

ixod start --pruning "everything" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
