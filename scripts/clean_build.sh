#!/usr/bin/env bash

rm -rf "$HOME"/.ixod
rm -rf "$HOME"/.ixocli

cd "$HOME"/go/src/github.com/ixofoundation/ixo-cosmos/ || exit
make install
