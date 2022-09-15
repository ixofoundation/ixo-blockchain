##
## Fetch Golang Dependencies
##
FROM golang:1.18 AS build-dependencies
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

FROM build-tests AS build-devel
RUN make install

# ENTRYPOINT [ "starport", "chain", "serve" ]

# # Build Golang files
# FROM build-dependencies AS build-binary
# COPY ./ ./

# ENV IXOD_HOME="/home/ixo/"
# RUN mkdir -p ${IXOD_HOME}/data

# RUN go mod tidy && go mod vendor
# RUN make install

# # RUN tar -xzf /dist/app/ixo-blockchain_linux_amd64.tar.gz -C /dist/app/

# # FROM ubuntu:latest AS setup-genisis
# # COPY ./.infra/dockerfiles/setups/setup-genisis.sh /setup/setup-genisis.sh
# # RUN /setup/setup-genisis.sh

# # RUN curl https://raw.githubusercontent.com/ixofoundation/genesis/master/pandora-4/genesis.json

# ENTRYPOINT ["ixod", "start"]

##
## Build Distrubutiable Blockchain
##

FROM alpine:latest

COPY --from=build-dependencies /usr/local/go/ /usr/local/go/
COPY --from=build-dependencies /app/ /app/

ENV O_PATH="${PATH}"
ENV PATH="/usr/local/go/bin:${PATH}"

RUN go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixod &&\
  rm -rf /usr/local/go/ &&\
  rm -rf /app/

ENV PATH="${O_PATH}"

  # FROM gcr.io/distroless/base-debian10 AS distributable

# COPY --from=build-binary /build/ixod /app/ixod
# # COPY --from=build-binary /dist/app/.ixod/ /app/.ixod/

# ONBUILD LABEL world.ixo.blockchain.image.authors="willem.mouton@ixo.world"

# ONBUILD WORKDIR /app

# ONBUILD EXPOSE 1317
# ONBUILD EXPOSE 26657

# ONBUILD ENV IXOD_ENV=prod
# ONBUILD ENV IXOD_HOME=/app/.ixod-$IXOD_ENV

ONBUILD ENTRYPOINT ["ixod", "--home", "${IXOD_HOME}", "start"]

# FROM distributable AS dist-devel
# # RUN echo "Fetch genisis file"

# FROM distributable AS dist-beta
# # RUN echo "Fetch genisis file"

# FROM distributable AS dist-rc
# # RUN echo "Fetch genisis file"

# FROM distributable AS dist-stable
# # RUN echo "Fetch genisis file"
