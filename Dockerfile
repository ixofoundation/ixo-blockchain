FROM iron/go:dev

ARG COMMIT_HASH=''

ENV FOUNDATION_PATH=/go/src/github.com/ixofoundation
ENV COSMOS_HOME=$FOUNDATION_PATH/ixo-cosmos

# Copy the local package files to the container's workspace.
ADD . $COSMOS_HOME
#COPY ./app $COSMOS_HOME/app
#COPY ./client $COSMOS_HOME/client
#COPY ./cmd $COSMOS_HOME/cmd
#COPY ./types $COSMOS_HOME/types
#COPY ./vendor $COSMOS_HOME/vendor
#COPY ./x $COSMOS_HOME/x
#COPY ./Makefile $COSMOS_HOME
#COPY ./bin/startBlockchain.sh $COSMOS_HOME/bin/

# Manage global dependencies.
RUN cd $FOUNDATION_PATH
RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/btcsuite/btcutil/base58
RUN go get github.com/ethereum/go-ethereum

# Manage vendor dependencies.
# RUN cd $COSMOS_HOME; make get_vendor_deps

# Build, Install
# COMMIT_HASH is an optional argument passed from the build-docker-image.sh utility script executed from within the Git repository
RUN cd $COSMOS_HOME; make COMMIT_HASH=$COMMIT_HASH build
RUN cd $COSMOS_HOME; make COMMIT_HASH=$COMMIT_HASH install

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
