package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/worker"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker to perform checks on systems",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			fx.Provide(NewStaticConfig, NewQueueConfig),
			telemetryfx.Module,

			fx.Provide(
				queue.NewWorkerQueue,
			),

			fx.Invoke(worker.InitWorker),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
