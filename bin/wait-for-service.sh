#!/bin/bash

SERVICE="$1"
CONFIRMATION_TEXT="$2"
RETRY_LIMIT=($3)

MATCH=$(docker logs $SERVICE | grep "$CONFIRMATION_TEXT")
MATCH_LENGTH=${#MATCH}

COUNTER=0
while [[ $MATCH_LENGTH -eq 0 && $COUNTER -lt $RETRY_LIMIT ]] ; do
    ((COUNTER++))
    sleep 1

    echo "waiting for $SERVICE (retry $COUNTER of $RETRY_LIMIT)"
    MATCH=$(docker logs $SERVICE | grep "$CONFIRMATION_TEXT")
    MATCH_LENGTH=${#MATCH}    
done

echo "$MATCH"