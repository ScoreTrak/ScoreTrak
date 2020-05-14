package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"github.com/jinzhu/gorm"
)

type roundRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewRoundRepo(db *gorm.DB, log logger.LogInfoFormat) round.Repo {
	return &roundRepo{db, log}
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
