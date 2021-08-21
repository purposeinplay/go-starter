FROM golang:alpine AS builder

RUN apk add --no-cache ca-certificates git build-base

WORKDIR /app
COPY .. .

RUN go build -o go-starter main.go

CMD tail -f /dev/null

FROM alpine:latest as release

LABEL maintainer="support@purposeinplay.com"

WORKDIR /app

COPY --from=builder /app/go-starter .
COPY --from=builder /app/config.yaml .

CMD ["./go-starter"]