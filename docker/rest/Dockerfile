FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos
ADD . $SRC_DIR
EXPOSE 1317
RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519; cd $SRC_DIR; make build; make install
CMD ["ixocli", "rest-server", "-n", "tcp://blockchain:46657", "-a", "tcp://0.0.0.0:1317"]