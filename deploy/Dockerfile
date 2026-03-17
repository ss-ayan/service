FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ayan-server ./cmd/service

FROM alpine:3.19
RUN apk add --no-cache ca-certificates git
WORKDIR /app

COPY --from=builder /ayan-server /usr/local/bin/ayan-server
COPY ayan.yaml /app/ayan.yaml

EXPOSE 8080
ENTRYPOINT ["ayan-server"]
