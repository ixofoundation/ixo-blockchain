## Download latest snapshot

### 1. Stop the validator

```bash
systemctl stop ixod
```

### 2. Backup validator state

If you are already running a validator, be sure you backed up your priv_validator_key.json prior to removing your current data or getting the new snapshot.

```
cp ~/.ixod/data/priv_validator_state.json ~/
```

### 3. Remove old data in directory `~/.ixod/data`

```bash
rm -rf ~/.ixod/data; \

mkdir -p ~/.ixod/data; \
cd ~/.ixod/data
```

### 4. Download snapshot

```bash
SNAP_NAME=$(curl -s https://snapshots.stake2.me/ixo/ | egrep -o ">ixo.*tar" | tr -d ">" | tail -n1); \
wget -O - https://snapshots.stake2.me/ixo/${SNAP_NAME} | tar xf -
```

### 5. Start service and check logs

```bash
systemctl start ixod; \
journalctl -u ixod -f --no-hostname
```
