#!/usr/bin/env bash
SWAGGER_DIR=./swagger-proto

set -eo pipefail

echo "Generating proto docs"

# wait 2 seconds
sleep 2

# command to generate docs using protoc-gen-doc (Markdown)
buf protoc \
  -I "proto" \
  -I "$SWAGGER_DIR/third_party" \
  --doc_out=./lib/docs/core \
  --doc_opt=markdown,proto-docs.md \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

# command to generate docs using protoc-gen-doc (HTML)
buf protoc \
  -I "proto" \
  -I "$SWAGGER_DIR/third_party" \
  --doc_out=./lib/docs/core \
  --doc_opt=html,proto-docs.html \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

# command to generate docs using protoc-gen-doc (JSON)
buf protoc \
  -I "proto" \
  -I "$SWAGGER_DIR/third_party" \
  --doc_out=./lib/docs/core \
  --doc_opt=json,proto-docs.json \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

# clean files
rm -rf "$SWAGGER_DIR"
