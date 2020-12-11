package util

import (
	repo5 "github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	repo6 "github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
	repo7 "github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"
	repo8 "github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_repo"
	repo11 "github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_repo"
	repo9 "github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"
	repo10 "github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
	repo3 "github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_repo"
	repo4 "github.com/ScoreTrak/ScoreTrak/pkg/team/team_repo"
	repo12 "github.com/ScoreTrak/ScoreTrak/pkg/user/user_repo"
)

type Store struct {
	Round        repo3.Repo
	Host         repo7.Repo
	HostGroup    repo8.Repo
	Service      repo2.Repo
	ServiceGroup service_group_repo.Repo
	Team         repo4.Repo
	Check        repo5.Repo
	Property     repo9.Repo
	Config       repo6.Repo
	Report       repo10.Repo
	Policy       repo11.Repo
	Users        repo12.Repo
}
