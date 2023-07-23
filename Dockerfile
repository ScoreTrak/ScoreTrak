FROM golang:1.20 as builder
ARG IMAGE_TAG
RUN mkdir -p /go/src/github.com/ScoreTrak/ScoreTrak
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak
COPY . .
RUN go mod download
RUN go build -o scoretrak -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'"
RUN chmod +x scoretrak


FROM debian:stable-slim

COPY --from=builder \
    /go/src/github.com/ScoreTrak/ScoreTrak/scoretrak \
    /bin/scoretrak

ENTRYPOINT ["/bin/scoretrak"]
