package util

import (
	repo5 "github.com/ScoreTrak/ScoreTrak/pkg/check/repo"
	repo6 "github.com/ScoreTrak/ScoreTrak/pkg/config/repo"
	repo7 "github.com/ScoreTrak/ScoreTrak/pkg/host/repo"
	repo8 "github.com/ScoreTrak/ScoreTrak/pkg/host_group/repo"
	repo11 "github.com/ScoreTrak/ScoreTrak/pkg/policy/repo"
	repo9 "github.com/ScoreTrak/ScoreTrak/pkg/property/repo"
	repo10 "github.com/ScoreTrak/ScoreTrak/pkg/report/repo"
	repo3 "github.com/ScoreTrak/ScoreTrak/pkg/round/repo"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/service/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/repo"
	repo4 "github.com/ScoreTrak/ScoreTrak/pkg/team/repo"
	repo12 "github.com/ScoreTrak/ScoreTrak/pkg/user/repo"
)

type Store struct {
	Round        repo3.Repo
	Host         repo7.Repo
	HostGroup    repo8.Repo
	Service      repo2.Repo
	ServiceGroup repo.Repo
	Team         repo4.Repo
	Check        repo5.Repo
	Property     repo9.Repo
	Config       repo6.Repo
	Report       repo10.Repo
	Policy       repo11.Repo
	Users        repo12.Repo
}
