"""
Usage:
This script generates a JSON object containing binary download URLs and their corresponding checksums 
for a given release tag of ixofoundation/ixo-blockchain or from a provided checksum URL.
The binary JSON is compatible with cosmovisor and with the chain registry.

You can run this script with the following commands:

❯ python create_binaries_json.py --checksums_url https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/sha256sum.txt

Output:
{
    "binaries": {
    "linux/arm64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/2.0.0/ixod-2.0.0-linux-arm64?checksum=<checksum>",
    "darwin/arm64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/2.0.0/ixod-2.0.0-darwin-arm64?checksum=<checksum>",
    "darwin/amd64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/2.0.0/ixod-2.0.0-darwin-amd64?checksum=<checksum>,
    "linux/amd64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/2.0.0/ixod-2.0.0-linux-amd64?checksum=><checksum>"
    }
}

Expects a checksum in the form:

<CHECKSUM>  ixod-<VERSION>-<OS>-<ARCH>[.tar.gz]
<CHECKSUM>  ixod-<VERSION>-<OS>-<ARCH>[.tar.gz]
...

Example:

0711bacaf0cee57f613796ba8c274011e22c3968e98755a105a1a500c87e19f5  ixod-2.0.0-linux-amd64
0859b596ca18257cf424223b35057a4a5296c81fe1e43164673b3344876daaeb  ixod-2.0.0-linux-amd64.tar.gz

(From: https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/sha256sum.txt)

❯ python create_binaries_json.py --tag v2.0.0

Output:
{
    "binaries": {
    "linux/arm64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/ixod-2.0.0-linux-arm64?checksum=<checksum>",
    "darwin/arm64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/ixod-2.0.0-darwin-arm64?checksum=<checksum>",
    "darwin/amd64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/ixod-2.0.0-darwin-amd64?checksum=<checksum>",
    "linux/amd64": "https://github.com/ixofoundation/ixo-blockchain/releases/download/v2.0.0/ixod-2.0.0-linux-amd64?checksum=><checksum>"
    }
}

Expect a checksum to be present at: 
https://github.com/ixofoundation/ixo-blockchain/releases/download/<TAG>/sha256sum.txt
"""

import requests
import json
import argparse
import re
import sys


def validate_tag(tag):
    pattern = '^v[0-9]+.[0-9]+.[0-9]+$'
    return bool(re.match(pattern, tag))


def download_checksums(checksums_url):

    response = requests.get(checksums_url)
    if response.status_code != 200:
        raise ValueError(
            f"Failed to fetch sha256sum.txt. Status code: {response.status_code}")
    return response.text


def checksums_to_binaries_json(checksums):

    binaries = {}

    # Parse the content and create the binaries dictionary
    for line in checksums.splitlines():
        checksum, filename = line.split('  ')

        # exclude tar.gz files
        if not filename.endswith('.tar.gz') and filename.startswith('ixod'):
            try:
                _, tag, platform, arch = filename.split('-')
            except ValueError:
                print(
                    f"Error: Expected binary name in the form: ixod-X.Y.Z-platform-architecture, but got {filename}")
                sys.exit(1)
            _, tag, platform, arch,  = filename.split('-')
            # exclude universal binaries and windows binaries
            if arch == 'all' or platform == 'windows':
                continue
            binaries[f"{platform}/{arch}"] = f"https://github.com/ixofoundation/ixo-blockchain/releases/download/v{tag}/{filename}?checksum=sha256:{checksum}"

    binaries_json = {
        "binaries": binaries
    }

    return json.dumps(binaries_json, indent=2)


def main():

    parser = argparse.ArgumentParser(description="Create binaries json")
    parser.add_argument('--tag', metavar='tag', type=str,
                        help='the tag to use (e.g v2.0.0)')
    parser.add_argument('--checksums_url', metavar='checksums_url',
                        type=str, help='URL to the checksum')

    args = parser.parse_args()

    # Validate the tag format
    if args.tag and not validate_tag(args.tag):
        print("Error: The provided tag does not follow the 'vX.Y.Z' format.")
        sys.exit(1)

    # Ensure that only one of --tag or --checksums_url is specified
    if not bool(args.tag) ^ bool(args.checksums_url):
        parser.error("Only one of tag or --checksums_url must be specified")
        sys.exit(1)

    checksums_url = args.checksums_url if args.checksums_url else f"https://github.com/ixofoundation/ixo-blockchain/releases/download/{args.tag}/sha256sum.txt"
    checksums = download_checksums(checksums_url)
    binaries_json = checksums_to_binaries_json(checksums)
    print(binaries_json)


if __name__ == "__main__":
    main()
