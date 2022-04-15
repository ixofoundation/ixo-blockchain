#!/bin/bash

echo "Ixo Chain Installer"
echo -e "-------------------------\n\n"
sleep 2

if [ "$USER" != "root" ]; then
        echo "You must be logged in as root to use this installer!"
        exit 0;
fi

echo "Starting IXO based install"
sleep 1
echo "Fetching Genisis Config"
apt-get update -y
apt-get upgrade -y
apt-get install -y git
cd ~


git clone 
mkdir /home/ixo/.ixod
mkdir /home/ixo/.ixod/config

cp /root/genesis/pandora-4/genesis.json /home/ixo/.ixod/config/genesis.json

chown -R ixo:ixo /home/ixo/.$DAEMONNAME
chown -R ixo:ixo /home/ixo/.$DAEMONNAME/config/
chown -R ixo:ixo /home/ixo/.$DAEMONNAME/config/genesis.json


su $USERNAME <<EOSU
$DAEMON init "Pandora node"
EOSU

sleep 5

cp /root/genesis/pandora-4/genesis.json /home/ixo/.ixod/config/genesis.json

chown -R ixo:ixo /home/ixo/.$DAEMONNAME/config/genesis.json


echo "---"
echo "Your peer ID:"
$DAEMON tendermint show-node-id
echo "---"

  cat << EOF > /etc/systemd/system/$DAEMONNAME.service
# /etc/systemd/system/$DAEMONNAME.service
[Unit]
Description=$DAEMONNAME Node
After=network.target
[Service]
Type=simple
User=$USERNAME
WorkingDirectory=$HOME
ExecStart=$DAEMON start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload

sleep 3

systemctl enable $DAEMONNAME.service

echo "Service created at /etc/systemd/system/$DAEMONNAME.service."
echo "Run 'systemctl start $DAEMONNAME.service' to start the node"