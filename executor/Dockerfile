FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main /app/main.go

FROM ubuntu:22.04 AS server

USER root

RUN apt-get update && apt-get install -y \
    python3 python3-pip \
    gcc g++ \
    curl \
    time \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN ln -s /usr/bin/python3 /usr/bin/python

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 50001

ENTRYPOINT [ "./main" ]