#!/bin/sh

# This is called to start the blockchain and if no genesis file exist then the blockchain is initalized

echo "***********************************"
echo "* IXO BLOCKCHAIN                  *"
echo "***********************************"
echo ""

echo "Check IXO Block is initialized"
config="/root/.ixo-node/config"
echo "Does $config exist? "
if [ -d "$config" ]
then
  echo "YES"
  echo "**************************************************************"
	echo "Blockchain already initialized"
  echo "**************************************************************"
  
  if [ -f "$config/init.lock" ]; then
    echo "**************************************************************"
    echo "Blockchain initialized but currently locked!"
    echo "**************************************************************"

    echo "**************************************************************"
    echo "* NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB!*"
    echo "* EDIT genesis.json WITH LEGITIMATE SMART CONTRACT ADDRESSES!*"
    echo "* AFTER PROVIDING ALL RELEVANT CONFIG DATA, DELETE init.lock *"
    echo "**************************************************************"
  else
    echo "**************************************************************"
    echo "Starting blockchain....."
    echo "**************************************************************"
    ixod start
    echo "**************************************************************"
    echo "Blockchain started!"
    echo "**************************************************************"
  fi
else
  echo "NO"
  echo "**************************************************************"
	echo "Initializing blockchain....."
  echo "**************************************************************"
  ixod init
  touch "$config/init.lock"
fi
