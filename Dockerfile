FROM golang:1.19.4-alpine3.17 as base
WORKDIR /usr/src/app
RUN --mount=type=cache,target=/var/cache/apk \
    ln -vs /var/cache/apk /etc/apk/cache && \
    apk add \
      git=2.38.2-r0

FROM base as dependencies
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

FROM dependencies as build
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /usr/local/bin

FROM base as unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go test -v ./...

FROM golangci/golangci-lint:v1.50.1-alpine AS lint-base

FROM base as lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run ./...

FROM alpine:3.17.0 as release
ENV TERM=xterm-256color
RUN --mount=type=cache,target=/var/cache/apk \
    ln -vs /var/cache/apk /etc/apk/cache && \
    apk add \
      git=2.38.2-r0
COPY --from=build /usr/local/bin/committed /usr/local/bin

FROM release as test
WORKDIR /root/repository
RUN git config --global user.email "you@example.com" && \
    git config --global user.name "Your Name" && \
    git init

ENTRYPOINT ["committed"]
