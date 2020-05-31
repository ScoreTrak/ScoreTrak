package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/team"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type teamRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewTeamRepo(db *gorm.DB, log logger.LogInfoFormat) team.Repo {
	return &teamRepo{db, log}
}

func (t *teamRepo) Delete(id string) error {
	t.log.Debugf("deleting the team with id : %d", id)

	result := t.db.Delete(&team.Team{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the team with id : %s", id)
		t.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for id"}
	}
	return nil
}

func (t *teamRepo) GetAll() ([]*team.Team, error) {
	t.log.Debug("get all the teams")
	teams := make([]*team.Team, 0)
	err := t.db.Find(&teams).Error
	if err != nil {
		t.log.Debug("not a single team found")
		return nil, err
	}
	return teams, nil
}

func (t *teamRepo) GetByID(id string) (*team.Team, error) {
	t.log.Debugf("get team details by id : %s", id)

	tea := &team.Team{}
	err := t.db.Where("id = ?", id).First(&tea).Error
	if err != nil {
		t.log.Errorf("team not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return tea, nil
}

func (t *teamRepo) Store(tm *team.Team) error {
	t.log.Debugf("creating the team with id : %v", tm.ID)
	err := t.db.Create(&tm).Error
	if err != nil {
		t.log.Errorf("error while creating the team, reason : %v", err)
		return err
	}
	return nil
}

func (t *teamRepo) Update(tm *team.Team) error {
	t.log.Debugf("updating the team, with id : %v", tm.ID)
	err := t.db.Model(&tm).Updates(team.Team{Enabled: tm.Enabled}).Error
	if err != nil {
		t.log.Errorf("error while updating the team, reason : %v", err)
		return err
	}
	return nil
}
