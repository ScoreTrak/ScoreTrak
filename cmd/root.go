package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/spf13/viper"
)

var cfgFile string
var C config.StaticConfig
var D config.DynamicConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scoretrak",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scoretrak.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".scoretrak" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/scoretrak")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".scoretrak")
	}

	// Scoretrak Defaults
	viper.SetDefault("adminUsername", "admin")
	viper.SetDefault("adminPassword", "changeme")
	viper.SetDefault("port", 33333)
	viper.SetDefault("prod", false)
	viper.SetDefault("databaseMaxTimeDriftSeconds", 2)
	viper.SetDefault("dynamicConfigPullSeconds", 5)

	// Database Defaults
	viper.SetDefault("db.use", "cockroach")

	// Cockroach Database Defaults
	viper.SetDefault("db.cockroach.host", "cockroach")
	viper.SetDefault("db.cockroach.port", 26257)
	viper.SetDefault("db.cockroach.username", "root")
	viper.SetDefault("db.cockroach.database", "scoretrak")
	viper.SetDefault("db.cockroach.configureZones", true)
	viper.SetDefault("db.cockroach.defaultZoneConfig.gcTtlseconds", 600)
	viper.SetDefault("db.cockroach.defaultZoneConfig.backpressueRangeSizeMultiplier", 0)

	// Queue Defaults
	viper.SetDefault("queue.use", false)
	viper.SetDefault("queue.nsq.producerNSQD", "nsqd:4150")
	viper.SetDefault("queue.nsq.ignoreAllScoresIfWorkerFails", true)
	viper.SetDefault("queue.nsq.topic", "default")
	viper.SetDefault("queue.nsq.maxInFlight", 200)
	viper.SetDefault("queue.nsq.concurrentHandlers", 200)
	viper.SetDefault("queue.nsq.NSQLookupd", []string{"nsqlookupd:4161"})
	viper.SetDefault("queue.nsq.consumerNSQDPool", []string{})

	// Platform Config
	viper.SetDefault("platform.use", "none")
	viper.SetDefault("platform.docker.name", "scoretrak")
	viper.SetDefault("platform.docker.host", "unix:///var/run/docker.sock")
	viper.SetDefault("platform.docker.network", "default")
	viper.SetDefault("platform.kubernetes.namespace", "default")

	// PubSubConfig
	viper.SetDefault("pubSubConfig.reportForceRefreshSeconds", 60)
	viper.SetDefault("pubSubConfig.channelPrefix", "master")

	// JWT Config
	viper.SetDefault("jwt.secret", "changeme")
	viper.SetDefault("jwt.timeooutInSeconds", 86400)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

		err := viper.Unmarshal(&C)
		if err != nil {
			log.Fatalf("unable to decode static config into struct, %v", err)
		}

		if C.Prod {
			log.SetFlags((log.LstdFlags | log.Lshortfile))
		}

		err = viper.Unmarshal(&D)
		if err != nil {
			log.Fatalf("unable to decode dynamic config into struct, %v", err)
		}
	}
}
