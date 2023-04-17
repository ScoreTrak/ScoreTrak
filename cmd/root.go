package cmd

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/creasty/defaults"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"

	"github.com/ScoreTrak/ScoreTrak/pkg/version"
)

var cfgFile string
var c config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scoretrak",
	Short: "Scoretrak, a cyber defense scoring engine!",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("scoretrak %s", version.Version) //nolint:forbidigo // Keeping fmt as log will add undesired timestamp
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is './.scoretrak.yaml', '$HOME/.scoretrak.yaml', '/etc/scoretrak/.scoretrak.yaml' in that order)")

	rootCmd.Flags().BoolP("version", "v", false, "version")
}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".scoretrak")
	}

	viper.SetEnvPrefix(config.ENV_PREFIX)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := defaults.Set(&c); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
}
