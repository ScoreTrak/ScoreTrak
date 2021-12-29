FROM golang:1.17
ARG IMAGE_TAG
RUN mkdir -p /go/src/github.com/ScoreTrak/ScoreTrak
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak
COPY pkg/ pkg/
COPY cmd/ cmd/
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN go build -o scoretrak -ldflags "-X 'github.com/Scoretrak/Scoretrak/pkg/version.Version=${IMAGE_TAG}'"
RUN chmod +x scoretrak
ENTRYPOINT ["./scoretrak"]
CMD ["master"]