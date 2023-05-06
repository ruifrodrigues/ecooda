# Base stage (installing dependencies)
FROM golang:1.20.4-buster as base

ARG GRPC_HEALTH_PROBE_VERSION=v0.4.18

ENV BUILD_HOME $GOPATH/src/github.com/ruifrodrigues/ecooda
ENV GOPRIVATE=github.com/dealroom/*

RUN apt update && apt upgrade -y && \
    apt install --no-install-recommends -y git make

RUN go install github.com/cosmtrek/air@latest

RUN curl -fLo /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
    && chmod +x /bin/grpc_health_probe

WORKDIR $BUILD_HOME
COPY . .

RUN --mount=type=secret,id=netrc,dst=/root/.netrc set -xe \
    && go mod download \
    && go mod verify

# Air stage (used for local env)
FROM base as air-server

CMD air -c .air.server.toml

# Air stage (used for local env)
FROM base as air-consumer

CMD air -c .air.consumer.toml

# Builder stage (builds app binary)
FROM base as builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

# Final stage (production-ready image)
FROM alpine:3

ENV APP_HOME /ecooda
ENV APP_PORT 50051

COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

WORKDIR $APP_HOME

COPY .env.example *.env ./
COPY --from=builder /go/bin/app $APP_HOME

EXPOSE $APP_PORT
