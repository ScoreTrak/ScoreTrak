package cmd

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		migrate := fx.New(
			fx.Provide(
				NewStaticConfig,
				NewDynamicConfig,
				NewStorageConfig,
			),

			storagefx.Module,

			fx.Invoke(util.AutoMigrate),
		)

		err := migrate.Start(ctx)
		if err != nil {
			log.Fatalf("Migrate failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
