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

func (r *roundRepo) Delete(id uint64) error {
	r.log.Debugf("deleting the round with id : %d", id)
	result := r.db.Delete(&round.Round{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the round with id : %d", id)
		r.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for id"}
	}

	return nil

}

func (r *roundRepo) GetAll() ([]*round.Round, error) {
	r.log.Debug("get all the rounds")
	rounds := make([]*round.Round, 0)
	err := r.db.Find(&rounds).Error
	if err != nil {
		r.log.Debug("not a single round found")
		return nil, err
	}
	return rounds, nil
}

func (r *roundRepo) GetByID(id uint64) (*round.Round, error) {
	r.log.Debugf("get round details by id : %s", id)

	tea := &round.Round{}
	err := r.db.Where("id = ?", id).First(&tea).Error
	if err != nil {
		r.log.Errorf("round not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return tea, nil
}

func (r *roundRepo) Store(tm *round.Round) error {
	if tm.ID == 0 {
		return errors.New("the ID should be provided")
	}
	r.log.Debugf("creating the round with id : %v", tm.ID)
	err := r.db.Create(&tm).Error
	if err != nil {
		r.log.Errorf("error while creating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) Update(tm *round.Round) error {
	r.log.Debugf("updating the round, with id : %v", tm.ID)
	err := r.db.Model(&tm).Updates(round.Round{Finish: tm.Finish}).Error
	if err != nil {
		r.log.Errorf("error while updating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) GetLastRound() (*round.Round, error) {
	rnd := &round.Round{}
	err := r.db.Where("\"finish\" IS NOT NULL").Last(&rnd).Error
	//r.db.Raw("SELECT * FROM rounds WHERE end is NOT NULL order by id desc limit 1").Scan(&rnd).Error
	if err != nil {
		r.log.Debug("not a single Round found")
		return nil, err
	}
	return rnd, nil
}
