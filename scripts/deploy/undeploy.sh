#!/usr/bin/env bash

# Disable the services
sudo /bin/systemctl disable ixo-rest-server.service
sudo /bin/systemctl disable ixod.service

# Stop the services
sudo /bin/systemctl stop ixo-rest-server.service
sudo /bin/systemctl stop ixod.service

# Remove unit files from /etc/systemd/system/
sudo rm /etc/systemd/system/ixo-rest-server.service
sudo rm /etc/systemd/system/ixod.service

# Reload all unit files
sudo /bin/systemctl daemon-reload
