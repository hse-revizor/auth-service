FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o srvr ./cmd/service/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/srvr .

EXPOSE 8779

CMD ["./srvr", "-env-mode=production", "-config-path=/cfg/config.yaml"]