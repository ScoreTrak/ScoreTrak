FROM golang:alpine as builder
ARG IMAGE_TAG

RUN mkdir -p /go/src/github.com/ScoreTrak/ScoreTrak
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak

COPY pkg/ pkg/
COPY cmd/ cmd/
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

ENV CGO_ENABLED=0
RUN go mod download && go build -o scoretrak -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'"

FROM alpine:3.14.3
RUN mkdir -p /scoretrak
WORKDIR /scoretrak
COPY --from=builder \
    /go/src/github.com/ScoreTrak/ScoreTrak/scoretrak \
    /scoretrak/scoretrak

ENTRYPOINT ["/scoretrak/scoretrak"]
CMD ["master"]