package di

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competition_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/config_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/service_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/team_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_service"

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
		orm.NewCheckRepo, check_service.NewCheckServ,
		orm.NewHostGroupRepo, host_group_service.NewHostGroupServ,
		orm.NewPropertyRepo, property_service.NewPropertyServ,
		orm.NewConfigRepo, config_service.NewConfigServ, config_service.NewStaticConfigServ,
		orm.NewHostRepo, host_service.NewHostServ,
		orm.NewRoundRepo, round_service.NewRoundServ,
		orm.NewServiceGroupRepo, service_group_service.NewServiceGroupServ,
		orm.NewServiceRepo, service_service.NewServiceServ,
		orm.NewTeamRepo, team_service.NewTeamServ,
		orm.NewUserRepo, user_service.NewUserServ,
		orm.NewPolicyRepo, policy_service.NewPolicyServ,
		orm.NewReportRepo, report_service.NewReportServ,
		competition_service.NewCompetitionServ,
		queue.NewWorkerQueue, platform.NewPlatform,
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
