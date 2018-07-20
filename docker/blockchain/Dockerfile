FROM iron/go:dev
ENV SRC_DIR=/go/src/github.com/ixofoundation/ixo-cosmos
COPY . $SRC_DIR
EXPOSE 46657
RUN go get github.com/btcsuite/btcutil/base58 golang.org/x/crypto/ed25519
RUN cd $SRC_DIR; make build; make install
CMD ["/go/src/github.com/ixofoundation/ixo-cosmos/bin/startBlockchain.sh"]
#CMD ["ixod", "start"]