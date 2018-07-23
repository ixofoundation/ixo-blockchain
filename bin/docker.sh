#!/bin/sh
echo "***********************************"
echo "* IXO BLOCKCHAIN                  *"
echo "***********************************"
echo ""
echo "Build IXO Block Sync"
CURRENT_DIR=`dirname $0`
ROOT_DIR=$CURRENT_DIR/..

docker build -t trustlab/ixo-blockchain -f - $ROOT_DIR < $ROOT_DIR/docker/blockchain/Dockerfile
docker-compose up -d

docker-compose up --no-start
# docker-compose create
docker-compose start block-sync-db

# attempting to wait for mongodb to be ready
$ROOT_DIR/bin/wait-for-service.sh block-sync-db 'waiting for connections on port' 10
docker-compose start blockchain
docker-compose start rest
docker-compose start block-sync

echo "Starting IXO Blockchain ..."
sleep 5
echo ${green} "done"
docker-compose logs --tail 13 blockchain
echo ""
echo "***********************************"
echo "* IXO BLOCKCHAIN COMPLETE          *"
echo "***********************************"
docker-compose ps