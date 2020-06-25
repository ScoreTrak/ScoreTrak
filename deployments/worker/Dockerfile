FROM golang:latest
WORKDIR /go/src/github.com/L1ghtman2k/ScoreTrak
COPY . .
RUN go mod tidy
ENTRYPOINT [ "go", "run", "./cmd/worker/main.go" ]
