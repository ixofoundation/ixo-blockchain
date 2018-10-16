FROM iron/go:dev

ARG COMMIT_HASH=''

ENV DESTINATION_PATH=/go/src/github.com/ixofoundation
ENV COSMOS_HOME=ixo-cosmos

# Copy the local package files to the container's workspace.
ADD . $DESTINATION_PATH/$COSMOS_HOME

# Manage global dependencies.
RUN cd $DESTINATION_PATH
RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/btcsuite/btcutil/base58
RUN go get github.com/ethereum/go-ethereum

# Manage vendor dependencies.
# RUN cd $DESTINATION_PATH/$COSMOS_HOME; make get_vendor_deps

# Build, Install
# COMMIT_HASH is an optional argument passed from the build-docker-image.sh utility script executed from within the Git repository
RUN cd $DESTINATION_PATH/$COSMOS_HOME; make COMMIT_HASH=$COMMIT_HASH build
RUN cd $DESTINATION_PATH/$COSMOS_HOME; make COMMIT_HASH=$COMMIT_HASH install

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
