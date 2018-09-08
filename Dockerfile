FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos

ARG COMMIT_HASH=''

COPY ./src $SRC_DIR
COPY ./bin/startBlockchain.sh $SRC_DIR/bin/
COPY ./data $SRC_DIR/data
COPY ./Makefile $SRC_DIR

RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519
# COMMIT_HASH is an argument passed from the build-docker-image.sh utility script executed from within the Git repository
RUN cd $SRC_DIR; make COMMIT_HASH=$COMMIT_HASH install

EXPOSE 46656
EXPOSE 46657
EXPOSE 1317
# CMD is configured in the docker-compose.yml