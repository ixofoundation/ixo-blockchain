#!/usr/bin/env bash

# Copy unit file to /etc/systemd/system/
sudo cp ixod.service /etc/systemd/system/

# Reload all unit files
sudo /bin/systemctl daemon-reload

# Enable and start the service
sudo /bin/systemctl enable ixod.service
sudo /bin/systemctl restart ixod.service
