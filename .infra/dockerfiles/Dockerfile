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
# ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
# ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
# RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 7d2239e9f25e96d0d4daba982ce92367aacf0cbd95d2facb8442268f2b1cc1fc
# RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep f6282df732a13dec836cda1f399dd874b1e3163504dbd9607c6af915b2740479

# CosmWasm: copy the right library according to architecture. The final location will be found by the linker flag `-lwasmvm_muslc`
# RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a

# RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build
COPY ./ .

ENTRYPOINT [ "starport", "chain", "serve" ]

# Build Golang files
FROM build-dependencies AS build-binary
COPY ./ ./

ENV IXOD_HOME="/home/ixo/"
RUN mkdir -p ${IXOD_HOME}/iid

RUN go mod tidy && go mod vendor
RUN make install

# RUN tar -xzf /dist/app/ixo-blockchain_linux_amd64.tar.gz -C /dist/app/

# FROM ubuntu:latest AS setup-genisis
# COPY ./.infra/dockerfiles/setups/setup-genisis.sh /setup/setup-genisis.sh
# RUN /setup/setup-genisis.sh

# RUN curl https://raw.githubusercontent.com/ixofoundation/genesis/master/pandora-4/genesis.json

ENTRYPOINT ["ixod", "start"]

##
## Build Distrubutiable Blockchain
##

FROM ubuntu:latest  as distributable

COPY --from=build-binary /go/bin/ixod /app/ixod
COPY --from=build-binary /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.0.0/api/ /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.0.0/api/

ONBUILD LABEL world.ixo.blockchain.image.authors="willem.mouton@ixo.world"

ONBUILD WORKDIR /app

ONBUILD EXPOSE 1317
ONBUILD EXPOSE 26657

ONBUILD ENTRYPOINT ["/app/ixod", "start", "--rpc.laddr", "tcp://0.0.0.0:26657"]

FROM distributable AS dist-devel
# RUN echo "Fetch genisis file"

FROM distributable AS dist-beta
# RUN echo "Fetch genisis file"

FROM distributable AS dist-rc
# RUN echo "Fetch genisis file"

FROM distributable AS dist-stable
# RUN echo "Fetch genisis file"
