# echo $(git describe --tags) | sed 's/^v//'

ARG GO_VERSION="1.22"
# ARG RUNNER_IMAGE="ubuntu:latest"
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"
ARG BUILD_TAGS="netgo,ledger,muslc"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine3.18 AS builder

# TODO: maybe extract below args to where called in cicd?
# git log -1 --format='%H'
ARG GIT_VERSION="4.0.0-rc.0"
ARG GIT_COMMIT="ca9e64cf2f8c29b8bb001281abcdcb942dc9fa01"

RUN apk add --no-cache \
  ca-certificates \
  build-base \
  linux-headers \
  binutils-gold

# Download go dependencies
WORKDIR /ixo
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/root/go/pkg/mod \
  go mod download

# RUN ARCH=$(uname -m)
# Cosmwasm - Download correct libwasmvm version
RUN ARCH=x86_64 WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && wget \
  https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
  -O /lib/libwasmvm_muslc.a && wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt \
  -O /tmp/checksums.txt && sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1)

# Copy the remaining files
COPY . .

# Build ixod binary
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/root/go/pkg/mod \
  GOWORK=off go build \
  -mod=readonly \
  -tags "netgo,ledger,muslc" \
  -ldflags \
  "-X github.com/cosmos/cosmos-sdk/version.Name="ixo" \
  -X github.com/cosmos/cosmos-sdk/version.AppName="ixod" \
  -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
  -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
  -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
  -w -s -linkmode=external -extldflags '-Wl,-z,muldefs --static -lm'" \
  -trimpath \
  -o /ixo/build/ixod \
  ./cmd/ixod/main.go

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ${RUNNER_IMAGE}

COPY --from=builder /ixo/build/ixod /bin/ixod

ENV HOME=/ixo
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 26660

ENTRYPOINT [ "ixod" ]
