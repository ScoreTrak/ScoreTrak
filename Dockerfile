FROM golang:latest
WORKDIR /go/src/github.com/ScoreTrak/ScoreTrak
COPY pkg/ pkg/
COPY cmd/ cmd/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod tidy
RUN go build -o master cmd/master/main.go
RUN go build -o worker cmd/worker/main.go
RUN go build -o jobs cmd/jobs/main.go
RUN chmod +x master worker jobs

#Set Context Path as ScoreTrak directory