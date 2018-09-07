FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos

COPY ./src $SRC_DIR
COPY ./bin/startBlockchain.sh $SRC_DIR/bin/
COPY ./data $SRC_DIR/data
COPY ./Makefile $SRC_DIR

RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519

# RUN cd $SRC_DIR/ixod; go install
RUN cd $SRC_DIR; make install

EXPOSE 46656
EXPOSE 46657
EXPOSE 1317
# CMD is configured in the docker-compose.yml