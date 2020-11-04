package util

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/team_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_repo"
)

func NewStore() *util.Store {
	var hostGroupRepo host_group_repo.Repo
	di.Invoke(func(re host_group_repo.Repo) {
		hostGroupRepo = re
	})
	var hostRepo host_repo.Repo
	di.Invoke(func(re host_repo.Repo) {
		hostRepo = re
	})
	var roundRepo round_repo.Repo
	di.Invoke(func(re round_repo.Repo) {
		roundRepo = re
	})
	var serviceRepo service_repo.Repo
	di.Invoke(func(re service_repo.Repo) {
		serviceRepo = re
	})
	var serviceGroupRepo service_group_repo.Repo
	di.Invoke(func(re service_group_repo.Repo) {
		serviceGroupRepo = re
	})
	var propertyRepo property_repo.Repo
	di.Invoke(func(re property_repo.Repo) {
		propertyRepo = re
	})
	var checkRepo check_repo.Repo
	di.Invoke(func(re check_repo.Repo) {
		checkRepo = re
	})
	var teamRepo team_repo.Repo
	di.Invoke(func(re team_repo.Repo) {
		teamRepo = re
	})
	var configRepo config_repo.Repo
	di.Invoke(func(re config_repo.Repo) {
		configRepo = re
	})
	var reportRepo report_repo.Repo
	di.Invoke(func(re report_repo.Repo) {
		reportRepo = re
	})

	var policyRepo policy_repo.Repo
	di.Invoke(func(re policy_repo.Repo) {
		policyRepo = re
	})

	var userRepo user_repo.Repo
	di.Invoke(func(re user_repo.Repo) {
		userRepo = re
	})

	return &util.Store{
		Round:        roundRepo,
		Host:         hostRepo,
		HostGroup:    hostGroupRepo,
		Service:      serviceRepo,
		ServiceGroup: serviceGroupRepo,
		Team:         teamRepo,
		Check:        checkRepo,
		Property:     propertyRepo,
		Config:       configRepo,
		Report:       reportRepo,
		Policy:       policyRepo,
		Users:        userRepo,
	}
}
