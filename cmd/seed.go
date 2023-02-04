package cmd

import (
	"context"
	seed2 "github.com/ScoreTrak/ScoreTrak/pkg/storage/seed"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed the database",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		seed := fx.New(
			fx.Provide(
				NewStaticConfig,
				NewStorageConfig,
			),

			storagefx.Module,

			fx.Invoke(seed2.DefaultSeed),
		)

		err := seed.Start(ctx)
		if err != nil {
			log.Fatalf("Command failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
