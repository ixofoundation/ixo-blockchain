FROM golang:1.19.4 as build

WORKDIR /app

COPY . .

RUN make build

FROM gcr.io/distroless/base-debian11 as run

COPY --from=build /app/build/ixod /bin/ixod
COPY --from=build /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.1.1/internal/api/ /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.1.1/internal/api/
COPY --from=build /lib/x86_64-linux-gnu/libgcc_s.so.1 /lib/x86_64-linux-gnu/libgcc_s.so.1

ENV HOME /ixo
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 26660

ENTRYPOINT [ "ixod" ]
