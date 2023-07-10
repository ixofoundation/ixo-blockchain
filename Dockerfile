FROM golang:1.19.4 as build

WORKDIR /app

COPY . .

RUN make build

FROM ubuntu:20.04 as run
# FROM gcr.io/distroless/base-debian11 as run

COPY --from=build /app/build/ixod /bin/ixod
COPY --from=build /lib/x86_64-linux-gnu/libgcc_s.so.1 /lib/x86_64-linux-gnu/libgcc_s.so.1

ENV HOME /ixo
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 26660

ENTRYPOINT [ "ixod" ]
