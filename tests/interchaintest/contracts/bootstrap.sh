#!/usr/bin/env bash
# Verify the bundled wasm artefacts are intact. The .wasm files in this
# directory are committed to the repo (sourced from
# ixo-multiclient-sdk/assets/contracts) so interchaintest runs offline
# without external downloads.
#
# If you need to bump a contract: replace the file, regenerate
# `checksums.txt` with `shasum -a 256 *.wasm > checksums.txt`, and commit.
set -euo pipefail

cd "$(dirname "$0")"

if [[ ! -f checksums.txt ]]; then
    echo "ERROR: checksums.txt missing" >&2
    exit 1
fi

if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 -c checksums.txt
elif command -v sha256sum >/dev/null 2>&1; then
    sha256sum -c checksums.txt
else
    echo "ERROR: neither shasum nor sha256sum available" >&2
    exit 1
fi

echo "Wasm artefacts verified:"
ls -lh ./*.wasm
