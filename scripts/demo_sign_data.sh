#!/usr/bin/env bash

echo "Getting sign data..."
REQ='{"msg":"0x7B2274797065223A226469642F416464446964222C2276616C7565223A7B22646964446F63223A7B22646964223A226469643A69786F3A55347453707A7A763931484871575731596D466B484A222C227075624B6579223A22466B65447565356974383274616568654D7072646150726374664B3344655656394E6E455059446777775247222C2263726564656E7469616C73223A5B5D7D7D7D","pub_key":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"}'
curl -X POST localhost:1317/sign_data --data-binary "$REQ"

# Decoded request (what we're actually requesting):
# {
#   "msg": {
#     "type": "did/AddDid",
#     "value": {
#       "didDoc": {
#         "did": "did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
#         "pubKey": "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
#         "credentials": []
#       }
#     }
#   },
#   "pub_key": "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG"
# }