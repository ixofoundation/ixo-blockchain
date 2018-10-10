FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos

ARG COMMIT_HASH=''

COPY ./app $SRC_DIR/app
COPY ./client $SRC_DIR/client
COPY ./cmd $SRC_DIR/cmd
COPY ./types $SRC_DIR/types
COPY ./vendor $SRC_DIR/vendor
COPY ./x $SRC_DIR/x
COPY ./Makefile $SRC_DIR
COPY ./bin/startBlockchain.sh $SRC_DIR/bin/

RUN go get github.com/btcsuite/btcutil/base58
RUN ls -alh $SRC_DIR
# COMMIT_HASH is an argument passed from the build-docker-image.sh utility script executed from within the Git repository
RUN cd $SRC_DIR; make COMMIT_HASH=$COMMIT_HASH install

EXPOSE 46656
EXPOSE 46657
EXPOSE 1317
# CMD is configured in the docker-compose.yml
