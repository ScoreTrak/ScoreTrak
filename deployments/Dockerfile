FROM golang:latest
WORKDIR /go/src/github.com/L1ghtman2k/ScoreTrak
COPY . .
RUN go mod tidy
RUN go build -o master cmd/master/main.go
RUN go build -o worker cmd/worker/main.go
RUN chmod +x master worker
RUN mv master /tmp/master
RUN mv worker /tmp/worker
RUN rm -rf *
RUN mv /tmp/master master
RUN mv /tmp/worker worker

#Set Context Path as ScoreTrak directory