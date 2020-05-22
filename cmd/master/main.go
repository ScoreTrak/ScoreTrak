package main

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/master"
	"fmt"
	"os"
)

func main() {

	if err := config.NewStaticConfig("configs/config.yml"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	if err := config.NewDynamicConfig("configs/config.yml"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	c := config.GetStaticConfig()
	c.Role = "Master"

	if err := master.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
