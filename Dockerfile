FROM golang:latest
ARG IMAGE_TAG
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak
COPY pkg/ pkg/
COPY cmd/ cmd/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod tidy
RUN go build -o master -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'" cmd/master/main.go
RUN go build -o worker -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'" cmd/worker/main.go
RUN go build -o cli -ldflags "-X 'github.com/ScoreTrak/ScoreTrak/pkg/version.Version=${IMAGE_TAG}'" cmd/cli/main.go
RUN chmod +x master worker

#Set Context Path as ScoreTrak directory