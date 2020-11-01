package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type teamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) repo.Repo {
	return &teamRepo{db}
}

func (t *teamRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := t.db.WithContext(ctx).Delete(&team.Team{}, "id = ?", id)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the team with id : %d", id)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (t *teamRepo) DeleteByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("you must specify the name of the team you are trying to update")
	}
	result := t.db.WithContext(ctx).Delete(&team.Team{}, "name = ?", name)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the team with name : %s", name)
		return errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (t *teamRepo) GetAll(ctx context.Context) ([]*team.Team, error) {
	teams := make([]*team.Team, 0)
	err := t.db.WithContext(ctx).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *teamRepo) GetByID(ctx context.Context, id uuid.UUID) (*team.Team, error) {
	tea := &team.Team{}
	err := t.db.WithContext(ctx).Where("id = ?", id).First(tea).Error
	if err != nil {
		return nil, err
	}
	return tea, nil
}

func (t *teamRepo) GetByName(ctx context.Context, name string) (*team.Team, error) {
	if name == "" {
		return nil, errors.New("you must specify the name of the team you are trying to update")
	}
	tea := &team.Team{}
	err := t.db.WithContext(ctx).Where("name = ?", name).First(tea).Error
	if err != nil {
		return nil, err
	}
	return tea, nil
}

func (t *teamRepo) Store(ctx context.Context, tm []*team.Team) error {
	err := t.db.WithContext(ctx).Create(tm).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *teamRepo) Upsert(ctx context.Context, usr []*team.Team) error {
	err := t.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(usr).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *teamRepo) Update(ctx context.Context, tm *team.Team) error {
	err := t.db.WithContext(ctx).Model(tm).Updates(team.Team{Enabled: tm.Enabled, Name: tm.Name}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *teamRepo) UpdateByName(ctx context.Context, tm *team.Team) error {
	if tm.Name == "" {
		return errors.New("you must specify the name of the team you are trying to update")
	}
	err := t.db.WithContext(ctx).Model(tm).Where("name = ?", tm.Name).Updates(team.Team{Enabled: tm.Enabled}).Error
	if err != nil {
		return err
	}
	return nil
}
