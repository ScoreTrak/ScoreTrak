package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"os"
)

func main() {
	path := flag.String("config", "configs/config.yml", "Please enter a path to config file")
	encodedConfig := flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	if *encodedConfig != "" {
		dec, err := base64.StdEncoding.DecodeString(*encodedConfig)
		handleErr(err)
		*path = "config.yml"
		f, err := os.Create(*path)
		handleErr(err)
		defer f.Close()
		_, err = f.Write(dec)
		handleErr(err)
		handleErr(f.Sync())
	} else if !configExists(*path) {
		handleErr(errors.New("you need to provide config"))
	}
	handleErr(config.NewStaticConfig(*path))
	l, err := logger.NewLogger(config.GetStaticConfig().Logger)
	handleErr(err)
	q, err := queue.NewQueue(config.GetStaticConfig(), l)
	handleErr(err)
	q.Receive()
}

func configExists(f string) bool {
	file, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else {
		return
	}
}
