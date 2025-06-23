FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tasks

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/tasks .

ENTRYPOINT ["./tasks"]
