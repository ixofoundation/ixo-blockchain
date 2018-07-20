#!/bin/sh

# This is called to start the blockchain and if no genesis file exist then the blockchain is initalized

echo "***********************************"
echo "* IXO BLOCKCHAIN                  *"
echo "***********************************"
echo ""

echo "Check IXO Block is initialized"
config="/root/.ixo-node/config"
if [ -f "$config" ]
then
	echo "Blockchain already initialized"
else
  echo "**************************************************************"
	echo "Initializing blockchain....."
  echo "**************************************************************"
  ixod init
  echo "**************************************************************"
  echo "Blockchain initialized!"
  echo "**************************************************************"
fi

echo ""
echo "**************************************************************"
echo "Starting Blockchain......"
echo "**************************************************************"
ixod start
