package util

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/di"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
)

type RepoStore struct {
	Round        round.Repo
	Host         host.Repo
	HostGroup    host_group.Repo
	Service      service.Repo
	ServiceGroup service_group.Repo
	Team         team.Repo
	Check        check.Repo
	Property     property.Repo
	Config       config.Repo
}

func NewRepoStore() RepoStore {
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

	return RepoStore{
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
}