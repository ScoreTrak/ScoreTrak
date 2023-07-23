package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/events/eventsfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/scheduler/schedulerfx"
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// flagbearerCmd represents the flagbearer command
var flagbearerCmd = &cobra.Command{
	Use: "flagbearer",
	//Short: "master runs the grpc server and runner if in single-node mode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting flagbearer")

		app := fx.New(
			// Create configs
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Create queueing components
			eventsfx.Module,

			// Create flagbearer
			schedulerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(flagbearerCmd)
}
