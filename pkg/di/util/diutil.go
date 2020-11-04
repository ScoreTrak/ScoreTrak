package util

import (
	repo5 "github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	repo6 "github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	repo7 "github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"
	repo8 "github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"
	repo9 "github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"
	repo10 "github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	repo3 "github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	repo4 "github.com/ScoreTrak/ScoreTrak/pkg/team/team_repo"
)

func NewStore() *util.Store {
	var hostGroupRepo repo8.Repo
	di.Invoke(func(re repo8.Repo) {
		hostGroupRepo = re
	})
	var hostRepo repo7.Repo
	di.Invoke(func(re repo7.Repo) {
		hostRepo = re
	})
	var roundRepo repo3.Repo
	di.Invoke(func(re repo3.Repo) {
		roundRepo = re
	})
	var serviceRepo repo2.Repo
	di.Invoke(func(re repo2.Repo) {
		serviceRepo = re
	})
	var serviceGroupRepo service_group_repo.Repo
	di.Invoke(func(re service_group_repo.Repo) {
		serviceGroupRepo = re
	})
	var propertyRepo repo9.Repo
	di.Invoke(func(re repo9.Repo) {
		propertyRepo = re
	})
	var checkRepo repo5.Repo
	di.Invoke(func(re repo5.Repo) {
		checkRepo = re
	})
	var teamRepo repo4.Repo
	di.Invoke(func(re repo4.Repo) {
		teamRepo = re
	})
	var configRepo repo6.Repo
	di.Invoke(func(re repo6.Repo) {
		configRepo = re
	})
	var reportRepo repo10.Repo
	di.Invoke(func(re repo10.Repo) {
		reportRepo = re
	})

	return &util.Store{
		Round:        roundRepo,
		HostGroup:    hostGroupRepo,
		Host:         hostRepo,
		Service:      serviceRepo,
		ServiceGroup: serviceGroupRepo,
		Property:     propertyRepo,
		Check:        checkRepo,
		Team:         teamRepo,
		Config:       configRepo,
		Report:       reportRepo,
	}
}
