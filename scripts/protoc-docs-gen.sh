#!/usr/bin/env bash

set -eo pipefail

# protoc_install_proto_gen_doc() {
#   echo "Installing protobuf protoc-gen-doc plugin"
#   (go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest 2>/dev/null)
# }
# protoc_install_proto_gen_doc

echo "Generating proto docs"

# command to generate docs using protoc-gen-doc (Markdown)
buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  --doc_out=./docs/core \
  --doc_opt=markdown,proto-docs.md \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

# command to generate docs using protoc-gen-doc (HTML)
buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  --doc_out=./docs/core \
  --doc_opt=html,proto-docs.html \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

# command to generate docs using protoc-gen-doc (JSON)
buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  --doc_out=./docs/core \
  --doc_opt=json,proto-docs.json \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')
