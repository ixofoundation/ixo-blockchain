# ixo Hub

## Setup a new ixo.hub node

### Prerequisites:
Ensure docker and docker compose are installed

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