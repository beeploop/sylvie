FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build API executable
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/api ./cmd/api

# build Processor executable
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/processor ./cmd/processor

# API final build
FROM gcr.io/distroless/static-debian12 AS api

COPY --from=builder /out/api /api

ENTRYPOINT ["/api"]

# Processor final build
FROM alpine:latest AS processor

RUN apk add --no-cache ffmpeg ca-certificates

COPY --from=builder /out/processor /processor

ENTRYPOINT ["/processor"]
