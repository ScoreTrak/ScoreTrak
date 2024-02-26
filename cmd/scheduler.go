package cmd

import (
	"github.com/scoretrak/scoretrak/pkg/eventsv2/eventsv2fx"
	"github.com/scoretrak/scoretrak/pkg/scheduler/schedulerfx"
	"log"

	"github.com/scoretrak/scoretrak/pkg/config/configfx"
	"github.com/scoretrak/scoretrak/pkg/storage/storagefx"
	"github.com/scoretrak/scoretrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// schedulerCmd represents the flagbearer command
var schedulerCmd = &cobra.Command{
	Use: "scheduler",
	//Short: "master runs the grpc server and runner if in single-node mode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting scheduler")

		app := fx.New(
			// Create configs
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Create events components
			eventsv2fx.Module,

			// Create Scheduler
			schedulerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(schedulerCmd)
}
