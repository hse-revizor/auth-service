FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o ./cmd/service/main ./cmd/service/main.go

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/templates
COPY --from=builder /app/cmd/service/main /app/cmd/service/main
COPY --from=builder /app/templates/* /app/templates/
