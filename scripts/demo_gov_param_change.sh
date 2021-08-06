#!/usr/bin/env bash

# IT IS RECOMMENDED TO RUN THE BLOCKCHAIN USING run_with_all_data_dev.sh SINCE
# THIS SETS GOVERNANCE PERIODS TO 30 seconds FOR FASTER GOVERNANCE

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(ixod status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(ixod status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

GAS_PRICES="0.025uixo"
CHAIN_ID="pandora-4"

ixod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  ixod tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

ixod_q() {
  ixod q "$@" --output=json | jq .
}

echo "Query transfer params before param change"
ixod_q ibc-transfer params

echo "Submitting param change proposal"
ixod_tx gov submit-proposal param-change demo_gov_param_change_proposal.json --from=miguel

echo "Query proposal 1"
ixod_q gov proposal 1

echo "Depositing 10000000uixo to reach minimum deposit"
ixod_tx gov deposit 1 10000000uixo --from=miguel

echo "Query proposal 1 deposits"
ixod_q gov deposits 1

echo "Voting yes for proposal"
ixod_tx gov vote 1 yes --from=miguel

echo "Query proposal 1 tally"
ixod_q gov tally 1

echo "Waiting for proposal to pass..."
while :; do
  RET=$(ixod_q gov proposal 1 2>&1)
  if [[ ($RET == *'PROPOSAL_STATUS_VOTING_PERIOD'*) ]]; then
    sleep 1
  else
    echo "A few more seconds..."
    sleep 6
    break
  fi
done

echo "Query transfer params (expected to be true and false)"
ixod_q ibc-transfer params
