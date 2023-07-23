package main

import (
	"github.com/ScoreTrak/ScoreTrak/cmd"
	_ "github.com/ScoreTrak/ScoreTrak/pkg/entities/runtime"
)

func main() {
	cmd.Execute()
}
