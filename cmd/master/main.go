package main

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master"
	"os"
)

func main() {

	if err := config.NewStaticConfig("configs/config.yml"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	if err := master.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start due to: %v", err)
		os.Exit(-1)
	}
}
