package orm

import (
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
	"gorm.io/gorm"
)

type teamRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewTeamRepo(db *gorm.DB, log logger.LogInfoFormat) team.Repo {
	return &teamRepo{db, log}
}

func (t *teamRepo) Delete(id uint64) error {
	t.log.Debugf("deleting the team with id : %d", id)

	result := t.db.Delete(&team.Team{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the team with id : %d", id)
		t.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (t *teamRepo) DeleteByName(name string) error {
	t.log.Debugf("deleting the team with name : %s", name)
	if name == "" {
		return errors.New("you must specify the name of the team you are trying to update")
	}
	result := t.db.Delete(&team.Team{}, "name = ?", name)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the team with name : %s", name)
		t.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
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

func (t *teamRepo) GetByID(id uint64) (*team.Team, error) {
	t.log.Debugf("get team details by id : %s", id)

	tea := &team.Team{}
	err := t.db.Where("id = ?", id).First(tea).Error
	if err != nil {
		t.log.Errorf("team not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return tea, nil
}

func (t *teamRepo) GetByName(name string) (*team.Team, error) {
	t.log.Debugf("get team details by name : %s", name)
	if name == "" {
		return nil, errors.New("you must specify the name of the team you are trying to update")
	}
	tea := &team.Team{}
	err := t.db.Where("name = ?", name).First(tea).Error
	if err != nil {
		t.log.Errorf("team not found with id : %d, reason : %v", name, err)
		return nil, err
	}
	return tea, nil
}

func (t *teamRepo) Store(tm *team.Team) error {
	t.log.Debugf("creating the team with id : %v", tm.ID)
	err := t.db.Create(tm).Error
	if err != nil {
		t.log.Errorf("error while creating the team, reason : %v", err)
		return err
	}
	return nil
}

func (t *teamRepo) Update(tm *team.Team) error {
	t.log.Debugf("updating the team, with id : %v", tm.ID)
	err := t.db.Model(tm).Updates(team.Team{Enabled: tm.Enabled, Name: tm.Name}).Error
	if err != nil {
		t.log.Errorf("error while updating the team, reason : %v", err)
		return err
	}
	return nil
}

func (t *teamRepo) UpdateByName(tm *team.Team) error {
	t.log.Debugf("updating the team, with id : %v", tm.ID)
	if tm.Name == "" {
		return errors.New("you must specify the name of the team you are trying to update")
	}
	err := t.db.Model(tm).Where("name = ?", tm.Name).Updates(team.Team{Enabled: tm.Enabled}).Error
	if err != nil {
		t.log.Errorf("error while updating the team, reason : %v", err)
		return err
	}
	return nil
}
