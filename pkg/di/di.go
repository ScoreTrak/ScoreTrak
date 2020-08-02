package di

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildMasterContainer() (*dig.Container, error) {
	var ctr []interface{}

	ctr = append(ctr,
		config.GetStaticConfig, config.GetLoggerConfig, config.GetDBConfig, config.GetQueueConfig, config.GetPlatformConfig,
		storage.LoadDB,
		logger.NewLogger,
		orm.NewCheckRepo, check.NewCheckServ,
		orm.NewHostGroupRepo, host_group.NewHostGroupServ,
		orm.NewPropertyRepo, property.NewPropertyServ,
		orm.NewConfigRepo, config.NewConfigServ, config.NewStaticConfigServ,
		orm.NewHostRepo, host.NewHostServ,
		orm.NewRoundRepo, round.NewRoundServ,
		orm.NewServiceGroupRepo, service_group.NewServiceGroupServ,
		orm.NewServiceRepo, service.NewServiceServ,
		orm.NewTeamRepo, team.NewTeamServ,
		orm.NewReportRepo, report.NewReportServ,
		queue.NewQueue, platform.NewPlatform,
	)

	for _, i := range ctr {
		err := container.Provide(i)
		if err != nil {
			return nil, err
		}
	}
	return container, nil
}

func BuildWorkerContainer() (*dig.Container, error) {
	var ctr []interface{}
	ctr = append(ctr, config.GetStaticConfig, config.GetLoggerConfig)

	for _, i := range ctr {
		err := container.Provide(i)
		if err != nil {
			return nil, err
		}
	}
	return container, nil
}

func Invoke(i interface{}) {
	err := container.Invoke(i)
	if err != nil {
		panic(err)
	}
}
