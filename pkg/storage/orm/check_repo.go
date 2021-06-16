package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type checkRepo struct {
	db *gorm.DB
}

func NewCheckRepo(db *gorm.DB) check_repo.Repo {
	return &checkRepo{db}
}

func (c *checkRepo) GetAllByRoundID(ctx context.Context, roundID uint64) ([]*check.Check, error) {
	var checks []*check.Check
	err := c.db.WithContext(ctx).Where("round_id = ?", roundID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*check.Check, error) {
	var checks []*check.Check
	err := c.db.WithContext(ctx).Where("service_id = ?", serviceID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetByRoundServiceID(ctx context.Context, roundID uint64, serviceID uuid.UUID) (*check.Check, error) {
	chk := &check.Check{}
	err := c.db.WithContext(ctx).Where("round_id = ?", roundID).Where("service_id = ?", serviceID).First(&chk).Error
	if err != nil {
		return nil, err
	}
	return chk, err
}

func (c *checkRepo) Delete(ctx context.Context, roundID uint64, serviceID uuid.UUID) error {
	result := c.db.WithContext(ctx).Where("round_id = ?", roundID).Where("service_id = ?", serviceID).Delete(&check.Check{})
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the check with rid, sid : %d, %d", roundID, serviceID)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for round ID, and check_service id provided"}
	}

	return nil
}

func (c *checkRepo) GetAll(ctx context.Context) ([]*check.Check, error) {
	checks := make([]*check.Check, 0)
	err := c.db.WithContext(ctx).Find(&checks).Error
	if err != nil {
		return nil, err
	}
	return checks, nil
}

func (c *checkRepo) GetByID(ctx context.Context, roundID uint64, serviceID uuid.UUID) (*check.Check, error) {
	chck := &check.Check{}
	err := c.db.WithContext(ctx).Where("round_id = ? AND service_id = ?", roundID, serviceID).First(&chck).Error
	if err != nil {
		return nil, err
	}
	return chck, nil
}

func (c *checkRepo) Store(ctx context.Context, chck []*check.Check) error {
	ReplaceInvalidCharacters(chck)
	err := c.db.WithContext(ctx).Create(chck).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *checkRepo) Upsert(ctx context.Context, chck []*check.Check) error {
	ReplaceInvalidCharacters(chck)
	err := c.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(chck).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *checkRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &check.Check{}, c.db)
	if err != nil {
		return err
	}
	return nil
}

func ReplaceInvalidCharacters(chck []*check.Check) {
	for i := range chck {
		chck[i].Log = strings.ToValidUTF8(chck[i].Log, "�")
		chck[i].Err = strings.ToValidUTF8(chck[i].Err, "�")
	}
}
