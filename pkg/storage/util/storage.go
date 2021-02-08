package util

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/team_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_repo"
)

type Store struct {
	Round        round_repo.Repo
	Host         host_repo.Repo
	HostGroup    host_group_repo.Repo
	Service      service_repo.Repo
	ServiceGroup service_group_repo.Repo
	Team         team_repo.Repo
	Check        check_repo.Repo
	Property     property_repo.Repo
	Config       config_repo.Repo
	Report       report_repo.Repo
	Policy       policy_repo.Repo
	Users        user_repo.Repo
}
