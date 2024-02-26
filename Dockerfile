# syntax=docker/dockerfile:1

ARG GO_CACHE_DIR=~/go/pkg/mod/
ARG GOLANGCI_LINT_VERSION=v1.52
ARG GO_VERSION=1.21
ARG APP_VERSION=v0.0.0-unknown

FROM golang:${GO_VERSION} as base
LABEL org.opencontainers.image.title="ScoreTrak"
LABEL org.opencontainers.image.source="github.com/scoretrak/scoretrak"
LABEL org.opencontainers.image.version="$APP_VERSION"
LABEL org.opencontainers.image.description="Scoring engine for security competitions"
WORKDIR /src
RUN --mount=type=cache,target=${GO_CACHE_DIR} \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

FROM golangci/golangci-lint:${GOLANGCI_LINT_VERSION} as lint
WORKDIR /src
RUN --mount=type=bind,target=. \
    golangci-lint run

FROM base AS test
RUN --mount=type=cache,target=${GO_CACHE_DIR} \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go install github.com/onsi/ginkgo/v2/ginkgo

FROM test AS test-now
# ginkgo requires a read/write bind mount as it writes temporary test files.
RUN --mount=type=cache,target=${GO_CACHE_DIR} \
    --mount=type=bind,rw,target=. \
    ginkgo -r --procs=2 --compilers=4 --randomize-all --randomize-suites --fail-on-pending --keep-going --cover --coverprofile=cover.profile --race --trace --json-report=report.json --timeout=10m --poll-progress-after=120s --poll-progress-interval=30s

FROM base AS build
RUN --mount=type=cache,target=${GO_CACHE_DIR} \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/scoretrak -ldflags "-X 'github.com/scoretrak/scoretrak/pkg/version.Version=${APP_VERSION}'"

FROM scratch as scoretrak
COPY --from=build \
    /bin/scoretrak \
    /bin/scoretrak
ENTRYPOINT ["/bin/scoretrak"]
