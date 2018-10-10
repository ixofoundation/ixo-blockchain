FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos
COPY . $SRC_DIR
RUN go get github.com/btcsuite/btcutil/base58 
RUN go get github.com/ethereum/go-ethereum
RUN cd $SRC_DIR; make build; make install
EXPOSE 46656
EXPOSE 46657
EXPOSE 1317
# CMD is configured in the docker-compose.yml