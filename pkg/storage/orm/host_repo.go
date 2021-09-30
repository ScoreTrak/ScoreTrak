package orm

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type hostRepo struct {
	db *gorm.DB
}

func NewHostRepo(db *gorm.DB) host_repo.Repo {
	return &hostRepo{db}
}

func (h *hostRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).Delete(&host.Host{}, "id = ?", id)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the host with id : %d", id)
		return errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (h *hostRepo) GetAll(ctx context.Context) ([]*host.Host, error) {
	hosts := make([]*host.Host, 0)
	err := h.db.WithContext(ctx).Find(&hosts).Error
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (h *hostRepo) GetByID(ctx context.Context, id uuid.UUID) (*host.Host, error) {
	hst := &host.Host{}
	err := h.db.WithContext(ctx).Where("id = ?", id).First(hst).Error
	if err != nil {
		return nil, err
	}
	return hst, nil
}

func (h *hostRepo) Store(ctx context.Context, hst []*host.Host) error {
	err := h.db.WithContext(ctx).Create(hst).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostRepo) Upsert(ctx context.Context, hst []*host.Host) error {
	err := h.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(hst).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostRepo) Update(ctx context.Context, hst *host.Host) error {
	err := h.db.WithContext(ctx).Model(hst).Updates(host.Host{Pause: hst.Pause, Hide: hst.Hide,
		Address: hst.Address, HostGroupID: hst.HostGroupID,
		TeamID: hst.TeamID, EditHost: hst.EditHost, AddressListRange: hst.AddressListRange,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *hostRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &host.Host{}, h.db)
	if err != nil {
		return err
	}
	return nil
}
