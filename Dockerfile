FROM docker.io/library/golang:1.21 AS builder

COPY example-plugin /tmp/example-plugin
WORKDIR /tmp/example-plugin
RUN go build -o example-plugin

FROM docker.io/kong/kong-gateway:3.4.1.0

COPY --from=builder /tmp/example-plugin/example-plugin /usr/local/bin/example-plugin
