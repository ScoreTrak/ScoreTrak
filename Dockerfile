FROM golang:1.17 as builder
ARG IMAGE_TAG
RUN mkdir -p /go/src/github.com/ScoreTrak/ScoreTrak
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak
COPY pkg/ pkg/
COPY cmd/ cmd/
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN go build -o scoretrak -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'"
RUN chmod +x scoretrak


FROM debian:bullseye-slim

COPY --from=builder \
    /go/src/github.com/ScoreTrak/ScoreTrak/scoretrak \
    /go/bin/scoretrak

# Setup grpc-health-probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.8 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENTRYPOINT ["/go/bin/scoretrak"]
CMD ["master"]