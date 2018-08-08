#!/bin/sh
echo "***********************************"
echo "* IXO BLOCKCHAIN                  *"
echo "***********************************"
echo ""
echo "Preparing IXO Block Sync"
CURRENT_DIR=`dirname $0`
ROOT_DIR=$CURRENT_DIR/..

if [ "$1" != "skip-building" ]
then
    echo "Building Images"
    docker build -t trustlab/ixo-blockchain $ROOT_DIR
fi

if [ "$2" != "expose-over-port80" ]
then
    echo "exposing on port 8080"
    docker-compose up --no-start
else
    echo "exposing on port 80"
    docker-compose -f $ROOT_DIR/docker-compose.yml -f $ROOT_DIR/docker-compose.port80.yml up --no-start
fi

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