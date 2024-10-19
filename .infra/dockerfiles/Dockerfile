# Dockerfile for docker environment to build and run the ixo blockchain

FROM --platform=linux/amd64 golang:latest AS starport
RUN curl https://get.starport.network/starport! | bash

FROM starport AS build-dependencies
WORKDIR /app
ONBUILD WORKDIR /app

FROM build-dependencies AS build-devel
COPY ./ .
