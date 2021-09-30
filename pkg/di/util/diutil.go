package util

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicerepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegrouprepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamrepo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
)

func NewStore() *util.Store {
	var hostGroupRepo hostgrouprepo.Repo
	di.Invoke(func(re hostgrouprepo.Repo) {
		hostGroupRepo = re
	})
	var hostRepo hostrepo.Repo
	di.Invoke(func(re hostrepo.Repo) {
		hostRepo = re
	})
	var roundRepo roundrepo.Repo
	di.Invoke(func(re roundrepo.Repo) {
		roundRepo = re
	})
	var serviceRepo servicerepo.Repo
	di.Invoke(func(re servicerepo.Repo) {
		serviceRepo = re
	})
	var serviceGroupRepo servicegrouprepo.Repo
	di.Invoke(func(re servicegrouprepo.Repo) {
		serviceGroupRepo = re
	})
	var propertyRepo propertyrepo.Repo
	di.Invoke(func(re propertyrepo.Repo) {
		propertyRepo = re
	})
	var checkRepo checkrepo.Repo
	di.Invoke(func(re checkrepo.Repo) {
		checkRepo = re
	})
	var teamRepo teamrepo.Repo
	di.Invoke(func(re teamrepo.Repo) {
		teamRepo = re
	})
	var configRepo configrepo.Repo
	di.Invoke(func(re configrepo.Repo) {
		configRepo = re
	})
	var reportRepo reportrepo.Repo
	di.Invoke(func(re reportrepo.Repo) {
		reportRepo = re
	})

	var policyRepo policyrepo.Repo
	di.Invoke(func(re policyrepo.Repo) {
		policyRepo = re
	})

	var userRepo userrepo.Repo
	di.Invoke(func(re userrepo.Repo) {
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
