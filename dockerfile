FROM golang:1.22

WORKDIR /gateway-work
COPY . ./

RUN apt-get update
WORKDIR /gateway-work/gateway
RUN go mod download