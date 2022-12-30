## Download latest snapshot  
Stop IXO service  
`systemctl stop ixo.service`  

Remove old data in directory `~/.ixod/data`  
```
rm -rf ~/.ixod/data; \
mkdir -p ~/.ixod/data; \
cd ~/.ixod/data
```

Download snapshot  
```bash
SNAP_NAME=$(curl -s https://snapshots.stake2.me/ixo/ | egrep -o ">ixo.*tar" | tr -d ">" | tail -n1); \
wget -O - https://snapshots.stake2.me/ixo/${SNAP_NAME} | tar xf -
```

Start service and check logs  
```
systemctl start ixo.service; \
journalctl -u ixo.service -f --no-hostname
```
