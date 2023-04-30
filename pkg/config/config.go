package config

import (
	"github.com/creasty/defaults"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// Config is a struct of settings that are set at the start of the application. It contains Configs from other packages defined under pkg/ directory.
// Config is read only at the moment, hence there is no lock / prevention to race conditions.
type Config struct {
	DB struct {
		Use string `default:"sqlite3"`
		DSN string `default:"file:ent?mode=memory&cache=shared&_fk=1"`
		//Migrate bool   `default:"false"`
		//Seed    bool   `default:"false"`
		// Cockroach struct {
		// 	ConfigureZones    bool `default:"true"`
		// 	DefaultZoneConfig struct {
		// TODO automate this outside of scoretrak
		// 		GcTtlseconds                    uint64 `default:"600"`
		// 		BackpressureRangeSizeMultiplier uint64 `default:"0"`
		// 	}
		// }
	}

	// This value ideally shouldn't be larger than few seconds
	DatabaseMaxTimeDriftSeconds uint `default:"2"`

	// How frequently to pull dynamic configs
	DynamicConfigPullSeconds uint `default:"5"`

	Queue struct {
		Use   string `default:"none"`
		Pool  int    `default:5`
		Kafka struct {
		}
		NSQ struct {
			Worker struct {
				NSQD        string `default:""`
				MaxInFlight int    `default:"200"` // This should be more than min(NumberOfChecks, #NSQD Nodes)
				Topic       string `default:"default"`
				Channel     string `default:""`
			}
			ProducerNSQD                 string   `default:"nsqd:4150"`
			IgnoreAllScoresIfWorkerFails bool     `default:"true"`
			Topic                        string   `default:"default"`
			MaxInFlight                  int      `default:"200"` // This should be more than min(NumberOfChecks, #NSQD Nodes)
			AuthSecret                   string   `default:""`
			ClientRootCA                 string   `default:""`
			ClientSSLKey                 string   `default:""`
			ClientSSLCert                string   `default:""`
			ConcurrentHandlers           int      `default:"200"`
			NSQLookupd                   []string `default:"[\"\"]"` // "[\"nsqlookupd:4160\"]"
			ConsumerNSQDPool             []string `default:"[\"\"]"` // "[\"nsqd:4150\"]"
		}
	}

	//Platform struct {
	//	Use    string `default:"none"`
	//	Docker struct {
	//		Name    string `default:"scoretrak"`
	//		Host    string `default:"unix:///var/run/docker.sock"`
	//		Network string `default:"default"`
	//	}
	//	Kubernetes struct {
	//		Namespace string `default:"default"`
	//	}
	//}

	PubSubConfig struct {
		ReportForceRefreshSeconds uint   `default:"60"`
		ChannelPrefix             string `default:"master"`
	}

	Server struct {
		Address string `default:"127.0.0.1"`
		Port    string `default:"3000"`
		TLS     struct {
			CertFile string
			KeyFile  string
		}
		Cors cors.Options
	}

	Prod bool `default:"false"`

	Auth struct {
		JWT struct {
			Secret           string `default:"changeme"`
			TimeoutInSeconds uint64 `default:"86400"`
		}
		Ory struct {
			SelfHosted  bool
			Slug        string
			CookieName  string `default:"ory_kratos_session"`
			AdminApiUrl string `default:"http://localhost:4433"`
		}
	}

	Debug bool `default:"false"`
}

const (
	ENV_PREFIX = "ST"
)

func NewScoreTrakConfig() (*Config, error) {
	c := &Config{}

	if err := defaults.Set(c); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, nil
}

func NewCorsConfig(c *Config) *cors.Cors {
	crsCfg := c.Server.Cors
	crsCfg.Debug = c.Debug
	crs := cors.New(crsCfg)
	return crs
}
