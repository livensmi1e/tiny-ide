FROM golang:1.24.0-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/ide /app/cmd/main.go

EXPOSE 8000

CMD [ "/app/cmd/ide" ]