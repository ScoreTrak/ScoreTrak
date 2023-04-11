package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/worker/workerfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker to perform checks on systems",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			// Create configs
			fx.Provide(
				fx.Annotate(
					func() string {
						return cfgFile
					},
					fx.ResultTags(`name:"cfgFile"`),
				),
			),
			configfx.Module,
			telemetryfx.Module,

			fx.Provide(
				queue.NewWorkerQueue,
			),

			workerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
