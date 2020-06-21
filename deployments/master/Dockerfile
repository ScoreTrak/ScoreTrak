FROM golang:latest
WORKDIR /go/src/github.com/L1ghtman2k/ScoreTrak
COPY . .
RUN go mod tidy
RUN scripts/init_db.sh
ENTRYPOINT [ "go", "run", "./cmd/master/main.go" ]
