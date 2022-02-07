## Starport

FROM golang:latest AS starport
RUN curl https://get.starport.network/starport! | bash

##
## Fetch Golang Dependencies
##
FROM starport AS build-dependencies
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod vendor

ONBUILD WORKDIR /app

##
## Fetch Golang Dependencies
##

FROM build-dependencies AS build-tests
RUN go test -v ./...

FROM build-dependencies AS build-devel
COPY ./ ./

ENTRYPOINT [ "starport", "chain", "serve" ]

# Build Golang files
FROM build-dependencies AS build-binary
COPY ./ ./

RUN starport chain build --release -o /dist/app -v -t linux:amd64 -t darwin:amd64
RUN tar -xzf /dist/app/ixo-blockchain_linux_amd64.tar.gz -C /dist/app/

# FROM ubuntu:latest AS setup-genisis
# COPY ./.infra/dockerfiles/setups/setup-genisis.sh /setup/setup-genisis.sh
# RUN /setup/setup-genisis.sh

# RUN curl https://raw.githubusercontent.com/ixofoundation/genesis/master/pandora-4/genesis.json

##
## Build Distrubutiable Blockchain
##

FROM gcr.io/distroless/base-debian10 AS distributable

COPY --from=build-binary /dist/app/ixo-blockchaind /app/ixod
# COPY --from=build-binary /dist/app/.ixod/ /app/.ixod/

ONBUILD LABEL world.ixo.blockchain.image.authors="willem.mouton@ixo.world"

ONBUILD WORKDIR /app

ONBUILD EXPOSE 1317
ONBUILD EXPOSE 26657

ONBUILD ENTRYPOINT ["/app/ixod", "--home", "/app/.ixod"]

FROM distributable AS dist-devel
# RUN echo "Fetch genisis file"

FROM distributable AS dist-beta
# RUN echo "Fetch genisis file"

FROM distributable AS dist-rc
# RUN echo "Fetch genisis file"

FROM distributable AS dist-stable
# RUN echo "Fetch genisis file"
