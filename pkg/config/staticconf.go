package config

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
)

// StaticConfig is a struct of settings that are set at the start of the application. It contains Configs from other packages defined under pkg/ directory.
// StaticConfig is read only at the moment, hence there is no lock / prevention to race conditions.
type StaticConfig struct {
	DB storage.Config
	// This value ideally shouldn't be larger than few seconds
	DatabaseMaxTimeDriftSeconds uint `default:"2"`
	// How frequently to pull dynamic configs
	DynamicConfigPullSeconds uint `default:"5"`

	Queue queueing.Config

	Platform platforming.Config

	PubSubConfig queueing.MasterConfig

	AdminUsername string `default:"admin"`

	AdminPassword string `default:"changeme"`

	Port string `default:"33333"`

	Prod bool `default:"false"`

	CertFile string `default:""`

	KeyFile string `default:""`

	JWT auth.Config
}
