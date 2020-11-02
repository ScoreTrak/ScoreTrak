package di

import (
	service5 "github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	service11 "github.com/ScoreTrak/ScoreTrak/pkg/competition/competition_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	service6 "github.com/ScoreTrak/ScoreTrak/pkg/config/config_service"
	service7 "github.com/ScoreTrak/ScoreTrak/pkg/host/host_service"
	service8 "github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_service"
	service13 "github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_service"
	service9 "github.com/ScoreTrak/ScoreTrak/pkg/property/property_service"
	service10 "github.com/ScoreTrak/ScoreTrak/pkg/report/report_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_service"
	service3 "github.com/ScoreTrak/ScoreTrak/pkg/service/service_service"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_service"
	service4 "github.com/ScoreTrak/ScoreTrak/pkg/team/team_service"
	service12 "github.com/ScoreTrak/ScoreTrak/pkg/user/user_service"

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
		orm.NewRoundRepo, round_service.NewRoundServ,
		orm.NewServiceGroupRepo, service2.NewServiceGroupServ,
		orm.NewServiceRepo, service3.NewServiceServ,
		orm.NewTeamRepo, service4.NewTeamServ,
		orm.NewUserRepo, service12.NewUserServ,
		orm.NewPolicyRepo, service13.NewPolicyServ,

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
