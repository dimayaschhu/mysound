FROM golang:1.16 AS build-env
USER root

RUN apt-get update
RUN apt-get install -y --no-install-recommends \
    protobuf-compiler

RUN groupadd --gid 1000 golang
RUN useradd --uid 1000 --gid golang --shell /bin/bash --create-home golang

RUN mkdir /var/log/rest
RUN chown golang:golang /var/log/rest

RUN mkdir /run/rest
RUN chown golang:golang /run/rest
RUN mkdir /run/rest/latest
RUN chown golang:golang /run/rest/latest

USER golang

WORKDIR /go