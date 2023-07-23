package config

import (
	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

// Config is a struct of settings that are set at the start of the application. It contains Configs from other packages defined under pkg/ directory.
// Config is read only at the moment, hence there is no lock / prevention to race conditions.
type Config struct {
	DB struct {
		Use string `default:"sqlite3"`
		DSN string `default:"file:ent?mode=memory&cache=shared&_fk=1"`
	}

	Queue struct {
		Use       string `default:"gochannel"`
		Pool      int    `default:"5"`
		GoChannel struct {
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
		Jobs    struct {
			RoundStarter struct {
				CronSpec string `default:"0 * * * * *"`
				Timeout  int    `default:"55"`
			}
			RoundFinisher struct {
				CronSpec string `default:"5 * * * * *"`
				Timeout  int    `default:"15"`
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

	Dev bool `default:"true"` // Is in Dev mode
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
