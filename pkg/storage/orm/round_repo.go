package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type roundRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewRoundRepo(db *gorm.DB, log logger.LogInfoFormat) round.Repo {
	return &roundRepo{db, log}
}

func (t *roundRepo) Delete(id uint64) error {
	t.log.Debugf("deleting the round with id : %s", id)

	if t.db.Delete(&round.Round{}, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the round with id : %s", id)
		t.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (t *roundRepo) GetAll() ([]*round.Round, error) {
	t.log.Debug("get all the rounds")
	rounds := make([]*round.Round, 0)
	err := t.db.Find(&rounds).Error
	if err != nil {
		t.log.Debug("not a single round found")
		return nil, err
	}
	return rounds, nil
}

func (t *roundRepo) GetByID(id uint64) (*round.Round, error) {
	t.log.Debugf("get round details by id : %s", id)

	tea := &round.Round{}
	err := t.db.Where("id = ?", id).First(&tea).Error
	if err != nil {
		t.log.Errorf("round not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return tea, nil
}

func (t *roundRepo) Store(tm *round.Round) error {
	t.log.Debugf("creating the round with id : %v", tm.ID)
	err := t.db.Create(&tm).Error
	if err != nil {
		t.log.Errorf("error while creating the round, reason : %v", err)
		return err
	}
	return nil
}

func (t *roundRepo) Update(tm *round.Round) error {
	t.log.Debugf("updating the round, with id : %v", tm.ID)
	err := t.db.Model(&tm).Updates(round.Round{End: tm.End}).Error
	if err != nil {
		t.log.Errorf("error while updating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) GetLastRound() (*round.Round, error) {
	rnd := &round.Round{}
	err := r.db.Where("\"end\" IS NOT NULL").Last(&rnd).Error
	//r.db.Raw("SELECT * FROM rounds WHERE end is NOT NULL order by id desc limit 1").Scan(&rnd).Error
	if err != nil {
		r.log.Debug("not a single Round found")
		return nil, err
	}
	return rnd, nil
}
