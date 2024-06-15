FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

EXPOSE 8040

CMD ["CompileDaemon", "-command=./first-go-crud"]
