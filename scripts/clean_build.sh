#!/usr/bin/env bash

rm -rf "$HOME"/.ixod
rm -rf "$HOME"/.ixocli  # TODO: remove me since there's no longer any ixocli

make install # assumes currently in project directory
