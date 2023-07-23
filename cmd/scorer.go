package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/events/eventsfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/scorer/scorerfx"
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// scorerCmd represents the scorer command
var scorerCmd = &cobra.Command{
	Use: "scorer",
	//Short: "master runs the grpc server and runner if in single-node mode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting scorer")

		app := fx.New(
			// Create configs
			configfx.Module,

			// Create telemetry components
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Create queueing components
			eventsfx.Module,

			// Create scorer components
			scorerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(scorerCmd)
}
