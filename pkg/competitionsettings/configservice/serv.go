package competitionsettingsservice

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/competitionsettings"
	"github.com/ScoreTrak/ScoreTrak/pkg/competitionsettings/competitionsettingsrepo"
)

type Serv interface {
	Get(ctx context.Context) (*competitionsettings.CompetitionSettings, error)
	Update(context.Context, *competitionsettings.CompetitionSettings) error
}

type competitionSettingServ struct {
	repo competitionsettingsrepo.Repo
}

func NewCompetitionSettingServ(repo competitionsettingsrepo.Repo) Serv {
	return &competitionSettingServ{
		repo: repo,
	}
}

func (svc *competitionSettingServ) Get(ctx context.Context) (*competitionsettings.CompetitionSettings, error) {
	return svc.repo.Get(ctx)
}

func (svc *competitionSettingServ) Update(ctx context.Context, cs *competitionsettings.CompetitionSettings) error {
	return svc.repo.Update(ctx, cs)
}
