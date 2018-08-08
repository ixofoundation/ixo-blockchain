#!/bin/bash

echo "***********************************"
echo "* IXO BLOCKCHAIN SHUTDOWN         *"
echo "***********************************"
echo ""
docker-compose stop
docker-compose rm
echo ""
echo "***********************************"
echo "* IXO BLOCKCHAIN SHUTDOWN COMPLETE*"
echo "***********************************"