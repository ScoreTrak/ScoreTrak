package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type roundRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewRoundRepo(db *gorm.DB, log logger.LogInfoFormat) round.Repo {
	return &roundRepo{db, log}
}

func (r *roundRepo) Delete(ctx context.Context, id uint) error {
	r.log.Debugf("deleting the round with id : %d", id)
	result := r.db.WithContext(ctx).Delete(&round.Round{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the round with id : %d", id)
		r.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil

}

func (r *roundRepo) GetAll(ctx context.Context) ([]*round.Round, error) {
	r.log.Debug("get all the rounds")
	rounds := make([]*round.Round, 0)
	err := r.db.WithContext(ctx).Find(&rounds).Error
	if err != nil {
		r.log.Debug("not a single round found")
		return nil, err
	}
	return rounds, nil
}

func (r *roundRepo) GetByID(ctx context.Context, id uint) (*round.Round, error) {
	r.log.Debugf("get round details by id : %s", id)

	tea := &round.Round{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(tea).Error
	if err != nil {
		r.log.Errorf("round not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return tea, nil
}

func (r *roundRepo) Store(ctx context.Context, rn *round.Round) error {
	if rn.ID == 0 {
		return errors.New("the ID should be provided")
	}
	r.log.Debugf("creating the round with id : %v", rn.ID)
	err := r.db.WithContext(ctx).Create(rn).Error
	if err != nil {
		r.log.Errorf("error while creating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) Upsert(ctx context.Context, rn []*round.Round) error {
	err := r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(rn).Error
	if err != nil {
		r.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) StoreMany(ctx context.Context, rn []*round.Round) error {
	err := r.db.WithContext(ctx).Create(rn).Error
	if err != nil {
		r.log.Errorf("error while creating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) Update(ctx context.Context, rn *round.Round) error {
	r.log.Debugf("updating the round, with id : %v", rn.ID)
	err := r.db.WithContext(ctx).Model(rn).Updates(round.Round{Finish: rn.Finish, Note: rn.Note, Err: rn.Err}).Error
	if err != nil {
		r.log.Errorf("error while updating the round, reason : %v", err)
		return err
	}
	return nil
}

func (r *roundRepo) GetLastNonElapsingRound(ctx context.Context) (*round.Round, error) {
	rnd := &round.Round{}
	err := r.db.WithContext(ctx).Where("\"finish\" IS NOT NULL").Last(rnd).Error
	if err != nil {
		r.log.Debug("not a single Round found")
		return nil, err
	}
	return rnd, nil
}

func (r *roundRepo) GetLastElapsingRound(ctx context.Context) (*round.Round, error) {
	rnd, err := r.GetLastRound(ctx)
	if err != nil {
		r.log.Debug("not a single Round found")
		return nil, err
	}
	if rnd.Finish == nil {
		return rnd, nil
	} else {
		return nil, errors.New("there is no round executing at the moment")
	}
}

func (r *roundRepo) GetLastRound(ctx context.Context) (*round.Round, error) {
	rnd := &round.Round{}
	err := r.db.WithContext(ctx).Last(rnd).Error
	if err != nil {
		r.log.Debug("not a single Round found")
		return nil, err
	}
	return rnd, nil
}

func (r *roundRepo) TruncateTable(ctx context.Context) (err error) {
	err = util.TruncateTable(ctx, &round.Round{}, r.db)
	if err != nil {
		return err
	}
	return nil
}
