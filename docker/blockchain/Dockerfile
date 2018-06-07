FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos
ADD . $SRC_DIR
EXPOSE 46657
RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519; cd $SRC_DIR; make build; make install; ixod init
CMD ["ixod", "start"]