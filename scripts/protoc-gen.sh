#!/usr/bin/env bash

set -eo pipefail

echo "Generating gogo proto code"
cd proto

proto_dirs=$(find ./ixo -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      echo "Generating gogo proto code for $file"
      buf generate $file --template buf.gen.gogo.yaml
    fi
  done
done

cd ..

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
cp -r github.com/ixofoundation/ixo-blockchain/v5/* ./
rm -rf github.com

# go mod tidy

# Uncomment once ORM/Pulsar support is needed.
# Ref: https://github.com/osmosis-labs/osmosis/pull/1589
