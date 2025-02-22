FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/ide /app/cmd/main.go

FROM alpine:3.21.3 AS server

WORKDIR /app

COPY --from=builder /app/cmd/ide .

EXPOSE 8000

ENTRYPOINT [ "./ide" ]