package di

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
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

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)
	// DB
	container.Provide(storage.NewDb)
	// logger
	container.Provide(logger.NewLogger)

	container.Provide(orm.NewCheckRepo)
	container.Provide(check.NewCheckServ)

	container.Provide(orm.NewHostGroupRepo)
	container.Provide(host_group.NewHostGroupServ)

	container.Provide(orm.NewPropertyRepo)
	container.Provide(property.NewPropertyServ)

	container.Provide(orm.NewConfigRepo)
	container.Provide(config.NewConfigServ)

	container.Provide(orm.NewHostRepo)
	container.Provide(host.NewHostServ)

	container.Provide(orm.NewRoundRepo)
	container.Provide(round.NewRoundServ)

	container.Provide(orm.NewServiceGroupRepo)
	container.Provide(service_group.NewServiceGroupServ)

	container.Provide(orm.NewServiceRepo)
	container.Provide(service.NewServiceServ)

	container.Provide(orm.NewSwarmRepo)
	container.Provide(swarm.NewSwarmServ)

	container.Provide(orm.NewTeamRepo)
	container.Provide(team.NewTeamServ)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
