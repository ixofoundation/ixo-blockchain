#!/usr/bin/env bash

echo "Getting sign data..."
REQ='{"msg":"0x7B2274797065223A226469642F416464446964222C2276616C7565223A7B22646964223A226469643A69786F3A55347453707A7A763931484871575731596D466B484A222C227075624B6579223A22466B65447565356974383274616568654D7072646150726374664B3344655656394E6E455059446777775247227D7D","pub_key":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"}'
curl -X POST localhost:1317/txs/sign_data --data-binary "$REQ"

# Decoded request (what we're actually requesting):
# {
#   "msg": {
#     "type": "did/AddDid",
#     "value": {
#       "did": "did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
#       "pubKey": "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"
#     }
#   },
#   "pub_key": "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"
# }