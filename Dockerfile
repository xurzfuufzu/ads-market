FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o app ./cmd

FROM alpine

WORKDIR /app

COPY --from=builder /build/app ./app
COPY --from=builder /build/migrations ./migrations

ENTRYPOINT ["./app"]