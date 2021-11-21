package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pong")
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
