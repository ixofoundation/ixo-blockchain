import requests
from base64 import b64decode
from typing import List, Dict
from urllib.parse import quote

URL = "http://localhost:26657"
IGNORED_EVENTS = \
    ["message", "transfer", "commission", "rewards", "proposer_reward"]


def print_events(events: List[Dict]):
    for evt in events:
        if evt['type'] in IGNORED_EVENTS:
            continue
        print('----------------------------', evt['type'])
        for attr in evt['attributes']:
            key = b64decode(attr['key']).decode('utf-8')
            if attr['value'] is not None:
                val = b64decode(attr['value']).decode('utf-8')
                print("{} = {}".format(key, val))
            else:
                print("{} = {}".format(key, None))


for height in range(1, 100):
    # Block events
    res = requests.get("{}/block_results?height={}".format(URL, height)).json()
    sections = ['begin_block', 'end_block']
    for section in sections:
        res_section = res['result']['results'][section]
        if 'events' in res_section:
            print_events(res_section['events'])

    # Transaction events
    query = quote("\"tx.height={}\"".format(height))
    params = "query=" + query + "&prove=true&page=1&per_page=30&order_by=asc"
    res = requests.get("{}/tx_search?{}".format(URL, params)).json()
    for tx in res['result']['txs']:
        print_events(tx['tx_result']['events'])
