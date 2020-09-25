#!/usr/bin/env bash

rm -rf "$HOME"/.ixod
rm -rf "$HOME"/.ixocli

make install # assumes currently in project directory
