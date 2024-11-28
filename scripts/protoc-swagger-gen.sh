#!/bin/sh
#
# This script is intended to be run inside the osmolabs/osmo-proto-gen:v0.8
# docker container: https://hub.docker.com/r/osmolabs/osmo-proto-gen

set -eo pipefail

# The directory where the final output are to be stored
docs_dir="./lib/docs"

# The directory where temporary swagger files are to be stored before they are
# combined. Will be deleted in the end
tmp_dir="./tmp-swagger-gen"
if [ -d $tmp_dir ]; then
  rm -rf $tmp_dir
fi
mkdir -p $tmp_dir

# Third-party proto dependencies
deps="github.com/cosmos/cosmos-sdk@v0.50.8"
deps="$deps github.com/cosmos/ibc-go/v8@v8.3.2"
deps="$deps github.com/CosmWasm/wasmd@v0.50.0"

# Download dependencies in go.mod
# Necessary for the `go list` commands in the next step to work
echo "Downloading dependencies..."
for dep in $deps; do
  echo $dep
  go mod download $dep
done

# Directories that contain protobuf files that are to be transpiled into swagger
# These include Mars modules and third party modules and services
dirs="./proto"
for dep in $deps; do
  dep_dir=$(go list -f '{{ .Dir }}' -m $dep)
  dirs="$dirs ${dep_dir}/proto"
done
proto_dirs=$(find $dirs -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

# Generate swagger files for `query.proto` and `service.proto`
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \)); do
    echo $file
    buf generate --template proto/buf.gen.swagger.yaml $file
  done
done

# Combine swagger files
# Uses nodejs package `swagger-combine`.
# All the individual swagger files need to be configured in `config.json` for merging
echo "Combining swagger files..."
swagger-combine ${docs_dir}/config.json \
  -o ${docs_dir}/swagger-ui/swagger.yaml \
  -f yaml \
  --continueOnConflictingPaths true \
  --includeDefinitions true

# # Clean swagger files
rm -rf $tmp_dir
