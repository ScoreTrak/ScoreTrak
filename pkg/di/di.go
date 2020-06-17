package di

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/queue"
	"ScoreTrak/pkg/report"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/storage"
	"ScoreTrak/pkg/storage/orm"
	"ScoreTrak/pkg/swarm"
	"ScoreTrak/pkg/team"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildMasterContainer() (*dig.Container, error) {
	cnf := config.GetStaticConfig()
	var ctr []interface{}

	ctr = append(ctr,
		config.GetStaticConfig,
		storage.LoadDB,
		logger.NewLogger,
		orm.NewCheckRepo, check.NewCheckServ,
		orm.NewHostGroupRepo, host_group.NewHostGroupServ,
		orm.NewPropertyRepo, property.NewPropertyServ,
		orm.NewConfigRepo, config.NewConfigServ,
		orm.NewHostRepo, host.NewHostServ,
		orm.NewRoundRepo, round.NewRoundServ,
		orm.NewServiceGroupRepo, service_group.NewServiceGroupServ,
		orm.NewServiceRepo, service.NewServiceServ,
		orm.NewTeamRepo, team.NewTeamServ,
		report.NewReportServ,
		queue.NewQueue,
	)

	if cnf.Platform == "swarm" {
		ctr = append(ctr, orm.NewSwarmRepo)
		ctr = append(ctr, swarm.NewSwarmServ)
	}

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
	ctr = append(ctr, config.GetStaticConfig)

	for _, i := range ctr {
		err := container.Provide(i)
		if err != nil {
			return nil, err
		}
	}
	return container, nil
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
