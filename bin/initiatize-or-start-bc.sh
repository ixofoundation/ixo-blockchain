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
  echo "**************************************************************"
	echo "Blockchain already initialized"
  echo "**************************************************************"
  ixod start
  echo "**************************************************************"
  echo "Blockchain started!"
  echo "**************************************************************"
else
  echo "**************************************************************"
	echo "Initializing blockchain....."
  echo "**************************************************************"
  ixod init
  echo "**************************************************************"
  echo "Blockchain initialized!"
  echo "**************************************************************"
  echo "* NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB! NB!*"
  echo "* EDIT THE genesis.json FILE BEFORE STARTING IXOD SERVICE !! *"
  echo "* CONTENTS OF genesis.json IS ONLY USED ON FIRST RUN     !!! *"
fi
