package master

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/run"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/queue"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/storage"
	"ScoreTrak/pkg/team"
	"github.com/lib/pq"
)

func Run() error {

	r := server.NewRouter()
	d, err := di.BuildMasterContainer()
	if err != nil {
		return err
	}
	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})
	svr := server.NewServer(r, d, l)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	dc := config.GetConfigCopy()
	db := storage.GetGlobalDB()
	err = db.Create(dc).Error
	if err != nil {
		serr, ok := err.(*pq.Error)
		if ok && serr.Code.Name() == "unique_violation" {
			dcc := config.DynamicConfig{}
			db.Take(&dcc)
			config.UpdateConfig(&dcc)
		} else {
			return err
		}
	}
	err = svr.Start()
	if err != nil {
		return err
	}

	q, err := queue.NewQueue(config.GetStaticConfig())

	if err != nil {
		return err
	}

	var hostGroupRepo host_group.Repo
	di.Invoke(func(re host_group.Repo) {
		hostGroupRepo = re
	})
	var hostRepo host.Repo
	di.Invoke(func(re host.Repo) {
		hostRepo = re
	})
	var roundRepo round.Repo
	di.Invoke(func(re round.Repo) {
		roundRepo = re
	})
	var serviceRepo service.Repo
	di.Invoke(func(re service.Repo) {
		serviceRepo = re
	})
	var serviceGroupRepo service_group.Repo
	di.Invoke(func(re service_group.Repo) {
		serviceGroupRepo = re
	})
	var propertyRepo property.Repo
	di.Invoke(func(re property.Repo) {
		propertyRepo = re
	})
	var checkRepo check.Repo
	di.Invoke(func(re check.Repo) {
		checkRepo = re
	})
	var teamRepo team.Repo
	di.Invoke(func(re team.Repo) {
		teamRepo = re
	})
	var configRepo config.Repo
	di.Invoke(func(re config.Repo) {
		configRepo = re
	})

	repoStore := run.RepoStore{
		Round:        roundRepo,
		HostGroup:    hostGroupRepo,
		Host:         hostRepo,
		Service:      serviceRepo,
		ServiceGroup: serviceGroupRepo,
		Property:     propertyRepo,
		Check:        checkRepo,
		Team:         teamRepo,
		Config:       configRepo,
	}

	dr := run.NewRunner(db, l, q, repoStore)

	return dr.MasterRunner()

}
