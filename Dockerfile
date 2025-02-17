FROM golang:1.24.0-alpine3.21

USER root

RUN apk update && apk add --no-cache \
    curl \
    bash \
    ca-certificates \
    device-mapper \
    iptables \
    lxc \
    tar \
    xz

# Cài Docker theo cách thủ công
RUN curl -fsSL https://download.docker.com/linux/static/stable/x86_64/docker-20.10.24.tgz -o /tmp/docker.tgz \
    && tar -xvzf /tmp/docker.tgz -C /usr/local/bin --strip-components=1 \
    && rm /tmp/docker.tgz

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/ide /app/cmd/main.go

EXPOSE 8000

CMD [ "/app/cmd/ide" ]