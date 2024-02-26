package main

import (
	"github.com/scoretrak/scoretrak/cmd"
	_ "github.com/scoretrak/scoretrak/internal/entities/runtime"
)

func main() {
	cmd.Execute()
}
