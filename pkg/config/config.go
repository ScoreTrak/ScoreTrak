package config

import (
	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

// Config is a struct of settings that are set at the start of the application. It contains Configs from other packages defined under pkg/ directory.
// Config is read only at the moment, hence there is no lock / prevention to race conditions.
type Config struct {
	DB struct {
		Use string `default:"sqlite3" json:"use"`
		DSN string `default:"file:ent?mode=memory&cache=shared&_fk=1" json:"dsn"`
	}

	Queue struct {
		Use       string `default:"nats"`
		Pool      int    `default:"5"`
		GoChannel struct {
			OutputChannelBuffer            int64 `default:"100000"`
			Persistent                     bool  `default:"true"`
			BlockPublishUntilSubscriberAck bool  `default: "false"`
		}
		NATS struct {
			Url              string `default:"nats://localhost:4222"`
			SubscriberCount  int    `default:"1"`
			QueueGroupPrefix string `default:"scorer"`
			Jetstream        struct {
				Disabled      bool `default:"false"`
				AutoProvision bool `default:"true"`
			}
		}
	}

	Scheduler struct {
		Enabled bool `default:"true"`
		Elector struct {
			Enabled     bool     `json:"enabled" default:"false"`
			Endpoints   []string `default:"\['127.0.0.1:2379'\]"`
			DialTimeout int
			Key         string `default:"/scoretrak/leader"`
		}
		Jobs struct {
			Ping struct {
				FrequencySecond int `default:"1"`
			}
			RoundStarter struct {
				FrequencySecond int `default:"60"`
				Timeout         int `default:"10"`
			}
			RoundFinisher struct {
				FrequencySecond int `default:"3"`
				Timeout         int `default:"10"`
			}
		}
	}

	Scorer struct {
		Enabled        bool `default:"true"`
		ScoringTimeout int  `default:"30"`
	}

	Server struct {
		Enabled bool   `default:"true"`
		Address string `default:"127.0.0.1"`
		Port    string `default:"3000"`
		TLS     struct {
			CertFile string
			KeyFile  string
		}
		Cors struct {
			Enabled              bool `default:"false"`
			AllowedOrigins       []string
			AllowedMethods       []string
			AllowedHeaders       []string
			ExposedHeaders       []string
			MaxAge               int
			AllowCredentials     bool
			AllowPrivateNetwork  bool
			OptionsPassthrough   bool
			OptionsSuccessStatus int
			Debug                bool
		}
	}

	Auth struct {
		Ory struct {
			SelfHosted  bool
			Slug        string
			CookieName  string `default:"ory_kratos_session"`
			AdminApiUrl string `default:"http://localhost:4433"`
		}
	}

	Log struct {
		Level string `default:"debug"`
	}
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
