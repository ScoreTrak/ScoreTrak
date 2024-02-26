package cmd

import (
	"github.com/scoretrak/scoretrak/pkg/auth/authfx"
	"github.com/scoretrak/scoretrak/pkg/config/configfx"
	"github.com/scoretrak/scoretrak/pkg/events/eventsfx"
	"github.com/scoretrak/scoretrak/pkg/server/serverfx"
	"github.com/scoretrak/scoretrak/pkg/storage/storagefx"
	"github.com/scoretrak/scoretrak/pkg/telemetry/telemetryfx"
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

			// Event components
			eventsfx.Module,

			// Server
			serverfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
