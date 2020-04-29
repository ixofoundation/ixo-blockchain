#!/usr/bin/env bash

ixod start --pruning "syncable" &
ixocli rest-server --chain-id pandora-1 --trust-node && fg
