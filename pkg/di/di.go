package di

import (
	service5 "github.com/ScoreTrak/ScoreTrak/pkg/check/service"
	service11 "github.com/ScoreTrak/ScoreTrak/pkg/competition/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	service6 "github.com/ScoreTrak/ScoreTrak/pkg/config/service"
	service7 "github.com/ScoreTrak/ScoreTrak/pkg/host/service"
	service8 "github.com/ScoreTrak/ScoreTrak/pkg/host_group/service"
	service9 "github.com/ScoreTrak/ScoreTrak/pkg/property/service"
	service10 "github.com/ScoreTrak/ScoreTrak/pkg/report/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/service"
	service3 "github.com/ScoreTrak/ScoreTrak/pkg/service/service"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service"
	service4 "github.com/ScoreTrak/ScoreTrak/pkg/team/service"

	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildMasterContainer() (*dig.Container, error) {
	var ctr []interface{}

	ctr = append(ctr,
		config.GetStaticConfig, config.GetDBConfig, config.GetQueueConfig, config.GetPlatformConfig,
		storage.LoadDB,
		orm.NewCheckRepo, service5.NewCheckServ,
		orm.NewHostGroupRepo, service8.NewHostGroupServ,
		orm.NewPropertyRepo, service9.NewPropertyServ,
		orm.NewConfigRepo, service6.NewConfigServ, service6.NewStaticConfigServ,
		orm.NewHostRepo, service7.NewHostServ,
		orm.NewRoundRepo, service.NewRoundServ,
		orm.NewServiceGroupRepo, service2.NewServiceGroupServ,
		orm.NewServiceRepo, service3.NewServiceServ,
		orm.NewTeamRepo, service4.NewTeamServ,
		orm.NewReportRepo, service10.NewReportServ,
		service11.NewCompetitionServ,
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
	ctr = append(ctr, config.GetStaticConfig)

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
