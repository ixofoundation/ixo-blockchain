##
## Fetch Golang Dependencies
##
FROM golang:latest AS fetch-dependencies
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod vendor

# ## ----------------------------------- ##

# ##
# ## Run Devel Build
# ##
# ARG GO_ENABLED=true
# ARG GO_VERSION=latest
# ARG NODE_ENABLED=true
# ARG NODE_VERSION=latest

# FROM golang:latest AS build


# # Setup Environment
# WORKDIR /app
# ENV DEBIAN_FRONTEND=noninteractive
# # RUN chsh --shell /bin/bash
# # Copy Source Files
# COPY ./ ./

# # Copy Dependencies
# COPY --from=fetch-dependencies-golang /app/vendor ./vendor
# RUN

# # Copy Entrypint script
# COPY ./.deploy/dockerfiles/devel.entrypoint.sh ./
# ENTRYPOINT [ "/bin/bash", "./devel.entrypoint.sh" ]

# Build Golang files
FROM golang:latest AS build-blockchain
WORKDIR /app

# COPY --from=fetch-dependencies /app ./
# RUN mkdir -p /dist/app/
COPY ./ ./

RUN curl https://get.starport.network/starport! | bash

COPY ./dist/app/ /dist/app/
# RUN starport chain build --release -o /dist/app -v -t linux:amd64 -t darwin:amd64
# RUN starport chain init --home /dist/app/.ixod

RUN tar -xzf /dist/app/ixo-blockchain_linux_amd64.tar.gz -C /dist/app/

# FROM ubuntu:latest AS setup-genisis
# COPY ./.infra/dockerfiles/setups/setup-genisis.sh /setup/setup-genisis.sh
# RUN /setup/setup-genisis.sh

# RUN curl https://raw.githubusercontent.com/ixofoundation/genesis/master/pandora-4/genesis.json


FROM gcr.io/distroless/base-debian10 AS run-app

WORKDIR /app

# COPY --from=setup-genisis /dist/app/ixo-blockchaind /app/ixod
# RUN 

# chown -R ixo:ixo /home/ixo/.$DAEMONNAME/config/
# chown -R ixo:ixo /home/ixo/.$DAEMONNAME/config/genesis.json

COPY --from=build-blockchain /dist/app/ixo-blockchaind /app/ixod
COPY --from=build-blockchain /dist/app/.ixod/ /app/.ixod/
# RUN chown -R nonroot:nonroot /app/.ixod

EXPOSE 8080

# USER nonroot:nonroot


ENTRYPOINT ["/app/ixod", "start", "--home", "/app/.ixod"]