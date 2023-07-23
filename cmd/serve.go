package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/authfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/events/eventsfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/scheduler/schedulerfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/scorer/scorerfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/serverfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"go.uber.org/fx"
	"log"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("hello")
		app := fx.New(
			// Config
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Database components
			storagefx.Module,

			// Auth components
			authfx.Module,

			// Server
			serverfx.Module,

			// Event components
			eventsfx.Module,
			scorerfx.Module, // TODO: REMOVE

			// Cron scheduler
			schedulerfx.Module, // TODO: REMOVE
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
