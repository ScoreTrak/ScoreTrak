package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type roundRepo struct {
	db *gorm.DB
}

func NewRoundRepo(db *gorm.DB) round_repo.Repo {
	return &roundRepo{db}
}

func (r *roundRepo) Delete(ctx context.Context, id uint64) error {
	result := r.db.WithContext(ctx).Delete(&round.Round{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the round with id : %d", id)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil

}

func (r *roundRepo) GetAll(ctx context.Context) ([]*round.Round, error) {
	rounds := make([]*round.Round, 0)
	err := r.db.WithContext(ctx).Find(&rounds).Error
	if err != nil {
		return nil, err
	}
	return rounds, nil
}

func (r *roundRepo) GetByID(ctx context.Context, id uint64) (*round.Round, error) {
	tea := &round.Round{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(tea).Error
	if err != nil {
		return nil, err
	}
	return tea, nil
}

func (r *roundRepo) Store(ctx context.Context, rn *round.Round) error {
	if rn.ID == 0 {
		return errors.New("the ID should be provided")
	}
	err := r.db.WithContext(ctx).Create(rn).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roundRepo) Upsert(ctx context.Context, rn []*round.Round) error {
	err := r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(rn).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roundRepo) StoreMany(ctx context.Context, rn []*round.Round) error {
	err := r.db.WithContext(ctx).Create(rn).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roundRepo) Update(ctx context.Context, rn *round.Round) error {
	err := r.db.WithContext(ctx).Model(rn).Updates(round.Round{Finish: rn.Finish, Note: rn.Note, Err: rn.Err}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roundRepo) GetLastNonElapsingRound(ctx context.Context) (*round.Round, error) {
	rnd := &round.Round{}
	err := r.db.WithContext(ctx).Where("\"finish\" IS NOT NULL").Last(rnd).Error
	if err != nil {
		return nil, err
	}
	return rnd, nil
}

func (r *roundRepo) GetLastElapsingRound(ctx context.Context) (*round.Round, error) {
	rnd, err := r.GetLastRound(ctx)
	if err != nil {
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
