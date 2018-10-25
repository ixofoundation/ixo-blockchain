# ixo Hub

## Setup a new ixo.hub node

### Prerequisites:
* Ensure docker and docker compose are installed
* Expose the following ports: 
`26656`, `26657`,`1317`


### Setup
* Create an `ixo` folder in a place of your choice on your host
* Copy the following files to the `ixo` folder created above
`purge.sh`
`start.sh`
`docker-compose.yml`

#### Connect to existing chain
* Create ixo node config
* From `ixo` folder run the following:
```
mkdir data
mkdir data\blockchain
```
* Copy your tendermint node config folder to `data\blockchain`

### Start the node
* Run `./start.sh`

### Purge all images (The next start will pull the latest images)
* Run `./purge.sh`


## Initial Configuration

### Steps
- run `ixod init` (This should be run on each node)
  - this creates a *genesis.json* file in the *ixo-node/config* folder
  - It also generates a nodeDid and Ethereum wallet
- Edit the *genesis.json* file from one of the nodes and setup the chainID, validators, nodes, fees and Ethereum wallets in that file. Copy the file to the other nodes.  It should be the same.
- The PDS needs to be configured with one of the nodeDID's. This will be used to identify the initiating node for project creation and fees. The Ethereum network also needs to correspond to the network is configured on the blockchain (see PDS config for more info)
- Fund the node's Ethereum wallet with ETH so that gas costs can be paid for authorising transactions
- Setup the environment variable on the blockchain server/container with the RPC value for the Ethereum connection. e.g. (mainnet, ropsten and localhost respectively)
  - ETH_URL=https://api.infura.io/v1/jsonrpc/mainnet
  - ETH_URL=https://api.infura.io/v1/jsonrpc/ropsten
  - ETH_URL=http://localhost:7545
