FROM golang:1.24.2-alpine3.21 AS base
WORKDIR /usr/src/app
RUN --mount=type=cache,target=/var/cache/apk \
    ln -vs /var/cache/apk /etc/apk/cache && \
    apk add \
      git~=2

FROM base AS dependencies
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

FROM dependencies AS build
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /usr/local/bin

FROM base AS unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go test -v ./...

FROM golangci/golangci-lint:v1.64.6-go1.24.1 AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run ./...

FROM alpine:3.21.3 AS release
ENV TERM=xterm-256color
RUN --mount=type=cache,target=/var/cache/apk \
    ln -vs /var/cache/apk /etc/apk/cache && \
    apk add \
      git~=2
COPY --from=build /usr/local/bin/committed /usr/local/bin

FROM release AS test
WORKDIR /root/repository
RUN git config --global user.email "you@example.com" && \
    git config --global user.name "Your Name" && \
    git init

ENTRYPOINT ["committed"]
