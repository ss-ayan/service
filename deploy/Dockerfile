# syntax=docker/dockerfile:1
FROM golang:1.27-bookworm AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
ENV GOPRIVATE=github.com/ss-ayan

# github.com/ss-ayan/ayan is private. Pass a token with repo read access as
# the `gh_token` build secret; it only lives for this RUN and never lands in a layer.
COPY go.mod go.sum ./
RUN --mount=type=secret,id=gh_token,required=true \
    GIT_CONFIG_COUNT=1 \
    GIT_CONFIG_KEY_0="url.https://x-access-token:$(cat /run/secrets/gh_token)@github.com/.insteadOf" \
    GIT_CONFIG_VALUE_0="https://github.com/" \
    go mod download

COPY . .
RUN GOOS=linux go build -o /ayan-server ./cmd/service

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app

COPY --from=builder /ayan-server /usr/local/bin/ayan-server
# Env-driven default config; mount your own over /app/ayan.yaml to customise.
COPY ayan.example.yaml /app/ayan.yaml

EXPOSE 8080
ENTRYPOINT ["ayan-server"]
