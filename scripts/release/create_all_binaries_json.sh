#!/bin/bash

tags=(
  "v2.0.0"
)

echo "# Cosmovisor binaries"

for tag in ${tags[@]}; do
  echo
  echo "## ${tag}"
  echo
  echo '```json'
  python create_binaries_json.py --tag $tag
  echo '```'
done
