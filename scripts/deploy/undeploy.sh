#!/usr/bin/env bash

# Disable, stop, and remove unit file
sudo /bin/systemctl disable ixod.service
sudo /bin/systemctl stop ixod.service
sudo rm /etc/systemd/system/ixod.service

# Reload all unit files and reset failed
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl reset-failed
