FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos
COPY ./app $SRC_DIR/app
# COPY ./bin/startBlockchain.sh $SRC_DIR/bin
COPY ./bin $SRC_DIR/bin
COPY ./cmd/ixod $SRC_DIR/ixod
COPY ./cmd/ixocli $SRC_DIR/ixocli
COPY ./data $SRC_DIR/data
COPY ./types $SRC_DIR/types
COPY ./x $SRC_DIR/x
COPY ./Makefile $SRC_DIR
COPY ./vendor $SRC_DIR/vendor
# COPY . $SRC_DIR

RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519

# RUN go get $SRC_DIR/ixod
# RUN go get $SRC_DIR/ixocli

RUN ls -alh $SRC_DIR/ixod
RUN cd $SRC_DIR/ixod; go install

EXPOSE 46656
EXPOSE 46657
EXPOSE 1317
# CMD is configured in the docker-compose.yml