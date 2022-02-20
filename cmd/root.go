package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
)

var cfgFile string
var encodedCfg string
var C config.StaticConfig
var D config.DynamicConfig

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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is './.scoretrak.yaml', '$HOME/.scoretrak.yaml', '/etc/scoretrak/.scoretrak.yaml' in that order)")
	rootCmd.PersistentFlags().StringVar(&encodedCfg, "encoded-config", "", "base64 encoded config")

	rootCmd.Flags().BoolP("version", "v", false, "version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".scoretrak" (without extension).
	viper.AddConfigPath(".")
	viper.AddConfigPath(home)
	viper.AddConfigPath("/etc/scoretrak")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".scoretrak")

	// Use config file flag value.
	viper.SetConfigFile(cfgFile)

	// Scoretrak Static Defaults
	viper.SetDefault("AdminUsername", "admin")
	viper.SetDefault("AdminPassword", "changeme")
	viper.SetDefault("Port", 33333)
	viper.SetDefault("Prod", false)
	viper.SetDefault("DatabaseMaxTimeDriftSeconds", 2)
	viper.SetDefault("DynamicConfigPullSeconds", 5)

	// Scoretrak Dynamic Defaults
	viper.SetDefault("Enabled", "false")
	viper.SetDefault("RoundDuration", 60)

	// Database Defaults
	viper.SetDefault("DB.Use", "cockroach")
	// viper.SetDefault("DB.Prefix", "st")

	// Cockroach Database Defaults
	viper.SetDefault("DB.Cockroach.Host", "cockroach")
	viper.SetDefault("DB.Cockroach.Port", 26257)
	viper.SetDefault("DB.Cockroach.Username", "root")
	viper.SetDefault("DB.Cockroach.Database", "scoretrak")
	viper.SetDefault("DB.Cockroach.ConfigureZones", true)
	viper.SetDefault("DB.Cockroach.DefaultZoneConfig.GcTtlseconds", 600)
	viper.SetDefault("DB.Cockroach.DefaultZoneConfig.BackpressueRangeSizeMultiplier", 0)

	// Queue Defaults
	viper.SetDefault("Queue.Use", "none")
	viper.SetDefault("Queue.NSQ.ProducerNSQD", "nsqd:4150")
	viper.SetDefault("Queue.NSQ.IgnoreAllScoresIfWorkerFails", true)
	viper.SetDefault("Queue.NSQ.Topic", "default")
	viper.SetDefault("Queue.NSQ.MaxInFlight", 200)
	viper.SetDefault("Queue.NSQ.ConcurrentHandlers", 200)
	viper.SetDefault("Queue.NSQ.NSQLookupd", []string{"nsqlookupd:4161"})
	viper.SetDefault("Queue.NSQ.ConsumerNSQDPool", []string{})

	// Platform Config
	viper.SetDefault("Platform.Use", "none")
	viper.SetDefault("Platform.Docker.Name", "scoretrak")
	viper.SetDefault("Platform.Docker.Host", "unix:///var/run/docker.sock")
	viper.SetDefault("Platform.Docker.Network", "default")
	viper.SetDefault("Platform.Kubernetes.Namespace", "default")

	// PubSubConfig
	viper.SetDefault("PubSubConfig.ReportForceRefreshSeconds", 60)
	viper.SetDefault("PubSubConfig.ChannelPrefix", "master")

	// JWT Config
	viper.SetDefault("JWT.Secret", "changeme")
	viper.SetDefault("JWT.TimeoutInSeconds", 86400)

	viper.AutomaticEnv() // read in environment variables that match

	// If an encodedConfig flag is provided, read it. Else read found config file
	if encodedCfg != "" {
		log.Printf("Using encoded config file: %s", encodedCfg)
		decodedCfg, err := base64.StdEncoding.DecodeString(encodedCfg)
		if err != nil {
			log.Printf("Error decoding string: %s ", err.Error())
			return
		}
		err = viper.ReadConfig(bytes.NewBuffer(decodedCfg))
		if err != nil {
			log.Printf("Error reading decoded config %s", err.Error())
		}
	} else if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			log.Fatalf("unable to print to standard error, %v", err)
		}
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("unable to decode static config into struct, %v", err)
	}
	config.SetStaticConfig(C)

	if C.Prod {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	err = viper.Unmarshal(&D)
	if err != nil {
		log.Fatalf("unable to decode dynamic config into struct, %v", err)
	}
}
