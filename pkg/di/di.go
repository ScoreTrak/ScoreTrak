package di

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/serviceservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"

	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildMasterContainer(c config.StaticConfig) (*dig.Container, error) {
	var ctr []interface{}

	ctr = append(ctr,
		func() config.StaticConfig {
			return c
		}, func() storage.Config {
			return c.DB
		}, func() queueing.Config {
			return c.Queue
		}, func() platforming.Config {
			return c.Platform
		},
		storage.LoadDB,
		orm.NewCheckRepo, checkservice.NewCheckServ,
		orm.NewHostGroupRepo, hostgroupservice.NewHostGroupServ,
		orm.NewPropertyRepo, propertyservice.NewPropertyServ,
		orm.NewConfigRepo, configservice.NewConfigServ, configservice.NewStaticConfigServ,
		orm.NewHostRepo, hostservice.NewHostServ,
		orm.NewRoundRepo, roundservice.NewRoundServ,
		orm.NewServiceGroupRepo, servicegroupservice.NewServiceGroupServ,
		orm.NewServiceRepo, serviceservice.NewServiceServ,
		orm.NewTeamRepo, teamservice.NewTeamServ,
		orm.NewUserRepo, userservice.NewUserServ,
		orm.NewPolicyRepo, policyservice.NewPolicyServ,
		orm.NewReportRepo, reportservice.NewReportServ,
		competitionservice.NewCompetitionServ,
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
	if err := container.Invoke(i); err != nil {
		panic(err)
	}
}
