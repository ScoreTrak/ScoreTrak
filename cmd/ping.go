package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("pong")
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
