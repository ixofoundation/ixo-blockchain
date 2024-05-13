# echo $(git describe --tags) | sed 's/^v//'
ARG GIT_VERSION="2.0.0-rc.0"
# git log -1 --format='%H'
ARG GIT_COMMIT="ca9e64cf2f8c29b8bb001281abcdcb942dc9fa01"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:1.22.3-alpine as builder

RUN apk add --no-cache \
  ca-certificates \
  build-base \
  linux-headers

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

# RUN make build
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

FROM ubuntu:20.04 as runner
# FROM gcr.io/distroless/base-debian11 as run

COPY --from=builder /ixo/build/ixod /bin/ixod

ENV HOME /ixo
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 26660

ENTRYPOINT [ "ixod" ]
