package main

import (
	"flag"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master"
	"os"
)

func main() {
	path := flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.Parse()
	if *path != "configs/config.yml" {
		_, err := os.Stat(*path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Provided config was not found")
			os.Exit(-1)
		}
	}
	if err := config.NewStaticConfig(*path); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	fmt.Println(*path)
	if err := master.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start due to: %v", err)
		os.Exit(-1)
	}
}
