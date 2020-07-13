package main

import (
	"encoding/base64"
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
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
		*path = "config.yml"
		f, err := os.Create(*path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
		if err := f.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	} else if !configExists(*path) {
		fmt.Fprintf(os.Stderr, "You need to provide config!")
		os.Exit(-1)
	}

	if err := config.NewStaticConfig(*path); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	l, err := logger.NewLogger(config.GetStaticConfig().Logger)
	if err != nil {
		panic(err)
	}
	q, err := queue.NewQueue(config.GetStaticConfig(), l)
	if err != nil {
		panic(err)
	}
	q.Receive()
}

func configExists(f string) bool {
	file, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}
