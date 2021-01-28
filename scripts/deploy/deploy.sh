#!/usr/bin/env bash

# Copy unit files to /etc/systemd/system/
sudo cp ixo-rest-server.service /etc/systemd/system/
sudo cp ixod.service /etc/systemd/system/

# Reload all unit files
sudo /bin/systemctl daemon-reload

# Enable the services
sudo /bin/systemctl enable ixo-rest-server.service
sudo /bin/systemctl enable ixod.service

# Start the services
sudo /bin/systemctl restart ixo-rest-server.service
sudo /bin/systemctl restart dixod.service
